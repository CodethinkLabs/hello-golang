package client

import (
	"context"

	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

func (cli *Client) Yo(message string) (*hellopb.Response, error) {
	msg := &hellopb.Send{
		Message: message,
	}

	return cli.yoClient.Yo(context.Background(), msg)
}
