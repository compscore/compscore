package models

// CredentialEdit is the struct used to unmarshal the JSON body of the credential edit request
//
// @Summary body of credential edit request
// @Description body of credential edit request
// @Tags credential
type CredentialEdit struct {
	Password string `json:"password"`
}
