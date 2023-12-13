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

// @Summary User require send email reset password
// @Description User can require send email reset password if user forgot password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param email formData string true "Email"
// @Success 200 {object} common.successRes "send successfully!"
// @Router /api/users/send-reset-password [post]
func SenResetPw(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mailInput common.ToUser

		if err := ctx.ShouldBindJSON(&mailInput); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		biz := userbiz.NewSendResetPwBiz(store, appCtx.GetEmailSender(), tokenProvider)

		if err := biz.SendResetPw(ctx.Request.Context(), mailInput); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
