package models

// cookie_s is the struct used to marshal the JSON response of the login request
// @Summary response of login request
// @Description response of login request
// @Tags auth
type Cookie struct {
	Name       string `json:"name"`
	Token      string `json:"token"`
	Expiration int    `json:"expiration"`
	Path       string `json:"path"`
	Domain     string `json:"domain"`
	Secure     bool   `json:"secure"`
	HttpOnly   bool   `json:"httponly"`
}
