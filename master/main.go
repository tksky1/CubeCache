package main

import (
	"cubeCache/cache"
	"cubeCache/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

var (
	server           *CubeMaster
	transportCredits credentials.TransportCredentials
)

func main() {
	println("CubeCache Master Initiating..")

	cubeCache := cache.New()
	defaultCubeReq := &rpc.CreateCubeRequest{
		CubeName:       "default",
		MaxBytes:       102400,
		CubeGetterFunc: nil,
		OnEvictedFunc:  nil,
		DelayWrite:     nil,
	}
	cubeCache.NewCube(defaultCubeReq)
	s := grpc.NewServer()

	server = &CubeMaster{
		mapper:   cache.NewMapper(3),
		cache:    cubeCache,
		cubeList: make([]*rpc.CreateCubeRequest, 1),
	}
	server.cubeList[0] = defaultCubeReq

	rpc.RegisterCubeControlServer(s, server)

	lis, err := net.Listen("tcp", "0.0.0.0:4011")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	println("cubeControl listening at 0.0.0.0:4011")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	runProxy()
}
