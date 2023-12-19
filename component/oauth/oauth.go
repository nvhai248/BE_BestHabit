package oauth

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/component/tokenprovider"
	"bestHabit/component/tokenprovider/jwt"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"encoding/json"
	"fmt"
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

		userStorage := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		userCreate := &usermodel.UserCreate{
			Email:  userInfo["email"].(string),
			GgID:   userInfo["id"].(string),
			Name:   userInfo["name"].(string),
			Avatar: common.NewImageFromGgAuth(userInfo["picture"].(string)),
			Role:   "user",
		}

		// find user by gg id
		userCheck, err := userStorage.FindByGgId(c.Request.Context(), userCreate.GgID)

		if err == common.ErrorNoRows {
			// if user not found => find user by email
			userCheckByMail, err := userStorage.FindByEmail(c.Request.Context(), userCreate.Email)

			if err == common.ErrorNoRows {
				// if not found => create new user and return
				err := userStorage.Create(c.Request.Context(), userCreate)

				if err != nil {
					c.JSON(http.StatusBadRequest, common.NewCustomError(err,
						"Cannot create user!",
						"CannotCreateUser"))
					return
				}

				userReturn, err := userStorage.FindByGgId(c.Request.Context(), userCreate.GgID)

				if err != nil {
					c.JSON(http.StatusBadRequest, common.NewCustomError(err,
						"Cannot find user after create user!",
						"CannotFindUser"))
					return
				}

				payload := tokenprovider.TokenPayload{
					UserId: userReturn.Id,
					Role:   *userReturn.Role,
				}

				accessToken, err := jwt.NewTokenJWTProvider(appCtx.SecretKey()).Generate(payload, 60*60*24*7)
				if err != nil {
					c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
					return
				}

				c.JSON(http.StatusOK, common.SimpleSuccessResponse(accessToken))
			}

			if err != nil {
				c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
				return
			}

			// if found => return err
			c.JSON(http.StatusBadRequest, common.NewCustomError(err,
				fmt.Sprintf("The email %s already used by another account!", *userCheckByMail.Email),
				"CannotFindUser"))
			return
		}

		// if find => return token
		payload := tokenprovider.TokenPayload{
			UserId: userCheck.Id,
			Role:   *userCheck.Role,
		}

		accessToken, err := jwt.NewTokenJWTProvider(appCtx.SecretKey()).Generate(payload, 60*60*24*7)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(accessToken))
	}
}
