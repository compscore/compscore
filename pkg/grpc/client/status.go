package client

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

func Status(ctx context.Context) (proto.StatusEnum, string, error) {
	response, err := compscoreClient.Status(ctx, &proto.StatusRequest{})
	return response.GetStatus(), response.GetMessage(), err
}
