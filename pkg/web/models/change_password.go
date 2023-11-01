package models

// change_password_s is the struct used to unmarshal the JSON body of the change password request
// @Summary body of change password request
// @Description body of change password request
// @Tags auth
type ChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
