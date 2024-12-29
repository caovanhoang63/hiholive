package subengine

import (
	"github.com/caovanhoang63/hiholive/shared/golang/asyncjob"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"golang.org/x/net/context"
)

type ConsumerJob struct {
	Title   string
	Handler func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx srvctx.ServiceContext
	pb     pubsub.Pubsub
	jobs   map[string][]ConsumerJob
	logger srvctx.Logger
}

func NewEngine(appCtx srvctx.ServiceContext, pb pubsub.Pubsub) *consumerEngine {
	logger := appCtx.Logger("subengine")
	return &consumerEngine{
		jobs:   make(map[string][]ConsumerJob),
		pb:     pb,
		logger: logger,
		appCtx: appCtx,
	}
}

func (engine *consumerEngine) Start() error {
	for topic, jobs := range engine.jobs {
		err := engine.startSubTopic(topic, true, jobs...)
		if err != nil {
			engine.logger.Infof("Error starting sub topic %s: %v", topic, err)
			continue
		}
	}
	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) Subscribe(topic string, consumerJobs ...ConsumerJob) {
	if _, ok := engine.jobs[topic]; !ok {
		engine.jobs[topic] = make([]ConsumerJob, 0)
	}
	engine.jobs[topic] = append(engine.jobs[topic], consumerJobs...)
}

func (engine *consumerEngine) startSubTopic(topic string, isConcurrent bool, consumerJobs ...ConsumerJob) error {
	c, _ := engine.pb.Subscribe(context.Background(), topic)
	for _, item := range consumerJobs {
		engine.logger.Infof("Set up consumer for: %s", item.Title)
	}

	getJobHandler := func(job *ConsumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			engine.logger.Infof("running job for %s for. Value %s \n", job.Title, message.Data)
			return job.Handler(ctx, message)
		}
	}

	go func() {
		defer core.AppRecover()
		for {
			msg := <-c
			engine.logger.Infof("Message: ", msg.Data)
			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}
			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				engine.logger.Errorf("Err:", err)
			}
		}
	}()

	return nil
}
