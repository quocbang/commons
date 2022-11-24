// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction_shipping_operating.proto

package mes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	code "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/code"
	error1 "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/code/error"
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

type MTxShipListOperatingRequest struct {
	ActionType code.FunctionModel `protobuf:"varint,1,opt,name=action_type,json=actionType,proto3,enum=proto.mes.code.FunctionModel" json:"action_type,omitempty"`
	TxId       string             `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	ItemId     string             `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SequenceNo int32              `protobuf:"varint,4,opt,name=sequence_no,json=sequenceNo,proto3" json:"sequence_no,omitempty"`
	//材料 Lot ID, 標示卡編號
	MaterialLotId string `protobuf:"bytes,5,opt,name=material_lot_id,json=materialLotId,proto3" json:"material_lot_id,omitempty"`
	//材料 ID ,材料,半製品,成品編號
	MaterialProductId string `protobuf:"bytes,6,opt,name=material_product_id,json=materialProductId,proto3" json:"material_product_id,omitempty"`
	//材料等級
	MaterialLevel string `protobuf:"bytes,7,opt,name=material_level,json=materialLevel,proto3" json:"material_level,omitempty"`
	//移轉數量
	Qty float32 `protobuf:"fixed32,8,opt,name=qty,proto3" json:"qty,omitempty"`
	//計量單位
	Unit string `protobuf:"bytes,9,opt,name=unit,proto3" json:"unit,omitempty"`
	//更新人
	UpdateBy string `protobuf:"bytes,10,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	//Compare Date
	OnlyToday            bool                `protobuf:"varint,11,opt,name=only_today,json=onlyToday,proto3" json:"only_today,omitempty"`
	Status               code.ShippingStatus `protobuf:"varint,12,opt,name=status,proto3,enum=proto.mes.code.ShippingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MTxShipListOperatingRequest) Reset()         { *m = MTxShipListOperatingRequest{} }
func (m *MTxShipListOperatingRequest) String() string { return proto.CompactTextString(m) }
func (*MTxShipListOperatingRequest) ProtoMessage()    {}
func (*MTxShipListOperatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_095c79d790d164da, []int{0}
}

func (m *MTxShipListOperatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MTxShipListOperatingRequest.Unmarshal(m, b)
}
func (m *MTxShipListOperatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MTxShipListOperatingRequest.Marshal(b, m, deterministic)
}
func (m *MTxShipListOperatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTxShipListOperatingRequest.Merge(m, src)
}
func (m *MTxShipListOperatingRequest) XXX_Size() int {
	return xxx_messageInfo_MTxShipListOperatingRequest.Size(m)
}
func (m *MTxShipListOperatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MTxShipListOperatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MTxShipListOperatingRequest proto.InternalMessageInfo

func (m *MTxShipListOperatingRequest) GetActionType() code.FunctionModel {
	if m != nil {
		return m.ActionType
	}
	return code.FunctionModel_FUNCTION_MODEL_UNSPECIFIED
}

func (m *MTxShipListOperatingRequest) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *MTxShipListOperatingRequest) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *MTxShipListOperatingRequest) GetSequenceNo() int32 {
	if m != nil {
		return m.SequenceNo
	}
	return 0
}

func (m *MTxShipListOperatingRequest) GetMaterialLotId() string {
	if m != nil {
		return m.MaterialLotId
	}
	return ""
}

func (m *MTxShipListOperatingRequest) GetMaterialProductId() string {
	if m != nil {
		return m.MaterialProductId
	}
	return ""
}

func (m *MTxShipListOperatingRequest) GetMaterialLevel() string {
	if m != nil {
		return m.MaterialLevel
	}
	return ""
}

func (m *MTxShipListOperatingRequest) GetQty() float32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *MTxShipListOperatingRequest) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *MTxShipListOperatingRequest) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *MTxShipListOperatingRequest) GetOnlyToday() bool {
	if m != nil {
		return m.OnlyToday
	}
	return false
}

func (m *MTxShipListOperatingRequest) GetStatus() code.ShippingStatus {
	if m != nil {
		return m.Status
	}
	return code.ShippingStatus_SHIPPING_STATUS_UNSPECIFIED
}

type MTxShipListOperatingReply struct {
	ShippingList         []*ShippingInfo             `protobuf:"bytes,1,rep,name=shippingList,proto3" json:"shippingList,omitempty"`
	Err                  error1.MtxShipListOperating `protobuf:"varint,2,opt,name=err,proto3,enum=proto.mes.code.error.MtxShipListOperating" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MTxShipListOperatingReply) Reset()         { *m = MTxShipListOperatingReply{} }
func (m *MTxShipListOperatingReply) String() string { return proto.CompactTextString(m) }
func (*MTxShipListOperatingReply) ProtoMessage()    {}
func (*MTxShipListOperatingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_095c79d790d164da, []int{1}
}

func (m *MTxShipListOperatingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MTxShipListOperatingReply.Unmarshal(m, b)
}
func (m *MTxShipListOperatingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MTxShipListOperatingReply.Marshal(b, m, deterministic)
}
func (m *MTxShipListOperatingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTxShipListOperatingReply.Merge(m, src)
}
func (m *MTxShipListOperatingReply) XXX_Size() int {
	return xxx_messageInfo_MTxShipListOperatingReply.Size(m)
}
func (m *MTxShipListOperatingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MTxShipListOperatingReply.DiscardUnknown(m)
}

var xxx_messageInfo_MTxShipListOperatingReply proto.InternalMessageInfo

func (m *MTxShipListOperatingReply) GetShippingList() []*ShippingInfo {
	if m != nil {
		return m.ShippingList
	}
	return nil
}

func (m *MTxShipListOperatingReply) GetErr() error1.MtxShipListOperating {
	if m != nil {
		return m.Err
	}
	return error1.MtxShipListOperating_MTXSHIPLIST_OPERATING_SUCCESS
}

type ShippingInfo struct {
	TxId                 string              `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	ItemId               string              `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SequenceNo           int32               `protobuf:"varint,3,opt,name=sequence_no,json=sequenceNo,proto3" json:"sequence_no,omitempty"`
	MaterialLotId        string              `protobuf:"bytes,4,opt,name=material_lot_id,json=materialLotId,proto3" json:"material_lot_id,omitempty"`
	MaterialProductId    string              `protobuf:"bytes,5,opt,name=material_product_id,json=materialProductId,proto3" json:"material_product_id,omitempty"`
	MaterialLevel        string              `protobuf:"bytes,6,opt,name=material_level,json=materialLevel,proto3" json:"material_level,omitempty"`
	Qty                  float32             `protobuf:"fixed32,7,opt,name=qty,proto3" json:"qty,omitempty"`
	Unit                 string              `protobuf:"bytes,8,opt,name=unit,proto3" json:"unit,omitempty"`
	UpdateAt             string              `protobuf:"bytes,9,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	UpdateBy             string              `protobuf:"bytes,10,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	InsertAt             string              `protobuf:"bytes,11,opt,name=insert_at,json=insertAt,proto3" json:"insert_at,omitempty"`
	InsertBy             string              `protobuf:"bytes,12,opt,name=insert_by,json=insertBy,proto3" json:"insert_by,omitempty"`
	Status               code.ShippingStatus `protobuf:"varint,13,opt,name=status,proto3,enum=proto.mes.code.ShippingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ShippingInfo) Reset()         { *m = ShippingInfo{} }
func (m *ShippingInfo) String() string { return proto.CompactTextString(m) }
func (*ShippingInfo) ProtoMessage()    {}
func (*ShippingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_095c79d790d164da, []int{2}
}

func (m *ShippingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShippingInfo.Unmarshal(m, b)
}
func (m *ShippingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShippingInfo.Marshal(b, m, deterministic)
}
func (m *ShippingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShippingInfo.Merge(m, src)
}
func (m *ShippingInfo) XXX_Size() int {
	return xxx_messageInfo_ShippingInfo.Size(m)
}
func (m *ShippingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ShippingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ShippingInfo proto.InternalMessageInfo

func (m *ShippingInfo) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ShippingInfo) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *ShippingInfo) GetSequenceNo() int32 {
	if m != nil {
		return m.SequenceNo
	}
	return 0
}

func (m *ShippingInfo) GetMaterialLotId() string {
	if m != nil {
		return m.MaterialLotId
	}
	return ""
}

func (m *ShippingInfo) GetMaterialProductId() string {
	if m != nil {
		return m.MaterialProductId
	}
	return ""
}

func (m *ShippingInfo) GetMaterialLevel() string {
	if m != nil {
		return m.MaterialLevel
	}
	return ""
}

func (m *ShippingInfo) GetQty() float32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *ShippingInfo) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *ShippingInfo) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func (m *ShippingInfo) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *ShippingInfo) GetInsertAt() string {
	if m != nil {
		return m.InsertAt
	}
	return ""
}

func (m *ShippingInfo) GetInsertBy() string {
	if m != nil {
		return m.InsertBy
	}
	return ""
}

func (m *ShippingInfo) GetStatus() code.ShippingStatus {
	if m != nil {
		return m.Status
	}
	return code.ShippingStatus_SHIPPING_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MTxShipListOperatingRequest)(nil), "proto.mes.MTxShipListOperatingRequest")
	proto.RegisterType((*MTxShipListOperatingReply)(nil), "proto.mes.MTxShipListOperatingReply")
	proto.RegisterType((*ShippingInfo)(nil), "proto.mes.ShippingInfo")
}

func init() {
	proto.RegisterFile("transaction_shipping_operating.proto", fileDescriptor_095c79d790d164da)
}

var fileDescriptor_095c79d790d164da = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xe5, 0x38, 0xff, 0x3c, 0x49, 0xf3, 0x7d, 0x4c, 0x17, 0x35, 0x89, 0x0a, 0x56, 0x04,
	0xc8, 0x62, 0x61, 0x4b, 0x41, 0xc0, 0x02, 0x84, 0x44, 0x16, 0x48, 0x91, 0x1a, 0x40, 0x4e, 0x56,
	0x6c, 0xac, 0x89, 0x3d, 0xa4, 0x23, 0xec, 0x19, 0xd7, 0x73, 0x5d, 0xe2, 0x17, 0xe1, 0x31, 0x78,
	0x0b, 0xde, 0x0b, 0xcd, 0x38, 0x4e, 0x43, 0x88, 0x4a, 0xbb, 0xca, 0xf8, 0x9c, 0x73, 0xcf, 0x75,
	0xec, 0x9f, 0x8c, 0x9e, 0x40, 0x4e, 0xb8, 0x24, 0x11, 0x30, 0xc1, 0x43, 0x79, 0xc9, 0xb2, 0x8c,
	0xf1, 0x75, 0x28, 0x32, 0x9a, 0x13, 0x60, 0x7c, 0xed, 0x65, 0xb9, 0x00, 0x81, 0x2d, 0xfd, 0xe3,
	0xa5, 0x54, 0x0e, 0x87, 0xfa, 0xe8, 0xa7, 0x54, 0xfa, 0x91, 0x88, 0xa9, 0x4f, 0x79, 0x91, 0xca,
	0x2a, 0x36, 0x74, 0x0e, 0xbd, 0x3c, 0x17, 0xf9, 0x7e, 0x62, 0xfc, 0xcb, 0x44, 0xa3, 0xf9, 0x72,
	0xb3, 0xb8, 0x64, 0xd9, 0x05, 0x93, 0xf0, 0xa9, 0xde, 0x13, 0xd0, 0xab, 0x82, 0x4a, 0xc0, 0xef,
	0x50, 0x6f, 0x7b, 0x2f, 0x50, 0x66, 0xd4, 0x36, 0x1c, 0xc3, 0x1d, 0x4c, 0xce, 0xbd, 0xdd, 0x7a,
	0x4f, 0xf5, 0x7a, 0x1f, 0x0a, 0xae, 0x43, 0x73, 0x11, 0xd3, 0x24, 0x40, 0xd5, 0xc4, 0xb2, 0xcc,
	0x28, 0x3e, 0x45, 0x2d, 0xd8, 0x84, 0x2c, 0xb6, 0x1b, 0x8e, 0xe1, 0x5a, 0x41, 0x13, 0x36, 0xb3,
	0x18, 0x9f, 0xa1, 0x0e, 0x03, 0x9a, 0x2a, 0xd9, 0xd4, 0x72, 0x5b, 0x5d, 0xce, 0x62, 0xfc, 0x18,
	0xf5, 0xa4, 0x5a, 0xcc, 0x23, 0x1a, 0x72, 0x61, 0x37, 0x1d, 0xc3, 0x6d, 0x05, 0xa8, 0x96, 0x3e,
	0x0a, 0xfc, 0x0c, 0xfd, 0x97, 0x12, 0xa0, 0x39, 0x23, 0x49, 0x98, 0x08, 0x50, 0x0d, 0x2d, 0xdd,
	0x70, 0x52, 0xcb, 0x17, 0x02, 0x66, 0x31, 0xf6, 0xd0, 0xe9, 0x2e, 0x97, 0xe5, 0x22, 0x2e, 0x22,
	0x9d, 0x6d, 0xeb, 0xec, 0x83, 0xda, 0xfa, 0x5c, 0x39, 0xb3, 0x18, 0x3f, 0x45, 0x83, 0x9b, 0x5e,
	0x7a, 0x4d, 0x13, 0xbb, 0x73, 0x50, 0xab, 0x44, 0xfc, 0x3f, 0x32, 0xaf, 0xa0, 0xb4, 0xbb, 0x8e,
	0xe1, 0x36, 0x02, 0x75, 0xc4, 0x18, 0x35, 0x0b, 0xce, 0xc0, 0xb6, 0xaa, 0xbf, 0xa7, 0xce, 0x78,
	0x84, 0xac, 0x22, 0x8b, 0x09, 0xd0, 0x70, 0x55, 0xda, 0x48, 0x1b, 0xdd, 0x4a, 0x98, 0x96, 0xf8,
	0x1c, 0x21, 0xc1, 0x93, 0x32, 0x04, 0x11, 0x93, 0xd2, 0xee, 0x39, 0x86, 0xdb, 0x0d, 0x2c, 0xa5,
	0x2c, 0x95, 0x80, 0x5f, 0xa1, 0xb6, 0x04, 0x02, 0x85, 0xb4, 0xfb, 0xfa, 0x51, 0x3f, 0x3a, 0x7c,
	0xd4, 0x8b, 0x2d, 0x12, 0x0b, 0x9d, 0x0a, 0xb6, 0xe9, 0xf1, 0x0f, 0x03, 0x3d, 0x3c, 0xfe, 0x1e,
	0xb3, 0xa4, 0xc4, 0x6f, 0x50, 0xbf, 0x46, 0x49, 0xb9, 0xb6, 0xe1, 0x98, 0x6e, 0x6f, 0x72, 0xb6,
	0xd7, 0x5d, 0xd7, 0xce, 0xf8, 0x57, 0x11, 0xfc, 0x11, 0xc6, 0x6f, 0x91, 0x49, 0xf3, 0x5c, 0xbf,
	0xc0, 0xc1, 0xe4, 0xf9, 0xe1, 0xfd, 0x68, 0xa4, 0xbc, 0x39, 0x1c, 0x59, 0xad, 0xc6, 0xc6, 0x3f,
	0x4d, 0xd4, 0xdf, 0x2f, 0xbf, 0x21, 0xc2, 0x38, 0x4e, 0x44, 0xe3, 0x36, 0x22, 0xcc, 0xbb, 0x10,
	0xd1, 0xbc, 0x07, 0x11, 0xad, 0xbb, 0x13, 0xd1, 0xbe, 0x85, 0x88, 0xce, 0xdf, 0x44, 0x74, 0x8f,
	0x12, 0x41, 0x6a, 0x54, 0xb6, 0x44, 0xbc, 0xff, 0x07, 0x2e, 0x23, 0x64, 0x31, 0x2e, 0x69, 0x0e,
	0x6a, 0xb2, 0x57, 0x99, 0x95, 0x50, 0x4d, 0x6e, 0xcd, 0x55, 0xa9, 0x79, 0xd9, 0x99, 0xd3, 0x7d,
	0x92, 0x4e, 0xee, 0x43, 0xd2, 0xf4, 0xf5, 0x97, 0x97, 0x6b, 0x06, 0x09, 0x59, 0x79, 0xdf, 0x28,
	0x8f, 0x89, 0x17, 0x89, 0xd4, 0x83, 0xef, 0xbe, 0xbe, 0xf0, 0x23, 0x91, 0xa6, 0x82, 0x4b, 0xff,
	0x7a, 0xe2, 0x57, 0x9f, 0x96, 0xb5, 0x48, 0x08, 0x5f, 0xab, 0x2f, 0xcc, 0xaa, 0xad, 0x95, 0x17,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x10, 0x55, 0x93, 0x5a, 0xc2, 0x04, 0x00, 0x00,
}
