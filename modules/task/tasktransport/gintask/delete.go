package gintask

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/task/taskbiz"
	"bestHabit/modules/task/taskstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Deleted Task
// @Description User Deleted Task after successful authentication.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "Task Id"
// @Success 200 {object} common.successRes "Successfully deleted task!"
// @Router /api/tasks/:id [delete]
func SoftDeleteTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		db := appCtx.GetMainDBConnection()
		store := taskstorage.NewSQLStore(db)
		biz := taskbiz.NewDeleteTaskBiz(store, appCtx.GetPubSub())

		err = biz.SoftDeleteTask(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
