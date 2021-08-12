package server

import (
	"context"
	"fmt"
	"net"

	proto "test-grpc/k8s-grpc-gateway-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			fmt.Printf("Failed to close %s %s: %v\n", network, address, err)
		}
	}()

	s := grpc.NewServer()
	proto.RegisterEchoServiceServer(s, newEchoServer())
	reflection.Register(s)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}
