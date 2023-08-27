package client

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

func Start(ctx context.Context) (string, error) {
	response, err := compscoreClient.Start(ctx, &proto.StartRequest{})
	return response.GetMessage(), err
}
