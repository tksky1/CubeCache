package main

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

func director(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	keys, ok := md["cube_cache_key"]
	if !ok || len(keys) != 1 {
		return nil, nil, errors.New("cube proxy error: no cube_cache_key header")
	}
	target := server.mapper.Get(keys[0])
	//TODO: DELETE ME
	target = "127.0.0.1:4012"

	outCtx := metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.DialContext(outCtx, target, grpc.WithCodec(proxy.Codec()), grpc.WithTransportCredentials(transportCredits))
	if err != nil {
		return nil, nil, err
	}
	return ctx, conn, nil
}

func runProxy() {
	c, err := tls.LoadX509KeyPair("./master/cert", "./master/key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	transportCredits = credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{c},
		InsecureSkipVerify: true,
	})

	proxyServer := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
		grpc.Creds(transportCredits),
	)

	println("rpc proxy listening on 0.0.0.0:4010")
	lis, err := net.Listen("tcp", "0.0.0.0:4010")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := proxyServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
