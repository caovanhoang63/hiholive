package cmd

import (
	"github.com/caovanhoang63/hiholive/services/video/strmcomposer"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/go/subengine"
	log "github.com/sirupsen/logrus"
)

func StartSubscriber(serviceCtx srvctx.ServiceContext) {
	service := strmcomposer.ComposeStreamSubscriber(serviceCtx)

	pb := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)

	engine := subengine.NewEngine(serviceCtx, pb)

	engine.Subscribe(core.TopicStreamStart, service.StartStream())

	go func() {
		err := engine.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
