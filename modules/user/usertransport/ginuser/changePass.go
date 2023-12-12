package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/component/hasher"
	"bestHabit/component/tokenprovider/jwt"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User change password
// @Description User can change password after authentication
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param password formData string true "Password"
// @Param new_password formData string true "New Password"
// @Success 200 {object} common.successRes "change password Successfully"
// @Router /api/users/change-password [patch]
func ChangePassword(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newPw usermodel.UpdatePassword

		if err := ctx.ShouldBindJSON(&newPw); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		biz := userbiz.NewChangePassBiz(store, md5, appCtx.GetEmailSender(), tokenProvider)

		if err := biz.ChangePass(ctx.Request.Context(), user.GetEmail(), user.GetId(), user.GetRole(), &newPw); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
