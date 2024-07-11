package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	pb "my-grpc-gateway/myservice2/service2"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedMyService2Server
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

func (s *server) EchoUpload(stream pb.MyService2_EchoUploadServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Fatalln("Failed to get metadata from context")
	}
	filename := md["filename"][0]
	fmt.Printf("server recv : %v\n", filename)
	filePath := "myservice2/server/upload/" + filename
	dst, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer dst.Close()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		dst.Write(req.Content[:req.Size])
	}
	stream.SendAndClose(&pb.UploadResponse{
		Path: filePath,
	})
	return nil
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
	pb.RegisterMyService2Server(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
