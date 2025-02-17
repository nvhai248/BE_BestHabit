// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/userservice/user_service.proto

package userproto

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
	UserService_UserUpdateProfile_FullMethodName     = "/bestHabit.UserService/UserUpdateProfile"
	UserService_UserUpload_FullMethodName            = "/bestHabit.UserService/UserUpload"
	UserService_UserUpdatePassword_FullMethodName    = "/bestHabit.UserService/UserUpdatePassword"
	UserService_UserUpdateDeviceToken_FullMethodName = "/bestHabit.UserService/UserUpdateDeviceToken"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserUpdateProfile(ctx context.Context, in *UserUpdateProfileRequest, opts ...grpc.CallOption) (*UserUpdateProfileResponse, error)
	UserUpload(ctx context.Context, in *UserUploadRequest, opts ...grpc.CallOption) (*UserUploadResponse, error)
	UserUpdatePassword(ctx context.Context, in *UserUpdatePasswordRequest, opts ...grpc.CallOption) (*UserUpdatePasswordResponse, error)
	UserUpdateDeviceToken(ctx context.Context, in *UserUpdateDeviceTokenRequest, opts ...grpc.CallOption) (*UserUpdateDeviceTokenResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserUpdateProfile(ctx context.Context, in *UserUpdateProfileRequest, opts ...grpc.CallOption) (*UserUpdateProfileResponse, error) {
	out := new(UserUpdateProfileResponse)
	err := c.cc.Invoke(ctx, UserService_UserUpdateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserUpload(ctx context.Context, in *UserUploadRequest, opts ...grpc.CallOption) (*UserUploadResponse, error) {
	out := new(UserUploadResponse)
	err := c.cc.Invoke(ctx, UserService_UserUpload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserUpdatePassword(ctx context.Context, in *UserUpdatePasswordRequest, opts ...grpc.CallOption) (*UserUpdatePasswordResponse, error) {
	out := new(UserUpdatePasswordResponse)
	err := c.cc.Invoke(ctx, UserService_UserUpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserUpdateDeviceToken(ctx context.Context, in *UserUpdateDeviceTokenRequest, opts ...grpc.CallOption) (*UserUpdateDeviceTokenResponse, error) {
	out := new(UserUpdateDeviceTokenResponse)
	err := c.cc.Invoke(ctx, UserService_UserUpdateDeviceToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserUpdateProfile(context.Context, *UserUpdateProfileRequest) (*UserUpdateProfileResponse, error)
	UserUpload(context.Context, *UserUploadRequest) (*UserUploadResponse, error)
	UserUpdatePassword(context.Context, *UserUpdatePasswordRequest) (*UserUpdatePasswordResponse, error)
	UserUpdateDeviceToken(context.Context, *UserUpdateDeviceTokenRequest) (*UserUpdateDeviceTokenResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserUpdateProfile(context.Context, *UserUpdateProfileRequest) (*UserUpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdateProfile not implemented")
}
func (UnimplementedUserServiceServer) UserUpload(context.Context, *UserUploadRequest) (*UserUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpload not implemented")
}
func (UnimplementedUserServiceServer) UserUpdatePassword(context.Context, *UserUpdatePasswordRequest) (*UserUpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdatePassword not implemented")
}
func (UnimplementedUserServiceServer) UserUpdateDeviceToken(context.Context, *UserUpdateDeviceTokenRequest) (*UserUpdateDeviceTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdateDeviceToken not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserUpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserUpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUpdateProfile(ctx, req.(*UserUpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUpload(ctx, req.(*UserUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserUpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserUpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUpdatePassword(ctx, req.(*UserUpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserUpdateDeviceToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateDeviceTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUpdateDeviceToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserUpdateDeviceToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUpdateDeviceToken(ctx, req.(*UserUpdateDeviceTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bestHabit.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserUpdateProfile",
			Handler:    _UserService_UserUpdateProfile_Handler,
		},
		{
			MethodName: "UserUpload",
			Handler:    _UserService_UserUpload_Handler,
		},
		{
			MethodName: "UserUpdatePassword",
			Handler:    _UserService_UserUpdatePassword_Handler,
		},
		{
			MethodName: "UserUpdateDeviceToken",
			Handler:    _UserService_UserUpdateDeviceToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/userservice/user_service.proto",
}
