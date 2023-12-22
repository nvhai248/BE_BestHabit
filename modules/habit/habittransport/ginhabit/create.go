package ginhabit

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/habit/habitbiz"
	"bestHabit/modules/habit/habitmodel"
	"bestHabit/modules/habit/habitstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Create New Habit
// @Description User Create New habit after successful authentication.
// @Tags Habits
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param name formData string true "habit Name"
// @Param description formData string true "Description"
// @Param start_date formData string true "StartDate"
// @Param end_date formData string true "EndDate"
// @Param type formData string true "Type"
// @Param reminder formData string true "Reminder"
// @Param is_count_based formData number true "IsCountBased"
// @Param days body common.Days true "IsCountBased"
// @Param completed_dates body common.CompleteDates true "IsCountBased"
// @Param target body common.Target true "Target"
// @Success 200 {object} habitmodel.HabitCreate "Successfully created habit!"
// @Router /api/habits [post]
func CreateHabit(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var habitData habitmodel.HabitCreate

		if err := ctx.ShouldBindJSON(&habitData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := habitbiz.NewCreateHabitBiz(store, appCtx.GetPubSub())

		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		err := biz.CreateHabit(ctx.Request.Context(), &habitData, user.GetId())

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(habitData))
	}
}
