package ginLikePost

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-photo/common"
	"social-photo/modules/userlikepost/biz"
	"social-photo/modules/userlikepost/storage"
)

func ListUserLiked(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		paging.Process()

		id, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewListUserLikePostBiz(store)

		result, err := business.ListUserLikePost(c.Request.Context(), int(id.GetLocalID()), &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
