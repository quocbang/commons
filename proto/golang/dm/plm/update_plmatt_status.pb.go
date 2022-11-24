// Code generated by protoc-gen-go. DO NOT EDIT.
// source: update_plmatt_status.proto

package plm

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

type UpdatePlmattStatusRequest struct {
	No                   int32    `protobuf:"varint,1,opt,name=no,proto3" json:"no,omitempty"`
	Seq                  int32    `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePlmattStatusRequest) Reset()         { *m = UpdatePlmattStatusRequest{} }
func (m *UpdatePlmattStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePlmattStatusRequest) ProtoMessage()    {}
func (*UpdatePlmattStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be6f755ded09ed3, []int{0}
}

func (m *UpdatePlmattStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePlmattStatusRequest.Unmarshal(m, b)
}
func (m *UpdatePlmattStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePlmattStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePlmattStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePlmattStatusRequest.Merge(m, src)
}
func (m *UpdatePlmattStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePlmattStatusRequest.Size(m)
}
func (m *UpdatePlmattStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePlmattStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePlmattStatusRequest proto.InternalMessageInfo

func (m *UpdatePlmattStatusRequest) GetNo() int32 {
	if m != nil {
		return m.No
	}
	return 0
}

func (m *UpdatePlmattStatusRequest) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *UpdatePlmattStatusRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type UpdatePlmattStatusReply struct {
	RowsAffect           int64    `protobuf:"varint,1,opt,name=rows_affect,json=rowsAffect,proto3" json:"rows_affect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePlmattStatusReply) Reset()         { *m = UpdatePlmattStatusReply{} }
func (m *UpdatePlmattStatusReply) String() string { return proto.CompactTextString(m) }
func (*UpdatePlmattStatusReply) ProtoMessage()    {}
func (*UpdatePlmattStatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be6f755ded09ed3, []int{1}
}

func (m *UpdatePlmattStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePlmattStatusReply.Unmarshal(m, b)
}
func (m *UpdatePlmattStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePlmattStatusReply.Marshal(b, m, deterministic)
}
func (m *UpdatePlmattStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePlmattStatusReply.Merge(m, src)
}
func (m *UpdatePlmattStatusReply) XXX_Size() int {
	return xxx_messageInfo_UpdatePlmattStatusReply.Size(m)
}
func (m *UpdatePlmattStatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePlmattStatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePlmattStatusReply proto.InternalMessageInfo

func (m *UpdatePlmattStatusReply) GetRowsAffect() int64 {
	if m != nil {
		return m.RowsAffect
	}
	return 0
}

func init() {
	proto.RegisterType((*UpdatePlmattStatusRequest)(nil), "proto.dm.plm.UpdatePlmattStatusRequest")
	proto.RegisterType((*UpdatePlmattStatusReply)(nil), "proto.dm.plm.UpdatePlmattStatusReply")
}

func init() { proto.RegisterFile("update_plmatt_status.proto", fileDescriptor_1be6f755ded09ed3) }

var fileDescriptor_1be6f755ded09ed3 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x46, 0xd9, 0x5d, 0x2c, 0x38, 0x8a, 0x48, 0x0e, 0xba, 0x7a, 0xb1, 0xf4, 0xd4, 0x53, 0x02,
	0x7a, 0x91, 0xde, 0xf4, 0x17, 0x48, 0xa4, 0x17, 0x2f, 0x4b, 0xba, 0x49, 0x17, 0x31, 0x93, 0x49,
	0x9b, 0x59, 0x17, 0xff, 0xbd, 0xec, 0xe8, 0xb1, 0xa7, 0xe4, 0x7b, 0x3c, 0x1e, 0x0c, 0xdc, 0x8f,
	0xd9, 0x3b, 0x0e, 0x5d, 0x8e, 0xe8, 0x98, 0xbb, 0xc2, 0x8e, 0xc7, 0xa2, 0xf3, 0x91, 0x98, 0xd4,
	0xa5, 0x3c, 0xda, 0xa3, 0xce, 0x11, 0x57, 0x5b, 0xb8, 0xdb, 0x8a, 0xfb, 0x26, 0xea, 0xbb, 0x98,
	0x36, 0x1c, 0xc6, 0x50, 0x58, 0x5d, 0x41, 0x9d, 0xa8, 0xad, 0x96, 0xd5, 0xfa, 0xcc, 0xd6, 0x89,
	0xd4, 0x35, 0x34, 0x25, 0x1c, 0xda, 0x5a, 0xc0, 0xfc, 0x55, 0x37, 0xb0, 0xf8, 0x8b, 0xb7, 0xcd,
	0xb2, 0x5a, 0x9f, 0xdb, 0xff, 0xb5, 0xda, 0xc0, 0xed, 0xa9, 0x6c, 0x8e, 0x3f, 0xea, 0x01, 0x2e,
	0x8e, 0x34, 0x95, 0xce, 0xed, 0xf7, 0xa1, 0x67, 0xa9, 0x37, 0x16, 0x66, 0xf4, 0x22, 0xe4, 0x75,
	0xf3, 0xf1, 0x3c, 0x7c, 0x72, 0x74, 0x3b, 0xfd, 0x15, 0x92, 0x77, 0xba, 0x27, 0xd4, 0x3c, 0x19,
	0x19, 0xa6, 0x27, 0x44, 0x4a, 0xc5, 0x7c, 0x3f, 0x1a, 0xb9, 0xc3, 0x0c, 0x14, 0x5d, 0x1a, 0x8c,
	0x47, 0x93, 0x23, 0xee, 0x16, 0x02, 0x9f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x32, 0x5d, 0xd7,
	0x1d, 0x01, 0x01, 0x00, 0x00,
}
