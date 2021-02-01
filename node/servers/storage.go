package servers

import (
	"context"

	"github.com/JDJFisher/distributed-storage/protos"
	"github.com/patrickmn/go-cache"
)

// StorageServer ...
type StorageServer struct {
	protos.UnimplementedStorageServer
	Cache *cache.Cache
}

func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {

	// TODO: Only serve if assigned a role in the chain

	value, _ := s.Cache.Get(req.Key)

	return &protos.ReadResponse{Value: value.(string)}, nil
}

func (s *StorageServer) Write(ctx context.Context, req *protos.WriteRequest) (*protos.WriteResponse, error) {

	// TODO: Only serve if assigned a role in the chain
	// TODO: Process value

	s.Cache.Set(req.Key, req.Value, cache.NoExpiration)

	return &protos.WriteResponse{Value: req.Value}, nil
}
