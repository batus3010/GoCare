package patientGin

import (
	"GoCare/common"
	"GoCare/components/appctx"
	patientBiz "GoCare/module/patient/biz"
	patientStorage "GoCare/module/patient/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListPatient(appCtx appctx.AppContext) func(gin *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		if err := paging.Process(); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}
		store := patientStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := patientBiz.NewListPatientBiz(store)

		result, err := biz.ListPatient(c.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging))
	}
}
