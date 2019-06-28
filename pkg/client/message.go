package client

import (
	"context"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

func (cli *Client) Message(message string) (*hellopb.Response, error) {

	msg := &hellopb.Send{
		Message: message,
	}

	return cli.messageClient.Message(context.Background(), msg)
}

func (cli *Client) MessageReverse(message string) (*hellopb.Response, error) {

	msg := &hellopb.Send{
		Message: message,
	}

	return cli.messageClient.MessageReverse(context.Background(), msg)
}
