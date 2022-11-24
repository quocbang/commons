// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_container.proto

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

// Delete Container Info
type DeleteContainerRequest struct {
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteContainerRequest) Reset()         { *m = DeleteContainerRequest{} }
func (m *DeleteContainerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteContainerRequest) ProtoMessage()    {}
func (*DeleteContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_428211ac705f1d24, []int{0}
}

func (m *DeleteContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteContainerRequest.Unmarshal(m, b)
}
func (m *DeleteContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteContainerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteContainerRequest.Merge(m, src)
}
func (m *DeleteContainerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteContainerRequest.Size(m)
}
func (m *DeleteContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteContainerRequest proto.InternalMessageInfo

func (m *DeleteContainerRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

type DeleteContainerReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteContainerReply) Reset()         { *m = DeleteContainerReply{} }
func (m *DeleteContainerReply) String() string { return proto.CompactTextString(m) }
func (*DeleteContainerReply) ProtoMessage()    {}
func (*DeleteContainerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_428211ac705f1d24, []int{1}
}

func (m *DeleteContainerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteContainerReply.Unmarshal(m, b)
}
func (m *DeleteContainerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteContainerReply.Marshal(b, m, deterministic)
}
func (m *DeleteContainerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteContainerReply.Merge(m, src)
}
func (m *DeleteContainerReply) XXX_Size() int {
	return xxx_messageInfo_DeleteContainerReply.Size(m)
}
func (m *DeleteContainerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteContainerReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteContainerReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DeleteContainerRequest)(nil), "proto.dm.mes.DeleteContainerRequest")
	proto.RegisterType((*DeleteContainerReply)(nil), "proto.dm.mes.DeleteContainerReply")
}

func init() { proto.RegisterFile("delete_container.proto", fileDescriptor_428211ac705f1d24) }

var fileDescriptor_428211ac705f1d24 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x8d, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x53, 0x7a, 0x29, 0xb9, 0x7a, 0xb9, 0xa9, 0xc5, 0x4a, 0xd6, 0x5c,
	0x62, 0x2e, 0x60, 0x75, 0xce, 0x30, 0x65, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x8a,
	0x5c, 0x3c, 0x70, 0xad, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xdc, 0x70,
	0x31, 0xcf, 0x14, 0x25, 0x31, 0x2e, 0x11, 0x0c, 0xcd, 0x05, 0x39, 0x95, 0x4e, 0x56, 0x51, 0x16,
	0xe9, 0x99, 0x25, 0x39, 0x89, 0x49, 0x7a, 0xd9, 0xa9, 0x79, 0x29, 0x89, 0x7a, 0xc9, 0xf9, 0xb9,
	0x7a, 0x25, 0xe5, 0xfa, 0x60, 0x8e, 0x7e, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x5e, 0xb1, 0x7e, 0x99,
	0x91, 0x3e, 0xd8, 0x25, 0xfa, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0xfa, 0x29, 0xb9, 0xfa, 0xb9,
	0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x41, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x19, 0x81,
	0xd7, 0xbf, 0x00, 0x00, 0x00,
}
