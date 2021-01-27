package network

import (
	"context"

	"github.com/JDJFisher/distributed-storage/protos"
)

type NetworkServer struct {
	protos.UnimplementedNetworkServer
	CandidateNodes []*CandidateNode
}

func (s *NetworkServer) RequestJoin(ctx context.Context, req *protos.RequestJoinRequest) (*protos.RequestJoinResponse, error) {
	newCandidateNode := &CandidateNode{name: req.ServiceName}
	s.CandidateNodes = append(s.CandidateNodes, newCandidateNode)

	return &protos.RequestJoinResponse{Ok: true}, nil
}
