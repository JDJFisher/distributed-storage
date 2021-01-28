package servers

import (
	"context"

	"github.com/JDJFisher/distributed-storage/protos"
)

// ChainServer ...
type ChainServer struct {
	protos.UnimplementedChainServer
}

// Assign ...
func (s *ChainServer) Assign(ctx context.Context, req *protos.AssignRequest) (*protos.AssignResponse, error) {
	return &protos.AssignResponse{}, nil
}
