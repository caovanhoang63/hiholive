package middlewares

import (
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type UserClient interface {
	GetUserRole(ctx context.Context, userId int) (string, error)
}

func Authorize(uc UserClient, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := core.GetRequester(c.Request.Context())
		role, err := uc.GetUserRole(c.Request.Context(), requester.GetUserId())
		if err != nil {
			core.WriteErrorResponse(c, err)
			c.Abort()
			return
		}

		for _, r := range roles {
			if role == r {
				c.Next()
				return
			}
		}

		core.WriteErrorResponse(c, core.ErrForbidden)
		c.Abort()
	}
}
