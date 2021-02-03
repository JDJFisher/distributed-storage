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
func (s *ChainServer) Register(ctx context.Context, req *protos.RegisterRequest) (*protos.RegisterResponse, error) {
	log.Printf("%s is requesting to join the network", req.Address)

	s.Chain.Lock()
	defer s.Chain.Unlock()

	chainLen := s.Chain.Len()

	var newNode *chain.Node
	//If the chain is empty we need to manually setup the node pointers
	if chainLen == 0 {
		newNode = &chain.Node{Address: req.Address, Successor: nil, Predecessor: nil}
		s.Chain.Head = newNode
		s.Chain.Tail = newNode
	} else {
		//Add to the tail!
		newNode = &chain.Node{Address: req.Address, Successor: nil, Predecessor: s.Chain.Tail}
		s.Chain.Tail.Successor = newNode
		s.Chain.Tail = newNode
	}
	s.Chain.Print()

	return &protos.RegisterResponse{}, nil
}
