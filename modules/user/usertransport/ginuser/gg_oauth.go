package ginuser

import (
	"bestHabit/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary OAuth by google
// @Description Simple access to application by Oauth by google
// @Accept  json
// @Produce  json
// @Success 200 {object} string "redirect to /api/auth/google/callback"
// @Router /api/auth/google [get]
func HandleGoogleLogin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := appCtx.GetGGOAuth().GetGoogleOauthConfig().AuthCodeURL(
			appCtx.GetGGOAuth().GetOauthStateString(),
		)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
