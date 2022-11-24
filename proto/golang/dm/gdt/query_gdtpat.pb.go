// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_gdtpat.proto

package gdt

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

type GetGdtpatRequest struct {
	Itnbr                string   `protobuf:"bytes,1,opt,name=itnbr,proto3" json:"itnbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGdtpatRequest) Reset()         { *m = GetGdtpatRequest{} }
func (m *GetGdtpatRequest) String() string { return proto.CompactTextString(m) }
func (*GetGdtpatRequest) ProtoMessage()    {}
func (*GetGdtpatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72ca446317f32c7e, []int{0}
}

func (m *GetGdtpatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGdtpatRequest.Unmarshal(m, b)
}
func (m *GetGdtpatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGdtpatRequest.Marshal(b, m, deterministic)
}
func (m *GetGdtpatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGdtpatRequest.Merge(m, src)
}
func (m *GetGdtpatRequest) XXX_Size() int {
	return xxx_messageInfo_GetGdtpatRequest.Size(m)
}
func (m *GetGdtpatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGdtpatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGdtpatRequest proto.InternalMessageInfo

func (m *GetGdtpatRequest) GetItnbr() string {
	if m != nil {
		return m.Itnbr
	}
	return ""
}

type GetGdtpatReply struct {
	Data                 []*Gdtpat `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetGdtpatReply) Reset()         { *m = GetGdtpatReply{} }
func (m *GetGdtpatReply) String() string { return proto.CompactTextString(m) }
func (*GetGdtpatReply) ProtoMessage()    {}
func (*GetGdtpatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_72ca446317f32c7e, []int{1}
}

func (m *GetGdtpatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGdtpatReply.Unmarshal(m, b)
}
func (m *GetGdtpatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGdtpatReply.Marshal(b, m, deterministic)
}
func (m *GetGdtpatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGdtpatReply.Merge(m, src)
}
func (m *GetGdtpatReply) XXX_Size() int {
	return xxx_messageInfo_GetGdtpatReply.Size(m)
}
func (m *GetGdtpatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGdtpatReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetGdtpatReply proto.InternalMessageInfo

func (m *GetGdtpatReply) GetData() []*Gdtpat {
	if m != nil {
		return m.Data
	}
	return nil
}

type Gdtpat struct {
	Itnbr                string   `protobuf:"bytes,1,opt,name=itnbr,proto3" json:"itnbr,omitempty"`
	Pallets              string   `protobuf:"bytes,2,opt,name=pallets,proto3" json:"pallets,omitempty"`
	Unit1                string   `protobuf:"bytes,3,opt,name=unit1,proto3" json:"unit1,omitempty"`
	Basewg               string   `protobuf:"bytes,4,opt,name=basewg,proto3" json:"basewg,omitempty"`
	Unit2                string   `protobuf:"bytes,5,opt,name=unit2,proto3" json:"unit2,omitempty"`
	Pallwg               string   `protobuf:"bytes,6,opt,name=pallwg,proto3" json:"pallwg,omitempty"`
	Indat                string   `protobuf:"bytes,7,opt,name=indat,proto3" json:"indat,omitempty"`
	Usrno                string   `protobuf:"bytes,8,opt,name=usrno,proto3" json:"usrno,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gdtpat) Reset()         { *m = Gdtpat{} }
func (m *Gdtpat) String() string { return proto.CompactTextString(m) }
func (*Gdtpat) ProtoMessage()    {}
func (*Gdtpat) Descriptor() ([]byte, []int) {
	return fileDescriptor_72ca446317f32c7e, []int{2}
}

func (m *Gdtpat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gdtpat.Unmarshal(m, b)
}
func (m *Gdtpat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gdtpat.Marshal(b, m, deterministic)
}
func (m *Gdtpat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gdtpat.Merge(m, src)
}
func (m *Gdtpat) XXX_Size() int {
	return xxx_messageInfo_Gdtpat.Size(m)
}
func (m *Gdtpat) XXX_DiscardUnknown() {
	xxx_messageInfo_Gdtpat.DiscardUnknown(m)
}

var xxx_messageInfo_Gdtpat proto.InternalMessageInfo

func (m *Gdtpat) GetItnbr() string {
	if m != nil {
		return m.Itnbr
	}
	return ""
}

func (m *Gdtpat) GetPallets() string {
	if m != nil {
		return m.Pallets
	}
	return ""
}

func (m *Gdtpat) GetUnit1() string {
	if m != nil {
		return m.Unit1
	}
	return ""
}

func (m *Gdtpat) GetBasewg() string {
	if m != nil {
		return m.Basewg
	}
	return ""
}

func (m *Gdtpat) GetUnit2() string {
	if m != nil {
		return m.Unit2
	}
	return ""
}

func (m *Gdtpat) GetPallwg() string {
	if m != nil {
		return m.Pallwg
	}
	return ""
}

func (m *Gdtpat) GetIndat() string {
	if m != nil {
		return m.Indat
	}
	return ""
}

func (m *Gdtpat) GetUsrno() string {
	if m != nil {
		return m.Usrno
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGdtpatRequest)(nil), "proto.dm.gdt.GetGdtpatRequest")
	proto.RegisterType((*GetGdtpatReply)(nil), "proto.dm.gdt.GetGdtpatReply")
	proto.RegisterType((*Gdtpat)(nil), "proto.dm.gdt.Gdtpat")
}

func init() { proto.RegisterFile("query_gdtpat.proto", fileDescriptor_72ca446317f32c7e) }

var fileDescriptor_72ca446317f32c7e = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x46, 0x71, 0x93, 0x38, 0xad, 0x5a, 0x4a, 0x11, 0xa1, 0xdc, 0x18, 0x3c, 0x79, 0x92, 0xa8,
	0xbb, 0x94, 0x8c, 0x5d, 0xb2, 0x67, 0xec, 0x52, 0xe4, 0x48, 0x08, 0x53, 0x5b, 0x72, 0xac, 0x73,
	0x4d, 0xfe, 0x5d, 0x7f, 0x5a, 0xd1, 0xc9, 0xa1, 0x19, 0x3a, 0x89, 0xf7, 0xf1, 0x78, 0x1c, 0x62,
	0xfc, 0x34, 0x9a, 0xe1, 0xfc, 0x69, 0x35, 0xf6, 0x0a, 0x45, 0x3f, 0x78, 0xf4, 0xfc, 0x81, 0x1e,
	0xa1, 0x3b, 0x61, 0x35, 0x16, 0x25, 0x7b, 0xda, 0x1b, 0xdc, 0x93, 0x70, 0x30, 0xa7, 0xd1, 0x04,
	0xe4, 0x1b, 0xb6, 0x6a, 0xd0, 0xd5, 0x03, 0x64, 0xdb, 0xac, 0xbc, 0x3b, 0x24, 0x28, 0x76, 0xec,
	0xf1, 0xca, 0xec, 0xdb, 0x33, 0x2f, 0xd9, 0x52, 0x2b, 0x54, 0x90, 0x6d, 0x17, 0xe5, 0x7d, 0xb5,
	0x11, 0xd7, 0x61, 0x31, 0x8b, 0x64, 0x14, 0x3f, 0x19, 0xcb, 0xd3, 0xf0, 0x7f, 0x9c, 0x03, 0x5b,
	0xf7, 0xaa, 0x6d, 0x0d, 0x06, 0xb8, 0xa1, 0xfd, 0x82, 0xd1, 0x1f, 0x5d, 0x83, 0x2f, 0xb0, 0x48,
	0x3e, 0x01, 0x7f, 0x66, 0x79, 0xad, 0x82, 0x99, 0x2c, 0x2c, 0x69, 0x9e, 0xe9, 0x62, 0x57, 0xb0,
	0xfa, 0xb3, 0xab, 0x68, 0xc7, 0xdc, 0x64, 0x21, 0x4f, 0x76, 0x22, 0xba, 0xc5, 0x69, 0x85, 0xb0,
	0x9e, 0x6f, 0x89, 0x40, 0x8d, 0x30, 0x38, 0x0f, 0xb7, 0x73, 0x23, 0xc2, 0xfb, 0xee, 0xe3, 0xcd,
	0x36, 0xd8, 0xaa, 0x5a, 0x7c, 0x19, 0xa7, 0x95, 0x38, 0xfa, 0x4e, 0xe0, 0x24, 0x09, 0xe4, 0xd1,
	0x77, 0x9d, 0x77, 0x41, 0x7e, 0x57, 0x92, 0x3e, 0x41, 0x5a, 0xdf, 0x2a, 0x67, 0xa5, 0xee, 0xa4,
	0xd5, 0x58, 0xe7, 0x34, 0xbe, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x32, 0x30, 0x2e, 0x58, 0x8f,
	0x01, 0x00, 0x00,
}
