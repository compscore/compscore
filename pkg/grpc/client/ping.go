package client

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

type pingClient_s struct{}

var (
	Ping       pingClient_s
	pingClient proto.PingSeviceClient
)

func (c pingClient_s) Ping(ctx context.Context, message string) (string, error) {
	response, err := pingClient.Ping(ctx, &proto.PingRequest{Message: message})
	if err != nil {
		return "", err
	}

	return response.GetMessage(), nil
}
