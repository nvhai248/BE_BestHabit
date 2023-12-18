package oauth

import (
	"bestHabit/component"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleGoogleLogin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := appCtx.GetGGOAuth().GetGoogleOauthConfig().AuthCodeURL(
			appCtx.GetGGOAuth().GetOauthStateString(),
		)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func HandleGoogleCallback(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		state := c.Query("state")
		if state != appCtx.GetGGOAuth().GetOauthStateString() {
			c.String(http.StatusBadRequest, "Invalid oauth state")
			return
		}

		code := c.Query("code")
		token, err := appCtx.GetGGOAuth().GetGoogleOauthConfig().Exchange(c, code)
		if err != nil {
			c.String(http.StatusBadRequest, "Failed to exchange token")
			return
		}

		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			c.String(http.StatusBadRequest, "Failed to get user info")
			return
		}

		defer response.Body.Close()

		var userInfo map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&userInfo)
		if err != nil {
			c.String(http.StatusBadRequest, "Failed to parse user info")
			return
		}

		fmt.Print(userInfo)
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("%s/api/ping", os.Getenv("SITE_DOMAIN")))
	}
}
