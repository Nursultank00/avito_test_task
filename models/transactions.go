package models

import "time"

type Transactions struct {
	TransactionId int       `json:"transasction_id"`
	ServiceId     int       `json:"service_id"`
	OrderId       int       `json:"order_id"`
	Amount        int       `json:"amount"`
	Datetime      time.Time `json:"datetime"`
}
