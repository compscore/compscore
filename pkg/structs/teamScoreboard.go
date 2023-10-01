package structs

type TeamCheck struct {
	Name   string `json:"name"`
	Status []int  `json:"status"`
}

type TeamScoreboard struct {
	Round  int         `json:"round"`
	Checks []TeamCheck `json:"checks"`
}
