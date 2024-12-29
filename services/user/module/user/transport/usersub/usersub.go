package usersub

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/golang/subengine"
	"golang.org/x/net/context"
)

type UserSub struct {
	biz        biz.UserBiz
	serviceCtx srvctx.ServiceContext
}

func NewUserSub(biz biz.UserBiz, serviceCtx srvctx.ServiceContext) *UserSub {
	return &UserSub{
		biz:        biz,
		serviceCtx: serviceCtx,
	}
}

func (u *UserSub) UpdateUserToStreamer() subengine.ConsumerJob {
	return subengine.ConsumerJob{
		Title: "Update user to a streamer",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			id, ok := message.Data.(float64)
			if ok {
				err := u.biz.UpdateToRoleStreamer(ctx, int(id))
				if err != nil {

					return err
				}
				return nil
			}
			return core.ErrInternalServerError
		},
	}
}
