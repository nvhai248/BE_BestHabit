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

func CreateHabit(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var taskData habitmodel.HabitCreate

		if err := ctx.ShouldBindJSON(&taskData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := habitbiz.NewCreateHabitBiz(store, appCtx.GetPubSub())

		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		err := biz.CreateHabit(ctx.Request.Context(), &taskData, user.GetId())

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(taskData))
	}
}
