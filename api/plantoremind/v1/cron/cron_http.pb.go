// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.0

package cron

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type CronHTTPServer interface {
	CreateCron(context.Context, *CreateCronRequest) (*emptypb.Empty, error)
}

func RegisterCronHTTPServer(s *http.Server, srv CronHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/cron/create", _Cron_CreateCron0_HTTP_Handler(srv))
}

func _Cron_CreateCron0_HTTP_Handler(srv CronHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCronRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.plantoremind.v1.cron.Cron/CreateCron")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCron(ctx, req.(*CreateCronRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type CronHTTPClient interface {
	CreateCron(ctx context.Context, req *CreateCronRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type CronHTTPClientImpl struct {
	cc *http.Client
}

func NewCronHTTPClient(client *http.Client) CronHTTPClient {
	return &CronHTTPClientImpl{client}
}

func (c *CronHTTPClientImpl) CreateCron(ctx context.Context, in *CreateCronRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/api/v1/cron/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.plantoremind.v1.cron.Cron/CreateCron"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
