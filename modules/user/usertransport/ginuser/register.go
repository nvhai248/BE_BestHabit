package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/component/hasher"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Basic Register
// @Description User create new Account by providing email and password
// @Accept  json
// @Produce  json
// @Param email formData string true "Email address"
// @Param password formData string true "Password"
// @Param phone formData string true "Phone"
// @Param name formData string true "Name"
// @Success 200 {object} usermodel.UserCreate "Sign up Success"
// @Router /api/register [post]
func BasicRegister(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewBasicRegisterBiz(store, md5)

		if err := biz.BasicRegister(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
