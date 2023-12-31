// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user/v1/user.proto

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

const (
	User_ListUser_FullMethodName               = "/bitMail.v1.User/ListUser"
	User_Ping_FullMethodName                   = "/bitMail.v1.User/Ping"
	User_Login_FullMethodName                  = "/bitMail.v1.User/Login"
	User_Register_FullMethodName               = "/bitMail.v1.User/Register"
	User_GetUserDetailsByID_FullMethodName     = "/bitMail.v1.User/GetUserDetailsByID"
	User_GetContactListByUserId_FullMethodName = "/bitMail.v1.User/GetContactListByUserId"
	User_UploadContact_FullMethodName          = "/bitMail.v1.User/UploadContact"
	User_DelContactByUserID_FullMethodName     = "/bitMail.v1.User/DelContactByUserID"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// Sends a greeting
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	GetUserDetailsByID(ctx context.Context, in *GetUserDetailsByIDRequest, opts ...grpc.CallOption) (*GetUserDetailsByIDReply, error)
	GetContactListByUserId(ctx context.Context, in *GetContactListByUserIdRequest, opts ...grpc.CallOption) (*GetContactListByUserIdReply, error)
	UploadContact(ctx context.Context, in *UploadContactRequest, opts ...grpc.CallOption) (*UploadContactReply, error)
	DelContactByUserID(ctx context.Context, in *DelContactByUserIDRequest, opts ...grpc.CallOption) (*DelContactByUserIDReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, User_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, User_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, User_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserDetailsByID(ctx context.Context, in *GetUserDetailsByIDRequest, opts ...grpc.CallOption) (*GetUserDetailsByIDReply, error) {
	out := new(GetUserDetailsByIDReply)
	err := c.cc.Invoke(ctx, User_GetUserDetailsByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetContactListByUserId(ctx context.Context, in *GetContactListByUserIdRequest, opts ...grpc.CallOption) (*GetContactListByUserIdReply, error) {
	out := new(GetContactListByUserIdReply)
	err := c.cc.Invoke(ctx, User_GetContactListByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UploadContact(ctx context.Context, in *UploadContactRequest, opts ...grpc.CallOption) (*UploadContactReply, error) {
	out := new(UploadContactReply)
	err := c.cc.Invoke(ctx, User_UploadContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DelContactByUserID(ctx context.Context, in *DelContactByUserIDRequest, opts ...grpc.CallOption) (*DelContactByUserIDReply, error) {
	out := new(DelContactByUserIDReply)
	err := c.cc.Invoke(ctx, User_DelContactByUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// Sends a greeting
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	GetUserDetailsByID(context.Context, *GetUserDetailsByIDRequest) (*GetUserDetailsByIDReply, error)
	GetContactListByUserId(context.Context, *GetContactListByUserIdRequest) (*GetContactListByUserIdReply, error)
	UploadContact(context.Context, *UploadContactRequest) (*UploadContactReply, error)
	DelContactByUserID(context.Context, *DelContactByUserIDRequest) (*DelContactByUserIDReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) ListUser(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) GetUserDetailsByID(context.Context, *GetUserDetailsByIDRequest) (*GetUserDetailsByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetailsByID not implemented")
}
func (UnimplementedUserServer) GetContactListByUserId(context.Context, *GetContactListByUserIdRequest) (*GetContactListByUserIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactListByUserId not implemented")
}
func (UnimplementedUserServer) UploadContact(context.Context, *UploadContactRequest) (*UploadContactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadContact not implemented")
}
func (UnimplementedUserServer) DelContactByUserID(context.Context, *DelContactByUserIDRequest) (*DelContactByUserIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelContactByUserID not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserDetailsByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserDetailsByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserDetailsByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserDetailsByID(ctx, req.(*GetUserDetailsByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetContactListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactListByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetContactListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetContactListByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetContactListByUserId(ctx, req.(*GetContactListByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UploadContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UploadContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UploadContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UploadContact(ctx, req.(*UploadContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DelContactByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelContactByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DelContactByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DelContactByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DelContactByUserID(ctx, req.(*DelContactByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bitMail.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _User_ListUser_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _User_Ping_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "GetUserDetailsByID",
			Handler:    _User_GetUserDetailsByID_Handler,
		},
		{
			MethodName: "GetContactListByUserId",
			Handler:    _User_GetContactListByUserId_Handler,
		},
		{
			MethodName: "UploadContact",
			Handler:    _User_UploadContact_Handler,
		},
		{
			MethodName: "DelContactByUserID",
			Handler:    _User_DelContactByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user.proto",
}
