package client

import (
	"context"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

func Heartbeat(ctx context.Context) error {
	_, err := compscoreClient.Heartbeat(ctx, &proto.HeartbeatRequest{})
	return err
}
