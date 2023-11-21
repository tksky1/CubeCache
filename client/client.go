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
	rpc.CubeClient
}

func NewCubeClient(target string, opts ...grpc.DialOption) *CubeClient {
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	ret := &CubeClient{Conn: conn, cli: rpc.NewCubeClient(conn)}
	return ret
}

func (c *CubeClient) Set(cubeName string, key string, value []byte, getterFunc string) (ok bool, msg string, connErr error) {

	ctx := context.Background()
	md := metadata.New(map[string]string{"cube_cache_key": key})

	// 将 Metadata 对象附加到上下文中
	ctx = metadata.NewOutgoingContext(ctx, md)

	req := &rpc.SetValueRequest{
		CubeName:   cubeName,
		Key:        key,
		Value:      value,
		GetterFunc: nil,
	}
	if getterFunc != "" {
		req.GetterFunc = &getterFunc
	}

	res, err := c.cli.Set(ctx, req)
	return res.Ok, res.Message, err
}
