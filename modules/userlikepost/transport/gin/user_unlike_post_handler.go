package ginLikePost

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-photo/common"
	"social-photo/modules/userlikepost/biz"
	"social-photo/modules/userlikepost/storage"
	"social-photo/pubsub"
)

func UnlikePost(db *gorm.DB, ps pubsub.PubSub) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSQLStore(db)
		business := biz.NewUserUnlikePostBiz(store, ps)

		if err := business.UnlikePost(c.Request.Context(), requester.GetUserId(), int(id.GetLocalID())); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
