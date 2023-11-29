package ginchallenge

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/challenge/challengebiz"
	"bestHabit/modules/challenge/challengemodel"
	"bestHabit/modules/challenge/challengestore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChallenge(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var challenge challengemodel.ChallengeCreate

		if err := ctx.ShouldBindJSON(&challenge); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := challengestore.NewSQLStore(db)
		biz := challengebiz.NewCreateChallengeBiz(store)

		err := biz.CreateChallenge(ctx.Request.Context(), &challenge)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(challenge))
	}
}
