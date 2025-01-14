// proto/quanta_admin_service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.20.3
// source: quanta_admin_service.proto

package quanta_admin_service

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

// Message definitions
type CommonGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommonGetRequest) Reset() {
	*x = CommonGetRequest{}
	mi := &file_quanta_admin_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonGetRequest) ProtoMessage() {}

func (x *CommonGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quanta_admin_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonGetRequest.ProtoReflect.Descriptor instead.
func (*CommonGetRequest) Descriptor() ([]byte, []int) {
	return file_quanta_admin_service_proto_rawDescGZIP(), []int{0}
}

func (x *CommonGetRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type GetInstanceConfigResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Configs       []*InstanceConfig      `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs,omitempty"` //配置信息
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInstanceConfigResponse) Reset() {
	*x = GetInstanceConfigResponse{}
	mi := &file_quanta_admin_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInstanceConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstanceConfigResponse) ProtoMessage() {}

func (x *GetInstanceConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quanta_admin_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstanceConfigResponse.ProtoReflect.Descriptor instead.
func (*GetInstanceConfigResponse) Descriptor() ([]byte, []int) {
	return file_quanta_admin_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetInstanceConfigResponse) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *GetInstanceConfigResponse) GetConfigs() []*InstanceConfig {
	if x != nil {
		return x.Configs
	}
	return nil
}

type InstanceConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SchemaText    string                 `protobuf:"bytes,1,opt,name=schema_text,json=schemaText,proto3" json:"schema_text,omitempty"`
	SchemaType    string                 `protobuf:"bytes,2,opt,name=schema_type,json=schemaType,proto3" json:"schema_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InstanceConfig) Reset() {
	*x = InstanceConfig{}
	mi := &file_quanta_admin_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceConfig) ProtoMessage() {}

func (x *InstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_quanta_admin_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceConfig.ProtoReflect.Descriptor instead.
func (*InstanceConfig) Descriptor() ([]byte, []int) {
	return file_quanta_admin_service_proto_rawDescGZIP(), []int{2}
}

func (x *InstanceConfig) GetSchemaText() string {
	if x != nil {
		return x.SchemaText
	}
	return ""
}

func (x *InstanceConfig) GetSchemaType() string {
	if x != nil {
		return x.SchemaType
	}
	return ""
}

type GetStartOrStopStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // 0-stop，1-start
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStartOrStopStatusResponse) Reset() {
	*x = GetStartOrStopStatusResponse{}
	mi := &file_quanta_admin_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStartOrStopStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStartOrStopStatusResponse) ProtoMessage() {}

func (x *GetStartOrStopStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quanta_admin_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStartOrStopStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStartOrStopStatusResponse) Descriptor() ([]byte, []int) {
	return file_quanta_admin_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetStartOrStopStatusResponse) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *GetStartOrStopStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_quanta_admin_service_proto protoreflect.FileDescriptor

var file_quanta_admin_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x61, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x33, 0x0a, 0x10, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22,
	0x74, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x36, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x52, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x22, 0x57, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x32, 0xd5, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x61, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x64, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x53, 0x74, 0x6f, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x2e, 0x2f,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x61, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3b, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x61, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_quanta_admin_service_proto_rawDescOnce sync.Once
	file_quanta_admin_service_proto_rawDescData = file_quanta_admin_service_proto_rawDesc
)

func file_quanta_admin_service_proto_rawDescGZIP() []byte {
	file_quanta_admin_service_proto_rawDescOnce.Do(func() {
		file_quanta_admin_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_quanta_admin_service_proto_rawDescData)
	})
	return file_quanta_admin_service_proto_rawDescData
}

var file_quanta_admin_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_quanta_admin_service_proto_goTypes = []any{
	(*CommonGetRequest)(nil),             // 0: grpc_service.CommonGetRequest
	(*GetInstanceConfigResponse)(nil),    // 1: grpc_service.GetInstanceConfigResponse
	(*InstanceConfig)(nil),               // 2: grpc_service.InstanceConfig
	(*GetStartOrStopStatusResponse)(nil), // 3: grpc_service.GetStartOrStopStatusResponse
}
var file_quanta_admin_service_proto_depIdxs = []int32{
	2, // 0: grpc_service.GetInstanceConfigResponse.configs:type_name -> grpc_service.InstanceConfig
	0, // 1: grpc_service.QuantaAdmin.GetStrategyInstanceConfig:input_type -> grpc_service.CommonGetRequest
	0, // 2: grpc_service.QuantaAdmin.GetStartOrStopFlag:input_type -> grpc_service.CommonGetRequest
	1, // 3: grpc_service.QuantaAdmin.GetStrategyInstanceConfig:output_type -> grpc_service.GetInstanceConfigResponse
	3, // 4: grpc_service.QuantaAdmin.GetStartOrStopFlag:output_type -> grpc_service.GetStartOrStopStatusResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_quanta_admin_service_proto_init() }
func file_quanta_admin_service_proto_init() {
	if File_quanta_admin_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_quanta_admin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quanta_admin_service_proto_goTypes,
		DependencyIndexes: file_quanta_admin_service_proto_depIdxs,
		MessageInfos:      file_quanta_admin_service_proto_msgTypes,
	}.Build()
	File_quanta_admin_service_proto = out.File
	file_quanta_admin_service_proto_rawDesc = nil
	file_quanta_admin_service_proto_goTypes = nil
	file_quanta_admin_service_proto_depIdxs = nil
}
