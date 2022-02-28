package main

import (
	"context"
	"time"

	"service/presentation/grpc/pb"

	"github.com/grpc-boot/base"
	"github.com/grpc-boot/betcd"
	clientv3 "go.etcd.io/etcd/client/v3"
	etcdresolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func callIndex(client pb.IndexClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := client.Index(ctx, &emptypb.Empty{})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			base.Red("client.Index(), code:%d msg:%s", st.Code(), st.Message())
			return
		}
		base.RedFatal("client.Index(), %v: ", err)
	}
	base.Green("Index: code:%d msg:%s", resp.Code, resp.Msg)
}

func callGw(client pb.GatewayClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := client.Gw(ctx, &emptypb.Empty{})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			base.Red("client.Gateway(), code:%d msg:%s", st.Code(), st.Message())
			return
		}
		base.RedFatal("client.Gateway(), %v: ", err)
	}
	base.Green("Gw: code:%d msg:%s data:%+v", resp.Code, resp.Msg, resp.Data)
}

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.16.49.131:2379"},
		DialTimeout: time.Second * 3,
	})

	if err != nil {
		base.RedFatal("init client err:%s", err.Error())
	}

	naming, err := betcd.NewNamingWithClient(client, "browser/service/app")
	if err != nil {
		base.RedFatal("init naming error:%s", err.Error())
	}

	rsl, err := etcdresolver.NewBuilder(client)
	if err != nil {
		base.RedFatal("init resolver err:%s", err.Error())
	}

	conn, err := naming.DialGrpc(grpc.WithUserAgent("grpc-boot"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithResolvers(rsl))
	if err != nil {
		base.RedFatal("connect grpc server err:%s", err.Error())
	}

	defer conn.Close()

	done := make(chan struct{}, 2)
	go func() {
		indexClient := pb.NewIndexClient(conn)
		for max := 10000; max > 0; max-- {
			callIndex(indexClient, time.Second*5)
		}
		done <- struct{}{}
	}()

	go func() {
		gwClient := pb.NewGatewayClient(conn)
		for max := 10000; max > 0; max-- {
			callGw(gwClient, time.Second)
		}

		done <- struct{}{}
	}()

	<-done
	<-done
}
