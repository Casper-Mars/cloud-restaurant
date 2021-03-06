// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.4

package v1

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

type FoodHTTPServer interface {
	AddFood(context.Context, *AddFoodReq) (*FoodModifyResp, error)
	List(context.Context, *emptypb.Empty) (*FoodList, error)
}

func RegisterFoodHTTPServer(s *http.Server, srv FoodHTTPServer) {
	r := s.Route("/")
	r.POST("/food", _Food_AddFood0_HTTP_Handler(srv))
	r.GET("/food/list", _Food_List0_HTTP_Handler(srv))
}

func _Food_AddFood0_HTTP_Handler(srv FoodHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddFoodReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/interface.v1.Food/AddFood")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddFood(ctx, req.(*AddFoodReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FoodModifyResp)
		return ctx.Result(200, reply)
	}
}

func _Food_List0_HTTP_Handler(srv FoodHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/interface.v1.Food/List")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FoodList)
		return ctx.Result(200, reply)
	}
}

type FoodHTTPClient interface {
	AddFood(ctx context.Context, req *AddFoodReq, opts ...http.CallOption) (rsp *FoodModifyResp, err error)
	List(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *FoodList, err error)
}

type FoodHTTPClientImpl struct {
	cc *http.Client
}

func NewFoodHTTPClient(client *http.Client) FoodHTTPClient {
	return &FoodHTTPClientImpl{client}
}

func (c *FoodHTTPClientImpl) AddFood(ctx context.Context, in *AddFoodReq, opts ...http.CallOption) (*FoodModifyResp, error) {
	var out FoodModifyResp
	pattern := "/food"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/interface.v1.Food/AddFood"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *FoodHTTPClientImpl) List(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*FoodList, error) {
	var out FoodList
	pattern := "/food/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/interface.v1.Food/List"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
