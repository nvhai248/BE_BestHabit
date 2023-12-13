package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/component/hasher"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User reset password
// @Description User reset password if user forgot the password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param password formData string true "Password"
// @Param token header string true "Authorization token"
// @Success 200 {object} common.successRes "reset password Successfully"
// @Router /api/users/reset-password [patch]
func ResetPassword(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newPw usermodel.ResetPassword

		if err := ctx.ShouldBindJSON(&newPw); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewResetPasswordBiz(store, md5)

		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		if err := biz.ResetPassword(ctx.Request.Context(), &newPw, user.GetId()); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
