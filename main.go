package main

import (
	_ "embed"
	"fmt"
	"runtime"

	"github.com/compscore/compscore/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.Execute()
}

//go:embed version
var version string

var gitCommit string
var gitBranch string
var buildDate string

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
	fmt.Println("Git Commit:", gitCommit)
	fmt.Println("Git Branch:", gitBranch)
	fmt.Println("Build Date:", buildDate)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("OS/Arch:", runtime.GOOS+"/"+runtime.GOARCH)
}
