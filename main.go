package main

import (
	"GoCare/components/appctx"
	"GoCare/middleware"
	"GoCare/module/patient/transport/gin"
	userStorage "GoCare/module/user/storage"
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

	authStore := userStorage.NewSQLStore(appCtx.GetMainDBConnection())

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/register", userGin.Register(appCtx))
		v1.POST("/authenticate", userGin.Login(appCtx))
		v1.GET("/profile", middleware.RequiredAuth(appCtx, authStore), userGin.Profile(appCtx))
	}

	// Protected /patients routes
	auth := v1.Group("/", middleware.RequiredAuth(appCtx, authStore))
	{
		// Receptionist-only CRUD
		rec := auth.Group("/", middleware.RequireRoles("receptionist"))
		rec.POST("/patients", patientGin.CreatePatient(appCtx))
		rec.DELETE("/patients/:id", patientGin.DeletePatient(appCtx))

		// Doctor-only read+update
		doc := auth.Group("/", middleware.RequireRoles("doctor"))
		doc.GET("/patients", patientGin.ListPatient(appCtx))
		doc.GET("/patients/:id", patientGin.GetPatient(appCtx))

		auth.PUT("/patients/:id",
			middleware.RequireRoles("receptionist", "doctor"),
			patientGin.UpdatePatient(appCtx),
		)
	}

	// Start server
	addr := "localhost:8080"
	log.Printf("listening on %s…", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
