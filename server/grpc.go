package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGRPCServer(port string) {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	log.Printf("gRPC server running on %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
