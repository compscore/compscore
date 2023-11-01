package models

// Status is the struct used to marshal the JSON response of the status endpoint
// @Summary status of the engine
// @Description status of the engine
// @Tags engine
type EngineStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
