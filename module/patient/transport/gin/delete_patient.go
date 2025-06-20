package patientGin

import (
	"GoCare/common"
	"GoCare/components/appctx"
	patientBiz "GoCare/module/patient/biz"
	patientStorage "GoCare/module/patient/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeletePatient(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := patientStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := patientBiz.NewDeletePatientBiz(store)

		if err := biz.DeletePatient(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
