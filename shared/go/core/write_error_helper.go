package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteErrorResponse(c *gin.Context, err error) {
	if errSt, ok := err.(StatusCodeCarrier); ok {
		c.JSON(errSt.StatusCode(), errSt)
		return
	}

	c.JSON(http.StatusInternalServerError, ErrInternalServerError.WithError(err.Error()))
}
