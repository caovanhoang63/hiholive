package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteErrorResponse(c *gin.Context, err error) {
	c.Error(err)
	if errSt, ok := err.(StatusCodeCarrier); ok {
		c.JSON(errSt.StatusCode(), errSt)
	} else {
		c.JSON(http.StatusInternalServerError, ErrInternalServerError.WithError(err.Error()))
	}
}
