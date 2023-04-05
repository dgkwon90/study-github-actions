package main

import (
	"fmt"
	"go-health/server/internal/healthsvc"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

var (
	// GitCommit, BuildTime, Get infos at build time the golang.
	GitCommit string
	BuildTime string
)

// buildInfoPrint display git short commit id, build time.
// It is used to know the version of the launched application.
func buildInfoPrint() {
	// default log
	log.Printf("Build Information: %v at %v\n", GitCommit, BuildTime)
	log.Println("Started at: ", time.Now().Format(time.RFC3339))
}

func main() {
	buildInfoPrint()
	server := grpc.NewServer()
	// register service
	if regiErr := healthsvc.RegisterService(server); regiErr != nil {
		log.Fatalf("failed to register: %v", regiErr)
	}

	hostname := os.Getenv("SERVICE_HOST_NAME")
	if hostname == "" {
		hostname = "" // default
	}
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "9999" // default
	}

	serverIP := fmt.Sprintf("%s:%s", hostname, port)
	lis, err := net.Listen("tcp", serverIP)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// main service logic
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
