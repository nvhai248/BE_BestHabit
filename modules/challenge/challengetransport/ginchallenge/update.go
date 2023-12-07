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

// @Summary User Update Challenge
// @Description User Update challenge after successful authentication.
// @Tags Challenges (requires admin)
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path string true "challenge Id"
// @Param name formData string true "challenge Name"
// @Param description formData string true "Description"
// @Param start_date formData string true "startDate"
// @Param end_date formData string true "endDate"
// @Param experience_point formData number true "ExperiencePoint"
// @Success 200 {object} challengemodel.ChallengeUpdate "Successfully update challenge!"
// @Router /api/challenges/:id [patch]
func UpdateChallenge(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		var newData challengemodel.ChallengeUpdate

		if err := ctx.ShouldBindJSON(&newData); err != nil {
			panic(common.ErrInternal(err))
		}

		store := challengestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := challengebiz.NewUpdateChallengeBiz(store)

		if err := biz.Update(ctx.Request.Context(), &newData, int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newData))
	}
}
