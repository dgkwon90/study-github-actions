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
	log.Printf("Build Information : %v at %v\n", GitCommit, BuildTime)
	log.Println("Started at :", time.Now().Format(time.RFC3339))
}

func SendPing() {
	serverIP := os.Getenv("SERVER_IP")
	if len(serverIP) <= 0 {
		serverIP = "9999" // default
	}

	i := 0
	for {
		conn, connErr := grpc.Dial(serverIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if connErr != nil {
			log.Fatalf("did not connect: %v", connErr)
		}
		c := health.NewHealthClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		req := &health.PingReq{
			RequesterName: "go-health-client",
			Msg:           fmt.Sprintf("ping-%d", i),
		}
		log.Println(">>> SEND Request SendPing")
		res, err := c.SendPing(ctx, req)
		if err != nil {
			log.Fatalf("could not get dokcing state: %v", err)
		}
		log.Printf("<<< RECV Response SendPing ResponserName: %s, Msg: %s", res.ResponserName, res.Msg)
		i++
		time.Sleep(10 * time.Second)
	}
}

func main() {
	buildInfoPrint()
	SendPing()
}
