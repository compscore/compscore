package client

import (
	"context"
	"time"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn            *grpc.ClientConn
	compscoreClient proto.CompscoreClient
)

func Open() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	_conn, err := grpc.DialContext(ctx, "unix:"+config.RunningConfig.Engine.Socket, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to gRPC server")
	}
	cancel()

	conn = _conn

	compscoreClient = proto.NewCompscoreClient(conn)
}

func Close() {
	err := conn.Close()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to close gRPC connection")
	}
}
