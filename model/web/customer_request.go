package web

type CustomerCreateRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomerLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}