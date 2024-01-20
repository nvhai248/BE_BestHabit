package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Admin banned user
// @Description Admin banned user after successful authentication.
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "User Id"
// @Success 200 {object} common.successRes "Successfully!"
// @Router /api/users/:id/banned [patch]
func BannedUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := userbiz.NewBannedUserBiz(store)

		if err := biz.BannedUser(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
