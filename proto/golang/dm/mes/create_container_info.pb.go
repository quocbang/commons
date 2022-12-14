// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_container_info.proto

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

// Create Container Info
type AddContainerRequest struct {
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	OwnedFtyId           int32    `protobuf:"varint,2,opt,name=owned_fty_id,json=ownedFtyId,proto3" json:"owned_fty_id,omitempty"`
	OwnedDeptId          string   `protobuf:"bytes,3,opt,name=owned_dept_id,json=ownedDeptId,proto3" json:"owned_dept_id,omitempty"`
	ConStat              string   `protobuf:"bytes,4,opt,name=con_stat,json=conStat,proto3" json:"con_stat,omitempty"`
	ConMtrl              string   `protobuf:"bytes,5,opt,name=con_mtrl,json=conMtrl,proto3" json:"con_mtrl,omitempty"`
	ConMtrlLevel         string   `protobuf:"bytes,6,opt,name=con_mtrl_level,json=conMtrlLevel,proto3" json:"con_mtrl_level,omitempty"`
	UpdateBy             string   `protobuf:"bytes,7,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddContainerRequest) Reset()         { *m = AddContainerRequest{} }
func (m *AddContainerRequest) String() string { return proto.CompactTextString(m) }
func (*AddContainerRequest) ProtoMessage()    {}
func (*AddContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_626488189d32c541, []int{0}
}

func (m *AddContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddContainerRequest.Unmarshal(m, b)
}
func (m *AddContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddContainerRequest.Marshal(b, m, deterministic)
}
func (m *AddContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddContainerRequest.Merge(m, src)
}
func (m *AddContainerRequest) XXX_Size() int {
	return xxx_messageInfo_AddContainerRequest.Size(m)
}
func (m *AddContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddContainerRequest proto.InternalMessageInfo

func (m *AddContainerRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *AddContainerRequest) GetOwnedFtyId() int32 {
	if m != nil {
		return m.OwnedFtyId
	}
	return 0
}

func (m *AddContainerRequest) GetOwnedDeptId() string {
	if m != nil {
		return m.OwnedDeptId
	}
	return ""
}

func (m *AddContainerRequest) GetConStat() string {
	if m != nil {
		return m.ConStat
	}
	return ""
}

func (m *AddContainerRequest) GetConMtrl() string {
	if m != nil {
		return m.ConMtrl
	}
	return ""
}

func (m *AddContainerRequest) GetConMtrlLevel() string {
	if m != nil {
		return m.ConMtrlLevel
	}
	return ""
}

func (m *AddContainerRequest) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

type AddContainerReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddContainerReply) Reset()         { *m = AddContainerReply{} }
func (m *AddContainerReply) String() string { return proto.CompactTextString(m) }
func (*AddContainerReply) ProtoMessage()    {}
func (*AddContainerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_626488189d32c541, []int{1}
}

func (m *AddContainerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddContainerReply.Unmarshal(m, b)
}
func (m *AddContainerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddContainerReply.Marshal(b, m, deterministic)
}
func (m *AddContainerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddContainerReply.Merge(m, src)
}
func (m *AddContainerReply) XXX_Size() int {
	return xxx_messageInfo_AddContainerReply.Size(m)
}
func (m *AddContainerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddContainerReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddContainerReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddContainerRequest)(nil), "proto.dm.mes.AddContainerRequest")
	proto.RegisterType((*AddContainerReply)(nil), "proto.dm.mes.AddContainerReply")
}

func init() { proto.RegisterFile("create_container_info.proto", fileDescriptor_626488189d32c541) }

var fileDescriptor_626488189d32c541 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4f, 0xfb, 0x30,
	0x10, 0x40, 0x95, 0xdf, 0x8f, 0x7e, 0x99, 0x80, 0x84, 0xbb, 0x04, 0x75, 0x29, 0x15, 0x43, 0x27,
	0x5b, 0x82, 0x05, 0xb1, 0x51, 0x10, 0x52, 0x25, 0x58, 0xca, 0xc6, 0x12, 0x39, 0xbe, 0x6b, 0x14,
	0xe1, 0x8f, 0x90, 0x5c, 0x5b, 0xe5, 0x2f, 0x67, 0x45, 0x71, 0x48, 0x05, 0x93, 0xe5, 0xf7, 0x9e,
	0x2d, 0xdd, 0xb1, 0x99, 0xae, 0x50, 0x11, 0xa6, 0xda, 0x3b, 0x52, 0x85, 0xc3, 0x2a, 0x2d, 0xdc,
	0xd6, 0x8b, 0xb2, 0xf2, 0xe4, 0x79, 0x1c, 0x0e, 0x01, 0x56, 0x58, 0xac, 0x17, 0x5f, 0x11, 0x9b,
	0x3e, 0x00, 0x3c, 0xf6, 0xe5, 0x06, 0x3f, 0x77, 0x58, 0x13, 0xbf, 0x62, 0xf1, 0xaf, 0xd7, 0x90,
	0x44, 0xf3, 0x68, 0x39, 0xd9, 0x9c, 0x1e, 0xd9, 0x1a, 0xf8, 0x9c, 0xc5, 0xfe, 0xe0, 0x10, 0xd2,
	0x2d, 0x35, 0x6d, 0xf2, 0x6f, 0x1e, 0x2d, 0x07, 0x1b, 0x16, 0xd8, 0x33, 0x35, 0x6b, 0xe0, 0x0b,
	0x76, 0xd6, 0x15, 0x80, 0x25, 0xb5, 0xc9, 0xff, 0xee, 0x97, 0x00, 0x9f, 0xb0, 0xa4, 0x35, 0xf0,
	0x4b, 0x36, 0xd6, 0xde, 0xa5, 0x35, 0x29, 0x4a, 0x4e, 0x82, 0x1e, 0x69, 0xef, 0xde, 0x48, 0x51,
	0xaf, 0x2c, 0x55, 0x26, 0x19, 0x1c, 0xd5, 0x2b, 0x55, 0x86, 0x5f, 0xb3, 0xf3, 0x5e, 0xa5, 0x06,
	0xf7, 0x68, 0x92, 0x61, 0x08, 0xe2, 0x9f, 0xe0, 0xa5, 0x65, 0x7c, 0xc6, 0x26, 0xbb, 0x12, 0xda,
	0x4d, 0x64, 0x4d, 0x32, 0x0a, 0xc1, 0xb8, 0x03, 0xab, 0x66, 0x31, 0x65, 0x17, 0x7f, 0x07, 0x2f,
	0x4d, 0xb3, 0xba, 0x7f, 0xbf, 0xcb, 0x0b, 0x32, 0x2a, 0x13, 0x1f, 0xe8, 0x40, 0x09, 0xed, 0xad,
	0xa0, 0x83, 0x0c, 0x17, 0xa9, 0xbd, 0xb5, 0xde, 0xd5, 0x72, 0x7f, 0x23, 0xc3, 0x0e, 0x65, 0xee,
	0x8d, 0x72, 0xb9, 0x04, 0x2b, 0x2d, 0xd6, 0xd9, 0x30, 0xc0, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7e, 0x7e, 0x71, 0x62, 0x7e, 0x01, 0x00, 0x00,
}
