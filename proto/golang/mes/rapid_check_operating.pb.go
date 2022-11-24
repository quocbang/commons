// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rapid_check_operating.proto

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

type RapidCheckOperatingRequest struct {
	// 請求類別
	ActionType code.FunctionModel `protobuf:"varint,1,opt,name=action_type,json=actionType,proto3,enum=proto.mes.code.FunctionModel" json:"action_type,omitempty"`
	// I = 進站 , O = 出站
	ActionMode code.Mode `protobuf:"varint,2,opt,name=action_mode,json=actionMode,proto3,enum=proto.mes.code.Mode" json:"action_mode,omitempty"`
	// 搜尋用
	LotId                string                                       `protobuf:"bytes,3,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	SheetIds             []string                                     `protobuf:"bytes,4,rep,name=sheet_ids,json=sheetIds,proto3" json:"sheet_ids,omitempty"`
	UpdateBy             string                                       `protobuf:"bytes,5,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	MeasurementItems     []*RapidCheckOperatingRequestMeasurementItem `protobuf:"bytes,6,rep,name=measurement_items,json=measurementItems,proto3" json:"measurement_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *RapidCheckOperatingRequest) Reset()         { *m = RapidCheckOperatingRequest{} }
func (m *RapidCheckOperatingRequest) String() string { return proto.CompactTextString(m) }
func (*RapidCheckOperatingRequest) ProtoMessage()    {}
func (*RapidCheckOperatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2ad98225f42789, []int{0}
}

func (m *RapidCheckOperatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RapidCheckOperatingRequest.Unmarshal(m, b)
}
func (m *RapidCheckOperatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RapidCheckOperatingRequest.Marshal(b, m, deterministic)
}
func (m *RapidCheckOperatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RapidCheckOperatingRequest.Merge(m, src)
}
func (m *RapidCheckOperatingRequest) XXX_Size() int {
	return xxx_messageInfo_RapidCheckOperatingRequest.Size(m)
}
func (m *RapidCheckOperatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RapidCheckOperatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RapidCheckOperatingRequest proto.InternalMessageInfo

func (m *RapidCheckOperatingRequest) GetActionType() code.FunctionModel {
	if m != nil {
		return m.ActionType
	}
	return code.FunctionModel_FUNCTION_MODEL_UNSPECIFIED
}

func (m *RapidCheckOperatingRequest) GetActionMode() code.Mode {
	if m != nil {
		return m.ActionMode
	}
	return code.Mode_MODE_UNSPECIFIED
}

func (m *RapidCheckOperatingRequest) GetLotId() string {
	if m != nil {
		return m.LotId
	}
	return ""
}

func (m *RapidCheckOperatingRequest) GetSheetIds() []string {
	if m != nil {
		return m.SheetIds
	}
	return nil
}

func (m *RapidCheckOperatingRequest) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *RapidCheckOperatingRequest) GetMeasurementItems() []*RapidCheckOperatingRequestMeasurementItem {
	if m != nil {
		return m.MeasurementItems
	}
	return nil
}

type RapidCheckOperatingRequestMeasurementItem struct {
	EquipmentId string `protobuf:"bytes,1,opt,name=equipment_id,json=equipmentId,proto3" json:"equipment_id,omitempty"`
	RecipeId    string `protobuf:"bytes,2,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	SheetId     string `protobuf:"bytes,3,opt,name=sheet_id,json=sheetId,proto3" json:"sheet_id,omitempty"`
	// TODO this should be a enum
	ItemName   string `protobuf:"bytes,4,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	SequenceNo string `protobuf:"bytes,5,opt,name=sequence_no,json=sequenceNo,proto3" json:"sequence_no,omitempty"`
	Value      string `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	// TODO this should be a enum
	Unit                 string   `protobuf:"bytes,7,opt,name=unit,proto3" json:"unit,omitempty"`
	Conformity           bool     `protobuf:"varint,8,opt,name=conformity,proto3" json:"conformity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RapidCheckOperatingRequestMeasurementItem) Reset() {
	*m = RapidCheckOperatingRequestMeasurementItem{}
}
func (m *RapidCheckOperatingRequestMeasurementItem) String() string {
	return proto.CompactTextString(m)
}
func (*RapidCheckOperatingRequestMeasurementItem) ProtoMessage() {}
func (*RapidCheckOperatingRequestMeasurementItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2ad98225f42789, []int{0, 0}
}

func (m *RapidCheckOperatingRequestMeasurementItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RapidCheckOperatingRequestMeasurementItem.Unmarshal(m, b)
}
func (m *RapidCheckOperatingRequestMeasurementItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RapidCheckOperatingRequestMeasurementItem.Marshal(b, m, deterministic)
}
func (m *RapidCheckOperatingRequestMeasurementItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RapidCheckOperatingRequestMeasurementItem.Merge(m, src)
}
func (m *RapidCheckOperatingRequestMeasurementItem) XXX_Size() int {
	return xxx_messageInfo_RapidCheckOperatingRequestMeasurementItem.Size(m)
}
func (m *RapidCheckOperatingRequestMeasurementItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RapidCheckOperatingRequestMeasurementItem.DiscardUnknown(m)
}

var xxx_messageInfo_RapidCheckOperatingRequestMeasurementItem proto.InternalMessageInfo

func (m *RapidCheckOperatingRequestMeasurementItem) GetEquipmentId() string {
	if m != nil {
		return m.EquipmentId
	}
	return ""
}

func (m *RapidCheckOperatingRequestMeasurementItem) GetRecipeId() string {
	if m != nil {
		return m.RecipeId
	}
	return ""
}

func (m *RapidCheckOperatingRequestMeasurementItem) GetSheetId() string {
	if m != nil {
		return m.SheetId
	}
	return ""
}

func (m *RapidCheckOperatingRequestMeasurementItem) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *RapidCheckOperatingRequestMeasurementItem) GetSequenceNo() string {
	if m != nil {
		return m.SequenceNo
	}
	return ""
}

func (m *RapidCheckOperatingRequestMeasurementItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RapidCheckOperatingRequestMeasurementItem) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *RapidCheckOperatingRequestMeasurementItem) GetConformity() bool {
	if m != nil {
		return m.Conformity
	}
	return false
}

type RapidCheckOperatingReply struct {
	// 錯誤代碼
	Err                  error1.RapidCheckOperating        `protobuf:"varint,1,opt,name=err,proto3,enum=proto.mes.code.error.RapidCheckOperating" json:"err,omitempty"`
	Sheets               []*RapidCheckOperatingReply_Sheet `protobuf:"bytes,2,rep,name=sheets,proto3" json:"sheets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *RapidCheckOperatingReply) Reset()         { *m = RapidCheckOperatingReply{} }
func (m *RapidCheckOperatingReply) String() string { return proto.CompactTextString(m) }
func (*RapidCheckOperatingReply) ProtoMessage()    {}
func (*RapidCheckOperatingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2ad98225f42789, []int{1}
}

func (m *RapidCheckOperatingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RapidCheckOperatingReply.Unmarshal(m, b)
}
func (m *RapidCheckOperatingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RapidCheckOperatingReply.Marshal(b, m, deterministic)
}
func (m *RapidCheckOperatingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RapidCheckOperatingReply.Merge(m, src)
}
func (m *RapidCheckOperatingReply) XXX_Size() int {
	return xxx_messageInfo_RapidCheckOperatingReply.Size(m)
}
func (m *RapidCheckOperatingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RapidCheckOperatingReply.DiscardUnknown(m)
}

var xxx_messageInfo_RapidCheckOperatingReply proto.InternalMessageInfo

func (m *RapidCheckOperatingReply) GetErr() error1.RapidCheckOperating {
	if m != nil {
		return m.Err
	}
	return error1.RapidCheckOperating_RAPID_CHECK_OPERATING_SUCCESS
}

func (m *RapidCheckOperatingReply) GetSheets() []*RapidCheckOperatingReply_Sheet {
	if m != nil {
		return m.Sheets
	}
	return nil
}

type RapidCheckOperatingReply_Sheet struct {
	Id     string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LotId  string         `protobuf:"bytes,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	Status code.LotStatus `protobuf:"varint,3,opt,name=status,proto3,enum=proto.mes.code.LotStatus" json:"status,omitempty"`
	// 途程名稱
	CurrentProcessName string         `protobuf:"bytes,4,opt,name=current_process_name,json=currentProcessName,proto3" json:"current_process_name,omitempty"`
	CurrentProcessType rs.ProcessType `protobuf:"varint,5,opt,name=current_process_type,json=currentProcessType,proto3,enum=proto.dm.rs.ProcessType" json:"current_process_type,omitempty"`
	// 途程名稱
	NextProcessName      string         `protobuf:"bytes,6,opt,name=next_process_name,json=nextProcessName,proto3" json:"next_process_name,omitempty"`
	FactoryId            rs.FactoryID   `protobuf:"varint,7,opt,name=factory_id,json=factoryId,proto3,enum=proto.dm.rs.FactoryID" json:"factory_id,omitempty"`
	ProductId            string         `protobuf:"bytes,8,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StageStatus          rs.StageStatus `protobuf:"varint,9,opt,name=stage_status,json=stageStatus,proto3,enum=proto.dm.rs.StageStatus" json:"stage_status,omitempty"`
	MajorVersion         string         `protobuf:"bytes,10,opt,name=major_version,json=majorVersion,proto3" json:"major_version,omitempty"`
	UpdateBy             string         `protobuf:"bytes,11,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	UpdateAt             string         `protobuf:"bytes,12,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RapidCheckOperatingReply_Sheet) Reset()         { *m = RapidCheckOperatingReply_Sheet{} }
func (m *RapidCheckOperatingReply_Sheet) String() string { return proto.CompactTextString(m) }
func (*RapidCheckOperatingReply_Sheet) ProtoMessage()    {}
func (*RapidCheckOperatingReply_Sheet) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d2ad98225f42789, []int{1, 0}
}

func (m *RapidCheckOperatingReply_Sheet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RapidCheckOperatingReply_Sheet.Unmarshal(m, b)
}
func (m *RapidCheckOperatingReply_Sheet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RapidCheckOperatingReply_Sheet.Marshal(b, m, deterministic)
}
func (m *RapidCheckOperatingReply_Sheet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RapidCheckOperatingReply_Sheet.Merge(m, src)
}
func (m *RapidCheckOperatingReply_Sheet) XXX_Size() int {
	return xxx_messageInfo_RapidCheckOperatingReply_Sheet.Size(m)
}
func (m *RapidCheckOperatingReply_Sheet) XXX_DiscardUnknown() {
	xxx_messageInfo_RapidCheckOperatingReply_Sheet.DiscardUnknown(m)
}

var xxx_messageInfo_RapidCheckOperatingReply_Sheet proto.InternalMessageInfo

func (m *RapidCheckOperatingReply_Sheet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RapidCheckOperatingReply_Sheet) GetLotId() string {
	if m != nil {
		return m.LotId
	}
	return ""
}

func (m *RapidCheckOperatingReply_Sheet) GetStatus() code.LotStatus {
	if m != nil {
		return m.Status
	}
	return code.LotStatus_LOT_STATUS_UNSPECIFIED
}

func (m *RapidCheckOperatingReply_Sheet) GetCurrentProcessName() string {
	if m != nil {
		return m.CurrentProcessName
	}
	return ""
}

func (m *RapidCheckOperatingReply_Sheet) GetCurrentProcessType() rs.ProcessType {
	if m != nil {
		return m.CurrentProcessType
	}
	return rs.ProcessType_PROCESS_TYPE_UNSPECIFIED
}

func (m *RapidCheckOperatingReply_Sheet) GetNextProcessName() string {
	if m != nil {
		return m.NextProcessName
	}
	return ""
}

func (m *RapidCheckOperatingReply_Sheet) GetFactoryId() rs.FactoryID {
	if m != nil {
		return m.FactoryId
	}
	return rs.FactoryID_FACTORY_UNSPECIFIED
}

func (m *RapidCheckOperatingReply_Sheet) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *RapidCheckOperatingReply_Sheet) GetStageStatus() rs.StageStatus {
	if m != nil {
		return m.StageStatus
	}
	return rs.StageStatus_STAGE_STATUS_UNSPECIFIED
}

func (m *RapidCheckOperatingReply_Sheet) GetMajorVersion() string {
	if m != nil {
		return m.MajorVersion
	}
	return ""
}

func (m *RapidCheckOperatingReply_Sheet) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *RapidCheckOperatingReply_Sheet) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func init() {
	proto.RegisterType((*RapidCheckOperatingRequest)(nil), "proto.mes.RapidCheckOperatingRequest")
	proto.RegisterType((*RapidCheckOperatingRequestMeasurementItem)(nil), "proto.mes.RapidCheckOperatingRequest.measurementItem")
	proto.RegisterType((*RapidCheckOperatingReply)(nil), "proto.mes.RapidCheckOperatingReply")
	proto.RegisterType((*RapidCheckOperatingReply_Sheet)(nil), "proto.mes.RapidCheckOperatingReply.Sheet")
}

func init() { proto.RegisterFile("rapid_check_operating.proto", fileDescriptor_6d2ad98225f42789) }

var fileDescriptor_6d2ad98225f42789 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x85, 0x3f, 0x63, 0x5d, 0x7b, 0xc9, 0x42, 0x64, 0x9b, 0xe2, 0x20, 0x9b, 0x97, 0xbd, 0x38,
	0x7b, 0x90, 0x36, 0x0f, 0xc6, 0x1e, 0x02, 0x0c, 0x48, 0x56, 0x04, 0x70, 0xd1, 0xa6, 0x85, 0x52,
	0xf4, 0xa1, 0x2f, 0x02, 0x2d, 0xde, 0x38, 0x6a, 0x44, 0x51, 0x21, 0xa9, 0xb4, 0xfa, 0x2b, 0xfd,
	0x33, 0xfd, 0x53, 0x7d, 0x2d, 0x50, 0x90, 0x92, 0x1d, 0x45, 0x4d, 0xd1, 0x27, 0x91, 0xe7, 0xdc,
	0xc3, 0x7b, 0xef, 0xb9, 0xa4, 0xe0, 0x40, 0xd2, 0x2c, 0x66, 0x61, 0x74, 0x8d, 0xd1, 0x4d, 0x28,
	0x32, 0x94, 0x54, 0xc7, 0xe9, 0xca, 0xcb, 0xa4, 0xd0, 0x82, 0x38, 0xf6, 0xe3, 0x71, 0x54, 0xe3,
	0xb1, 0x5d, 0xfa, 0x1c, 0x95, 0x1f, 0x09, 0x86, 0x3e, 0xa6, 0x39, 0x57, 0x65, 0xd8, 0x78, 0xd2,
	0xe4, 0xa4, 0x14, 0xf2, 0x41, 0xc4, 0x2f, 0x65, 0x04, 0xe3, 0xbe, 0x54, 0x75, 0xe2, 0xe8, 0x63,
	0x17, 0xc6, 0x81, 0xa9, 0xe0, 0x7f, 0x53, 0xc0, 0x8b, 0x75, 0xfe, 0x00, 0x6f, 0x73, 0x54, 0x9a,
	0xfc, 0x07, 0x43, 0x1a, 0xe9, 0x58, 0xa4, 0xa1, 0x2e, 0x32, 0x74, 0x5b, 0x93, 0xd6, 0x74, 0x7b,
	0x76, 0xe8, 0x6d, 0xca, 0xf2, 0x4c, 0x3e, 0xef, 0x3c, 0x4f, 0x6d, 0xd0, 0x73, 0xc1, 0x30, 0x09,
	0xa0, 0x54, 0xbc, 0x2a, 0x32, 0x24, 0xf3, 0x8d, 0x9e, 0x0b, 0x86, 0x6e, 0xdb, 0xea, 0xf7, 0x9a,
	0x7a, 0xa3, 0x5b, 0xcb, 0xcc, 0x9a, 0xfc, 0x04, 0xfd, 0x44, 0xe8, 0x30, 0x66, 0x6e, 0x67, 0xd2,
	0x9a, 0x3a, 0x41, 0x2f, 0x11, 0x7a, 0xc1, 0xc8, 0x01, 0x38, 0xea, 0x1a, 0xd1, 0x10, 0xca, 0xed,
	0x4e, 0x3a, 0x53, 0x27, 0x18, 0x58, 0x60, 0xc1, 0x94, 0x21, 0xf3, 0x8c, 0x51, 0x8d, 0xe1, 0xb2,
	0x70, 0x7b, 0x56, 0x36, 0x28, 0x81, 0xb3, 0x82, 0x2c, 0x61, 0x97, 0x23, 0x55, 0xb9, 0x44, 0x8e,
	0xa9, 0x0e, 0x63, 0x8d, 0x5c, 0xb9, 0xfd, 0x49, 0x67, 0x3a, 0x9c, 0xcd, 0x6b, 0xd5, 0x7c, 0xdb,
	0x09, 0xaf, 0x26, 0x5f, 0x68, 0xe4, 0xc1, 0x8f, 0x0d, 0x40, 0x8d, 0x3f, 0xb5, 0x60, 0xa7, 0x01,
	0x92, 0xdf, 0x61, 0x84, 0xb7, 0x79, 0x9c, 0x95, 0x59, 0x99, 0x35, 0xd0, 0x09, 0x86, 0x1b, 0xac,
	0x6c, 0x4a, 0x62, 0x14, 0x67, 0x68, 0xf8, 0x76, 0x59, 0x77, 0x09, 0x2c, 0x18, 0xd9, 0x87, 0xc1,
	0xba, 0xe3, 0xca, 0x8a, 0xad, 0xaa, 0x61, 0xa3, 0x33, 0x6d, 0x84, 0x29, 0xe5, 0xe8, 0x76, 0x4b,
	0x9d, 0x01, 0x2e, 0x28, 0x47, 0xf2, 0x1b, 0x0c, 0x95, 0x29, 0x3c, 0x8d, 0x30, 0x4c, 0x45, 0x65,
	0x07, 0xac, 0xa1, 0x0b, 0x41, 0xf6, 0xa0, 0x77, 0x47, 0x93, 0x1c, 0xdd, 0x7e, 0x69, 0xb0, 0xdd,
	0x10, 0x02, 0xdd, 0x3c, 0x8d, 0xb5, 0xbb, 0x65, 0x41, 0xbb, 0x26, 0xbf, 0x02, 0x44, 0x22, 0xbd,
	0x12, 0x92, 0xc7, 0xba, 0x70, 0x07, 0x93, 0xd6, 0x74, 0x10, 0xd4, 0x90, 0xa3, 0x0f, 0x3d, 0x70,
	0x1f, 0xf5, 0x2d, 0x4b, 0x0a, 0x72, 0x02, 0x1d, 0x94, 0xb2, 0xba, 0x37, 0xc7, 0xcd, 0xb9, 0xdb,
	0x7b, 0xfa, 0xa8, 0xe9, 0x46, 0x45, 0x4e, 0xa1, 0x6f, 0x9b, 0x55, 0x6e, 0xdb, 0x4e, 0xea, 0xf8,
	0x7b, 0x93, 0xca, 0x92, 0xc2, 0xbb, 0x34, 0x8a, 0xa0, 0x12, 0x8e, 0x3f, 0x77, 0xa0, 0x67, 0x11,
	0xb2, 0x0d, 0xed, 0x8d, 0xff, 0xed, 0x98, 0xd5, 0xae, 0x58, 0xbb, 0x7e, 0xc5, 0xfe, 0x86, 0xbe,
	0xd2, 0x54, 0xe7, 0xca, 0xda, 0xbd, 0x3d, 0xdb, 0x6f, 0xd6, 0xfc, 0x4c, 0xe8, 0x4b, 0x1b, 0x10,
	0x54, 0x81, 0xe4, 0x2f, 0xd8, 0x8b, 0x72, 0x29, 0xcd, 0x84, 0x33, 0x29, 0x22, 0x54, 0xaa, 0x3e,
	0x13, 0x52, 0x71, 0x2f, 0x4b, 0xca, 0x4e, 0xe7, 0xe9, 0xd7, 0x0a, 0xfb, 0xbc, 0x7a, 0x36, 0xa5,
	0x5b, 0xa5, 0x64, 0xdc, 0x93, 0xca, 0xab, 0x74, 0xe6, 0x35, 0x35, 0xcf, 0xb2, 0x2f, 0xec, 0x4f,
	0xd8, 0x4d, 0xf1, 0x7d, 0x23, 0x75, 0x39, 0xd4, 0x1d, 0x43, 0xd4, 0xf3, 0xce, 0x01, 0xae, 0x68,
	0xa4, 0x85, 0x2c, 0x4c, 0xdf, 0x5b, 0x36, 0xdb, 0xcf, 0x0f, 0xb2, 0x9d, 0x97, 0xf4, 0xe2, 0x49,
	0xe0, 0x54, 0x91, 0x0b, 0x46, 0x0e, 0x01, 0x32, 0x29, 0x58, 0x1e, 0x59, 0xbb, 0x06, 0xf6, 0x6c,
	0xa7, 0x42, 0x16, 0x8c, 0x9c, 0xc0, 0x48, 0x69, 0xba, 0xc2, 0xb0, 0x32, 0xce, 0x79, 0xa4, 0x8b,
	0x4b, 0x13, 0x50, 0xf9, 0x36, 0x54, 0xf7, 0x1b, 0xf2, 0x07, 0xfc, 0xc0, 0xe9, 0x5b, 0x21, 0xc3,
	0x3b, 0x94, 0x2a, 0x16, 0xa9, 0x0b, 0xf6, 0xf8, 0x91, 0x05, 0x5f, 0x97, 0xd8, 0xc3, 0xa7, 0x3d,
	0x6c, 0x3c, 0xed, 0x7b, 0x92, 0x6a, 0x77, 0x54, 0x27, 0x4f, 0xf5, 0xd9, 0xbf, 0x6f, 0xe6, 0xab,
	0x58, 0x27, 0x74, 0xe9, 0xdd, 0x60, 0xca, 0xa8, 0x17, 0x09, 0xee, 0xe9, 0x77, 0xbe, 0xdd, 0xf8,
	0x91, 0xe0, 0x5c, 0xa4, 0xca, 0xbf, 0x9b, 0xf9, 0xe5, 0xef, 0x71, 0x25, 0x12, 0x9a, 0xae, 0xcc,
	0x7f, 0x74, 0xd9, 0xb7, 0xc8, 0x3f, 0x5f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xb4, 0x03, 0xab,
	0x9f, 0x05, 0x00, 0x00,
}