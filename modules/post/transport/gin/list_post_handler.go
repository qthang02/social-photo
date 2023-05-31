package ginPost

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-photo/common"
	"social-photo/modules/post/biz"
	"social-photo/modules/post/storage"
)

func ListPost(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		paging.Process()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSQLStore(db)
		business := biz.NewListPostBiz(store, requester)

		data, err := business.ListPostBiz(c.Request.Context(), &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		for i := range data {
			data[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
