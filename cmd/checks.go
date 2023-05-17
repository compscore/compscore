package cmd

import (
	"fmt"

	"github.com/compscore/compscore/checks"
	"github.com/spf13/cobra"
)

var checksCmd = &cobra.Command{
	Use:   "checks",
	Short: "Displays all available checks",
	Args:  cobra.NoArgs,
	Run:   versionCmdRun,
}

func init() {
	RootCmd.AddCommand(checksCmd)
}

func versionCmdRun(cmd *cobra.Command, args []string) {
	longestCheckName := len("name")
	for _, check := range checks.Checks {
		if len(check.Name) > longestCheckName {
			longestCheckName = len(check.Name)
		}
	}

	fmt.Printf("%*s : %s\n\n", longestCheckName, "Name", "Description")
	for _, check := range checks.Checks {
		fmt.Printf("%*s : %s\n", longestCheckName, check.Name, check.Description)
	}
}
