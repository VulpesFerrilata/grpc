// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: protoc/auth/auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Auth service

func NewAuthEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Auth service

type AuthService interface {
	CreateUserCredential(ctx context.Context, in *UserCredentialRequest, opts ...client.CallOption) (*empty.Empty, error)
	Authenticate(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*ClaimResponse, error)
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

func (c *authService) CreateUserCredential(ctx context.Context, in *UserCredentialRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateUserCredential", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Authenticate(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*ClaimResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Authenticate", in)
	out := new(ClaimResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	CreateUserCredential(context.Context, *UserCredentialRequest, *empty.Empty) error
	Authenticate(context.Context, *TokenRequest, *ClaimResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		CreateUserCredential(ctx context.Context, in *UserCredentialRequest, out *empty.Empty) error
		Authenticate(ctx context.Context, in *TokenRequest, out *ClaimResponse) error
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

func (h *authHandler) CreateUserCredential(ctx context.Context, in *UserCredentialRequest, out *empty.Empty) error {
	return h.AuthHandler.CreateUserCredential(ctx, in, out)
}

func (h *authHandler) Authenticate(ctx context.Context, in *TokenRequest, out *ClaimResponse) error {
	return h.AuthHandler.Authenticate(ctx, in, out)
}
