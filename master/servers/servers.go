package servers

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
)

// Network wrapper

type NetworkServer struct {
	protos.UnimplementedNetworkServer
}

func (s *NetworkServer) JoinNetwork(ctx context.Context, req *protos.NetworkJoinRequest) (*protos.NetworkJoinResponse, error) {
	log.Println("Received a network join request")
	return &protos.NetworkJoinResponse{Type: protos.NetworkJoinResponse_NORMAL}, nil
}

// Storage wrapper

type StorageServer struct {
	protos.UnimplementedStorageServer
}

func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {
	return &protos.ReadResponse{Message: "foo"}, nil
}
