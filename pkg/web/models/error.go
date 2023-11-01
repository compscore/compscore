package models

// error_s is the struct used to marshal the JSON response of the login request
// @Summary response of login request
// @Description response of login request
// @Tags auth
type Error struct {
	Error string `json:"error"`
}
