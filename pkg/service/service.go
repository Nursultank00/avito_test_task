package service

import (
	"github.com/Nursultank00/avito_test_task/models"
	"github.com/Nursultank00/avito_test_task/pkg/repository"
)

type Account interface {
	CreateAccount(account_id int) error
	GetAllAccounts() ([]*models.Account, error)
}

type Balance interface {
	GetBalance(accountId int) (map[string]int, error)
	AccrueBalance(accountId int, amount int, description string) error
	DebitBalance(accountId int, amount int, description string) error
	ReserveBalance(transId int, accountId int, serviceId int, orderId int, amount int, description string) error
	ConfirmationBalance(transId int, accountId int, serviceId int, orderId int, amount int, description string) error
	TransferBalance(receiverId int, senderId int, amount int, description string) error
	GetTransactionHistory(account_id int) ([]*models.Transactions, error)
}

type Service struct {
	Account
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
		Balance: NewBalanceService(repos.Balance),
	}
}
