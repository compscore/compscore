package main

import (
	_ "embed"
	"fmt"

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
	fmt.Println(version)
}
