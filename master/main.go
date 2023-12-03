package main

import (
	"crypto/tls"
	"cubeCache/cache"
	"cubeCache/rpc"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"path/filepath"
)

var server *CubeMaster

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

	server = &CubeMaster{
		mapper: cache.NewMapper(3),
		cache:  cubeCache,
	}

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

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				NextProtos:         []string{"h2"},
				InsecureSkipVerify: true,
			},
		},
	}

	proxy := &httputil.ReverseProxy{
		Director: proxyDirector,
		ErrorHandler: func(writer http.ResponseWriter, request *http.Request, err error) {
			log.Println(err.Error() + "\nreq:" + request.URL.RawPath)
		},
		Transport: client.Transport,
	}
	//proxy.Transport = &http.Transport{
	//	TLSClientConfig: &tls.Config{
	//		InsecureSkipVerify: true,
	//	},
	//	ForceAttemptHTTP2: true,
	//}

	server := &http.Server{
		Addr:         ":4010",
		Handler:      proxy,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	err = http2.ConfigureServer(server, nil)
	if err != nil {
		log.Fatal("configure http2 error:", err.Error())
	}

	println("rpc proxy listening on 0.0.0.0:4010")
	if err := server.ListenAndServeTLS(filepath.Join("master", "cert"), filepath.Join("master", "key")); err != nil {
		log.Fatal("Start proxy server failed:", err)
	}
}
