package structs

type Check struct {
	Name     string `json:"name"`
	Statuses []int  `json:"teams"`
}

type Scoreboard struct {
	Round  int     `json:"round"`
	Scores []int   `json:"scores"`
	Checks []Check `json:"checks"`
}
