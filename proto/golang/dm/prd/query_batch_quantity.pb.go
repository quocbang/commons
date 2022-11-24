// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_batch_quantity.proto

package prd

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetBatchQuantityRequest struct {
	Subno                string   `protobuf:"bytes,1,opt,name=subno,proto3" json:"subno,omitempty"`
	Partno               string   `protobuf:"bytes,2,opt,name=partno,proto3" json:"partno,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBatchQuantityRequest) Reset()         { *m = GetBatchQuantityRequest{} }
func (m *GetBatchQuantityRequest) String() string { return proto.CompactTextString(m) }
func (*GetBatchQuantityRequest) ProtoMessage()    {}
func (*GetBatchQuantityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bb71a1714d19ec9, []int{0}
}

func (m *GetBatchQuantityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBatchQuantityRequest.Unmarshal(m, b)
}
func (m *GetBatchQuantityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBatchQuantityRequest.Marshal(b, m, deterministic)
}
func (m *GetBatchQuantityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBatchQuantityRequest.Merge(m, src)
}
func (m *GetBatchQuantityRequest) XXX_Size() int {
	return xxx_messageInfo_GetBatchQuantityRequest.Size(m)
}
func (m *GetBatchQuantityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBatchQuantityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBatchQuantityRequest proto.InternalMessageInfo

func (m *GetBatchQuantityRequest) GetSubno() string {
	if m != nil {
		return m.Subno
	}
	return ""
}

func (m *GetBatchQuantityRequest) GetPartno() string {
	if m != nil {
		return m.Partno
	}
	return ""
}

type GetBatchQuantityReply struct {
	BatchQty             float32  `protobuf:"fixed32,1,opt,name=batch_qty,json=batchQty,proto3" json:"batch_qty,omitempty"`
	BatchUnit            string   `protobuf:"bytes,2,opt,name=batch_unit,json=batchUnit,proto3" json:"batch_unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBatchQuantityReply) Reset()         { *m = GetBatchQuantityReply{} }
func (m *GetBatchQuantityReply) String() string { return proto.CompactTextString(m) }
func (*GetBatchQuantityReply) ProtoMessage()    {}
func (*GetBatchQuantityReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bb71a1714d19ec9, []int{1}
}

func (m *GetBatchQuantityReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBatchQuantityReply.Unmarshal(m, b)
}
func (m *GetBatchQuantityReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBatchQuantityReply.Marshal(b, m, deterministic)
}
func (m *GetBatchQuantityReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBatchQuantityReply.Merge(m, src)
}
func (m *GetBatchQuantityReply) XXX_Size() int {
	return xxx_messageInfo_GetBatchQuantityReply.Size(m)
}
func (m *GetBatchQuantityReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBatchQuantityReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBatchQuantityReply proto.InternalMessageInfo

func (m *GetBatchQuantityReply) GetBatchQty() float32 {
	if m != nil {
		return m.BatchQty
	}
	return 0
}

func (m *GetBatchQuantityReply) GetBatchUnit() string {
	if m != nil {
		return m.BatchUnit
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBatchQuantityRequest)(nil), "proto.dm.prd.GetBatchQuantityRequest")
	proto.RegisterType((*GetBatchQuantityReply)(nil), "proto.dm.prd.GetBatchQuantityReply")
}

func init() { proto.RegisterFile("query_batch_quantity.proto", fileDescriptor_7bb71a1714d19ec9) }

var fileDescriptor_7bb71a1714d19ec9 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xbf, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xb9, 0x03, 0x0f, 0x6f, 0xb0, 0x0a, 0xfe, 0x38, 0x14, 0x41, 0xae, 0xb2, 0x4a, 0x40,
	0x1b, 0xb1, 0xbc, 0xe6, 0xea, 0x5b, 0xb1, 0xb1, 0x59, 0xb2, 0x9b, 0xb0, 0x06, 0x37, 0x33, 0xd9,
	0xec, 0x44, 0xc9, 0x7f, 0x2f, 0x66, 0xd3, 0x59, 0x0d, 0xef, 0xcd, 0xc7, 0xc7, 0x0c, 0xdc, 0x4e,
	0xc9, 0xc6, 0xdc, 0x76, 0x9a, 0xfb, 0xcf, 0x76, 0x4a, 0x1a, 0xd9, 0x71, 0x96, 0x21, 0x12, 0x93,
	0xb8, 0x28, 0x43, 0x1a, 0x2f, 0x43, 0x34, 0xfb, 0x23, 0xdc, 0x1c, 0x2d, 0x1f, 0xfe, 0xc0, 0x53,
	0xe5, 0x1a, 0x3b, 0x25, 0x3b, 0xb3, 0xb8, 0x84, 0xb3, 0x39, 0x75, 0x48, 0xbb, 0xd5, 0xc3, 0xea,
	0x71, 0xdb, 0x2c, 0x41, 0x5c, 0xc3, 0x26, 0xe8, 0xc8, 0x48, 0xbb, 0x75, 0xa9, 0x6b, 0xda, 0xbf,
	0xc1, 0xd5, 0x7f, 0x51, 0x18, 0xb3, 0xb8, 0x83, 0x6d, 0xbd, 0x83, 0x73, 0x51, 0xad, 0x9b, 0xf3,
	0x52, 0x9c, 0x38, 0x8b, 0x7b, 0x80, 0x65, 0x99, 0xd0, 0x71, 0x35, 0x2e, 0xf8, 0x3b, 0x3a, 0x3e,
	0xbc, 0x7e, 0xbc, 0x0c, 0x8e, 0x47, 0xdd, 0xc9, 0x2f, 0x8b, 0x46, 0xcb, 0x9e, 0xbc, 0xe4, 0x1f,
	0x55, 0x82, 0xea, 0xc9, 0x7b, 0xc2, 0x59, 0x7d, 0x3f, 0xa9, 0xf2, 0x92, 0x1a, 0x68, 0xd4, 0x38,
	0x28, 0xe3, 0x55, 0x88, 0xa6, 0xdb, 0x94, 0xf2, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x06, 0xfb,
	0x26, 0x3d, 0x0c, 0x01, 0x00, 0x00,
}