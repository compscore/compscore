package cmd

import (
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Subcommand for check management",
	Args:  cobra.ArbitraryArgs,
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func init() {
	RootCmd.AddCommand(checkCmd)
}
