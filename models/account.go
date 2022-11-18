package models

type Account struct {
	AccountId      int `json:"account_id" db:"account_id"`
	MainBalance    int `json:"main_balance" db:"main_balance"`
	ReserveBalance int `json:"reserve_balance" db:"reserve_balance"`
}
