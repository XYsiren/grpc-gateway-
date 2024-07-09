package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	gw "my-grpc-gateway/myservice1/service1"
	"net/http"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterMyService1HandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	return http.ListenAndServe("8081", mux)
}
