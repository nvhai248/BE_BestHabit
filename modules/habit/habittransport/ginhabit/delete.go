package ginhabit

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/habit/habitbiz"
	"bestHabit/modules/habit/habitstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Deleted Habit
// @Description User Deleted habit after successful authentication.
// @Tags Habits
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "habit Id"
// @Success 200 {object} common.successRes "Successfully deleted habit!"
// @Router /api/habits/:id [delete]
func SoftDeleteHabit(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := habitbiz.NewSoftDeleteHabitBiz(store, appCtx.GetPubSub())

		if err := biz.SoftDeleteHabit(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
