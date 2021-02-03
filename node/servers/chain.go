package servers

import (
	"context"

	"github.com/JDJFisher/distributed-storage/protos"
)

// ChainServer ...
type ChainServer struct {
	protos.UnimplementedChainServer
}

// UpdateNeighbours ...
func (s *ChainServer) UpdateNeighbours(ctx context.Context, req *protos.NeighbourInfo) (*protos.OkReponse, error) {
	// TODO: Update neighbour addresses
	return &protos.OkReponse{}, nil
}
