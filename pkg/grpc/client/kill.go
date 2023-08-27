package client

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

func Kill(ctx context.Context) (string, error) {
	response, err := compscoreClient.Kill(ctx, &proto.KillRequest{})
	return response.GetMessage(), err
}
