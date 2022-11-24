// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_prd_man.proto

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

type PrdManRequest struct {
	Subno                string   `protobuf:"bytes,1,opt,name=subno,proto3" json:"subno,omitempty"`
	Factory              string   `protobuf:"bytes,2,opt,name=factory,proto3" json:"factory,omitempty"`
	Depno                string   `protobuf:"bytes,3,opt,name=depno,proto3" json:"depno,omitempty"`
	Pday                 string   `protobuf:"bytes,4,opt,name=pday,proto3" json:"pday,omitempty"`
	Class                string   `protobuf:"bytes,5,opt,name=class,proto3" json:"class,omitempty"`
	Empno                string   `protobuf:"bytes,6,opt,name=empno,proto3" json:"empno,omitempty"`
	Code                 string   `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	Codd                 string   `protobuf:"bytes,8,opt,name=codd,proto3" json:"codd,omitempty"`
	Partno               string   `protobuf:"bytes,9,opt,name=partno,proto3" json:"partno,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrdManRequest) Reset()         { *m = PrdManRequest{} }
func (m *PrdManRequest) String() string { return proto.CompactTextString(m) }
func (*PrdManRequest) ProtoMessage()    {}
func (*PrdManRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfeebd03a900ab13, []int{0}
}

func (m *PrdManRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrdManRequest.Unmarshal(m, b)
}
func (m *PrdManRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrdManRequest.Marshal(b, m, deterministic)
}
func (m *PrdManRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrdManRequest.Merge(m, src)
}
func (m *PrdManRequest) XXX_Size() int {
	return xxx_messageInfo_PrdManRequest.Size(m)
}
func (m *PrdManRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrdManRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrdManRequest proto.InternalMessageInfo

func (m *PrdManRequest) GetSubno() string {
	if m != nil {
		return m.Subno
	}
	return ""
}

func (m *PrdManRequest) GetFactory() string {
	if m != nil {
		return m.Factory
	}
	return ""
}

func (m *PrdManRequest) GetDepno() string {
	if m != nil {
		return m.Depno
	}
	return ""
}

func (m *PrdManRequest) GetPday() string {
	if m != nil {
		return m.Pday
	}
	return ""
}

func (m *PrdManRequest) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *PrdManRequest) GetEmpno() string {
	if m != nil {
		return m.Empno
	}
	return ""
}

func (m *PrdManRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PrdManRequest) GetCodd() string {
	if m != nil {
		return m.Codd
	}
	return ""
}

func (m *PrdManRequest) GetPartno() string {
	if m != nil {
		return m.Partno
	}
	return ""
}

type PrdManReply struct {
	Lists                []*PrdManReply_List `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PrdManReply) Reset()         { *m = PrdManReply{} }
func (m *PrdManReply) String() string { return proto.CompactTextString(m) }
func (*PrdManReply) ProtoMessage()    {}
func (*PrdManReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfeebd03a900ab13, []int{1}
}

func (m *PrdManReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrdManReply.Unmarshal(m, b)
}
func (m *PrdManReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrdManReply.Marshal(b, m, deterministic)
}
func (m *PrdManReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrdManReply.Merge(m, src)
}
func (m *PrdManReply) XXX_Size() int {
	return xxx_messageInfo_PrdManReply.Size(m)
}
func (m *PrdManReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PrdManReply.DiscardUnknown(m)
}

var xxx_messageInfo_PrdManReply proto.InternalMessageInfo

func (m *PrdManReply) GetLists() []*PrdManReply_List {
	if m != nil {
		return m.Lists
	}
	return nil
}

type PrdManReply_List struct {
	Subno                string   `protobuf:"bytes,1,opt,name=subno,proto3" json:"subno,omitempty"`
	Factory              string   `protobuf:"bytes,2,opt,name=factory,proto3" json:"factory,omitempty"`
	Depno                string   `protobuf:"bytes,3,opt,name=depno,proto3" json:"depno,omitempty"`
	Pday                 string   `protobuf:"bytes,4,opt,name=pday,proto3" json:"pday,omitempty"`
	Class                string   `protobuf:"bytes,5,opt,name=class,proto3" json:"class,omitempty"`
	Empno                string   `protobuf:"bytes,6,opt,name=empno,proto3" json:"empno,omitempty"`
	Code                 string   `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	Codd                 string   `protobuf:"bytes,8,opt,name=codd,proto3" json:"codd,omitempty"`
	Manuno               string   `protobuf:"bytes,9,opt,name=manuno,proto3" json:"manuno,omitempty"`
	Unit                 string   `protobuf:"bytes,10,opt,name=unit,proto3" json:"unit,omitempty"`
	Mano                 string   `protobuf:"bytes,11,opt,name=mano,proto3" json:"mano,omitempty"`
	Partno               string   `protobuf:"bytes,12,opt,name=partno,proto3" json:"partno,omitempty"`
	Stype                string   `protobuf:"bytes,13,opt,name=stype,proto3" json:"stype,omitempty"`
	Preanum              string   `protobuf:"bytes,14,opt,name=preanum,proto3" json:"preanum,omitempty"`
	Actanum              string   `protobuf:"bytes,15,opt,name=actanum,proto3" json:"actanum,omitempty"`
	Godanum              string   `protobuf:"bytes,16,opt,name=godanum,proto3" json:"godanum,omitempty"`
	Disanum              string   `protobuf:"bytes,17,opt,name=disanum,proto3" json:"disanum,omitempty"`
	Recanum              string   `protobuf:"bytes,18,opt,name=recanum,proto3" json:"recanum,omitempty"`
	Cyctim               string   `protobuf:"bytes,19,opt,name=cyctim,proto3" json:"cyctim,omitempty"`
	Biotim               string   `protobuf:"bytes,20,opt,name=biotim,proto3" json:"biotim,omitempty"`
	Rttim                string   `protobuf:"bytes,21,opt,name=rttim,proto3" json:"rttim,omitempty"`
	Rate                 string   `protobuf:"bytes,22,opt,name=rate,proto3" json:"rate,omitempty"`
	Ptype                string   `protobuf:"bytes,23,opt,name=ptype,proto3" json:"ptype,omitempty"`
	Machnd               string   `protobuf:"bytes,24,opt,name=machnd,proto3" json:"machnd,omitempty"`
	Indat                string   `protobuf:"bytes,25,opt,name=indat,proto3" json:"indat,omitempty"`
	Usrno                string   `protobuf:"bytes,26,opt,name=usrno,proto3" json:"usrno,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrdManReply_List) Reset()         { *m = PrdManReply_List{} }
func (m *PrdManReply_List) String() string { return proto.CompactTextString(m) }
func (*PrdManReply_List) ProtoMessage()    {}
func (*PrdManReply_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfeebd03a900ab13, []int{1, 0}
}

func (m *PrdManReply_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrdManReply_List.Unmarshal(m, b)
}
func (m *PrdManReply_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrdManReply_List.Marshal(b, m, deterministic)
}
func (m *PrdManReply_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrdManReply_List.Merge(m, src)
}
func (m *PrdManReply_List) XXX_Size() int {
	return xxx_messageInfo_PrdManReply_List.Size(m)
}
func (m *PrdManReply_List) XXX_DiscardUnknown() {
	xxx_messageInfo_PrdManReply_List.DiscardUnknown(m)
}

var xxx_messageInfo_PrdManReply_List proto.InternalMessageInfo

func (m *PrdManReply_List) GetSubno() string {
	if m != nil {
		return m.Subno
	}
	return ""
}

func (m *PrdManReply_List) GetFactory() string {
	if m != nil {
		return m.Factory
	}
	return ""
}

func (m *PrdManReply_List) GetDepno() string {
	if m != nil {
		return m.Depno
	}
	return ""
}

func (m *PrdManReply_List) GetPday() string {
	if m != nil {
		return m.Pday
	}
	return ""
}

func (m *PrdManReply_List) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *PrdManReply_List) GetEmpno() string {
	if m != nil {
		return m.Empno
	}
	return ""
}

func (m *PrdManReply_List) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PrdManReply_List) GetCodd() string {
	if m != nil {
		return m.Codd
	}
	return ""
}

func (m *PrdManReply_List) GetManuno() string {
	if m != nil {
		return m.Manuno
	}
	return ""
}

func (m *PrdManReply_List) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *PrdManReply_List) GetMano() string {
	if m != nil {
		return m.Mano
	}
	return ""
}

func (m *PrdManReply_List) GetPartno() string {
	if m != nil {
		return m.Partno
	}
	return ""
}

func (m *PrdManReply_List) GetStype() string {
	if m != nil {
		return m.Stype
	}
	return ""
}

func (m *PrdManReply_List) GetPreanum() string {
	if m != nil {
		return m.Preanum
	}
	return ""
}

func (m *PrdManReply_List) GetActanum() string {
	if m != nil {
		return m.Actanum
	}
	return ""
}

func (m *PrdManReply_List) GetGodanum() string {
	if m != nil {
		return m.Godanum
	}
	return ""
}

func (m *PrdManReply_List) GetDisanum() string {
	if m != nil {
		return m.Disanum
	}
	return ""
}

func (m *PrdManReply_List) GetRecanum() string {
	if m != nil {
		return m.Recanum
	}
	return ""
}

func (m *PrdManReply_List) GetCyctim() string {
	if m != nil {
		return m.Cyctim
	}
	return ""
}

func (m *PrdManReply_List) GetBiotim() string {
	if m != nil {
		return m.Biotim
	}
	return ""
}

func (m *PrdManReply_List) GetRttim() string {
	if m != nil {
		return m.Rttim
	}
	return ""
}

func (m *PrdManReply_List) GetRate() string {
	if m != nil {
		return m.Rate
	}
	return ""
}

func (m *PrdManReply_List) GetPtype() string {
	if m != nil {
		return m.Ptype
	}
	return ""
}

func (m *PrdManReply_List) GetMachnd() string {
	if m != nil {
		return m.Machnd
	}
	return ""
}

func (m *PrdManReply_List) GetIndat() string {
	if m != nil {
		return m.Indat
	}
	return ""
}

func (m *PrdManReply_List) GetUsrno() string {
	if m != nil {
		return m.Usrno
	}
	return ""
}

func init() {
	proto.RegisterType((*PrdManRequest)(nil), "proto.dm.prd.PrdManRequest")
	proto.RegisterType((*PrdManReply)(nil), "proto.dm.prd.PrdManReply")
	proto.RegisterType((*PrdManReply_List)(nil), "proto.dm.prd.PrdManReply.List")
}

func init() { proto.RegisterFile("query_prd_man.proto", fileDescriptor_bfeebd03a900ab13) }

var fileDescriptor_bfeebd03a900ab13 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x92, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x15, 0x9a, 0xa4, 0xd4, 0x69, 0xf9, 0x99, 0x96, 0x72, 0xe9, 0x02, 0x55, 0x5d, 0x75,
	0xe5, 0x91, 0x0a, 0x0b, 0xc4, 0x92, 0x35, 0x48, 0x28, 0x4b, 0x36, 0x95, 0xe3, 0x6b, 0xc2, 0x88,
	0xf1, 0x4f, 0x6c, 0x0f, 0x68, 0x5e, 0x86, 0xb7, 0x42, 0xe2, 0x71, 0xd0, 0xbd, 0x77, 0x02, 0x79,
	0x85, 0xae, 0xe2, 0xef, 0x3b, 0xc7, 0xd2, 0xf8, 0x28, 0xea, 0x7c, 0x37, 0xb8, 0x3c, 0xde, 0xa7,
	0x8c, 0xf7, 0xde, 0x04, 0x9d, 0x72, 0xac, 0xb1, 0x39, 0xe5, 0x1f, 0x8d, 0x5e, 0xa7, 0x8c, 0x37,
	0xbf, 0x67, 0xea, 0xec, 0x73, 0xc6, 0x4f, 0x26, 0xac, 0xdd, 0x6e, 0x70, 0xa5, 0x36, 0x17, 0x6a,
	0x51, 0x86, 0x4d, 0x88, 0x30, 0xbb, 0x9e, 0xdd, 0x9e, 0xac, 0x05, 0x1a, 0x50, 0xc7, 0x5f, 0x8d,
	0xad, 0x31, 0x8f, 0xf0, 0x88, 0xfd, 0x1e, 0xa9, 0x8f, 0x2e, 0x85, 0x08, 0x47, 0xd2, 0x67, 0x68,
	0x1a, 0x35, 0x4f, 0x68, 0x46, 0x98, 0xb3, 0xe4, 0x33, 0x35, 0x6d, 0x6f, 0x4a, 0x81, 0x85, 0x34,
	0x19, 0xc8, 0x3a, 0x4f, 0xf7, 0x97, 0x62, 0x19, 0xe8, 0xbe, 0x8d, 0xe8, 0xe0, 0x58, 0xee, 0xd3,
	0x79, 0x72, 0x08, 0x8f, 0xff, 0x39, 0x6c, 0x2e, 0xd5, 0x32, 0x99, 0x5c, 0x43, 0x84, 0x13, 0xb6,
	0x13, 0xdd, 0xfc, 0x5a, 0xa8, 0xd5, 0xfe, 0x5d, 0xa9, 0x1f, 0x9b, 0xb7, 0x6a, 0xd1, 0x77, 0xa5,
	0x16, 0x98, 0x5d, 0x1f, 0xdd, 0xae, 0xee, 0x5e, 0xeb, 0xc3, 0x15, 0xf4, 0x41, 0x53, 0x7f, 0xec,
	0x4a, 0x5d, 0x4b, 0xf9, 0xea, 0xcf, 0x5c, 0xcd, 0x89, 0x1f, 0xde, 0x28, 0xde, 0x84, 0xe1, 0xff,
	0x28, 0x42, 0xd4, 0x1d, 0x42, 0x57, 0x41, 0x49, 0x97, 0xce, 0xe4, 0xbc, 0x09, 0x11, 0x56, 0xe2,
	0xe8, 0x7c, 0x30, 0xea, 0xe9, 0xe1, 0xa8, 0xbc, 0x42, 0x1d, 0x93, 0x83, 0xb3, 0x69, 0x05, 0x02,
	0x5a, 0x21, 0x65, 0x67, 0xc2, 0xe0, 0xe1, 0x89, 0xac, 0x30, 0x21, 0x25, 0xc6, 0x56, 0x4e, 0x9e,
	0x4a, 0x32, 0x21, 0x25, 0xdb, 0x88, 0x9c, 0x3c, 0x93, 0x64, 0x42, 0x4a, 0xb0, 0x2b, 0x9c, 0x3c,
	0x97, 0x64, 0x42, 0x4a, 0xb2, 0xb3, 0x9c, 0x34, 0x92, 0x4c, 0x48, 0xdf, 0x6b, 0x47, 0x5b, 0x3b,
	0x0f, 0xe7, 0xf2, 0xbd, 0x42, 0xe4, 0x37, 0x5d, 0x24, 0x7f, 0x21, 0x5e, 0x88, 0xde, 0x91, 0x2b,
	0xe9, 0x17, 0xf2, 0x0e, 0x06, 0x5a, 0x22, 0x9b, 0xea, 0xe0, 0x52, 0x96, 0xa0, 0x33, 0x35, 0x13,
	0xbf, 0xf8, 0xa5, 0x34, 0x19, 0x64, 0x5f, 0xfb, 0x2d, 0x20, 0xc0, 0x7e, 0x5f, 0x22, 0x6a, 0x77,
	0x01, 0x4d, 0x85, 0x57, 0xd2, 0x66, 0x20, 0x3b, 0x94, 0x1c, 0x22, 0x5c, 0x89, 0x65, 0xf8, 0xf0,
	0xfe, 0xcb, 0xbb, 0x6d, 0x57, 0x7b, 0xb3, 0xd1, 0xdf, 0x5d, 0x40, 0xa3, 0x6d, 0xf4, 0xba, 0xfe,
	0x6c, 0x19, 0x5a, 0x1b, 0xbd, 0x8f, 0xa1, 0xb4, 0x3f, 0xee, 0x5a, 0xfe, 0x9f, 0xb6, 0xdb, 0xd8,
	0x9b, 0xb0, 0x6d, 0xd1, 0xb7, 0x29, 0xe3, 0x66, 0xc9, 0xf2, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x03, 0x9c, 0x89, 0x88, 0xe0, 0x03, 0x00, 0x00,
}