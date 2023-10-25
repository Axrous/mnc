package domain

type Customer struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ID implements simdb.Entity.
func (c *Customer) ID() (jsonField string, value interface{}) {
	value = c.Id
	jsonField = "id"
	return
}
