package servers

import (
	"context"
	"sync"

	"google.golang.org/grpc/peer"

	"github.com/JDJFisher/distributed-storage/protos"
)

// ChainServer ...
type ChainServer struct {
	protos.UnimplementedChainServer
	sync.RWMutex
	Chain *Chain
}

// Register ...
func (s *ChainServer) Register(ctx context.Context, req *protos.RegisterRequest) (*protos.RegisterResponse, error) {
	p, _ := peer.FromContext(ctx)

	s.Lock()
	defer s.Unlock()

	// Create a new node
	node := &Node{
		debug:       req.Name,
		address:     p.Addr.String(),
		successor:   nil,
		predecessor: nil}

	// No chain exists yet, lets create one!
	if s.Chain == nil {
		s.Chain = NewChain(node)
	} else {
		// Add to the chain, one already exists
		node.predecessor = s.Chain.Tail

		// TODO: Transfer tailness to the new tail
		s.Chain.Tail.successor = node
		s.Chain.Tail = node
	}

	preAddr := ""
	sucAddr := ""
	if node.predecessor != nil {
		sucAddr = node.predecessor.address
	}
	if node.successor != nil {
		sucAddr = node.successor.address
	}

	s.Chain.Print()

	return &protos.RegisterResponse{Predecessor: preAddr, Successor: sucAddr}, nil
}
