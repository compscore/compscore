package cmd

import (
	"github.com/compscore/compscore/pkg/grpc/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: serverRun,
}

func serverRun(cmd *cobra.Command, args []string) {
	server.Serve()
	defer server.Close()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
