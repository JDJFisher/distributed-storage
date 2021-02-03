package chain

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

// UpdateNeighbours - Tell a node about it's neighbour nodes changing
func (node *Node) UpdateNeighbours(predAddress string, succAddress string) {
	// Open a connection to the predecessor node
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(node.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to node - %v", err.Error())
	}
	defer conn.Close()

	// Create storage client
	networkClient := protos.NewChainClient(conn)

	request := &protos.NeighbourInfo{
		PredAddress: predAddress,
		SuccAddress: succAddress,
	}

	_, err = networkClient.UpdateNeighbours(context.Background(), request)
	if err != nil {
		log.Fatalf("Error updating node - %v", err.Error())
	}
}
