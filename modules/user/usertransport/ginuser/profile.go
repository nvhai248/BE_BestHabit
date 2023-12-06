package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Get Profile
// @Description User get profile after successful authentication.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} common.Requester "Update Profile Successfully"
// @Router /api/users/profile [get]
func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
