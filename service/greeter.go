package service

import (
	"context"

	"github.com/grpc-boot/service/proto/pb"
)

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
