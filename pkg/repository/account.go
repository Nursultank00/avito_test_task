package repository

import (
	"fmt"

	"github.com/Nursultank00/avito_test_task/models"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (acc *AccountPostgres) CreateAccount(account_id int) error {
	query := fmt.Sprint("INSERTO INTO accounts (account_id, main_balance, reserve_balance)" +
		"VALUES ($1, 0, 0)")
	if _, err := acc.db.Query(query, account_id); err != nil {
		return err
	}
	return nil
}

func (acc *AccountPostgres) GetAllAccounts() ([]*models.Account, error) {
	var all_accounts []*models.Account
	query := "SELECT * FROM accounts"
	err := acc.db.Select(&all_accounts, query)
	if err != nil {
		return nil, err
	}
	return all_accounts, nil
}
