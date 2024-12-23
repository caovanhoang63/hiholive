package usercomposer

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/services/user/module/user/repository/mysql"
	"github.com/caovanhoang63/hiholive/services/user/module/user/transport/usersub"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
)

func ComposeUserSubscriber(serviceCtx srvctx.ServiceContext) *usersub.UserSub {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)

	userRepo := mysql.NewMySQLRepository(db.GetDB())
	userBiz := biz.NewBiz(userRepo)

	userSub := usersub.NewUserSub(userBiz, serviceCtx)
	return userSub
}
