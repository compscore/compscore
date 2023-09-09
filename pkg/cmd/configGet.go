package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g", "show", "read"},
	Short:   "Get current configuration",
	Long:    "Get current configuration",
	Run:     configGetRun,
}

func init() {
	configCmd.AddCommand(configGetCmd)
}

func configGetRun(cmd *cobra.Command, args []string) {
	config.Init()

	client.Open()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.RunningConfig.Engine.Timeout)*time.Second)
	defer cancel()

	conf, err := client.GetConfig(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("Error getting configuration")
	}

	fmt.Println(conf)
}
