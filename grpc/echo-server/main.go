package main

import (
	"google.golang.org/grpc"
	"grpc/echo"
	"grpc/echo-server/server"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, server.NewEchoServer())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
