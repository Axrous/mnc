package web

import "github.com/Axrous/mnc/model/domain"

type TransactionResponse struct {
	Id       string      `json:"id"`
	Customer domain.Customer `json:"customer"`
	Merchant domain.Merchant    `json:"merchant"`
	Amount   int         `json:"amount"`
}