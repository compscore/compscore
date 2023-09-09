package server

import (
	"context"

	"github.com/compscore/compscore/pkg/engine"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (*compscoreServer_s) Start(ctx context.Context, in *proto.StartRequest) (*proto.StartResponse, error) {
	logrus.Info("Received start request")

	err := engine.Run()
	if err != nil {
		return &proto.StartResponse{Message: err.Error()}, nil
	}
	return &proto.StartResponse{Message: "started"}, nil
}
