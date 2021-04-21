package service

import (
	"context"
	proto_generate "github.com/grpc-boot/service/proto-generate"
)

func (s *Server) SayHello(ctx context.Context, in *proto_generate.HelloRequest) (*proto_generate.HelloReply, error) {
	return &proto_generate.HelloReply{Message: "Hello " + in.Name}, nil
}
