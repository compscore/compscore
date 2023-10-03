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
	Round int     `json:"round"`
	Teams []Check `json:"teams"`
}

type Status struct {
	Error  string `json:"error"`
	Time   string `json:"time"`
	Status int    `json:"status"`
}

type HistoryStatus struct {
	Team    string   `json:"team"`
	Round   int      `json:"round"`
	Check   string   `json:"check"`
	History []Status `json:"history"`
}
