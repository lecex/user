// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	// 用户验证授权
	Auth(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// token 验证
	ValidateToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 只限微服务之间调用
	// 根据用户ID获取授权
	AuthById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AuthByMobile(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Auth(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.Auth", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ValidateToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.ValidateToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) AuthById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.AuthById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) AuthByMobile(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.AuthByMobile", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	// 用户验证授权
	Auth(context.Context, *Request, *Response) error
	// token 验证
	ValidateToken(context.Context, *Request, *Response) error
	// 只限微服务之间调用
	// 根据用户ID获取授权
	AuthById(context.Context, *Request, *Response) error
	AuthByMobile(context.Context, *Request, *Response) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Auth(ctx context.Context, in *Request, out *Response) error
		ValidateToken(ctx context.Context, in *Request, out *Response) error
		AuthById(ctx context.Context, in *Request, out *Response) error
		AuthByMobile(ctx context.Context, in *Request, out *Response) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Auth(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.Auth(ctx, in, out)
}

func (h *authHandler) ValidateToken(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}

func (h *authHandler) AuthById(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.AuthById(ctx, in, out)
}

func (h *authHandler) AuthByMobile(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.AuthByMobile(ctx, in, out)
}
