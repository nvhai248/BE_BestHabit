package ginhabit

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/habit/habitbiz"
	"bestHabit/modules/habit/habitstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Find Habit
// @Description User Find habit after successful authentication.
// @Tags Habits
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "habit Id"
// @Success 200 {object} habitmodel.HabitFind "Successfully!"
// @Router /api/habits/:id [get]
func FindHabit(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := habitbiz.NewFindHabitBiz(store)

		result, err := biz.FindHabit(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
