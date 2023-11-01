package models

// Check is the struct used to marshal the JSON response of the check endpoint
//
// @Summary score check
// @Description score check
// @Tags check
type Check struct {
	Name  string `json:"name"`
	Edges struct {
		Status     []interface{} `json:"status,omitempty"`
		Credential []interface{} `json:"credential,omitempty"`
	} `json:"edges"`
}

// Credential is the struct used to marshal the JSON response of the credential endpoint
//
// @Summary credential of a check
// @Description credential of a check
// @Tags credential
type Credential struct {
	Password string `json:"password"`
	Edges    struct {
		Check interface{} `json:"check,omitempty"`
		Team  interface{} `json:"team,omitempty"`
	} `json:"edges"`
}

// Round is the struct used to marshal the JSON response of the round endpoint
//
// @Summary scoring round
// @Description scoring round
// @Tags round
type Round struct {
	Number   int  `json:"number"`
	Complete bool `json:"complete"`
	Edges    struct {
		Status []interface{} `json:"status,omitempty"`
	} `json:"edges"`
}

// Status is the struct used to marshal the JSON response of the status endpoint
//
// @Summary status of a check
// @Description status of a check
// @Tags status
type Status struct {
	Error  string `json:"error,omitempty"`
	Status string `json:"status"`
	Time   string `json:"time"`
	Points int    `json:"points"`
	Edges  struct {
		Check interface{} `json:"check,omitempty"`
		Team  interface{} `json:"team,omitempty"`
		Round interface{} `json:"round,omitempty"`
	} `json:"edges"`
}

// Team is the struct used to marshal the JSON response of the team endpoint
//
// @Summary team
// @Description team
// @Tags team
type Team struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name"`
	Roles  string `json:"roles"`
	Edges  struct {
		Status     []interface{} `json:"status,omitempty"`
		Credential []interface{} `json:"credential,omitempty"`
	} `json:"edges"`
}
