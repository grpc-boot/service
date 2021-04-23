package main

import (
	"github.com/grpc-boot/service/conf"
	"log"
	"net"

	"github.com/grpc-boot/service/proto/pb"
	"github.com/grpc-boot/service/service"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	conf.Init()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen： %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &service.Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}
