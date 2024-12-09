package composer

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authbiz"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/repository/authgrpcrepo"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/repository/authmysql"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/transport/authapi"
	"github.com/caovanhoang63/hiholive/shared/go/shared"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register() func(c *gin.Context)
}

func ComposeAuthAPIService(serviceCtx srvctx.ServiceContext) AuthService {
	db := serviceCtx.MustGet(shared.KeyCompMySQL).(shared.GormComponent)

	userClient := authgrpcrepo.NewClient(ComposeUserRPCClient(serviceCtx))
	authRepo := authmysql.NewMySQLRepository(db.GetDB())
	authBiz := authbiz.NewAuthBiz(serviceCtx, authRepo, userClient)
	userService := authapi.NewGinAPI(serviceCtx, authBiz)
	return userService
}
