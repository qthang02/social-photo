package ginUser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-photo/common"
)

func Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
