package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	rawClient, err := clientv3.New(clientv3.Config{
		Context:     ctx,
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		panic(err)
	}
	rawClient.Put(ctx, "pty", "ptyVal")
	rawClient.Put(ctx, "pty/1", "ptyVal1")
	rawClient.Delete(ctx, "pty")
	getResponse, err := rawClient.Get(ctx, "pty", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	for _, v := range getResponse.Kvs {
		fmt.Printf("%+v\n", v)
	}
	fmt.Printf("getResponse:%+v\n", getResponse)
	cancel()
	<-ctx.Done()
	fmt.Println("done")
}
