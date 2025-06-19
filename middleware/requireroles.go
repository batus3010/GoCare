package middleware

import (
	"GoCare/common"
	userModel "GoCare/module/user/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get(common.CurrentUser)
		if !exists {
			panic(common.NewUnauthorizedErrorResponse(
				nil,
				"authentication required",
				"ErrUnauthorized"))
		}
		user, ok := u.(*userModel.User)
		if !ok {
			panic(common.NewUnauthorizedErrorResponse(
				nil,
				"authentication required",
				"ErrUnauthorized",
			))
		}

		for _, role := range allowedRoles {
			if user.Role == role {
				c.Next()
				return
			}
		}

		panic(common.NewFullErrorResponse(
			http.StatusForbidden,
			nil,
			fmt.Sprintf("role %q cannot access this resource", user.Role),
			fmt.Sprintf("ForbiddenRole:%s", user.Role),
			"ErrForbidden",
		))
	}
}
