package cmd

import (
	"github.com/caovanhoang63/hiholive/services/video/videocomposer"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/golang/subengine"
	log "github.com/sirupsen/logrus"
)

func StartSubscriber(serviceCtx srvctx.ServiceContext) {
	streamService := videocomposer.ComposeStreamSubscriber(serviceCtx)
	categoryService := videocomposer.ComposeCategorySubscriber(serviceCtx)
	channelService := videocomposer.ComposeChannelSubscriber(serviceCtx)
	pb := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)

	engine := subengine.NewEngine(serviceCtx, pb)

	engine.Subscribe(core.TopicStreamStart, streamService.StartStream())
	engine.Subscribe(core.TopicStreamCreate, categoryService.IncreaseTotalContent())
	engine.Subscribe(core.TopicUpdateStreamViewCount, streamService.UpdateStreamViewCount())
	engine.Subscribe(core.TopicStreamEnded, streamService.EndStream())
	engine.Subscribe(core.TopicUpdateUserName, channelService.UpdateChannelName())
	engine.Subscribe(core.TopicUpdateChannelImage, channelService.UpdateChannelImage())
	go func() {
		err := engine.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
