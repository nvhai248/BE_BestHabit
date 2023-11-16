package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
