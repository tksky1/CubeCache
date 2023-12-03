package main

import (
	"context"
	"cubeCache/cube"
	"cubeCache/rpc"
)

type CubeNode struct {
	*rpc.UnimplementedCubeServer
	cube *cube.Cube
}

func (n CubeNode) Get(ctx context.Context, req *rpc.GetValueRequest) (res *rpc.GetValueResponse, err error) {
	panic("implement me")
}

func (n CubeNode) Set(ctx context.Context, req *rpc.SetValueRequest) (res *rpc.SetValueResponse, err error) {
	//TODO implement me
	panic("implement me")
}
