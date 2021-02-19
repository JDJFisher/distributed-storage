package servers

import (
	"context"

	"github.com/JDJFisher/distributed-storage/protos"
)

// Neighbours - The current state of this nodes neighbours
type Neighbours struct {
	PredAddress string
	SuccAddress string
}

// ChainServer - ...
type ChainServer struct {
	protos.UnimplementedChainServer
	Neighbours *Neighbours
}

// NewChainServer - Create a new chain server object
func NewChainServer(neighbours *Neighbours) *ChainServer {
	return &ChainServer{Neighbours: neighbours}
}

// UpdateNeighbours - ...
func (s *ChainServer) UpdateNeighbours(ctx context.Context, req *protos.NeighbourInfo) (*protos.OkResponse, error) {
	s.Neighbours.PredAddress = req.PredAddress
	s.Neighbours.SuccAddress = req.SuccAddress
	return &protos.OkResponse{}, nil
}
