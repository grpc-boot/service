package main

import (
	proto_generate "github.com/grpc-boot/service/proto-generate"
	"github.com/grpc-boot/service/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen： %v", err)
	}
	s := grpc.NewServer()
	proto_generate.RegisterGreeterServer(s, &service.Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}
