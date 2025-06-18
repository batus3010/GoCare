package patientGin

import (
	"GoCare/common"
	"GoCare/components/appctx"
	"GoCare/module/patient/biz"
	"GoCare/module/patient/model"
	"GoCare/module/patient/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePatient(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var newData patientModel.PatientCreate

		if err := c.ShouldBind(&newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := patientStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := patientBiz.NewCreateNewPatientBiz(store)
		if err := biz.CreateNewPatient(c.Request.Context(), &newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(newData.Id))
	}
}
