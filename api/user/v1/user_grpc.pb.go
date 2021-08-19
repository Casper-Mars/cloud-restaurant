// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Heath(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListUer(ctx context.Context, in *OnePageUserReq, opts ...grpc.CallOption) (*UserListResp, error)
	AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*UserModifyResp, error)
	ListUserByIds(ctx context.Context, in *ListUserByIdReq, opts ...grpc.CallOption) (*UserListResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Heath(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.v1.User/Heath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUer(ctx context.Context, in *OnePageUserReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, "/user.v1.User/ListUer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*UserModifyResp, error) {
	out := new(UserModifyResp)
	err := c.cc.Invoke(ctx, "/user.v1.User/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUserByIds(ctx context.Context, in *ListUserByIdReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, "/user.v1.User/ListUserByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Heath(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ListUer(context.Context, *OnePageUserReq) (*UserListResp, error)
	AddUser(context.Context, *AddUserReq) (*UserModifyResp, error)
	ListUserByIds(context.Context, *ListUserByIdReq) (*UserListResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Heath(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heath not implemented")
}
func (UnimplementedUserServer) ListUer(context.Context, *OnePageUserReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUer not implemented")
}
func (UnimplementedUserServer) AddUser(context.Context, *AddUserReq) (*UserModifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserServer) ListUserByIds(context.Context, *ListUserByIdReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserByIds not implemented")
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

func _User_Heath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Heath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/Heath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Heath(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnePageUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/ListUer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUer(ctx, req.(*OnePageUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUser(ctx, req.(*AddUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUserByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUserByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/ListUserByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUserByIds(ctx, req.(*ListUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heath",
			Handler:    _User_Heath_Handler,
		},
		{
			MethodName: "ListUer",
			Handler:    _User_ListUer_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _User_AddUser_Handler,
		},
		{
			MethodName: "ListUserByIds",
			Handler:    _User_ListUserByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/user.proto",
}
