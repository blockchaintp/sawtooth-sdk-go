// Copyright 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protobuf/client_peers_pb2/client_peers.proto

package client_peer

import (
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

type ClientPeersGetResponse_Status int32

const (
	ClientPeersGetResponse_STATUS_UNSET ClientPeersGetResponse_Status = 0
	ClientPeersGetResponse_OK           ClientPeersGetResponse_Status = 1
	ClientPeersGetResponse_ERROR        ClientPeersGetResponse_Status = 2
)

// Enum value maps for ClientPeersGetResponse_Status.
var (
	ClientPeersGetResponse_Status_name = map[int32]string{
		0: "STATUS_UNSET",
		1: "OK",
		2: "ERROR",
	}
	ClientPeersGetResponse_Status_value = map[string]int32{
		"STATUS_UNSET": 0,
		"OK":           1,
		"ERROR":        2,
	}
)

func (x ClientPeersGetResponse_Status) Enum() *ClientPeersGetResponse_Status {
	p := new(ClientPeersGetResponse_Status)
	*p = x
	return p
}

func (x ClientPeersGetResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientPeersGetResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_client_peers_pb2_client_peers_proto_enumTypes[0].Descriptor()
}

func (ClientPeersGetResponse_Status) Type() protoreflect.EnumType {
	return &file_protobuf_client_peers_pb2_client_peers_proto_enumTypes[0]
}

func (x ClientPeersGetResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientPeersGetResponse_Status.Descriptor instead.
func (ClientPeersGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_client_peers_pb2_client_peers_proto_rawDescGZIP(), []int{1, 0}
}

type ClientPeersGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientPeersGetRequest) Reset() {
	*x = ClientPeersGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_peers_pb2_client_peers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPeersGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPeersGetRequest) ProtoMessage() {}

func (x *ClientPeersGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_peers_pb2_client_peers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPeersGetRequest.ProtoReflect.Descriptor instead.
func (*ClientPeersGetRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_client_peers_pb2_client_peers_proto_rawDescGZIP(), []int{0}
}

type ClientPeersGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ClientPeersGetResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ClientPeersGetResponse_Status" json:"status,omitempty"`
	Peers  []string                      `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *ClientPeersGetResponse) Reset() {
	*x = ClientPeersGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_client_peers_pb2_client_peers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPeersGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPeersGetResponse) ProtoMessage() {}

func (x *ClientPeersGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_client_peers_pb2_client_peers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPeersGetResponse.ProtoReflect.Descriptor instead.
func (*ClientPeersGetResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_client_peers_pb2_client_peers_proto_rawDescGZIP(), []int{1}
}

func (x *ClientPeersGetResponse) GetStatus() ClientPeersGetResponse_Status {
	if x != nil {
		return x.Status
	}
	return ClientPeersGetResponse_STATUS_UNSET
}

func (x *ClientPeersGetResponse) GetPeers() []string {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_protobuf_client_peers_pb2_client_peers_proto protoreflect.FileDescriptor

var file_protobuf_client_peers_pb2_client_peers_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x5f, 0x70, 0x62, 0x32, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17,
	0x0a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x16, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x22, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42,
	0x26, 0x0a, 0x15, 0x73, 0x61, 0x77, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x0b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_client_peers_pb2_client_peers_proto_rawDescOnce sync.Once
	file_protobuf_client_peers_pb2_client_peers_proto_rawDescData = file_protobuf_client_peers_pb2_client_peers_proto_rawDesc
)

func file_protobuf_client_peers_pb2_client_peers_proto_rawDescGZIP() []byte {
	file_protobuf_client_peers_pb2_client_peers_proto_rawDescOnce.Do(func() {
		file_protobuf_client_peers_pb2_client_peers_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_client_peers_pb2_client_peers_proto_rawDescData)
	})
	return file_protobuf_client_peers_pb2_client_peers_proto_rawDescData
}

var file_protobuf_client_peers_pb2_client_peers_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_client_peers_pb2_client_peers_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_client_peers_pb2_client_peers_proto_goTypes = []interface{}{
	(ClientPeersGetResponse_Status)(0), // 0: ClientPeersGetResponse.Status
	(*ClientPeersGetRequest)(nil),      // 1: ClientPeersGetRequest
	(*ClientPeersGetResponse)(nil),     // 2: ClientPeersGetResponse
}
var file_protobuf_client_peers_pb2_client_peers_proto_depIdxs = []int32{
	0, // 0: ClientPeersGetResponse.status:type_name -> ClientPeersGetResponse.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_client_peers_pb2_client_peers_proto_init() }
func file_protobuf_client_peers_pb2_client_peers_proto_init() {
	if File_protobuf_client_peers_pb2_client_peers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_client_peers_pb2_client_peers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPeersGetRequest); i {
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
		file_protobuf_client_peers_pb2_client_peers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPeersGetResponse); i {
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
			RawDescriptor: file_protobuf_client_peers_pb2_client_peers_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_client_peers_pb2_client_peers_proto_goTypes,
		DependencyIndexes: file_protobuf_client_peers_pb2_client_peers_proto_depIdxs,
		EnumInfos:         file_protobuf_client_peers_pb2_client_peers_proto_enumTypes,
		MessageInfos:      file_protobuf_client_peers_pb2_client_peers_proto_msgTypes,
	}.Build()
	File_protobuf_client_peers_pb2_client_peers_proto = out.File
	file_protobuf_client_peers_pb2_client_peers_proto_rawDesc = nil
	file_protobuf_client_peers_pb2_client_peers_proto_goTypes = nil
	file_protobuf_client_peers_pb2_client_peers_proto_depIdxs = nil
}