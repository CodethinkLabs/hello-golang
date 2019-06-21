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

func (cli *Client) MessageReverse(message string) (*hellopb.Response, error) {

	msg := &hellopb.Send{
		Message: Reverse(message),
	}

	return cli.messageClient.MessageReverse(context.Background(), msg)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
