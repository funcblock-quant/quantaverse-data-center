// proto/instance_service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.20.3
// source: instance_service.proto

package instance_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type InstanceType int32

const (
	InstanceType_OBSERVER_INSTANCE InstanceType = 0 //观察者实例
	InstanceType_TRADER_INSTANCE   InstanceType = 1 //交易者实例
)

// Enum value maps for InstanceType.
var (
	InstanceType_name = map[int32]string{
		0: "OBSERVER_INSTANCE",
		1: "TRADER_INSTANCE",
	}
	InstanceType_value = map[string]int32{
		"OBSERVER_INSTANCE": 0,
		"TRADER_INSTANCE":   1,
	}
)

func (x InstanceType) Enum() *InstanceType {
	p := new(InstanceType)
	*p = x
	return p
}

func (x InstanceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceType) Descriptor() protoreflect.EnumDescriptor {
	return file_instance_service_proto_enumTypes[0].Descriptor()
}

func (InstanceType) Type() protoreflect.EnumType {
	return &file_instance_service_proto_enumTypes[0]
}

func (x InstanceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceType.Descriptor instead.
func (InstanceType) EnumDescriptor() ([]byte, []int) {
	return file_instance_service_proto_rawDescGZIP(), []int{0}
}

// Message definitions
type StartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`                                       // 中台实例id
	InstanceType  InstanceType           `protobuf:"varint,2,opt,name=instance_type,json=instanceType,proto3,enum=grpc_service.InstanceType" json:"instance_type,omitempty"` // 实例类型，
	ConfigYaml    *structpb.Struct       `protobuf:"bytes,3,opt,name=config_yaml,json=configYaml,proto3" json:"config_yaml,omitempty"`                                       // 策略启动实例的参数封装
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	mi := &file_instance_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instance_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_instance_service_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *StartRequest) GetInstanceType() InstanceType {
	if x != nil {
		return x.InstanceType
	}
	return InstanceType_OBSERVER_INSTANCE
}

func (x *StartRequest) GetConfigYaml() *structpb.Struct {
	if x != nil {
		return x.ConfigYaml
	}
	return nil
}

type StartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	mi := &file_instance_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instance_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_instance_service_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type StopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	mi := &file_instance_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instance_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_instance_service_proto_rawDescGZIP(), []int{2}
}

func (x *StopRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceIds   string                 `protobuf:"bytes,1,opt,name=instance_ids,json=instanceIds,proto3" json:"instance_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_instance_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instance_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_instance_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetInstanceIds() string {
	if x != nil {
		return x.InstanceIds
	}
	return ""
}

type GetRealtimeInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRealtimeInfoRequest) Reset() {
	*x = GetRealtimeInfoRequest{}
	mi := &file_instance_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRealtimeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealtimeInfoRequest) ProtoMessage() {}

func (x *GetRealtimeInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instance_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealtimeInfoRequest.ProtoReflect.Descriptor instead.
func (*GetRealtimeInfoRequest) Descriptor() ([]byte, []int) {
	return file_instance_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetRealtimeInfoRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type GetRealtimeInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     int64                  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data          *structpb.Struct       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // 通用数据字段
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRealtimeInfoResponse) Reset() {
	*x = GetRealtimeInfoResponse{}
	mi := &file_instance_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRealtimeInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealtimeInfoResponse) ProtoMessage() {}

func (x *GetRealtimeInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instance_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealtimeInfoResponse.ProtoReflect.Descriptor instead.
func (*GetRealtimeInfoResponse) Descriptor() ([]byte, []int) {
	return file_instance_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetRealtimeInfoResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GetRealtimeInfoResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_instance_service_proto protoreflect.FileDescriptor

var file_instance_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x79,
	0x61, 0x6d, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d, 0x6c, 0x22, 0x30,
	0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x2e, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x31, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x73, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x64,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x2a, 0x3a, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x42, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54,
	0x52, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x01,
	0x32, 0xc4, 0x02, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x66, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61,
	0x6c, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x2e, 0x2f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_instance_service_proto_rawDescOnce sync.Once
	file_instance_service_proto_rawDescData = file_instance_service_proto_rawDesc
)

func file_instance_service_proto_rawDescGZIP() []byte {
	file_instance_service_proto_rawDescOnce.Do(func() {
		file_instance_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_instance_service_proto_rawDescData)
	})
	return file_instance_service_proto_rawDescData
}

var file_instance_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_instance_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_instance_service_proto_goTypes = []any{
	(InstanceType)(0),               // 0: grpc_service.InstanceType
	(*StartRequest)(nil),            // 1: grpc_service.StartRequest
	(*StartResponse)(nil),           // 2: grpc_service.StartResponse
	(*StopRequest)(nil),             // 3: grpc_service.StopRequest
	(*ListResponse)(nil),            // 4: grpc_service.ListResponse
	(*GetRealtimeInfoRequest)(nil),  // 5: grpc_service.GetRealtimeInfoRequest
	(*GetRealtimeInfoResponse)(nil), // 6: grpc_service.GetRealtimeInfoResponse
	(*structpb.Struct)(nil),         // 7: google.protobuf.Struct
	(*emptypb.Empty)(nil),           // 8: google.protobuf.Empty
}
var file_instance_service_proto_depIdxs = []int32{
	0, // 0: grpc_service.StartRequest.instance_type:type_name -> grpc_service.InstanceType
	7, // 1: grpc_service.StartRequest.config_yaml:type_name -> google.protobuf.Struct
	7, // 2: grpc_service.GetRealtimeInfoResponse.data:type_name -> google.protobuf.Struct
	1, // 3: grpc_service.Instance.StartInstance:input_type -> grpc_service.StartRequest
	3, // 4: grpc_service.Instance.StopInstance:input_type -> grpc_service.StopRequest
	8, // 5: grpc_service.Instance.ListInstances:input_type -> google.protobuf.Empty
	5, // 6: grpc_service.Instance.GetInstanceRealtimeInfo:input_type -> grpc_service.GetRealtimeInfoRequest
	2, // 7: grpc_service.Instance.StartInstance:output_type -> grpc_service.StartResponse
	8, // 8: grpc_service.Instance.StopInstance:output_type -> google.protobuf.Empty
	4, // 9: grpc_service.Instance.ListInstances:output_type -> grpc_service.ListResponse
	6, // 10: grpc_service.Instance.GetInstanceRealtimeInfo:output_type -> grpc_service.GetRealtimeInfoResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_instance_service_proto_init() }
func file_instance_service_proto_init() {
	if File_instance_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_instance_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_instance_service_proto_goTypes,
		DependencyIndexes: file_instance_service_proto_depIdxs,
		EnumInfos:         file_instance_service_proto_enumTypes,
		MessageInfos:      file_instance_service_proto_msgTypes,
	}.Build()
	File_instance_service_proto = out.File
	file_instance_service_proto_rawDesc = nil
	file_instance_service_proto_goTypes = nil
	file_instance_service_proto_depIdxs = nil
}
