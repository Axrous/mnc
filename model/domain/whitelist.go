package domain

type Whitelist struct {
	Id string `json:"id"`
}

// ID implements simdb.Entity.
func (w *Whitelist) ID() (jsonField string, value interface{}) {
	value = w.Id
	jsonField = "id"
	return
}
