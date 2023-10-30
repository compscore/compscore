package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/data"
	"github.com/compscore/compscore/pkg/grpc/server"
	"github.com/compscore/compscore/pkg/web"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Run the server",
	Long:    "Run the server",
	Aliases: []string{"s", "serve"},

	Run: serverRun,
}

// serverRun runs the server
func serverRun(cmd *cobra.Command, args []string) {
	config.Init()
	data.Init()

	go web.Start()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	go server.Serve()

	// close procedure
	force := false
	for {
		<-sigChannel
		if force {
			os.Exit(1)
		} else {
			force = true
			server.Close()
		}
	}
}
