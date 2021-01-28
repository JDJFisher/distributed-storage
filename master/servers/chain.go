package servers

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
)

type CandidateNode struct {
	name string
}

// ChainServer ...
type ChainServer struct {
	protos.UnimplementedChainServer
	CandidateNodes []*CandidateNode
}

// Register ...
func (s *ChainServer) Register(ctx context.Context, req *protos.RegisterRequest) (*protos.RegisterResponse, error) {
	newCandidateNode := &CandidateNode{name: req.ServiceName}
	s.CandidateNodes = append(s.CandidateNodes, newCandidateNode)

	log.Printf("Added %v to candidates", req.ServiceName)

	return &protos.RegisterResponse{}, nil
}
