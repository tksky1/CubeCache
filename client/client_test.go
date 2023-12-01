package client

import (
	"cubeCache/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestClient(t *testing.T) {
	cli := NewCubeClient("127.0.0.1:4010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	_, _, err := cli.CreateCube(&rpc.CreateCubeRequest{
		CubeName:       "test",
		MaxBytes:       10000,
		CubeGetterFunc: nil,
		OnEvictedFunc:  nil,
		DelayWrite:     nil,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	ok, _, err := cli.Set(&rpc.SetValueRequest{
		CubeName:   "test",
		Key:        "key1",
		Value:      []byte("abc"),
		GetterFunc: nil,
	})
	if !ok || err != nil {
		t.Fatalf(err.Error())
	}
	ok, value, _, err := cli.Get(&rpc.GetValueRequest{
		CubeName: "test",
		Key:      "key1",
	})
	if !ok || err != nil {
		t.Fatalf(err.Error())
	}
	if string(value) != "abc" {
		t.Fail()
	}

}
