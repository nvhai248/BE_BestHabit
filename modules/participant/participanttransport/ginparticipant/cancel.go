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

// @Summary User Cancel the Challenge
// @Description User Cancel the challenge after successful authentication.
// @Tags Challenges
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "challenge Id"
// @Success 200 {object} common.successRes "Successfully!"
// @Router /api/challenges/:id/user-cancel [delete]
func CancelParticipant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get cl id
		uid, err := common.FromBase58(ctx.Param("id"))
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		participant := &participantmodel.ParticipantCancel{
			UserId:      user.GetId(),
			ChallengeId: int(uid.GetLocalID()),
		}

		db := appCtx.GetMainDBConnection()
		store := participantstore.NewSQLStore(db)
		biz := participantbiz.NewCancelParticipantBiz(store, appCtx.GetPubSub())

		err = biz.CancelChallenge(ctx.Request.Context(), participant)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
