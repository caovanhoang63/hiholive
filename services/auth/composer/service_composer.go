package composer

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/biz"
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

	authRepo := mysql.NewMySQLRepository(db.GetDB())
	userBiz := biz.NewAuthBiz(serviceCtx, authRepo)

	userService := ginapi.NewGinAPI(serviceCtx, userBiz)
	return userService
}
