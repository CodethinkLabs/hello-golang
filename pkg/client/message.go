package client

import (
	"context"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

type Client struct {
	name          string
	messageClient hellopb.MessageServiceClient
	yoClient      hellopb.YoServiceClient
}

func NewClient(name string, messageClient hellopb.MessageServiceClient, yoClient hellopb.YoServiceClient) *Client {
	return &Client{
		name:          name,
		messageClient: messageClient,
		yoClient:      yoClient,
	}
}

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

func (cli *Client) Yo(message string) (*hellopb.Response, error) {
	msg := &hellopb.Send{
		Message: message,
	}

	return cli.yoClient.Yo(context.Background(), msg)
}
