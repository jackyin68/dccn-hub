// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: usermgr/v1/micro/usermgr.proto

/*
Package usermgr is a generated protocol buffer package.

It is generated from these files:
	usermgr/v1/micro/usermgr.proto

It has these top-level messages:
	User
	LoginRequest
	LoginResponse
	LogoutRequest
	Token
	NewTokenResponse
*/
package usermgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common_proto1 "github.com/Ankr-network/dccn-common/protos/common"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = common_proto1.Error{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserMgr service

type UserMgrService interface {
	// Register Create a new user
	Register(ctx context.Context, in *User, opts ...client.CallOption) (*common_proto1.Error, error)
	// Login login
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// Logout need verify permission
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*common_proto1.Error, error)
	// Auth  validates user
	NewToken(ctx context.Context, in *User, opts ...client.CallOption) (*NewTokenResponse, error)
	// VerifyToken Validated Token
	VerifyToken(ctx context.Context, in *Token, opts ...client.CallOption) (*common_proto1.Error, error)
}

type userMgrService struct {
	c    client.Client
	name string
}

func NewUserMgrService(name string, c client.Client) UserMgrService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "usermgr"
	}
	return &userMgrService{
		c:    c,
		name: name,
	}
}

func (c *userMgrService) Register(ctx context.Context, in *User, opts ...client.CallOption) (*common_proto1.Error, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Register", in)
	out := new(common_proto1.Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*common_proto1.Error, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Logout", in)
	out := new(common_proto1.Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) NewToken(ctx context.Context, in *User, opts ...client.CallOption) (*NewTokenResponse, error) {
	req := c.c.NewRequest(c.name, "UserMgr.NewToken", in)
	out := new(NewTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) VerifyToken(ctx context.Context, in *Token, opts ...client.CallOption) (*common_proto1.Error, error) {
	req := c.c.NewRequest(c.name, "UserMgr.VerifyToken", in)
	out := new(common_proto1.Error)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserMgr service

type UserMgrHandler interface {
	// Register Create a new user
	Register(context.Context, *User, *common_proto1.Error) error
	// Login login
	Login(context.Context, *LoginRequest, *LoginResponse) error
	// Logout need verify permission
	Logout(context.Context, *LogoutRequest, *common_proto1.Error) error
	// Auth  validates user
	NewToken(context.Context, *User, *NewTokenResponse) error
	// VerifyToken Validated Token
	VerifyToken(context.Context, *Token, *common_proto1.Error) error
}

func RegisterUserMgrHandler(s server.Server, hdlr UserMgrHandler, opts ...server.HandlerOption) error {
	type userMgr interface {
		Register(ctx context.Context, in *User, out *common_proto1.Error) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Logout(ctx context.Context, in *LogoutRequest, out *common_proto1.Error) error
		NewToken(ctx context.Context, in *User, out *NewTokenResponse) error
		VerifyToken(ctx context.Context, in *Token, out *common_proto1.Error) error
	}
	type UserMgr struct {
		userMgr
	}
	h := &userMgrHandler{hdlr}
	return s.Handle(s.NewHandler(&UserMgr{h}, opts...))
}

type userMgrHandler struct {
	UserMgrHandler
}

func (h *userMgrHandler) Register(ctx context.Context, in *User, out *common_proto1.Error) error {
	return h.UserMgrHandler.Register(ctx, in, out)
}

func (h *userMgrHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserMgrHandler.Login(ctx, in, out)
}

func (h *userMgrHandler) Logout(ctx context.Context, in *LogoutRequest, out *common_proto1.Error) error {
	return h.UserMgrHandler.Logout(ctx, in, out)
}

func (h *userMgrHandler) NewToken(ctx context.Context, in *User, out *NewTokenResponse) error {
	return h.UserMgrHandler.NewToken(ctx, in, out)
}

func (h *userMgrHandler) VerifyToken(ctx context.Context, in *Token, out *common_proto1.Error) error {
	return h.UserMgrHandler.VerifyToken(ctx, in, out)
}