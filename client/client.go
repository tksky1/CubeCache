package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)
import "cubeCache/rpc"

type CubeClient struct {
	Conn *grpc.ClientConn
	cli  rpc.CubeClient
}

func NewCubeClient(target string, opts ...grpc.DialOption) *CubeClient {
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	ret := &CubeClient{Conn: conn, cli: rpc.NewCubeClient(conn)}
	return ret
}

func (c *CubeClient) Set(request *rpc.SetValueRequest) (ok bool, msg string, connErr error) {
	ctx := context.Background()
	md := metadata.New(map[string]string{"cube_cache_key": request.Key})

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := c.cli.Set(ctx, request)
	return res.Ok, res.Message, err
}

func (c *CubeClient) Get(request *rpc.GetValueRequest) (ok bool, value []byte, msg string, connErr error) {
	ctx := context.Background()
	md := metadata.New(map[string]string{"cube_cache_key": request.Key})

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := c.cli.Get(ctx, request)
	if err != nil {
		return false, nil, "rpc error", err
	}
	return res.Ok, res.Value, res.Message, nil
}

func (c *CubeClient) CreateCube(request *rpc.CreateCubeRequest) (ok bool, msg string, connErr error) {
	ctx := context.Background()
	md := metadata.New(map[string]string{"cube_cache_create_cube": "true"})

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := c.cli.CreateCube(ctx, request)
	if err != nil {
		return false, "rpc error", err
	}
	return res.Success, res.Message, nil
}
