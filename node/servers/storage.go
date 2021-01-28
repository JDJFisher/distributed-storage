package servers

import (
	"context"
	// "log"

	"github.com/JDJFisher/distributed-storage/protos"
)

// StorageServer ...
type StorageServer struct {
	protos.UnimplementedStorageServer
}

func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {
	// log.Println("Foo")
	return &protos.ReadResponse{}, nil
}

func (s *StorageServer) Write(ctx context.Context, req *protos.WriteRequest) (*protos.WriteResponse, error) {
	// log.Println("Bar")
	return &protos.WriteResponse{}, nil
}
