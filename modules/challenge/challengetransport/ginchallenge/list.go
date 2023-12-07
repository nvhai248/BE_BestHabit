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

// @Summary Get List User's Challenge
// @Description User Get List User's challenge after successful authentication.
// @Tags Challenges
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param name path string true "challenge's name"
// @Param page path number true "Page number"
// @Param limit path number true "Limit of challenges returned!"
// @Param cursor path string true "challenge Id"
// @Param deadline path string true "Deadline"
// @Success 200 {object} []challengemodel.Challenge "Successfully!"
// @Router /api/challenges [get]
func ListChallengeByConditions(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter challengemodel.ChallengeFilter

		// Bind query parameters to the filter struct
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		// Bind query parameters to the filter struct
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		var conditions common.Conditions

		// Bind conditions from
		if err := c.ShouldBind(&conditions); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		mapConditions := common.ConvertToMap(conditions)

		store := challengestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := challengebiz.NewListChallengeBiz(store)

		result, err := biz.ListChallenge(c.Request.Context(), &filter, &paging, mapConditions)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeID.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter, 200, "Get list challenge successfully!"))
	}
}
