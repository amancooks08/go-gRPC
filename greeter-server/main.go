package main

import (
	"context"
	"fmt"
	proto "go-grpc/greeting-app/proto/greeting"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myGRPCServer struct {
	proto.UnimplementedGreetingServiceServer
}

// actual implementation of the Greet method
func (s *myGRPCServer) Greet(ctx context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	fmt.Println("Request received to greet: " + req.Name)
	return &proto.GreetingResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	fmt.Println("Starting server on port :8080")

	//start a listener on the port you want to start the gRPC server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// pass the address of the struct which implements the gRPC GreetingService interface
	// to the RegisterGreetingServiceServer method of the gRPC server object

	proto.RegisterGreetingServiceServer(grpcServer, &myGRPCServer{})

	// start the gRPC server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}

	// close the listener and the gRPC server
	defer lis.Close()
	defer grpcServer.Stop()
}
