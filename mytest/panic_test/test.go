package main

import (
	"time"
)

func main(){
	i:=0
	for i<100{
		go func(){
			i++
			if i==10{
				panic("xxxxx")
			}
			time.Sleep(time.Second*100)
		}()
	}
	time.Sleep(time.Second*1000)
}