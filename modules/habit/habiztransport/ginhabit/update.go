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

func UpdateTask(appCtx component.AppContext) gin.HandlerFunc {
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
