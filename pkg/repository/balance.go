package repository

import (
	"fmt"

	"github.com/Nursultank00/avito_test_task/models"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (b *BalancePostgres) GetBalance(accountId int) (map[string]int, error) {
	acc := &models.Account{}
	query := "SELECT * FROM accounts WHERE account_id = $1"
	if err := b.db.Get(acc, query, accountId); err != nil {
		return nil, err
	}
	return map[string]int{
		"main_balance":    acc.MainBalance,
		"reserve_balance": acc.ReserveBalance,
	}, nil
}

func (b *BalancePostgres) AccrueBalance(accountId int, amount int, description string) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	updateQuery := "UPDATE accounts SET main_balance = main_balance + $2 WHERE account_id = $1"

	if _, err := tx.Exec(updateQuery, accountId, amount); err != nil {
		tx.Rollback()
		return err
	}

	transactionQuery := fmt.Sprint("INSERT INTO transactions" +
		"(account_id, amount, trans_type, description)" +
		"VALUES ($1, $2, $3, $4)")

	_, err = tx.Exec(transactionQuery, accountId, amount, "Accrual", description)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (b *BalancePostgres) DebitBalance(accountId int, amount int, description string) error {
	curBalance, err := b.GetBalance(accountId)
	if err != nil {
		return err
	}

	if curBalance["main_balance"] < amount {
		return fmt.Errorf("not enough money on main_balance")
	}

	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	updateQuery := "UPDATE accounts SET main_balance = main_balance - $2 WHERE account_id = $1"

	if _, err := tx.Exec(updateQuery, accountId, amount); err != nil {
		tx.Rollback()
		return err
	}

	transactionQuery := fmt.Sprint("INSERT INTO transactions" +
		"(account_id, amount, trans_type, description)" +
		"VALUES ($1, $2, $3, $4)")

	_, err = tx.Exec(transactionQuery, accountId, amount, "Debit", description)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (b *BalancePostgres) ReserveBalance(transId int, accountId int, serviceId int, orderId int, amount int, description string) error {
	curBalance, err := b.GetBalance(accountId)

	if err != nil {
		return err
	}

	if curBalance["main_balance"] < amount {
		return fmt.Errorf("not enough money on main_balance")
	}

	tx, err := b.db.Begin()

	if err != nil {
		return err
	}
	updateMainBalanceQuery := "UPDATE accounts SET main_balance = main_balance - $2 WHERE account_id = $1"

	if _, err := tx.Exec(updateMainBalanceQuery, accountId, amount); err != nil {
		tx.Rollback()
		return err
	}

	updateReserveBalanceQuery := "UPDATE accounts SET reserve_balance = reserve_balance + $2 WHERE account_id = $1"

	if _, err := tx.Exec(updateReserveBalanceQuery, accountId, amount); err != nil {
		tx.Rollback()
		return err
	}

	transactionQuery := fmt.Sprint("INSERT INTO transactions" +
		"(account_id, service_id, order_id, amount, trans_type, description)" +
		"VALUES ($1, $2, $3, $4, $5, $6)")

	_, err = tx.Exec(transactionQuery, accountId, serviceId, orderId, amount, "Reservation", description)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (b *BalancePostgres) ConfirmationBalance(transId int, accountId int, serviceId int, orderId int, amount int, description string) error {
	tx, err := b.db.Begin()

	if err != nil {
		return err
	}

	updateReserveBalanceQuery := "UPDATE accounts SET reserve_balance = reserve_balance - $2 WHERE account_id = $1"

	if _, err := tx.Exec(updateReserveBalanceQuery, accountId, amount); err != nil {
		tx.Rollback()
		return err
	}

	transactionQuery := "INSERT INTO transactions" +
		"(account_id, service_id, order_id, amount, trans_type, description)" +
		"VALUES ($1, $2, $3, $4, $5, $6)"

	_, err = tx.Exec(transactionQuery, accountId, serviceId, orderId, amount, "Confirmation of the reservation", description)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (b *BalancePostgres) TransferBalance(receiverId int, senderId int, amount int, description string) error {
	senderBalance, err := b.GetBalance(senderId)

	if err != nil {
		return err
	}

	if senderBalance["main_balance"] < amount {
		return fmt.Errorf("not enough money on main_balance")
	}

	if _, err := b.GetBalance(receiverId); err != nil {
		return err
	}

	tx, err := b.db.Begin()

	if err != nil {
		return err
	}

	updateSenderBalanceQuery := "UPDATE accounts SET main_balance = main_balance - $2 WHERE account_id = $1"
	updateReceiverBalanceQuery := "UPDATE accounts SET main_balance = main_balance + $2 WHERE account_id = $1"

	if _, err := tx.Exec(updateSenderBalanceQuery, senderId, amount); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(updateReceiverBalanceQuery, receiverId, amount); err != nil {
		tx.Rollback()
		return err
	}

	senderTransactionQuery := "INSERT INTO transactions" +
		"(account_id, amount, trans_type, description)" +
		"VALUES ($1, $2, $3, $4)"
	receiverTransactionQuery := "INSERT INTO transactions" +
		"(account_id, amount, trans_type, description)" +
		"VALUES ($1, $2, $3, $4)"

	_, err = tx.Exec(senderTransactionQuery, senderId, amount, "Send transfer", description)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(receiverTransactionQuery, receiverId, amount, "Receive transfer", description)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (b *BalancePostgres) GetTransactionHistory(account_id int) ([]*models.Transactions, error) {
	transactions := make([]*models.Transactions, 0)

	query := "SELECT * FROM transactions WHERE account_id = $1"

	err := b.db.Select(&transactions, query, account_id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
