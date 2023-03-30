package healthsvc

import (
	"context"
	"fmt"
	"go-health/api/health"
	"log"

	"google.golang.org/grpc"
)

type Health struct {
	health.UnimplementedHealthServer
}

func RegisterService(source *grpc.Server) error {
	s := new(Health)
	health.RegisterHealthServer(source, s)
	return nil
}

func (s *Health) SendPing(ctx context.Context, in *health.PingReq) (*health.PingRes, error) {
	log.Printf("<<< RECV Request SendPing RequesterName: %v, msg: %v", in.RequesterName, in.Msg)
	log.Println(">>> SEND Response Pong")
	return &health.PingRes{
		ResponserName: "go-health-server",
		Ok:            true,
		Msg:           fmt.Sprintf("%s-pong", in.Msg),
	}, nil
}
