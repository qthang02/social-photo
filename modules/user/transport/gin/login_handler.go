package ginUser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-photo/common"
	"social-photo/component/tokenprovider/jwt"
	"social-photo/modules/user/biz"
	"social-photo/modules/user/model"
	"social-photo/modules/user/storage"
)

func Login(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.UserLogin

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		tokenProvider := jwt.NewTokenProvider("jwt", "secretKey")

		store := storage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		business := biz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)

		account, err := business.Login(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
