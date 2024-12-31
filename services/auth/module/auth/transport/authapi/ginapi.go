package authapi

import (
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authbiz"
	"github.com/caovanhoang63/hiholive/services/auth/module/auth/authmodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ginAPI struct {
	biz        authbiz.AuthBiz
	serviceCtx srvctx.ServiceContext
}

func (g *ginAPI) ForgotPassword() func(c *gin.Context) {
	return func(c *gin.Context) {
		type Email struct {
			Email string `json:"email"`
		}
		var email Email
		if err := c.ShouldBind(&email); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
			return
		}
		if err := g.biz.ForgotPassword(c.Request.Context(), email.Email); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (g *ginAPI) ResetPassword() func(c *gin.Context) {
	return func(c *gin.Context) {
		type Pin struct {
			Pin      string `json:"pin"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var pin Pin
		if err := c.ShouldBind(&pin); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
			return
		}

		if err := g.biz.ResetPasswordWithPin(c.Request.Context(), pin.Email, pin.Pin, pin.Password); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (g *ginAPI) CheckForgotPasswordPin() func(c *gin.Context) {
	return func(c *gin.Context) {
		type Pin struct {
			Pin   string `json:"pin"`
			Email string `json:"email"`
		}
		var pin Pin
		if err := c.ShouldBind(&pin); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest)
			return
		}
		if err := g.biz.CheckForgotPasswordPin(c.Request.Context(), pin.Email, pin.Pin); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(true))

	}
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
func (g *ginAPI) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data authmodel.AuthEmailPassword

		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		token, err := g.biz.Login(c.Request.Context(), &data)

		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(token))

	}
}
