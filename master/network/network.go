package network

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
)

// Implement the interface of the grpc server

type NetworkServer struct {
	protos.UnimplementedNetworkServer
}

func (s *NetworkServer) JoinNetwork(ctx context.Context, req *protos.NetworkJoinRequest) (*protos.NetworkJoinResponse, error) {
	log.Printf("Received a network join request from: %v", req.ServiceName)
	return &protos.NetworkJoinResponse{Type: protos.NetworkJoinResponse_NORMAL}, nil
}
