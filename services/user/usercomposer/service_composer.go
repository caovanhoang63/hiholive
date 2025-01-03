package usercomposer

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/services/user/module/user/repository/mysql"
	"github.com/caovanhoang63/hiholive/services/user/module/user/transport/ginapi"
	"github.com/caovanhoang63/hiholive/services/user/module/user/transport/grpc"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserById() func(c *gin.Context)
	ListUser() func(c *gin.Context)
	GetUserProfile() func(c *gin.Context)
	UpdateUserData() func(c *gin.Context)
}

func ComposeUserAPIService(serviceCtx srvctx.ServiceContext) UserService {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	userRepo := mysql.NewMySQLRepository(db.GetDB())
	userBiz := biz.NewBiz(userRepo, ps)

	userService := ginapi.NewGinAPI(serviceCtx, userBiz)
	return userService
}

func ComposeUserGRPCService(serviceCtx srvctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(core.KeyCompMySQL).(core.GormComponent)
	ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)
	userRepo := mysql.NewMySQLRepository(db.GetDB())
	userBiz := biz.NewBiz(userRepo, ps)
	userService := grpc.NewService(userBiz)

	return userService
}
