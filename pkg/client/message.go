package client

import (
	"context"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

type Client struct {
	name          string
	messageClient hellopb.MessageServiceClient
}

func NewClient(name string, messageClient hellopb.MessageServiceClient) *Client {
	return &Client{
		name:          name,
		messageClient: messageClient,
	}
}

func (cli *Client) Message(message string) (*hellopb.Response, error) {

	msg := &hellopb.Send{
		Message: message,
	}

	return cli.messageClient.Message(context.Background(), msg)
}
