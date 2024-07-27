package server

import (
	"context"
	"fmt"
	"grpc/echo"
	"io"
	"log"
)

func NewEchoServer() echo.EchoServer {
	return &echoServer{}
}

type echoServer struct {
	echo.UnimplementedEchoServer
}

func (s *echoServer) UnaryEcho(ctx context.Context, in *echo.EchoMsg) (*echo.EchoMsg, error) {
	fmt.Printf("server recv message:%+v\n", in)
	return in, nil
}
func (s *echoServer) ClientStreamEcho(stream echo.Echo_ClientStreamEchoServer) error {
	for {
		in, err := stream.Recv()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("server recv message:%+v\n", in)
	}
	err := stream.SendAndClose(&echo.EchoResponse{
		Ok: true,
	})
	return err
}
