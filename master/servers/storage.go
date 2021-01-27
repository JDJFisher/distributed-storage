package servers

import (
	"context"

	"github.com/JDJFisher/distributed-storage/protos"
)

// Storage server

type StorageServer struct {
	protos.UnimplementedStorageServer
}

func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {

	// TODO: Request from the tail

	return &protos.ReadResponse{}, nil
}

func (s *StorageServer) Write(ctx context.Context, req *protos.WriteRequest) (*protos.WriteResponse, error) {

	// TODO: Write to the head

	return &protos.WriteResponse{}, nil
}
