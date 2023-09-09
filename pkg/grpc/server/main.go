package server

import (
	"net"
	"os"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/engine"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type compscoreServer_s struct {
	proto.UnimplementedCompscoreServer
}

var (
	lis             net.Listener
	grpcServer      *grpc.Server
	compscoreServer compscoreServer_s = compscoreServer_s{}
	kill            chan struct{}     = make(chan struct{}, 1)
)

func Serve() {
	err := os.Remove(config.RunningConfig.Engine.Socket)
	if err != nil && !os.IsNotExist(err) {
		logrus.WithError(err).Fatal("Failed to remove existing socket")
	}

	_lis, err := net.Listen("unix", config.RunningConfig.Engine.Socket)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to listen on socket")
	}
	lis = _lis

	_grpcServer := grpc.NewServer()
	grpcServer = _grpcServer

	proto.RegisterCompscoreServer(grpcServer, &compscoreServer)

	logrus.Info("Starting gRPC server")

	closed := make(chan struct{})

	go handleClose(closed)

	err = grpcServer.Serve(lis)
	if err != nil && err != grpc.ErrServerStopped {
		logrus.WithError(err).Fatal("Failed to serve gRPC server")
	}
}

func handleClose(closed chan struct{}) {
	<-kill

	// Normal Close
	time.Sleep(time.Second)

	err := engine.Stop()
	if err != nil {
		logrus.WithError(err).Error("Failed to stop engine")
	}

	grpcServer.GracefulStop()

	// Force Close
	time.Sleep(time.Duration(config.RunningConfig.Scoring.Interval) + 1)

	if grpcServer != nil {
		grpcServer.Stop()
		if err != nil {
			logrus.WithError(err).Error("Failed to pause engine")
		}
	}

	// Force Exit
	time.Sleep(time.Duration(config.RunningConfig.Scoring.Interval) + 1)

	os.Exit(1)
}

func Close() {
	grpcServer.GracefulStop()
}

func ForceClose() {
	grpcServer.Stop()
}
