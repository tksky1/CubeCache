package main

import (
	"cubeCache/cache"
	"cubeCache/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	println("CubeCache Node Initiating..")

	cubeCache := cache.New()
	cubeCache.NewCube(&rpc.CreateCubeRequest{
		CubeName:       "default",
		MaxBytes:       102400,
		CubeGetterFunc: nil,
		OnEvictedFunc:  nil,
		DelayWrite:     nil,
	})
	s := grpc.NewServer()

	rpc.RegisterCubeServer(s, &CubeNode{})

	lis, err := net.Listen("tcp", "0.0.0.0:4011")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	println("listening at 0.0.0.0:4011")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
