// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package sas

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

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x53, 0x7a, 0x29, 0xb9,
	0x7a, 0xc5, 0x89, 0xc5, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99,
	0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0xb5, 0x52, 0xd2,
	0x50, 0x59, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0x22, 0x69, 0x14,
	0xc5, 0xc5, 0x1c, 0xec, 0x18, 0x2c, 0x14, 0xcc, 0xc5, 0xed, 0x0a, 0x12, 0xf5, 0x4d, 0x2d, 0xc9,
	0xc8, 0x4f, 0x11, 0x12, 0xd3, 0x83, 0xe8, 0xd1, 0x83, 0xe9, 0xd1, 0x03, 0xcb, 0x4a, 0xe1, 0x10,
	0x57, 0x12, 0x69, 0xba, 0xfc, 0x64, 0x32, 0x13, 0x9f, 0x12, 0xa7, 0x7e, 0x71, 0x62, 0xb1, 0x7e,
	0x49, 0x6a, 0x71, 0x89, 0x15, 0xa3, 0x96, 0x93, 0x55, 0x94, 0x45, 0x7a, 0x66, 0x49, 0x4e, 0x62,
	0x92, 0x5e, 0x76, 0x6a, 0x5e, 0x4a, 0xa2, 0x5e, 0x72, 0x7e, 0xae, 0x5e, 0x49, 0xb9, 0x3e, 0x98,
	0xa3, 0x9f, 0x9c, 0x9f, 0x9b, 0x9b, 0x9f, 0x57, 0xac, 0x5f, 0x66, 0x04, 0x71, 0x98, 0x7e, 0x7a,
	0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x7e, 0x4a, 0x2e, 0xc8, 0x8c, 0x24, 0x36, 0xb0, 0xa0, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0xcb, 0x58, 0x17, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SASClient is the client API for SAS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SASClient interface {
	// EmptyMethod ??????????????????, ???????????????????????????.
	EmptyMethod(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sASClient struct {
	cc *grpc.ClientConn
}

func NewSASClient(cc *grpc.ClientConn) SASClient {
	return &sASClient{cc}
}

func (c *sASClient) EmptyMethod(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.dm.sas.SAS/EmptyMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SASServer is the server API for SAS service.
type SASServer interface {
	// EmptyMethod ??????????????????, ???????????????????????????.
	EmptyMethod(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterSASServer(s *grpc.Server, srv SASServer) {
	s.RegisterService(&_SAS_serviceDesc, srv)
}

func _SAS_EmptyMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SASServer).EmptyMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.sas.SAS/EmptyMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SASServer).EmptyMethod(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SAS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.dm.sas.SAS",
	HandlerType: (*SASServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyMethod",
			Handler:    _SAS_EmptyMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
