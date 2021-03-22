package servers

import (
	"context"
	"log"
	"os"

	"github.com/JDJFisher/distributed-storage/protos"
	"github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

// StorageServer ...
type StorageServer struct {
	protos.UnimplementedStorageServer
	Neighbours *Neighbours
	Cache      *cache.Cache
	Staging    map[uuid.UUID]*protos.WriteRequest
}

// NewStorageServer - Create a new storage server object
func NewStorageServer(neighbours *Neighbours, cache *cache.Cache) *StorageServer {
	return &StorageServer{
		Neighbours: neighbours,
		Cache:      cache,
		Staging:    make(map[uuid.UUID]*protos.WriteRequest),
	}
}

// Read - ...
func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {
	value, found := s.Cache.Get(req.Key)

	if found {
		return &protos.ReadResponse{Value: value.(string)}, nil
	}

	return &protos.ReadResponse{}, nil
}

// Write - ...
func (s *StorageServer) Write(ctx context.Context, req *protos.WriteRequest) (*protos.OkResponse, error) {
	//
	uid, _ := uuid.FromString(req.Uuid)

	// Stage the write
	s.Staging[uid] = req

	var grpcHost string

	// Determine host based on if tail
	if s.Neighbours.SuccAddress == "" {
		grpcHost = os.Getenv("host")
	} else {
		grpcHost = s.Neighbours.SuccAddress
	}

	// Open a connection
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting - %v", err.Error())
	}
	defer conn.Close()

	// Create storage client
	storageClient := protos.NewStorageClient(conn)

	if s.Neighbours.SuccAddress == "" {

		// Persist the write to the replica
		s.Cache.Set(req.Key, req.Value, cache.NoExpiration)

		// Response to the master
		req := &protos.ProcessedRequest{Uuid: uid.String()}
		_, err = storageClient.Processed(context.Background(), req)
		if err != nil {
			log.Fatalf("Error ... to the master - %v", err.Error())
		}

	} else {
		// Propagate the write request down the chain
		_, err = storageClient.Write(context.Background(), req)
		if err != nil {
			log.Fatalf("Error forwarding write request to successor node - %v", err.Error())
		}
	}

	return &protos.OkResponse{}, nil
}

//
func (s *StorageServer) Persist(ctx context.Context, req *protos.ProcessedRequest) (*protos.OkResponse, error) {
	//
	uid, _ := uuid.FromString(req.Uuid)

	//
	r := s.Staging[uid]

	// Persist the write to the replica
	s.Cache.Set(r.Key, r.Value, cache.NoExpiration)

	//
	delete(s.Staging, uid)

	return &protos.OkResponse{}, nil
}

func (s *StorageServer) GetTailData(req *protos.RequestData, stream protos.Storage_GetTailDataServer) error {
	for k, v := range s.Cache.Items() {
		log.Printf("Sending %s: %s to the new node", k, v.Object.(string))
		data := &protos.RequestDataResponse{Key: k, Value: v.Object.(string)}

		if err := stream.Send(data); err != nil {
			log.Fatalf("Error sending key value to the new node - %v", err.Error())
		}
	}
	log.Println("Sent batch data to the new tail")
	return nil
}
