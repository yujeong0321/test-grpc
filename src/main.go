package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"test-grpc/k8s-grpc-gateway-example/src/gateway"
	"test-grpc/k8s-grpc-gateway-example/src/server"
)

var (
	serverAddr  = flag.String("serverAddr", ":50051", "endpoint of the gRPC server")
	gatewayAddr = flag.String("gatewayAddr", ":50050", "endpoint of the gRPC gateway")
	network     = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		cancel()
		fmt.Println("Exiting server on ", sig)
		os.Exit(0)
	}()

	go func() {
		fmt.Println("Starting HTTP gateway on", *gatewayAddr)
		if err := gateway.Run(ctx, *gatewayAddr, "0.0.0.0"+*serverAddr); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Starting gRPC server on", *serverAddr, *serverAddr)
	if err := server.Run(ctx, *network, *serverAddr); err != nil {
		log.Fatal(err)
	}
}
