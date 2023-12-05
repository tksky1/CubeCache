package main

import (
	"context"
	"cubeCache/cache"
	"cubeCache/rpc"
	"errors"
)

type CubeNode struct {
	*rpc.UnimplementedCubeServer
	cache *cache.CubeCache
}

func (n CubeNode) Get(ctx context.Context, req *rpc.GetValueRequest) (res *rpc.GetValueResponse, err error) {
	panic("implement me")
}

func (n CubeNode) Set(ctx context.Context, req *rpc.SetValueRequest) (res *rpc.SetValueResponse, err error) {
	cube, ok := n.cache.GetCube(req.CubeName)
	if ok {
		cube.Set(req.Key, req.Value, req.GetterFunc)
		return nil, nil
	}
	return nil, errors.New("cube not exist")
}
