// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_meqptcon.proto

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

type GetMeqptconRequest struct {
	EqptId               string   `protobuf:"bytes,1,opt,name=eqpt_id,json=eqptId,proto3" json:"eqpt_id,omitempty"`
	ContainerId          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMeqptconRequest) Reset()         { *m = GetMeqptconRequest{} }
func (m *GetMeqptconRequest) String() string { return proto.CompactTextString(m) }
func (*GetMeqptconRequest) ProtoMessage()    {}
func (*GetMeqptconRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0862480bb5c93632, []int{0}
}

func (m *GetMeqptconRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMeqptconRequest.Unmarshal(m, b)
}
func (m *GetMeqptconRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMeqptconRequest.Marshal(b, m, deterministic)
}
func (m *GetMeqptconRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMeqptconRequest.Merge(m, src)
}
func (m *GetMeqptconRequest) XXX_Size() int {
	return xxx_messageInfo_GetMeqptconRequest.Size(m)
}
func (m *GetMeqptconRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMeqptconRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMeqptconRequest proto.InternalMessageInfo

func (m *GetMeqptconRequest) GetEqptId() string {
	if m != nil {
		return m.EqptId
	}
	return ""
}

func (m *GetMeqptconRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

type GetMeqptconReply struct {
	Info                 []*GetMeqptconReply_Meqptcon `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetMeqptconReply) Reset()         { *m = GetMeqptconReply{} }
func (m *GetMeqptconReply) String() string { return proto.CompactTextString(m) }
func (*GetMeqptconReply) ProtoMessage()    {}
func (*GetMeqptconReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0862480bb5c93632, []int{1}
}

func (m *GetMeqptconReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMeqptconReply.Unmarshal(m, b)
}
func (m *GetMeqptconReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMeqptconReply.Marshal(b, m, deterministic)
}
func (m *GetMeqptconReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMeqptconReply.Merge(m, src)
}
func (m *GetMeqptconReply) XXX_Size() int {
	return xxx_messageInfo_GetMeqptconReply.Size(m)
}
func (m *GetMeqptconReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMeqptconReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMeqptconReply proto.InternalMessageInfo

func (m *GetMeqptconReply) GetInfo() []*GetMeqptconReply_Meqptcon {
	if m != nil {
		return m.Info
	}
	return nil
}

type GetMeqptconReply_Meqptcon struct {
	EqptId               string   `protobuf:"bytes,1,opt,name=eqpt_id,json=eqptId,proto3" json:"eqpt_id,omitempty"`
	ContainerId          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	UpdateBy             string   `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	UpdateAt             string   `protobuf:"bytes,4,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMeqptconReply_Meqptcon) Reset()         { *m = GetMeqptconReply_Meqptcon{} }
func (m *GetMeqptconReply_Meqptcon) String() string { return proto.CompactTextString(m) }
func (*GetMeqptconReply_Meqptcon) ProtoMessage()    {}
func (*GetMeqptconReply_Meqptcon) Descriptor() ([]byte, []int) {
	return fileDescriptor_0862480bb5c93632, []int{1, 0}
}

func (m *GetMeqptconReply_Meqptcon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMeqptconReply_Meqptcon.Unmarshal(m, b)
}
func (m *GetMeqptconReply_Meqptcon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMeqptconReply_Meqptcon.Marshal(b, m, deterministic)
}
func (m *GetMeqptconReply_Meqptcon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMeqptconReply_Meqptcon.Merge(m, src)
}
func (m *GetMeqptconReply_Meqptcon) XXX_Size() int {
	return xxx_messageInfo_GetMeqptconReply_Meqptcon.Size(m)
}
func (m *GetMeqptconReply_Meqptcon) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMeqptconReply_Meqptcon.DiscardUnknown(m)
}

var xxx_messageInfo_GetMeqptconReply_Meqptcon proto.InternalMessageInfo

func (m *GetMeqptconReply_Meqptcon) GetEqptId() string {
	if m != nil {
		return m.EqptId
	}
	return ""
}

func (m *GetMeqptconReply_Meqptcon) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *GetMeqptconReply_Meqptcon) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *GetMeqptconReply_Meqptcon) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMeqptconRequest)(nil), "proto.dm.mes.GetMeqptconRequest")
	proto.RegisterType((*GetMeqptconReply)(nil), "proto.dm.mes.GetMeqptconReply")
	proto.RegisterType((*GetMeqptconReply_Meqptcon)(nil), "proto.dm.mes.GetMeqptconReply.Meqptcon")
}

func init() { proto.RegisterFile("query_meqptcon.proto", fileDescriptor_0862480bb5c93632) }

var fileDescriptor_0862480bb5c93632 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0x3d, 0x4f, 0xf3, 0x30,
	0x14, 0x85, 0xe5, 0xb7, 0x55, 0xdf, 0xd6, 0xed, 0x80, 0x2c, 0x24, 0x22, 0x58, 0x4a, 0x17, 0x3a,
	0x5d, 0x4b, 0x65, 0x41, 0x30, 0xd1, 0x05, 0x75, 0x40, 0x42, 0x1d, 0x59, 0x22, 0x27, 0xbe, 0x44,
	0x11, 0xb1, 0x6f, 0x3e, 0x6e, 0x40, 0xd9, 0xf8, 0x8d, 0xfc, 0x22, 0x14, 0x03, 0x55, 0x60, 0x65,
	0xb2, 0xce, 0xf3, 0x1c, 0x5b, 0xc7, 0xf2, 0xb8, 0x6a, 0xb1, 0xee, 0x62, 0x87, 0x55, 0xc9, 0x29,
	0x79, 0x28, 0x6b, 0x62, 0x52, 0x8b, 0x70, 0x80, 0x75, 0xe0, 0xb0, 0x59, 0x3d, 0x48, 0x75, 0x87,
	0x7c, 0xff, 0x55, 0xd9, 0x63, 0xd5, 0x62, 0xc3, 0xea, 0x44, 0xfe, 0xef, 0x49, 0x9c, 0xdb, 0x48,
	0x2c, 0xc5, 0x7a, 0xb6, 0x9f, 0xf4, 0x71, 0x67, 0xd5, 0xb9, 0x5c, 0xa4, 0xe4, 0xd9, 0xe4, 0x1e,
	0xeb, 0xde, 0xfe, 0x0b, 0x76, 0x7e, 0x60, 0x3b, 0xbb, 0x7a, 0x17, 0xf2, 0xe8, 0xc7, 0x93, 0x65,
	0xd1, 0xa9, 0x1b, 0x39, 0xce, 0xfd, 0x13, 0x45, 0x62, 0x39, 0x5a, 0xcf, 0x37, 0x17, 0x30, 0xdc,
	0x00, 0xbf, 0xdb, 0x70, 0x48, 0xe1, 0xd2, 0xe9, 0x9b, 0x90, 0xd3, 0x6f, 0xf4, 0x97, 0x69, 0xea,
	0x4c, 0xce, 0xda, 0xd2, 0x1a, 0xc6, 0x38, 0xe9, 0xa2, 0x51, 0xf0, 0xd3, 0x4f, 0xb0, 0xed, 0x06,
	0xd2, 0x70, 0x34, 0x1e, 0xca, 0x5b, 0xde, 0x5e, 0x3f, 0x5e, 0x65, 0x39, 0x17, 0x26, 0x81, 0x67,
	0xf4, 0xd6, 0x40, 0x4a, 0x0e, 0xf8, 0x55, 0x87, 0xa0, 0x53, 0x72, 0x8e, 0x7c, 0xa3, 0x5f, 0x36,
	0x3a, 0xfc, 0x4b, 0x67, 0x54, 0x18, 0x9f, 0x69, 0xeb, 0xb4, 0xc3, 0x26, 0x99, 0x04, 0x78, 0xf9,
	0x11, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xe7, 0x92, 0xd4, 0x8f, 0x01, 0x00, 0x00,
}
