package main

import (
	"context"
	"cubeCache/cache"
	"cubeCache/cluster"
	"cubeCache/rpc"
	"net/http/httputil"
)

type CubeMaster struct {
	*rpc.UnimplementedCubeControlServer
	*cluster.UnimplementedClusterServer
	mapper   *cache.Mapper
	cache    *cache.CubeCache
	proxy    *httputil.ReverseProxy
	cubeList []*rpc.CreateCubeRequest
}

func (m *CubeMaster) CreateCube(ctx context.Context, in *rpc.CreateCubeRequest) (*rpc.CreateCubeResponse, error) {
	//TODO implement me
	return &rpc.CreateCubeResponse{}, nil
}
