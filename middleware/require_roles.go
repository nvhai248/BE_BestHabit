package middleware

import (
	"bestHabit/common"
	"bestHabit/component"

	"github.com/gin-gonic/gin"
)

func RequireRoles(appCtx component.AppContext, roles ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for i := range roles {
			if u.GetRole() == roles[i] {
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(nil))
	}
}
