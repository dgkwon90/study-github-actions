package main

import (
	"context"
	"fmt"
	"go-health/api/health"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendPing() {
	i := 0
	for {
		conn, connErr := grpc.Dial(":9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		log.Println("SEND Request SendPing")
		res, err := c.SendPing(ctx, req)
		if err != nil {
			log.Fatalf("could not get dokcing state: %v", err)
		}
		log.Printf("RECV Response SendPing ResponserName: %s, Msg: %s", res.ResponserName, res.Msg)
		i++
		time.Sleep(10 * time.Second)
	}
}

func main() {
	SendPing()
}
