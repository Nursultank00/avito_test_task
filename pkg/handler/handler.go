package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	services service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/balance", h.getBalance)
	router.POST("/accrue", h.accrueBalance)
	router.POST("/debit", h.debitBalance)
	router.POST("/reserve", h.reserveBalance)
	router.POST("/transfer", h.transferBalance)
	router.POST("/transactions", h.showTransactions)

	return router
}
