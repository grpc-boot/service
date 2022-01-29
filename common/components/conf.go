package components

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"strings"

	"service/common/define/constant"
	"service/common/model"

	"github.com/grpc-boot/base"
	_ "github.com/grpc-boot/betcd"
	_ "github.com/grpc-boot/gateway"
)

var (
	conf model.Conf
)

// GetConf 获取配置
func GetConf() *model.Conf {
	return &conf
}

// InitConf 初始化配置
func InitConf(filename string) error {
	var (
		kind     = ".yml"
		rootPath = filepath.Dir(os.Args[0]) + "/"
		position int
	)

	if filename == "" {
		confFile := rootPath + constant.ConfigYaml
		if !base.FileExists(confFile) {
			kind = ".json"
			filename = rootPath + constant.ConfJson
		} else {
			filename = confFile
		}
	} else {
		if filename[0] != '/' {
			filename = rootPath + filename
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

// InitConfWithFlag with flag初始化配置
func InitConfWithFlag() error {
	var (
		filename string
	)

	flag.StringVar(&filename, "c", "", "配置文件路径，默认为conf/app.yml或conf/app.json")
	flag.Parse()

	return InitConf(filename)
}
