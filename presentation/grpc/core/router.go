package core

import (
	"google.golang.org/grpc"

	v1 "service/presentation/grpc/api/v1"
)

func Route(server *grpc.Server) {
	bind(server)

	return
}

func bind(server *grpc.Server) {
	v1.RegisterIndex(server)
	v1.RegisterGateway(server)
}
