package domain

type Transaction struct {
	Id       string   `json:"id"`
	Customer Customer `json:"customer"`
	Merchant Merchant `json:"merchant"`
	Amount   int      `json:"amount"`
}

// ID implements simdb.Entity.
func (t *Transaction) ID() (jsonField string, value interface{}) {
	value = t.Id
	jsonField = "id"
	return
}
