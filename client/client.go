package client

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
)
import "cubeCache/rpc"

type CubeClient struct {
	Conn    *grpc.ClientConn
	cli     rpc.CubeClient
	ctrlCli rpc.CubeControlClient
}

func NewCubeClient(cubeServer string, ctrlServer string, opts ...grpc.DialOption) *CubeClient {
	c, err := tls.LoadX509KeyPair("../master/cert", "../master/key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	transportCredits := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{c},
		InsecureSkipVerify: true,
	})
	options := append(make([]grpc.DialOption, 0))
	conn, err := grpc.Dial(cubeServer, append(append(options, opts...), grpc.WithTransportCredentials(transportCredits))...)
	if err != nil {
		log.Fatalf("Failed to connect to cube server: %v", err)
	}
	connCtrl, err := grpc.Dial(ctrlServer, append(append(options, opts...), grpc.WithInsecure())...)
	if err != nil {
		log.Fatalf("Failed to connect to cube control server: %v", err)
	}
	ret := &CubeClient{Conn: conn, cli: rpc.NewCubeClient(conn), ctrlCli: rpc.NewCubeControlClient(connCtrl)}
	return ret
}

func (c *CubeClient) Set(request *rpc.SetValueRequest) (err error) {
	ctx := context.Background()
	md := metadata.New(map[string]string{"cube_cache_key": request.Key, "content-type": "application/grpc"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err = c.cli.Set(ctx, request)
	return err
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

	res, err := c.ctrlCli.CreateCube(ctx, request)
	if err != nil {
		return false, "rpc error", err
	}
	return res.Success, res.Message, nil
}
