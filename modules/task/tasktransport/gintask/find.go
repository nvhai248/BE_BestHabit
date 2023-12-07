package gintask

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/task/taskbiz"
	"bestHabit/modules/task/taskstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Find Task
// @Description User Find Task after successful authentication.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "Task Id"
// @Success 200 {object} taskmodel.TaskFind "Successfully!"
// @Router /api/tasks/:id [get]
func FindTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := taskstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := taskbiz.NewFindTaskBiz(store)

		result, err := biz.FindTask(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
