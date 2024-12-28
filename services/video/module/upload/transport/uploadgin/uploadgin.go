package uploadgin

import (
	"github.com/caovanhoang63/hiholive/services/video/module/upload/uploadbiz"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UploadGin struct {
	biz        uploadbiz.UploadBiz
	serviceCtx srvctx.ServiceContext
}

func NewUploadGin(biz uploadbiz.UploadBiz, serviceCtx srvctx.ServiceContext) *UploadGin {
	return &UploadGin{
		biz:        biz,
		serviceCtx: serviceCtx,
	}
}

func (u *UploadGin) UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		defer file.Close()

		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		dataBytes := make([]byte, fileHeader.Size)

		if _, err = file.Read(dataBytes); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		img, err := u.biz.UploadImage(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(img))
	}
}
