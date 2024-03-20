// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: service_booknest.proto

package pb

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

const (
	Booknest_CreateUser_FullMethodName  = "/pb.Booknest/CreateUser"
	Booknest_LoginUser_FullMethodName   = "/pb.Booknest/LoginUser"
	Booknest_VerifyEmail_FullMethodName = "/pb.Booknest/VerifyEmail"
)

// BooknestClient is the client API for Booknest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooknestClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
}

type booknestClient struct {
	cc grpc.ClientConnInterface
}

func NewBooknestClient(cc grpc.ClientConnInterface) BooknestClient {
	return &booknestClient{cc}
}

func (c *booknestClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Booknest_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booknestClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Booknest_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booknestClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, Booknest_VerifyEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooknestServer is the server API for Booknest service.
// All implementations must embed UnimplementedBooknestServer
// for forward compatibility
type BooknestServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	mustEmbedUnimplementedBooknestServer()
}

// UnimplementedBooknestServer must be embedded to have forward compatible implementations.
type UnimplementedBooknestServer struct {
}

func (UnimplementedBooknestServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedBooknestServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedBooknestServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedBooknestServer) mustEmbedUnimplementedBooknestServer() {}

// UnsafeBooknestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooknestServer will
// result in compilation errors.
type UnsafeBooknestServer interface {
	mustEmbedUnimplementedBooknestServer()
}

func RegisterBooknestServer(s grpc.ServiceRegistrar, srv BooknestServer) {
	s.RegisterService(&Booknest_ServiceDesc, srv)
}

func _Booknest_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooknestServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booknest_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooknestServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Booknest_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooknestServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booknest_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooknestServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Booknest_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooknestServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booknest_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooknestServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Booknest_ServiceDesc is the grpc.ServiceDesc for Booknest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Booknest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Booknest",
	HandlerType: (*BooknestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Booknest_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Booknest_LoginUser_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _Booknest_VerifyEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_booknest.proto",
}