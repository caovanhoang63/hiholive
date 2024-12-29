package cmd

import (
	"github.com/caovanhoang63/hiholive/services/user/usercomposer"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/golang/subengine"
	log "github.com/sirupsen/logrus"
)

func StartSubscriber(serviceCtx srvctx.ServiceContext) {
	service := usercomposer.ComposeUserSubscriber(serviceCtx)

	pb := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)

	engine := subengine.NewEngine(serviceCtx, pb)

	engine.Subscribe("Test", service.TestHandler())
	engine.Subscribe(core.TopicCreateChannel, service.UpdateUserToStreamer())
	go func() {
		err := engine.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
