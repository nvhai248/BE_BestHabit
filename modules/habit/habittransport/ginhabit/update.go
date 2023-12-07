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

// @Summary User Update Habit
// @Description User Update habit after successful authentication.
// @Tags Habits
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "habit Id"
// @Param name formData string true "habit Name"
// @Param description formData string true "Description"
// @Param start_date formData string true "StartDate"
// @Param end_date formData string true "EndDate"
// @Param type formData string true "Type"
// @Param reminder formData string true "Reminder"
// @Param is_count_based formData number true "IsCountBased"
// @Param days body common.Days true "IsCountBased"
// @Param completed_dates body common.Dates true "IsCountBased"
// @Param target body common.Target true "Target"
// @Success 200 {object} habitmodel.HabitUpdate "Successfully update habit!"
// @Router /api/habits/:id [patch]
func UpdateHabit(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		var newData habitmodel.HabitUpdate

		if err := ctx.ShouldBindJSON(&newData); err != nil {
			panic(common.ErrInternal(err))
		}

		store := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := habitbiz.NewUpdateHabitBiz(store)

		if err := biz.Update(ctx.Request.Context(), &newData, int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newData))
	}
}
