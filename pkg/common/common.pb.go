// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: common.proto

package common

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

type RPCStatus int32

const (
	RPCStatus_SUCCESS RPCStatus = 0
	RPCStatus_FAILURE RPCStatus = 1
)

// Enum value maps for RPCStatus.
var (
	RPCStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	RPCStatus_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x RPCStatus) Enum() *RPCStatus {
	p := new(RPCStatus)
	*p = x
	return p
}

func (x RPCStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RPCStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (RPCStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x RPCStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RPCStatus.Descriptor instead.
func (RPCStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type ProviderFeature int32

const (
	ProviderFeature_DEPLOY  ProviderFeature = 0
	ProviderFeature_DESTROY ProviderFeature = 1
)

// Enum value maps for ProviderFeature.
var (
	ProviderFeature_name = map[int32]string{
		0: "DEPLOY",
		1: "DESTROY",
	}
	ProviderFeature_value = map[string]int32{
		"DEPLOY":  0,
		"DESTROY": 1,
	}
)

func (x ProviderFeature) Enum() *ProviderFeature {
	p := new(ProviderFeature)
	*p = x
	return p
}

func (x ProviderFeature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProviderFeature) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (ProviderFeature) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x ProviderFeature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProviderFeature.Descriptor instead.
func (ProviderFeature) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

// Handshake (DO NOT MODIFY)
type HandshakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientVersion string `protobuf:"bytes,1,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
}

func (x *HandshakeRequest) Reset() {
	*x = HandshakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeRequest) ProtoMessage() {}

func (x *HandshakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeRequest.ProtoReflect.Descriptor instead.
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakeRequest) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

type HandshakeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerVersion string `protobuf:"bytes,1,opt,name=server_version,json=serverVersion,proto3" json:"server_version,omitempty"`
}

func (x *HandshakeReply) Reset() {
	*x = HandshakeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeReply) ProtoMessage() {}

func (x *HandshakeReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeReply.ProtoReflect.Descriptor instead.
func (*HandshakeReply) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *HandshakeReply) GetServerVersion() string {
	if x != nil {
		return x.ServerVersion
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39,
	0x0a, 0x10, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0e, 0x48, 0x61, 0x6e,
	0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2a, 0x25, 0x0a, 0x09, 0x52, 0x50, 0x43, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x2a, 0x2a, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x53, 0x54,
	0x52, 0x4f, 0x59, 0x10, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x62, 0x6c, 0x65, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x63, 0x62, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_proto_goTypes = []interface{}{
	(RPCStatus)(0),           // 0: RPCStatus
	(ProviderFeature)(0),     // 1: ProviderFeature
	(*HandshakeRequest)(nil), // 2: HandshakeRequest
	(*HandshakeReply)(nil),   // 3: HandshakeReply
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeRequest); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeReply); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
