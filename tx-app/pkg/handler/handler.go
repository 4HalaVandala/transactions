package handler

import (
	"github.com/4halavandala/transactions/tx-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		tx := api.Group("/transactions")
		{
			tx.GET("/", h.GetAllTransactions)
			tx.POST("/", h.CreateTransaction)
			tx.GET("/:id", h.GetTransactionById)
		}
	}

	return router
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}
