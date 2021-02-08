package servers

import (
	"context"
	"fmt"
	"log"

	"github.com/JDJFisher/distributed-storage/master/chain"
	"github.com/JDJFisher/distributed-storage/protos"
)

// ChainServer ...
type ChainServer struct {
	protos.UnimplementedChainServer
	Chain *chain.Chain
}

func NewChainServer(chain *chain.Chain) *ChainServer {
	return &ChainServer{Chain: chain}
}

// Register ...
func (s *ChainServer) Register(ctx context.Context, req *protos.RegisterRequest) (*protos.NeighbourInfo, error) {

	// TODO: Handle if the node is already in the chain ...
	node := s.Chain.GetNode(req.Address)
	if node != nil {
		fmt.Printf("Node %v is already in the network (probably hasn't been cleaned up in the health check yet)", node.Address)
		return nil, nil
	}

	log.Printf("%s is requesting to join the network", req.Address)
	s.Chain.Lock()
	defer s.Chain.Unlock()

	// Add the node to the chain
	node = s.Chain.AddNode(req.Address)
	s.Chain.Print()

	response := &protos.NeighbourInfo{
		PredAddress: node.GetPredAddress(),
		SuccAddress: node.GetSuccAddress(),
	}

	return response, nil
}
