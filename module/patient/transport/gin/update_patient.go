package patientGin

import (
	"GoCare/common"
	"GoCare/components/appctx"
	patientBiz "GoCare/module/patient/biz"
	patientModel "GoCare/module/patient/model"
	patientStorage "GoCare/module/patient/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdatePatient(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}
		var data patientModel.PatientUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := patientStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := patientBiz.NewUpdatePatientBiz(store)

		if err := biz.UpdatePatient(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
