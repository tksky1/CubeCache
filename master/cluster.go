package main

import (
	"context"
	"cubeCache/cache"
	"cubeCache/rpc"
	"net/http/httputil"
)

type CubeMaster struct {
	*rpc.UnimplementedCubeControlServer
	mapper *cache.Mapper
	cache  *cache.CubeCache
	proxy  *httputil.ReverseProxy
}

func (m *CubeMaster) CreateCube(ctx context.Context, in *rpc.CreateCubeRequest) (*rpc.CreateCubeResponse, error) {
	//TODO implement me
	return &rpc.CreateCubeResponse{}, nil
}
