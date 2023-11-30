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

func CreateParticipant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get cl id
		uid, err := common.FromBase58(ctx.Param(":id"))
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		participant := &participantmodel.ParticipantCreate{
			UserId:      user.GetId(),
			ChallengeId: int(uid.GetLocalID()),
		}

		db := appCtx.GetMainDBConnection()
		store := participantstore.NewSQLStore(db)
		biz := participantbiz.NewCreateParticipantBiz(store, appCtx.GetPubSub())

		err = biz.CreateParticipant(ctx.Request.Context(), participant)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
