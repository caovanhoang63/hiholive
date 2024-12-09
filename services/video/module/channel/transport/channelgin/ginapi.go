package channelgin

import (
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginAPI struct {
	biz        channelbiz.ChannelBiz
	serviceCtx srvctx.ServiceContext
}

func NewChannelGinApi(biz channelbiz.ChannelBiz, serviceCtx srvctx.ServiceContext) *ginAPI {
	return &ginAPI{
		biz:        biz,
		serviceCtx: serviceCtx,
	}
}

func (g *ginAPI) CreateChannel() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data channelmodel.ChannelCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)

		if err := g.biz.Create(c.Request.Context(), requester, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(true))

	}
}
