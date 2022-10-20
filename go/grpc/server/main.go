package main

import (
	"context"
	"fmt"
	pb "github.com/pantianying/learn/go/grpc/proto"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":50051"
)

type server struct {
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
func (s *server) GetUser(ctx context.Context, req *pb.HelloReq) (resp *pb.HelloReply1, err error) {
	fmt.Println("get user", *req)
	return &pb.HelloReply1{Message: "hello pty"}, nil
}
