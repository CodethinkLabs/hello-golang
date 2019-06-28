package client

import hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"

type Client struct {
	name          string
	messageClient hellopb.MessageServiceClient
	yoClient      hellopb.YoServiceClient
	concatClient  hellopb.ConcatServiceClient
}

func NewClient(name string, messageClient hellopb.MessageServiceClient, yoClient hellopb.YoServiceClient, concatClient hellopb.ConcatServiceClient) *Client {
	return &Client{
		name:          name,
		messageClient: messageClient,
		yoClient:      yoClient,
		concatClient:  concatClient,
	}
}
