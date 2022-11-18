package handler

import (
	"net/http"

	"github.com/Nursultank00/avito_test_task/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createAccount(c *gin.Context) {
	acc := &models.Account{}

	err := h.Services.CreateAccount(acc.AccountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, acc.AccountId)
}

func (h *Handler) getAllAccounts(c *gin.Context) {
	accs, err := h.Services.GetAllAccounts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, accs)
}
