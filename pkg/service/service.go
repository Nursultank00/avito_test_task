package service

import (
	"github.com/Nursultank00/avito_test_task/models"
	"github.com/Nursultank00/avito_test_task/pkg/repository"
)

type Accounts interface {
	createAccount(account models.Account) (int, error)
	getAllAccounts() ([]models.Account, error)
	deleteAccount(accountId int) error
}

type Balance interface {
	getBalance(accountId int) (int, error)
	accrueBalance(accountId, int, amount int, description string) error
	debitBalance(accountId int, amount int, description string) error
	reserveBalance(acountId int, amount int, description string) error
}

type Service struct {
	Accounts
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
