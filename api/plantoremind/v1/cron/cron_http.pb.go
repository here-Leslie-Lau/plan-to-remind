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
	DeleteCron(context.Context, *DeleteCronRequest) (*emptypb.Empty, error)
	GetCron(context.Context, *GetCronRequest) (*GetCronReply, error)
	ListCron(context.Context, *ListCronRequest) (*ListCronReply, error)
	UpdateCron(context.Context, *UpdateCronRequest) (*emptypb.Empty, error)
}

func RegisterCronHTTPServer(s *http.Server, srv CronHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/cron/create", _Cron_CreateCron0_HTTP_Handler(srv))
	r.POST("/api/v1/cron/update", _Cron_UpdateCron0_HTTP_Handler(srv))
	r.POST("/api/v1/cron/delete", _Cron_DeleteCron0_HTTP_Handler(srv))
	r.POST("/api/v1/cron/get", _Cron_GetCron0_HTTP_Handler(srv))
	r.POST("/api/v1/cron/list", _Cron_ListCron0_HTTP_Handler(srv))
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

func _Cron_UpdateCron0_HTTP_Handler(srv CronHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCronRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.plantoremind.v1.cron.Cron/UpdateCron")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCron(ctx, req.(*UpdateCronRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Cron_DeleteCron0_HTTP_Handler(srv CronHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCronRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.plantoremind.v1.cron.Cron/DeleteCron")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCron(ctx, req.(*DeleteCronRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Cron_GetCron0_HTTP_Handler(srv CronHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCronRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.plantoremind.v1.cron.Cron/GetCron")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCron(ctx, req.(*GetCronRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCronReply)
		return ctx.Result(200, reply)
	}
}

func _Cron_ListCron0_HTTP_Handler(srv CronHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCronRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.plantoremind.v1.cron.Cron/ListCron")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCron(ctx, req.(*ListCronRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCronReply)
		return ctx.Result(200, reply)
	}
}

type CronHTTPClient interface {
	CreateCron(ctx context.Context, req *CreateCronRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteCron(ctx context.Context, req *DeleteCronRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetCron(ctx context.Context, req *GetCronRequest, opts ...http.CallOption) (rsp *GetCronReply, err error)
	ListCron(ctx context.Context, req *ListCronRequest, opts ...http.CallOption) (rsp *ListCronReply, err error)
	UpdateCron(ctx context.Context, req *UpdateCronRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
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

func (c *CronHTTPClientImpl) DeleteCron(ctx context.Context, in *DeleteCronRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/api/v1/cron/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.plantoremind.v1.cron.Cron/DeleteCron"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CronHTTPClientImpl) GetCron(ctx context.Context, in *GetCronRequest, opts ...http.CallOption) (*GetCronReply, error) {
	var out GetCronReply
	pattern := "/api/v1/cron/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.plantoremind.v1.cron.Cron/GetCron"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CronHTTPClientImpl) ListCron(ctx context.Context, in *ListCronRequest, opts ...http.CallOption) (*ListCronReply, error) {
	var out ListCronReply
	pattern := "/api/v1/cron/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.plantoremind.v1.cron.Cron/ListCron"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CronHTTPClientImpl) UpdateCron(ctx context.Context, in *UpdateCronRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/api/v1/cron/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.plantoremind.v1.cron.Cron/UpdateCron"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
