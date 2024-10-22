package middlewares

import (
	"github.com/gin-gonic/gin"
	"hiholive/shared/go/srvctx"
	"time"
)

func Logger(serviceCtx srvctx.ServiceContext) gin.HandlerFunc {
	logger := serviceCtx.Logger("GIN")
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now().Sub(start)
		fields := srvctx.Field{
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"query":      c.Request.URL.RawQuery,
			"statusCode": c.Writer.Status(),
			"ip":         c.ClientIP(),
			"duration":   end.String(),
		}

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.WithSrc().WithField(fields).Error(e)
			}
		} else {
			logger.WithField(fields).Info()
		}
	}
}
