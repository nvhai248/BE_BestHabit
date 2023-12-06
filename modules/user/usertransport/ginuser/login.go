package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/component/hasher"
	"bestHabit/component/tokenprovider/jwt"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Basic Login
// @Description User use email and password to login to system
// @Accept  json
// @Produce  json
// @Param email formData string true "Email address"
// @Param password formData string true "Password"
// @Success 200 {object} tokenprovider.Token "Login Successfully"
// @Router /api/login [post]
func BasicLogin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBindJSON(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewBasicLoginBiz(store, tokenProvider, md5, 60*60*24*7)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
