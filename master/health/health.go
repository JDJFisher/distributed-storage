package health

import (
	"context"
	"log"
	"time"

	"github.com/JDJFisher/distributed-storage/protos"
)

type HealthServer struct {
	protos.UnimplementedHealthServer
	monitoredNodes map[string]time.Time //Nodes which are sending health checks, their service name and the latest health check received from them
}

func NewHealthServer() *HealthServer {
	return &HealthServer{monitoredNodes: make(map[string]time.Time)}
}

func (s *HealthServer) Alive(ctx context.Context, req *protos.HealthCheckRequest) (*protos.HealthCheckResponse, error) {
	log.Printf("Received health check from: %s", req.Service)

	currentTime := time.Now()
	serviceName := req.Service
	s.monitoredNodes[serviceName] = currentTime

	return &protos.HealthCheckResponse{Status: protos.HealthCheckResponse_WAITING}, nil
}

func (s *HealthServer) CheckNodes(interval uint8) {
	for {
		<-time.After(time.Duration(interval) * time.Second)
		now := time.Now()
		for k, v := range s.monitoredNodes {
			if now.Sub(v).Seconds() > float64(interval) {
				log.Printf("Node %s hasnt sent a health check within %d seconds!", k, interval)
			}
		}
	}

}
