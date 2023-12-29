package main

import (
	"context"
	"cubeCache/cache"
	"cubeCache/rpc"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type CubeNode struct {
	*rpc.UnimplementedCubeServer
	cache *cache.CubeCache
}

func (n CubeNode) Get(ctx context.Context, req *rpc.GetValueRequest) (res *rpc.GetValueResponse, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	keys, ok := md["cube_cache_key"]
	cube, ok := n.cache.GetCube(req.CubeName)
	if ok {
		ret, ok := cube.Get(keys[0])
		logrus.Debugf("get value for key %s success: %s", req.Key, string(ret))
		return &rpc.GetValueResponse{
			Ok:      ok,
			Value:   ret,
			Message: "",
		}, nil
	}
	return &rpc.GetValueResponse{}, errors.New("cube not exist")
}

func (n CubeNode) Set(ctx context.Context, req *rpc.SetValueRequest) (res *rpc.SetValueResponse, err error) {
	cube, ok := n.cache.GetCube(req.CubeName)
	if ok {
		cube.Set(req.Key, req.Value, req.GetterFunc)
		return &rpc.SetValueResponse{}, nil
	}
	return &rpc.SetValueResponse{}, errors.New("cube not exist")
}
