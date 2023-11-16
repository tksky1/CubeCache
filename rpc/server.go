package rpc

import (
	"context"
	"cubeCache/cache"
)

type Server struct {
	cubeCache *cache.CubeCache
}

func (s *Server) Get(ctx context.Context, request *GetValueRequest) (*GetValueResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CreateCube(ctx context.Context, request *CreateCubeRequest) (*CreateCubeResponse, error) {
	//TODO implement me
	panic("implement me")
}
