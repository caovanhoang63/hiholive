package usersub

import (
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/go/subengine"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func (u *UserSub) TestHandler() subengine.ConsumerJob {
	return subengine.ConsumerJob{
		Title: "Test",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			log.Print("test")
			return nil
		},
	}
}
