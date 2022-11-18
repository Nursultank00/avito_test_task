package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	AccountId int `json:"account_id"`
}

type Request_accrue_debit struct {
	AccountId   int    `json:"account_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type Request_reserve_confirm struct {
	TransactionId int    `json:"transaction_id"`
	AccountId     int    `json:"account_id"`
	ServiceId     int    `json:"service_id"`
	OrderId       int    `json:"order_id"`
	Amount        int    `json:"amount"`
	Description   string `json:"description"`
}

type Request_transfer struct {
	ReceiverId  int    `json:"receiver_id"`
	SenderId    int    `json:"sender_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func (h *Handler) getBalance(c *gin.Context) {
	var input Request

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad request")
		return
	}

	balance, err := h.Services.GetBalance(input.AccountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't find account")
		return
	}
	c.JSON(http.StatusOK, balance)
}

func (h *Handler) accrueBalance(c *gin.Context) {
	var input Request_accrue_debit

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.Services.AccrueBalance(input.AccountId, input.Amount, input.Description)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Successful acrrual")
}

func (h *Handler) debitBalance(c *gin.Context) {
	var input Request_accrue_debit

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.Services.DebitBalance(input.AccountId, input.Amount, input.Description)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Successful debit")
}

func (h *Handler) reserveBalance(c *gin.Context) {
	var input Request_reserve_confirm

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.Services.ReserveBalance(input.TransactionId, input.AccountId, input.ServiceId,
		input.OrderId, input.Amount, input.Description)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Successful reservation")
}

func (h *Handler) confirmationBalance(c *gin.Context) {
	var input Request_reserve_confirm

	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.Services.ConfirmationBalance(input.TransactionId, input.AccountId, input.ServiceId,
		input.OrderId, input.Amount, input.Description)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Successful confirmation")
}

func (h *Handler) transferBalance(c *gin.Context) {
	var input Request_transfer

	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.Services.TransferBalance(input.ReceiverId, input.SenderId, input.Amount, input.Description)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Successful transfer")
}

func (h *Handler) getTransactionsHistory(c *gin.Context) {
	var input Request
	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	transactions, err := h.Services.GetTransactionHistory(input.AccountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, transactions)
}
