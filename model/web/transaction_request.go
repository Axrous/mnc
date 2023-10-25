package web

type TransactionPaymentRequest struct {
	MerchantId string `json:"merchant_id"`
	Amount     int    `json:"amount"`
}