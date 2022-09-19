// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.7.1
// source: api/thinktank/v1/common.proto

package v1

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Mine int32

const (
	Mine_MINE_DEFAULT Mine = 0
	Mine_MINE_MINE    Mine = 1
)

// Enum value maps for Mine.
var (
	Mine_name = map[int32]string{
		0: "MINE_DEFAULT",
		1: "MINE_MINE",
	}
	Mine_value = map[string]int32{
		"MINE_DEFAULT": 0,
		"MINE_MINE":    1,
	}
)

func (x Mine) Enum() *Mine {
	p := new(Mine)
	*p = x
	return p
}

func (x Mine) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mine) Descriptor() protoreflect.EnumDescriptor {
	return file_api_thinktank_v1_common_proto_enumTypes[0].Descriptor()
}

func (Mine) Type() protoreflect.EnumType {
	return &file_api_thinktank_v1_common_proto_enumTypes[0]
}

func (x Mine) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mine.Descriptor instead.
func (Mine) EnumDescriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{0}
}

type IDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDReq) Reset() {
	*x = IDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReq) ProtoMessage() {}

func (x *IDReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReq.ProtoReflect.Descriptor instead.
func (*IDReq) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *IDReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IDReqs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IDReqs) Reset() {
	*x = IDReqs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReqs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReqs) ProtoMessage() {}

func (x *IDReqs) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReqs.ProtoReflect.Descriptor instead.
func (*IDReqs) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *IDReqs) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BaseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 模糊搜索
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// 分页
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 分页
	PageNum int32 `protobuf:"varint,3,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	// id
	Id int64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BaseReq) Reset() {
	*x = BaseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseReq) ProtoMessage() {}

func (x *BaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseReq.ProtoReflect.Descriptor instead.
func (*BaseReq) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *BaseReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *BaseReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BaseReq) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *BaseReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type KeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyReq) Reset() {
	*x = KeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyReq) ProtoMessage() {}

func (x *KeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyReq.ProtoReflect.Descriptor instead.
func (*KeyReq) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *KeyReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type UpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回信息
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDResp) Reset() {
	*x = IDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDResp) ProtoMessage() {}

func (x *IDResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDResp.ProtoReflect.Descriptor instead.
func (*IDResp) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{5}
}

func (x *IDResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MapResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapResp) Reset() {
	*x = MapResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapResp) ProtoMessage() {}

func (x *MapResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapResp.ProtoReflect.Descriptor instead.
func (*MapResp) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{6}
}

func (x *MapResp) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 模糊搜索
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// 分页
	PageSize int64 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 分页
	PageNum int64 `protobuf:"varint,4,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	// 开始时间
	Stime int64 `protobuf:"varint,5,opt,name=stime,proto3" json:"stime,omitempty"`
	// 结束时间
	Etime int64 `protobuf:"varint,6,opt,name=etime,proto3" json:"etime,omitempty"`
	// 排序字段
	Orderby string `protobuf:"bytes,7,opt,name=orderby,proto3" json:"orderby,omitempty"`
	// 排序字段
	Asc string `protobuf:"bytes,8,opt,name=asc,proto3" json:"asc,omitempty"`
	// 我的
	Mine int64 `protobuf:"varint,9,opt,name=mine,proto3" json:"mine,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_common_proto_rawDescGZIP(), []int{7}
}

func (x *ListReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListReq) GetStime() int64 {
	if x != nil {
		return x.Stime
	}
	return 0
}

func (x *ListReq) GetEtime() int64 {
	if x != nil {
		return x.Etime
	}
	return 0
}

func (x *ListReq) GetOrderby() string {
	if x != nil {
		return x.Orderby
	}
	return ""
}

func (x *ListReq) GetAsc() string {
	if x != nil {
		return x.Asc
	}
	return ""
}

func (x *ListReq) GetMine() int64 {
	if x != nil {
		return x.Mine
	}
	return 0
}

var File_api_thinktank_v1_common_proto protoreflect.FileDescriptor

var file_api_thinktank_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x1a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x05, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x27, 0xea, 0xde, 0x1f, 0x02, 0x69, 0x64, 0xf2, 0xde, 0x1f, 0x1d, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x22, 0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a,
	0x06, 0x49, 0x44, 0x52, 0x65, 0x71, 0x73, 0x12, 0x3b, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x42, 0x29, 0xea, 0xde, 0x1f, 0x03, 0x69, 0x64, 0x73, 0xf2, 0xde, 0x1f,
	0x1e, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x69, 0x64, 0x73, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x07, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x27, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xea,
	0xde, 0x1f, 0x03, 0x6b, 0x65, 0x79, 0xf2, 0xde, 0x1f, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22,
	0x6b, 0x65, 0x79, 0x22, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x21, 0xea, 0xde, 0x1f,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0xf2, 0xde, 0x1f, 0x10, 0x66, 0x6f,
	0x72, 0x6d, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1f, 0xea, 0xde, 0x1f, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0xf2, 0xde, 0x1f, 0x0f, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x22, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x22, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x13, 0xea, 0xde, 0x1f, 0x02, 0x69, 0x64, 0xf2, 0xde, 0x1f, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x06, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0xf2, 0xde, 0x1f, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6b, 0x65, 0x79, 0x22, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x33, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x25, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xea, 0xde, 0x1f, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x06, 0xea, 0xde, 0x1f, 0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x07,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e,
	0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x07, 0xea, 0xde, 0x1f, 0x03,
	0x6d, 0x61, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x8c, 0x04, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x23,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x13, 0xea, 0xde, 0x1f, 0x02,
	0x69, 0x64, 0xf2, 0xde, 0x1f, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0xea, 0xde, 0x1f, 0x01, 0x73, 0xf2, 0xde, 0x1f, 0x08, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x22, 0x73, 0x22, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x21, 0xea, 0xde, 0x1f, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0xf2, 0xde, 0x1f, 0x10, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f, 0xea, 0xde, 0x1f, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0xf2, 0xde, 0x1f, 0x0f, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x22, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x58, 0x0a, 0x05, 0x73, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x42, 0xea, 0xde, 0x1f, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0xf2, 0xde, 0x1f, 0x11, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0xfa, 0xde, 0x1f, 0x1b, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x05, 0x73, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x05,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x3e, 0xea, 0xde, 0x1f,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0xf2, 0xde, 0x1f, 0x0f, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x22, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0xfa, 0xde, 0x1f, 0x1b,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x05, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1d, 0xea, 0xde, 0x1f, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x79,
	0xf2, 0xde, 0x1f, 0x0e, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x79, 0x22, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x79, 0x12, 0x27, 0x0a, 0x03, 0x61,
	0x73, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xea, 0xde, 0x1f, 0x03, 0x61, 0x73,
	0x63, 0xf2, 0xde, 0x1f, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x61, 0x73, 0x63, 0x22, 0x52,
	0x03, 0x61, 0x73, 0x63, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x17, 0xea, 0xde, 0x1f, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0xf2, 0xde, 0x1f, 0x0b,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6d, 0x69, 0x6e, 0x65, 0x22, 0x52, 0x04, 0x6d, 0x69, 0x6e,
	0x65, 0x2a, 0x27, 0x0a, 0x04, 0x4d, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x49, 0x4e,
	0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d,
	0x49, 0x4e, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x42, 0x23, 0x5a, 0x1d, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x68, 0x69, 0x6e,
	0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xc8, 0xe1, 0x1e, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_thinktank_v1_common_proto_rawDescOnce sync.Once
	file_api_thinktank_v1_common_proto_rawDescData = file_api_thinktank_v1_common_proto_rawDesc
)

func file_api_thinktank_v1_common_proto_rawDescGZIP() []byte {
	file_api_thinktank_v1_common_proto_rawDescOnce.Do(func() {
		file_api_thinktank_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_thinktank_v1_common_proto_rawDescData)
	})
	return file_api_thinktank_v1_common_proto_rawDescData
}

var file_api_thinktank_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_thinktank_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_thinktank_v1_common_proto_goTypes = []interface{}{
	(Mine)(0),          // 0: api.thinktank.v1.Mine
	(*IDReq)(nil),      // 1: api.thinktank.v1.IDReq
	(*IDReqs)(nil),     // 2: api.thinktank.v1.IDReqs
	(*BaseReq)(nil),    // 3: api.thinktank.v1.BaseReq
	(*KeyReq)(nil),     // 4: api.thinktank.v1.KeyReq
	(*UpdateResp)(nil), // 5: api.thinktank.v1.UpdateResp
	(*IDResp)(nil),     // 6: api.thinktank.v1.IDResp
	(*MapResp)(nil),    // 7: api.thinktank.v1.MapResp
	(*ListReq)(nil),    // 8: api.thinktank.v1.ListReq
	nil,                // 9: api.thinktank.v1.MapResp.DataEntry
}
var file_api_thinktank_v1_common_proto_depIdxs = []int32{
	9, // 0: api.thinktank.v1.MapResp.data:type_name -> api.thinktank.v1.MapResp.DataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_thinktank_v1_common_proto_init() }
func file_api_thinktank_v1_common_proto_init() {
	if File_api_thinktank_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_thinktank_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReq); i {
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
		file_api_thinktank_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReqs); i {
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
		file_api_thinktank_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseReq); i {
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
		file_api_thinktank_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyReq); i {
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
		file_api_thinktank_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResp); i {
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
		file_api_thinktank_v1_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDResp); i {
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
		file_api_thinktank_v1_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapResp); i {
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
		file_api_thinktank_v1_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
			RawDescriptor: file_api_thinktank_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_thinktank_v1_common_proto_goTypes,
		DependencyIndexes: file_api_thinktank_v1_common_proto_depIdxs,
		EnumInfos:         file_api_thinktank_v1_common_proto_enumTypes,
		MessageInfos:      file_api_thinktank_v1_common_proto_msgTypes,
	}.Build()
	File_api_thinktank_v1_common_proto = out.File
	file_api_thinktank_v1_common_proto_rawDesc = nil
	file_api_thinktank_v1_common_proto_goTypes = nil
	file_api_thinktank_v1_common_proto_depIdxs = nil
}
