// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/usermgr/usermgr.proto

/*
Package go_micro_srv_usermgr is a generated protocol buffer package.

It is generated from these files:
	proto/usermgr/usermgr.proto

It has these top-level messages:
	Email
	ID
	User
	Response
	TokenString
*/
package go_micro_srv_usermgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
	// Create Create a new user
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// Gets the specified user
	Get(ctx context.Context, in *ID, opts ...client.CallOption) (*User, error)
	// Gets user by email
	GetByEmail(ctx context.Context, in *Email, opts ...client.CallOption) (*User, error)
	// Auth  validates user
	NewToken(ctx context.Context, in *User, opts ...client.CallOption) (*TokenString, error)
	// VerifyToken Validated Token
	VerifyToken(ctx context.Context, in *TokenString, opts ...client.CallOption) (*Response, error)
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
		name = "go.micro.srv.usermgr"
	}
	return &userMgrService{
		c:    c,
		name: name,
	}
}

func (c *userMgrService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) Get(ctx context.Context, in *ID, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.Get", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) GetByEmail(ctx context.Context, in *Email, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserMgr.GetByEmail", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) NewToken(ctx context.Context, in *User, opts ...client.CallOption) (*TokenString, error) {
	req := c.c.NewRequest(c.name, "UserMgr.NewToken", in)
	out := new(TokenString)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrService) VerifyToken(ctx context.Context, in *TokenString, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserMgr.VerifyToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserMgr service

type UserMgrHandler interface {
	// Create Create a new user
	Create(context.Context, *User, *Response) error
	// Gets the specified user
	Get(context.Context, *ID, *User) error
	// Gets user by email
	GetByEmail(context.Context, *Email, *User) error
	// Auth  validates user
	NewToken(context.Context, *User, *TokenString) error
	// VerifyToken Validated Token
	VerifyToken(context.Context, *TokenString, *Response) error
}

func RegisterUserMgrHandler(s server.Server, hdlr UserMgrHandler, opts ...server.HandlerOption) error {
	type userMgr interface {
		Create(ctx context.Context, in *User, out *Response) error
		Get(ctx context.Context, in *ID, out *User) error
		GetByEmail(ctx context.Context, in *Email, out *User) error
		NewToken(ctx context.Context, in *User, out *TokenString) error
		VerifyToken(ctx context.Context, in *TokenString, out *Response) error
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

func (h *userMgrHandler) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserMgrHandler.Create(ctx, in, out)
}

func (h *userMgrHandler) Get(ctx context.Context, in *ID, out *User) error {
	return h.UserMgrHandler.Get(ctx, in, out)
}

func (h *userMgrHandler) GetByEmail(ctx context.Context, in *Email, out *User) error {
	return h.UserMgrHandler.GetByEmail(ctx, in, out)
}

func (h *userMgrHandler) NewToken(ctx context.Context, in *User, out *TokenString) error {
	return h.UserMgrHandler.NewToken(ctx, in, out)
}

func (h *userMgrHandler) VerifyToken(ctx context.Context, in *TokenString, out *Response) error {
	return h.UserMgrHandler.VerifyToken(ctx, in, out)
}