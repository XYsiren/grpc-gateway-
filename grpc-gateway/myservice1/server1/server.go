package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "my-grpc-gateway/myservice1/service1"
	"net"
)

type server struct {
	pb.UnimplementedMyService1Server
}

func (s *server) Echo(ctx context.Context, in *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	fmt.Printf("server recv : #{in}\n")
	return in, nil
}

func (s *server) EchoBody(ctx context.Context, in *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	fmt.Printf("server recv : #{in}\n")
	return in, nil
}

func (s *server) EchoDelete(ctx context.Context, in *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	fmt.Printf("server recv : #{in}\n")
	return in, nil
}

var (
	port = flag.Int("port", 50051, "server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyService1Server(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
