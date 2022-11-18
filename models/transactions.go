package models

import "time"

type Transactions struct {
	TransactionId int       `json:"transasction_id" db:"transasction_id"`
	AccountId     int       `json:"account_id" db:"account_id"`
	ServiceId     int       `json:"service_id" db:"service_id"`
	OrderId       int       `json:"order_id" db:"order_id"`
	Amount        int       `json:"amount" db:"amount"`
	Datetime      time.Time `json:"datetime" db:"datetime"`
	TransType     string    `json:"trans_type" db:"trans_type"`
	Description   string    `json:"description" db:"description"`
}
