package models

type Check struct {
	Name  string `json:"name"`
	Edges struct {
		Status     []interface{} `json:"status,omitempty"`
		Credential []interface{} `json:"credential,omitempty"`
	} `json:"edges"`
}
