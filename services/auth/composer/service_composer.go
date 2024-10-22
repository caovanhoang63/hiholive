package composer

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/biz"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/repository/grpc"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/repository/mysql"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/transport/ginapi"
	"github.com/caovanhoang63/hiholive/shared/go/shared"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register() func(c *gin.Context)
}

func ComposeAuthAPIService(serviceCtx srvctx.ServiceContext) AuthService {
	db := serviceCtx.MustGet(shared.KeyCompMySQL).(shared.GormComponent)

	userClient := grpc.NewClient(ComposeUserRPCClient(serviceCtx))
	authRepo := mysql.NewMySQLRepository(db.GetDB())
	authBiz := biz.NewAuthBiz(serviceCtx, authRepo, userClient)
	userService := ginapi.NewGinAPI(serviceCtx, authBiz)
	return userService
}
