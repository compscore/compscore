package cmd

import (
	"github.com/spf13/cobra"
)

var engineCmd = &cobra.Command{
	Use: "engine",
	Run: engineRun,
}

func engineRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	rootCmd.AddCommand(engineCmd)
}
