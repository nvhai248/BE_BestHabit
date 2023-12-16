package oauth

import (
	"bestHabit/component"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/api/auth/google/callback",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "random"
)

func HandleGoogleLogin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(googleOauthConfig.ClientID, googleOauthConfig.ClientSecret)

		url := googleOauthConfig.AuthCodeURL(oauthStateString)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func HandleGoogleCallback(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		state := c.Query("state")
		if state != oauthStateString {
			c.String(http.StatusBadRequest, "Invalid oauth state")
			return
		}

		code := c.Query("code")
		token, err := googleOauthConfig.Exchange(c, code)
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
		c.Redirect(http.StatusSeeOther, "/api/ping")
	}
}
