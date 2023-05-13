package ginPost

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-photo/common"
	"social-photo/modules/post/biz"
	"social-photo/modules/post/storage"
	"strconv"
)

func DeletePostById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewDeletePostBiz(store)

		if err := business.DeletePostById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
