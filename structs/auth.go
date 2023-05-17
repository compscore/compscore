package structs

// Auth is passed to the scorecheck and stores all authentication parameters that ScoreCheck will need
type Auth struct {
	Host     string
	Port     int
	Username string
	Password string
}
