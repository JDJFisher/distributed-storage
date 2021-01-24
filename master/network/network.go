package network

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedNetworkServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message from client: %s", message.Body)
	return &Message{Body: "Hello from the master!"}, nil
}
