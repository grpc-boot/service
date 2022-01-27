#!/usr/bin/env bash
app_name="service"
app_path=$(dirname $(cd $(dirname $0); pwd))
go=$GOROOT/bin/go
gin() {
  cd $app_path/presentation/gin
  $go build -o $"${app_path}/${app_name}-gin"
}

fasthttp() {
  cd $app_path/presentation/fasthttp
  $go build -o $"${app_path}/${app_name}-fasthttp"
}

grpc() {
  cd $app_path/presentation/grpc
  $go build -o $"${app_path}/${app_name}-grpc"
}

case "$1" in

gin)

gin

;;

fasthttp)

fasthttp

;;

grpc)

grpc

;;

*)

echo $"Usage: {gin|fasthttp|grpc|help}"

exit 1

esac