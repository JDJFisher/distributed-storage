package servers

import (
	"context"
	"log"

	"google.golang.org/grpc/peer"

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
	p, _ := peer.FromContext(ctx)
	addr := p.Addr.String()

	newCandidateNode := &CandidateNode{name: addr}
	s.CandidateNodes = append(s.CandidateNodes, newCandidateNode)

	log.Printf("Added %v to candidates", addr)

	return &protos.RegisterResponse{}, nil
}
