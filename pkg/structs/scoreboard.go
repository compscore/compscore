package structs

type Check struct {
	Name  string `json:"name"`
	Teams []int  `json:"teams"`
}

type Scoreboard struct {
	Round  int     `json:"round"`
	Checks []Check `json:"checks"`
}
