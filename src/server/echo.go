package server

import (
	"context"
	"os"

	proto "test-grpc/k8s-grpc-gateway-example/proto"
)

type echoServer struct{}

func newEchoServer() proto.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, req *proto.RequestMessage) (*proto.ResponseMessage, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	if req.Message == "" {
		return &proto.ResponseMessage{
			Host:    hostname,
			Message: ":(",
		}, nil
	}

	return &proto.ResponseMessage{
		Host:    hostname,
		Message: req.Message,
	}, nil
}
