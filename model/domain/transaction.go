package domain

type Transaction struct {
	Id       string   `json:"id"`
	Customer Customer `json:"customer"`
	Merchant Merchant `json:"merchant"`
	Amount   int      `json:"amount"`
}