package cmd

import "github.com/spf13/cobra"

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"conf", "c"},
	Short:   "Manage compscore configuration",
	Long:    "Manage compscore configuration",
	Run:     configRun,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func configRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}
