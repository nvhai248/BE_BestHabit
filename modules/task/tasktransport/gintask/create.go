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

// @Summary User Create New Task
// @Description User Create New Task after successful authentication.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param name formData string true "Task Name"
// @Param description formData string true "Description"
// @Param deadline formData string true "Deadline"
// @Param reminder formData string true "Reminder"
// @Success 200 {object} taskmodel.TaskCreate "Successfully created task!"
// @Router /api/tasks [post]
func CreateTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var taskData taskmodel.TaskCreate

		if err := ctx.ShouldBindJSON(&taskData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := taskstorage.NewSQLStore(db)
		biz := taskbiz.NewCreateTaskBiz(store, appCtx.GetPubSub())

		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		err := biz.CreateTask(ctx.Request.Context(), &taskData, user.GetId())

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(taskData))
	}
}
