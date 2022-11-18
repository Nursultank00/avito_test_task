package service

import (
	"fmt"

	"github.com/Nursultank00/avito_test_task/models"
	"github.com/Nursultank00/avito_test_task/pkg/repository"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (bs *BalanceService) GetBalance(accountId int) (map[string]int, error) {
	return bs.repo.GetBalance(accountId)
}

func (bs *BalanceService) AccrueBalance(accountId int, amount int, description string) error {
	if amount <= 0 {
		return fmt.Errorf("amount should be greater than 0")
	}
	return bs.repo.AccrueBalance(accountId, amount, description)
}

func (bs *BalanceService) DebitBalance(accountId int, amount int, description string) error {
	if amount <= 0 {
		return fmt.Errorf("amount should be greater than 0")
	}
	return bs.repo.DebitBalance(accountId, amount, description)
}

func (bs *BalanceService) ReserveBalance(transId int, accountId int, serviceId int, orderId int, amount int, description string) error {
	if amount <= 0 {
		return fmt.Errorf("amount should be greater than 0")
	}
	return bs.repo.ReserveBalance(transId, accountId, serviceId, orderId, amount, description)
}

func (bs *BalanceService) ConfirmationBalance(transId int, accountId int, serviceId int, orderId int, amount int, description string) error {
	if amount <= 0 {
		return fmt.Errorf("amount should be greater than 0")
	}
	return bs.ConfirmationBalance(transId, accountId, serviceId, orderId, amount, description)
}

func (bs *BalanceService) TransferBalance(receiverId int, senderId int, amount int, description string) error {
	if amount <= 0 {
		return fmt.Errorf("amount should be greater than 0")
	}
	return bs.TransferBalance(receiverId, senderId, amount, description)
}

func (bs *BalanceService) GetTransactionHistory(account_id int) ([]*models.Transactions, error) {
	return bs.GetTransactionHistory(account_id)
}
