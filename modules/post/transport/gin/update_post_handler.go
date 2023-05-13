package ginPost

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-photo/common"
	"social-photo/modules/post/biz"
	"social-photo/modules/post/model"
	"social-photo/modules/post/storage"
	"strconv"
)

func UpdatePost(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadGateway, common.ErrInvalidRequest(err))

			return
		}

		var data model.PostUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadGateway, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdatePostBiz(store)

		if err := business.UpdatePostById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadGateway, err)

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
