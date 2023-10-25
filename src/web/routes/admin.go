package routes

import (
	"go-base-template/src/handler"
	"go-base-template/src/internal/history"
	"go-base-template/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAdminRoutes(r *gin.Engine, db *gorm.DB) {
	admin := r.Group("/admin")
	admin.Use(middleware.Authorization("ADMIN"))

	// example setup route
	historyRepository := history.NewRepository(db)
	historyService := history.NewService(historyRepository)
	historyHandler := handler.NewHistory(historyService)
	historyHandler.Mount(admin)

}
