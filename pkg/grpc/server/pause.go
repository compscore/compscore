package server

import (
	"context"

	"github.com/compscore/compscore/pkg/engine"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (*compscoreServer_s) Pause(ctx context.Context, in *proto.PauseRequest) (*proto.PauseResponse, error) {
	logrus.Info("Received pause request")

	err := engine.Pause()
	if err != nil {
		return &proto.PauseResponse{Message: err.Error()}, nil
	}
	return &proto.PauseResponse{Message: "pausing"}, nil
}
