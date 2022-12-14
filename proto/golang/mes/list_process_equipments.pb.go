// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_process_equipments.proto

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

type ListProcessEquipmentsRequest struct {
	RecipeId string `protobuf:"bytes,1,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	// 途程名稱
	ProcessName          string         `protobuf:"bytes,2,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	ProcessType          rs.ProcessType `protobuf:"varint,3,opt,name=process_type,json=processType,proto3,enum=proto.dm.rs.ProcessType" json:"process_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListProcessEquipmentsRequest) Reset()         { *m = ListProcessEquipmentsRequest{} }
func (m *ListProcessEquipmentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListProcessEquipmentsRequest) ProtoMessage()    {}
func (*ListProcessEquipmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb195a4c97408bb9, []int{0}
}

func (m *ListProcessEquipmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProcessEquipmentsRequest.Unmarshal(m, b)
}
func (m *ListProcessEquipmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProcessEquipmentsRequest.Marshal(b, m, deterministic)
}
func (m *ListProcessEquipmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProcessEquipmentsRequest.Merge(m, src)
}
func (m *ListProcessEquipmentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListProcessEquipmentsRequest.Size(m)
}
func (m *ListProcessEquipmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProcessEquipmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProcessEquipmentsRequest proto.InternalMessageInfo

func (m *ListProcessEquipmentsRequest) GetRecipeId() string {
	if m != nil {
		return m.RecipeId
	}
	return ""
}

func (m *ListProcessEquipmentsRequest) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *ListProcessEquipmentsRequest) GetProcessType() rs.ProcessType {
	if m != nil {
		return m.ProcessType
	}
	return rs.ProcessType_PROCESS_TYPE_UNSPECIFIED
}

type ListProcessEquipmentsReply struct {
	ProductEquipmentIds  []string `protobuf:"bytes,1,rep,name=product_equipment_ids,json=productEquipmentIds,proto3" json:"product_equipment_ids,omitempty"`
	SettingEquipmentIds  []string `protobuf:"bytes,2,rep,name=setting_equipment_ids,json=settingEquipmentIds,proto3" json:"setting_equipment_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProcessEquipmentsReply) Reset()         { *m = ListProcessEquipmentsReply{} }
func (m *ListProcessEquipmentsReply) String() string { return proto.CompactTextString(m) }
func (*ListProcessEquipmentsReply) ProtoMessage()    {}
func (*ListProcessEquipmentsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb195a4c97408bb9, []int{1}
}

func (m *ListProcessEquipmentsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProcessEquipmentsReply.Unmarshal(m, b)
}
func (m *ListProcessEquipmentsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProcessEquipmentsReply.Marshal(b, m, deterministic)
}
func (m *ListProcessEquipmentsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProcessEquipmentsReply.Merge(m, src)
}
func (m *ListProcessEquipmentsReply) XXX_Size() int {
	return xxx_messageInfo_ListProcessEquipmentsReply.Size(m)
}
func (m *ListProcessEquipmentsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProcessEquipmentsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListProcessEquipmentsReply proto.InternalMessageInfo

func (m *ListProcessEquipmentsReply) GetProductEquipmentIds() []string {
	if m != nil {
		return m.ProductEquipmentIds
	}
	return nil
}

func (m *ListProcessEquipmentsReply) GetSettingEquipmentIds() []string {
	if m != nil {
		return m.SettingEquipmentIds
	}
	return nil
}

func init() {
	proto.RegisterType((*ListProcessEquipmentsRequest)(nil), "proto.mes.ListProcessEquipmentsRequest")
	proto.RegisterType((*ListProcessEquipmentsReply)(nil), "proto.mes.ListProcessEquipmentsReply")
}

func init() { proto.RegisterFile("list_process_equipments.proto", fileDescriptor_bb195a4c97408bb9) }

var fileDescriptor_bb195a4c97408bb9 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xf4, 0x30,
	0x10, 0xc6, 0xc9, 0x2e, 0xbc, 0xbc, 0x8d, 0xe2, 0xa1, 0x22, 0x96, 0x55, 0x61, 0xdd, 0x53, 0x4f,
	0x09, 0x54, 0xc4, 0x83, 0x37, 0xc1, 0xc3, 0x82, 0x88, 0x14, 0x4f, 0x5e, 0x4a, 0xb7, 0x19, 0x4a,
	0xb0, 0xf9, 0xb3, 0x99, 0xa9, 0xd2, 0xbb, 0xdf, 0xc2, 0x2f, 0x2b, 0xdb, 0xd6, 0xae, 0x08, 0x9e,
	0x92, 0x79, 0x9e, 0xf9, 0xcd, 0x93, 0x0c, 0xbf, 0x68, 0x34, 0x52, 0xe1, 0x83, 0xab, 0x00, 0xb1,
	0x80, 0x6d, 0xab, 0xbd, 0x01, 0x4b, 0x28, 0x7c, 0x70, 0xe4, 0xe2, 0xa8, 0x3f, 0x84, 0x01, 0x5c,
	0x9c, 0xf6, 0x57, 0xa9, 0x8c, 0x0c, 0x28, 0xc1, 0xb6, 0x66, 0xec, 0x59, 0x7d, 0x32, 0x7e, 0xfe,
	0xa0, 0x91, 0x9e, 0x86, 0x21, 0xf7, 0xd3, 0x8c, 0x1c, 0xb6, 0x2d, 0x20, 0xc5, 0x67, 0x3c, 0x0a,
	0x50, 0x69, 0x0f, 0x85, 0x56, 0x09, 0x5b, 0xb2, 0x34, 0xca, 0xff, 0x0f, 0xc2, 0x5a, 0xc5, 0x97,
	0xfc, 0xf0, 0x3b, 0xdd, 0x96, 0x06, 0x92, 0x59, 0xef, 0x1f, 0x8c, 0xda, 0x63, 0x69, 0x20, 0xbe,
	0xdd, 0xb7, 0x50, 0xe7, 0x21, 0x99, 0x2f, 0x59, 0x7a, 0x94, 0x25, 0x43, 0xbc, 0x50, 0x46, 0x04,
	0x14, 0x63, 0xf8, 0x73, 0xe7, 0x61, 0x82, 0x77, 0xc5, 0xea, 0x83, 0xf1, 0xc5, 0x1f, 0xaf, 0xf3,
	0x4d, 0x17, 0x67, 0xfc, 0xc4, 0x07, 0xa7, 0xda, 0x8a, 0xf6, 0x9f, 0x2f, 0xb4, 0xc2, 0x84, 0x2d,
	0xe7, 0x69, 0x94, 0x1f, 0x8f, 0xe6, 0x84, 0xad, 0x15, 0xee, 0x18, 0x04, 0x22, 0x6d, 0xeb, 0x5f,
	0xcc, 0x6c, 0x60, 0x46, 0xf3, 0x27, 0x73, 0x77, 0xf3, 0x72, 0x5d, 0x6b, 0x6a, 0xca, 0x8d, 0x78,
	0x05, 0xab, 0x4a, 0x51, 0x39, 0x23, 0xe8, 0x5d, 0xf6, 0x85, 0xac, 0x9c, 0x31, 0xce, 0xa2, 0x7c,
	0xcb, 0xe4, 0xb0, 0xe4, 0xda, 0x35, 0xa5, 0xad, 0xa5, 0x01, 0xdc, 0xfc, 0xeb, 0x95, 0xab, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xa7, 0x60, 0xc0, 0xa9, 0x01, 0x00, 0x00,
}
