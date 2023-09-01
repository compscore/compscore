package cmd

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/spf13/cobra"
)

var configRegenerateCmd = &cobra.Command{
	Use:     "regenerate",
	Aliases: []string{"regen"},
	Short:   "Regenerate running configuration",
	Long:    "Regenerate running configuration",
	Run:     configRegenerateRun,
}

func init() {
	configCmd.AddCommand(configRegenerateCmd)
}

func configRegenerateRun(cmd *cobra.Command, args []string) {
	config.RegenerateConfiguration()
}
