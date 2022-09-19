// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.7.1
// source: api/thinktank/v1/doc.proto

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

type CreateDocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// title
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateDocRequest) Reset() {
	*x = CreateDocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_doc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDocRequest) ProtoMessage() {}

func (x *CreateDocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_doc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDocRequest.ProtoReflect.Descriptor instead.
func (*CreateDocRequest) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_doc_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDocRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateDocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// title
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateDocRequest) Reset() {
	*x = UpdateDocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_doc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDocRequest) ProtoMessage() {}

func (x *UpdateDocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_doc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDocRequest.ProtoReflect.Descriptor instead.
func (*UpdateDocRequest) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_doc_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDocRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDocRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetDocReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// uuid
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// 标题
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// 创建时间
	Ctime int64 `protobuf:"varint,4,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// 修改时间
	Mtime int64 `protobuf:"varint,5,opt,name=mtime,proto3" json:"mtime,omitempty"`
}

func (x *GetDocReply) Reset() {
	*x = GetDocReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_doc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDocReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocReply) ProtoMessage() {}

func (x *GetDocReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_doc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocReply.ProtoReflect.Descriptor instead.
func (*GetDocReply) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_doc_proto_rawDescGZIP(), []int{2}
}

func (x *GetDocReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetDocReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetDocReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetDocReply) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *GetDocReply) GetMtime() int64 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

type ListDocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 模糊搜索
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// title
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// 分页
	PageSize int64 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 分页
	PageNum int64 `protobuf:"varint,5,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	// 排序字段
	Orderby string `protobuf:"bytes,6,opt,name=orderby,proto3" json:"orderby,omitempty"`
	// 排序顺序
	Asc string `protobuf:"bytes,7,opt,name=asc,proto3" json:"asc,omitempty"`
	// 我的 0.全部 1.我的
	Mine Mine `protobuf:"varint,8,opt,name=mine,proto3,enum=api.thinktank.v1.Mine" json:"mine,omitempty"`
}

func (x *ListDocRequest) Reset() {
	*x = ListDocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_doc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDocRequest) ProtoMessage() {}

func (x *ListDocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_doc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDocRequest.ProtoReflect.Descriptor instead.
func (*ListDocRequest) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_doc_proto_rawDescGZIP(), []int{3}
}

func (x *ListDocRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListDocRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListDocRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListDocRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDocRequest) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListDocRequest) GetOrderby() string {
	if x != nil {
		return x.Orderby
	}
	return ""
}

func (x *ListDocRequest) GetAsc() string {
	if x != nil {
		return x.Asc
	}
	return ""
}

func (x *ListDocRequest) GetMine() Mine {
	if x != nil {
		return x.Mine
	}
	return Mine_MINE_DEFAULT
}

type ListDocReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*DocForList `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Count int64         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListDocReply) Reset() {
	*x = ListDocReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_doc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDocReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDocReply) ProtoMessage() {}

func (x *ListDocReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_doc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDocReply.ProtoReflect.Descriptor instead.
func (*ListDocReply) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_doc_proto_rawDescGZIP(), []int{4}
}

func (x *ListDocReply) GetData() []*DocForList {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListDocReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DocForList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// uuid
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// 标题
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// 创建时间
	Ctime int64 `protobuf:"varint,4,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// 修改时间
	Mtime int64 `protobuf:"varint,5,opt,name=mtime,proto3" json:"mtime,omitempty"`
}

func (x *DocForList) Reset() {
	*x = DocForList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thinktank_v1_doc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocForList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocForList) ProtoMessage() {}

func (x *DocForList) ProtoReflect() protoreflect.Message {
	mi := &file_api_thinktank_v1_doc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocForList.ProtoReflect.Descriptor instead.
func (*DocForList) Descriptor() ([]byte, []int) {
	return file_api_thinktank_v1_doc_proto_rawDescGZIP(), []int{5}
}

func (x *DocForList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DocForList) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DocForList) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DocForList) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *DocForList) GetMtime() int64 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

var File_api_thinktank_v1_doc_proto protoreflect.FileDescriptor

var file_api_thinktank_v1_doc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x68,
	0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x38, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x73, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0xd8, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x73, 0x63, 0x12,
	0x2a, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x65, 0x22, 0x56, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63,
	0x46, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x72, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x46, 0x6f, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xe3, 0x03, 0x0a, 0x03, 0x44, 0x6f, 0x63, 0x12,
	0x65, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x12, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x64, 0x6f, 0x63, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x63, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74,
	0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68,
	0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x1a, 0x0b, 0x2f,
	0x64, 0x6f, 0x63, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74,
	0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0b, 0x2f, 0x64, 0x6f, 0x63, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x63, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74,
	0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x64, 0x6f, 0x63, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x5e, 0x0a,
	0x07, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0b, 0x12, 0x09, 0x2f, 0x64, 0x6f, 0x63, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x23, 0x5a,
	0x1d, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x68, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xc8, 0xe1,
	0x1e, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_thinktank_v1_doc_proto_rawDescOnce sync.Once
	file_api_thinktank_v1_doc_proto_rawDescData = file_api_thinktank_v1_doc_proto_rawDesc
)

func file_api_thinktank_v1_doc_proto_rawDescGZIP() []byte {
	file_api_thinktank_v1_doc_proto_rawDescOnce.Do(func() {
		file_api_thinktank_v1_doc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_thinktank_v1_doc_proto_rawDescData)
	})
	return file_api_thinktank_v1_doc_proto_rawDescData
}

var file_api_thinktank_v1_doc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_thinktank_v1_doc_proto_goTypes = []interface{}{
	(*CreateDocRequest)(nil), // 0: api.thinktank.v1.CreateDocRequest
	(*UpdateDocRequest)(nil), // 1: api.thinktank.v1.UpdateDocRequest
	(*GetDocReply)(nil),      // 2: api.thinktank.v1.GetDocReply
	(*ListDocRequest)(nil),   // 3: api.thinktank.v1.ListDocRequest
	(*ListDocReply)(nil),     // 4: api.thinktank.v1.ListDocReply
	(*DocForList)(nil),       // 5: api.thinktank.v1.DocForList
	(Mine)(0),                // 6: api.thinktank.v1.Mine
	(*IDReq)(nil),            // 7: api.thinktank.v1.IDReq
	(*UpdateResp)(nil),       // 8: api.thinktank.v1.UpdateResp
}
var file_api_thinktank_v1_doc_proto_depIdxs = []int32{
	6, // 0: api.thinktank.v1.ListDocRequest.mine:type_name -> api.thinktank.v1.Mine
	5, // 1: api.thinktank.v1.ListDocReply.data:type_name -> api.thinktank.v1.DocForList
	0, // 2: api.thinktank.v1.Doc.CreateDoc:input_type -> api.thinktank.v1.CreateDocRequest
	1, // 3: api.thinktank.v1.Doc.UpdateDoc:input_type -> api.thinktank.v1.UpdateDocRequest
	7, // 4: api.thinktank.v1.Doc.DeleteDoc:input_type -> api.thinktank.v1.IDReq
	7, // 5: api.thinktank.v1.Doc.GetDoc:input_type -> api.thinktank.v1.IDReq
	3, // 6: api.thinktank.v1.Doc.ListDoc:input_type -> api.thinktank.v1.ListDocRequest
	8, // 7: api.thinktank.v1.Doc.CreateDoc:output_type -> api.thinktank.v1.UpdateResp
	8, // 8: api.thinktank.v1.Doc.UpdateDoc:output_type -> api.thinktank.v1.UpdateResp
	8, // 9: api.thinktank.v1.Doc.DeleteDoc:output_type -> api.thinktank.v1.UpdateResp
	2, // 10: api.thinktank.v1.Doc.GetDoc:output_type -> api.thinktank.v1.GetDocReply
	4, // 11: api.thinktank.v1.Doc.ListDoc:output_type -> api.thinktank.v1.ListDocReply
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_thinktank_v1_doc_proto_init() }
func file_api_thinktank_v1_doc_proto_init() {
	if File_api_thinktank_v1_doc_proto != nil {
		return
	}
	file_api_thinktank_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_thinktank_v1_doc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDocRequest); i {
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
		file_api_thinktank_v1_doc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDocRequest); i {
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
		file_api_thinktank_v1_doc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDocReply); i {
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
		file_api_thinktank_v1_doc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDocRequest); i {
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
		file_api_thinktank_v1_doc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDocReply); i {
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
		file_api_thinktank_v1_doc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocForList); i {
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
			RawDescriptor: file_api_thinktank_v1_doc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_thinktank_v1_doc_proto_goTypes,
		DependencyIndexes: file_api_thinktank_v1_doc_proto_depIdxs,
		MessageInfos:      file_api_thinktank_v1_doc_proto_msgTypes,
	}.Build()
	File_api_thinktank_v1_doc_proto = out.File
	file_api_thinktank_v1_doc_proto_rawDesc = nil
	file_api_thinktank_v1_doc_proto_goTypes = nil
	file_api_thinktank_v1_doc_proto_depIdxs = nil
}