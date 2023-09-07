package server

import (
	"context"

	"github.com/compscore/compscore/pkg/engine"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (*compscoreServer_s) Status(ctx context.Context, in *proto.StatusRequest) (*proto.StatusResponse, error) {
	logrus.Info("Received status request")
	return &proto.StatusResponse{Message: "Status Response", Status: engine.Status}, nil
}
