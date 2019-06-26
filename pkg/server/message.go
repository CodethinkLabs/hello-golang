package server

import (
	"context"
	"fmt"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

func NewMessageService(name string) hellopb.MessageServiceServer {
	ms := &messageService{
		name: name,
	}
	return ms
}

type messageService struct {
	name string
}

func (ms *messageService) Message(ctx context.Context, in *hellopb.Send) (*hellopb.Response, error) {
	fmt.Println(ms.name)
	return &hellopb.Response{Message: in.Message}, nil
}

func (ms *messageService) MessageReverse(ctx context.Context, in *hellopb.Send) (*hellopb.Response, error) {
	return &hellopb.Response{Message: reverse(in.Message)}, nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
