package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/component/tokenprovider/jwt"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User require send verification
// @Description User require send verification account.
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} usermodel.UserCreate "Sign up Success"
// @Router /api/users/send-verification [post]
func SendVerification(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		biz := userbiz.NewSendVerificationBiz(store, appCtx.GetEmailSender(), tokenProvider)

		user := c.MustGet(common.CurrentUser).(common.Requester)

		if err := biz.SendVerification(c.Request.Context(), user.GetId(), user.GetRole()); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
