// proto/trigger_service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.20.3
// source: trigger_service.proto

package trigger_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
type StartTriggerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"` // 中台实例id
	OpenPrice     string                 `protobuf:"bytes,2,opt,name=open_price,json=openPrice,proto3" json:"open_price,omitempty"`    // 开仓价格
	ClosePrice    string                 `protobuf:"bytes,3,opt,name=close_price,json=closePrice,proto3" json:"close_price,omitempty"` // 平仓价格
	Side          string                 `protobuf:"bytes,4,opt,name=side,proto3" json:"side,omitempty"`                               // 开仓方向
	Amount        string                 `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`                           //开仓数量
	Symbol        string                 `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`                           // 交易对
	StopTime      string                 `protobuf:"bytes,7,opt,name=stop_time,json=stopTime,proto3" json:"stop_time,omitempty"`       //停止时间
	ApiConfig     *APIConfig             `protobuf:"bytes,8,opt,name=api_config,json=apiConfig,proto3" json:"api_config,omitempty"`    //Apikey相关
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartTriggerRequest) Reset() {
	*x = StartTriggerRequest{}
	mi := &file_trigger_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartTriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTriggerRequest) ProtoMessage() {}

func (x *StartTriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trigger_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTriggerRequest.ProtoReflect.Descriptor instead.
func (*StartTriggerRequest) Descriptor() ([]byte, []int) {
	return file_trigger_service_proto_rawDescGZIP(), []int{0}
}

func (x *StartTriggerRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *StartTriggerRequest) GetOpenPrice() string {
	if x != nil {
		return x.OpenPrice
	}
	return ""
}

func (x *StartTriggerRequest) GetClosePrice() string {
	if x != nil {
		return x.ClosePrice
	}
	return ""
}

func (x *StartTriggerRequest) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *StartTriggerRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *StartTriggerRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *StartTriggerRequest) GetStopTime() string {
	if x != nil {
		return x.StopTime
	}
	return ""
}

func (x *StartTriggerRequest) GetApiConfig() *APIConfig {
	if x != nil {
		return x.ApiConfig
	}
	return nil
}

type CheckApiKeyHealthyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiKey        string                 `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	SecretKey     string                 `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Exchange      string                 `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckApiKeyHealthyRequest) Reset() {
	*x = CheckApiKeyHealthyRequest{}
	mi := &file_trigger_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckApiKeyHealthyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckApiKeyHealthyRequest) ProtoMessage() {}

func (x *CheckApiKeyHealthyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trigger_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckApiKeyHealthyRequest.ProtoReflect.Descriptor instead.
func (*CheckApiKeyHealthyRequest) Descriptor() ([]byte, []int) {
	return file_trigger_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckApiKeyHealthyRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *CheckApiKeyHealthyRequest) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *CheckApiKeyHealthyRequest) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type APIConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiKey        string                 `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	SecretKey     string                 `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Exchange      string                 `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *APIConfig) Reset() {
	*x = APIConfig{}
	mi := &file_trigger_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APIConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIConfig) ProtoMessage() {}

func (x *APIConfig) ProtoReflect() protoreflect.Message {
	mi := &file_trigger_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIConfig.ProtoReflect.Descriptor instead.
func (*APIConfig) Descriptor() ([]byte, []int) {
	return file_trigger_service_proto_rawDescGZIP(), []int{2}
}

func (x *APIConfig) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *APIConfig) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *APIConfig) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type CheckApiKeyHealthyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsHealth      bool                   `protobuf:"varint,1,opt,name=is_health,json=isHealth,proto3" json:"is_health,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckApiKeyHealthyResponse) Reset() {
	*x = CheckApiKeyHealthyResponse{}
	mi := &file_trigger_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckApiKeyHealthyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckApiKeyHealthyResponse) ProtoMessage() {}

func (x *CheckApiKeyHealthyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trigger_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckApiKeyHealthyResponse.ProtoReflect.Descriptor instead.
func (*CheckApiKeyHealthyResponse) Descriptor() ([]byte, []int) {
	return file_trigger_service_proto_rawDescGZIP(), []int{3}
}

func (x *CheckApiKeyHealthyResponse) GetIsHealth() bool {
	if x != nil {
		return x.IsHealth
	}
	return false
}

type StartTriggerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartTriggerResponse) Reset() {
	*x = StartTriggerResponse{}
	mi := &file_trigger_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartTriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTriggerResponse) ProtoMessage() {}

func (x *StartTriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trigger_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTriggerResponse.ProtoReflect.Descriptor instead.
func (*StartTriggerResponse) Descriptor() ([]byte, []int) {
	return file_trigger_service_proto_rawDescGZIP(), []int{4}
}

func (x *StartTriggerResponse) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type StopTriggerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopTriggerRequest) Reset() {
	*x = StopTriggerRequest{}
	mi := &file_trigger_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopTriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTriggerRequest) ProtoMessage() {}

func (x *StopTriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trigger_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTriggerRequest.ProtoReflect.Descriptor instead.
func (*StopTriggerRequest) Descriptor() ([]byte, []int) {
	return file_trigger_service_proto_rawDescGZIP(), []int{5}
}

func (x *StopTriggerRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type TriggerListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceIds   []string               `protobuf:"bytes,1,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TriggerListResponse) Reset() {
	*x = TriggerListResponse{}
	mi := &file_trigger_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TriggerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerListResponse) ProtoMessage() {}

func (x *TriggerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trigger_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerListResponse.ProtoReflect.Descriptor instead.
func (*TriggerListResponse) Descriptor() ([]byte, []int) {
	return file_trigger_service_proto_rawDescGZIP(), []int{6}
}

func (x *TriggerListResponse) GetInstanceIds() []string {
	if x != nil {
		return x.InstanceIds
	}
	return nil
}

var File_trigger_service_proto protoreflect.FileDescriptor

var file_trigger_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x70, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0a,
	0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x50, 0x49, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x61, 0x70, 0x69, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x6f, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x5f, 0x0a, 0x09, 0x41, 0x50, 0x49, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x22, 0x37, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x12, 0x53, 0x74,
	0x6f, 0x70, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x38, 0x0a, 0x13, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x73, 0x32, 0x87, 0x02, 0x0a, 0x0f,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x56, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x41, 0x50, 0x49, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x28, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x2e, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_trigger_service_proto_rawDescOnce sync.Once
	file_trigger_service_proto_rawDescData = file_trigger_service_proto_rawDesc
)

func file_trigger_service_proto_rawDescGZIP() []byte {
	file_trigger_service_proto_rawDescOnce.Do(func() {
		file_trigger_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_trigger_service_proto_rawDescData)
	})
	return file_trigger_service_proto_rawDescData
}

var file_trigger_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_trigger_service_proto_goTypes = []any{
	(*StartTriggerRequest)(nil),        // 0: grpc_service.StartTriggerRequest
	(*CheckApiKeyHealthyRequest)(nil),  // 1: grpc_service.CheckApiKeyHealthyRequest
	(*APIConfig)(nil),                  // 2: grpc_service.APIConfig
	(*CheckApiKeyHealthyResponse)(nil), // 3: grpc_service.CheckApiKeyHealthyResponse
	(*StartTriggerResponse)(nil),       // 4: grpc_service.StartTriggerResponse
	(*StopTriggerRequest)(nil),         // 5: grpc_service.StopTriggerRequest
	(*TriggerListResponse)(nil),        // 6: grpc_service.TriggerListResponse
	(*emptypb.Empty)(nil),              // 7: google.protobuf.Empty
}
var file_trigger_service_proto_depIdxs = []int32{
	2, // 0: grpc_service.StartTriggerRequest.api_config:type_name -> grpc_service.APIConfig
	0, // 1: grpc_service.TriggerInstance.StartInstance:input_type -> grpc_service.StartTriggerRequest
	7, // 2: grpc_service.TriggerInstance.ListInstances:input_type -> google.protobuf.Empty
	2, // 3: grpc_service.TriggerInstance.CheckApiKey:input_type -> grpc_service.APIConfig
	4, // 4: grpc_service.TriggerInstance.StartInstance:output_type -> grpc_service.StartTriggerResponse
	6, // 5: grpc_service.TriggerInstance.ListInstances:output_type -> grpc_service.TriggerListResponse
	3, // 6: grpc_service.TriggerInstance.CheckApiKey:output_type -> grpc_service.CheckApiKeyHealthyResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trigger_service_proto_init() }
func file_trigger_service_proto_init() {
	if File_trigger_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trigger_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trigger_service_proto_goTypes,
		DependencyIndexes: file_trigger_service_proto_depIdxs,
		MessageInfos:      file_trigger_service_proto_msgTypes,
	}.Build()
	File_trigger_service_proto = out.File
	file_trigger_service_proto_rawDesc = nil
	file_trigger_service_proto_goTypes = nil
	file_trigger_service_proto_depIdxs = nil
}
