package ginLikePost

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-photo/common"
	"social-photo/modules/userlikepost/biz"
	"social-photo/modules/userlikepost/model"
	"social-photo/modules/userlikepost/storage"
	"social-photo/pubsub"
	"time"
)

func LikePost(db *gorm.DB, ps pubsub.PubSub) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSQLStore(db)
		//postStore := postStorage.NewSQLStore(db)
		business := biz.NewUserLikePostBiz(store, ps)

		now := time.Now().UTC()

		if err := business.LikePost(c.Request.Context(), &model.Like{
			UserId:    requester.GetUserId(),
			PostId:    int(id.GetLocalID()),
			CreatedAt: &now,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
