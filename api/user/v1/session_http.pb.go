// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.8
// source: api/user/v1/session.proto

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

const OperationSessionCreateSession = "/api.user.v1.Session/CreateSession"

type SessionHTTPServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionReply, error)
}

func RegisterSessionHTTPServer(s *http.Server, srv SessionHTTPServer) {
	r := s.Route("/")
	r.POST("/douyin/user/register", _Session_CreateSession0_HTTP_Handler(srv))
}

func _Session_CreateSession0_HTTP_Handler(srv SessionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSessionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSessionCreateSession)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSession(ctx, req.(*CreateSessionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateSessionReply)
		return ctx.Result(200, reply)
	}
}

type SessionHTTPClient interface {
	CreateSession(ctx context.Context, req *CreateSessionRequest, opts ...http.CallOption) (rsp *CreateSessionReply, err error)
}

type SessionHTTPClientImpl struct {
	cc *http.Client
}

func NewSessionHTTPClient(client *http.Client) SessionHTTPClient {
	return &SessionHTTPClientImpl{client}
}

func (c *SessionHTTPClientImpl) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...http.CallOption) (*CreateSessionReply, error) {
	var out CreateSessionReply
	pattern := "/douyin/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSessionCreateSession))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
