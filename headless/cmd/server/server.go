package main

import (
	"context"
	"fmt"
	"github.com/tangxusc/istio-test/headless/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

type server struct {
	hostname string
}

func (s *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Printf("SayHello form hostname:%v\n", request.Hostname)
	return &proto.HelloReply{
		Hostname: s.hostname,
	}, nil
}

func main() {
	port := os.Getenv("PORT")
	hostname, _ := os.Hostname()
	fmt.Printf("Starting server on hostname:%v,port:%v\n", hostname, port)
	sock, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	proto.RegisterHelloServiceServer(s, &server{hostname: hostname})
	if err := s.Serve(sock); err != nil {
		panic(err)
	}
}
