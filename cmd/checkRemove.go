package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removw a check to the scoring engine",
	Args:  cobra.NoArgs,
	Run:   checkRemove,
}

func init() {
	checkCmd.AddCommand(checkRemoveCmd)
}

func checkRemove(cmd *cobra.Command, args []string) {
	fmt.Println("Removing a check")
}
