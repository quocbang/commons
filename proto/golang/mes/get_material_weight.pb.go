// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_material_weight.proto

package mes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	rs "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	container "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs/container"
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

// 取得BOM清單請求
type GetMaterialWeightRequest struct {
	// 廠別
	FactoryId rs.FactoryID `protobuf:"varint,1,opt,name=factory_id,json=factoryId,proto3,enum=proto.dm.rs.FactoryID" json:"factory_id,omitempty"`
	// 配合表ID
	RecipeId string `protobuf:"bytes,2,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	// 生產機台
	ProductEquipmentId string `protobuf:"bytes,3,opt,name=product_equipment_id,json=productEquipmentId,proto3" json:"product_equipment_id,omitempty"`
	// 設定機台 (指下工程所需機台, 只有密煉使用)
	SetEquipmentId string `protobuf:"bytes,4,opt,name=set_equipment_id,json=setEquipmentId,proto3" json:"set_equipment_id,omitempty"`
	// 途程名稱
	ProcessName          string   `protobuf:"bytes,5,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMaterialWeightRequest) Reset()         { *m = GetMaterialWeightRequest{} }
func (m *GetMaterialWeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetMaterialWeightRequest) ProtoMessage()    {}
func (*GetMaterialWeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e91f559b3f2269f, []int{0}
}

func (m *GetMaterialWeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMaterialWeightRequest.Unmarshal(m, b)
}
func (m *GetMaterialWeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMaterialWeightRequest.Marshal(b, m, deterministic)
}
func (m *GetMaterialWeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMaterialWeightRequest.Merge(m, src)
}
func (m *GetMaterialWeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetMaterialWeightRequest.Size(m)
}
func (m *GetMaterialWeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMaterialWeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMaterialWeightRequest proto.InternalMessageInfo

func (m *GetMaterialWeightRequest) GetFactoryId() rs.FactoryID {
	if m != nil {
		return m.FactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

func (m *GetMaterialWeightRequest) GetRecipeId() string {
	if m != nil {
		return m.RecipeId
	}
	return ""
}

func (m *GetMaterialWeightRequest) GetProductEquipmentId() string {
	if m != nil {
		return m.ProductEquipmentId
	}
	return ""
}

func (m *GetMaterialWeightRequest) GetSetEquipmentId() string {
	if m != nil {
		return m.SetEquipmentId
	}
	return ""
}

func (m *GetMaterialWeightRequest) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

// 取得BOM清單回應
type GetMaterialWeightReply struct {
	// 材料清單
	Materials            []*GetMaterialWeightReply_Material `protobuf:"bytes,2,rep,name=materials,proto3" json:"materials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GetMaterialWeightReply) Reset()         { *m = GetMaterialWeightReply{} }
func (m *GetMaterialWeightReply) String() string { return proto.CompactTextString(m) }
func (*GetMaterialWeightReply) ProtoMessage()    {}
func (*GetMaterialWeightReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e91f559b3f2269f, []int{1}
}

func (m *GetMaterialWeightReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMaterialWeightReply.Unmarshal(m, b)
}
func (m *GetMaterialWeightReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMaterialWeightReply.Marshal(b, m, deterministic)
}
func (m *GetMaterialWeightReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMaterialWeightReply.Merge(m, src)
}
func (m *GetMaterialWeightReply) XXX_Size() int {
	return xxx_messageInfo_GetMaterialWeightReply.Size(m)
}
func (m *GetMaterialWeightReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMaterialWeightReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMaterialWeightReply proto.InternalMessageInfo

func (m *GetMaterialWeightReply) GetMaterials() []*GetMaterialWeightReply_Material {
	if m != nil {
		return m.Materials
	}
	return nil
}

// 材料
type GetMaterialWeightReply_Material struct {
	// 材料類型
	Type rs.ProductType `protobuf:"varint,1,opt,name=type,proto3,enum=proto.dm.rs.ProductType" json:"type,omitempty"`
	// 材料ID
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// 材料等級
	Level string `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	// 標準耗用量
	Value float32 `protobuf:"fixed32,4,opt,name=value,proto3" json:"value,omitempty"`
	// 計量單位 : 重量、長度、個數
	Unit string `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	// 允許誤差
	ErrorMargin float32 `protobuf:"fixed32,6,opt,name=error_margin,json=errorMargin,proto3" json:"error_margin,omitempty"`
	// 桶槽類型 : 有指定桶槽代表為中央供應料
	ContainerType container.Type `protobuf:"varint,7,opt,name=containerType,proto3,enum=proto.dm.rs.container.Type" json:"containerType,omitempty"`
	// 指定recipe id (如密煉使用之藥包需指定特定藥包配合表版本)
	SpecifiedRecipeId    string   `protobuf:"bytes,8,opt,name=specified_recipe_id,json=specifiedRecipeId,proto3" json:"specified_recipe_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMaterialWeightReply_Material) Reset()         { *m = GetMaterialWeightReply_Material{} }
func (m *GetMaterialWeightReply_Material) String() string { return proto.CompactTextString(m) }
func (*GetMaterialWeightReply_Material) ProtoMessage()    {}
func (*GetMaterialWeightReply_Material) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e91f559b3f2269f, []int{1, 0}
}

func (m *GetMaterialWeightReply_Material) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMaterialWeightReply_Material.Unmarshal(m, b)
}
func (m *GetMaterialWeightReply_Material) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMaterialWeightReply_Material.Marshal(b, m, deterministic)
}
func (m *GetMaterialWeightReply_Material) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMaterialWeightReply_Material.Merge(m, src)
}
func (m *GetMaterialWeightReply_Material) XXX_Size() int {
	return xxx_messageInfo_GetMaterialWeightReply_Material.Size(m)
}
func (m *GetMaterialWeightReply_Material) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMaterialWeightReply_Material.DiscardUnknown(m)
}

var xxx_messageInfo_GetMaterialWeightReply_Material proto.InternalMessageInfo

func (m *GetMaterialWeightReply_Material) GetType() rs.ProductType {
	if m != nil {
		return m.Type
	}
	return rs.ProductType_PRODUCT_TYPE_UNSPECIFIED
}

func (m *GetMaterialWeightReply_Material) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetMaterialWeightReply_Material) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *GetMaterialWeightReply_Material) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *GetMaterialWeightReply_Material) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *GetMaterialWeightReply_Material) GetErrorMargin() float32 {
	if m != nil {
		return m.ErrorMargin
	}
	return 0
}

func (m *GetMaterialWeightReply_Material) GetContainerType() container.Type {
	if m != nil {
		return m.ContainerType
	}
	return container.Type_CONTAINER_TYPE_UNSPECIFIED
}

func (m *GetMaterialWeightReply_Material) GetSpecifiedRecipeId() string {
	if m != nil {
		return m.SpecifiedRecipeId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMaterialWeightRequest)(nil), "proto.mes.GetMaterialWeightRequest")
	proto.RegisterType((*GetMaterialWeightReply)(nil), "proto.mes.GetMaterialWeightReply")
	proto.RegisterType((*GetMaterialWeightReply_Material)(nil), "proto.mes.GetMaterialWeightReply.Material")
}

func init() { proto.RegisterFile("get_material_weight.proto", fileDescriptor_7e91f559b3f2269f) }

var fileDescriptor_7e91f559b3f2269f = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0xb4, 0x1b, 0xad, 0x0b, 0x15, 0x98, 0x69, 0x84, 0xee, 0xa6, 0x4c, 0x42, 0xaa,
	0x10, 0x72, 0x50, 0xd1, 0xc4, 0x35, 0x88, 0x7f, 0xbd, 0x18, 0x42, 0x11, 0x12, 0x12, 0x37, 0x91,
	0x17, 0x9f, 0x05, 0x8b, 0xf8, 0xcf, 0x6c, 0xa7, 0x53, 0xaf, 0x79, 0x0f, 0xde, 0x90, 0x77, 0x40,
	0x39, 0x4e, 0x33, 0x3a, 0xed, 0x2a, 0xf6, 0x77, 0x7e, 0x9f, 0x7c, 0xce, 0x97, 0x43, 0x9e, 0x56,
	0x10, 0x0a, 0xc5, 0x03, 0x38, 0xc9, 0xeb, 0xe2, 0x1a, 0x64, 0xf5, 0x33, 0x30, 0xeb, 0x4c, 0x30,
	0x74, 0x82, 0x1f, 0xa6, 0xc0, 0xcf, 0x9f, 0xe0, 0x31, 0x13, 0x2a, 0x73, 0x3e, 0x03, 0xdd, 0x28,
	0x1f, 0x99, 0xf9, 0xf3, 0xff, 0x0b, 0xa5, 0xd1, 0x81, 0x4b, 0x0d, 0x0e, 0x91, 0x22, 0x6c, 0x2d,
	0x44, 0xec, 0xf4, 0xef, 0x80, 0xa4, 0x9f, 0x20, 0x9c, 0x77, 0xef, 0x7c, 0xc7, 0x67, 0x72, 0xb8,
	0x6a, 0xc0, 0x07, 0x7a, 0x46, 0xc8, 0x25, 0x2f, 0x83, 0x71, 0xdb, 0x42, 0x8a, 0x74, 0xb0, 0x18,
	0x2c, 0x67, 0xab, 0xe3, 0x68, 0x64, 0x42, 0x31, 0xe7, 0xd9, 0xc7, 0x58, 0x5e, 0xbf, 0xcf, 0x27,
	0x1d, 0xb9, 0x16, 0xf4, 0x84, 0x4c, 0x1c, 0x94, 0xd2, 0x42, 0xeb, 0x4a, 0x16, 0x83, 0xe5, 0x24,
	0x1f, 0x47, 0x61, 0x2d, 0xe8, 0x2b, 0x72, 0x64, 0x9d, 0x11, 0x4d, 0x19, 0x0a, 0xb8, 0x6a, 0xa4,
	0x55, 0xa0, 0x43, 0xcb, 0x0d, 0x91, 0xa3, 0x5d, 0xed, 0xc3, 0xae, 0xb4, 0x16, 0x74, 0x49, 0x1e,
	0x7a, 0xb8, 0x45, 0x8f, 0x90, 0x9e, 0x79, 0xd8, 0x23, 0x9f, 0x91, 0xfb, 0xd6, 0x99, 0x12, 0xbc,
	0x2f, 0x34, 0x57, 0x90, 0x1e, 0x20, 0x35, 0xed, 0xb4, 0x2f, 0x5c, 0xc1, 0xe9, 0xef, 0x21, 0x39,
	0xbe, 0x63, 0x5e, 0x5b, 0x6f, 0xe9, 0x67, 0x32, 0xd9, 0xc5, 0xed, 0xd3, 0x64, 0x31, 0x5c, 0x4e,
	0x57, 0x2f, 0x58, 0x9f, 0x34, 0xbb, 0xdb, 0xc5, 0x76, 0x5a, 0x7e, 0x63, 0x9e, 0xff, 0x49, 0xc8,
	0x78, 0xa7, 0xd3, 0x97, 0x64, 0xd4, 0xe6, 0xdd, 0xc5, 0x97, 0xee, 0xc5, 0xf7, 0x35, 0x4e, 0xfb,
	0x6d, 0x6b, 0x21, 0x47, 0x8a, 0xce, 0x48, 0xd2, 0x87, 0x96, 0x48, 0x41, 0x8f, 0xc8, 0x41, 0x0d,
	0x1b, 0xa8, 0xbb, 0x7c, 0xe2, 0xa5, 0x55, 0x37, 0xbc, 0x6e, 0x00, 0x73, 0x48, 0xf2, 0x78, 0xa1,
	0x94, 0x8c, 0x1a, 0x2d, 0x43, 0x37, 0x36, 0x9e, 0xdb, 0x48, 0xc0, 0x39, 0xe3, 0x0a, 0xc5, 0x5d,
	0x25, 0x75, 0x7a, 0x88, 0x86, 0x29, 0x6a, 0xe7, 0x28, 0xd1, 0xb7, 0xe4, 0x41, 0xbf, 0x1f, 0x6d,
	0x27, 0xe9, 0x3d, 0xec, 0xf4, 0x64, 0xaf, 0xd3, 0x9e, 0x60, 0xd8, 0xec, 0xbe, 0x83, 0x32, 0xf2,
	0xd8, 0x5b, 0x28, 0xe5, 0xa5, 0x04, 0x51, 0xdc, 0xfc, 0xfb, 0x31, 0x36, 0xf2, 0xa8, 0x2f, 0xe5,
	0xdd, 0x12, 0xbc, 0x7b, 0xf3, 0xe3, 0xac, 0x92, 0xa1, 0xe6, 0x17, 0xec, 0x17, 0x68, 0xc1, 0x59,
	0x69, 0x14, 0x0b, 0xd7, 0x19, 0x5e, 0xb2, 0xd2, 0x28, 0x65, 0xb4, 0xcf, 0x36, 0xab, 0x2c, 0xee,
	0x70, 0x65, 0x6a, 0xae, 0xab, 0x4c, 0x81, 0xbf, 0x38, 0x44, 0xe5, 0xf5, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0x4b, 0x86, 0x09, 0x1d, 0x03, 0x00, 0x00,
}
