// Setup gRPC protobuf meta

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: provider.proto

package provider

import (
	common "github.com/cble-platform/cble-provider-grpc/pkg/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Deploy
type DeployRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId string           `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Blueprint    []byte           `protobuf:"bytes,2,opt,name=blueprint,proto3" json:"blueprint,omitempty"`
	TemplateVars *structpb.Struct `protobuf:"bytes,3,opt,name=templateVars,proto3" json:"templateVars,omitempty"`
}

func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployRequest) ProtoMessage() {}

func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployRequest.ProtoReflect.Descriptor instead.
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{0}
}

func (x *DeployRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *DeployRequest) GetBlueprint() []byte {
	if x != nil {
		return x.Blueprint
	}
	return nil
}

func (x *DeployRequest) GetTemplateVars() *structpb.Struct {
	if x != nil {
		return x.TemplateVars
	}
	return nil
}

type DeployReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId    string           `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Status          common.RPCStatus `protobuf:"varint,2,opt,name=status,proto3,enum=RPCStatus" json:"status,omitempty"`
	DeploymentState *structpb.Struct `protobuf:"bytes,3,opt,name=deployment_state,json=deploymentState,proto3" json:"deployment_state,omitempty"`
	DeploymentVars  *structpb.Struct `protobuf:"bytes,4,opt,name=deployment_vars,json=deploymentVars,proto3" json:"deployment_vars,omitempty"`
}

func (x *DeployReply) Reset() {
	*x = DeployReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployReply) ProtoMessage() {}

func (x *DeployReply) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployReply.ProtoReflect.Descriptor instead.
func (*DeployReply) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{1}
}

func (x *DeployReply) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *DeployReply) GetStatus() common.RPCStatus {
	if x != nil {
		return x.Status
	}
	return common.RPCStatus(0)
}

func (x *DeployReply) GetDeploymentState() *structpb.Struct {
	if x != nil {
		return x.DeploymentState
	}
	return nil
}

func (x *DeployReply) GetDeploymentVars() *structpb.Struct {
	if x != nil {
		return x.DeploymentVars
	}
	return nil
}

// Destroy
type DestroyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId    string           `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Blueprint       []byte           `protobuf:"bytes,2,opt,name=blueprint,proto3" json:"blueprint,omitempty"`
	TemplateVars    *structpb.Struct `protobuf:"bytes,3,opt,name=templateVars,proto3" json:"templateVars,omitempty"`
	DeploymentState *structpb.Struct `protobuf:"bytes,4,opt,name=deployment_state,json=deploymentState,proto3" json:"deployment_state,omitempty"`
	DeploymentVars  *structpb.Struct `protobuf:"bytes,5,opt,name=deployment_vars,json=deploymentVars,proto3" json:"deployment_vars,omitempty"`
}

func (x *DestroyRequest) Reset() {
	*x = DestroyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyRequest) ProtoMessage() {}

func (x *DestroyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyRequest.ProtoReflect.Descriptor instead.
func (*DestroyRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{2}
}

func (x *DestroyRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *DestroyRequest) GetBlueprint() []byte {
	if x != nil {
		return x.Blueprint
	}
	return nil
}

func (x *DestroyRequest) GetTemplateVars() *structpb.Struct {
	if x != nil {
		return x.TemplateVars
	}
	return nil
}

func (x *DestroyRequest) GetDeploymentState() *structpb.Struct {
	if x != nil {
		return x.DeploymentState
	}
	return nil
}

func (x *DestroyRequest) GetDeploymentVars() *structpb.Struct {
	if x != nil {
		return x.DeploymentVars
	}
	return nil
}

type DestroyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId    string           `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Status          common.RPCStatus `protobuf:"varint,2,opt,name=status,proto3,enum=RPCStatus" json:"status,omitempty"`
	DeploymentState *structpb.Struct `protobuf:"bytes,3,opt,name=deployment_state,json=deploymentState,proto3" json:"deployment_state,omitempty"`
	DeploymentVars  *structpb.Struct `protobuf:"bytes,4,opt,name=deployment_vars,json=deploymentVars,proto3" json:"deployment_vars,omitempty"`
}

func (x *DestroyReply) Reset() {
	*x = DestroyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyReply) ProtoMessage() {}

func (x *DestroyReply) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyReply.ProtoReflect.Descriptor instead.
func (*DestroyReply) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{3}
}

func (x *DestroyReply) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *DestroyReply) GetStatus() common.RPCStatus {
	if x != nil {
		return x.Status
	}
	return common.RPCStatus(0)
}

func (x *DestroyReply) GetDeploymentState() *structpb.Struct {
	if x != nil {
		return x.DeploymentState
	}
	return nil
}

func (x *DestroyReply) GetDeploymentVars() *structpb.Struct {
	if x != nil {
		return x.DeploymentVars
	}
	return nil
}

var File_provider_proto protoreflect.FileDescriptor

var file_provider_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a,
	0x0d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x73, 0x22, 0xdc,
	0x01, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x73, 0x22, 0x96, 0x02,
	0x0a, 0x0e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56,
	0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x73,
	0x12, 0x42, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x56, 0x61, 0x72, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x52,
	0x50, 0x43, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x42, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x56, 0x61, 0x72, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x09, 0x48, 0x61, 0x6e,
	0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x11, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61,
	0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x48, 0x61, 0x6e, 0x64,
	0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x06,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x0e, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x12, 0x0f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x62, 0x6c, 0x65, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x63, 0x62, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provider_proto_rawDescOnce sync.Once
	file_provider_proto_rawDescData = file_provider_proto_rawDesc
)

func file_provider_proto_rawDescGZIP() []byte {
	file_provider_proto_rawDescOnce.Do(func() {
		file_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_provider_proto_rawDescData)
	})
	return file_provider_proto_rawDescData
}

var file_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_provider_proto_goTypes = []interface{}{
	(*DeployRequest)(nil),           // 0: DeployRequest
	(*DeployReply)(nil),             // 1: DeployReply
	(*DestroyRequest)(nil),          // 2: DestroyRequest
	(*DestroyReply)(nil),            // 3: DestroyReply
	(*structpb.Struct)(nil),         // 4: google.protobuf.Struct
	(common.RPCStatus)(0),           // 5: RPCStatus
	(*common.HandshakeRequest)(nil), // 6: HandshakeRequest
	(*common.HandshakeReply)(nil),   // 7: HandshakeReply
}
var file_provider_proto_depIdxs = []int32{
	4,  // 0: DeployRequest.templateVars:type_name -> google.protobuf.Struct
	5,  // 1: DeployReply.status:type_name -> RPCStatus
	4,  // 2: DeployReply.deployment_state:type_name -> google.protobuf.Struct
	4,  // 3: DeployReply.deployment_vars:type_name -> google.protobuf.Struct
	4,  // 4: DestroyRequest.templateVars:type_name -> google.protobuf.Struct
	4,  // 5: DestroyRequest.deployment_state:type_name -> google.protobuf.Struct
	4,  // 6: DestroyRequest.deployment_vars:type_name -> google.protobuf.Struct
	5,  // 7: DestroyReply.status:type_name -> RPCStatus
	4,  // 8: DestroyReply.deployment_state:type_name -> google.protobuf.Struct
	4,  // 9: DestroyReply.deployment_vars:type_name -> google.protobuf.Struct
	6,  // 10: ProviderServer.Handshake:input_type -> HandshakeRequest
	0,  // 11: ProviderServer.Deploy:input_type -> DeployRequest
	2,  // 12: ProviderServer.Destroy:input_type -> DestroyRequest
	7,  // 13: ProviderServer.Handshake:output_type -> HandshakeReply
	1,  // 14: ProviderServer.Deploy:output_type -> DeployReply
	3,  // 15: ProviderServer.Destroy:output_type -> DestroyReply
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_provider_proto_init() }
func file_provider_proto_init() {
	if File_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployRequest); i {
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
		file_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployReply); i {
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
		file_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyRequest); i {
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
		file_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyReply); i {
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
			RawDescriptor: file_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provider_proto_goTypes,
		DependencyIndexes: file_provider_proto_depIdxs,
		MessageInfos:      file_provider_proto_msgTypes,
	}.Build()
	File_provider_proto = out.File
	file_provider_proto_rawDesc = nil
	file_provider_proto_goTypes = nil
	file_provider_proto_depIdxs = nil
}
