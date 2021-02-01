package health

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/JDJFisher/distributed-storage/master/chain"
	"github.com/JDJFisher/distributed-storage/protos"
)

type HealthServer struct {
	protos.UnimplementedHealthServer
	monitoredNodes map[string]time.Time //Nodes which are sending health checks, their service name and the latest health check received from them
	sync.RWMutex
	chain *chain.Chain
}

func NewHealthServer(chain *chain.Chain) *HealthServer {
	return &HealthServer{monitoredNodes: make(map[string]time.Time), chain: chain}
}

//Alive (Node -> Master) - Health check ping coming from the node
func (s *HealthServer) Alive(ctx context.Context, req *protos.HealthCheckRequest) (*protos.HealthCheckResponse, error) {
	//log.Printf("Received health check from: %s", req.Name)
	//TODO - check if we actually care about this node, if we dont reply the node should kill itself or try to rejoin the chain
	s.Lock()
	defer s.Unlock()

	currentTime := time.Now()
	serviceName := req.Name
	s.monitoredNodes[serviceName] = currentTime

	return &protos.HealthCheckResponse{Status: protos.HealthCheckResponse_WAITING}, nil
}

//CheckNodes - Checks if the nodes in the chain have recently made a health check
func (s *HealthServer) CheckNodes(interval uint8) {
	for {
		<-time.After(time.Duration(interval) * time.Second)
		now := time.Now()
		for k, v := range s.monitoredNodes {
			if now.Sub(v).Seconds() > float64(interval) {
				log.Printf("Node %s hasnt sent a health check within %d seconds!", k, interval)

				//TODO - Remove the node from the chain and
				
			}
		}
	}

}
