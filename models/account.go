package models

type Account struct {
	AccountId      int `json:"account_id"`
	MainBalance    int `json:"main_balance"`
	ReserveBalance int `json:"reserve_balance"`
}
