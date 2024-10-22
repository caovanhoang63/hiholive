package composer

import (
	"github.com/caovanhoang63/hiholive/shared/shared"
	"github.com/caovanhoang63/hiholive/shared/srvctx"
	"github.com/caovanhoang63/hiholive/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/user/module/user/repository/mysql"
	"github.com/caovanhoang63/hiholive/user/module/user/transport/ginapi"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserProfile() func(c *gin.Context)
}

func ComposeUserAPIService(serviceCtx srvctx.ServiceContext) UserService {
	db := serviceCtx.MustGet(shared.KeyCompMySQL).(shared.GormComponent)

	userRepo := mysql.NewMySQLRepository(db.GetDB())
	userBiz := biz.NewBiz(userRepo)

	userService := ginapi.NewGinAPI(serviceCtx, userBiz)
	return userService
}
