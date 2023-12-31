// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: protobuf/cube.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Cube_Get_FullMethodName = "/Cube/Get"
	Cube_Set_FullMethodName = "/Cube/Set"
)

// CubeClient is the client API for Cube service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CubeClient interface {
	Get(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error)
	Set(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error)
}

type cubeClient struct {
	cc grpc.ClientConnInterface
}

func NewCubeClient(cc grpc.ClientConnInterface) CubeClient {
	return &cubeClient{cc}
}

func (c *cubeClient) Get(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error) {
	out := new(GetValueResponse)
	err := c.cc.Invoke(ctx, Cube_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cubeClient) Set(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error) {
	out := new(SetValueResponse)
	err := c.cc.Invoke(ctx, Cube_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CubeServer is the server API for Cube service.
// All implementations should embed UnimplementedCubeServer
// for forward compatibility
type CubeServer interface {
	Get(context.Context, *GetValueRequest) (*GetValueResponse, error)
	Set(context.Context, *SetValueRequest) (*SetValueResponse, error)
}

// UnimplementedCubeServer should be embedded to have forward compatible implementations.
type UnimplementedCubeServer struct {
}

func (UnimplementedCubeServer) Get(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCubeServer) Set(context.Context, *SetValueRequest) (*SetValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}

// UnsafeCubeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CubeServer will
// result in compilation errors.
type UnsafeCubeServer interface {
	mustEmbedUnimplementedCubeServer()
}

func RegisterCubeServer(s grpc.ServiceRegistrar, srv CubeServer) {
	s.RegisterService(&Cube_ServiceDesc, srv)
}

func _Cube_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CubeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cube_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CubeServer).Get(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cube_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CubeServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cube_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CubeServer).Set(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cube_ServiceDesc is the grpc.ServiceDesc for Cube service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cube_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Cube",
	HandlerType: (*CubeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Cube_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Cube_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/cube.proto",
}

const (
	CubeControl_CreateCube_FullMethodName = "/CubeControl/CreateCube"
)

// CubeControlClient is the client API for CubeControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CubeControlClient interface {
	CreateCube(ctx context.Context, in *CreateCubeRequest, opts ...grpc.CallOption) (*CreateCubeResponse, error)
}

type cubeControlClient struct {
	cc grpc.ClientConnInterface
}

func NewCubeControlClient(cc grpc.ClientConnInterface) CubeControlClient {
	return &cubeControlClient{cc}
}

func (c *cubeControlClient) CreateCube(ctx context.Context, in *CreateCubeRequest, opts ...grpc.CallOption) (*CreateCubeResponse, error) {
	out := new(CreateCubeResponse)
	err := c.cc.Invoke(ctx, CubeControl_CreateCube_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CubeControlServer is the server API for CubeControl service.
// All implementations should embed UnimplementedCubeControlServer
// for forward compatibility
type CubeControlServer interface {
	CreateCube(context.Context, *CreateCubeRequest) (*CreateCubeResponse, error)
}

// UnimplementedCubeControlServer should be embedded to have forward compatible implementations.
type UnimplementedCubeControlServer struct {
}

func (UnimplementedCubeControlServer) CreateCube(context.Context, *CreateCubeRequest) (*CreateCubeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCube not implemented")
}

// UnsafeCubeControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CubeControlServer will
// result in compilation errors.
type UnsafeCubeControlServer interface {
	mustEmbedUnimplementedCubeControlServer()
}

func RegisterCubeControlServer(s grpc.ServiceRegistrar, srv CubeControlServer) {
	s.RegisterService(&CubeControl_ServiceDesc, srv)
}

func _CubeControl_CreateCube_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCubeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CubeControlServer).CreateCube(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CubeControl_CreateCube_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CubeControlServer).CreateCube(ctx, req.(*CreateCubeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CubeControl_ServiceDesc is the grpc.ServiceDesc for CubeControl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CubeControl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CubeControl",
	HandlerType: (*CubeControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCube",
			Handler:    _CubeControl_CreateCube_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/cube.proto",
}
