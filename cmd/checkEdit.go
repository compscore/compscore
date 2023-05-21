package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a check to the scoring engine",
	Args:  cobra.NoArgs,
	Run:   checkEdit,
}

func init() {
	checkCmd.AddCommand(checkEditCmd)
}

func checkEdit(cmd *cobra.Command, args []string) {
	fmt.Println("Editting a check")
}
