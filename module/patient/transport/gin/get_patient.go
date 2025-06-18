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

func GetPatient(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := patientStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := patientBiz.NewGetPatientBiz(store)
		data, err := biz.GetPatient(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
