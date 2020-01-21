package main

import (
	"fmt"
	"time"
)

var ch1 chan *reqStu

type reqStu struct {
	done chan struct{}
	data string
}

func main() {
	ch1 = make(chan *reqStu, 100)

	for i := 0; i < 100; i++ {
		go work()
	}

	for {
		time.Sleep(1 * time.Second)
		req := &reqStu{
			done: make(chan struct{}),
			data: "data example",
		}
		ch1 <- req

		select {
		case <-req.done:
			fmt.Println("做完了")
		case <-time.After(3 * time.Second):
			fmt.Println("超时了")
		}
	}
}
func work() {
	for {
		req := <-ch1
		time.Sleep(1 * time.Second)
		close(req.done) //代表做完了
	}
}
