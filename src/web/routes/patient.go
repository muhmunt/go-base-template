package routes

import (
	"go-base-template/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPatientRoutes(r *gin.Engine, db *gorm.DB) {
	patient := r.Group("/patient")
	patient.Use(middleware.Authorization("PATIENT"))

	// example setup route
	// historyRepository := history.NewRepository(db)
	// historyService := history.NewService(historyRepository)
	// historyHandler := handler.NewHistory(historyService)
	// historyHandler.Mount(admin)

}
