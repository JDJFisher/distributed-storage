package network

import (
	"context"
	"log"
)

// Implement the interface of the grpc server
type NetServer struct {
	UnimplementedNetworkServer
}

func (server *NetServer) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message from client: %s", message.Body)
	return &Message{Body: "Hello from the master!"}, nil
}
