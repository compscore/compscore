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
