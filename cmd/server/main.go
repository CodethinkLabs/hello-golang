package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

type messageService struct {
	name string
}

func NewMessageService(name string) hellopb.MessageServiceServer {
	ms := &messageService{
		name: name,
	}
	return ms
}

func (ms *messageService) Message(ctx context.Context, in *hellopb.Send) (*hellopb.Response, error) {
	return &hellopb.Response{Message: "bar"}, nil
}

func main() {
	hostname, _ := os.Hostname()
	msgService := NewMessageService(hostname)

	s := grpc.NewServer()
	hellopb.RegisterMessageServiceServer(s, msgService)

	sock, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to create listening socket: ", err)
	}
	if err := s.Serve(sock); err != nil {
		log.Fatal("Failed to serve RPC server: ", err)
	}
}
