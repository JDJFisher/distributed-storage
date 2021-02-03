package servers

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
)

// StorageServer ...
type StorageServer struct {
	protos.UnimplementedStorageServer
	Neighbours *Neighbours
	Cache      *cache.Cache
}

// NewStorageServer - Create a new storage server object
func NewStorageServer(neighbours *Neighbours, cache *cache.Cache) *StorageServer {
	return &StorageServer{Neighbours: neighbours, Cache: cache}
}

// Read - ...
func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {
	// TODO: Only serve if assigned a role in the chain

	value, found := s.Cache.Get(req.Key)

	if found {
		return &protos.ReadResponse{Value: value.(string)}, nil
	}

	return &protos.ReadResponse{}, nil
}

// Write - ...
func (s *StorageServer) Write(ctx context.Context, req *protos.WriteRequest) (*protos.WriteResponse, error) {
	// TODO: Only serve if assigned a role in the chain
	// TODO: Process value

	s.Cache.Set(req.Key, req.Value, cache.NoExpiration)

	// Propagate the write request down the chain
	if s.Neighbours.SuccAddress != "" {

		// Open a connection to the successor node
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(s.Neighbours.SuccAddress, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Error connecting to successor node - %v", err.Error())
		}
		defer conn.Close()

		// Create storage client
		storageClient := protos.NewStorageClient(conn)

		// Forward write request to successor
		response, err := storageClient.Write(context.Background(), req)
		if err != nil {
			log.Fatalf("Error forwarding write request to successor node - %v", err.Error())
		}

		return response, nil
	}

	return &protos.WriteResponse{Value: req.Value}, nil
}
