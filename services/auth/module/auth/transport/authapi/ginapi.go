package authapi

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authbiz"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginAPI struct {
	biz        authbiz.AuthBiz
	serviceCtx srvctx.ServiceContext
}

func NewGinAPI(serviceCtx srvctx.ServiceContext, b authbiz.AuthBiz) *ginAPI {
	return &ginAPI{b, serviceCtx}
}

func (g *ginAPI) Register() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data authmodel.AuthRegister

		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		err := g.biz.Register(c.Request.Context(), &data)

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))

	}
}
