package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get List User
// @Description Admin Get List User after successful authentication.
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param name path string true "User's name"
// @Param page path number true "Page number"
// @Param limit path number true "Limit of user returned!"
// @Param cursor path string true "User Id"
// @Param deadline path string true "Deadline"
// @Success 200 {object} []usermodel.User "Successfully!"
// @Router /api/users [get]
func ListUserByConditions(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter usermodel.UserFilter

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

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewListUserBiz(store)

		result, err := biz.ListTask(c.Request.Context(), &filter, &paging, mapConditions)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeID.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter, 200, "Successful!"))
	}
}
