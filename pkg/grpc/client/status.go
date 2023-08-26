package client

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

type statusClient_s struct{}

var (
	Status       statusClient_s
	statusClient proto.StatusServiceClient
)

func (statusClient_s) Status(ctx context.Context) (proto.StatusEnum, string, error) {
	response, err := statusClient.Status(ctx, &proto.StatusRequest{})
	return response.GetStatus(), response.GetMessage(), err
}
