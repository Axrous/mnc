package domain

type Merchant struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// ID implements simdb.Entity.
func (m *Merchant) ID() (jsonField string, value interface{}) {
	value = m.Id
	jsonField = "id"
	return
}