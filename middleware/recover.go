package middleware

import (
	"GoCare/common"
	"GoCare/components/appctx"
	"github.com/gin-gonic/gin"
)

func Recover(ac appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				// check if it's AppError
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err) // panic to log out error trace, gin default doesn't print it out
					return
				}
				// if not, then it's go error, not server error
				appErr := common.ErrorInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()
		c.Next()
	}
}
