package middlewares

import (
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"strings"
)

type AuthClient interface {
	IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error)
}

func RequireAuth(ac AuthClient) func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			core.WriteErrorResponse(c, err)
			c.Abort()
			return
		}

		sub, tid, err := ac.IntrospectToken(c.Request.Context(), token)

		if err != nil {
			core.WriteErrorResponse(c, core.ErrUnauthorized.WithDebug(err.Error()))
			c.Abort()
			return
		}
		c.Set(core.KeyRequester, core.NewRequester(sub, tid))
		c.Next()
	}
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" || strings.TrimSpace(parts[1]) == "null" {
		return "", core.ErrUnauthorized.WithError("missing access token")
	}

	return parts[1], nil
}
