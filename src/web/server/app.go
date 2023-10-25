package server

import (
	"go-base-template/src/middleware"
	"go-base-template/src/web/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	router := gin.Default()

	port := "3030"
	router.Use(middleware.HandleAuthentication())

	routes.SetupDoctorRoutes(router, db)
	routes.SetupPatientRoutes(router, db)
	routes.SetupAdminRoutes(router, db)
	routes.SetupClinicRoutes(router, db)

	router.Run(":" + port)
}
