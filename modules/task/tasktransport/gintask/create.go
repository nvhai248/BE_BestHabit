package gintask

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/task/taskbiz"
	"bestHabit/modules/task/taskmodel"
	"bestHabit/modules/task/taskstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var taskData taskmodel.TaskCreate

		if err := ctx.ShouldBindJSON(&taskData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := taskstorage.NewSQLStore(db)
		biz := taskbiz.NewTaskBiz(store)

		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		err := biz.CreateTask(ctx.Request.Context(), &taskData, user.GetId())

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(taskData))
	}
}
