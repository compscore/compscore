package server

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

type pingServer_s struct {
	proto.UnimplementedPingSeviceServer
}

func (s *pingServer_s) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PongResponse, error) {
	logrus.Info("Received message: ", in.GetMessage())
	return &proto.PongResponse{Message: in.GetMessage()}, nil
}
