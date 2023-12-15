// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: protobuf/cluster.proto

package cluster

import (
	rpc "cubecache/rpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RegisterNodeRequest) Reset() {
	*x = RegisterNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeRequest) ProtoMessage() {}

func (x *RegisterNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeRequest.ProtoReflect.Descriptor instead.
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterNodeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RegisterNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Cubes   []*rpc.CreateCubeRequest `protobuf:"bytes,2,rep,name=cubes,proto3" json:"cubes,omitempty"`
}

func (x *RegisterNodeResponse) Reset() {
	*x = RegisterNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeResponse) ProtoMessage() {}

func (x *RegisterNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeResponse.ProtoReflect.Descriptor instead.
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterNodeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterNodeResponse) GetCubes() []*rpc.CreateCubeRequest {
	if x != nil {
		return x.Cubes
	}
	return nil
}

type SendHeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // address with cube-service port
}

func (x *SendHeartbeatRequest) Reset() {
	*x = SendHeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendHeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendHeartbeatRequest) ProtoMessage() {}

func (x *SendHeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendHeartbeatRequest.ProtoReflect.Descriptor instead.
func (*SendHeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *SendHeartbeatRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type SendHeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CleanCubes bool `protobuf:"varint,1,opt,name=cleanCubes,proto3" json:"cleanCubes,omitempty"`
}

func (x *SendHeartbeatResponse) Reset() {
	*x = SendHeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendHeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendHeartbeatResponse) ProtoMessage() {}

func (x *SendHeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendHeartbeatResponse.ProtoReflect.Descriptor instead.
func (*SendHeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *SendHeartbeatResponse) GetCleanCubes() bool {
	if x != nil {
		return x.CleanCubes
	}
	return false
}

var File_protobuf_cluster_proto protoreflect.FileDescriptor

var file_protobuf_cluster_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5a,
	0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x28, 0x0a, 0x05, 0x63, 0x75, 0x62, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x05, 0x63, 0x75, 0x62, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x65,
	0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x15,
	0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x75,
	0x62, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x6e,
	0x43, 0x75, 0x62, 0x65, 0x73, 0x32, 0x86, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12,
	0x15, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13,
	0x5a, 0x11, 0x63, 0x75, 0x62, 0x65, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_cluster_proto_rawDescOnce sync.Once
	file_protobuf_cluster_proto_rawDescData = file_protobuf_cluster_proto_rawDesc
)

func file_protobuf_cluster_proto_rawDescGZIP() []byte {
	file_protobuf_cluster_proto_rawDescOnce.Do(func() {
		file_protobuf_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_cluster_proto_rawDescData)
	})
	return file_protobuf_cluster_proto_rawDescData
}

var file_protobuf_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_cluster_proto_goTypes = []interface{}{
	(*RegisterNodeRequest)(nil),   // 0: RegisterNodeRequest
	(*RegisterNodeResponse)(nil),  // 1: RegisterNodeResponse
	(*SendHeartbeatRequest)(nil),  // 2: SendHeartbeatRequest
	(*SendHeartbeatResponse)(nil), // 3: SendHeartbeatResponse
	(*rpc.CreateCubeRequest)(nil), // 4: CreateCubeRequest
}
var file_protobuf_cluster_proto_depIdxs = []int32{
	4, // 0: RegisterNodeResponse.cubes:type_name -> CreateCubeRequest
	0, // 1: Cluster.RegisterNode:input_type -> RegisterNodeRequest
	2, // 2: Cluster.SendHeartbeat:input_type -> SendHeartbeatRequest
	1, // 3: Cluster.RegisterNode:output_type -> RegisterNodeResponse
	3, // 4: Cluster.SendHeartbeat:output_type -> SendHeartbeatResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_cluster_proto_init() }
func file_protobuf_cluster_proto_init() {
	if File_protobuf_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterNodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendHeartbeatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendHeartbeatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_cluster_proto_goTypes,
		DependencyIndexes: file_protobuf_cluster_proto_depIdxs,
		MessageInfos:      file_protobuf_cluster_proto_msgTypes,
	}.Build()
	File_protobuf_cluster_proto = out.File
	file_protobuf_cluster_proto_rawDesc = nil
	file_protobuf_cluster_proto_goTypes = nil
	file_protobuf_cluster_proto_depIdxs = nil
}
