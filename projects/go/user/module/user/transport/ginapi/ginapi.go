package ginapi

import (
	"github.com/gin-gonic/gin"
	"hiholive/projects/go/user/module/user/biz"
	"hiholive/shared/go/srvctx"
)

type ginAPI struct {
	biz        biz.UserBiz
	serviceCtx srvctx.ServiceContext
}

func (g ginAPI) GetUserProfile() func(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewGinAPI(serviceCtx srvctx.ServiceContext, b biz.UserBiz) *ginAPI {
	return &ginAPI{b, serviceCtx}
}
