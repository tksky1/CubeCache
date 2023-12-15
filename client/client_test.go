package client

import (
	"cubeCache/rpc"
	"google.golang.org/grpc"
	"testing"
)

func TestClient(t *testing.T) {
	option := make([]grpc.DialOption, 0)
	cli := NewCubeClient("127.0.0.1:4010", "127.0.0.1:4011", option...)
	_, _, err := cli.CreateCube(&rpc.CreateCubeRequest{
		CubeName:       "default",
		MaxBytes:       10000,
		CubeGetterFunc: nil,
		OnEvictedFunc:  nil,
		DelayWrite:     nil,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = cli.Set(&rpc.SetValueRequest{
		CubeName:   "default",
		Key:        "key1",
		Value:      []byte("abc"),
		GetterFunc: nil,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	ok, value, _, err := cli.Get(&rpc.GetValueRequest{
		CubeName: "default",
		Key:      "key1",
	})
	if !ok || err != nil {
		t.Fatalf(err.Error())
	}
	if string(value) != "abc" {
		t.Fail()
	}

}
