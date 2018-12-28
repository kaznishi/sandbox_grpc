package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/codes"

	pb "github.com/kaznishi/sandbox_grpc/proto/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	r, err = c.SayHelloHoge(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		printDetail(err)
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}

func printDetail(err error) {
	s, _ := status.FromError(err)
	if s.Code() == codes.PermissionDenied {
		dt := s.Details()
		pd, _ := dt[0].(*pb.PermissionDeniedDetail)
		fmt.Println(pd.Type)
		fmt.Println(pd.Code)
	} else {
		fmt.Println("fugafuga")
	}
}
