// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.23.4
// source: user/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserDelContactByUserID = "/bitMail.v1.User/DelContactByUserID"
const OperationUserGetContactListByUserId = "/bitMail.v1.User/GetContactListByUserId"
const OperationUserGetUserDetailsByID = "/bitMail.v1.User/GetUserDetailsByID"
const OperationUserListUser = "/bitMail.v1.User/ListUser"
const OperationUserLogin = "/bitMail.v1.User/Login"
const OperationUserPing = "/bitMail.v1.User/Ping"
const OperationUserRegister = "/bitMail.v1.User/Register"
const OperationUserUploadContact = "/bitMail.v1.User/UploadContact"

type UserHTTPServer interface {
	DelContactByUserID(context.Context, *DelContactByUserIDRequest) (*DelContactByUserIDReply, error)
	GetContactListByUserId(context.Context, *GetContactListByUserIdRequest) (*GetContactListByUserIdReply, error)
	GetUserDetailsByID(context.Context, *GetUserDetailsByIDRequest) (*GetUserDetailsByIDReply, error)
	// ListUser Sends a greeting
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	UploadContact(context.Context, *UploadContactRequest) (*UploadContactReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.GET("/bit-mail/user/listUser", _User_ListUser0_HTTP_Handler(srv))
	r.GET("/bit-mail/ping", _User_Ping0_HTTP_Handler(srv))
	r.POST("/bit-mail/user/login", _User_Login0_HTTP_Handler(srv))
	r.POST("/bit-mail/user/register", _User_Register0_HTTP_Handler(srv))
	r.GET("/bit-mail/user/getUserDetails", _User_GetUserDetailsByID0_HTTP_Handler(srv))
	r.GET("/bit-mail/contact/getContactListByUserId", _User_GetContactListByUserId0_HTTP_Handler(srv))
	r.POST("/bit-mail/contact/uploadContact", _User_UploadContact0_HTTP_Handler(srv))
	r.POST("/bit-mail/contact/delContactByUserID", _User_DelContactByUserID0_HTTP_Handler(srv))
}

func _User_ListUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_Ping0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserPing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*PingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PingReply)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_Register0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserDetailsByID0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserDetailsByIDRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserDetailsByID)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserDetailsByID(ctx, req.(*GetUserDetailsByIDRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserDetailsByIDReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetContactListByUserId0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetContactListByUserIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetContactListByUserId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetContactListByUserId(ctx, req.(*GetContactListByUserIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetContactListByUserIdReply)
		return ctx.Result(200, reply)
	}
}

func _User_UploadContact0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadContactRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUploadContact)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadContact(ctx, req.(*UploadContactRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadContactReply)
		return ctx.Result(200, reply)
	}
}

func _User_DelContactByUserID0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelContactByUserIDRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDelContactByUserID)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelContactByUserID(ctx, req.(*DelContactByUserIDRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelContactByUserIDReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	DelContactByUserID(ctx context.Context, req *DelContactByUserIDRequest, opts ...http.CallOption) (rsp *DelContactByUserIDReply, err error)
	GetContactListByUserId(ctx context.Context, req *GetContactListByUserIdRequest, opts ...http.CallOption) (rsp *GetContactListByUserIdReply, err error)
	GetUserDetailsByID(ctx context.Context, req *GetUserDetailsByIDRequest, opts ...http.CallOption) (rsp *GetUserDetailsByIDReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Ping(ctx context.Context, req *PingRequest, opts ...http.CallOption) (rsp *PingReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	UploadContact(ctx context.Context, req *UploadContactRequest, opts ...http.CallOption) (rsp *UploadContactReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) DelContactByUserID(ctx context.Context, in *DelContactByUserIDRequest, opts ...http.CallOption) (*DelContactByUserIDReply, error) {
	var out DelContactByUserIDReply
	pattern := "/bit-mail/contact/delContactByUserID"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDelContactByUserID))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetContactListByUserId(ctx context.Context, in *GetContactListByUserIdRequest, opts ...http.CallOption) (*GetContactListByUserIdReply, error) {
	var out GetContactListByUserIdReply
	pattern := "/bit-mail/contact/getContactListByUserId"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetContactListByUserId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserDetailsByID(ctx context.Context, in *GetUserDetailsByIDRequest, opts ...http.CallOption) (*GetUserDetailsByIDReply, error) {
	var out GetUserDetailsByIDReply
	pattern := "/bit-mail/user/getUserDetails"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserDetailsByID))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/bit-mail/user/listUser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/bit-mail/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Ping(ctx context.Context, in *PingRequest, opts ...http.CallOption) (*PingReply, error) {
	var out PingReply
	pattern := "/bit-mail/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserPing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/bit-mail/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UploadContact(ctx context.Context, in *UploadContactRequest, opts ...http.CallOption) (*UploadContactReply, error) {
	var out UploadContactReply
	pattern := "/bit-mail/contact/uploadContact"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUploadContact))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
