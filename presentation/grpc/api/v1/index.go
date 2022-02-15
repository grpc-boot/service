package v1

import (
	"context"

	"service/common/define/constant"
	"service/presentation/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type index struct {
	pb.UnimplementedIndexServer
}

func (i *index) Index(ctx context.Context, in *emptypb.Empty) (*pb.IndexReply, error) {
	data := &pb.IndexReply{
		Code: constant.Success,
		Msg:  "Hello World!",
	}

	gateway(ctx, data, constant.Success)
	return data, nil
}

func RegisterIndex(server *grpc.Server) {
	pb.RegisterIndexServer(server, &index{})
}
