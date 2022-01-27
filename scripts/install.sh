#!/usr/bin/env bash
app_path=$(dirname $(cd $(dirname $0); pwd))

case "$1" in
gin)
app="gin"
;;
fasthttp)
app="fasthttp"
;;
grpc)
app="grpc"
;;
*)
echo $"Usage: {gin|fasthttp|grpc} {test|product}"
exit 1
esac

case "$2" in
test)
env="test"
;;
product)
env="product"
;;
*)
echo $"Usage: {gin|fasthttp|grpc} {test|product}"
exit 1
esac

sh $app_path/scripts/build.sh $app

yaml_conf="$app_path/presentation/$app/conf/$env.yml"
json_conf="$app_path/presentation/$app/conf/$env.json"

if [ ! -f "$yaml_conf" ]; then
  cp -rf $json_conf "$app_path/conf/app.json"
  exit 0
fi

cp -rf $yaml_conf "$app_path/conf/app.yml"