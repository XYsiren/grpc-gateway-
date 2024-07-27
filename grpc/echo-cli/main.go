package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc/echo"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := echo.NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	in := &echo.EchoMsg{
		Name: "siren",
		Addr: &echo.Addr{
			Province: "河北",
			City:     "石家庄",
		},
		Birthday: timestamppb.New(time.Now()),
		Data:     []byte("越努力，越幸运"),
		Gender:   echo.Gender_FEMAL,
		Hobby:    []string{"看小说", "看综艺", "追男团"},
	}
	res, err := client.UnaryEcho(ctx, in)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(res)

	stream, err := client.ClientStreamEcho(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	for i := 0; i < 5; i++ {
		err = stream.Send(in)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		result, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Println(result)
	}
}
