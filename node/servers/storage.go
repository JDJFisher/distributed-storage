package servers

import (
	"context"

	"github.com/JDJFisher/distributed-storage/protos"
)

// StorageServer ...
type StorageServer struct {
	protos.UnimplementedStorageServer
}

func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {

	// TODO: Only serve if assigned a role in the chain

	return &protos.ReadResponse{}, nil
}

func (s *StorageServer) Write(ctx context.Context, req *protos.WriteRequest) (*protos.WriteResponse, error) {

	// TODO: Only serve if assigned a role in the chain

	return &protos.WriteResponse{}, nil
}
