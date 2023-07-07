package main

import (
	pb "github.com/vinith/grpc-demo/proto"
	"golang.org/x/net/context"
)

func (s *helloServer) SaysHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
