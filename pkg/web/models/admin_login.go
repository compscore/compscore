package models

// AdminLogin is the JSON body for the admin login request
//
// @Summary body of admin login request
// @Description body of admin login request
// @Tags admin
type AdminLogin struct {
	Team string `json:"team"`
}
