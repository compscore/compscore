package structs

// Check is a struct for a check
//
// @Summary Check
// @Description Check
// @Tags scoreboard
type Check struct {
	Name   string `json:"name"`
	Status []int  `json:"status"`
}

// Scoreboard is a struct for the main scoreboard
//
// @Summary Main Scoreboard
// @Description Main Scoreboard
// @Tags scoreboard
type Scoreboard struct {
	Round  int     `json:"round"`
	Scores []int   `json:"scores"`
	Checks []Check `json:"checks"`
}

// TeamScoreboard is a struct for the scoreboard for a given team
//
// @Summary Team Scoreboard
// @Description Team Scoreboard
// @Tags scoreboard
type TeamScoreboard struct {
	Round  int     `json:"round"`
	Checks []Check `json:"checks"`
}

// CheckScoreboard is a struct for the scoreboard for a given check
//
// @Summary Check Scoreboard
// @Description Check Scoreboard
// @Tags scoreboard
type CheckScoreboard struct {
	Round int     `json:"round"`
	Teams []Check `json:"teams"`
}

// Status is a struct for a status entry
//
// @Summary Status
// @Description Status
// @Tags status
type Status struct {
	Error  string `json:"error"`
	Time   string `json:"time"`
	Status int    `json:"status"`
	Round  int    `json:"round"`
}
