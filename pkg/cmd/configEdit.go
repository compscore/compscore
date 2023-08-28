package cmd

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/helpers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the config file",
	Long:  "Edit the config file",
	Run:   configEditRun,
}

var editor string

func init() {
	configEditCmd.Flags().StringVarP(&editor, "editor", "e", "", "Editor to use")
	configCmd.AddCommand(configEditCmd)
}

func configEditRun(cmd *cobra.Command, args []string) {
	if editor == "" {
		editor = helpers.GetEditor()
	}

	if editor == "" {
		logrus.Fatal("No editor found, try using -e/--editor to specify one")
	}

	err := helpers.EditFile(config.ConfigFile, editor)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to edit config file")
	}

	logrus.Info("Config file updated and loaded")
}
