package userGin

import (
	"GoCare/common"
	"GoCare/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
