package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var engineStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check status of scoring engine",
	Args:  cobra.NoArgs,
	Run:   engineStatus,
}

func init() {
	engineCmd.AddCommand(engineStatusCmd)
}

func engineStatus(cmd *cobra.Command, args []string) {
	fmt.Println("Checking status of engine")
}
