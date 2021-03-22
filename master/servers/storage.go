package servers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JDJFisher/distributed-storage/master/chain"
	"github.com/JDJFisher/distributed-storage/protos"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

type StorageServer struct {
	protos.UnimplementedStorageServer
	Chain *chain.Chain
}

func NewStorageServer(chain *chain.Chain) *StorageServer {
	return &StorageServer{Chain: chain}
}

func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {
	log.Println("Received a read request")

	// Add to the pending in the master
	uid := uuid.NewV4()
	s.Chain.Pending = append(s.Chain.Pending, uid)

	// Open a connection to the tail node
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(s.Chain.GetTailAddress(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the tail node - %v", err.Error())
	}
	defer conn.Close()

	// Create storage client
	storageClient := protos.NewStorageClient(conn)

	// Forward read request to the tail
	response, err := storageClient.Read(context.Background(), req)
	if err != nil {
		log.Fatalf("Error forwarding read request to the tail node - %v", err.Error())
	} else {
		s.Chain.RemoveUUIDFromPending(uid)
	}

	return response, nil
}

func (s *StorageServer) Write(ctx context.Context, req *protos.WriteRequest) (*protos.OkResponse, error) {
	log.Println("Received a write request")

	// Add to the pending in the master
	uid := uuid.NewV4()
	s.Chain.Pending = append(s.Chain.Pending, uid)
	req.Uuid = uid.String()

	// Open a connection to the head node
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(s.Chain.GetHeadAddress(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the head node - %v", err.Error())
	}
	defer conn.Close()

	// Create storage client
	storageClient := protos.NewStorageClient(conn)

	// Forward write request to the head
	_, err = storageClient.Write(context.Background(), req)
	if err != nil {
		log.Fatalf("Error forwarding write request to the head node - %v", err.Error())
	}

	start := time.Now()

	for s.Chain.IsInPending(uid) {
		if time.Since(start) > 5*time.Second {
			return nil, fmt.Errorf("timeout")
		}
	}

	return &protos.OkResponse{}, nil
}

func (s *StorageServer) Processed(ctx context.Context, req *protos.ProcessedRequest) (*protos.OkResponse, error) {
	uid, err := uuid.FromString(req.Uuid)
	if err != nil {
		panic("big issues")
	}

	// Tell the nodes that the operation can be persisted
	for _, node := range s.Chain.Nodes {

		// Open a connection to the node
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(node.Address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Error connecting to node - %v", err.Error())
		}
		defer conn.Close()

		// Create storage client
		storageClient := protos.NewStorageClient(conn)

		_, err = storageClient.Persist(context.Background(), req)
		if err != nil {
			// not an actual problem because it died (due to assumptions)
		}
	}

	s.Chain.RemoveUUIDFromPending(uid)

	return &protos.OkResponse{}, nil
}
