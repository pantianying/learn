package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	var a = make(chan int, 0)
	go func() {
		log.Println(http.ListenAndServe("localhost:10000", nil))
	}()
	<-a

}
