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

// @Summary User Update Task
// @Description User Update Task after successful authentication.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "Task Id"
// @Param name formData string true "Task Name"
// @Param description formData string true "Description"
// @Param deadline formData string true "Deadline"
// @Param reminder formData string true "Reminder"
// @Param status formData string true "Status"
// @Success 200 {object} taskmodel.TaskUpdate "Successfully update task!"
// @Router /api/task/:id [patch]
func UpdateTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		var newData taskmodel.TaskUpdate

		if err := ctx.ShouldBindJSON(&newData); err != nil {
			panic(common.ErrInternal(err))
		}

		store := taskstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := taskbiz.NewUpdateTaskBiz(store)

		if err := biz.Update(ctx.Request.Context(), &newData, int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newData))
	}
}
