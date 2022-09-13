// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: protos/Query/query.proto

package query

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

type TokenName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
}

func (x *TokenName) Reset() {
	*x = TokenName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_Query_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenName) ProtoMessage() {}

func (x *TokenName) ProtoReflect() protoreflect.Message {
	mi := &file_protos_Query_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenName.ProtoReflect.Descriptor instead.
func (*TokenName) Descriptor() ([]byte, []int) {
	return file_protos_Query_query_proto_rawDescGZIP(), []int{0}
}

func (x *TokenName) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName *TokenName `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	Account   string     `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_Query_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_Query_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_protos_Query_query_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceRequest) GetTokenName() *TokenName {
	if x != nil {
		return x.TokenName
	}
	return nil
}

func (x *GetBalanceRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance uint64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_Query_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_Query_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_protos_Query_query_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceResponse) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type TokenInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName *TokenName `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
}

func (x *TokenInfoRequest) Reset() {
	*x = TokenInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_Query_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfoRequest) ProtoMessage() {}

func (x *TokenInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_Query_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfoRequest.ProtoReflect.Descriptor instead.
func (*TokenInfoRequest) Descriptor() ([]byte, []int) {
	return file_protos_Query_query_proto_rawDescGZIP(), []int{3}
}

func (x *TokenInfoRequest) GetTokenName() *TokenName {
	if x != nil {
		return x.TokenName
	}
	return nil
}

type TokenInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName   string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	Symbol      string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimal     string `protobuf:"bytes,3,opt,name=decimal,proto3" json:"decimal,omitempty"`
	TotalSupply string `protobuf:"bytes,4,opt,name=totalSupply,proto3" json:"totalSupply,omitempty"`
}

func (x *TokenInfoResponse) Reset() {
	*x = TokenInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_Query_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfoResponse) ProtoMessage() {}

func (x *TokenInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_Query_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfoResponse.ProtoReflect.Descriptor instead.
func (*TokenInfoResponse) Descriptor() ([]byte, []int) {
	return file_protos_Query_query_proto_rawDescGZIP(), []int{4}
}

func (x *TokenInfoResponse) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *TokenInfoResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenInfoResponse) GetDecimal() string {
	if x != nil {
		return x.Decimal
	}
	return ""
}

func (x *TokenInfoResponse) GetTotalSupply() string {
	if x != nil {
		return x.TotalSupply
	}
	return ""
}

type AllowanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName *TokenName `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	Owner     string     `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Spender   string     `protobuf:"bytes,3,opt,name=spender,proto3" json:"spender,omitempty"`
}

func (x *AllowanceRequest) Reset() {
	*x = AllowanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_Query_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowanceRequest) ProtoMessage() {}

func (x *AllowanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_Query_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowanceRequest.ProtoReflect.Descriptor instead.
func (*AllowanceRequest) Descriptor() ([]byte, []int) {
	return file_protos_Query_query_proto_rawDescGZIP(), []int{5}
}

func (x *AllowanceRequest) GetTokenName() *TokenName {
	if x != nil {
		return x.TokenName
	}
	return nil
}

func (x *AllowanceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AllowanceRequest) GetSpender() string {
	if x != nil {
		return x.Spender
	}
	return ""
}

type AllowanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance uint64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *AllowanceResponse) Reset() {
	*x = AllowanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_Query_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowanceResponse) ProtoMessage() {}

func (x *AllowanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_Query_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowanceResponse.ProtoReflect.Descriptor instead.
func (*AllowanceResponse) Descriptor() ([]byte, []int) {
	return file_protos_Query_query_proto_rawDescGZIP(), []int{6}
}

func (x *AllowanceResponse) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_protos_Query_query_proto protoreflect.FileDescriptor

var file_protos_Query_query_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x72, 0x63, 0x32,
	0x30, 0x22, 0x29, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x85, 0x01, 0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x22, 0x72, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x11, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xd0, 0x01, 0x0a, 0x05, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65,
	0x72, 0x63, 0x32, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x65, 0x72, 0x63,
	0x32, 0x30, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6e, 0x67,
	0x67, 0x79, 0x75, 0x4c, 0x69, 0x6d, 0x2f, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_Query_query_proto_rawDescOnce sync.Once
	file_protos_Query_query_proto_rawDescData = file_protos_Query_query_proto_rawDesc
)

func file_protos_Query_query_proto_rawDescGZIP() []byte {
	file_protos_Query_query_proto_rawDescOnce.Do(func() {
		file_protos_Query_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_Query_query_proto_rawDescData)
	})
	return file_protos_Query_query_proto_rawDescData
}

var file_protos_Query_query_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_Query_query_proto_goTypes = []interface{}{
	(*TokenName)(nil),          // 0: erc20.tokenName
	(*GetBalanceRequest)(nil),  // 1: erc20.GetBalanceRequest
	(*GetBalanceResponse)(nil), // 2: erc20.GetBalanceResponse
	(*TokenInfoRequest)(nil),   // 3: erc20.TokenInfoRequest
	(*TokenInfoResponse)(nil),  // 4: erc20.TokenInfoResponse
	(*AllowanceRequest)(nil),   // 5: erc20.AllowanceRequest
	(*AllowanceResponse)(nil),  // 6: erc20.AllowanceResponse
}
var file_protos_Query_query_proto_depIdxs = []int32{
	0, // 0: erc20.GetBalanceRequest.tokenName:type_name -> erc20.tokenName
	0, // 1: erc20.TokenInfoRequest.tokenName:type_name -> erc20.tokenName
	0, // 2: erc20.AllowanceRequest.tokenName:type_name -> erc20.tokenName
	1, // 3: erc20.Query.GetBalance:input_type -> erc20.GetBalanceRequest
	3, // 4: erc20.Query.GetTokenInfo:input_type -> erc20.TokenInfoRequest
	5, // 5: erc20.Query.GetAllowance:input_type -> erc20.AllowanceRequest
	2, // 6: erc20.Query.GetBalance:output_type -> erc20.GetBalanceResponse
	4, // 7: erc20.Query.GetTokenInfo:output_type -> erc20.TokenInfoResponse
	6, // 8: erc20.Query.GetAllowance:output_type -> erc20.AllowanceResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_Query_query_proto_init() }
func file_protos_Query_query_proto_init() {
	if File_protos_Query_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_Query_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenName); i {
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
		file_protos_Query_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_protos_Query_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
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
		file_protos_Query_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfoRequest); i {
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
		file_protos_Query_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfoResponse); i {
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
		file_protos_Query_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowanceRequest); i {
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
		file_protos_Query_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowanceResponse); i {
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
			RawDescriptor: file_protos_Query_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_Query_query_proto_goTypes,
		DependencyIndexes: file_protos_Query_query_proto_depIdxs,
		MessageInfos:      file_protos_Query_query_proto_msgTypes,
	}.Build()
	File_protos_Query_query_proto = out.File
	file_protos_Query_query_proto_rawDesc = nil
	file_protos_Query_query_proto_goTypes = nil
	file_protos_Query_query_proto_depIdxs = nil
}
