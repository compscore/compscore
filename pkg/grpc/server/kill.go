package server

import (
	"context"

	"github.com/compscore/compscore/pkg/engine"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (*compscoreServer_s) Kill(ctx context.Context, in *proto.KillRequest) (*proto.KillResponse, error) {
	logrus.Info("Received kill request")

	engine.Status = proto.StatusEnum_UNKNOWN
	kill <- struct{}{}

	return &proto.KillResponse{Message: "killing"}, nil
}
