package models

// AdminLogin is the JSON body for the admin password reset request
//
// @Summary body of admin password reset request
// @Description body of admin password reset request
// @Tags admin
type AdminPasswordReset struct {
	Team     string `json:"team"`
	Password string `json:"password"`
}
