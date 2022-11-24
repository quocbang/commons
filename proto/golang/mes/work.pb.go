// Code generated by protoc-gen-go. DO NOT EDIT.
// source: work.proto

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

// 投料請求
// sheetCount 預設值為1
// moveWeight 不給時則依BOM計算
// materialLotId 除成型工程外其餘工程不須給值
type MoveInRequest struct {
	//投料收料基本資訊
	BasicInfo *MoveBasicInfo `protobuf:"bytes,1,opt,name=basic_info,json=basicInfo,proto3" json:"basic_info,omitempty"`
	//投入材料集合
	Materials            []*MoveInRequest_Material `protobuf:"bytes,2,rep,name=materials,proto3" json:"materials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MoveInRequest) Reset()         { *m = MoveInRequest{} }
func (m *MoveInRequest) String() string { return proto.CompactTextString(m) }
func (*MoveInRequest) ProtoMessage()    {}
func (*MoveInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8ca8b2cad64666d, []int{0}
}

func (m *MoveInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveInRequest.Unmarshal(m, b)
}
func (m *MoveInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveInRequest.Marshal(b, m, deterministic)
}
func (m *MoveInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveInRequest.Merge(m, src)
}
func (m *MoveInRequest) XXX_Size() int {
	return xxx_messageInfo_MoveInRequest.Size(m)
}
func (m *MoveInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveInRequest proto.InternalMessageInfo

func (m *MoveInRequest) GetBasicInfo() *MoveBasicInfo {
	if m != nil {
		return m.BasicInfo
	}
	return nil
}

func (m *MoveInRequest) GetMaterials() []*MoveInRequest_Material {
	if m != nil {
		return m.Materials
	}
	return nil
}

// 標示卡資訊
type MoveInRequest_Material struct {
	//材料標示卡條碼
	LotId string `protobuf:"bytes,1,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	//投入量 : 密煉製程計量單位為kg, 精度小數三位, 其他製程待訂定
	UseQty float32 `protobuf:"fixed32,2,opt,name=use_qty,json=useQty,proto3" json:"use_qty,omitempty"`
	//材料 ID
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	//材料等級
	Level string `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	//是否為真實重量
	RealWeight           bool     `protobuf:"varint,5,opt,name=realWeight,proto3" json:"realWeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveInRequest_Material) Reset()         { *m = MoveInRequest_Material{} }
func (m *MoveInRequest_Material) String() string { return proto.CompactTextString(m) }
func (*MoveInRequest_Material) ProtoMessage()    {}
func (*MoveInRequest_Material) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8ca8b2cad64666d, []int{0, 0}
}

func (m *MoveInRequest_Material) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveInRequest_Material.Unmarshal(m, b)
}
func (m *MoveInRequest_Material) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveInRequest_Material.Marshal(b, m, deterministic)
}
func (m *MoveInRequest_Material) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveInRequest_Material.Merge(m, src)
}
func (m *MoveInRequest_Material) XXX_Size() int {
	return xxx_messageInfo_MoveInRequest_Material.Size(m)
}
func (m *MoveInRequest_Material) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveInRequest_Material.DiscardUnknown(m)
}

var xxx_messageInfo_MoveInRequest_Material proto.InternalMessageInfo

func (m *MoveInRequest_Material) GetLotId() string {
	if m != nil {
		return m.LotId
	}
	return ""
}

func (m *MoveInRequest_Material) GetUseQty() float32 {
	if m != nil {
		return m.UseQty
	}
	return 0
}

func (m *MoveInRequest_Material) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *MoveInRequest_Material) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *MoveInRequest_Material) GetRealWeight() bool {
	if m != nil {
		return m.RealWeight
	}
	return false
}

//投料回應
type MoveInReply struct {
	// 錯誤代碼
	Err                  error1.MoveIn `protobuf:"varint,1,opt,name=err,proto3,enum=proto.mes.code.error.MoveIn" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MoveInReply) Reset()         { *m = MoveInReply{} }
func (m *MoveInReply) String() string { return proto.CompactTextString(m) }
func (*MoveInReply) ProtoMessage()    {}
func (*MoveInReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8ca8b2cad64666d, []int{1}
}

func (m *MoveInReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveInReply.Unmarshal(m, b)
}
func (m *MoveInReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveInReply.Marshal(b, m, deterministic)
}
func (m *MoveInReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveInReply.Merge(m, src)
}
func (m *MoveInReply) XXX_Size() int {
	return xxx_messageInfo_MoveInReply.Size(m)
}
func (m *MoveInReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveInReply.DiscardUnknown(m)
}

var xxx_messageInfo_MoveInReply proto.InternalMessageInfo

func (m *MoveInReply) GetErr() error1.MoveIn {
	if m != nil {
		return m.Err
	}
	return error1.MoveIn_MOVE_IN_SUCCESS
}

//收料請求
type MoveOutRequest struct {
	// 收料基本資訊
	BasicInfo            *MoveBasicInfo `protobuf:"bytes,1,opt,name=basic_info,json=basicInfo,proto3" json:"basic_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MoveOutRequest) Reset()         { *m = MoveOutRequest{} }
func (m *MoveOutRequest) String() string { return proto.CompactTextString(m) }
func (*MoveOutRequest) ProtoMessage()    {}
func (*MoveOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8ca8b2cad64666d, []int{2}
}

func (m *MoveOutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveOutRequest.Unmarshal(m, b)
}
func (m *MoveOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveOutRequest.Marshal(b, m, deterministic)
}
func (m *MoveOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveOutRequest.Merge(m, src)
}
func (m *MoveOutRequest) XXX_Size() int {
	return xxx_messageInfo_MoveOutRequest.Size(m)
}
func (m *MoveOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveOutRequest proto.InternalMessageInfo

func (m *MoveOutRequest) GetBasicInfo() *MoveBasicInfo {
	if m != nil {
		return m.BasicInfo
	}
	return nil
}

//收料回應
type MoveOutReply struct {
	// 錯誤代碼
	Err                  error1.MoveOut `protobuf:"varint,1,opt,name=err,proto3,enum=proto.mes.code.error.MoveOut" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MoveOutReply) Reset()         { *m = MoveOutReply{} }
func (m *MoveOutReply) String() string { return proto.CompactTextString(m) }
func (*MoveOutReply) ProtoMessage()    {}
func (*MoveOutReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8ca8b2cad64666d, []int{3}
}

func (m *MoveOutReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveOutReply.Unmarshal(m, b)
}
func (m *MoveOutReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveOutReply.Marshal(b, m, deterministic)
}
func (m *MoveOutReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveOutReply.Merge(m, src)
}
func (m *MoveOutReply) XXX_Size() int {
	return xxx_messageInfo_MoveOutReply.Size(m)
}
func (m *MoveOutReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveOutReply.DiscardUnknown(m)
}

var xxx_messageInfo_MoveOutReply proto.InternalMessageInfo

func (m *MoveOutReply) GetErr() error1.MoveOut {
	if m != nil {
		return m.Err
	}
	return error1.MoveOut_MOVE_OUT_SUCCESS
}

// 投料/收料結束請求
// 發出請求以改變工單狀態
type MoveEndRequest struct {
	// 請求類別
	// I = 投料 , O = 收料
	ActionMode code.Mode `protobuf:"varint,1,opt,name=action_mode,json=actionMode,proto3,enum=proto.mes.code.Mode" json:"action_mode,omitempty"`
	//工單號
	LotId string `protobuf:"bytes,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	//途程名稱
	ProcessName string `protobuf:"bytes,3,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	//操作者
	UpdateBy             string   `protobuf:"bytes,4,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveEndRequest) Reset()         { *m = MoveEndRequest{} }
func (m *MoveEndRequest) String() string { return proto.CompactTextString(m) }
func (*MoveEndRequest) ProtoMessage()    {}
func (*MoveEndRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8ca8b2cad64666d, []int{4}
}

func (m *MoveEndRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveEndRequest.Unmarshal(m, b)
}
func (m *MoveEndRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveEndRequest.Marshal(b, m, deterministic)
}
func (m *MoveEndRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveEndRequest.Merge(m, src)
}
func (m *MoveEndRequest) XXX_Size() int {
	return xxx_messageInfo_MoveEndRequest.Size(m)
}
func (m *MoveEndRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveEndRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveEndRequest proto.InternalMessageInfo

func (m *MoveEndRequest) GetActionMode() code.Mode {
	if m != nil {
		return m.ActionMode
	}
	return code.Mode_MODE_UNSPECIFIED
}

func (m *MoveEndRequest) GetLotId() string {
	if m != nil {
		return m.LotId
	}
	return ""
}

func (m *MoveEndRequest) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *MoveEndRequest) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

// 投料/收料結束回應
type MoveEndReply struct {
	// 錯誤代碼
	Err                  error1.MoveEnd `protobuf:"varint,1,opt,name=err,proto3,enum=proto.mes.code.error.MoveEnd" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MoveEndReply) Reset()         { *m = MoveEndReply{} }
func (m *MoveEndReply) String() string { return proto.CompactTextString(m) }
func (*MoveEndReply) ProtoMessage()    {}
func (*MoveEndReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8ca8b2cad64666d, []int{5}
}

func (m *MoveEndReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveEndReply.Unmarshal(m, b)
}
func (m *MoveEndReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveEndReply.Marshal(b, m, deterministic)
}
func (m *MoveEndReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveEndReply.Merge(m, src)
}
func (m *MoveEndReply) XXX_Size() int {
	return xxx_messageInfo_MoveEndReply.Size(m)
}
func (m *MoveEndReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveEndReply.DiscardUnknown(m)
}

var xxx_messageInfo_MoveEndReply proto.InternalMessageInfo

func (m *MoveEndReply) GetErr() error1.MoveEnd {
	if m != nil {
		return m.Err
	}
	return error1.MoveEnd_MOVE_END_SUCCESS
}

func init() {
	proto.RegisterType((*MoveInRequest)(nil), "proto.mes.MoveInRequest")
	proto.RegisterType((*MoveInRequest_Material)(nil), "proto.mes.MoveInRequest.Material")
	proto.RegisterType((*MoveInReply)(nil), "proto.mes.MoveInReply")
	proto.RegisterType((*MoveOutRequest)(nil), "proto.mes.MoveOutRequest")
	proto.RegisterType((*MoveOutReply)(nil), "proto.mes.MoveOutReply")
	proto.RegisterType((*MoveEndRequest)(nil), "proto.mes.MoveEndRequest")
	proto.RegisterType((*MoveEndReply)(nil), "proto.mes.MoveEndReply")
}

func init() { proto.RegisterFile("work.proto", fileDescriptor_b8ca8b2cad64666d) }

var fileDescriptor_b8ca8b2cad64666d = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0x86, 0x71, 0xb6, 0x49, 0xe3, 0xc9, 0x76, 0x0f, 0x62, 0x4b, 0x4d, 0xda, 0x2d, 0x49, 0x4e,
	0x39, 0xc9, 0x90, 0xb2, 0xec, 0xa9, 0x2c, 0x04, 0x7a, 0xc8, 0x21, 0x2d, 0xf5, 0xa5, 0xd0, 0x8b,
	0x51, 0xac, 0xd9, 0xd4, 0xac, 0x25, 0x65, 0x25, 0x39, 0x4b, 0x9e, 0xa2, 0x0f, 0x50, 0xe8, 0xb3,
	0x16, 0xc9, 0x72, 0x62, 0x02, 0x3d, 0x94, 0x9e, 0x22, 0xfd, 0xfc, 0xff, 0xe4, 0x1b, 0x8d, 0x07,
	0xe0, 0x59, 0xe9, 0x47, 0xba, 0xd3, 0xca, 0x2a, 0x12, 0xfb, 0x1f, 0x2a, 0xd0, 0x8c, 0x27, 0xfe,
	0x98, 0x0a, 0x34, 0x69, 0xa1, 0x38, 0xa6, 0xa8, 0xb5, 0xd2, 0x29, 0xca, 0x5a, 0x98, 0xc6, 0x3c,
	0x9e, 0x9e, 0x1c, 0x3b, 0xad, 0x78, 0x5d, 0xd8, 0x52, 0xc9, 0xb4, 0x50, 0x42, 0x28, 0x19, 0x2c,
	0xe3, 0xf3, 0x22, 0xa7, 0xf8, 0xec, 0x57, 0x0f, 0x5e, 0xad, 0xd5, 0x1e, 0x57, 0x32, 0xc3, 0xa7,
	0x1a, 0x8d, 0x25, 0x77, 0x00, 0x1b, 0x66, 0xca, 0x22, 0x2f, 0xe5, 0x83, 0x4a, 0xa2, 0x49, 0x34,
	0x1f, 0x2d, 0x12, 0x7a, 0x44, 0xa2, 0xce, 0xbd, 0x74, 0x86, 0x95, 0x7c, 0x50, 0x59, 0xbc, 0x69,
	0x8f, 0xe4, 0x1e, 0x62, 0xc1, 0x2c, 0xea, 0x92, 0x55, 0x26, 0xe9, 0x4d, 0x2e, 0xe6, 0xa3, 0xc5,
	0xf4, 0x2c, 0x77, 0xfc, 0x17, 0xba, 0x0e, 0xce, 0xec, 0x94, 0x19, 0xff, 0x8c, 0x60, 0xd8, 0xea,
	0xe4, 0x35, 0x0c, 0x2a, 0x65, 0xf3, 0x92, 0x7b, 0x84, 0x38, 0xeb, 0x57, 0xca, 0xae, 0x38, 0x79,
	0x03, 0x2f, 0x6b, 0x83, 0xf9, 0x93, 0x3d, 0x24, 0xbd, 0x49, 0x34, 0xef, 0x65, 0x83, 0xda, 0xe0,
	0x57, 0x7b, 0x20, 0x37, 0x00, 0xa1, 0x7f, 0x97, 0xb9, 0xf0, 0x99, 0x38, 0x28, 0x2b, 0x4e, 0xae,
	0xa1, 0x5f, 0xe1, 0x1e, 0xab, 0xe4, 0x45, 0xa8, 0xe6, 0x2e, 0xe4, 0x3d, 0x80, 0x46, 0x56, 0x7d,
	0xc3, 0x72, 0xfb, 0xc3, 0x26, 0xfd, 0x49, 0x34, 0x1f, 0x66, 0x1d, 0x65, 0xf6, 0x11, 0x46, 0x2d,
	0xf6, 0xae, 0x3a, 0x10, 0x0a, 0x17, 0xa8, 0xb5, 0x07, 0xba, 0x5a, 0xbc, 0xeb, 0xf4, 0xe6, 0x9e,
	0x95, 0xfa, 0xd9, 0xb4, 0x6d, 0x3a, 0xe3, 0x6c, 0x05, 0x57, 0xee, 0xfa, 0xa5, 0xb6, 0xff, 0xfb,
	0xb8, 0xb3, 0x7b, 0xb8, 0x3c, 0x96, 0x72, 0x28, 0x69, 0x17, 0xe5, 0xe6, 0xef, 0x28, 0x2e, 0xe0,
	0x59, 0x7e, 0x47, 0x0d, 0xcc, 0x27, 0xc9, 0x5b, 0x98, 0x5b, 0x18, 0x31, 0xff, 0xb9, 0xe4, 0x42,
	0x71, 0x0c, 0xb5, 0xae, 0xcf, 0x6b, 0xad, 0x15, 0xc7, 0x0c, 0x1a, 0xa3, 0x3b, 0x77, 0x26, 0xd3,
	0xeb, 0x4e, 0x66, 0x0a, 0x97, 0x3b, 0xad, 0x0a, 0x34, 0x26, 0x97, 0x4c, 0x60, 0x18, 0xc1, 0x28,
	0x68, 0x9f, 0x99, 0x40, 0xf2, 0x16, 0xe2, 0x7a, 0xc7, 0x99, 0xc5, 0x7c, 0x73, 0x08, 0x83, 0x18,
	0x36, 0xc2, 0xf2, 0xd0, 0x76, 0xe8, 0xf9, 0xfe, 0xa5, 0x43, 0x17, 0x70, 0xce, 0xe5, 0xdd, 0xf7,
	0xdb, 0x6d, 0x69, 0x2b, 0xb6, 0xa1, 0x8f, 0x28, 0x39, 0xa3, 0x85, 0x12, 0xd4, 0x3e, 0xa7, 0xfe,
	0x12, 0x16, 0xc2, 0xa4, 0xfb, 0x45, 0xda, 0x6c, 0xc3, 0x56, 0x55, 0x4c, 0x6e, 0xdd, 0x52, 0x6c,
	0x06, 0x5e, 0xf9, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xd9, 0xa1, 0x4c, 0x84, 0x03, 0x00,
	0x00,
}