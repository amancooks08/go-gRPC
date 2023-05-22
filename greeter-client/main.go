package main

import (
	"context"
	"fmt"
	proto "go-grpc/greeting-app/proto/greeting"
	"log"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting gRPC client application...")

	// use grpc.Dial to connect to the running gRPC server
	// use grpc.WithBlock() to block until the connection is established
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
		return
	}

	defer conn.Close()

	// use the NewGreetingServiceClient method of the generated in the .pb.go file
	// to create a new GreetingServiceClient object
	// this object can be used to call methods implemented in the grpc server

	client := proto.NewGreetingServiceClient(conn)

	// create a new GreetingRequest object
	request := &proto.GreetingRequest{
		Name: "John",
	}

	// call the Greet method of the GreetingServiceClient object
	response, err := client.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
		return
	}

	fmt.Println("Response from server: " + response.String())
}