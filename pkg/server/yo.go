package server

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

func NewYoService(name string) hellopb.YoServiceServer {
	yo := &messageService{
		name: name,
	}
	return yo
}

func (ms *messageService) Yo(ctx context.Context, in *hellopb.Send) (*hellopb.Response, error) {
	fmt.Println("Received a call: ", in)
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
	fmt.Println("Work finished for: ", in)
	return &hellopb.Response{Message: in.Message}, nil
}
