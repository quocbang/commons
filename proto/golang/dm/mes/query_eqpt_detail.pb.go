// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_eqpt_detail.proto

package mes

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

type GetEqptDetailRequest struct {
	EqptId               string   `protobuf:"bytes,1,opt,name=eqpt_id,json=eqptId,proto3" json:"eqpt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEqptDetailRequest) Reset()         { *m = GetEqptDetailRequest{} }
func (m *GetEqptDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetEqptDetailRequest) ProtoMessage()    {}
func (*GetEqptDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcabe32cb3329d60, []int{0}
}

func (m *GetEqptDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEqptDetailRequest.Unmarshal(m, b)
}
func (m *GetEqptDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEqptDetailRequest.Marshal(b, m, deterministic)
}
func (m *GetEqptDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEqptDetailRequest.Merge(m, src)
}
func (m *GetEqptDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetEqptDetailRequest.Size(m)
}
func (m *GetEqptDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEqptDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEqptDetailRequest proto.InternalMessageInfo

func (m *GetEqptDetailRequest) GetEqptId() string {
	if m != nil {
		return m.EqptId
	}
	return ""
}

type GetEqptDetailReply struct {
	ProdTypes            string   `protobuf:"bytes,1,opt,name=prod_types,json=prodTypes,proto3" json:"prod_types,omitempty"`
	BayId                string   `protobuf:"bytes,2,opt,name=bay_id,json=bayId,proto3" json:"bay_id,omitempty"`
	Volume               float32  `protobuf:"fixed32,3,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEqptDetailReply) Reset()         { *m = GetEqptDetailReply{} }
func (m *GetEqptDetailReply) String() string { return proto.CompactTextString(m) }
func (*GetEqptDetailReply) ProtoMessage()    {}
func (*GetEqptDetailReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcabe32cb3329d60, []int{1}
}

func (m *GetEqptDetailReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEqptDetailReply.Unmarshal(m, b)
}
func (m *GetEqptDetailReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEqptDetailReply.Marshal(b, m, deterministic)
}
func (m *GetEqptDetailReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEqptDetailReply.Merge(m, src)
}
func (m *GetEqptDetailReply) XXX_Size() int {
	return xxx_messageInfo_GetEqptDetailReply.Size(m)
}
func (m *GetEqptDetailReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEqptDetailReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetEqptDetailReply proto.InternalMessageInfo

func (m *GetEqptDetailReply) GetProdTypes() string {
	if m != nil {
		return m.ProdTypes
	}
	return ""
}

func (m *GetEqptDetailReply) GetBayId() string {
	if m != nil {
		return m.BayId
	}
	return ""
}

func (m *GetEqptDetailReply) GetVolume() float32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func init() {
	proto.RegisterType((*GetEqptDetailRequest)(nil), "proto.dm.mes.GetEqptDetailRequest")
	proto.RegisterType((*GetEqptDetailReply)(nil), "proto.dm.mes.GetEqptDetailReply")
}

func init() { proto.RegisterFile("query_eqpt_detail.proto", fileDescriptor_dcabe32cb3329d60) }

var fileDescriptor_dcabe32cb3329d60 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x40, 0xd9, 0x13, 0x57, 0x2e, 0x58, 0x05, 0xf5, 0xae, 0x11, 0x8e, 0xab, 0xae, 0x4a, 0x40,
	0x1b, 0xb1, 0x14, 0x45, 0xae, 0x5d, 0xac, 0x6c, 0x96, 0xe4, 0x66, 0x58, 0x16, 0x33, 0x3b, 0xd9,
	0xcd, 0xec, 0x49, 0xfe, 0x5e, 0x36, 0x5c, 0x65, 0x15, 0xde, 0xcb, 0xbc, 0x81, 0x51, 0x9b, 0x71,
	0xc6, 0x29, 0xb7, 0x38, 0x46, 0x69, 0x01, 0xc5, 0xf5, 0xc1, 0xc4, 0x89, 0x85, 0xf5, 0x6d, 0x79,
	0x0c, 0x90, 0x21, 0x4c, 0x7b, 0xab, 0xee, 0x3e, 0x51, 0x3e, 0xc6, 0x28, 0xef, 0x65, 0xa8, 0xc1,
	0x71, 0xc6, 0x24, 0x7a, 0xa3, 0x6e, 0x4a, 0xda, 0xc3, 0xb6, 0xda, 0x55, 0x87, 0x75, 0x53, 0x2f,
	0x78, 0x84, 0xbd, 0x57, 0xfa, 0x5f, 0x10, 0x43, 0xd6, 0x8f, 0x4a, 0xc5, 0x89, 0xa1, 0x95, 0x1c,
	0x31, 0x5d, 0x8a, 0xf5, 0x62, 0xbe, 0x16, 0xa1, 0xef, 0x55, 0xed, 0x5d, 0x5e, 0x96, 0xad, 0xca,
	0xd7, 0xb5, 0x77, 0xf9, 0x08, 0xfa, 0x41, 0xd5, 0x67, 0x0e, 0x33, 0xe1, 0xf6, 0x6a, 0x57, 0x1d,
	0x56, 0xcd, 0x85, 0xde, 0x5e, 0xbf, 0x5f, 0xba, 0x5e, 0x82, 0xf3, 0xe6, 0x07, 0x07, 0x70, 0xe6,
	0xc4, 0x64, 0xe4, 0xd7, 0x16, 0xb0, 0x27, 0x26, 0xe2, 0x21, 0xd9, 0xf3, 0x93, 0x2d, 0x97, 0xd8,
	0x8e, 0x83, 0x1b, 0x3a, 0x0b, 0x64, 0x09, 0x93, 0xaf, 0x8b, 0x7c, 0xfe, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x49, 0x3d, 0x5c, 0x00, 0x01, 0x00, 0x00,
}