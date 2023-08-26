package server

import (
	"context"
	"fmt"

	"github.com/compscore/compscore/pkg/grpc/proto"
)

type pingServer_s struct {
	proto.UnimplementedPingSeviceServer
}

func (s *pingServer_s) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PongResponse, error) {
	fmt.Println("Received message:", in.GetMessage())
	return &proto.PongResponse{Message: in.GetMessage()}, nil
}
