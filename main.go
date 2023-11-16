package main

import (
	"cubeCache/cache"
	"cubeCache/http"
	"cubeCache/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	cubeCache := cache.New()
	cubeCache.NewCube("default", nil, nil, 10000)
	s := grpc.NewServer()
	rpc.RegisterCreateCubeServer(s, &rpc.Server{})
	rpc.RegisterGetValueServer(s, &rpc.Server{})
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	http.Start()
}
