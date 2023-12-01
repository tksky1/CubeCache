package main

import (
	"cubeCache/cache"
	"cubeCache/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	println("CubeCache Master Initiating..")

	cubeCache := cache.New()
	cubeCache.NewCube(&rpc.CreateCubeRequest{
		CubeName:       "default",
		MaxBytes:       102400,
		CubeGetterFunc: nil,
		OnEvictedFunc:  nil,
		DelayWrite:     nil,
	})
	s := grpc.NewServer()

	proxy := &httputil.ReverseProxy{
		Director: ,
	}

	rpc.RegisterCubeControlServer(s, &CubeMaster{
		mapper: cache.NewMapper(3),
		cache:  cubeCache,
	})

	lis, err := net.Listen("tcp", "0.0.0.0:4010")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	println("listening at 0.0.0.0:4010")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	//TODO: grpc拦截器
}
