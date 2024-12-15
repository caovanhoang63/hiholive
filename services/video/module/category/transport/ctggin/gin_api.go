package ctggin

import (
	"github.com/caovanhoang63/hiholive/services/video/module/category/ctgbiz"
	"github.com/caovanhoang63/hiholive/services/video/module/category/ctgmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/gin-gonic/gin"
)

type categoryApi struct {
	biz ctgbiz.CategoryBiz
}

func NewCategoryApi(biz ctgbiz.CategoryBiz) *categoryApi {
	return &categoryApi{biz: biz}
}

func (a *categoryApi) CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var create ctgmodel.CategoryCreate

		if err := c.ShouldBindJSON(&create); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)

		if err := a.biz.CreateCategory(c.Request.Context(), requester, &create); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		create.Mask(core.DbTypeCategory)
		c.JSON(200, core.ResponseData(&create.Uid))
	}
}

func (a *categoryApi) UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var update ctgmodel.CategoryUpdate

		uid, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err = c.ShouldBindJSON(&update); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)

		if err = a.biz.UpdateCategory(c.Request.Context(), requester, int(uid.GetLocalID()), &update); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		c.JSON(200, core.ResponseData(true))

	}
}

func (a *categoryApi) DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)

		if err = a.biz.DeleteCategory(c.Request.Context(), requester, int(uid.GetLocalID())); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(200, core.ResponseData(true))
	}
}

func (a *categoryApi) FindCategories() gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter ctgmodel.CategoryFilter
		var paging core.Paging

		if err := c.ShouldBindQuery(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithWrap(err))
			return
		}

		if err := c.ShouldBindQuery(paging); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithWrap(err))
			return
		}

		paging.Process()

		data, err := a.biz.FindCategories(c.Request.Context(), &filter, &paging)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		for i := range data {
			data[i].Mask()
		}
		c.JSON(200, core.SuccessResponse(data, paging, filter))

	}
}

func (a *categoryApi) FindCategoryById() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}
		data, err := a.biz.FindCategory(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		data.Mask()
		c.JSON(200, core.ResponseData(data))
	}
}
