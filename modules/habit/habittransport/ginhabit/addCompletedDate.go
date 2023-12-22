package ginhabit

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/habit/habitbiz"
	"bestHabit/modules/habit/habitstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Add Completed date of Habit
// @Description User Add Completed date of Habit after successful authentication.
// @Tags Habits
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "habit Id"
// @Param date body common.CompleteDate true "Date"
// @Success 200 {object} common.CompleteDate "Successfully!"
// @Router /api/habits/:id/confirm-completed [patch]
func AddCompletedDate(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data common.CompleteDate

		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := habitbiz.NewAddCompletedDateBiz(store, appCtx.GetPubSub())

		err = biz.AddCompletedDate(ctx.Request.Context(), &data, int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
