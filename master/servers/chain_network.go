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
	var node *chain.Node
	// // TODO: Check if the node is already in the chain
	// node := s.Chain.GetNode(req.Address)
	// if node != nil {
	// 	s.Chain.RemoveNode(node) // TODO: We should also remove the node from the health map
	// }

	log.Printf("%s is requesting to join the network", req.Address)
	s.Chain.Lock()
	defer s.Chain.Unlock()

	chainLen := s.Chain.Len()

	//If the chain is empty we need to manually setup the node pointers
	if chainLen == 0 {
		node = chain.NewNode(req.Address, nil, nil)
		s.Chain.Head = node
		s.Chain.Tail = node
	} else {
		//Add to the tail!
		node = chain.NewNode(req.Address, nil, s.Chain.Tail)
		s.Chain.Tail.SetSucc(node)
		s.Chain.Tail = node
	}
	s.Chain.Print()

	response := &protos.NeighbourInfo{
		PredAddress: node.GetPredAddress(),
		SuccAddress: node.GetSuccAddress(),
	}

	return response, nil
}
