package main

import (
	"crypto/tls"
	"cubeCache/cache"
	"cubeCache/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	println("CubeCache Node Initiating..")

	cubeCache := cache.New()
	cubeCache.NewCube(&rpc.CreateCubeRequest{
		CubeName:       "default",
		MaxBytes:       102400,
		CubeGetterFunc: nil,
		OnEvictedFunc:  nil,
		DelayWrite:     nil,
	})

	c, err := tls.LoadX509KeyPair("master/cert", "master/key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	transportCredits := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{c},
		InsecureSkipVerify: true,
	})
	s := grpc.NewServer(grpc.Creds(transportCredits))

	rpc.RegisterCubeServer(s, &CubeNode{cache: cubeCache})

	lis, err := net.Listen("tcp", "0.0.0.0:4012")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	println("listening at 0.0.0.0:4012")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
