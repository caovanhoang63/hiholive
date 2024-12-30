package usercomposer

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/services/user/module/user/repository/mysql"
	"github.com/caovanhoang63/hiholive/services/user/module/user/transport/usersub"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
)

func ComposeUserSubscriber(serviceCtx srvctx.ServiceContext) *usersub.UserSub {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)

	userRepo := mysql.NewMySQLRepository(db.GetDB())
	userBiz := biz.NewBiz(userRepo, ps)

	userSub := usersub.NewUserSub(userBiz, serviceCtx)
	return userSub
}
