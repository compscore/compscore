package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Version   string
	GitCommit string
	GitBranch string
	BuildDate string
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Run:     versionRun,
}

// version command for running "compscore version"
func versionRun(cmd *cobra.Command, args []string) {
	fmt.Println("Version:", Version)

	if GitCommit == "" {
		GitCommit = gitCommit()
	}
	fmt.Println("Git Commit:", GitCommit)

	if GitBranch == "" {
		GitBranch = gitBranch()
	}
	fmt.Println("Git Branch:", GitBranch)

	if BuildDate == "" {
		BuildDate = buildDate()
	}
	fmt.Println("Build Date:", BuildDate)

	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("OS/Arch:", runtime.GOOS+"/"+runtime.GOARCH)
}

// gitCommit gets the current git commit hash
func gitCommit() string {
	command, _ := exec.Command("git", "rev-parse", "HEAD").Output()
	return strings.TrimSpace(string(command))
}

// gitBranch gets the current git branch
func gitBranch() string {
	command, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.TrimSpace(string(command))
}

// buildDate gets the current build date
func buildDate() string {
	command, _ := exec.Command("date", "+%Y-%m-%d").Output()
	return strings.TrimSpace(string(command))
}
