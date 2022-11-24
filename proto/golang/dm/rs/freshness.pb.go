// Code generated by protoc-gen-go. DO NOT EDIT.
// source: freshness.proto

package rs

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

type GetFreshnessRequest struct {
	// 集團碼
	SubCompany string `protobuf:"bytes,1,opt,name=subCompany,proto3" json:"subCompany,omitempty"`
	// 產品/材料類別
	ProductType ProductType `protobuf:"varint,2,opt,name=product_type,json=productType,proto3,enum=proto.dm.rs.ProductType" json:"product_type,omitempty"`
	// resource name: could be process name or part of material ID. e.g., 10, 79700-9
	Resource string `protobuf:"bytes,3,opt,name=Resource,proto3" json:"Resource,omitempty"`
	// 配方ID
	RecipeID             string   `protobuf:"bytes,4,opt,name=recipeID,proto3" json:"recipeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFreshnessRequest) Reset()         { *m = GetFreshnessRequest{} }
func (m *GetFreshnessRequest) String() string { return proto.CompactTextString(m) }
func (*GetFreshnessRequest) ProtoMessage()    {}
func (*GetFreshnessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_176928e597cde3a1, []int{0}
}

func (m *GetFreshnessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFreshnessRequest.Unmarshal(m, b)
}
func (m *GetFreshnessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFreshnessRequest.Marshal(b, m, deterministic)
}
func (m *GetFreshnessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFreshnessRequest.Merge(m, src)
}
func (m *GetFreshnessRequest) XXX_Size() int {
	return xxx_messageInfo_GetFreshnessRequest.Size(m)
}
func (m *GetFreshnessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFreshnessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFreshnessRequest proto.InternalMessageInfo

func (m *GetFreshnessRequest) GetSubCompany() string {
	if m != nil {
		return m.SubCompany
	}
	return ""
}

func (m *GetFreshnessRequest) GetProductType() ProductType {
	if m != nil {
		return m.ProductType
	}
	return ProductType_PRODUCT_TYPE_UNSPECIFIED
}

func (m *GetFreshnessRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *GetFreshnessRequest) GetRecipeID() string {
	if m != nil {
		return m.RecipeID
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFreshnessRequest)(nil), "proto.dm.rs.GetFreshnessRequest")
}

func init() { proto.RegisterFile("freshness.proto", fileDescriptor_176928e597cde3a1) }

var fileDescriptor_176928e597cde3a1 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x59, 0x15, 0xd1, 0x9c, 0x28, 0xac, 0x85, 0xe1, 0x0a, 0x39, 0xac, 0xae, 0x9a, 0xc0,
	0x59, 0x88, 0xd8, 0xa9, 0x28, 0x76, 0xb2, 0x58, 0xd9, 0xc8, 0x6e, 0x32, 0xae, 0x87, 0x97, 0x4c,
	0xcc, 0x24, 0xca, 0xfe, 0x25, 0x7f, 0xa5, 0x98, 0x78, 0xde, 0x56, 0xc3, 0xf7, 0xde, 0x9b, 0x19,
	0x9e, 0x38, 0x7a, 0x0d, 0xc8, 0x6f, 0x0e, 0x99, 0xc1, 0x07, 0x8a, 0x54, 0x4f, 0xf2, 0x00, 0x63,
	0x21, 0xf0, 0xf4, 0x24, 0x83, 0x32, 0x56, 0x05, 0x56, 0xe8, 0x92, 0xfd, 0x4b, 0x9d, 0x7d, 0x57,
	0xe2, 0xf8, 0x1e, 0xe3, 0xdd, 0x7a, 0xb9, 0xc1, 0x8f, 0x84, 0x1c, 0xeb, 0x53, 0x21, 0x38, 0x75,
	0x37, 0x64, 0x7d, 0xeb, 0x06, 0x59, 0xcd, 0xaa, 0xf9, 0x7e, 0x33, 0x52, 0xea, 0x2b, 0x71, 0xe0,
	0x03, 0x99, 0xa4, 0xe3, 0x4b, 0x1c, 0x3c, 0xca, 0xad, 0x59, 0x35, 0x3f, 0x5c, 0x48, 0x18, 0x3d,
	0x85, 0xc7, 0x12, 0x78, 0x1a, 0x3c, 0x36, 0x13, 0xbf, 0x81, 0x7a, 0x2a, 0xf6, 0x1a, 0x64, 0x4a,
	0x41, 0xa3, 0xdc, 0xce, 0xa7, 0xff, 0xf9, 0xd7, 0x0b, 0xa8, 0x97, 0x1e, 0x1f, 0x6e, 0xe5, 0x4e,
	0xf1, 0xd6, 0x7c, 0x7d, 0xf9, 0x7c, 0xd1, 0x2f, 0xe3, 0xaa, 0xed, 0xe0, 0x1d, 0x9d, 0x69, 0x41,
	0x93, 0x85, 0xf8, 0xa5, 0x32, 0x28, 0x4d, 0xd6, 0x92, 0x63, 0xf5, 0xb9, 0x50, 0xa5, 0x6c, 0x4f,
	0xab, 0xd6, 0xf5, 0xa5, 0x73, 0xb7, 0x9b, 0xb5, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d,
	0xdb, 0xa4, 0x74, 0x27, 0x01, 0x00, 0x00,
}
