package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/component/tokenprovider/jwt"
	"bestHabit/modules/user/userbiz"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User require send verification
// @Description User require send verification account.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} common.successRes "Success!"
// @Router /api/users/send-verification [post]
func SendVerification(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		biz := userbiz.NewSendVerificationBiz(appCtx.GetEmailSender(), tokenProvider)

		user := c.MustGet(common.CurrentUser).(common.Requester)

		if err := biz.SendVerification(c.Request.Context(), user.GetEmail(), user.GetId(), user.GetRole()); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
