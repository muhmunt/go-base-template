package routes

import (
	"go-base-template/src/handler"
	"go-base-template/src/internal/history"
	"go-base-template/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupDoctorRoutes(r *gin.Engine, db *gorm.DB) {
	doctor := r.Group("/doctor")
	doctor.Use(middleware.Authorization("DOCTOR"))

	// example setup route
	historyRepository := history.NewRepository(db)
	historyService := history.NewService(historyRepository)
	historyHandler := handler.NewHistory(historyService)
	historyHandler.Mount(doctor)

}
