package ginparticipant

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/participant/participantbiz"
	"bestHabit/modules/participant/participantstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Find participant
// @Description User Find participant after successful authentication.
// @Tags Participants
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "challenge Id"
// @Success 200 {object} participantmodel.ParticipantFind "Successfully!"
// @Router /api/participants/:id [get]
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
