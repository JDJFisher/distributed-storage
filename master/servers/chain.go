package servers

import (
	"context"
	"log"

	"google.golang.org/grpc/peer"

	"github.com/JDJFisher/distributed-storage/protos"
)

type CandidateNode struct {
	address string
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

	newCandidateNode := &CandidateNode{address: req.Address}
	s.CandidateNodes = append(s.CandidateNodes, newCandidateNode)

	log.Printf("Added %v (%v) to candidates", req.Address, addr)

	return &protos.RegisterResponse{}, nil
}
