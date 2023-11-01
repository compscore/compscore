package main

import (
	_ "embed"

	"github.com/compscore/compscore/pkg/cmd"
)

//go:embed version
var version string

var gitCommit string
var gitBranch string
var buildDate string

func init() {
	cmd.Version = version
	cmd.GitCommit = gitCommit
	cmd.GitBranch = gitBranch
	cmd.BuildDate = buildDate
}

func main() {
	cmd.Execute()
}

// @title Compscore API
//
// @version 1.0
// @BasePath /api
//
// @description This is the API for the Compscore application
//
// @contact.name 1nv8rzim
// @contact.url https://github.com/compsore/compscore/issues
//
// @securityDefinitions.apiKey ServiceAuth
// @tokenUrl /api/login
// @in cookie
// @name auth
// @description JWT for authentication
