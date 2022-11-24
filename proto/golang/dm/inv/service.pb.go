// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package inv

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetInventoryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInventoryRequest) Reset()         { *m = GetInventoryRequest{} }
func (m *GetInventoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetInventoryRequest) ProtoMessage()    {}
func (*GetInventoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *GetInventoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInventoryRequest.Unmarshal(m, b)
}
func (m *GetInventoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInventoryRequest.Marshal(b, m, deterministic)
}
func (m *GetInventoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInventoryRequest.Merge(m, src)
}
func (m *GetInventoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetInventoryRequest.Size(m)
}
func (m *GetInventoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInventoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInventoryRequest proto.InternalMessageInfo

func (m *GetInventoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetInventoryReply struct {
	Specification        string   `protobuf:"bytes,1,opt,name=specification,proto3" json:"specification,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInventoryReply) Reset()         { *m = GetInventoryReply{} }
func (m *GetInventoryReply) String() string { return proto.CompactTextString(m) }
func (*GetInventoryReply) ProtoMessage()    {}
func (*GetInventoryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *GetInventoryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInventoryReply.Unmarshal(m, b)
}
func (m *GetInventoryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInventoryReply.Marshal(b, m, deterministic)
}
func (m *GetInventoryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInventoryReply.Merge(m, src)
}
func (m *GetInventoryReply) XXX_Size() int {
	return xxx_messageInfo_GetInventoryReply.Size(m)
}
func (m *GetInventoryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInventoryReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetInventoryReply proto.InternalMessageInfo

func (m *GetInventoryReply) GetSpecification() string {
	if m != nil {
		return m.Specification
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInventoryRequest)(nil), "proto.dm.inv.GetInventoryRequest")
	proto.RegisterType((*GetInventoryReply)(nil), "proto.dm.inv.GetInventoryReply")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x53, 0x7a, 0x29, 0xb9,
	0x7a, 0x99, 0x79, 0x65, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99,
	0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0xb5, 0x4a, 0xaa,
	0x5c, 0xc2, 0xee, 0xa9, 0x25, 0x9e, 0x79, 0x65, 0xa9, 0x79, 0x25, 0xf9, 0x45, 0x95, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0x96, 0x5c, 0x82, 0xa8, 0xca, 0x0a, 0x72, 0x2a, 0x85, 0x54,
	0xb8, 0x78, 0x8b, 0x0b, 0x52, 0x93, 0x33, 0xd3, 0x32, 0x93, 0xc1, 0x66, 0x42, 0xd5, 0xa3, 0x0a,
	0x1a, 0x95, 0x71, 0x31, 0x7b, 0xfa, 0x85, 0x09, 0xe5, 0x73, 0xf1, 0x20, 0x9b, 0x20, 0xa4, 0xa8,
	0x87, 0xec, 0x4a, 0x3d, 0x2c, 0x8e, 0x90, 0x92, 0xc7, 0xa7, 0xa4, 0x20, 0xa7, 0x52, 0x49, 0xba,
	0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xa2, 0x42, 0xc2, 0xfa, 0x99, 0x79, 0x65, 0x20, 0x0c, 0x91, 0xd4,
	0xaf, 0xce, 0x4c, 0xa9, 0x75, 0xb2, 0x8a, 0xb2, 0x48, 0xcf, 0x2c, 0xc9, 0x49, 0x4c, 0xd2, 0xcb,
	0x4e, 0xcd, 0x4b, 0x49, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x2b, 0x29, 0xd7, 0x07, 0x73, 0xf4, 0x93,
	0xf3, 0x73, 0x73, 0xf3, 0xf3, 0x8a, 0xf5, 0xcb, 0x8c, 0xf4, 0xc1, 0x76, 0xe8, 0xa7, 0xe7, 0xe7,
	0x24, 0xe6, 0xa5, 0xeb, 0xa7, 0xe4, 0x82, 0x8c, 0x49, 0x62, 0x03, 0x0b, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x26, 0xfa, 0x0a, 0x62, 0x59, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// INVClient is the client API for INV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type INVClient interface {
	// GetInventory 查詢成代資訊.
	GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*GetInventoryReply, error)
}

type iNVClient struct {
	cc *grpc.ClientConn
}

func NewINVClient(cc *grpc.ClientConn) INVClient {
	return &iNVClient{cc}
}

func (c *iNVClient) GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*GetInventoryReply, error) {
	out := new(GetInventoryReply)
	err := c.cc.Invoke(ctx, "/proto.dm.inv.INV/GetInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// INVServer is the server API for INV service.
type INVServer interface {
	// GetInventory 查詢成代資訊.
	GetInventory(context.Context, *GetInventoryRequest) (*GetInventoryReply, error)
}

func RegisterINVServer(s *grpc.Server, srv INVServer) {
	s.RegisterService(&_INV_serviceDesc, srv)
}

func _INV_GetInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(INVServer).GetInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.inv.INV/GetInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(INVServer).GetInventory(ctx, req.(*GetInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _INV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.dm.inv.INV",
	HandlerType: (*INVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInventory",
			Handler:    _INV_GetInventory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
