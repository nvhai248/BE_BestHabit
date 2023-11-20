package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newProfile usermodel.UserUpdate

		if err := ctx.ShouldBindJSON(&newProfile); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewUpdateProfileBiz(store)

		if err := biz.UpdateProfile(ctx.Request.Context(), &newProfile,
			ctx.MustGet(common.CurrentUser).(common.Requester).GetId()); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newProfile))
	}
}
