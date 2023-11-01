package models

type Check struct {
	Name  string `json:"name"`
	Edges struct {
		Status     []interface{} `json:"status,omitempty"`
		Credential []interface{} `json:"credential,omitempty"`
	} `json:"edges"`
}

type Credential struct {
	Password string `json:"password"`
	Edges    struct {
		Check interface{} `json:"check,omitempty"`
		Team  interface{} `json:"team,omitempty"`
	} `json:"edges"`
}

type Round struct {
	Number   int  `json:"number"`
	Complete bool `json:"complete"`
	Edges    struct {
		Status []interface{} `json:"status,omitempty"`
	} `json:"edges"`
}
