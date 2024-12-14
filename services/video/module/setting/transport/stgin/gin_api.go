package stgin

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/video/module/setting/stbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/setting/stmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/gin-gonic/gin"
)

type ginApi struct {
	biz stbiz.SystemSettingBiz
}

func NewGinApi(biz stbiz.SystemSettingBiz) *ginApi {
	return &ginApi{biz: biz}
}

func (api *ginApi) CreateSystemSetting() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data stmodel.SettingCreate
		if err := c.ShouldBindJSON(&data); err != nil {
			fmt.Println(err.Error())
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)

		if err := api.biz.Create(c.Request.Context(), requester, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(200, core.ResponseData(true))
	}
}

func (api *ginApi) UpdateSystemSetting() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data stmodel.SettingUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		name := c.Param("name")
		requester := c.MustGet(core.KeyRequester).(core.Requester)

		if err := api.biz.Update(c.Request.Context(), requester, name, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(200, core.ResponseData(true))
	}
}

func (api *ginApi) FindSystemSettingByName() gin.HandlerFunc {
	return func(c *gin.Context) {

		name := c.Param("name")

		data, err := api.biz.FindByName(c.Request.Context(), name)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask(core.DbTypeSystemSetting)

		c.JSON(200, core.ResponseData(data))

	}
}

func (api *ginApi) FindSystemSetting() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging core.Paging
		var filter stmodel.Filter

		if err := c.ShouldBindQuery(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithWrap(err))
			return
		}
		if err := c.ShouldBindQuery(paging); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithWrap(err))
			return
		}
		paging.Process()

		data, err := api.biz.FindByCondition(c.Request.Context(), &filter, &paging)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		for i := range data {
			data[i].Mask(core.DbTypeSystemSetting)
		}
		c.JSON(200, core.SuccessResponse(data, paging, filter))
	}
}
