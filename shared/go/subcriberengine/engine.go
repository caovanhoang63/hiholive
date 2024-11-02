package subcriberengine

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/go/asyncjob"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type consumerJob struct {
	Title   string
	Handler func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx srvctx.ServiceContext
	pb     pubsub.Pubsub
}

func NewEngine(appCtx srvctx.ServiceContext) *consumerEngine {
	return &consumerEngine{
		appCtx: appCtx,
	}
}

func (engine *consumerEngine) Start() error {

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic string, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.pb.Subscribe(context.Background(), topic)
	for _, item := range consumerJobs {
		log.Printf("Set up consumer for: %s", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Printf("running job for %s for. Value %s \n", job.Title, message.Data)
			return job.Handler(ctx, message)
		}
	}

	go func() {
		core.AppRecover()
		for {
			msg := <-c
			fmt.Println("Message: ", msg.Data)
			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}
			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println("Err:", err)
			}
		}
	}()

	return nil
}
