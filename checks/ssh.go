package checks

import "github.com/compscore/compscore/structs"

func sshFunction(auth structs.Auth) bool {
	return true
}

var (
	sshDescription string = "Runs a basic ssh check"
	sshName        string = "SSH"
)

var sshCheck structs.Scorecheck = structs.Scorecheck{
	Name:        sshName,
	Description: sshDescription,
	Function:    sshFunction,
}
