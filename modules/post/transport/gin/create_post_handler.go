package ginPost

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-photo/modules/post/biz"
	"social-photo/modules/post/model"
	"social-photo/modules/post/storage"
)

func CreatePost(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.PostCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreatePostBiz(store)

		if err := business.CreateNewPost(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}
