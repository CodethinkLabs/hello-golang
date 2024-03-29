package main

import (
	pb "proto"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const address = "localhost:8080"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("connection problem")
	}
	defer conn.Close()
	c := pb.NewConfessionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1000)
	defer cancel()


	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')

	for {
		fmt.Print("Your message: ")
		message, _ := reader.ReadString('\n')

		r, err := c.Confess(ctx, &pb.SubmitMessage{Name: name, Message: message})
		if err != nil {
			log.Fatal("connection refused")
		}

		if r.Message == "" {
			fmt.Println("No Messages")
		} else {
			fmt.Printf("Last message from %v: %v\n", r.Name, r.Message)
		}
	}
}