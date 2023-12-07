package ginchallenge

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/challenge/challengebiz"
	"bestHabit/modules/challenge/challengestore"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User Deleted Challenge
// @Description User Deleted challenge after successful authentication.
// @Tags Challenges (requires admin)
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "challenge Id"
// @Success 200 {object} common.successRes "Successfully deleted challenge!"
// @Router /api/challenges/:id [delete]
func DeleteChallenge(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := challengestore.NewSQLStore(appCtx.GetMainDBConnection())

		biz := challengebiz.NewDeleteChallengeBiz(store)

		err = biz.SoftDeleteChallenge(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
