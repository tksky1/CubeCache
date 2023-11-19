package rpc

import "google.golang.org/grpc"

func Get() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

}
