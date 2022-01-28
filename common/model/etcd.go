package model

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// Etcd 配置
type Etcd struct {
	Endpoints            []string `json:"endpoints" yaml:"endpoints"`
	UserName             string   `json:"userName" yaml:"userName"`
	Password             string   `json:"password" yaml:"password"`
	DialTimeout          int64    `json:"dialTimeout" yaml:"dialTimeout"`
	DialKeepAliveTime    int64    `json:"dialKeepAliveTime" yaml:"dialKeepAliveTime"`
	DialKeepAliveTimeout int64    `json:"dialKeepAliveTimeout" yaml:"dialKeepAliveTimeout"`
}

// ToConfig 转换为etcd client config
func (e *Etcd) ToConfig() (conf clientv3.Config) {
	return clientv3.Config{
		Endpoints:            e.Endpoints,
		Username:             e.UserName,
		Password:             e.Password,
		DialTimeout:          time.Duration(e.DialTimeout) * time.Second,
		DialKeepAliveTime:    time.Duration(e.DialKeepAliveTime) * time.Second,
		DialKeepAliveTimeout: time.Duration(e.DialKeepAliveTimeout) * time.Second,
	}
}
