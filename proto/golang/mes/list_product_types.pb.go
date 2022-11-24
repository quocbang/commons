// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_product_types.proto

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

type ListProductTypesRequest struct {
	FactoryId            rs.FactoryID `protobuf:"varint,1,opt,name=factory_id,json=factoryId,proto3,enum=proto.dm.rs.FactoryID" json:"factory_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListProductTypesRequest) Reset()         { *m = ListProductTypesRequest{} }
func (m *ListProductTypesRequest) String() string { return proto.CompactTextString(m) }
func (*ListProductTypesRequest) ProtoMessage()    {}
func (*ListProductTypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7845652fed7e61c, []int{0}
}

func (m *ListProductTypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductTypesRequest.Unmarshal(m, b)
}
func (m *ListProductTypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductTypesRequest.Marshal(b, m, deterministic)
}
func (m *ListProductTypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductTypesRequest.Merge(m, src)
}
func (m *ListProductTypesRequest) XXX_Size() int {
	return xxx_messageInfo_ListProductTypesRequest.Size(m)
}
func (m *ListProductTypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductTypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductTypesRequest proto.InternalMessageInfo

func (m *ListProductTypesRequest) GetFactoryId() rs.FactoryID {
	if m != nil {
		return m.FactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

type ListProductTypesReply struct {
	ProductTypes         []*ListProductTypesReply_ProductType `protobuf:"bytes,1,rep,name=product_types,json=productTypes,proto3" json:"product_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ListProductTypesReply) Reset()         { *m = ListProductTypesReply{} }
func (m *ListProductTypesReply) String() string { return proto.CompactTextString(m) }
func (*ListProductTypesReply) ProtoMessage()    {}
func (*ListProductTypesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7845652fed7e61c, []int{1}
}

func (m *ListProductTypesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductTypesReply.Unmarshal(m, b)
}
func (m *ListProductTypesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductTypesReply.Marshal(b, m, deterministic)
}
func (m *ListProductTypesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductTypesReply.Merge(m, src)
}
func (m *ListProductTypesReply) XXX_Size() int {
	return xxx_messageInfo_ListProductTypesReply.Size(m)
}
func (m *ListProductTypesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductTypesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductTypesReply proto.InternalMessageInfo

func (m *ListProductTypesReply) GetProductTypes() []*ListProductTypesReply_ProductType {
	if m != nil {
		return m.ProductTypes
	}
	return nil
}

type ListProductTypesReply_ProductType struct {
	Id                   rs.ProductType `protobuf:"varint,1,opt,name=id,proto3,enum=proto.dm.rs.ProductType" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListProductTypesReply_ProductType) Reset()         { *m = ListProductTypesReply_ProductType{} }
func (m *ListProductTypesReply_ProductType) String() string { return proto.CompactTextString(m) }
func (*ListProductTypesReply_ProductType) ProtoMessage()    {}
func (*ListProductTypesReply_ProductType) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7845652fed7e61c, []int{1, 0}
}

func (m *ListProductTypesReply_ProductType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductTypesReply_ProductType.Unmarshal(m, b)
}
func (m *ListProductTypesReply_ProductType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductTypesReply_ProductType.Marshal(b, m, deterministic)
}
func (m *ListProductTypesReply_ProductType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductTypesReply_ProductType.Merge(m, src)
}
func (m *ListProductTypesReply_ProductType) XXX_Size() int {
	return xxx_messageInfo_ListProductTypesReply_ProductType.Size(m)
}
func (m *ListProductTypesReply_ProductType) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductTypesReply_ProductType.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductTypesReply_ProductType proto.InternalMessageInfo

func (m *ListProductTypesReply_ProductType) GetId() rs.ProductType {
	if m != nil {
		return m.Id
	}
	return rs.ProductType_PRODUCT_TYPE_UNSPECIFIED
}

func (m *ListProductTypesReply_ProductType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ListProductTypesRequest)(nil), "proto.mes.ListProductTypesRequest")
	proto.RegisterType((*ListProductTypesReply)(nil), "proto.mes.ListProductTypesReply")
	proto.RegisterType((*ListProductTypesReply_ProductType)(nil), "proto.mes.ListProductTypesReply.ProductType")
}

func init() { proto.RegisterFile("list_product_types.proto", fileDescriptor_d7845652fed7e61c) }

var fileDescriptor_d7845652fed7e61c = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0x15, 0xa1, 0x99, 0x7a, 0x08, 0xe8, 0xc2, 0x4e, 0x65, 0xa7, 0x1e, 0xe4, 0x05,
	0x2a, 0xc3, 0xbb, 0x88, 0x30, 0xf4, 0x30, 0x8b, 0x27, 0x2f, 0x25, 0x6b, 0x62, 0x09, 0x36, 0x3f,
	0x4c, 0x52, 0xa5, 0x7f, 0x99, 0xff, 0x9e, 0x98, 0x8e, 0xd1, 0xe1, 0x4e, 0x79, 0x79, 0xdf, 0xf7,
	0x79, 0x7c, 0x78, 0x88, 0x74, 0xd2, 0x87, 0xda, 0x3a, 0xc3, 0xfb, 0x26, 0xd4, 0x61, 0xb0, 0xc2,
	0x83, 0x75, 0x26, 0x18, 0x9c, 0xc5, 0x07, 0x94, 0xf0, 0x8b, 0x79, 0x2c, 0x29, 0x57, 0xd4, 0x79,
	0x2a, 0x74, 0xaf, 0x76, 0x33, 0xcb, 0x0d, 0x9a, 0x3f, 0x4b, 0x1f, 0x36, 0x23, 0xfe, 0xfa, 0x47,
	0x57, 0xe2, 0xb3, 0x17, 0x3e, 0xe0, 0x15, 0x42, 0xef, 0xac, 0x09, 0xc6, 0x0d, 0xb5, 0xe4, 0x24,
	0xc9, 0x93, 0xe2, 0xb2, 0xbc, 0x1e, 0x31, 0xe0, 0x0a, 0x9c, 0x87, 0xc7, 0x31, 0x5e, 0x3f, 0x54,
	0xd9, 0x6e, 0x72, 0xcd, 0x97, 0x3f, 0x09, 0xba, 0xfa, 0xbf, 0xd2, 0x76, 0x03, 0x7e, 0x41, 0x17,
	0x07, 0x9a, 0x24, 0xc9, 0x4f, 0x8a, 0x59, 0x79, 0x03, 0x7b, 0x4f, 0x38, 0x0a, 0xc2, 0xa4, 0x53,
	0x9d, 0xdb, 0x49, 0xbc, 0x78, 0x42, 0xb3, 0x49, 0x88, 0x0b, 0x94, 0xee, 0x55, 0xc9, 0x81, 0xea,
	0x74, 0x45, 0x2a, 0x39, 0xc6, 0xe8, 0x54, 0x33, 0x25, 0x48, 0x9a, 0x27, 0x45, 0x56, 0xc5, 0xfa,
	0xfe, 0xee, 0x6d, 0xd5, 0xca, 0xd0, 0xb1, 0x2d, 0x7c, 0x08, 0xcd, 0x19, 0x34, 0x46, 0x41, 0xf8,
	0xa6, 0xf1, 0x43, 0x1b, 0xa3, 0x94, 0xd1, 0x9e, 0x7e, 0x95, 0x74, 0xbc, 0x65, 0x6b, 0x3a, 0xa6,
	0x5b, 0xaa, 0x84, 0xdf, 0x9e, 0xc5, 0xce, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x6a,
	0x4d, 0x9b, 0x8b, 0x01, 0x00, 0x00,
}