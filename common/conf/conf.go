package conf

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/grpc-boot/base"
	_ "github.com/grpc-boot/betcd"
	_ "github.com/grpc-boot/gateway"
	"github.com/grpc-boot/gedis"
	"github.com/grpc-boot/orm"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	ConfJson = `json`
	ConfYaml = `yml`
)

const (
	JsonConfig = `conf/app.json`
	YamlConfig = `conf/app.yml`
)

var (
	conf Conf
)

// Conf 配置
type Conf struct {
	App   App                 `json:"app" yaml:"app"`
	Db    orm.GroupOption     `json:"db" yaml:"db"`
	Redis []gedis.GroupOption `json:"redis" yaml:"redis"`
	Etcd  clientv3.Config     `json:"etcd" yaml:"etcd"`
}

func initJsonConf(filename string) error {
	if filename == "" {
		filename = JsonConfig
	}
	return base.JsonDecodeFile(filename, &conf)
}

func initYamlConf(filename string) error {
	if filename == "" {
		filename = YamlConfig
	}
	return base.YamlDecodeFile(filename, &conf)
}

func Init(kind string, filename string) error {
	if filename != "" && filename[0] != '/' {
		filename = filepath.Dir(os.Args[0]) + "/" + filename
	}

	switch kind {
	case ConfYaml:
		return initYamlConf(filename)
	}
	return initJsonConf(filename)
}

func InitWithFlag() error {
	var (
		kind     string
		filename string
	)

	flag.StringVar(&kind, "k", "yml", "配置文件类型，默认为yml")
	flag.StringVar(&filename, "c", "", "配置文件路径，默认为conf/app.yml或conf/app.json")
	flag.Parse()

	return Init(kind, filename)
}

func GetConf() *Conf {
	return &conf
}
