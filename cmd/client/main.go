package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"google.golang.org/grpc"

	"github.com/CodethinkLabs/hello-golang/pkg/client"
	hellopb "github.com/CodethinkLabs/hello-golang/pkg/proto"
)

const address = "localhost:8080"

func yoCall(cli *client.Client, call string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, _ := cli.Yo(call)
	fmt.Println(response)
}

func main() {

	var wg sync.WaitGroup

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connection problem: ", err)
	}
	defer conn.Close()

	mc := hellopb.NewMessageServiceClient(conn)
	yo := hellopb.NewYoServiceClient(conn)
	cli := client.NewClient("foo", mc, yo)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go yoCall(cli, strconv.Itoa(i), &wg)

	}

	wg.Wait()
	fmt.Println("Done.")

}
