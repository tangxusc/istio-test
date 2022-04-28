package main

import (
	"context"
	"fmt"
	_ "github.com/Jille/grpc-multi-resolver"
	"github.com/tangxusc/istio-test/headless/pkg/proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/health"
	"log"
	"os"
	"time"
)

func main() {
	addr := os.Getenv("ADDR")
	hostname, _ := os.Hostname()
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewHelloServiceClient(conn)
	for {
		fmt.Printf("[%s] Sending request\n", hostname)
		hello, err := client.SayHello(context.TODO(), &proto.HelloRequest{Hostname: hostname})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Printf("response on host: %s\n", hello.Hostname)
		time.Sleep(time.Second)
	}
}
