package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Update Profile
// @Description User update profile after successful authentication.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param device_token formData string true "Device Token from FMC"
// @Success 200 {object} common.DvToken "Update Profile Successfully"
// @Router /api/users/profile [patch]
func UpdateDeviceToken(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dvToken common.DvToken
		if err := ctx.ShouldBindJSON(&dvToken); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewUpdateDVTokenBiz(store)

		if err := biz.UpdateDVToken(ctx.Request.Context(), &dvToken,
			ctx.MustGet(common.CurrentUser).(common.Requester).GetId()); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(dvToken))
	}
}
