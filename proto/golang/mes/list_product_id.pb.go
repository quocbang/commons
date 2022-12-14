// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_product_id.proto

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

type ListMaterialProductIDsRequest struct {
	FactoryId            rs.FactoryID   `protobuf:"varint,1,opt,name=factory_id,json=factoryId,proto3,enum=proto.dm.rs.FactoryID" json:"factory_id,omitempty"`
	ProductType          rs.ProductType `protobuf:"varint,2,opt,name=product_type,json=productType,proto3,enum=proto.dm.rs.ProductType" json:"product_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListMaterialProductIDsRequest) Reset()         { *m = ListMaterialProductIDsRequest{} }
func (m *ListMaterialProductIDsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMaterialProductIDsRequest) ProtoMessage()    {}
func (*ListMaterialProductIDsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7779759e6a5db98d, []int{0}
}

func (m *ListMaterialProductIDsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMaterialProductIDsRequest.Unmarshal(m, b)
}
func (m *ListMaterialProductIDsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMaterialProductIDsRequest.Marshal(b, m, deterministic)
}
func (m *ListMaterialProductIDsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMaterialProductIDsRequest.Merge(m, src)
}
func (m *ListMaterialProductIDsRequest) XXX_Size() int {
	return xxx_messageInfo_ListMaterialProductIDsRequest.Size(m)
}
func (m *ListMaterialProductIDsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMaterialProductIDsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMaterialProductIDsRequest proto.InternalMessageInfo

func (m *ListMaterialProductIDsRequest) GetFactoryId() rs.FactoryID {
	if m != nil {
		return m.FactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

func (m *ListMaterialProductIDsRequest) GetProductType() rs.ProductType {
	if m != nil {
		return m.ProductType
	}
	return rs.ProductType_PRODUCT_TYPE_UNSPECIFIED
}

type ListMaterialProductIDsReply struct {
	MaterialProducts     []*ListMaterialProductIDsReply_Product `protobuf:"bytes,1,rep,name=material_products,json=materialProducts,proto3" json:"material_products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *ListMaterialProductIDsReply) Reset()         { *m = ListMaterialProductIDsReply{} }
func (m *ListMaterialProductIDsReply) String() string { return proto.CompactTextString(m) }
func (*ListMaterialProductIDsReply) ProtoMessage()    {}
func (*ListMaterialProductIDsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7779759e6a5db98d, []int{1}
}

func (m *ListMaterialProductIDsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMaterialProductIDsReply.Unmarshal(m, b)
}
func (m *ListMaterialProductIDsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMaterialProductIDsReply.Marshal(b, m, deterministic)
}
func (m *ListMaterialProductIDsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMaterialProductIDsReply.Merge(m, src)
}
func (m *ListMaterialProductIDsReply) XXX_Size() int {
	return xxx_messageInfo_ListMaterialProductIDsReply.Size(m)
}
func (m *ListMaterialProductIDsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMaterialProductIDsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListMaterialProductIDsReply proto.InternalMessageInfo

func (m *ListMaterialProductIDsReply) GetMaterialProducts() []*ListMaterialProductIDsReply_Product {
	if m != nil {
		return m.MaterialProducts
	}
	return nil
}

type ListMaterialProductIDsReply_Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMaterialProductIDsReply_Product) Reset()         { *m = ListMaterialProductIDsReply_Product{} }
func (m *ListMaterialProductIDsReply_Product) String() string { return proto.CompactTextString(m) }
func (*ListMaterialProductIDsReply_Product) ProtoMessage()    {}
func (*ListMaterialProductIDsReply_Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_7779759e6a5db98d, []int{1, 0}
}

func (m *ListMaterialProductIDsReply_Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMaterialProductIDsReply_Product.Unmarshal(m, b)
}
func (m *ListMaterialProductIDsReply_Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMaterialProductIDsReply_Product.Marshal(b, m, deterministic)
}
func (m *ListMaterialProductIDsReply_Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMaterialProductIDsReply_Product.Merge(m, src)
}
func (m *ListMaterialProductIDsReply_Product) XXX_Size() int {
	return xxx_messageInfo_ListMaterialProductIDsReply_Product.Size(m)
}
func (m *ListMaterialProductIDsReply_Product) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMaterialProductIDsReply_Product.DiscardUnknown(m)
}

var xxx_messageInfo_ListMaterialProductIDsReply_Product proto.InternalMessageInfo

func (m *ListMaterialProductIDsReply_Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListMaterialProductIDsReply_Product) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func init() {
	proto.RegisterType((*ListMaterialProductIDsRequest)(nil), "proto.mes.ListMaterialProductIDsRequest")
	proto.RegisterType((*ListMaterialProductIDsReply)(nil), "proto.mes.ListMaterialProductIDsReply")
	proto.RegisterType((*ListMaterialProductIDsReply_Product)(nil), "proto.mes.ListMaterialProductIDsReply.Product")
}

func init() { proto.RegisterFile("list_product_id.proto", fileDescriptor_7779759e6a5db98d) }

var fileDescriptor_7779759e6a5db98d = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4f, 0x4b, 0x3b, 0x31,
	0x10, 0x65, 0xfb, 0xe3, 0xa7, 0x6c, 0x2a, 0x45, 0x83, 0x7f, 0x96, 0x8a, 0x50, 0x7a, 0xea, 0x29,
	0x81, 0x4a, 0xf1, 0xe0, 0x4d, 0x8a, 0x50, 0x50, 0x90, 0xe0, 0x49, 0x0f, 0x25, 0xdd, 0xc4, 0x25,
	0x98, 0x6c, 0x62, 0x26, 0x5b, 0xd9, 0xcf, 0xe1, 0xc7, 0xf0, 0x4b, 0x0a, 0xc9, 0xae, 0xd2, 0x83,
	0x9e, 0x32, 0x2f, 0xf3, 0xde, 0xcc, 0x7b, 0x0c, 0x3a, 0xd1, 0x0a, 0xc2, 0xda, 0x79, 0x2b, 0x9a,
	0x32, 0xac, 0x95, 0x20, 0xce, 0xdb, 0x60, 0x71, 0x1e, 0x1f, 0x62, 0x24, 0x8c, 0xcf, 0x62, 0x49,
	0x85, 0xa1, 0x1e, 0xa8, 0xac, 0x1b, 0x03, 0x89, 0x33, 0xfd, 0xc8, 0xd0, 0xc5, 0x9d, 0x82, 0x70,
	0xcf, 0x83, 0xf4, 0x8a, 0xeb, 0x87, 0x34, 0x64, 0xb5, 0x04, 0x26, 0xdf, 0x1a, 0x09, 0x01, 0x2f,
	0x10, 0x7a, 0xe1, 0x65, 0xb0, 0xbe, 0x5d, 0x2b, 0x51, 0x64, 0x93, 0x6c, 0x36, 0x9a, 0x9f, 0x26,
	0x35, 0x11, 0x86, 0x78, 0x20, 0xb7, 0xa9, 0xbd, 0x5a, 0xb2, 0xbc, 0x63, 0xae, 0x04, 0xbe, 0x46,
	0x07, 0xbd, 0xa1, 0xd0, 0x3a, 0x59, 0x0c, 0xa2, 0xb0, 0xd8, 0x11, 0x76, 0xcb, 0x1e, 0x5b, 0x27,
	0xd9, 0xd0, 0xfd, 0x80, 0xe9, 0x67, 0x86, 0xce, 0x7f, 0x73, 0xe5, 0x74, 0x8b, 0x9f, 0xd1, 0x91,
	0xe9, 0x5a, 0x7d, 0x6c, 0x28, 0xb2, 0xc9, 0xbf, 0xd9, 0x70, 0x4e, 0xc8, 0x77, 0x6a, 0xf2, 0xc7,
	0x88, 0x7e, 0x37, 0x3b, 0x34, 0xbb, 0x04, 0x18, 0x53, 0xb4, 0xdf, 0xd5, 0x78, 0x84, 0x06, 0x5d,
	0xe6, 0x9c, 0x0d, 0x94, 0xc0, 0xc7, 0xe8, 0xbf, 0x96, 0x5b, 0xa9, 0x63, 0x9a, 0x9c, 0x25, 0x70,
	0x73, 0xf5, 0xb4, 0xa8, 0x54, 0xd0, 0x7c, 0x43, 0x5e, 0x65, 0x2d, 0x38, 0x29, 0xad, 0x21, 0xe1,
	0x9d, 0x46, 0x40, 0x4b, 0x6b, 0x8c, 0xad, 0x81, 0x6e, 0xe7, 0x34, 0xdd, 0xa0, 0xb2, 0x9a, 0xd7,
	0x15, 0x35, 0x12, 0x36, 0x7b, 0xf1, 0xe7, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x42, 0x86, 0xd9,
	0xfd, 0xc0, 0x01, 0x00, 0x00,
}
