package adder

import (
	"context"
	"fmt"
	"go-grpc-graph/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{}

// ADD ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	fmt.Println(req.GetX(), req.GetY())
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
