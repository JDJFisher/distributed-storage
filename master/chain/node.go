package chain

import (
	"context"
	"log"

	"github.com/JDJFisher/distributed-storage/protos"
	"google.golang.org/grpc"
)

// Node - A server in the chain (either a head, replica or tail)
type Node struct {
	Address     string
	successor   *Node
	predecessor *Node
}

// NewNode - Create a new node object
func NewNode(address string, successor *Node, predecessor *Node) *Node {
	return &Node{address, successor, predecessor}
}

//Predecessor stuff

func (node *Node) GetPred() *Node {
	return node.predecessor
}

func (node *Node) GetPredAddress() string {
	if node.predecessor != nil {
		return node.predecessor.Address
	}
	return ""
}

func (node *Node) SetPred(newPred *Node) {
	node.predecessor = newPred
}

//Successor stuff

func (node *Node) GetSucc() *Node {
	return node.successor
}

func (node *Node) GetSuccAddress() string {
	if node.successor != nil {
		return node.successor.Address
	}
	return ""
}

func (node *Node) SetSucc(newSucc *Node) {
	node.successor = newSucc
}

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
