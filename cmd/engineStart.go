package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var engineStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the scoring engine",
	Args:  cobra.NoArgs,
	Run:   engineStart,
}

func init() {
	engineCmd.AddCommand(engineStartCmd)
}

func engineStart(cmd *cobra.Command, args []string) {
	fmt.Println("Starting engine")
}
