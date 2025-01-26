// proto/observe_service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.20.3
// source: observer_service.proto

package observe_service

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

type DexType int32

const (
	DexType_RAY_AMM  DexType = 0
	DexType_RAY_CLMM DexType = 1
)

// Enum value maps for DexType.
var (
	DexType_name = map[int32]string{
		0: "RAY_AMM",
		1: "RAY_CLMM",
	}
	DexType_value = map[string]int32{
		"RAY_AMM":  0,
		"RAY_CLMM": 1,
	}
)

func (x DexType) Enum() *DexType {
	p := new(DexType)
	*p = x
	return p
}

func (x DexType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DexType) Descriptor() protoreflect.EnumDescriptor {
	return file_observer_service_proto_enumTypes[0].Descriptor()
}

func (DexType) Type() protoreflect.EnumType {
	return &file_observer_service_proto_enumTypes[0]
}

func (x DexType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DexType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DexType(num)
	return nil
}

// Deprecated: Use DexType.Descriptor instead.
func (DexType) EnumDescriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{0}
}

// Message definitions
type StartRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	AmberConfig     *AmberConfig           `protobuf:"bytes,1,req,name=amber_config,json=amberConfig" json:"amber_config,omitempty"`
	DexConfig       *DexConfig             `protobuf:"bytes,2,req,name=dex_config,json=dexConfig" json:"dex_config,omitempty"`
	ArbitrageConfig *ArbitrageConfig       `protobuf:"bytes,3,req,name=arbitrage_config,json=arbitrageConfig" json:"arbitrage_config,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	mi := &file_observer_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[0]
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
	return file_observer_service_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetAmberConfig() *AmberConfig {
	if x != nil {
		return x.AmberConfig
	}
	return nil
}

func (x *StartRequest) GetDexConfig() *DexConfig {
	if x != nil {
		return x.DexConfig
	}
	return nil
}

func (x *StartRequest) GetArbitrageConfig() *ArbitrageConfig {
	if x != nil {
		return x.ArbitrageConfig
	}
	return nil
}

type StartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    *string                `protobuf:"bytes,1,req,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	mi := &file_observer_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[1]
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
	return file_observer_service_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetInstanceId() string {
	if x != nil && x.InstanceId != nil {
		return *x.InstanceId
	}
	return ""
}

type StopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    *string                `protobuf:"bytes,1,req,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	mi := &file_observer_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[2]
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
	return file_observer_service_proto_rawDescGZIP(), []int{2}
}

func (x *StopRequest) GetInstanceId() string {
	if x != nil && x.InstanceId != nil {
		return *x.InstanceId
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    *string                `protobuf:"bytes,1,req,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_observer_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetInstanceId() string {
	if x != nil && x.InstanceId != nil {
		return *x.InstanceId
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Observers     []*ObserverInfo        `protobuf:"bytes,1,rep,name=observers" json:"observers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_observer_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[4]
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
	return file_observer_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetObservers() []*ObserverInfo {
	if x != nil {
		return x.Observers
	}
	return nil
}

type GetStateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    *string                `protobuf:"bytes,1,req,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStateRequest) Reset() {
	*x = GetStateRequest{}
	mi := &file_observer_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateRequest) ProtoMessage() {}

func (x *GetStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateRequest.ProtoReflect.Descriptor instead.
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetStateRequest) GetInstanceId() string {
	if x != nil && x.InstanceId != nil {
		return *x.InstanceId
	}
	return ""
}

type GetStateResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Timestamp         *int64                 `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	ProfitOfBuyOnDex  *float64               `protobuf:"fixed64,2,req,name=profit_of_buy_on_dex,json=profitOfBuyOnDex" json:"profit_of_buy_on_dex,omitempty"`
	ProfitOfSellOnDex *float64               `protobuf:"fixed64,3,req,name=profit_of_sell_on_dex,json=profitOfSellOnDex" json:"profit_of_sell_on_dex,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetStateResponse) Reset() {
	*x = GetStateResponse{}
	mi := &file_observer_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateResponse) ProtoMessage() {}

func (x *GetStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateResponse.ProtoReflect.Descriptor instead.
func (*GetStateResponse) Descriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetStateResponse) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *GetStateResponse) GetProfitOfBuyOnDex() float64 {
	if x != nil && x.ProfitOfBuyOnDex != nil {
		return *x.ProfitOfBuyOnDex
	}
	return 0
}

func (x *GetStateResponse) GetProfitOfSellOnDex() float64 {
	if x != nil && x.ProfitOfSellOnDex != nil {
		return *x.ProfitOfSellOnDex
	}
	return 0
}

type AmberConfig struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	ExchangeType        *string                `protobuf:"bytes,1,req,name=exchange_type,json=exchangeType" json:"exchange_type,omitempty"` // Binance
	BaseTokenOrderbook  *AmberOrderBookConfig  `protobuf:"bytes,2,req,name=base_token_orderbook,json=baseTokenOrderbook" json:"base_token_orderbook,omitempty"`
	QuoteTokenOrderbook *AmberOrderBookConfig  `protobuf:"bytes,3,req,name=quote_token_orderbook,json=quoteTokenOrderbook" json:"quote_token_orderbook,omitempty"`
	TakerFee            *float64               `protobuf:"fixed64,4,req,name=taker_fee,json=takerFee" json:"taker_fee,omitempty"` // 0.00023
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AmberConfig) Reset() {
	*x = AmberConfig{}
	mi := &file_observer_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AmberConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmberConfig) ProtoMessage() {}

func (x *AmberConfig) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmberConfig.ProtoReflect.Descriptor instead.
func (*AmberConfig) Descriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{7}
}

func (x *AmberConfig) GetExchangeType() string {
	if x != nil && x.ExchangeType != nil {
		return *x.ExchangeType
	}
	return ""
}

func (x *AmberConfig) GetBaseTokenOrderbook() *AmberOrderBookConfig {
	if x != nil {
		return x.BaseTokenOrderbook
	}
	return nil
}

func (x *AmberConfig) GetQuoteTokenOrderbook() *AmberOrderBookConfig {
	if x != nil {
		return x.QuoteTokenOrderbook
	}
	return nil
}

func (x *AmberConfig) GetTakerFee() float64 {
	if x != nil && x.TakerFee != nil {
		return *x.TakerFee
	}
	return 0
}

type DexConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DexType       *DexType               `protobuf:"varint,1,req,name=dex_type,json=dexType,enum=grpc_service.DexType" json:"dex_type,omitempty"`
	AmmPool       *string                `protobuf:"bytes,2,req,name=amm_pool,json=ammPool" json:"amm_pool,omitempty"`
	TokenMint     *string                `protobuf:"bytes,3,req,name=token_mint,json=tokenMint" json:"token_mint,omitempty"`
	SlippageBps   *uint64                `protobuf:"varint,4,req,name=slippage_bps,json=slippageBps" json:"slippage_bps,omitempty"`
	MaxArraySize  *uint32                `protobuf:"varint,5,opt,name=max_array_size,json=maxArraySize" json:"max_array_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DexConfig) Reset() {
	*x = DexConfig{}
	mi := &file_observer_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DexConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DexConfig) ProtoMessage() {}

func (x *DexConfig) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DexConfig.ProtoReflect.Descriptor instead.
func (*DexConfig) Descriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{8}
}

func (x *DexConfig) GetDexType() DexType {
	if x != nil && x.DexType != nil {
		return *x.DexType
	}
	return DexType_RAY_AMM
}

func (x *DexConfig) GetAmmPool() string {
	if x != nil && x.AmmPool != nil {
		return *x.AmmPool
	}
	return ""
}

func (x *DexConfig) GetTokenMint() string {
	if x != nil && x.TokenMint != nil {
		return *x.TokenMint
	}
	return ""
}

func (x *DexConfig) GetSlippageBps() uint64 {
	if x != nil && x.SlippageBps != nil {
		return *x.SlippageBps
	}
	return 0
}

func (x *DexConfig) GetMaxArraySize() uint32 {
	if x != nil && x.MaxArraySize != nil {
		return *x.MaxArraySize
	}
	return 0
}

type ArbitrageConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	QuoteAmount   *float64               `protobuf:"fixed64,1,req,name=quote_amount,json=quoteAmount" json:"quote_amount,omitempty"` // 1.0
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ArbitrageConfig) Reset() {
	*x = ArbitrageConfig{}
	mi := &file_observer_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArbitrageConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArbitrageConfig) ProtoMessage() {}

func (x *ArbitrageConfig) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArbitrageConfig.ProtoReflect.Descriptor instead.
func (*ArbitrageConfig) Descriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{9}
}

func (x *ArbitrageConfig) GetQuoteAmount() float64 {
	if x != nil && x.QuoteAmount != nil {
		return *x.QuoteAmount
	}
	return 0
}

type AmberOrderBookConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Symbol        *string                `protobuf:"bytes,1,req,name=symbol" json:"symbol,omitempty"` // PNUT/USDT
	AskDepth      *int32                 `protobuf:"varint,2,opt,name=ask_depth,json=askDepth,def=20" json:"ask_depth,omitempty"`
	BidDepth      *int32                 `protobuf:"varint,3,opt,name=bid_depth,json=bidDepth,def=20" json:"bid_depth,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// Default values for AmberOrderBookConfig fields.
const (
	Default_AmberOrderBookConfig_AskDepth = int32(20)
	Default_AmberOrderBookConfig_BidDepth = int32(20)
)

func (x *AmberOrderBookConfig) Reset() {
	*x = AmberOrderBookConfig{}
	mi := &file_observer_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AmberOrderBookConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmberOrderBookConfig) ProtoMessage() {}

func (x *AmberOrderBookConfig) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmberOrderBookConfig.ProtoReflect.Descriptor instead.
func (*AmberOrderBookConfig) Descriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{10}
}

func (x *AmberOrderBookConfig) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *AmberOrderBookConfig) GetAskDepth() int32 {
	if x != nil && x.AskDepth != nil {
		return *x.AskDepth
	}
	return Default_AmberOrderBookConfig_AskDepth
}

func (x *AmberOrderBookConfig) GetBidDepth() int32 {
	if x != nil && x.BidDepth != nil {
		return *x.BidDepth
	}
	return Default_AmberOrderBookConfig_BidDepth
}

type ObserverInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    *string                `protobuf:"bytes,1,req,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	DexType       *DexType               `protobuf:"varint,2,req,name=dex_type,json=dexType,enum=grpc_service.DexType" json:"dex_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObserverInfo) Reset() {
	*x = ObserverInfo{}
	mi := &file_observer_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObserverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObserverInfo) ProtoMessage() {}

func (x *ObserverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_observer_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObserverInfo.ProtoReflect.Descriptor instead.
func (*ObserverInfo) Descriptor() ([]byte, []int) {
	return file_observer_service_proto_rawDescGZIP(), []int{11}
}

func (x *ObserverInfo) GetInstanceId() string {
	if x != nil && x.InstanceId != nil {
		return *x.InstanceId
	}
	return ""
}

func (x *ObserverInfo) GetDexType() DexType {
	if x != nil && x.DexType != nil {
		return *x.DexType
	}
	return DexType_RAY_AMM
}

var File_observer_service_proto protoreflect.FileDescriptor

var file_observer_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x61, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6d, 0x62, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x61, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x36, 0x0a, 0x0a, 0x64, 0x65, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x09, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48, 0x0a, 0x10, 0x61, 0x72,
	0x62, 0x69, 0x74, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0f, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x30, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x22, 0x32, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x01, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x4f, 0x66, 0x42,
	0x75, 0x79, 0x4f, 0x6e, 0x44, 0x65, 0x78, 0x12, 0x30, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x01, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x4f, 0x66,
	0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x6e, 0x44, 0x65, 0x78, 0x22, 0xfd, 0x01, 0x0a, 0x0b, 0x41, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x54,
	0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x12, 0x62, 0x61, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x56, 0x0a, 0x15, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x13, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x01, 0x52,
	0x08, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x46, 0x65, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x09, 0x44, 0x65,
	0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x08, 0x64, 0x65, 0x78, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6d, 0x6d,
	0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6d, 0x6d,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4d,
	0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x62, 0x70, 0x73, 0x18, 0x04, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x6c, 0x69, 0x70, 0x70,
	0x61, 0x67, 0x65, 0x42, 0x70, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x6d, 0x61, 0x78, 0x41, 0x72, 0x72, 0x61, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x34, 0x0a, 0x0f,
	0x41, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x21, 0x0a, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x01, 0x52, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x70, 0x0a, 0x14, 0x41, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x32, 0x30, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x44, 0x65,
	0x70, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x32, 0x30, 0x52, 0x08, 0x62, 0x69, 0x64, 0x44,
	0x65, 0x70, 0x74, 0x68, 0x22, 0x61, 0x0a, 0x0c, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x64, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07,
	0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x24, 0x0a, 0x07, 0x44, 0x65, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x41, 0x59, 0x5f, 0x41, 0x4d, 0x4d, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x41, 0x59, 0x5f, 0x43, 0x4c, 0x4d, 0x4d, 0x10, 0x01, 0x32, 0xaf, 0x02,
	0x0a, 0x08, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x23, 0x5a, 0x21, 0x2e, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3b, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65,
}

var (
	file_observer_service_proto_rawDescOnce sync.Once
	file_observer_service_proto_rawDescData = file_observer_service_proto_rawDesc
)

func file_observer_service_proto_rawDescGZIP() []byte {
	file_observer_service_proto_rawDescOnce.Do(func() {
		file_observer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_observer_service_proto_rawDescData)
	})
	return file_observer_service_proto_rawDescData
}

var file_observer_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_observer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_observer_service_proto_goTypes = []any{
	(DexType)(0),                 // 0: grpc_service.DexType
	(*StartRequest)(nil),         // 1: grpc_service.StartRequest
	(*StartResponse)(nil),        // 2: grpc_service.StartResponse
	(*StopRequest)(nil),          // 3: grpc_service.StopRequest
	(*ListRequest)(nil),          // 4: grpc_service.ListRequest
	(*ListResponse)(nil),         // 5: grpc_service.ListResponse
	(*GetStateRequest)(nil),      // 6: grpc_service.GetStateRequest
	(*GetStateResponse)(nil),     // 7: grpc_service.GetStateResponse
	(*AmberConfig)(nil),          // 8: grpc_service.AmberConfig
	(*DexConfig)(nil),            // 9: grpc_service.DexConfig
	(*ArbitrageConfig)(nil),      // 10: grpc_service.ArbitrageConfig
	(*AmberOrderBookConfig)(nil), // 11: grpc_service.AmberOrderBookConfig
	(*ObserverInfo)(nil),         // 12: grpc_service.ObserverInfo
	(*emptypb.Empty)(nil),        // 13: google.protobuf.Empty
}
var file_observer_service_proto_depIdxs = []int32{
	8,  // 0: grpc_service.StartRequest.amber_config:type_name -> grpc_service.AmberConfig
	9,  // 1: grpc_service.StartRequest.dex_config:type_name -> grpc_service.DexConfig
	10, // 2: grpc_service.StartRequest.arbitrage_config:type_name -> grpc_service.ArbitrageConfig
	12, // 3: grpc_service.ListResponse.observers:type_name -> grpc_service.ObserverInfo
	11, // 4: grpc_service.AmberConfig.base_token_orderbook:type_name -> grpc_service.AmberOrderBookConfig
	11, // 5: grpc_service.AmberConfig.quote_token_orderbook:type_name -> grpc_service.AmberOrderBookConfig
	0,  // 6: grpc_service.DexConfig.dex_type:type_name -> grpc_service.DexType
	0,  // 7: grpc_service.ObserverInfo.dex_type:type_name -> grpc_service.DexType
	1,  // 8: grpc_service.Observer.StartObserver:input_type -> grpc_service.StartRequest
	3,  // 9: grpc_service.Observer.StopObserver:input_type -> grpc_service.StopRequest
	13, // 10: grpc_service.Observer.ListObservers:input_type -> google.protobuf.Empty
	6,  // 11: grpc_service.Observer.GetObserverState:input_type -> grpc_service.GetStateRequest
	2,  // 12: grpc_service.Observer.StartObserver:output_type -> grpc_service.StartResponse
	13, // 13: grpc_service.Observer.StopObserver:output_type -> google.protobuf.Empty
	5,  // 14: grpc_service.Observer.ListObservers:output_type -> grpc_service.ListResponse
	7,  // 15: grpc_service.Observer.GetObserverState:output_type -> grpc_service.GetStateResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_observer_service_proto_init() }
func file_observer_service_proto_init() {
	if File_observer_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_observer_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_observer_service_proto_goTypes,
		DependencyIndexes: file_observer_service_proto_depIdxs,
		EnumInfos:         file_observer_service_proto_enumTypes,
		MessageInfos:      file_observer_service_proto_msgTypes,
	}.Build()
	File_observer_service_proto = out.File
	file_observer_service_proto_rawDesc = nil
	file_observer_service_proto_goTypes = nil
	file_observer_service_proto_depIdxs = nil
}
