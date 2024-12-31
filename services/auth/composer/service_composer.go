package composer

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authbiz"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/repository/authgrpcrepo"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/repository/authmysql"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/transport/authapi"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/transport/authgrpc"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register() func(c *gin.Context)
	Login() func(c *gin.Context)
	ForgotPassword() func(c *gin.Context)
	ResetPassword() func(c *gin.Context)
	CheckForgotPasswordPin() func(c *gin.Context)
}

func ComposeAuthAPIService(serviceCtx srvctx.ServiceContext) AuthService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
	jwtComp := serviceCtx.MustGet(core.KeyCompJWT).(core.JWTProvider)
	userClient := authgrpcrepo.NewClient(ComposeUserRPCClient(serviceCtx))
	authRepo := authmysql.NewMySQLRepository(db.GetDB(), rd.GetClient())
	authBiz := authbiz.NewAuthBiz(serviceCtx, authRepo, userClient, jwtComp, core.NewSha256Hash(), ps)
	userService := authapi.NewGinAPI(serviceCtx, authBiz)
	return userService
}

func ComposeAuthGRPCService(serviceCtx srvctx.ServiceContext) pb.AuthServiceServer {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
	jwtComp := serviceCtx.MustGet(core.KeyCompJWT).(core.JWTProvider)
	userClient := authgrpcrepo.NewClient(ComposeUserRPCClient(serviceCtx))
	authRepo := authmysql.NewMySQLRepository(db.GetDB(), rd.GetClient())
	authBiz := authbiz.NewAuthBiz(serviceCtx, authRepo, userClient, jwtComp, core.NewSha256Hash(), ps)
	service := authgrpc.NewAuthGRPCService(authBiz)
	return service
}
