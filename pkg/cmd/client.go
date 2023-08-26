package cmd

import (
	"context"
	"fmt"

	"github.com/compscore/compscore/pkg/grpc/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use: "client",
	Run: clientRun,
}

func clientRun(cmd *cobra.Command, args []string) {
	client.Open()
	defer client.Close()

	message, err := client.Ping.Ping(context.Background(), "Hello World!")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to ping server")
	}

	fmt.Println(message)
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
