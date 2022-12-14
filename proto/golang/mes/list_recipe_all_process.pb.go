// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_recipe_all_process.proto

package mes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	rs "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
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

type ListRecipeAllProcessesRequest struct {
	RecipeId             string   `protobuf:"bytes,1,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRecipeAllProcessesRequest) Reset()         { *m = ListRecipeAllProcessesRequest{} }
func (m *ListRecipeAllProcessesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRecipeAllProcessesRequest) ProtoMessage()    {}
func (*ListRecipeAllProcessesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d046d31eb363d5, []int{0}
}

func (m *ListRecipeAllProcessesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecipeAllProcessesRequest.Unmarshal(m, b)
}
func (m *ListRecipeAllProcessesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecipeAllProcessesRequest.Marshal(b, m, deterministic)
}
func (m *ListRecipeAllProcessesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecipeAllProcessesRequest.Merge(m, src)
}
func (m *ListRecipeAllProcessesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRecipeAllProcessesRequest.Size(m)
}
func (m *ListRecipeAllProcessesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecipeAllProcessesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecipeAllProcessesRequest proto.InternalMessageInfo

func (m *ListRecipeAllProcessesRequest) GetRecipeId() string {
	if m != nil {
		return m.RecipeId
	}
	return ""
}

type ListRecipeAllProcessesReply struct {
	Processes            []*ListRecipeAllProcessesReply_Process `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *ListRecipeAllProcessesReply) Reset()         { *m = ListRecipeAllProcessesReply{} }
func (m *ListRecipeAllProcessesReply) String() string { return proto.CompactTextString(m) }
func (*ListRecipeAllProcessesReply) ProtoMessage()    {}
func (*ListRecipeAllProcessesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d046d31eb363d5, []int{1}
}

func (m *ListRecipeAllProcessesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecipeAllProcessesReply.Unmarshal(m, b)
}
func (m *ListRecipeAllProcessesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecipeAllProcessesReply.Marshal(b, m, deterministic)
}
func (m *ListRecipeAllProcessesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecipeAllProcessesReply.Merge(m, src)
}
func (m *ListRecipeAllProcessesReply) XXX_Size() int {
	return xxx_messageInfo_ListRecipeAllProcessesReply.Size(m)
}
func (m *ListRecipeAllProcessesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecipeAllProcessesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecipeAllProcessesReply proto.InternalMessageInfo

func (m *ListRecipeAllProcessesReply) GetProcesses() []*ListRecipeAllProcessesReply_Process {
	if m != nil {
		return m.Processes
	}
	return nil
}

type ListRecipeAllProcessesReply_Process struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 rs.ProcessType `protobuf:"varint,2,opt,name=type,proto3,enum=proto.dm.rs.ProcessType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListRecipeAllProcessesReply_Process) Reset()         { *m = ListRecipeAllProcessesReply_Process{} }
func (m *ListRecipeAllProcessesReply_Process) String() string { return proto.CompactTextString(m) }
func (*ListRecipeAllProcessesReply_Process) ProtoMessage()    {}
func (*ListRecipeAllProcessesReply_Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d046d31eb363d5, []int{1, 0}
}

func (m *ListRecipeAllProcessesReply_Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecipeAllProcessesReply_Process.Unmarshal(m, b)
}
func (m *ListRecipeAllProcessesReply_Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecipeAllProcessesReply_Process.Marshal(b, m, deterministic)
}
func (m *ListRecipeAllProcessesReply_Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecipeAllProcessesReply_Process.Merge(m, src)
}
func (m *ListRecipeAllProcessesReply_Process) XXX_Size() int {
	return xxx_messageInfo_ListRecipeAllProcessesReply_Process.Size(m)
}
func (m *ListRecipeAllProcessesReply_Process) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecipeAllProcessesReply_Process.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecipeAllProcessesReply_Process proto.InternalMessageInfo

func (m *ListRecipeAllProcessesReply_Process) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListRecipeAllProcessesReply_Process) GetType() rs.ProcessType {
	if m != nil {
		return m.Type
	}
	return rs.ProcessType_PROCESS_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*ListRecipeAllProcessesRequest)(nil), "proto.mes.ListRecipeAllProcessesRequest")
	proto.RegisterType((*ListRecipeAllProcessesReply)(nil), "proto.mes.ListRecipeAllProcessesReply")
	proto.RegisterType((*ListRecipeAllProcessesReply_Process)(nil), "proto.mes.ListRecipeAllProcessesReply.Process")
}

func init() { proto.RegisterFile("list_recipe_all_process.proto", fileDescriptor_a0d046d31eb363d5) }

var fileDescriptor_a0d046d31eb363d5 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0x87, 0x89, 0x1e, 0x6a, 0x56, 0xb0, 0xd8, 0xc6, 0x70, 0xc7, 0x41, 0xb8, 0x2a, 0x85, 0xcc,
	0x42, 0x44, 0x6c, 0x6c, 0xb4, 0x13, 0xaf, 0x90, 0x60, 0x65, 0x13, 0x72, 0xc9, 0x10, 0x16, 0xf7,
	0x9f, 0x3b, 0x7b, 0x4a, 0xde, 0xcc, 0xc7, 0x13, 0x36, 0x7b, 0x5a, 0x69, 0xb5, 0x33, 0x1f, 0x33,
	0xdf, 0xfc, 0x58, 0xb6, 0x56, 0x92, 0x42, 0xeb, 0xb1, 0x97, 0x0e, 0xdb, 0x4e, 0xa9, 0xd6, 0x79,
	0xdb, 0x23, 0x11, 0x38, 0x6f, 0x83, 0xe5, 0x79, 0x7c, 0x40, 0x23, 0x2d, 0x2f, 0x63, 0x29, 0x06,
	0x2d, 0x3c, 0x09, 0x34, 0x7b, 0x9d, 0x66, 0x36, 0x77, 0x6c, 0xbd, 0x95, 0x14, 0x9a, 0xe8, 0xb8,
	0x57, 0xea, 0x79, 0x36, 0x20, 0x35, 0xf8, 0xbe, 0x47, 0x0a, 0x7c, 0xc5, 0xf2, 0x74, 0x40, 0x0e,
	0x45, 0x56, 0x66, 0x55, 0xde, 0x9c, 0xcd, 0xe0, 0x71, 0xd8, 0x7c, 0x65, 0x6c, 0xf5, 0xd7, 0xba,
	0x53, 0x13, 0xdf, 0xb2, 0xdc, 0x1d, 0x48, 0x91, 0x95, 0xc7, 0xd5, 0x79, 0x0d, 0xf0, 0x93, 0x0a,
	0xfe, 0x59, 0x85, 0xd4, 0x36, 0xbf, 0x82, 0xe5, 0x13, 0x3b, 0x4d, 0x94, 0x73, 0xb6, 0x30, 0x9d,
	0xc6, 0x14, 0x28, 0xd6, 0xfc, 0x8a, 0x2d, 0xc2, 0xe4, 0xb0, 0x38, 0x2a, 0xb3, 0xea, 0xa2, 0x2e,
	0xd2, 0x9d, 0x41, 0x83, 0xa7, 0x83, 0xed, 0x65, 0x72, 0xd8, 0xc4, 0xa9, 0x87, 0xdb, 0xd7, 0x9b,
	0x51, 0x06, 0xd5, 0xed, 0xe0, 0x0d, 0xcd, 0xd0, 0x41, 0x6f, 0x35, 0x84, 0x4f, 0x11, 0x1b, 0xd1,
	0x5b, 0xad, 0xad, 0x21, 0xf1, 0x51, 0x8b, 0xf9, 0xe3, 0x46, 0xab, 0x3a, 0x33, 0x0a, 0x8d, 0xb4,
	0x3b, 0x89, 0xe4, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x51, 0xe8, 0xda, 0x7d, 0x01, 0x00,
	0x00,
}
