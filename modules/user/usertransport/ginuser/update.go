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

// @Summary User Update Profile
// @Description User update profile after successful authentication.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param phone formData string true "Phone"
// @Param name formData string true "Name"
// @Param settings formData string true "Settings"
// @Param image formData string true "Image"
// @Success 200 {object} usermodel.UserUpdate "Update Profile Successfully"
// @Router /api/users/profile [patch]
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
