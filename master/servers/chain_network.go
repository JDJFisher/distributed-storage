package servers

import (
	"context"
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
		log.Printf("Node %v is attempting to join the network but appears to have not been cleaned up yet (wait for a health check and try again)\n", node.Address)
		return &protos.NeighbourInfo{
			Success:     false,
			PredAddress: "",
			SuccAddress: "",
		}, nil
	}

	log.Printf("%s is requesting to join the network", req.Address)
	s.Chain.Lock()
	defer s.Chain.Unlock()

	// Add the node to the chain TODO: Handle error
	node, _ = s.Chain.AddNode(req.Address)
	s.Chain.Print()

	response := &protos.NeighbourInfo{
		Success:     true,
		PredAddress: node.GetPredAddress(),
		SuccAddress: node.GetSuccAddress(),
	}

	return response, nil
}
