package structs

// Scorecheck store the state of a score check and can be used to run a scorecheck on a host
type Scorecheck struct {
	Name        string
	Description string
	Function    func(Auth) bool
}

// Runs a single instance of a scorecheck
func (s *Scorecheck) Run(auth Auth) bool {
	return s.Function(auth)
}
