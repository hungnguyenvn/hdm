package hdm

import (
	"context"
	"hdm/pb"
)

func NewGRPCHandler() pb.HDMServer {
	handler := grpcHandler{}
	return handler
}

type grpcHandler struct {
}

func (grpcHandler) Hotels(ctx context.Context, request *pb.HotelsRequest) (*pb.HotelsResponse, error) {
	return &pb.HotelsResponse{Msg: request.Msg + "____ response"}, nil
}
