package ginstatistical

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/challenge/challengestore"
	"bestHabit/modules/habit/habitstorage"
	"bestHabit/modules/statistical/statisticalbiz"
	"bestHabit/modules/statistical/statisticalmodel"
	"bestHabit/modules/task/taskstorage"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatistical(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter statisticalmodel.Filter

		// Bind query parameters to the filter struct
		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		us := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		ts := taskstorage.NewSQLStore(appCtx.GetMainDBConnection())
		hs := habitstorage.NewSQLStore(appCtx.GetMainDBConnection())
		cs := challengestore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := statisticalbiz.NewStatisticalBiz(us, hs, ts, cs)

		result, err := biz.GetStatistical(filter.Time)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
