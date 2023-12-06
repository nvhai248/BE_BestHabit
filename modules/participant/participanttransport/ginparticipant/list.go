package ginparticipant

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/participant/participantbiz"
	"bestHabit/modules/participant/participantstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListChallengeJoined(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		// Bind query parameters to the filter struct
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := participantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := participantbiz.NewListParticipantBiz(store)

		result, err := biz.ListChallengeJoined(c.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeID.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil, 200, "Successful!"))
	}
}
