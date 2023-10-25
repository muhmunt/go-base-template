package handler

import (
	"go-base-template/src/internal/history"

	"github.com/gin-gonic/gin"
)

type historyHandler struct {
	historyService history.Service
}

func NewHistory(service history.Service) *historyHandler {
	return &historyHandler{service}
}

func (h *historyHandler) Mount(doctor *gin.RouterGroup) {
	history := doctor.Group("/history")
	history.POST("/store", h.StoreHistory)
}

func (h *historyHandler) StoreHistory(c *gin.Context) {

	c.JSON(200, "halo")
}
