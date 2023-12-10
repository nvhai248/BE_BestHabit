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

// @Summary Basic Register
// @Description User create new Account by providing email and password
// @Accept  json
// @Produce  json
// @Param email formData string true "Email address"
// @Param password formData string true "Password"
// @Param phone formData string true "Phone"
// @Param name formData string true "Name"
// @Param avatar body common.Image true "Avatar"
// @Param settings body common.Settings true "Settings"
// @Success 200 {object} usermodel.UserCreate "Sign up Success"
// @Router /api/register [post]
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
