// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_replace_mtrl.proto

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

type ReplaceMtrl struct {
	ReplaceById          string   `protobuf:"bytes,1,opt,name=replace_by_id,json=replaceById,proto3" json:"replace_by_id,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Ratio                float32  `protobuf:"fixed32,3,opt,name=ratio,proto3" json:"ratio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplaceMtrl) Reset()         { *m = ReplaceMtrl{} }
func (m *ReplaceMtrl) String() string { return proto.CompactTextString(m) }
func (*ReplaceMtrl) ProtoMessage()    {}
func (*ReplaceMtrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_722b570a034a0dbd, []int{0}
}

func (m *ReplaceMtrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceMtrl.Unmarshal(m, b)
}
func (m *ReplaceMtrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceMtrl.Marshal(b, m, deterministic)
}
func (m *ReplaceMtrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceMtrl.Merge(m, src)
}
func (m *ReplaceMtrl) XXX_Size() int {
	return xxx_messageInfo_ReplaceMtrl.Size(m)
}
func (m *ReplaceMtrl) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceMtrl.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceMtrl proto.InternalMessageInfo

func (m *ReplaceMtrl) GetReplaceById() string {
	if m != nil {
		return m.ReplaceById
	}
	return ""
}

func (m *ReplaceMtrl) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *ReplaceMtrl) GetRatio() float32 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

type GetReplaceMtrlRequest struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReplaceMtrlRequest) Reset()         { *m = GetReplaceMtrlRequest{} }
func (m *GetReplaceMtrlRequest) String() string { return proto.CompactTextString(m) }
func (*GetReplaceMtrlRequest) ProtoMessage()    {}
func (*GetReplaceMtrlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_722b570a034a0dbd, []int{1}
}

func (m *GetReplaceMtrlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReplaceMtrlRequest.Unmarshal(m, b)
}
func (m *GetReplaceMtrlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReplaceMtrlRequest.Marshal(b, m, deterministic)
}
func (m *GetReplaceMtrlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReplaceMtrlRequest.Merge(m, src)
}
func (m *GetReplaceMtrlRequest) XXX_Size() int {
	return xxx_messageInfo_GetReplaceMtrlRequest.Size(m)
}
func (m *GetReplaceMtrlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReplaceMtrlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetReplaceMtrlRequest proto.InternalMessageInfo

func (m *GetReplaceMtrlRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *GetReplaceMtrlRequest) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type GetReplaceMtrlReply struct {
	RpMtrls              []*ReplaceMtrl `protobuf:"bytes,1,rep,name=rp_mtrls,json=rpMtrls,proto3" json:"rp_mtrls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetReplaceMtrlReply) Reset()         { *m = GetReplaceMtrlReply{} }
func (m *GetReplaceMtrlReply) String() string { return proto.CompactTextString(m) }
func (*GetReplaceMtrlReply) ProtoMessage()    {}
func (*GetReplaceMtrlReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_722b570a034a0dbd, []int{2}
}

func (m *GetReplaceMtrlReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReplaceMtrlReply.Unmarshal(m, b)
}
func (m *GetReplaceMtrlReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReplaceMtrlReply.Marshal(b, m, deterministic)
}
func (m *GetReplaceMtrlReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReplaceMtrlReply.Merge(m, src)
}
func (m *GetReplaceMtrlReply) XXX_Size() int {
	return xxx_messageInfo_GetReplaceMtrlReply.Size(m)
}
func (m *GetReplaceMtrlReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReplaceMtrlReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReplaceMtrlReply proto.InternalMessageInfo

func (m *GetReplaceMtrlReply) GetRpMtrls() []*ReplaceMtrl {
	if m != nil {
		return m.RpMtrls
	}
	return nil
}

func init() {
	proto.RegisterType((*ReplaceMtrl)(nil), "proto.dm.mes.ReplaceMtrl")
	proto.RegisterType((*GetReplaceMtrlRequest)(nil), "proto.dm.mes.GetReplaceMtrlRequest")
	proto.RegisterType((*GetReplaceMtrlReply)(nil), "proto.dm.mes.GetReplaceMtrlReply")
}

func init() { proto.RegisterFile("query_replace_mtrl.proto", fileDescriptor_722b570a034a0dbd) }

var fileDescriptor_722b570a034a0dbd = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x4f, 0x4b, 0xc4, 0x30,
	0x14, 0xc4, 0xe9, 0x2e, 0xfe, 0xd9, 0x54, 0x2f, 0x51, 0x21, 0x1e, 0x84, 0xd2, 0x53, 0x4f, 0x09,
	0xac, 0x1e, 0xc4, 0xe3, 0x5e, 0x44, 0xd4, 0x4b, 0x8f, 0x82, 0x94, 0xb4, 0x79, 0x94, 0x62, 0xd2,
	0x64, 0x93, 0xd7, 0x95, 0x7e, 0x7b, 0xd9, 0x74, 0xc5, 0x22, 0xec, 0xe9, 0x31, 0xbf, 0x64, 0x66,
	0x18, 0xc2, 0xb6, 0x03, 0xf8, 0xb1, 0xf2, 0xe0, 0xb4, 0x6c, 0xa0, 0x32, 0xe8, 0x35, 0x77, 0xde,
	0xa2, 0xa5, 0x17, 0xf1, 0x70, 0x65, 0xb8, 0x81, 0x90, 0x7f, 0x92, 0xb4, 0x9c, 0xfe, 0xbc, 0xa3,
	0xd7, 0x34, 0x27, 0x97, 0xbf, 0x96, 0x7a, 0xac, 0x3a, 0xc5, 0x92, 0x2c, 0x29, 0x56, 0x65, 0x7a,
	0x80, 0x9b, 0xf1, 0x45, 0xd1, 0x6b, 0x72, 0xa2, 0x61, 0x07, 0x9a, 0x2d, 0xe2, 0xdb, 0x24, 0xf6,
	0xd4, 0x4b, 0xec, 0x2c, 0x5b, 0x66, 0x49, 0xb1, 0x28, 0x27, 0x91, 0xbf, 0x91, 0x9b, 0x67, 0xc0,
	0x59, 0x43, 0x09, 0xdb, 0x01, 0x02, 0xd2, 0x3b, 0x42, 0x9c, 0xb7, 0x6a, 0x68, 0xf0, 0xaf, 0x65,
	0x75, 0x20, 0xc7, 0x3a, 0xf2, 0x57, 0x72, 0xf5, 0x3f, 0xcd, 0xe9, 0x91, 0x3e, 0x90, 0x73, 0xef,
	0xe2, 0xc4, 0xc0, 0x92, 0x6c, 0x59, 0xa4, 0xeb, 0x5b, 0x3e, 0x1f, 0xc9, 0xe7, 0x8e, 0x33, 0xef,
	0xf6, 0x37, 0x6c, 0x9e, 0x3e, 0x1e, 0xdb, 0x0e, 0xb5, 0xac, 0xf9, 0x17, 0xf4, 0x4a, 0xf2, 0xc6,
	0x1a, 0x8e, 0xdf, 0x22, 0x0a, 0xd1, 0x58, 0x63, 0x6c, 0x1f, 0xc4, 0x6e, 0x2d, 0x62, 0x92, 0x68,
	0xad, 0x96, 0x7d, 0x2b, 0x94, 0x11, 0x06, 0x42, 0x7d, 0x1a, 0xe1, 0xfd, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x6c, 0xb6, 0x13, 0x66, 0x01, 0x00, 0x00,
}
