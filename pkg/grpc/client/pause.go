package client

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

func Pause(ctx context.Context) (string, error) {
	response, err := compscoreClient.Pause(ctx, &proto.PauseRequest{})
	return response.GetMessage(), err
}
