package middleware

import (
	"bestHabit/common"
	"bestHabit/component"
	"bestHabit/modules/user/userstorage"
	"errors"

	"github.com/gin-gonic/gin"
)

func IsVerifiedUser(appCtx component.AppContext) func(c *gin.Context) {
	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	return func(c *gin.Context) {

		user := c.MustGet(common.CurrentUser).(common.Requester)

		userInfo, err := userStore.FindById(c.Request.Context(), user.GetId())

		if err != nil {
			panic(err)
		}

		if userInfo.Status == common.UserNotVerified {
			panic(common.ErrNoPermission(errors.New("user not verified!")))
		}

		c.Next()
	}
}
