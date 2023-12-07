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

// @Summary User Create New Challenge
// @Description User Create New challenge after successful authentication.
// @Tags Challenges (requires admin)
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param name formData string true "challenge Name"
// @Param description formData string true "Description"
// @Param start_date formData string true "startDate"
// @Param end_date formData string true "endDate"
// @Param experience_point formData number true "ExperiencePoint"
// @Success 200 {object} challengemodel.ChallengeCreate "Successfully created challenge!"
// @Router /api/challenges [post]
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
