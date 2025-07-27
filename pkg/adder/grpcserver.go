package adder

import (
	"context"

	"github.com/lypolix/grpc-calculator/pkg/api"
)

// GRPCServer ...

type GRPCServer struct {
	api.UnimplementedAdderServer // Обязательное встраивание!
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error){
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

