package conf

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/client/v3"
)

var (
	cache sync.Map
)

func Init() {
	go watch()
}

func watch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.16.49.131:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, "config", clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Fatal("get failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		cache.Store(string(ev.Key), ev.Value)
	}

	cache.Range(func(key, value interface{}) bool {
		fmt.Printf("%s:%q\n", key, value)
		return true
	})
	fmt.Println("")

	rch := cli.Watch(context.Background(), "config", clientv3.WithPrefix(), clientv3.WithFromKey())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				cache.Store(string(ev.Kv.Key), ev.Kv.Value)
			case mvccpb.DELETE:
				cache.Delete(string(ev.Kv.Key))
			}

			cache.Range(func(key, value interface{}) bool {
				fmt.Printf("%s:%q\n", key, value)
				return true
			})
			fmt.Println("")
		}
	}
}
