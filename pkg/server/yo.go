package server

import (
	"context"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

func NewYoService(name string) hellopb.YoServiceServer {
	yo := &messageService{
		name: name,
	}
	return yo
}

func (ms *messageService) Yo(ctx context.Context, in *hellopb.Send) (*hellopb.Response, error) {
	return &hellopb.Response{Message: ("Yo!")}, nil
}
