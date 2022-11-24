// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_authentication.proto

package utility

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("service_authentication.proto", fileDescriptor_72b14d11a6d9addd) }

var fileDescriptor_72b14d11a6d9addd = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x89, 0x07, 0x59, 0x06, 0x94, 0x3a, 0x16, 0xd1, 0xec, 0x16, 0xdc, 0x51, 0x51, 0xbb,
	0x30, 0x03, 0xf5, 0xcf, 0x61, 0x5d, 0x14, 0xb5, 0x8b, 0x08, 0x0a, 0x52, 0x10, 0xa4, 0x97, 0x65,
	0xda, 0xbc, 0xc6, 0xc1, 0x4c, 0x26, 0x66, 0xde, 0xac, 0x94, 0xe2, 0xc5, 0x8b, 0x07, 0xc1, 0x8b,
	0x78, 0xf6, 0xe2, 0x07, 0xf0, 0xbb, 0xf8, 0x15, 0xfc, 0x20, 0x92, 0x49, 0xb6, 0x4d, 0xdc, 0x66,
	0x37, 0x3d, 0xbd, 0xbc, 0x3c, 0x0f, 0x79, 0x7e, 0x79, 0x66, 0x18, 0xb2, 0x65, 0x21, 0x3d, 0x54,
	0x53, 0x38, 0x90, 0x19, 0xbe, 0x83, 0x18, 0xd5, 0x54, 0xa2, 0x32, 0x31, 0x4f, 0x52, 0x83, 0x86,
	0x76, 0xdc, 0xe0, 0x81, 0xe6, 0x19, 0xaa, 0x48, 0xe1, 0xcc, 0xdf, 0x0a, 0x8d, 0x09, 0x23, 0x10,
	0x32, 0x51, 0x42, 0xc6, 0xb1, 0x41, 0x67, 0xb7, 0x85, 0xdf, 0xdf, 0x2c, 0x55, 0xb7, 0x4d, 0xb2,
	0xb7, 0x02, 0x74, 0x82, 0xb3, 0x52, 0xdc, 0x76, 0x43, 0x04, 0x5a, 0x94, 0x1f, 0x13, 0x09, 0xa4,
	0x5a, 0x59, 0xbb, 0xc8, 0x1b, 0xfc, 0xdc, 0x20, 0xe7, 0x1f, 0xd7, 0x40, 0xe8, 0x0f, 0x8f, 0x74,
	0x5e, 0x66, 0x28, 0x11, 0x5e, 0x2d, 0xdc, 0x74, 0x87, 0xff, 0x0f, 0xc6, 0x9d, 0x47, 0x99, 0x78,
	0xe9, 0x1a, 0xc1, 0x87, 0x0c, 0x2c, 0xfa, 0x37, 0x8f, 0x9b, 0x97, 0xa6, 0x21, 0xa0, 0x54, 0xd1,
	0x08, 0x92, 0x68, 0xc6, 0x76, 0x3e, 0xff, 0xf9, 0xfb, 0xfd, 0xcc, 0x8d, 0xc1, 0x95, 0x15, 0x80,
	0x56, 0xcc, 0x55, 0xf0, 0x69, 0x77, 0x43, 0x97, 0x31, 0x34, 0x25, 0x9d, 0x21, 0x44, 0x50, 0xc3,
	0xba, 0x76, 0x52, 0xd2, 0x11, 0xce, 0x25, 0x5e, 0x94, 0xc4, 0x8f, 0x4a, 0xe2, 0xfb, 0x79, 0x49,
	0x6c, 0xdb, 0xa5, 0x6f, 0xf6, 0x9b, 0xd3, 0xe9, 0x2f, 0x8f, 0x9c, 0x7b, 0x06, 0xb8, 0x6e, 0x62,
	0xeb, 0x02, 0xf6, 0x1d, 0xc2, 0x23, 0xda, 0x8c, 0x30, 0x66, 0xac, 0xd7, 0x28, 0x8a, 0x10, 0x70,
	0xd7, 0xeb, 0xd3, 0x2f, 0x1e, 0xb9, 0xf0, 0x42, 0xd9, 0x0a, 0xe7, 0xf3, 0xa1, 0xa5, 0x0d, 0xff,
	0xed, 0xf7, 0x4e, 0xa2, 0xb3, 0xec, 0x81, 0x63, 0xba, 0x47, 0xbb, 0xab, 0x62, 0xc7, 0x3d, 0x76,
	0x79, 0x25, 0x4e, 0x49, 0xf2, 0xd5, 0x23, 0xdd, 0xa7, 0x29, 0x48, 0x84, 0xd7, 0x16, 0xd2, 0x4a,
	0x6f, 0x57, 0x8f, 0x87, 0xd6, 0x1d, 0xfe, 0xa9, 0x0e, 0x26, 0x1c, 0xd9, 0x6d, 0x76, 0x7d, 0x41,
	0x90, 0x59, 0x48, 0xad, 0x98, 0xe7, 0xe3, 0x20, 0xef, 0xa3, 0x42, 0x94, 0xd3, 0x7c, 0xf3, 0x48,
	0xb7, 0xb8, 0x33, 0x6b, 0xd3, 0x34, 0x5d, 0x9a, 0x3d, 0xc7, 0x70, 0xbf, 0x7f, 0xb7, 0x0d, 0x83,
	0x98, 0x2f, 0x97, 0x5c, 0xa3, 0xbf, 0x3d, 0x72, 0x31, 0x3f, 0xa8, 0x7a, 0x98, 0x6d, 0xc1, 0x73,
	0xca, 0xa1, 0xbd, 0x71, 0x58, 0x23, 0xda, 0xaa, 0x9a, 0x71, 0x9f, 0xdd, 0x6a, 0x85, 0x1f, 0x02,
	0x3e, 0x79, 0x38, 0xde, 0x0b, 0x15, 0x46, 0x72, 0xc2, 0xdf, 0x43, 0x1c, 0x48, 0x3e, 0x35, 0x9a,
	0xe3, 0x47, 0xe1, 0x16, 0x31, 0x35, 0x5a, 0xe7, 0xc6, 0xc3, 0x41, 0xf1, 0x02, 0x89, 0xd0, 0x44,
	0x32, 0x0e, 0x2b, 0x4f, 0xce, 0xe4, 0xac, 0x13, 0xee, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x5a, 0xa1, 0x73, 0xf8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	// MutatePermission adds/removes the columns that the permission can access.
	// Also you can assign star sign to column name to make whole table accessible or not.
	MutatePermission(ctx context.Context, in *MutationPermissionRequest, opts ...grpc.CallOption) (*PermissionDetailReply, error)
	// DeletePermission deletes specific permission.
	DeletePermission(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// GetPermission returns the tables that the permission can access and the columns that the permission cannot access.
	GetPermission(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*PermissionDetailReply, error)
	// ListPermissionIDs 查詢所有權限代號.
	ListPermissionIDs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Permissions, error)
	// CreateUserPermission 程式權限授權給使用者.
	CreateUserPermission(ctx context.Context, in *UserPermission, opts ...grpc.CallOption) (*UserPermission, error)
	// DeleteUserPermission 刪除使用者權限.
	DeleteUserPermission(ctx context.Context, in *UserPermission, opts ...grpc.CallOption) (*empty.Empty, error)
	// ListUserPermissions 查詢使用者所有權限.
	ListUserPermissions(ctx context.Context, in *UserPermission, opts ...grpc.CallOption) (*Permissions, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) MutatePermission(ctx context.Context, in *MutationPermissionRequest, opts ...grpc.CallOption) (*PermissionDetailReply, error) {
	out := new(PermissionDetailReply)
	err := c.cc.Invoke(ctx, "/proto.dm.utility.Authentication/MutatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) DeletePermission(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.dm.utility.Authentication/DeletePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GetPermission(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*PermissionDetailReply, error) {
	out := new(PermissionDetailReply)
	err := c.cc.Invoke(ctx, "/proto.dm.utility.Authentication/GetPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) ListPermissionIDs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Permissions, error) {
	out := new(Permissions)
	err := c.cc.Invoke(ctx, "/proto.dm.utility.Authentication/ListPermissionIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) CreateUserPermission(ctx context.Context, in *UserPermission, opts ...grpc.CallOption) (*UserPermission, error) {
	out := new(UserPermission)
	err := c.cc.Invoke(ctx, "/proto.dm.utility.Authentication/CreateUserPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) DeleteUserPermission(ctx context.Context, in *UserPermission, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.dm.utility.Authentication/DeleteUserPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) ListUserPermissions(ctx context.Context, in *UserPermission, opts ...grpc.CallOption) (*Permissions, error) {
	out := new(Permissions)
	err := c.cc.Invoke(ctx, "/proto.dm.utility.Authentication/ListUserPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	// MutatePermission adds/removes the columns that the permission can access.
	// Also you can assign star sign to column name to make whole table accessible or not.
	MutatePermission(context.Context, *MutationPermissionRequest) (*PermissionDetailReply, error)
	// DeletePermission deletes specific permission.
	DeletePermission(context.Context, *PermissionRequest) (*empty.Empty, error)
	// GetPermission returns the tables that the permission can access and the columns that the permission cannot access.
	GetPermission(context.Context, *PermissionRequest) (*PermissionDetailReply, error)
	// ListPermissionIDs 查詢所有權限代號.
	ListPermissionIDs(context.Context, *empty.Empty) (*Permissions, error)
	// CreateUserPermission 程式權限授權給使用者.
	CreateUserPermission(context.Context, *UserPermission) (*UserPermission, error)
	// DeleteUserPermission 刪除使用者權限.
	DeleteUserPermission(context.Context, *UserPermission) (*empty.Empty, error)
	// ListUserPermissions 查詢使用者所有權限.
	ListUserPermissions(context.Context, *UserPermission) (*Permissions, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_MutatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutationPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).MutatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.utility.Authentication/MutatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).MutatePermission(ctx, req.(*MutationPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.utility.Authentication/DeletePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).DeletePermission(ctx, req.(*PermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.utility.Authentication/GetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetPermission(ctx, req.(*PermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_ListPermissionIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).ListPermissionIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.utility.Authentication/ListPermissionIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).ListPermissionIDs(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_CreateUserPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).CreateUserPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.utility.Authentication/CreateUserPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).CreateUserPermission(ctx, req.(*UserPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_DeleteUserPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).DeleteUserPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.utility.Authentication/DeleteUserPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).DeleteUserPermission(ctx, req.(*UserPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_ListUserPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).ListUserPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.utility.Authentication/ListUserPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).ListUserPermissions(ctx, req.(*UserPermission))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.dm.utility.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutatePermission",
			Handler:    _Authentication_MutatePermission_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _Authentication_DeletePermission_Handler,
		},
		{
			MethodName: "GetPermission",
			Handler:    _Authentication_GetPermission_Handler,
		},
		{
			MethodName: "ListPermissionIDs",
			Handler:    _Authentication_ListPermissionIDs_Handler,
		},
		{
			MethodName: "CreateUserPermission",
			Handler:    _Authentication_CreateUserPermission_Handler,
		},
		{
			MethodName: "DeleteUserPermission",
			Handler:    _Authentication_DeleteUserPermission_Handler,
		},
		{
			MethodName: "ListUserPermissions",
			Handler:    _Authentication_ListUserPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_authentication.proto",
}