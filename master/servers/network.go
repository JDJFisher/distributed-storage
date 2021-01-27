package servers

import (
	"context"

	"github.com/JDJFisher/distributed-storage/protos"
)

// NetworkServer ...
type NetworkServer struct {
	protos.UnimplementedNetworkServer
	CandidateNodes []*CandidateNode
}

func (s *NetworkServer) RequestJoin(ctx context.Context, req *protos.RequestJoinRequest) (*protos.RequestJoinResponse, error) {
	newCandidateNode := &CandidateNode{name: req.ServiceName}
	s.CandidateNodes = append(s.CandidateNodes, newCandidateNode)

	return &protos.RequestJoinResponse{Ok: true}, nil
}

// // JoinNetwork ...
// func (s *NetworkServer) JoinNetwork(ctx context.Context, req *protos.) (*protos.NetworkJoinResponse, error) {
// 	log.Println("Received a network join request")
// 	return &protos.NetworkJoinResponse{Type: protos.NetworkJoinResponse_NORMAL}, nil
// }
