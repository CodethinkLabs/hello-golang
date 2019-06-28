package server

import (
	"io"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

func NewConcatService(name string) hellopb.ConcatServiceServer {
	conc := &messageService{
		name: name,
	}
	return conc
}

func (ms *messageService) Concat(stream hellopb.ConcatService_ConcatServer) error {

	concat := ""
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&hellopb.Response{
				Message: concat,
			})
		}
		if err != nil {
			return err
		}
		concat += message.Message
	}
}
