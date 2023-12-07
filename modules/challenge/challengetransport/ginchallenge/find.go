package ginchallenge

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/challenge/challengebiz"
	"bestHabit/modules/challenge/challengestore"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Find Challenge
// @Description User Find challenge after successful authentication.
// @Tags Challenges
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "challenge Id"
// @Success 200 {object} challengemodel.ChallengeFind "Successfully!"
// @Router /api/challenges/:id [get]
func FindChallenge(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := challengestore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := challengebiz.NewFindChallengeBiz(store)

		result, err := biz.FindChallenge(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
