package server

import (
	"context"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func (*compscoreServer_s) GetConfig(ctx context.Context, in *proto.GetConfigRequest) (*proto.GetConfigResponse, error) {
	logrus.Info("Received get config request")

	conf, err := yaml.Marshal(config.RunningConfig)
	if err != nil {
		return nil, err
	}

	return &proto.GetConfigResponse{Config: string(conf)}, nil
}
