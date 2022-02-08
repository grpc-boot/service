package model

import "github.com/grpc-boot/base"

// SnowFlake 雪花算法配置
type SnowFlake struct {
	Mode  uint8  `json:"mode" yaml:"mode"`   //时钟回拨处理模式，1等待5ms，2使用最大时间，3抛出异常
	Begin string `json:"begin" yaml:"begin"` //开始时间，格式如：2022-01-01
}

// Build 初始化
func (sf *SnowFlake) Build() {
	if sf.Mode == 0 {
		sf.Mode = base.ModeWait
	}

	if sf.Begin == "" {
		sf.Begin = `2022-01-01`
	}
}
