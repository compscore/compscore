package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var engineStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the scoring engine",
	Args:  cobra.NoArgs,
	Run:   engineStop,
}

func init() {
	engineCmd.AddCommand(engineStopCmd)
}

func engineStop(cmd *cobra.Command, args []string) {
	fmt.Println("Stopping engine")
}
