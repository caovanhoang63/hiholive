package ginapi

import (
	"github.com/caovanhoang63/hiholive/projects/go/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
)

type ginAPI struct {
	biz        biz.UserBiz
	serviceCtx srvctx.ServiceContext
}

func (g ginAPI) GetUserProfile() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func NewGinAPI(serviceCtx srvctx.ServiceContext, b biz.UserBiz) *ginAPI {
	return &ginAPI{b, serviceCtx}
}
