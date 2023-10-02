package structs

type Check struct {
	Name   string `json:"name"`
	Status []int  `json:"status"`
}

type Scoreboard struct {
	Round  int     `json:"round"`
	Scores []int   `json:"scores"`
	Checks []Check `json:"checks"`
}

type TeamScoreboard struct {
	Round  int     `json:"round"`
	Checks []Check `json:"checks"`
}

type CheckScoreboard struct {
	Round  int     `json:"round"`
	Checks []Check `json:"checks"`
}
