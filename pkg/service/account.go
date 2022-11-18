package service

import (
	"github.com/Nursultank00/avito_test_task/models"
	"github.com/Nursultank00/avito_test_task/pkg/repository"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (as *AccountService) CreateAccount(account_id int) error {
	return as.repo.CreateAccount(account_id)
}

func (as *AccountService) GetAllAccounts() ([]*models.Account, error) {
	return as.repo.GetAllAccounts()
}
