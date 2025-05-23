package main

import (
	"log"
	"net"

	pb "github.com/usama1031/go-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8083"
)

type helloServer struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
