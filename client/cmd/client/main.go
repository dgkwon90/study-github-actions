package main

import (
	"context"
	"fmt"
	"go-health/api/health"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func SendPing() {
	serverIP := os.Getenv("SERVER_IP")
	if serverIP == "" {
		serverIP = "9999" // default
	}

	i := 0
	conn, connErr := grpc.Dial(serverIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if connErr != nil {
		log.Fatalf("did not connect: %v", connErr)
	}
	c := health.NewHealthClient(conn)

	timeoutSecs := 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSecs)*time.Second)
	defer cancel()
	for {
		req := &health.PingReq{
			RequesterName: "go-health-client",
			Msg:           fmt.Sprintf("ping-%d", i),
		}
		log.Println(">>> SEND Request SendPing")
		res, err := c.SendPing(ctx, req)
		if err != nil {
			log.Panicf("could not get dokcing state: %v", err)
		}
		log.Printf("<<< RECV Response SendPing ResponserName: %s, Msg: %s", res.ResponserName, res.Msg)
		i++
		sleepSecs := 10
		time.Sleep(time.Duration(sleepSecs) * time.Second)
	}
}

func main() {
	buildInfoPrint()
	SendPing()
}
