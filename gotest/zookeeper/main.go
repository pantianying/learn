package main

import (
	"fmt"
	"github.com/TuyaInc/pulsar-client-go/pkg/log"
	"github.com/samuel/go-zookeeper/zk"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Info(http.ListenAndServe("localhost:8809", nil))
	}()
	a := make(chan int, 0)
	var (
		conn *zk.Conn
		err  error
	)
	for {
		conn, _, err = zk.Connect([]string{"127.0.0.1:2181"}, 3*time.Second)
		fmt.Println(err)
		ss, stat, c, e := conn.ChildrenW("/")
		fmt.Println(ss, stat, c, e)
		ss, stat, c, e = conn.ChildrenW("/dubbo")
		fmt.Println(ss, stat, c, e)
		conn.Close()
		time.Sleep(1 * time.Second)
	}
	a <- 0
}
