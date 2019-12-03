package main

import (
	"fmt"
	cache "github.com/patrickmn/go-cache"
	"time"
)

func main() {
	c := cache.New(5*time.Second, 10*time.Second)
	c.SetDefault("k", "v")
	for {
		fmt.Println(c.Get("k"))
		time.Sleep(1 * time.Second)
	}
}
