package servers

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
)

// Network server

type NetworkServer struct {
	protos.UnimplementedNetworkServer
}

func (s *NetworkServer) JoinNetwork(ctx context.Context, req *protos.NetworkJoinRequest) (*protos.NetworkJoinResponse, error) {
	log.Println("Received a network join request")
	return &protos.NetworkJoinResponse{Type: protos.NetworkJoinResponse_NORMAL}, nil
}
