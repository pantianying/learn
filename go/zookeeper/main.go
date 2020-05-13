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
	a := make(chan int, 0)
	go func() {
		log.Info(http.ListenAndServe("localhost:8809", nil))
	}()

	conn, event, err := zk.Connect([]string{"127.0.0.1:2181"}, 3*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		e := <-event
		fmt.Println("event get ", e)
	}()
	s1, stat1, event1, e1 := conn.ChildrenW("/")
	fmt.Println("conn.ChildrenW(/)", s1, stat1, event1, e1)
	go func() {
		e := <-event1
		fmt.Println("event1 get ", e)
	}()

	s2, stat2, event2, e2 := conn.ChildrenW("/dubbo")
	fmt.Println("conn.ChildrenW(/dubbo)", s2, stat2, event2, e2)
	go func() {
		e := <-event2
		fmt.Println("event2 get ", e)
	}()
	s3, stat3, event3, e3 := conn.ChildrenW("/dubbo/com.tuya.provider.service.IHelloService")
	fmt.Println("conn.ChildrenW(/dubbocom.tuya.provider.service.IHelloService)", s3, stat3, event3, e3)
	go func() {
		e := <-event3
		fmt.Println("event3 get ", e)
	}()
	a <- 0
}
