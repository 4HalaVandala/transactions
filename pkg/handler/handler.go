package handler

import (
	"github.com/4halavandala/transactions/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/transactions", h.GetAllTransactions)
		api.POST("/transactions", h.CreateTransaction)
		api.GET("/transactions/:id", h.GetTransactionById)
	}

	return router
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}
