// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: sys.proto

package sysclient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Sys_UserInfo_FullMethodName    = "/sysclient.Sys/UserInfo"
	Sys_UserAdd_FullMethodName     = "/sysclient.Sys/UserAdd"
	Sys_RedisAdd_FullMethodName    = "/sysclient.Sys/RedisAdd"
	Sys_RedisDelete_FullMethodName = "/sysclient.Sys/RedisDelete"
	Sys_RedisUpdate_FullMethodName = "/sysclient.Sys/RedisUpdate"
	Sys_RedisGet_FullMethodName    = "/sysclient.Sys/RedisGet"
)

// SysClient is the client API for Sys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysClient interface {
	UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
	UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
	// redis增删改查
	RedisAdd(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error)
	RedisDelete(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error)
	RedisUpdate(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error)
	RedisGet(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error)
}

type sysClient struct {
	cc grpc.ClientConnInterface
}

func NewSysClient(cc grpc.ClientConnInterface) SysClient {
	return &sysClient{cc}
}

func (c *sysClient) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	out := new(InfoResp)
	err := c.cc.Invoke(ctx, Sys_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	out := new(UserAddResp)
	err := c.cc.Invoke(ctx, Sys_UserAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RedisAdd(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error) {
	out := new(RedisResp)
	err := c.cc.Invoke(ctx, Sys_RedisAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RedisDelete(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error) {
	out := new(RedisResp)
	err := c.cc.Invoke(ctx, Sys_RedisDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RedisUpdate(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error) {
	out := new(RedisResp)
	err := c.cc.Invoke(ctx, Sys_RedisUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RedisGet(ctx context.Context, in *RedisReq, opts ...grpc.CallOption) (*RedisResp, error) {
	out := new(RedisResp)
	err := c.cc.Invoke(ctx, Sys_RedisGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysServer is the server API for Sys service.
// All implementations must embed UnimplementedSysServer
// for forward compatibility
type SysServer interface {
	UserInfo(context.Context, *InfoReq) (*InfoResp, error)
	UserAdd(context.Context, *UserAddReq) (*UserAddResp, error)
	// redis增删改查
	RedisAdd(context.Context, *RedisReq) (*RedisResp, error)
	RedisDelete(context.Context, *RedisReq) (*RedisResp, error)
	RedisUpdate(context.Context, *RedisReq) (*RedisResp, error)
	RedisGet(context.Context, *RedisReq) (*RedisResp, error)
	mustEmbedUnimplementedSysServer()
}

// UnimplementedSysServer must be embedded to have forward compatible implementations.
type UnimplementedSysServer struct {
}

func (UnimplementedSysServer) UserInfo(context.Context, *InfoReq) (*InfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedSysServer) UserAdd(context.Context, *UserAddReq) (*UserAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (UnimplementedSysServer) RedisAdd(context.Context, *RedisReq) (*RedisResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedisAdd not implemented")
}
func (UnimplementedSysServer) RedisDelete(context.Context, *RedisReq) (*RedisResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedisDelete not implemented")
}
func (UnimplementedSysServer) RedisUpdate(context.Context, *RedisReq) (*RedisResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedisUpdate not implemented")
}
func (UnimplementedSysServer) RedisGet(context.Context, *RedisReq) (*RedisResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedisGet not implemented")
}
func (UnimplementedSysServer) mustEmbedUnimplementedSysServer() {}

// UnsafeSysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysServer will
// result in compilation errors.
type UnsafeSysServer interface {
	mustEmbedUnimplementedSysServer()
}

func RegisterSysServer(s grpc.ServiceRegistrar, srv SysServer) {
	s.RegisterService(&Sys_ServiceDesc, srv)
}

func _Sys_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserInfo(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserAdd(ctx, req.(*UserAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RedisAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedisReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RedisAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_RedisAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RedisAdd(ctx, req.(*RedisReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RedisDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedisReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RedisDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_RedisDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RedisDelete(ctx, req.(*RedisReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RedisUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedisReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RedisUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_RedisUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RedisUpdate(ctx, req.(*RedisReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RedisGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedisReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RedisGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_RedisGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RedisGet(ctx, req.(*RedisReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sys_ServiceDesc is the grpc.ServiceDesc for Sys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sysclient.Sys",
	HandlerType: (*SysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfo",
			Handler:    _Sys_UserInfo_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _Sys_UserAdd_Handler,
		},
		{
			MethodName: "RedisAdd",
			Handler:    _Sys_RedisAdd_Handler,
		},
		{
			MethodName: "RedisDelete",
			Handler:    _Sys_RedisDelete_Handler,
		},
		{
			MethodName: "RedisUpdate",
			Handler:    _Sys_RedisUpdate_Handler,
		},
		{
			MethodName: "RedisGet",
			Handler:    _Sys_RedisGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys.proto",
}
