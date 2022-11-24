// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction_operating.proto

package mes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	rs "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
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

type TransactionListOperatingRequest struct {
	ActionType           code.FunctionModel                          `protobuf:"varint,1,opt,name=action_type,json=actionType,proto3,enum=proto.mes.code.FunctionModel" json:"action_type,omitempty"`
	TxId                 string                                      `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	TxType               code.TransactionType                        `protobuf:"varint,3,opt,name=tx_type,json=txType,proto3,enum=proto.mes.code.TransactionType" json:"tx_type,omitempty"`
	Materials            []*TransactionListOperatingRequest_Material `protobuf:"bytes,4,rep,name=materials,proto3" json:"materials,omitempty"`
	CurrentFactoryId     rs.FactoryID                                `protobuf:"varint,5,opt,name=current_factory_id,json=currentFactoryId,proto3,enum=proto.dm.rs.FactoryID" json:"current_factory_id,omitempty"`
	CurrentDepartmentId  string                                      `protobuf:"bytes,6,opt,name=current_department_id,json=currentDepartmentId,proto3" json:"current_department_id,omitempty"`
	NextFactoryId        rs.FactoryID                                `protobuf:"varint,7,opt,name=next_factory_id,json=nextFactoryId,proto3,enum=proto.dm.rs.FactoryID" json:"next_factory_id,omitempty"`
	NextDepartmentId     string                                      `protobuf:"bytes,8,opt,name=next_department_id,json=nextDepartmentId,proto3" json:"next_department_id,omitempty"`
	Codd                 string                                      `protobuf:"bytes,9,opt,name=codd,proto3" json:"codd,omitempty"`
	Status               code.TransactionStatus                      `protobuf:"varint,10,opt,name=status,proto3,enum=proto.mes.code.TransactionStatus" json:"status,omitempty"`
	Comment              string                                      `protobuf:"bytes,11,opt,name=comment,proto3" json:"comment,omitempty"`
	UpdateBy             string                                      `protobuf:"bytes,12,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	OnlyToday            bool                                        `protobuf:"varint,13,opt,name=only_today,json=onlyToday,proto3" json:"only_today,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *TransactionListOperatingRequest) Reset()         { *m = TransactionListOperatingRequest{} }
func (m *TransactionListOperatingRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionListOperatingRequest) ProtoMessage()    {}
func (*TransactionListOperatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c84113a8cb1c0ee1, []int{0}
}

func (m *TransactionListOperatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionListOperatingRequest.Unmarshal(m, b)
}
func (m *TransactionListOperatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionListOperatingRequest.Marshal(b, m, deterministic)
}
func (m *TransactionListOperatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionListOperatingRequest.Merge(m, src)
}
func (m *TransactionListOperatingRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionListOperatingRequest.Size(m)
}
func (m *TransactionListOperatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionListOperatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionListOperatingRequest proto.InternalMessageInfo

func (m *TransactionListOperatingRequest) GetActionType() code.FunctionModel {
	if m != nil {
		return m.ActionType
	}
	return code.FunctionModel_FUNCTION_MODEL_UNSPECIFIED
}

func (m *TransactionListOperatingRequest) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *TransactionListOperatingRequest) GetTxType() code.TransactionType {
	if m != nil {
		return m.TxType
	}
	return code.TransactionType_TRANSACTION_TYPE_UNSPECIFIED
}

func (m *TransactionListOperatingRequest) GetMaterials() []*TransactionListOperatingRequest_Material {
	if m != nil {
		return m.Materials
	}
	return nil
}

func (m *TransactionListOperatingRequest) GetCurrentFactoryId() rs.FactoryID {
	if m != nil {
		return m.CurrentFactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

func (m *TransactionListOperatingRequest) GetCurrentDepartmentId() string {
	if m != nil {
		return m.CurrentDepartmentId
	}
	return ""
}

func (m *TransactionListOperatingRequest) GetNextFactoryId() rs.FactoryID {
	if m != nil {
		return m.NextFactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

func (m *TransactionListOperatingRequest) GetNextDepartmentId() string {
	if m != nil {
		return m.NextDepartmentId
	}
	return ""
}

func (m *TransactionListOperatingRequest) GetCodd() string {
	if m != nil {
		return m.Codd
	}
	return ""
}

func (m *TransactionListOperatingRequest) GetStatus() code.TransactionStatus {
	if m != nil {
		return m.Status
	}
	return code.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED
}

func (m *TransactionListOperatingRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *TransactionListOperatingRequest) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *TransactionListOperatingRequest) GetOnlyToday() bool {
	if m != nil {
		return m.OnlyToday
	}
	return false
}

type TransactionListOperatingRequest_Material struct {
	ItemId               string                 `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ProductType          rs.ProductType         `protobuf:"varint,2,opt,name=product_type,json=productType,proto3,enum=proto.dm.rs.ProductType" json:"product_type,omitempty"`
	Id                   string                 `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Level                string                 `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	Status               code.TransactionStatus `protobuf:"varint,5,opt,name=status,proto3,enum=proto.mes.code.TransactionStatus" json:"status,omitempty"`
	Qty                  float32                `protobuf:"fixed32,6,opt,name=qty,proto3" json:"qty,omitempty"`
	Unit                 string                 `protobuf:"bytes,7,opt,name=unit,proto3" json:"unit,omitempty"`
	Comment              string                 `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TransactionListOperatingRequest_Material) Reset() {
	*m = TransactionListOperatingRequest_Material{}
}
func (m *TransactionListOperatingRequest_Material) String() string { return proto.CompactTextString(m) }
func (*TransactionListOperatingRequest_Material) ProtoMessage()    {}
func (*TransactionListOperatingRequest_Material) Descriptor() ([]byte, []int) {
	return fileDescriptor_c84113a8cb1c0ee1, []int{0, 0}
}

func (m *TransactionListOperatingRequest_Material) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionListOperatingRequest_Material.Unmarshal(m, b)
}
func (m *TransactionListOperatingRequest_Material) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionListOperatingRequest_Material.Marshal(b, m, deterministic)
}
func (m *TransactionListOperatingRequest_Material) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionListOperatingRequest_Material.Merge(m, src)
}
func (m *TransactionListOperatingRequest_Material) XXX_Size() int {
	return xxx_messageInfo_TransactionListOperatingRequest_Material.Size(m)
}
func (m *TransactionListOperatingRequest_Material) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionListOperatingRequest_Material.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionListOperatingRequest_Material proto.InternalMessageInfo

func (m *TransactionListOperatingRequest_Material) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *TransactionListOperatingRequest_Material) GetProductType() rs.ProductType {
	if m != nil {
		return m.ProductType
	}
	return rs.ProductType_PRODUCT_TYPE_UNSPECIFIED
}

func (m *TransactionListOperatingRequest_Material) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TransactionListOperatingRequest_Material) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *TransactionListOperatingRequest_Material) GetStatus() code.TransactionStatus {
	if m != nil {
		return m.Status
	}
	return code.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED
}

func (m *TransactionListOperatingRequest_Material) GetQty() float32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *TransactionListOperatingRequest_Material) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *TransactionListOperatingRequest_Material) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type TransactionListOperatingReply struct {
	TransactionLists     []*TransactionListOperatingReply_TransactionList `protobuf:"bytes,1,rep,name=transactionLists,proto3" json:"transactionLists,omitempty"`
	Err                  error1.TransactionOperating                      `protobuf:"varint,3,opt,name=err,proto3,enum=proto.mes.code.error.TransactionOperating" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *TransactionListOperatingReply) Reset()         { *m = TransactionListOperatingReply{} }
func (m *TransactionListOperatingReply) String() string { return proto.CompactTextString(m) }
func (*TransactionListOperatingReply) ProtoMessage()    {}
func (*TransactionListOperatingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c84113a8cb1c0ee1, []int{1}
}

func (m *TransactionListOperatingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionListOperatingReply.Unmarshal(m, b)
}
func (m *TransactionListOperatingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionListOperatingReply.Marshal(b, m, deterministic)
}
func (m *TransactionListOperatingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionListOperatingReply.Merge(m, src)
}
func (m *TransactionListOperatingReply) XXX_Size() int {
	return xxx_messageInfo_TransactionListOperatingReply.Size(m)
}
func (m *TransactionListOperatingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionListOperatingReply.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionListOperatingReply proto.InternalMessageInfo

func (m *TransactionListOperatingReply) GetTransactionLists() []*TransactionListOperatingReply_TransactionList {
	if m != nil {
		return m.TransactionLists
	}
	return nil
}

func (m *TransactionListOperatingReply) GetErr() error1.TransactionOperating {
	if m != nil {
		return m.Err
	}
	return error1.TransactionOperating_TRANSACTION_OPERATING_SUCCESS
}

type TransactionListOperatingReply_TransactionList struct {
	//轉移編號 (廠+移出部門+單據日期+三碼流水號)
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	//1->生產 (same factory) 2->調撥 (diff factory) 3->代工 4->報廢 5->買賣 6->研發
	TxType code.TransactionType `protobuf:"varint,2,opt,name=tx_type,json=txType,proto3,enum=proto.mes.code.TransactionType" json:"tx_type,omitempty"`
	//項次
	ItemId string `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	//材料類別
	ProductType rs.ProductType `protobuf:"varint,4,opt,name=product_type,json=productType,proto3,enum=proto.dm.rs.ProductType" json:"product_type,omitempty"`
	//材料 ID ,材料,半製品,成品編號
	MaterialProductId string `protobuf:"bytes,5,opt,name=material_product_id,json=materialProductId,proto3" json:"material_product_id,omitempty"`
	//材料等級
	MaterialLevel string `protobuf:"bytes,6,opt,name=material_level,json=materialLevel,proto3" json:"material_level,omitempty"`
	//移出廠
	CurrentFactoryId rs.FactoryID `protobuf:"varint,7,opt,name=current_factory_id,json=currentFactoryId,proto3,enum=proto.dm.rs.FactoryID" json:"current_factory_id,omitempty"`
	//移出單位代號
	CurrentDepartmentId string `protobuf:"bytes,8,opt,name=current_department_id,json=currentDepartmentId,proto3" json:"current_department_id,omitempty"`
	//移入廠
	NextFactoryId rs.FactoryID `protobuf:"varint,9,opt,name=next_factory_id,json=nextFactoryId,proto3,enum=proto.dm.rs.FactoryID" json:"next_factory_id,omitempty"`
	//移入單位代號
	NextDepartmentId string `protobuf:"bytes,10,opt,name=next_department_id,json=nextDepartmentId,proto3" json:"next_department_id,omitempty"`
	//內外購
	Codd string `protobuf:"bytes,11,opt,name=codd,proto3" json:"codd,omitempty"`
	//移轉數量
	Qty float32 `protobuf:"fixed32,12,opt,name=qty,proto3" json:"qty,omitempty"`
	//計量單位
	Unit string `protobuf:"bytes,13,opt,name=unit,proto3" json:"unit,omitempty"`
	//確認日期
	ConfirmAt string `protobuf:"bytes,14,opt,name=confirm_at,json=confirmAt,proto3" json:"confirm_at,omitempty"`
	//確認人
	ConfirmBy string `protobuf:"bytes,15,opt,name=confirm_by,json=confirmBy,proto3" json:"confirm_by,omitempty"`
	//移轉狀態
	Status code.TransactionStatus `protobuf:"varint,16,opt,name=status,proto3,enum=proto.mes.code.TransactionStatus" json:"status,omitempty"`
	//註解
	TxComments string `protobuf:"bytes,17,opt,name=tx_comments,json=txComments,proto3" json:"tx_comments,omitempty"`
	//材料註解
	Comments string `protobuf:"bytes,18,opt,name=comments,proto3" json:"comments,omitempty"`
	//更新日期
	UpdateAt string `protobuf:"bytes,19,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	//更新人
	UpdateBy string `protobuf:"bytes,20,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	//建檔日期
	InsertAt string `protobuf:"bytes,21,opt,name=insert_at,json=insertAt,proto3" json:"insert_at,omitempty"`
	//建檔人
	InsertBy             string   `protobuf:"bytes,22,opt,name=insert_by,json=insertBy,proto3" json:"insert_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionListOperatingReply_TransactionList) Reset() {
	*m = TransactionListOperatingReply_TransactionList{}
}
func (m *TransactionListOperatingReply_TransactionList) String() string {
	return proto.CompactTextString(m)
}
func (*TransactionListOperatingReply_TransactionList) ProtoMessage() {}
func (*TransactionListOperatingReply_TransactionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c84113a8cb1c0ee1, []int{1, 0}
}

func (m *TransactionListOperatingReply_TransactionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionListOperatingReply_TransactionList.Unmarshal(m, b)
}
func (m *TransactionListOperatingReply_TransactionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionListOperatingReply_TransactionList.Marshal(b, m, deterministic)
}
func (m *TransactionListOperatingReply_TransactionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionListOperatingReply_TransactionList.Merge(m, src)
}
func (m *TransactionListOperatingReply_TransactionList) XXX_Size() int {
	return xxx_messageInfo_TransactionListOperatingReply_TransactionList.Size(m)
}
func (m *TransactionListOperatingReply_TransactionList) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionListOperatingReply_TransactionList.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionListOperatingReply_TransactionList proto.InternalMessageInfo

func (m *TransactionListOperatingReply_TransactionList) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetTxType() code.TransactionType {
	if m != nil {
		return m.TxType
	}
	return code.TransactionType_TRANSACTION_TYPE_UNSPECIFIED
}

func (m *TransactionListOperatingReply_TransactionList) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetProductType() rs.ProductType {
	if m != nil {
		return m.ProductType
	}
	return rs.ProductType_PRODUCT_TYPE_UNSPECIFIED
}

func (m *TransactionListOperatingReply_TransactionList) GetMaterialProductId() string {
	if m != nil {
		return m.MaterialProductId
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetMaterialLevel() string {
	if m != nil {
		return m.MaterialLevel
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetCurrentFactoryId() rs.FactoryID {
	if m != nil {
		return m.CurrentFactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

func (m *TransactionListOperatingReply_TransactionList) GetCurrentDepartmentId() string {
	if m != nil {
		return m.CurrentDepartmentId
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetNextFactoryId() rs.FactoryID {
	if m != nil {
		return m.NextFactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

func (m *TransactionListOperatingReply_TransactionList) GetNextDepartmentId() string {
	if m != nil {
		return m.NextDepartmentId
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetCodd() string {
	if m != nil {
		return m.Codd
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetQty() float32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *TransactionListOperatingReply_TransactionList) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetConfirmAt() string {
	if m != nil {
		return m.ConfirmAt
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetConfirmBy() string {
	if m != nil {
		return m.ConfirmBy
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetStatus() code.TransactionStatus {
	if m != nil {
		return m.Status
	}
	return code.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED
}

func (m *TransactionListOperatingReply_TransactionList) GetTxComments() string {
	if m != nil {
		return m.TxComments
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetInsertAt() string {
	if m != nil {
		return m.InsertAt
	}
	return ""
}

func (m *TransactionListOperatingReply_TransactionList) GetInsertBy() string {
	if m != nil {
		return m.InsertBy
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionListOperatingRequest)(nil), "proto.mes.TransactionListOperatingRequest")
	proto.RegisterType((*TransactionListOperatingRequest_Material)(nil), "proto.mes.TransactionListOperatingRequest.Material")
	proto.RegisterType((*TransactionListOperatingReply)(nil), "proto.mes.TransactionListOperatingReply")
	proto.RegisterType((*TransactionListOperatingReply_TransactionList)(nil), "proto.mes.TransactionListOperatingReply.TransactionList")
}

func init() { proto.RegisterFile("transaction_operating.proto", fileDescriptor_c84113a8cb1c0ee1) }

var fileDescriptor_c84113a8cb1c0ee1 = []byte{
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6a, 0xeb, 0x46,
	0x10, 0xc6, 0xff, 0xd2, 0x38, 0x76, 0x7c, 0xd6, 0xe7, 0x67, 0x71, 0x30, 0x71, 0x0f, 0x14, 0x4c,
	0x29, 0x12, 0xf8, 0x50, 0x7a, 0x4a, 0xcb, 0x81, 0xb8, 0x21, 0x60, 0x48, 0x68, 0xab, 0xe6, 0xaa,
	0x37, 0x42, 0xd6, 0x6e, 0x8c, 0xa8, 0xa4, 0x55, 0x56, 0xab, 0xd4, 0xba, 0xec, 0x3b, 0xf4, 0xbe,
	0xcf, 0xd5, 0xb7, 0x29, 0xbb, 0x2b, 0xd9, 0x92, 0x93, 0x26, 0x35, 0xed, 0x95, 0x77, 0xe7, 0x9b,
	0xf9, 0x66, 0x76, 0x77, 0xe6, 0xb3, 0xe0, 0x4c, 0x70, 0x2f, 0x4e, 0x3d, 0x5f, 0x04, 0x2c, 0x76,
	0x59, 0x42, 0xb9, 0x27, 0x82, 0x78, 0x63, 0x25, 0x9c, 0x09, 0x86, 0x4c, 0xf5, 0x63, 0x45, 0x34,
	0x9d, 0x4c, 0xd4, 0xd2, 0x8e, 0x68, 0x6a, 0xfb, 0x8c, 0x50, 0x9b, 0xc6, 0x59, 0x94, 0x6a, 0xb7,
	0xc9, 0xec, 0x10, 0xe3, 0x9c, 0xf1, 0x9a, 0xc7, 0x3b, 0xed, 0x41, 0x22, 0x9b, 0xa7, 0x55, 0xe0,
	0xfd, 0x1f, 0x3d, 0x38, 0xbf, 0xdd, 0x57, 0x70, 0x1d, 0xa4, 0xe2, 0x87, 0xb2, 0x08, 0x87, 0xde,
	0x67, 0x34, 0x15, 0xe8, 0x13, 0xf4, 0x8b, 0xfa, 0x44, 0x9e, 0x50, 0xdc, 0x98, 0x35, 0xe6, 0xc3,
	0xc5, 0xd4, 0xda, 0xd5, 0x66, 0xc9, 0xa4, 0xd6, 0x55, 0x16, 0x2b, 0xa7, 0x1b, 0x46, 0x68, 0xe8,
	0x80, 0x8e, 0xb8, 0xcd, 0x13, 0x8a, 0xc6, 0xd0, 0x11, 0x5b, 0x37, 0x20, 0xb8, 0x39, 0x6b, 0xcc,
	0x4d, 0xa7, 0x2d, 0xb6, 0x2b, 0x82, 0x3e, 0x42, 0x4f, 0x6c, 0x35, 0x61, 0x4b, 0x11, 0x9e, 0x1f,
	0x12, 0x56, 0xca, 0x92, 0x34, 0x4e, 0x57, 0x6c, 0x15, 0xdd, 0x4f, 0x60, 0x46, 0x9e, 0xa0, 0x3c,
	0xf0, 0xc2, 0x14, 0xb7, 0x67, 0xad, 0x79, 0x7f, 0xf1, 0xa1, 0x12, 0xfb, 0xc2, 0x69, 0xac, 0x9b,
	0x22, 0xd6, 0xd9, 0xb3, 0xa0, 0x4b, 0x40, 0x7e, 0xc6, 0x39, 0x8d, 0x85, 0x7b, 0xe7, 0xf9, 0x82,
	0xf1, 0x5c, 0x96, 0xdb, 0x51, 0x75, 0xbd, 0x2d, 0xb8, 0x49, 0x64, 0xf1, 0xd4, 0xba, 0xd2, 0xf0,
	0xea, 0xd2, 0x19, 0x15, 0x11, 0xa5, 0x85, 0xa0, 0x05, 0xbc, 0x29, 0x59, 0x08, 0x4d, 0x3c, 0x2e,
	0x22, 0xb9, 0x0c, 0x08, 0xee, 0xaa, 0x73, 0x8f, 0x0b, 0xf0, 0x72, 0x87, 0xad, 0x08, 0xfa, 0x04,
	0xa7, 0x31, 0xdd, 0xd6, 0xd2, 0xf6, 0x9e, 0x4d, 0x3b, 0x90, 0xee, 0xfb, 0x9c, 0x5f, 0x02, 0x52,
	0xf1, 0xf5, 0x84, 0x86, 0x4a, 0x38, 0x92, 0x48, 0x2d, 0x1b, 0x82, 0xb6, 0xcf, 0x08, 0xc1, 0xa6,
	0x7e, 0x08, 0xb9, 0x46, 0xdf, 0x40, 0x37, 0x15, 0x9e, 0xc8, 0x52, 0x0c, 0x2a, 0xf1, 0x67, 0xcf,
	0xbc, 0xc3, 0xcf, 0xca, 0xd1, 0x29, 0x02, 0x10, 0x86, 0x9e, 0xcf, 0x22, 0xc9, 0x8d, 0xfb, 0x8a,
	0xb1, 0xdc, 0xa2, 0x33, 0x30, 0xb3, 0x84, 0x78, 0x82, 0xba, 0xeb, 0x1c, 0x9f, 0x28, 0xcc, 0xd0,
	0x86, 0x65, 0x8e, 0xa6, 0x00, 0x2c, 0x0e, 0x73, 0x57, 0x30, 0xe2, 0xe5, 0x78, 0x30, 0x6b, 0xcc,
	0x0d, 0xc7, 0x94, 0x96, 0x5b, 0x69, 0x98, 0xfc, 0xde, 0x04, 0xa3, 0x7c, 0x24, 0xf4, 0x0e, 0x7a,
	0x81, 0xa0, 0x91, 0x3c, 0x54, 0x43, 0xd1, 0x74, 0xe5, 0x76, 0x45, 0xd0, 0xb7, 0x70, 0x92, 0x70,
	0x46, 0x32, 0x5f, 0xe8, 0x26, 0x6a, 0xaa, 0xe2, 0x71, 0xed, 0xd6, 0x7e, 0xd4, 0x0e, 0xaa, 0x7b,
	0xfa, 0xc9, 0x7e, 0x83, 0x86, 0xd0, 0x0c, 0x88, 0xea, 0x3b, 0xd3, 0x69, 0x06, 0x04, 0xbd, 0x86,
	0x4e, 0x48, 0x1f, 0x68, 0x88, 0xdb, 0xca, 0xa4, 0x37, 0x95, 0x9b, 0xe9, 0x1c, 0x7b, 0x33, 0x23,
	0x68, 0xdd, 0x8b, 0x5c, 0x3d, 0x7c, 0xd3, 0x91, 0x4b, 0x79, 0xf5, 0x59, 0x1c, 0x08, 0xf5, 0xba,
	0xa6, 0xa3, 0xd6, 0xd5, 0xfb, 0x33, 0x6a, 0xf7, 0xf7, 0xfe, 0x4f, 0x03, 0xa6, 0xff, 0xdc, 0xc8,
	0x49, 0x98, 0x23, 0x02, 0x23, 0x51, 0x77, 0x48, 0x71, 0x43, 0x0d, 0xc3, 0xc7, 0x7f, 0x35, 0x0c,
	0x49, 0x98, 0x1f, 0xa2, 0xce, 0x23, 0x46, 0xf4, 0x1d, 0xb4, 0x28, 0xe7, 0xc5, 0x84, 0x7e, 0x71,
	0x78, 0x7e, 0xa5, 0x33, 0x55, 0x96, 0x3d, 0xbf, 0x0c, 0x9b, 0xfc, 0xd5, 0x85, 0xd3, 0x83, 0x1c,
	0x7b, 0x31, 0x68, 0x3c, 0x2d, 0x06, 0xcd, 0xe3, 0xc4, 0xa0, 0xd2, 0x1f, 0xad, 0x67, 0xfb, 0xa3,
	0x7d, 0x4c, 0x7f, 0x58, 0x30, 0x2e, 0xc5, 0xc1, 0x2d, 0x59, 0x0a, 0x41, 0x30, 0x9d, 0x57, 0x25,
	0x54, 0x84, 0xaf, 0x08, 0xfa, 0x1c, 0x86, 0x3b, 0x7f, 0xdd, 0x48, 0x7a, 0xe4, 0x07, 0xa5, 0xf5,
	0x5a, 0x35, 0xd4, 0xd3, 0x32, 0xd3, 0xfb, 0xbf, 0x64, 0xc6, 0x38, 0x4a, 0x66, 0xcc, 0xff, 0x2e,
	0x33, 0xf0, 0x82, 0xcc, 0xf4, 0x2b, 0x32, 0x53, 0x4c, 0xc4, 0xc9, 0xe3, 0x89, 0x18, 0x54, 0x26,
	0x62, 0x0a, 0xe0, 0xb3, 0xf8, 0x2e, 0xe0, 0x91, 0xeb, 0x09, 0x3c, 0x54, 0x88, 0x59, 0x58, 0x2e,
	0x6a, 0xf0, 0x3a, 0xc7, 0xa7, 0x35, 0x78, 0x99, 0x57, 0x06, 0x76, 0x74, 0xec, 0xc0, 0x9e, 0x43,
	0x5f, 0x6c, 0xdd, 0x62, 0xfc, 0x52, 0xfc, 0x4a, 0x51, 0x83, 0xd8, 0x7e, 0x5f, 0x58, 0xd0, 0x04,
	0x8c, 0x1d, 0x8a, 0xb4, 0xa0, 0x95, 0xfb, 0x8a, 0xda, 0x79, 0x02, 0x8f, 0xab, 0x6a, 0x77, 0x71,
	0x20, 0x85, 0xaf, 0x0f, 0xa4, 0xf0, 0x0c, 0xcc, 0x20, 0x4e, 0x29, 0x17, 0x32, 0xf2, 0x8d, 0x06,
	0xb5, 0x41, 0x47, 0x16, 0xe0, 0x3a, 0xc7, 0x6f, 0xab, 0xe0, 0x32, 0x5f, 0x7e, 0xfd, 0xcb, 0x57,
	0x9b, 0x40, 0x84, 0xde, 0xda, 0xfa, 0x95, 0xc6, 0xc4, 0xb3, 0x7c, 0x16, 0x59, 0xe2, 0x37, 0x5b,
	0x6d, 0x6c, 0x59, 0x19, 0x8b, 0x53, 0xfb, 0x61, 0x61, 0xeb, 0x3f, 0xfe, 0x0d, 0x0b, 0xbd, 0x78,
	0x23, 0xbf, 0x10, 0xd6, 0x5d, 0x65, 0xf9, 0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xce,
	0x92, 0x8e, 0x79, 0x08, 0x00, 0x00,
}