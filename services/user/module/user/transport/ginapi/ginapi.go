package ginapi

import (
	"github.com/caovanhoang63/hiholive/services/user/module/user/biz"
	"github.com/caovanhoang63/hiholive/services/user/module/user/usermodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginAPI struct {
	biz        biz.UserBiz
	serviceCtx srvctx.ServiceContext
}

func NewGinAPI(serviceCtx srvctx.ServiceContext, b biz.UserBiz) *ginAPI {
	return &ginAPI{b, serviceCtx}
}

func (g *ginAPI) GetUserById() func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("id"))

		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		user, err := g.biz.FindUserById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}
		user.Mask()
		c.JSON(http.StatusOK, core.ResponseData(user))
	}
}

func (g *ginAPI) GetUserProfile() func(c *gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(core.KeyRequester).(core.Requester)

		user, err := g.biz.FindUserById(c.Request.Context(), requester.GetUserId())
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}
		user.Mask()
		c.JSON(http.StatusOK, core.ResponseData(user))
	}
}

func (g *ginAPI) ListUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging core.Paging
		var filter usermodel.UserFilter

		if err := c.ShouldBindQuery(&filter); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
			return
		}
		if err := c.ShouldBindQuery(&paging); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
			return
		}

		paging.Process()

		users, err := g.biz.FindUsersWithCondition(c.Request.Context(), &filter, &paging)
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		for i := range users {
			users[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(users, paging, filter))
	}
}
