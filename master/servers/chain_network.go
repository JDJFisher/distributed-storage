package servers

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/JDJFisher/distributed-storage/master/chain"
	"github.com/JDJFisher/distributed-storage/protos"
)

// ChainServer ...
type ChainServer struct {
	protos.UnimplementedChainServer
	Chain       *chain.Chain
	IsExtending bool
}

func NewChainServer(chain *chain.Chain) *ChainServer {
	return &ChainServer{Chain: chain}
}

// GetTail ...
func (s *ChainServer) GetTail(ctx context.Context, req *protos.TailRequest) (*protos.TailResponse, error) {
	if s.IsExtending {
		return nil, errors.New("Chain is extending")
	}

	s.IsExtending = true

	return &protos.TailResponse{
		Address: s.Chain.GetTailAddress(),
	}, nil
}

// Join ...
func (s *ChainServer) Join(ctx context.Context, req *protos.JoinRequest) (*protos.OkResponse, error) {
	log.Printf("%s is requesting to join the network", req.Address)

	s.Chain.Lock()
	defer s.Chain.Unlock()

	// Verify tail address is correct
	if req.TailAddress != s.Chain.GetTailAddress() {
		return nil, errors.New("Whoops")
	}

	node := s.Chain.GetNode(req.Address)
	if node != nil {
		err := errors.New(fmt.Sprintf("Node %v is attempting to join the network but appears to have not been cleaned up yet (wait for a health check and try again)\n", node.Address))
		return nil, err
	}

	_, err := s.Chain.AddNode(req.Address)
	if err != nil {
		//
	}
	s.Chain.Print()

	s.IsExtending = false
	return &protos.OkResponse{}, nil
}
