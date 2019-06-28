package client

import (
	"context"
	"io"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

func (cli *Client) Concat(messages [3]string) (*hellopb.Response, error) {

	stream, err := cli.concatClient.Concat(context.Background())

	if err != nil {
		return nil, err
	}

	for _, msg := range messages {
		message := &hellopb.Send{
			Message: msg,
		}
		if err := stream.Send(message); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return stream.CloseAndRecv()
}
