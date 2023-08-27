package server

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (*compscoreServer_s) Heartbeat(ctx context.Context, in *proto.HeartbeatRequest) (*proto.HeartbeatResponse, error) {
	logrus.Info("Received heartbeat request")
	return &proto.HeartbeatResponse{}, nil
}
