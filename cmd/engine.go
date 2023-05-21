package cmd

import (
	"github.com/spf13/cobra"
)

var engineCmd = &cobra.Command{
	Use:   "engine",
	Short: "Subcommand for engine management",
	Args:  cobra.ArbitraryArgs,
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func init() {
	RootCmd.AddCommand(engineCmd)
}
