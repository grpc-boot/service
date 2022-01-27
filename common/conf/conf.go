package conf

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/grpc-boot/base"
	_ "github.com/grpc-boot/betcd"
	_ "github.com/grpc-boot/gateway"
)

const (
	JsonConfig = `conf/app.json`
	YamlConfig = `conf/app.yml`
)

var (
	conf Conf
)

// GetConf 获取配置
func GetConf() *Conf {
	return &conf
}

// Init 初始化配置
func Init(filename string) error {
	var (
		kind     = ".yml"
		rootPath = filepath.Dir(os.Args[0])
		position int
	)

	if filename == "" {
		confFile := rootPath + "/" + YamlConfig
		if !base.FileExists(confFile) {
			kind = ".json"
			filename = rootPath + "/" + JsonConfig
		}
	} else {
		if filename[0] != '/' {
			filename = rootPath + "/" + filename
		}

		// 解析文件扩展名
		position = strings.LastIndexByte(filename, '.')
		if position > 0 {
			kind = filename[position:]
		}
	}

	switch kind {
	case ".yml", ".yaml":
		return base.YamlDecodeFile(filename, &conf)
	case ".json":
		return base.JsonDecodeFile(filename, &conf)
	}

	return errors.New("仅支持扩展名为[.yml,.yaml,.json]的配置文件")
}

// InitWithFlag with flag初始化配置
func InitWithFlag() error {
	var (
		filename string
	)

	flag.StringVar(&filename, "c", "", "配置文件路径，默认为conf/app.yml或conf/app.json")
	flag.Parse()

	return Init(filename)
}
