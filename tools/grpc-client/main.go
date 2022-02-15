package main

import (
	"context"
	"time"

	"service/presentation/grpc/pb"

	"github.com/grpc-boot/base"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func callIndex(client pb.IndexClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := client.Index(ctx, &emptypb.Empty{})
	if err != nil {
		base.RedFatal("client.Index(), %v: ", err)
	}
	base.Green("Index: code:%d msg:%s", resp.Code, resp.Msg)
}

func callGw(client pb.GatewayClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err := client.Gw(ctx, &emptypb.Empty{})
	if err != nil {
		base.RedFatal("client.Gateway(), %v: ", err)
	}
	base.Green("Gw: code:%d msg:%s data:%+v", resp.Code, resp.Msg, resp.Data)
}

func main() {
	opts := []grpc.DialOption{
		grpc.WithUserAgent("grpc-boot"),
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
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
