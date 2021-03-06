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
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserList, error)
	Add(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserModifyResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/interface.v1.User/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Add(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserModifyResp, error) {
	out := new(UserModifyResp)
	err := c.cc.Invoke(ctx, "/interface.v1.User/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	List(context.Context, *emptypb.Empty) (*UserList, error)
	Add(context.Context, *UserAddReq) (*UserModifyResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) List(context.Context, *emptypb.Empty) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserServer) Add(context.Context, *UserAddReq) (*UserModifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
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

func _User_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.User/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interface.v1.User/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Add(ctx, req.(*UserAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interface.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _User_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _User_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interface/v1/user.proto",
}
