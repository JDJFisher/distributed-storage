package servers

import (
	"context"
	"log"
	"os"

	"github.com/JDJFisher/distributed-storage/master/chain"
	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

// StorageServer ...
type StorageServer struct {
	protos.UnimplementedStorageServer
	Chain *chain.Chain
}

func NewStorageServer(chain *chain.Chain) *StorageServer {
	return &StorageServer{Chain: chain}
}

func (s *StorageServer) Read(ctx context.Context, req *protos.ReadRequest) (*protos.ReadResponse, error) {
	log.Println("Received a read request")

	// Different grpc connection info depending on if it's running in docker or not
	grpcHost := ":7000"
	if os.Getenv("docker") == "true" {
		grpcHost = s.Chain.Tail.Address
	}

	// Open a connection to the tail node
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
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
	}

	return response, nil
}

func (s *StorageServer) Write(ctx context.Context, req *protos.WriteRequest) (*protos.WriteResponse, error) {
	log.Println("Received a write request")

	// Different grpc connection info depending on if it's running in docker or not
	grpcHost := ":7000"
	if os.Getenv("docker") == "true" {
		grpcHost = s.Chain.Head.Address
	}

	// Open a connection to the head node
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the head node - %v", err.Error())
	}
	defer conn.Close()

	// Create storage client
	storageClient := protos.NewStorageClient(conn)

	// Forward write request to the head
	response, err := storageClient.Write(context.Background(), req)
	if err != nil {
		log.Fatalf("Error forwarding write request to the head node - %v", err.Error())
	}

	return response, nil
}
