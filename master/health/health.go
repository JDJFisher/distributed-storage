package health

import (
	"context"
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
	log.Printf("Received health check from: %s", req.Address)
	//TODO - check if we actually care about this node, if we dont reply the node should kill itself or try to rejoin the chain
	s.Lock()
	defer s.Unlock()

	currentTime := time.Now()
	s.monitoredNodes[req.Address] = currentTime

	return &protos.HealthCheckResponse{Status: protos.HealthCheckResponse_WAITING}, nil
}

//CheckNodes - Checks if the nodes in the chain have recently made a health check
func (s *HealthServer) CheckNodes(interval uint8) {
	for {
		<-time.After(time.Duration(interval) * time.Second)
		now := time.Now()
		for address, latestTime := range s.monitoredNodes {
			if now.Sub(latestTime).Seconds() > float64(interval) {
				log.Printf("Node %s hasnt sent a health check within %d seconds!", address, interval)

				s.chain.RemoveNode(address)

				// TODO: Tell the pred that its new succ is this nodes succ

			}
		}
	}

}
