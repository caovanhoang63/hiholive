package composer

import (
	"github.com/gin-gonic/gin"
	"hiholive/projects/go/user/module/user/biz"
	"hiholive/projects/go/user/module/user/repository/mysql"
	"hiholive/projects/go/user/module/user/transport/ginapi"
	"hiholive/shared/go/shared"
	"hiholive/shared/go/srvctx"
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
