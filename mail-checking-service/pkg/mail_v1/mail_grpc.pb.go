// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: mail.proto

package mail_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MailV1Client is the client API for MailV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailV1Client interface {
	CheckUniqueCode(ctx context.Context, in *CheckUniqueCodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mailV1Client struct {
	cc grpc.ClientConnInterface
}

func NewMailV1Client(cc grpc.ClientConnInterface) MailV1Client {
	return &mailV1Client{cc}
}

func (c *mailV1Client) CheckUniqueCode(ctx context.Context, in *CheckUniqueCodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mail_v1.MailV1/CheckUniqueCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailV1Server is the server API for MailV1 service.
// All implementations must embed UnimplementedMailV1Server
// for forward compatibility
type MailV1Server interface {
	CheckUniqueCode(context.Context, *CheckUniqueCodeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMailV1Server()
}

// UnimplementedMailV1Server must be embedded to have forward compatible implementations.
type UnimplementedMailV1Server struct {
}

func (UnimplementedMailV1Server) CheckUniqueCode(context.Context, *CheckUniqueCodeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUniqueCode not implemented")
}
func (UnimplementedMailV1Server) mustEmbedUnimplementedMailV1Server() {}

// UnsafeMailV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailV1Server will
// result in compilation errors.
type UnsafeMailV1Server interface {
	mustEmbedUnimplementedMailV1Server()
}

func RegisterMailV1Server(s grpc.ServiceRegistrar, srv MailV1Server) {
	s.RegisterService(&MailV1_ServiceDesc, srv)
}

func _MailV1_CheckUniqueCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUniqueCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailV1Server).CheckUniqueCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mail_v1.MailV1/CheckUniqueCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailV1Server).CheckUniqueCode(ctx, req.(*CheckUniqueCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailV1_ServiceDesc is the grpc.ServiceDesc for MailV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mail_v1.MailV1",
	HandlerType: (*MailV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUniqueCode",
			Handler:    _MailV1_CheckUniqueCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mail.proto",
}