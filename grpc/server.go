package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"

	pb "proto"

	"github.com/golang/protobuf/ptypes"
)

type message struct {
	name      string
	message   string
	timestamp time.Time
}

type server struct {
	messages []message
}

func (s *server) Confess(ctx context.Context, in *pb.SubmitMessage) (*pb.ReceiveReply, error) {
	timeStart := time.Now()

	lastMessage := s.messages[len(s.messages)-1]

	newMessage := message{
		name:      strings.TrimSpace(in.Name),
		message:   strings.TrimSpace(in.Message),
		timestamp: timeStart,
	}

	timestampProto, _ := ptypes.TimestampProto(lastMessage.timestamp)

	reply := pb.ReceiveReply{
		Message:   lastMessage.message,
		Name:      lastMessage.name,
		Timestamp: timestampProto,
	}

	s.messages = append(s.messages, newMessage)
	log.Printf("Received message from %v: %v", newMessage.name, newMessage.message)
	return &reply, nil
}

func main() {
	srv, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to attach: %v", err)
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterConfessionServer(grpcSrv, &server{messages: make([]message, 10)})
	err = grpcSrv.Serve(srv)
	if err != nil {
		log.Fatal("FAILED")
	}
	log.Print("Server started.")
}
