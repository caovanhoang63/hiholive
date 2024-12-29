package middlewares

import (
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type UserClient interface {
	GetUserRole(ctx context.Context, userId int) (string, error)
}

func Authorize(uc UserClient, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		role, err := uc.GetUserRole(c.Request.Context(), requester.GetUserId())
		if err != nil {

			core.WriteErrorResponse(c, err)
			c.Abort()
			return
		}

		for _, r := range roles {
			if role == r {
				requester.SetRole(role)
				c.Set(core.KeyRequester, requester)
				c.Next()
				return
			}
		}

		core.WriteErrorResponse(c, core.ErrForbidden)
		c.Abort()
	}
}
