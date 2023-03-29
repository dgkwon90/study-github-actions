package main

import (
	"go-health/server/healthsvc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	// register service
	if regiErr := healthsvc.RegisterService(server); regiErr != nil {
		log.Fatalf("failed to register: %v", regiErr)
	}

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// main service logic
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
