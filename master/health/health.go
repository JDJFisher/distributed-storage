package health

import (
	"context"
	"log"
)

type HealthCheckServer struct {
	UnimplementedHealthServer
}

func (s *HealthCheckServer) JoinNetwork(ctx context.Context, req *JoinNetworkRequest) (*JoinNetworkResponse, error) {
	log.Println("[MASTER] Received a join network request")

	return &JoinNetworkResponse{
		Status: JoinNetworkResponse_TAIL,
	}, nil
}
