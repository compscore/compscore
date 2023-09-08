package server

import (
	"context"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"gopkg.in/yaml.v3"
)

func (*compscoreServer_s) GetConfig(ctx context.Context, in *proto.GetConfigRequest) (*proto.GetConfigResponse, error) {
	confBytes, err := yaml.Marshal(config.RunningConfig)
	if err != nil {
		return nil, err
	}
	confString := string(confBytes)
	return &proto.GetConfigResponse{Config: confString}, nil
}
