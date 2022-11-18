package handler

import (
	"github.com/Nursultank00/avito_test_task/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	balance := router.Group("/balance")
	{
		balance.GET("", h.getBalance)
		balance.POST("/accrue", h.accrueBalance)
		balance.POST("/debit", h.debitBalance)
		balance.POST("/reserve", h.reserveBalance)
		balance.POST("/confirmation", h.confirmationBalance)
		balance.POST("/transfer", h.transferBalance)
		balance.GET("/transactions", h.getTransactionsHistory)
	}
	accounts := router.Group("accounts")
	{
		accounts.GET("/all", h.getAllAccounts)
		accounts.POST("/create-account", h.createAccount)
	}
	return router
}
