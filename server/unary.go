package main

import (
	"context"

	pb "github.com/usama1031/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, reg *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
