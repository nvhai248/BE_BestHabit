package ginparticipant

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/participant/participantbiz"
	"bestHabit/modules/participant/participantmodel"
	"bestHabit/modules/participant/participantstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateParticipant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		var newData participantmodel.ParticipantUpdate

		if err := ctx.ShouldBindJSON(&newData); err != nil {
			panic(common.ErrInternal(err))
		}

		store := participantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := participantbiz.NewUpdateParticipantBiz(store)

		if err := biz.Update(ctx.Request.Context(), &newData, int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newData))
	}
}
