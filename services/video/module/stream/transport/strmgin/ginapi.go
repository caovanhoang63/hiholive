package strmgin

import (
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streambiz"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streammodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginAPI struct {
	biz    streambiz.StreamBiz
	srvctx srvctx.ServiceContext
}

func NewStreamApi(biz streambiz.StreamBiz, srvctx srvctx.ServiceContext) *ginAPI {
	return &ginAPI{
		biz:    biz,
		srvctx: srvctx,
	}
}

func (g *ginAPI) CreateStream() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data streammodel.StreamCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)

		res, err := g.biz.Create(c.Request.Context(), requester, &data)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(res))

	}
}

func (g *ginAPI) GetStreamById() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		data, err := g.biz.FindStreamById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		data.Mask()

		c.JSON(http.StatusOK, core.ResponseData(data))
	}

}
