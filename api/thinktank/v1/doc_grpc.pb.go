// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.7.1
// source: api/thinktank/v1/doc.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DocClient is the client API for Doc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocClient interface {
	CreateDoc(ctx context.Context, in *CreateDocRequest, opts ...grpc.CallOption) (*UpdateResp, error)
	UpdateDoc(ctx context.Context, in *UpdateDocRequest, opts ...grpc.CallOption) (*UpdateResp, error)
	DeleteDoc(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UpdateResp, error)
	GetDoc(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*GetDocReply, error)
	ListDoc(ctx context.Context, in *ListDocRequest, opts ...grpc.CallOption) (*ListDocReply, error)
}

type docClient struct {
	cc grpc.ClientConnInterface
}

func NewDocClient(cc grpc.ClientConnInterface) DocClient {
	return &docClient{cc}
}

func (c *docClient) CreateDoc(ctx context.Context, in *CreateDocRequest, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/api.thinktank.v1.Doc/CreateDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *docClient) UpdateDoc(ctx context.Context, in *UpdateDocRequest, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/api.thinktank.v1.Doc/UpdateDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *docClient) DeleteDoc(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/api.thinktank.v1.Doc/DeleteDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *docClient) GetDoc(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*GetDocReply, error) {
	out := new(GetDocReply)
	err := c.cc.Invoke(ctx, "/api.thinktank.v1.Doc/GetDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *docClient) ListDoc(ctx context.Context, in *ListDocRequest, opts ...grpc.CallOption) (*ListDocReply, error) {
	out := new(ListDocReply)
	err := c.cc.Invoke(ctx, "/api.thinktank.v1.Doc/ListDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocServer is the server API for Doc service.
// All implementations must embed UnimplementedDocServer
// for forward compatibility
type DocServer interface {
	CreateDoc(context.Context, *CreateDocRequest) (*UpdateResp, error)
	UpdateDoc(context.Context, *UpdateDocRequest) (*UpdateResp, error)
	DeleteDoc(context.Context, *IDReq) (*UpdateResp, error)
	GetDoc(context.Context, *IDReq) (*GetDocReply, error)
	ListDoc(context.Context, *ListDocRequest) (*ListDocReply, error)
	mustEmbedUnimplementedDocServer()
}

// UnimplementedDocServer must be embedded to have forward compatible implementations.
type UnimplementedDocServer struct {
}

func (UnimplementedDocServer) CreateDoc(context.Context, *CreateDocRequest) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDoc not implemented")
}
func (UnimplementedDocServer) UpdateDoc(context.Context, *UpdateDocRequest) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoc not implemented")
}
func (UnimplementedDocServer) DeleteDoc(context.Context, *IDReq) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoc not implemented")
}
func (UnimplementedDocServer) GetDoc(context.Context, *IDReq) (*GetDocReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoc not implemented")
}
func (UnimplementedDocServer) ListDoc(context.Context, *ListDocRequest) (*ListDocReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDoc not implemented")
}
func (UnimplementedDocServer) mustEmbedUnimplementedDocServer() {}

// UnsafeDocServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocServer will
// result in compilation errors.
type UnsafeDocServer interface {
	mustEmbedUnimplementedDocServer()
}

func RegisterDocServer(s grpc.ServiceRegistrar, srv DocServer) {
	s.RegisterService(&Doc_ServiceDesc, srv)
}

func _Doc_CreateDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocServer).CreateDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.thinktank.v1.Doc/CreateDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocServer).CreateDoc(ctx, req.(*CreateDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doc_UpdateDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocServer).UpdateDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.thinktank.v1.Doc/UpdateDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocServer).UpdateDoc(ctx, req.(*UpdateDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doc_DeleteDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocServer).DeleteDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.thinktank.v1.Doc/DeleteDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocServer).DeleteDoc(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doc_GetDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocServer).GetDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.thinktank.v1.Doc/GetDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocServer).GetDoc(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doc_ListDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocServer).ListDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.thinktank.v1.Doc/ListDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocServer).ListDoc(ctx, req.(*ListDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Doc_ServiceDesc is the grpc.ServiceDesc for Doc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Doc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.thinktank.v1.Doc",
	HandlerType: (*DocServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDoc",
			Handler:    _Doc_CreateDoc_Handler,
		},
		{
			MethodName: "UpdateDoc",
			Handler:    _Doc_UpdateDoc_Handler,
		},
		{
			MethodName: "DeleteDoc",
			Handler:    _Doc_DeleteDoc_Handler,
		},
		{
			MethodName: "GetDoc",
			Handler:    _Doc_GetDoc_Handler,
		},
		{
			MethodName: "ListDoc",
			Handler:    _Doc_ListDoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/thinktank/v1/doc.proto",
}
