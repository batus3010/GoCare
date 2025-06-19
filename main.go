package main

import (
	"GoCare/components/appctx"
	"GoCare/middleware"
	"GoCare/module/patient/transport/gin"
	userGin "GoCare/module/user/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	secretKey := os.Getenv("SYSTEM_SECRET")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)

	db = db.Debug()

	appCtx := appctx.NewAppContext(db, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/register", userGin.Register(appCtx))
		v1.POST("/authenticate", userGin.Login(appCtx))
		patient := v1.Group("/patients")
		{
			patient.POST("", patientGin.CreatePatient(appCtx))
			patient.GET("/:id", patientGin.GetPatient(appCtx))
			patient.GET("", patientGin.ListPatient(appCtx))
			patient.PUT("/:id", patientGin.UpdatePatient(appCtx))
			patient.DELETE("/:id", patientGin.DeletePatient(appCtx))
		}
	}

	r.Run("localhost:8080")
}
