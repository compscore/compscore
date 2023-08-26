package server

import (
	"net"
	"os"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	lis        net.Listener
	grpcServer *grpc.Server
)

func Serve() {
	err := os.Remove(config.UnixSocket)
	if err != nil && !os.IsNotExist(err) {
		logrus.WithError(err).Fatal("Failed to remove existing socket")
	}

	_lis, err := net.Listen("unix", config.UnixSocket)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to listen on socket")
	}
	lis = _lis

	_grpcServer := grpc.NewServer()
	grpcServer = _grpcServer

	proto.RegisterPingSeviceServer(grpcServer, &pingServer_s{})
	proto.RegisterStatusServiceServer(grpcServer, &statusServer_s{})

	logrus.Info("Serving gRPC server")

	err = grpcServer.Serve(lis)
	if err != nil && err != grpc.ErrServerStopped {
		logrus.WithError(err).Fatal("Failed to serve gRPC server")
	}
}

func Close() {
	grpcServer.GracefulStop()
}

func ForceClose() {
	grpcServer.Stop()
}
