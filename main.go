package main

import (
	_ "embed"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/compscore/compscore/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.Execute()
}

//go:embed version
var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays current version of Compscore",
	Args:  cobra.NoArgs,
	Run:   versionCmdRun,
}

func init() {
	cmd.RootCmd.AddCommand(versionCmd)
}

func versionCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("Version:", version)
	fmt.Println("Git Commit:", getGitCommit())
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("OS/Arch:", runtime.GOOS+"/"+runtime.GOARCH)
}

func getGitCommit() string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error getting git commit:", err)
		return ""
	}

	return strings.TrimSpace(string(output))
}
