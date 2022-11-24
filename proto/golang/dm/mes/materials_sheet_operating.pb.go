// Code generated by protoc-gen-go. DO NOT EDIT.
// source: materials_sheet_operating.proto

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

// 查詢
type GetMaterialsSheetRequest struct {
	LotId                string   `protobuf:"bytes,1,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	ShtNo                string   `protobuf:"bytes,2,opt,name=sht_no,json=shtNo,proto3" json:"sht_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMaterialsSheetRequest) Reset()         { *m = GetMaterialsSheetRequest{} }
func (m *GetMaterialsSheetRequest) String() string { return proto.CompactTextString(m) }
func (*GetMaterialsSheetRequest) ProtoMessage()    {}
func (*GetMaterialsSheetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_566077b4d2a2f202, []int{0}
}

func (m *GetMaterialsSheetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMaterialsSheetRequest.Unmarshal(m, b)
}
func (m *GetMaterialsSheetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMaterialsSheetRequest.Marshal(b, m, deterministic)
}
func (m *GetMaterialsSheetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMaterialsSheetRequest.Merge(m, src)
}
func (m *GetMaterialsSheetRequest) XXX_Size() int {
	return xxx_messageInfo_GetMaterialsSheetRequest.Size(m)
}
func (m *GetMaterialsSheetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMaterialsSheetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMaterialsSheetRequest proto.InternalMessageInfo

func (m *GetMaterialsSheetRequest) GetLotId() string {
	if m != nil {
		return m.LotId
	}
	return ""
}

func (m *GetMaterialsSheetRequest) GetShtNo() string {
	if m != nil {
		return m.ShtNo
	}
	return ""
}

type GetMaterialsSheetReply struct {
	Mtrls                []*MtrlSheetInfo `protobuf:"bytes,1,rep,name=mtrls,proto3" json:"mtrls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetMaterialsSheetReply) Reset()         { *m = GetMaterialsSheetReply{} }
func (m *GetMaterialsSheetReply) String() string { return proto.CompactTextString(m) }
func (*GetMaterialsSheetReply) ProtoMessage()    {}
func (*GetMaterialsSheetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_566077b4d2a2f202, []int{1}
}

func (m *GetMaterialsSheetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMaterialsSheetReply.Unmarshal(m, b)
}
func (m *GetMaterialsSheetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMaterialsSheetReply.Marshal(b, m, deterministic)
}
func (m *GetMaterialsSheetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMaterialsSheetReply.Merge(m, src)
}
func (m *GetMaterialsSheetReply) XXX_Size() int {
	return xxx_messageInfo_GetMaterialsSheetReply.Size(m)
}
func (m *GetMaterialsSheetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMaterialsSheetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMaterialsSheetReply proto.InternalMessageInfo

func (m *GetMaterialsSheetReply) GetMtrls() []*MtrlSheetInfo {
	if m != nil {
		return m.Mtrls
	}
	return nil
}

// 新增
type AddMaterialsSheetRequest struct {
	MtrlLotId            string   `protobuf:"bytes,1,opt,name=mtrl_lot_id,json=mtrlLotId,proto3" json:"mtrl_lot_id,omitempty"`
	MtrlProductId        string   `protobuf:"bytes,2,opt,name=mtrl_product_id,json=mtrlProductId,proto3" json:"mtrl_product_id,omitempty"`
	MtrlLevel            string   `protobuf:"bytes,3,opt,name=mtrl_level,json=mtrlLevel,proto3" json:"mtrl_level,omitempty"`
	LotId                string   `protobuf:"bytes,4,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	ShtNo                string   `protobuf:"bytes,5,opt,name=sht_no,json=shtNo,proto3" json:"sht_no,omitempty"`
	Qty                  float32  `protobuf:"fixed32,6,opt,name=qty,proto3" json:"qty,omitempty"`
	Unit                 string   `protobuf:"bytes,7,opt,name=unit,proto3" json:"unit,omitempty"`
	UpdateAt             string   `protobuf:"bytes,8,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	UpdateBy             string   `protobuf:"bytes,9,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMaterialsSheetRequest) Reset()         { *m = AddMaterialsSheetRequest{} }
func (m *AddMaterialsSheetRequest) String() string { return proto.CompactTextString(m) }
func (*AddMaterialsSheetRequest) ProtoMessage()    {}
func (*AddMaterialsSheetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_566077b4d2a2f202, []int{2}
}

func (m *AddMaterialsSheetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMaterialsSheetRequest.Unmarshal(m, b)
}
func (m *AddMaterialsSheetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMaterialsSheetRequest.Marshal(b, m, deterministic)
}
func (m *AddMaterialsSheetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMaterialsSheetRequest.Merge(m, src)
}
func (m *AddMaterialsSheetRequest) XXX_Size() int {
	return xxx_messageInfo_AddMaterialsSheetRequest.Size(m)
}
func (m *AddMaterialsSheetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMaterialsSheetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMaterialsSheetRequest proto.InternalMessageInfo

func (m *AddMaterialsSheetRequest) GetMtrlLotId() string {
	if m != nil {
		return m.MtrlLotId
	}
	return ""
}

func (m *AddMaterialsSheetRequest) GetMtrlProductId() string {
	if m != nil {
		return m.MtrlProductId
	}
	return ""
}

func (m *AddMaterialsSheetRequest) GetMtrlLevel() string {
	if m != nil {
		return m.MtrlLevel
	}
	return ""
}

func (m *AddMaterialsSheetRequest) GetLotId() string {
	if m != nil {
		return m.LotId
	}
	return ""
}

func (m *AddMaterialsSheetRequest) GetShtNo() string {
	if m != nil {
		return m.ShtNo
	}
	return ""
}

func (m *AddMaterialsSheetRequest) GetQty() float32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *AddMaterialsSheetRequest) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *AddMaterialsSheetRequest) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func (m *AddMaterialsSheetRequest) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

type AddMaterialsSheetReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMaterialsSheetReply) Reset()         { *m = AddMaterialsSheetReply{} }
func (m *AddMaterialsSheetReply) String() string { return proto.CompactTextString(m) }
func (*AddMaterialsSheetReply) ProtoMessage()    {}
func (*AddMaterialsSheetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_566077b4d2a2f202, []int{3}
}

func (m *AddMaterialsSheetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMaterialsSheetReply.Unmarshal(m, b)
}
func (m *AddMaterialsSheetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMaterialsSheetReply.Marshal(b, m, deterministic)
}
func (m *AddMaterialsSheetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMaterialsSheetReply.Merge(m, src)
}
func (m *AddMaterialsSheetReply) XXX_Size() int {
	return xxx_messageInfo_AddMaterialsSheetReply.Size(m)
}
func (m *AddMaterialsSheetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMaterialsSheetReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddMaterialsSheetReply proto.InternalMessageInfo

type MtrlSheetInfo struct {
	MtrlLotId            string   `protobuf:"bytes,1,opt,name=mtrl_lot_id,json=mtrlLotId,proto3" json:"mtrl_lot_id,omitempty"`
	MtrlProductId        string   `protobuf:"bytes,2,opt,name=mtrl_product_id,json=mtrlProductId,proto3" json:"mtrl_product_id,omitempty"`
	MtrlLevel            string   `protobuf:"bytes,3,opt,name=mtrl_level,json=mtrlLevel,proto3" json:"mtrl_level,omitempty"`
	LotId                string   `protobuf:"bytes,4,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	ShtNo                string   `protobuf:"bytes,5,opt,name=sht_no,json=shtNo,proto3" json:"sht_no,omitempty"`
	Qty                  float32  `protobuf:"fixed32,6,opt,name=qty,proto3" json:"qty,omitempty"`
	Unit                 string   `protobuf:"bytes,7,opt,name=unit,proto3" json:"unit,omitempty"`
	UpdateAt             string   `protobuf:"bytes,8,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	UpdateBy             string   `protobuf:"bytes,9,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MtrlSheetInfo) Reset()         { *m = MtrlSheetInfo{} }
func (m *MtrlSheetInfo) String() string { return proto.CompactTextString(m) }
func (*MtrlSheetInfo) ProtoMessage()    {}
func (*MtrlSheetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_566077b4d2a2f202, []int{4}
}

func (m *MtrlSheetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MtrlSheetInfo.Unmarshal(m, b)
}
func (m *MtrlSheetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MtrlSheetInfo.Marshal(b, m, deterministic)
}
func (m *MtrlSheetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MtrlSheetInfo.Merge(m, src)
}
func (m *MtrlSheetInfo) XXX_Size() int {
	return xxx_messageInfo_MtrlSheetInfo.Size(m)
}
func (m *MtrlSheetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MtrlSheetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MtrlSheetInfo proto.InternalMessageInfo

func (m *MtrlSheetInfo) GetMtrlLotId() string {
	if m != nil {
		return m.MtrlLotId
	}
	return ""
}

func (m *MtrlSheetInfo) GetMtrlProductId() string {
	if m != nil {
		return m.MtrlProductId
	}
	return ""
}

func (m *MtrlSheetInfo) GetMtrlLevel() string {
	if m != nil {
		return m.MtrlLevel
	}
	return ""
}

func (m *MtrlSheetInfo) GetLotId() string {
	if m != nil {
		return m.LotId
	}
	return ""
}

func (m *MtrlSheetInfo) GetShtNo() string {
	if m != nil {
		return m.ShtNo
	}
	return ""
}

func (m *MtrlSheetInfo) GetQty() float32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *MtrlSheetInfo) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *MtrlSheetInfo) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func (m *MtrlSheetInfo) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMaterialsSheetRequest)(nil), "proto.dm.mes.GetMaterialsSheetRequest")
	proto.RegisterType((*GetMaterialsSheetReply)(nil), "proto.dm.mes.GetMaterialsSheetReply")
	proto.RegisterType((*AddMaterialsSheetRequest)(nil), "proto.dm.mes.AddMaterialsSheetRequest")
	proto.RegisterType((*AddMaterialsSheetReply)(nil), "proto.dm.mes.AddMaterialsSheetReply")
	proto.RegisterType((*MtrlSheetInfo)(nil), "proto.dm.mes.MtrlSheetInfo")
}

func init() { proto.RegisterFile("materials_sheet_operating.proto", fileDescriptor_566077b4d2a2f202) }

var fileDescriptor_566077b4d2a2f202 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x49, 0xfa, 0xe7, 0x6d, 0xb6, 0x6f, 0x51, 0x16, 0x2c, 0x0b, 0x45, 0x2d, 0x39, 0x48,
	0x4f, 0x1b, 0xac, 0x17, 0xf1, 0xd6, 0x5e, 0xb4, 0x68, 0x45, 0xe2, 0xcd, 0x4b, 0xd8, 0x76, 0xd7,
	0x34, 0xb8, 0x9b, 0x4d, 0xb3, 0x93, 0x4a, 0xee, 0xde, 0xfd, 0xca, 0x92, 0x0d, 0x42, 0x0a, 0xf9,
	0x06, 0x9e, 0x32, 0xf9, 0x3d, 0xcf, 0x0c, 0xcf, 0x0c, 0x8b, 0x2e, 0x15, 0x03, 0x91, 0x27, 0x4c,
	0x9a, 0xc8, 0xec, 0x84, 0x80, 0x48, 0x67, 0x22, 0x67, 0x90, 0xa4, 0x31, 0xcd, 0x72, 0x0d, 0x1a,
	0xff, 0xb7, 0x1f, 0xca, 0x15, 0x55, 0xc2, 0xf8, 0x0f, 0x88, 0xdc, 0x0b, 0x58, 0xff, 0xf6, 0xbc,
	0x56, 0x2d, 0xa1, 0xd8, 0x17, 0xc2, 0x00, 0x3e, 0x43, 0x7d, 0xa9, 0x21, 0x4a, 0x38, 0x71, 0xa6,
	0xce, 0xcc, 0x0b, 0x7b, 0x52, 0xc3, 0x8a, 0x57, 0xd8, 0xec, 0x20, 0x4a, 0x35, 0x71, 0x6b, 0x6c,
	0x76, 0xf0, 0xac, 0xfd, 0x47, 0x34, 0x6e, 0x99, 0x94, 0xc9, 0x12, 0x5f, 0xa3, 0x9e, 0x82, 0x5c,
	0x1a, 0xe2, 0x4c, 0x3b, 0xb3, 0xe1, 0x7c, 0x42, 0x9b, 0x09, 0xe8, 0x1a, 0x72, 0x69, 0xcd, 0xab,
	0xf4, 0x5d, 0x87, 0xb5, 0xd3, 0xff, 0x76, 0x11, 0x59, 0x70, 0xde, 0x9e, 0xeb, 0x02, 0x0d, 0x2b,
	0x57, 0x74, 0x14, 0xce, 0xab, 0xd0, 0x93, 0x0d, 0x78, 0x85, 0x4e, 0xac, 0x9e, 0xe5, 0x9a, 0x17,
	0x5b, 0xeb, 0xa9, 0x93, 0x8e, 0x2a, 0xfc, 0x52, 0xd3, 0x15, 0xc7, 0xe7, 0x08, 0xd5, 0x73, 0xc4,
	0x41, 0x48, 0xd2, 0x69, 0x8c, 0xa9, 0x40, 0x63, 0xfd, 0x6e, 0xfb, 0xfa, 0xbd, 0xc6, 0xfa, 0xf8,
	0x14, 0x75, 0xf6, 0x50, 0x92, 0xfe, 0xd4, 0x99, 0xb9, 0x61, 0x55, 0x62, 0x8c, 0xba, 0x45, 0x9a,
	0x00, 0xf9, 0x67, 0x6d, 0xb6, 0xc6, 0x13, 0xe4, 0x15, 0x19, 0x67, 0x20, 0x22, 0x06, 0x64, 0x60,
	0x85, 0x41, 0x0d, 0x16, 0x4d, 0x71, 0x53, 0x12, 0xaf, 0x29, 0x2e, 0x4b, 0x9f, 0xa0, 0x71, 0xcb,
	0x41, 0x32, 0x59, 0xfa, 0x5f, 0x2e, 0x1a, 0x1d, 0x1d, 0xf1, 0x4f, 0x1e, 0x68, 0x79, 0xf7, 0x76,
	0x1b, 0x27, 0x20, 0xd9, 0x86, 0x7e, 0x88, 0x94, 0x33, 0xba, 0xd5, 0x8a, 0xc2, 0x67, 0x60, 0x7f,
	0x82, 0xad, 0x56, 0x4a, 0xa7, 0x26, 0x38, 0xcc, 0x03, 0xfb, 0xf8, 0x82, 0x58, 0x4b, 0x96, 0xc6,
	0x01, 0x57, 0x81, 0x12, 0x66, 0xd3, 0xb7, 0xf0, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x97, 0xf4,
	0xcd, 0x9a, 0x3d, 0x03, 0x00, 0x00,
}