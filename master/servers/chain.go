package servers

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc/peer"

	"github.com/JDJFisher/distributed-storage/protos"
)

// Enum representing the possible types a node can be
type NodeType int

const (
	HEAD      NodeType = 0 << iota
	TAIL      NodeType = 1
	STANDARD  NodeType = 2
	HEAD_TAIL NodeType = 3
	CANDIDATE NodeType = 4
)

// ChainServer ...
type ChainServer struct {
	protos.UnimplementedChainServer
	sync.RWMutex
	CandidateNodes []*Node
	Chain          *Chain
}

// Register ...
func (s *ChainServer) Register(ctx context.Context, req *protos.RegisterRequest) (*protos.RegisterResponse, error) {
	p, _ := peer.FromContext(ctx)
	addr := p.Addr.String()
	newCandidateNode := &Node{address: req.Address, successor: nil, predecessor: nil, nodeType: CANDIDATE}

	//No chain exists yet, lets create one!
	if s.Chain == nil {
		s.Chain = NewChain(newCandidateNode)
		log.Printf("A new chain has been created!")
	}

	s.Lock()
	s.CandidateNodes = append(s.CandidateNodes, newCandidateNode)
	defer s.Unlock()

	log.Printf("Added %v (%v) to candidates (candidate size=%d)", req.Address, addr, len(s.CandidateNodes))

	return &protos.RegisterResponse{}, nil
}
