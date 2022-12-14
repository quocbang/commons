// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_bom_materials.proto

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

// 查詢
type GetBomRequest struct {
	FactoryId   rs.FactoryID `protobuf:"varint,1,opt,name=factory_id,json=factoryId,proto3,enum=proto.dm.rs.FactoryID" json:"factory_id,omitempty"`
	EquipmentId string       `protobuf:"bytes,2,opt,name=equipment_id,json=equipmentId,proto3" json:"equipment_id,omitempty"`
	RecipeId    string       `protobuf:"bytes,3,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	// 途程名稱
	ProcessName          string   `protobuf:"bytes,4,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	RecipeEquipmentId    string   `protobuf:"bytes,5,opt,name=recipe_equipment_id,json=recipeEquipmentId,proto3" json:"recipe_equipment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBomRequest) Reset()         { *m = GetBomRequest{} }
func (m *GetBomRequest) String() string { return proto.CompactTextString(m) }
func (*GetBomRequest) ProtoMessage()    {}
func (*GetBomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d547b7fe597de32, []int{0}
}

func (m *GetBomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBomRequest.Unmarshal(m, b)
}
func (m *GetBomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBomRequest.Marshal(b, m, deterministic)
}
func (m *GetBomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBomRequest.Merge(m, src)
}
func (m *GetBomRequest) XXX_Size() int {
	return xxx_messageInfo_GetBomRequest.Size(m)
}
func (m *GetBomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBomRequest proto.InternalMessageInfo

func (m *GetBomRequest) GetFactoryId() rs.FactoryID {
	if m != nil {
		return m.FactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

func (m *GetBomRequest) GetEquipmentId() string {
	if m != nil {
		return m.EquipmentId
	}
	return ""
}

func (m *GetBomRequest) GetRecipeId() string {
	if m != nil {
		return m.RecipeId
	}
	return ""
}

func (m *GetBomRequest) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *GetBomRequest) GetRecipeEquipmentId() string {
	if m != nil {
		return m.RecipeEquipmentId
	}
	return ""
}

type GetBomReply struct {
	BomDetails           []*GetBomReply_BomDetail `protobuf:"bytes,1,rep,name=bom_details,json=bomDetails,proto3" json:"bom_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetBomReply) Reset()         { *m = GetBomReply{} }
func (m *GetBomReply) String() string { return proto.CompactTextString(m) }
func (*GetBomReply) ProtoMessage()    {}
func (*GetBomReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d547b7fe597de32, []int{1}
}

func (m *GetBomReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBomReply.Unmarshal(m, b)
}
func (m *GetBomReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBomReply.Marshal(b, m, deterministic)
}
func (m *GetBomReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBomReply.Merge(m, src)
}
func (m *GetBomReply) XXX_Size() int {
	return xxx_messageInfo_GetBomReply.Size(m)
}
func (m *GetBomReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBomReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBomReply proto.InternalMessageInfo

func (m *GetBomReply) GetBomDetails() []*GetBomReply_BomDetail {
	if m != nil {
		return m.BomDetails
	}
	return nil
}

type GetBomReply_BomDetail struct {
	MaterialProductId    string   `protobuf:"bytes,1,opt,name=material_product_id,json=materialProductId,proto3" json:"material_product_id,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Qty                  float32  `protobuf:"fixed32,3,opt,name=qty,proto3" json:"qty,omitempty"`
	ErrorMargin          float32  `protobuf:"fixed32,4,opt,name=error_margin,json=errorMargin,proto3" json:"error_margin,omitempty"`
	SpecifiedRecipeId    string   `protobuf:"bytes,5,opt,name=specified_recipe_id,json=specifiedRecipeId,proto3" json:"specified_recipe_id,omitempty"`
	AlternativeMaterials string   `protobuf:"bytes,6,opt,name=alternative_materials,json=alternativeMaterials,proto3" json:"alternative_materials,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBomReply_BomDetail) Reset()         { *m = GetBomReply_BomDetail{} }
func (m *GetBomReply_BomDetail) String() string { return proto.CompactTextString(m) }
func (*GetBomReply_BomDetail) ProtoMessage()    {}
func (*GetBomReply_BomDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d547b7fe597de32, []int{1, 0}
}

func (m *GetBomReply_BomDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBomReply_BomDetail.Unmarshal(m, b)
}
func (m *GetBomReply_BomDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBomReply_BomDetail.Marshal(b, m, deterministic)
}
func (m *GetBomReply_BomDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBomReply_BomDetail.Merge(m, src)
}
func (m *GetBomReply_BomDetail) XXX_Size() int {
	return xxx_messageInfo_GetBomReply_BomDetail.Size(m)
}
func (m *GetBomReply_BomDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBomReply_BomDetail.DiscardUnknown(m)
}

var xxx_messageInfo_GetBomReply_BomDetail proto.InternalMessageInfo

func (m *GetBomReply_BomDetail) GetMaterialProductId() string {
	if m != nil {
		return m.MaterialProductId
	}
	return ""
}

func (m *GetBomReply_BomDetail) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *GetBomReply_BomDetail) GetQty() float32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *GetBomReply_BomDetail) GetErrorMargin() float32 {
	if m != nil {
		return m.ErrorMargin
	}
	return 0
}

func (m *GetBomReply_BomDetail) GetSpecifiedRecipeId() string {
	if m != nil {
		return m.SpecifiedRecipeId
	}
	return ""
}

func (m *GetBomReply_BomDetail) GetAlternativeMaterials() string {
	if m != nil {
		return m.AlternativeMaterials
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBomRequest)(nil), "proto.mes.GetBomRequest")
	proto.RegisterType((*GetBomReply)(nil), "proto.mes.GetBomReply")
	proto.RegisterType((*GetBomReply_BomDetail)(nil), "proto.mes.GetBomReply.BomDetail")
}

func init() { proto.RegisterFile("get_bom_materials.proto", fileDescriptor_7d547b7fe597de32) }

var fileDescriptor_7d547b7fe597de32 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x95, 0x96, 0x4d, 0xc4, 0x01, 0xc4, 0xcc, 0x60, 0xd5, 0xb8, 0x94, 0x9d, 0x7a, 0xb2,
	0xa5, 0x4e, 0x13, 0x67, 0xaa, 0x01, 0xea, 0x61, 0x08, 0xf9, 0xc8, 0x25, 0x72, 0xe3, 0x67, 0x91,
	0x85, 0x1f, 0x3b, 0xb5, 0xdd, 0xa2, 0x7e, 0x39, 0xbe, 0x07, 0x9f, 0x81, 0x2f, 0x81, 0x62, 0x27,
	0xed, 0x38, 0xe5, 0xf1, 0xff, 0x45, 0x4f, 0xfc, 0x93, 0xc9, 0x55, 0x0b, 0xb1, 0xde, 0x38, 0xac,
	0x51, 0x46, 0xf0, 0x5a, 0x9a, 0xc0, 0x3a, 0xef, 0xa2, 0xa3, 0x65, 0xfa, 0x30, 0x84, 0x70, 0x7d,
	0x95, 0x46, 0xae, 0x90, 0xfb, 0xc0, 0xc1, 0xee, 0x70, 0xc8, 0xdc, 0xfc, 0x29, 0xc8, 0xcb, 0xaf,
	0x10, 0x57, 0x0e, 0x05, 0x6c, 0x77, 0x10, 0x22, 0xbd, 0x23, 0xe4, 0x51, 0x36, 0xd1, 0xf9, 0x43,
	0xad, 0xd5, 0xac, 0x98, 0x17, 0x8b, 0x57, 0xcb, 0x77, 0x39, 0xcd, 0x14, 0x32, 0x1f, 0xd8, 0x97,
	0x6c, 0xaf, 0xef, 0x45, 0x39, 0x24, 0xd7, 0x8a, 0x7e, 0x20, 0x2f, 0x60, 0xbb, 0xd3, 0x1d, 0x82,
	0x8d, 0x7d, 0x71, 0x32, 0x2f, 0x16, 0xa5, 0xa8, 0x8e, 0xda, 0x5a, 0xd1, 0xf7, 0xa4, 0xf4, 0xd0,
	0xe8, 0x0e, 0x7a, 0x7f, 0x9a, 0xfc, 0xe7, 0x59, 0xc8, 0xfd, 0xce, 0xbb, 0x06, 0x42, 0xa8, 0xad,
	0x44, 0x98, 0x3d, 0xcb, 0xfd, 0x41, 0xfb, 0x26, 0x11, 0x28, 0x23, 0x6f, 0x86, 0xfe, 0x7f, 0x9b,
	0xce, 0x52, 0xf2, 0x22, 0x5b, 0x9f, 0x4f, 0xfb, 0x6e, 0x7e, 0x4f, 0x48, 0x35, 0xde, 0xad, 0x33,
	0x07, 0xfa, 0x89, 0x54, 0x3d, 0x26, 0x05, 0x51, 0x6a, 0x13, 0x66, 0xc5, 0x7c, 0xba, 0xa8, 0x96,
	0x73, 0x76, 0xa4, 0xc4, 0x9e, 0x84, 0xd9, 0xca, 0xe1, 0x7d, 0x0a, 0x0a, 0xb2, 0x19, 0xc7, 0x70,
	0xfd, 0xb7, 0x20, 0xe5, 0xd1, 0xe9, 0x7f, 0x68, 0x64, 0x5e, 0x77, 0xde, 0xa9, 0x5d, 0x13, 0x47,
	0x66, 0xa5, 0xb8, 0x18, 0xad, 0xef, 0xd9, 0x59, 0x2b, 0x7a, 0x49, 0xce, 0x0c, 0xec, 0xc1, 0x0c,
	0x70, 0xf2, 0x81, 0xbe, 0x26, 0xd3, 0x6d, 0x3c, 0x24, 0x20, 0x13, 0xd1, 0x8f, 0x89, 0xa5, 0xf7,
	0xce, 0xd7, 0x28, 0x7d, 0xab, 0x6d, 0x62, 0x31, 0x11, 0x55, 0xd2, 0x1e, 0x92, 0xd4, 0xaf, 0x0e,
	0x1d, 0x34, 0xfa, 0x51, 0x83, 0xaa, 0x4f, 0x54, 0x07, 0x16, 0x47, 0x4b, 0x8c, 0x78, 0x6f, 0xc9,
	0x5b, 0x69, 0x22, 0x78, 0x2b, 0xa3, 0xde, 0xc3, 0xe9, 0xa9, 0xcc, 0xce, 0x53, 0xe3, 0xf2, 0x89,
	0xf9, 0x30, 0x7a, 0xab, 0x8f, 0x3f, 0xee, 0x5a, 0x1d, 0x8d, 0xdc, 0xb0, 0x9f, 0x60, 0x95, 0x64,
	0x8d, 0x43, 0x16, 0x7f, 0xf1, 0x74, 0xe0, 0x8d, 0x43, 0x74, 0x36, 0xf0, 0xfd, 0x92, 0xe7, 0xc7,
	0xd5, 0x3a, 0x23, 0x6d, 0xcb, 0x11, 0xc2, 0xe6, 0x3c, 0x29, 0xb7, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x29, 0x0c, 0x44, 0xbc, 0x9b, 0x02, 0x00, 0x00,
}
