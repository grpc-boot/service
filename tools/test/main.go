package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/grpc-boot/base"
)

func main() {
	for i := 0; i < 8; i++ {
		go func() {
			for {
				request("http://127.0.0.1:8080/")
			}
		}()
	}

	var wa chan struct{}
	<-wa
}

func request(url string) {
	resp, err := http.Get(url)
	if err != nil {
		base.Red("request %s err:%s", url, err.Error())
		return
	}

	body, er := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if er != nil {
		base.Red("read body err:%s", err.Error())
		return
	}

	base.Green("%s %s", time.Now().Format("2006-01-02 15:04:05"), body)
}
