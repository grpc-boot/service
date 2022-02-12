#!/usr/bin/env bash
app_path=$(dirname $(cd $(dirname $0); pwd))
out_path=${app_path}/presentation/grpc
proto_path=${app_path}/presentation/grpc/proto

protoc -I=${proto_path} --go-grpc_out=${out_path} --go_out=${out_path} ${proto_path}/*.proto