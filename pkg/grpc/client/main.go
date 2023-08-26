package client

import (
	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn *grpc.ClientConn
)

func Open() {
	_conn, err := grpc.Dial("unix:"+config.UnixSocket, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to gRPC server")
	}

	conn = _conn

	pingClient = proto.NewPingSeviceClient(conn)
	statusClient = proto.NewStatusServiceClient(conn)
}

func Close() {
	err := conn.Close()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to close gRPC connection")
	}
}
