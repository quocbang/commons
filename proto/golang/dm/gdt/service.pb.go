// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package gdt

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

type GetMtrlSGRequest struct {
	Itnbr                string   `protobuf:"bytes,1,opt,name=itnbr,proto3" json:"itnbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMtrlSGRequest) Reset()         { *m = GetMtrlSGRequest{} }
func (m *GetMtrlSGRequest) String() string { return proto.CompactTextString(m) }
func (*GetMtrlSGRequest) ProtoMessage()    {}
func (*GetMtrlSGRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *GetMtrlSGRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMtrlSGRequest.Unmarshal(m, b)
}
func (m *GetMtrlSGRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMtrlSGRequest.Marshal(b, m, deterministic)
}
func (m *GetMtrlSGRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMtrlSGRequest.Merge(m, src)
}
func (m *GetMtrlSGRequest) XXX_Size() int {
	return xxx_messageInfo_GetMtrlSGRequest.Size(m)
}
func (m *GetMtrlSGRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMtrlSGRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMtrlSGRequest proto.InternalMessageInfo

func (m *GetMtrlSGRequest) GetItnbr() string {
	if m != nil {
		return m.Itnbr
	}
	return ""
}

type GetMtrlSGReply struct {
	Sg                   float32  `protobuf:"fixed32,1,opt,name=sg,proto3" json:"sg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMtrlSGReply) Reset()         { *m = GetMtrlSGReply{} }
func (m *GetMtrlSGReply) String() string { return proto.CompactTextString(m) }
func (*GetMtrlSGReply) ProtoMessage()    {}
func (*GetMtrlSGReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *GetMtrlSGReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMtrlSGReply.Unmarshal(m, b)
}
func (m *GetMtrlSGReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMtrlSGReply.Marshal(b, m, deterministic)
}
func (m *GetMtrlSGReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMtrlSGReply.Merge(m, src)
}
func (m *GetMtrlSGReply) XXX_Size() int {
	return xxx_messageInfo_GetMtrlSGReply.Size(m)
}
func (m *GetMtrlSGReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMtrlSGReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMtrlSGReply proto.InternalMessageInfo

func (m *GetMtrlSGReply) GetSg() float32 {
	if m != nil {
		return m.Sg
	}
	return 0
}

type GetMtrlsRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMtrlsRequest) Reset()         { *m = GetMtrlsRequest{} }
func (m *GetMtrlsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMtrlsRequest) ProtoMessage()    {}
func (*GetMtrlsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *GetMtrlsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMtrlsRequest.Unmarshal(m, b)
}
func (m *GetMtrlsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMtrlsRequest.Marshal(b, m, deterministic)
}
func (m *GetMtrlsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMtrlsRequest.Merge(m, src)
}
func (m *GetMtrlsRequest) XXX_Size() int {
	return xxx_messageInfo_GetMtrlsRequest.Size(m)
}
func (m *GetMtrlsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMtrlsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMtrlsRequest proto.InternalMessageInfo

func (m *GetMtrlsRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetMtrlsReply struct {
	Mtrls                []*GetMtrlsReply_Mtrl `protobuf:"bytes,1,rep,name=mtrls,proto3" json:"mtrls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetMtrlsReply) Reset()         { *m = GetMtrlsReply{} }
func (m *GetMtrlsReply) String() string { return proto.CompactTextString(m) }
func (*GetMtrlsReply) ProtoMessage()    {}
func (*GetMtrlsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *GetMtrlsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMtrlsReply.Unmarshal(m, b)
}
func (m *GetMtrlsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMtrlsReply.Marshal(b, m, deterministic)
}
func (m *GetMtrlsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMtrlsReply.Merge(m, src)
}
func (m *GetMtrlsReply) XXX_Size() int {
	return xxx_messageInfo_GetMtrlsReply.Size(m)
}
func (m *GetMtrlsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMtrlsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMtrlsReply proto.InternalMessageInfo

func (m *GetMtrlsReply) GetMtrls() []*GetMtrlsReply_Mtrl {
	if m != nil {
		return m.Mtrls
	}
	return nil
}

type GetMtrlsReply_Mtrl struct {
	Matno                string   `protobuf:"bytes,1,opt,name=matno,proto3" json:"matno,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMtrlsReply_Mtrl) Reset()         { *m = GetMtrlsReply_Mtrl{} }
func (m *GetMtrlsReply_Mtrl) String() string { return proto.CompactTextString(m) }
func (*GetMtrlsReply_Mtrl) ProtoMessage()    {}
func (*GetMtrlsReply_Mtrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3, 0}
}

func (m *GetMtrlsReply_Mtrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMtrlsReply_Mtrl.Unmarshal(m, b)
}
func (m *GetMtrlsReply_Mtrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMtrlsReply_Mtrl.Marshal(b, m, deterministic)
}
func (m *GetMtrlsReply_Mtrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMtrlsReply_Mtrl.Merge(m, src)
}
func (m *GetMtrlsReply_Mtrl) XXX_Size() int {
	return xxx_messageInfo_GetMtrlsReply_Mtrl.Size(m)
}
func (m *GetMtrlsReply_Mtrl) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMtrlsReply_Mtrl.DiscardUnknown(m)
}

var xxx_messageInfo_GetMtrlsReply_Mtrl proto.InternalMessageInfo

func (m *GetMtrlsReply_Mtrl) GetMatno() string {
	if m != nil {
		return m.Matno
	}
	return ""
}

func (m *GetMtrlsReply_Mtrl) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMtrlSGRequest)(nil), "proto.dm.gdt.GetMtrlSGRequest")
	proto.RegisterType((*GetMtrlSGReply)(nil), "proto.dm.gdt.GetMtrlSGReply")
	proto.RegisterType((*GetMtrlsRequest)(nil), "proto.dm.gdt.GetMtrlsRequest")
	proto.RegisterType((*GetMtrlsReply)(nil), "proto.dm.gdt.GetMtrlsReply")
	proto.RegisterType((*GetMtrlsReply_Mtrl)(nil), "proto.dm.gdt.GetMtrlsReply.Mtrl")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x25, 0x99, 0x5d, 0x71, 0x5b, 0x77, 0x56, 0xda, 0x3d, 0x8c, 0x71, 0xd5, 0xd0, 0x20, 0x0c,
	0x1e, 0xba, 0x21, 0x82, 0xc8, 0x1e, 0x45, 0xc8, 0xc9, 0x4b, 0xf4, 0xe4, 0x45, 0x3a, 0x93, 0xa6,
	0x09, 0xf6, 0xd7, 0xa6, 0x6b, 0x47, 0x82, 0x37, 0xff, 0x82, 0x67, 0x7f, 0x95, 0x7f, 0xc1, 0x1f,
	0x22, 0xe9, 0x4e, 0xc6, 0x95, 0xc9, 0xcc, 0x29, 0xfd, 0xea, 0x3d, 0xea, 0xd5, 0xab, 0x14, 0x3a,
	0xf7, 0xa2, 0xdb, 0xb6, 0x1b, 0x41, 0x5d, 0x67, 0xc1, 0xe2, 0x87, 0xe1, 0x43, 0x1b, 0x4d, 0x65,
	0x03, 0xd9, 0x95, 0xb4, 0x56, 0x2a, 0xc1, 0xb8, 0x6b, 0x19, 0x37, 0xc6, 0x02, 0x87, 0xd6, 0x1a,
	0x1f, 0xb5, 0xd9, 0x8b, 0xf0, 0x61, 0x8d, 0x66, 0xb2, 0x01, 0x76, 0x73, 0x2b, 0xba, 0xfe, 0x8b,
	0x6c, 0xa0, 0xe6, 0xdd, 0x71, 0x81, 0xe3, 0x10, 0x05, 0x64, 0x8d, 0x1e, 0x95, 0x02, 0x3e, 0x40,
	0xa7, 0x3e, 0x96, 0x95, 0xb8, 0xb9, 0x15, 0x1e, 0xf0, 0x25, 0x3a, 0x6d, 0xc1, 0xd4, 0xdd, 0x2a,
	0xc9, 0x93, 0xf5, 0x59, 0x15, 0x01, 0xc9, 0xd1, 0xf2, 0x8e, 0xd2, 0xa9, 0x1e, 0x2f, 0x51, 0xea,
	0x65, 0x10, 0xa5, 0x55, 0xea, 0x25, 0x79, 0x89, 0x2e, 0x46, 0x85, 0x9f, 0x5a, 0x61, 0x74, 0x02,
	0xbd, 0x13, 0x63, 0xa7, 0xf0, 0x26, 0xdf, 0xd1, 0xf9, 0x3f, 0xd9, 0xd0, 0xe7, 0x0d, 0x3a, 0xd5,
	0x03, 0x5a, 0x25, 0xf9, 0x62, 0xfd, 0xa0, 0xc8, 0xe9, 0xdd, 0x0d, 0xd0, 0xff, 0xb4, 0x74, 0x78,
	0x56, 0x51, 0x9e, 0x15, 0xe8, 0x64, 0x80, 0xc3, 0xbc, 0x9a, 0x83, 0xb1, 0xd3, 0xbc, 0x01, 0x0c,
	0x55, 0x25, 0xb6, 0x42, 0xad, 0xd2, 0x58, 0x0d, 0xa0, 0xf8, 0xb5, 0x40, 0x8b, 0xf2, 0xfd, 0x27,
	0x2c, 0xd0, 0xd9, 0x2e, 0x0d, 0x7e, 0x3e, 0xeb, 0xb8, 0x5b, 0x48, 0x76, 0x75, 0x90, 0x77, 0xaa,
	0x27, 0xd9, 0x8f, 0xdf, 0x7f, 0x7e, 0xa6, 0x97, 0xe4, 0x22, 0xec, 0x58, 0x07, 0x86, 0x49, 0x01,
	0xd7, 0xc9, 0x2b, 0xcc, 0xd1, 0xfd, 0x69, 0x7e, 0xfc, 0xec, 0x50, 0xae, 0x68, 0xf2, 0xf4, 0x48,
	0x6c, 0xf2, 0x24, 0x78, 0x3c, 0x26, 0xcb, 0x9d, 0x87, 0x9f, 0x2c, 0x62, 0x92, 0x32, 0xfc, 0xf5,
	0x99, 0x24, 0x91, 0x38, 0x9c, 0x64, 0xe2, 0xf7, 0x93, 0xc4, 0x43, 0xda, 0xb3, 0x71, 0x1c, 0xe6,
	0x6d, 0x1c, 0x87, 0xa3, 0x36, 0x81, 0x9f, 0xb5, 0x71, 0x1c, 0x46, 0x9b, 0x77, 0xd7, 0x9f, 0xdf,
	0xca, 0x16, 0x14, 0xaf, 0xe9, 0x57, 0x61, 0x1a, 0x4e, 0x37, 0x56, 0x53, 0xf8, 0xc6, 0x02, 0x60,
	0x1b, 0xab, 0xb5, 0x35, 0x9e, 0x6d, 0x0b, 0x16, 0xef, 0x5a, 0x5a, 0xc5, 0x8d, 0x1c, 0xcf, 0xbb,
	0xbe, 0x17, 0x8a, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x82, 0x9c, 0xa8, 0x1d, 0x51, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GDTClient is the client API for GDT service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GDTClient interface {
	//回傳原料比重
	GetMtrlSG(ctx context.Context, in *GetMtrlSGRequest, opts ...grpc.CallOption) (*GetMtrlSGReply, error)
	//回傳原料大類對應的5碼料號+等級
	GetMtrls(ctx context.Context, in *GetMtrlsRequest, opts ...grpc.CallOption) (*GetMtrlsReply, error)
	// Obtain data using mtrlLotID from Gdtbar
	GetGdtbar(ctx context.Context, in *GetGdtbarRequest, opts ...grpc.CallOption) (*GetGdtbarReply, error)
	// Obtain minimum usage quantity of queried material
	GetGdtpat(ctx context.Context, in *GetGdtpatRequest, opts ...grpc.CallOption) (*GetGdtpatReply, error)
}

type gDTClient struct {
	cc *grpc.ClientConn
}

func NewGDTClient(cc *grpc.ClientConn) GDTClient {
	return &gDTClient{cc}
}

func (c *gDTClient) GetMtrlSG(ctx context.Context, in *GetMtrlSGRequest, opts ...grpc.CallOption) (*GetMtrlSGReply, error) {
	out := new(GetMtrlSGReply)
	err := c.cc.Invoke(ctx, "/proto.dm.gdt.GDT/GetMtrlSG", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gDTClient) GetMtrls(ctx context.Context, in *GetMtrlsRequest, opts ...grpc.CallOption) (*GetMtrlsReply, error) {
	out := new(GetMtrlsReply)
	err := c.cc.Invoke(ctx, "/proto.dm.gdt.GDT/GetMtrls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gDTClient) GetGdtbar(ctx context.Context, in *GetGdtbarRequest, opts ...grpc.CallOption) (*GetGdtbarReply, error) {
	out := new(GetGdtbarReply)
	err := c.cc.Invoke(ctx, "/proto.dm.gdt.GDT/GetGdtbar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gDTClient) GetGdtpat(ctx context.Context, in *GetGdtpatRequest, opts ...grpc.CallOption) (*GetGdtpatReply, error) {
	out := new(GetGdtpatReply)
	err := c.cc.Invoke(ctx, "/proto.dm.gdt.GDT/GetGdtpat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GDTServer is the server API for GDT service.
type GDTServer interface {
	//回傳原料比重
	GetMtrlSG(context.Context, *GetMtrlSGRequest) (*GetMtrlSGReply, error)
	//回傳原料大類對應的5碼料號+等級
	GetMtrls(context.Context, *GetMtrlsRequest) (*GetMtrlsReply, error)
	// Obtain data using mtrlLotID from Gdtbar
	GetGdtbar(context.Context, *GetGdtbarRequest) (*GetGdtbarReply, error)
	// Obtain minimum usage quantity of queried material
	GetGdtpat(context.Context, *GetGdtpatRequest) (*GetGdtpatReply, error)
}

func RegisterGDTServer(s *grpc.Server, srv GDTServer) {
	s.RegisterService(&_GDT_serviceDesc, srv)
}

func _GDT_GetMtrlSG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMtrlSGRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GDTServer).GetMtrlSG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.gdt.GDT/GetMtrlSG",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GDTServer).GetMtrlSG(ctx, req.(*GetMtrlSGRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GDT_GetMtrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMtrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GDTServer).GetMtrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.gdt.GDT/GetMtrls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GDTServer).GetMtrls(ctx, req.(*GetMtrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GDT_GetGdtbar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGdtbarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GDTServer).GetGdtbar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.gdt.GDT/GetGdtbar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GDTServer).GetGdtbar(ctx, req.(*GetGdtbarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GDT_GetGdtpat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGdtpatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GDTServer).GetGdtpat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.gdt.GDT/GetGdtpat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GDTServer).GetGdtpat(ctx, req.(*GetGdtpatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GDT_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.dm.gdt.GDT",
	HandlerType: (*GDTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMtrlSG",
			Handler:    _GDT_GetMtrlSG_Handler,
		},
		{
			MethodName: "GetMtrls",
			Handler:    _GDT_GetMtrls_Handler,
		},
		{
			MethodName: "GetGdtbar",
			Handler:    _GDT_GetGdtbar_Handler,
		},
		{
			MethodName: "GetGdtpat",
			Handler:    _GDT_GetGdtpat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}