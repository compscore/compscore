package cmd

import (
	"context"
	"fmt"
	"plugin"

	"github.com/compscore/compscore/pkg/helpers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var loadCheckCmd = &cobra.Command{
	Use:   "load",
	Short: "Load check",
	Long:  `Load check`,
	Run:   loadCheckRun,
}

func init() {
	rootCmd.AddCommand(loadCheckCmd)
}

func loadCheckRun(cmd *cobra.Command, args []string) {
	file, err := helpers.GetReleaseAsset("compscore", "check-template", "")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to download release asset")
	}

	plugin, err := plugin.Open(file)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to open plugin")
	}

	runSymbol, err := plugin.Lookup("Run")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to lookup Run")
	}

	runFunc, ok := runSymbol.(func(ctx context.Context, target string, command string, expectedOutput string, username string, password string) (bool, string))
	if !ok {
		logrus.Fatal("Failed to cast Run to func")
	}

	fmt.Println(runFunc(context.Background(), "localhost", "echo", "hello world", "", ""))
}
