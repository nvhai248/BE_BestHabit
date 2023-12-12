package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Verify User
// @Description User can verify after authentication (only you can verify your account)
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param token path string true "Token"
// @Success 200 {object} common.successRes "Verify Successfully"
// @Router /api/users/verify/:token [patch]
func Verify(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewVerifyUserBiz(store)

		if err := biz.VerifyBiz(ctx.Request.Context(), user.GetId()); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
