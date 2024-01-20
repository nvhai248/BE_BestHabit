package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Admin find user
// @Description Admin find user after successful authentication.
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "User Id"
// @Success 200 {object} common.successRes "Successfully!"
// @Router /api/users/:id [get]
func FindUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := userbiz.NewFindUserBiz(store)

		user, err := biz.FindUser(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
