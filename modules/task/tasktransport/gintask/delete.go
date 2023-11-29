package gintask

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/task/taskbiz"
	"bestHabit/modules/task/taskstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SoftDeleteTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		db := appCtx.GetMainDBConnection()
		store := taskstorage.NewSQLStore(db)
		biz := taskbiz.NewDeleteTaskBiz(store, appCtx.GetPubSub())

		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		err := biz.SoftDeleteTask(ctx.Request.Context(), user.GetId())

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
