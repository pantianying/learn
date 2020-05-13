package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var cxt = context.Background()
	cxt1, f1 := context.WithCancel(cxt)
	go func(cxt context.Context) {
		cxt2, _ := context.WithCancel(cxt1)
		go func(cxt context.Context) {
			for {
				fmt.Println("go routine 2 running")
				select {
				case <-cxt.Done():
					fmt.Println("main tell routine 2 stop")
					return
				default:
					<-time.Tick(1 * time.Second)
				}
			}
		}(cxt2)

		for {
			fmt.Println("go routine 1 running")
			select {
			case <-cxt.Done():
				fmt.Println("main tell routine 1 stop")
				return
			default:
				<-time.Tick(1 * time.Second)
			}

		}
	}(cxt1)
	time.Sleep(time.Second * 10)
	f1()
	fmt.Println("f() done")
	time.Sleep(time.Second * 10)
}
