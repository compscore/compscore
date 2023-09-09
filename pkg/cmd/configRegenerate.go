package cmd

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/helpers"
	"github.com/spf13/cobra"
)

var configRegenerateCmd = &cobra.Command{
	Use:     "regenerate",
	Aliases: []string{"regen", "r"},
	Short:   "Regenerate running configuration",
	Long:    "Regenerate running configuration",
	Run:     configRegenerateRun,
}

func init() {
	configCmd.AddCommand(configRegenerateCmd)
}

func configRegenerateRun(cmd *cobra.Command, args []string) {
	confirmed := helpers.UserConfirm("Regerating running configuration will overwrite any changes you have made to the running configuration (DO NOT DO THIS MID-COMPETITION).\nAre you sure you want to continue?")

	if confirmed {
		config.RegenerateConfiguration()
	}
}
