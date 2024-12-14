package core

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteErrorResponse(c *gin.Context, err error) {
	var defaultError *DefaultError
	if errors.As(err, &defaultError) && defaultError.err != nil {
		c.Error(defaultError.err)
	} else {
		c.Error(err)
	}
	if errSt, ok := err.(StatusCodeCarrier); ok {
		c.JSON(errSt.StatusCode(), errSt)
	} else {
		c.JSON(http.StatusInternalServerError, ErrInternalServerError.WithError(err.Error()))
	}
}
