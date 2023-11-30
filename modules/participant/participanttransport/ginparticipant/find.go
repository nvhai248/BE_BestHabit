package ginparticipant

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/participant/participantbiz"
	"bestHabit/modules/participant/participantstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindParticipant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get cl id
		uid, err := common.FromBase58(ctx.Param("id"))

		db := appCtx.GetMainDBConnection()
		store := participantstore.NewSQLStore(db)
		biz := participantbiz.NewFindParticipantBiz(store)

		result, err := biz.FindChallengeJoined(ctx.Request.Context(), uid.GetObjectType())

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
