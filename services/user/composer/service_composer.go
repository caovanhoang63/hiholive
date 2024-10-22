package composer

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/services/user/module/user/repository/mysql"
	"github.com/caovanhoang63/hiholive/services/user/module/user/transport/ginapi"
	"github.com/caovanhoang63/hiholive/services/user/module/user/transport/grpc"
	"github.com/caovanhoang63/hiholive/services/user/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/go/shared"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
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

func ComposeUserGRPCService(serviceCtx srvctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(shared.KeyCompMySQL).(shared.GormComponent)

	userRepo := mysql.NewMySQLRepository(db.GetDB())
	userBiz := biz.NewBiz(userRepo)
	userService := grpc.NewService(userBiz)

	return userService
}
