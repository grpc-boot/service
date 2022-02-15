package v1

import (
	"context"

	"service/common/components"
	"service/common/define/constant"
	"service/presentation/grpc/pb"
	"service/presentation/grpc/utils"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type gwServer struct {
	pb.UnimplementedGatewayServer
}

// Gw 获取网关信息
func (g *gwServer) Gw(ctx context.Context, in *emptypb.Empty) (*pb.GwReply, error) {
	gw, ok := components.GetGateway(constant.ContGateway)
	if !ok {
		code := constant.ErrGatewayNotExists
		gateway(ctx, nil, code)
		return nil, components.Code2Error(code)
	}

	info := gw.Info()
	data := utils.ConvertGatewayInfo(&info)

	gateway(ctx, data, data.Code)
	return data, nil
}

func RegisterGateway(server *grpc.Server) {
	pb.RegisterGatewayServer(server, &gwServer{})
}
