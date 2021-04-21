package main

import (
	"context"
	proto_generate "github.com/grpc-boot/service/proto-generate"
	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect： %v", err)
	}

	defer conn.Close()
	c := proto_generate.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto_generate.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet： %v", err)
	}
	log.Printf("Greeting： %s", r.Message)
}
