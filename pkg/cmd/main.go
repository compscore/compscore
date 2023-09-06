package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compscore",
	Short: "Compscore is a scoring engine for Red/Blue Competitions",
	Long:  "Compscore is a scoring engine for Red/Blue Competitions",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the root command
func Execute() {
	rootCmd.Execute()
}
