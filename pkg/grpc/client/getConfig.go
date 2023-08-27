package client

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

func GetConfig(ctx context.Context) (string, error) {
	response, err := compscoreClient.GetConfig(ctx, &proto.GetConfigRequest{})
	return response.GetConfig(), err
}
