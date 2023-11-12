package server

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/gql/graph"
	"golang.org/x/term"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Run the server",
	Long:    "Run the server",
	Aliases: []string{"s", "serve"},

	Run: run,
}

// serverRun runs the server
func run(cmd *cobra.Command, args []string) {
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	go func() {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		logrus.Printf("Starting server on http://localhost:%d", config.Port)
		err := http.ListenAndServe(
			fmt.Sprintf(":%d", config.Port),
			nil,
		)
		if err != nil {
			logrus.WithError(err).Fatal("failed to start server")
		} else {
			logrus.Info("Server stopped")
		}
	}()

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		logrus.WithError(err).Fatal("failed to make terminal raw")
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Press 'q' to stop the server\r")
	fmt.Println("Press 'o' to open the browser\r")
	fmt.Println("\r")
	fmt.Println("===== LOGS =====\r")

	keystrokeChan := make(chan rune)
	defer close(keystrokeChan)

	go func() {
		for {
			// Read a single character
			buf := make([]byte, 1)
			_, err = os.Stdin.Read(buf)
			if err != nil {
				logrus.WithError(err).Fatal("failed to read from stdin")
			}

			// Send the character to the keystroke channel
			keystrokeChan <- rune(buf[0])
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	for {
		select {
		case keystroke := <-keystrokeChan:
			if keystroke == 'q' || keystroke == 'Q' {
				logrus.Info("Stopping server")
				return
			}

			if keystroke == 'o' || keystroke == 'O' {
				err := open(fmt.Sprintf("http://localhost:%d", config.Port))
				if err != nil {
					logrus.WithError(err).Error("failed to open browser")
				}
			}
		case <-sigs:
			logrus.Info("Stopping server")
			return
		}
	}
}

// open opens the given url in the default browser
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
