package composer

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authbiz"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/repository/authgrpcrepo"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/repository/authmysql"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/transport/authapi"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register() func(c *gin.Context)
	Login() func(c *gin.Context)
}

func ComposeAuthAPIService(serviceCtx srvctx.ServiceContext) AuthService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	jwtComp := serviceCtx.MustGet(core.KeyCompJWT).(core.JWTProvider)
	userClient := authgrpcrepo.NewClient(ComposeUserRPCClient(serviceCtx))
	authRepo := authmysql.NewMySQLRepository(db.GetDB())
	authBiz := authbiz.NewAuthBiz(serviceCtx, authRepo, userClient, jwtComp, core.NewSha256Hash())
	userService := authapi.NewGinAPI(serviceCtx, authBiz)
	return userService
}
