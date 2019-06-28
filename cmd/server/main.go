package main

import (
	"log"
	"net"
	"os"

	"github.com/CodethinkLabs/hello-golang/pkg/server"
	"google.golang.org/grpc"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

func main() {
	hostname, _ := os.Hostname()
	msgService := server.NewMessageService(hostname)
	yoService := server.NewYoService(hostname)
	concatService := server.NewConcatService(hostname)

	s := grpc.NewServer()
	hellopb.RegisterMessageServiceServer(s, msgService)
	hellopb.RegisterYoServiceServer(s, yoService)
	hellopb.RegisterConcatServiceServer(s, concatService)

	sock, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to create listening socket: ", err)
	}
	if err := s.Serve(sock); err != nil {
		log.Fatal("Failed to serve RPC server: ", err)
	}
}
