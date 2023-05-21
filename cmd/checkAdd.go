package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new check to the scoring engine",
	Args:  cobra.NoArgs,
	Run:   checkAdd,
}

func init() {
	checkCmd.AddCommand(checkAddCmd)
}

func checkAdd(cmd *cobra.Command, args []string) {
	fmt.Println("Adding a check")
}
