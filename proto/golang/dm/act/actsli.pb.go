// Code generated by protoc-gen-go. DO NOT EDIT.
// source: actsli.proto

package act

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//actsli
type KeepAccountRequest struct {
	Subno                string   `protobuf:"bytes,1,opt,name=subno,proto3" json:"subno,omitempty"`
	Acctno               string   `protobuf:"bytes,2,opt,name=acctno,proto3" json:"acctno,omitempty"`
	StartAccdat          string   `protobuf:"bytes,3,opt,name=start_accdat,json=startAccdat,proto3" json:"start_accdat,omitempty"`
	EndAccdat            string   `protobuf:"bytes,4,opt,name=end_accdat,json=endAccdat,proto3" json:"end_accdat,omitempty"`
	Curr                 string   `protobuf:"bytes,5,opt,name=curr,proto3" json:"curr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeepAccountRequest) Reset()         { *m = KeepAccountRequest{} }
func (m *KeepAccountRequest) String() string { return proto.CompactTextString(m) }
func (*KeepAccountRequest) ProtoMessage()    {}
func (*KeepAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb453a1ab6d65b4, []int{0}
}

func (m *KeepAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeepAccountRequest.Unmarshal(m, b)
}
func (m *KeepAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeepAccountRequest.Marshal(b, m, deterministic)
}
func (m *KeepAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeepAccountRequest.Merge(m, src)
}
func (m *KeepAccountRequest) XXX_Size() int {
	return xxx_messageInfo_KeepAccountRequest.Size(m)
}
func (m *KeepAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeepAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeepAccountRequest proto.InternalMessageInfo

func (m *KeepAccountRequest) GetSubno() string {
	if m != nil {
		return m.Subno
	}
	return ""
}

func (m *KeepAccountRequest) GetAcctno() string {
	if m != nil {
		return m.Acctno
	}
	return ""
}

func (m *KeepAccountRequest) GetStartAccdat() string {
	if m != nil {
		return m.StartAccdat
	}
	return ""
}

func (m *KeepAccountRequest) GetEndAccdat() string {
	if m != nil {
		return m.EndAccdat
	}
	return ""
}

func (m *KeepAccountRequest) GetCurr() string {
	if m != nil {
		return m.Curr
	}
	return ""
}

//??????????????????
type SingleKeepAccount struct {
	Subno                string               `protobuf:"bytes,1,opt,name=subno,proto3" json:"subno,omitempty"`
	Acctno               string               `protobuf:"bytes,2,opt,name=acctno,proto3" json:"acctno,omitempty"`
	Accdat               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=accdat,proto3" json:"accdat,omitempty"`
	Curr                 string               `protobuf:"bytes,4,opt,name=curr,proto3" json:"curr,omitempty"`
	Rate                 float32              `protobuf:"fixed32,5,opt,name=rate,proto3" json:"rate,omitempty"`
	Orgamt               float32              `protobuf:"fixed32,6,opt,name=orgamt,proto3" json:"orgamt,omitempty"`
	Peoamt               float32              `protobuf:"fixed32,7,opt,name=peoamt,proto3" json:"peoamt,omitempty"`
	Reference            string               `protobuf:"bytes,8,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SingleKeepAccount) Reset()         { *m = SingleKeepAccount{} }
func (m *SingleKeepAccount) String() string { return proto.CompactTextString(m) }
func (*SingleKeepAccount) ProtoMessage()    {}
func (*SingleKeepAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb453a1ab6d65b4, []int{1}
}

func (m *SingleKeepAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleKeepAccount.Unmarshal(m, b)
}
func (m *SingleKeepAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleKeepAccount.Marshal(b, m, deterministic)
}
func (m *SingleKeepAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleKeepAccount.Merge(m, src)
}
func (m *SingleKeepAccount) XXX_Size() int {
	return xxx_messageInfo_SingleKeepAccount.Size(m)
}
func (m *SingleKeepAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleKeepAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SingleKeepAccount proto.InternalMessageInfo

func (m *SingleKeepAccount) GetSubno() string {
	if m != nil {
		return m.Subno
	}
	return ""
}

func (m *SingleKeepAccount) GetAcctno() string {
	if m != nil {
		return m.Acctno
	}
	return ""
}

func (m *SingleKeepAccount) GetAccdat() *timestamp.Timestamp {
	if m != nil {
		return m.Accdat
	}
	return nil
}

func (m *SingleKeepAccount) GetCurr() string {
	if m != nil {
		return m.Curr
	}
	return ""
}

func (m *SingleKeepAccount) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *SingleKeepAccount) GetOrgamt() float32 {
	if m != nil {
		return m.Orgamt
	}
	return 0
}

func (m *SingleKeepAccount) GetPeoamt() float32 {
	if m != nil {
		return m.Peoamt
	}
	return 0
}

func (m *SingleKeepAccount) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

//??????????????????
type MultiKeepAccount struct {
	Dataset              []*SingleKeepAccount `protobuf:"bytes,1,rep,name=dataset,proto3" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MultiKeepAccount) Reset()         { *m = MultiKeepAccount{} }
func (m *MultiKeepAccount) String() string { return proto.CompactTextString(m) }
func (*MultiKeepAccount) ProtoMessage()    {}
func (*MultiKeepAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb453a1ab6d65b4, []int{2}
}

func (m *MultiKeepAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiKeepAccount.Unmarshal(m, b)
}
func (m *MultiKeepAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiKeepAccount.Marshal(b, m, deterministic)
}
func (m *MultiKeepAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiKeepAccount.Merge(m, src)
}
func (m *MultiKeepAccount) XXX_Size() int {
	return xxx_messageInfo_MultiKeepAccount.Size(m)
}
func (m *MultiKeepAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiKeepAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MultiKeepAccount proto.InternalMessageInfo

func (m *MultiKeepAccount) GetDataset() []*SingleKeepAccount {
	if m != nil {
		return m.Dataset
	}
	return nil
}

func init() {
	proto.RegisterType((*KeepAccountRequest)(nil), "proto.dm.act.KeepAccountRequest")
	proto.RegisterType((*SingleKeepAccount)(nil), "proto.dm.act.SingleKeepAccount")
	proto.RegisterType((*MultiKeepAccount)(nil), "proto.dm.act.MultiKeepAccount")
}

func init() { proto.RegisterFile("actsli.proto", fileDescriptor_9bb453a1ab6d65b4) }

var fileDescriptor_9bb453a1ab6d65b4 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3f, 0x4f, 0xeb, 0x30,
	0x14, 0xc5, 0x95, 0xfe, 0x7d, 0x75, 0x3b, 0xbc, 0x67, 0x3d, 0xa1, 0xa8, 0x02, 0xb5, 0x74, 0xea,
	0x64, 0x4b, 0x61, 0x01, 0xb6, 0xb2, 0xa2, 0x2e, 0x81, 0x89, 0x05, 0x39, 0xce, 0x6d, 0x14, 0x11,
	0xdb, 0xc1, 0xb9, 0x81, 0xaf, 0xc2, 0x77, 0xe4, 0x4b, 0xa0, 0xdc, 0x24, 0xa5, 0x12, 0x13, 0x53,
	0xee, 0xf9, 0xdd, 0x23, 0xe5, 0xf8, 0xd8, 0x6c, 0xa1, 0x34, 0x56, 0x45, 0x2e, 0x4a, 0xef, 0xd0,
	0xf1, 0x05, 0x7d, 0x44, 0x6a, 0x84, 0xd2, 0xb8, 0x5c, 0x65, 0xce, 0x65, 0x05, 0x48, 0x82, 0x49,
	0x7d, 0x90, 0x98, 0x1b, 0xa8, 0x50, 0x99, 0xb2, 0xb5, 0x6f, 0x3e, 0x02, 0xc6, 0xef, 0x01, 0xca,
	0x9d, 0xd6, 0xae, 0xb6, 0x18, 0xc3, 0x6b, 0x0d, 0x15, 0xf2, 0xff, 0x6c, 0x5c, 0xd5, 0x89, 0x75,
	0x61, 0xb0, 0x0e, 0xb6, 0xb3, 0xb8, 0x15, 0xfc, 0x8c, 0x4d, 0x94, 0xd6, 0x68, 0x5d, 0x38, 0x20,
	0xdc, 0x29, 0x7e, 0xc9, 0x16, 0x15, 0x2a, 0x8f, 0xcf, 0x4a, 0xeb, 0x54, 0x61, 0x38, 0xa4, 0xed,
	0x9c, 0xd8, 0x8e, 0x10, 0xbf, 0x60, 0x0c, 0x6c, 0xda, 0x1b, 0x46, 0x64, 0x98, 0x81, 0x4d, 0xbb,
	0x35, 0x67, 0x23, 0x5d, 0x7b, 0x1f, 0x8e, 0x69, 0x41, 0xf3, 0xe6, 0x33, 0x60, 0xff, 0x1e, 0x72,
	0x9b, 0x15, 0x70, 0x12, 0xf0, 0x97, 0xc9, 0x22, 0xe2, 0x7d, 0xa6, 0x79, 0xb4, 0x14, 0x6d, 0x21,
	0xa2, 0x2f, 0x44, 0x3c, 0xf6, 0x85, 0xc4, 0x9d, 0xf3, 0x98, 0x65, 0xf4, 0x9d, 0xa5, 0x61, 0x5e,
	0x21, 0x50, 0xbe, 0x41, 0x4c, 0x73, 0xf3, 0x4f, 0xe7, 0x33, 0x65, 0x30, 0x9c, 0x10, 0xed, 0x54,
	0xc3, 0x4b, 0x70, 0x0d, 0x9f, 0xb6, 0xbc, 0x55, 0xfc, 0x9c, 0xcd, 0x3c, 0x1c, 0xc0, 0x83, 0xd5,
	0x10, 0xfe, 0x69, 0x1b, 0x38, 0x82, 0xcd, 0x9e, 0xfd, 0xdd, 0xd7, 0x05, 0xe6, 0xa7, 0x67, 0xbd,
	0x61, 0xd3, 0x54, 0xa1, 0xaa, 0x00, 0xc3, 0x60, 0x3d, 0xdc, 0xce, 0xa3, 0x95, 0x38, 0xbd, 0x5d,
	0xf1, 0xa3, 0x9d, 0xb8, 0xf7, 0xdf, 0xdd, 0x3e, 0x5d, 0x67, 0x39, 0x16, 0x2a, 0x11, 0x2f, 0x60,
	0x53, 0x25, 0xb4, 0x33, 0x02, 0xdf, 0x25, 0x09, 0xa9, 0x9d, 0x31, 0xce, 0x56, 0xf2, 0x2d, 0x6a,
	0x1f, 0x86, 0xcc, 0x5c, 0xa1, 0x6c, 0x26, 0x53, 0x23, 0x95, 0xc6, 0x64, 0x42, 0xf0, 0xea, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0x8e, 0x88, 0x1b, 0x59, 0x02, 0x00, 0x00,
}
