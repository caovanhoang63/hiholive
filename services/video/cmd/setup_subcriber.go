package cmd

import (
	"github.com/caovanhoang63/hiholive/services/video/videocomposer"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/go/subengine"
	log "github.com/sirupsen/logrus"
)

func StartSubscriber(serviceCtx srvctx.ServiceContext) {
	streamService := videocomposer.ComposeStreamSubscriber(serviceCtx)
	categoryService := videocomposer.ComposeCategorySubscriber(serviceCtx)

	pb := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)

	engine := subengine.NewEngine(serviceCtx, pb)

	engine.Subscribe(core.TopicStreamStart, streamService.StartStream())
	engine.Subscribe(core.TopicStreamCreate, categoryService.IncreaseTotalContent())
	engine.Subscribe(core.TopicUpdateStreamViewCount, streamService.UpdateStreamViewCount())

	go func() {
		err := engine.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
