package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/compscore/compscore/pkg/grpc/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Run the server",
	Long:    "Run the server",
	Aliases: []string{"s", "serve"},

	Run: serverRun,
}

func serverRun(cmd *cobra.Command, args []string) {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		force := false
		<-sigChannel
		if force {
			server.ForceClose()
		} else {
			force = true
			server.Close()
		}
	}()

	server.Serve()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
