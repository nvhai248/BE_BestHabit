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

// @Summary Get List User's Habit
// @Description User Get List User's habit after successful authentication.
// @Tags Habits
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param name path string true "Habit's name"
// @Param page path number true "Page number"
// @Param limit path number true "Limit of habits returned!"
// @Param cursor path string true "habit Id"
// @Param deadline path string true "Deadline"
// @Success 200 {object} []habitmodel.Habit "Successfully!"
// @Router /api/habits [get]
func ListHabitByConditions(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter habitmodel.HabitFilter

		// Bind query parameters to the filter struct
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		// Bind query parameters to the filter struct
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		var conditions common.Conditions

		// Bind conditions from
		if err := c.ShouldBind(&conditions); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		mapConditions := common.ConvertToMap(conditions)

		mapConditions["user_id"] = c.MustGet(common.CurrentUser).(common.Requester).GetId()

		store := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := habitbiz.NewListHabitBiz(store)

		result, err := biz.ListHabit(c.Request.Context(), &filter, &paging, mapConditions)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeID.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter, 200, "Get list habit successfully!"))
	}
}
