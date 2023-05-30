package ginUpload

import (
	"github.com/gin-gonic/gin"
	"social-photo/common"
	"social-photo/component/uploadprovider"
	"social-photo/modules/upload/biz"
)

func Upload(uploadProvider uploadprovider.UploadProvider) func(*gin.Context) {
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

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//imgStore := storage.NewSQLStore(db)
		business := biz.NewUploadBiz(uploadProvider, nil)
		img, err := business.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
