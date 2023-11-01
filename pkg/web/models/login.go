package models

// login_s is the struct used to unmarshal the JSON body of the login request
// @Summary body of login request
// @Description body of login request
// @Tags auth
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
