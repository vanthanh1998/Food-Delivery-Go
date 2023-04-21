package ginupload

import (
	"Food-Delivery/common"
	"Food-Delivery/component/appctx"
	"Food-Delivery/module/upload/uploadbusiness"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
)

func Upload(appCtx appctx.AppContext) func(*gin.Context) { // gin.HandlerFunc ~~ c *gin.Contex
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close hear

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbusiness.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
