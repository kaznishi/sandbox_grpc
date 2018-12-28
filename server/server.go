package main

import (
	"log"
	"net"

	pb "github.com/kaznishi/sandbox_grpc/proto/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello again " + in.Name}, nil
}

func (s *server) SayHelloHoge(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	st := status.New(codes.PermissionDenied, "permission denied")

	details := &pb.PermissionDeniedDetail{
		Type: pb.PermissionDeniedDetail_TYPE_HOGE,
		Code: pb.PermissionDeniedDetail_CODE_FUGA,
	}

	if stwd, err := st.WithDetails(details); err == nil {
		return nil, stwd.Err()
	}

	return nil, st.Err()
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
