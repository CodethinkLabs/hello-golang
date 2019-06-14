package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/CodethinkLabs/hello-golang/pkg/client"
	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

const address = "localhost:8080"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connection problem: ", err)
	}
	defer conn.Close()

	mc := hellopb.NewMessageServiceClient(conn)
	cli := client.NewClient("foo", mc)
	response, _ := cli.Message("Hey")

	fmt.Printf("%s", response.Message)

}
