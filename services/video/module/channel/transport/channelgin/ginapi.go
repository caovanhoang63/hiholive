package channelgin

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelmodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
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

		data.Mask(core.DbTypeChannel)

		c.JSON(http.StatusOK, core.ResponseData(&data.Uid))
	}
}

func (g *ginAPI) FindUserChannel() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		uid, err := core.FromBase58(c.Param("id"))

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		channel, err := g.biz.FindUserChannel(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		channel.Mask()

		c.JSON(http.StatusOK, core.ResponseData(channel))
	}

}

func (g *ginAPI) FindChannelById() func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		channel, err := g.biz.FindChannelById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		channel.Mask()

		c.JSON(http.StatusOK, core.ResponseData(channel))
	}
}

func (g *ginAPI) FindChannels() func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging core.Paging
		var filter channelmodel.ChannelFilter

		if err := c.ShouldBindQuery(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
			return
		}
		if err := c.ShouldBindQuery(&paging); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
			return
		}

		paging.Process()

		channels, err := g.biz.FindChannels(c.Request.Context(), &filter, &paging)

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		for i := 0; i < len(channels); i++ {
			channels[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(channels, paging, filter))
	}
}
