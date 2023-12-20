package ginuser

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/component/tokenprovider/jwt"
	"bestHabit/modules/user/userbiz"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Callback OAuth by google
// @Description Simple access to application by Oauth by google
// @Accept  json
// @Produce  json
// @Success 200 {object} tokenprovider.Token "Success!"
// @Router /api/auth/google/callback [get]
func HandleGoogleCallback(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		state := c.Query("state")
		if state != appCtx.GetGGOAuth().GetOauthStateString() {
			panic(common.NewCustomError(nil, "Invalid state!", "InvalidState"))
		}

		code := c.Query("code")
		token, err := appCtx.GetGGOAuth().GetGoogleOauthConfig().Exchange(c, code)
		if err != nil {
			panic(err)
		}

		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		var userInfo map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&userInfo)
		if err != nil {
			panic(err)
		}

		userCreate := &usermodel.UserCreate{
			Email:  userInfo["email"].(string),
			GgID:   userInfo["id"].(string),
			Name:   userInfo["name"].(string),
			Avatar: common.NewImageFromGgAuth(userInfo["picture"].(string)),
			Role:   "user",
		}

		userStorage := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		biz := userbiz.NewGgOauthCallbackBiz(userStorage, appCtx.GetEmailSender(), tokenProvider)

		if token, err := biz.GgOauthCallback(c.Request.Context(), userCreate); err != nil {
			panic(err)
		} else {
			c.JSON(http.StatusOK, common.SimpleSuccessResponse(token))
		}
	}
}
