// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_log_record.proto

package mes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LogType int32

const (
	LogType_NO_THIS_Log_Type LogType = 0
	LogType_ERROR            LogType = 1
	LogType_MODIFY           LogType = 2
)

var LogType_name = map[int32]string{
	0: "NO_THIS_Log_Type",
	1: "ERROR",
	2: "MODIFY",
}

var LogType_value = map[string]int32{
	"NO_THIS_Log_Type": 0,
	"ERROR":            1,
	"MODIFY":           2,
}

func (x LogType) String() string {
	return proto.EnumName(LogType_name, int32(x))
}

func (LogType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_addd2eb7948f2eba, []int{0}
}

type AddLogRecordRequest struct {
	Type                 LogType              `protobuf:"varint,1,opt,name=type,proto3,enum=proto.dm.mes.LogType" json:"type,omitempty"`
	Factory              string               `protobuf:"bytes,2,opt,name=factory,proto3" json:"factory,omitempty"`
	ProductCate          string               `protobuf:"bytes,3,opt,name=product_cate,json=productCate,proto3" json:"product_cate,omitempty"`
	ProductId            string               `protobuf:"bytes,4,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	MtrlId               string               `protobuf:"bytes,5,opt,name=mtrl_id,json=mtrlId,proto3" json:"mtrl_id,omitempty"`
	LotId                string               `protobuf:"bytes,6,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	EqptId               string               `protobuf:"bytes,7,opt,name=eqpt_id,json=eqptId,proto3" json:"eqpt_id,omitempty"`
	OldState             string               `protobuf:"bytes,8,opt,name=old_state,json=oldState,proto3" json:"old_state,omitempty"`
	NewState             string               `protobuf:"bytes,9,opt,name=new_state,json=newState,proto3" json:"new_state,omitempty"`
	OldLimdat            string               `protobuf:"bytes,10,opt,name=old_limdat,json=oldLimdat,proto3" json:"old_limdat,omitempty"`
	NewLimdat            string               `protobuf:"bytes,11,opt,name=new_limdat,json=newLimdat,proto3" json:"new_limdat,omitempty"`
	OverDay              int32                `protobuf:"varint,12,opt,name=over_day,json=overDay,proto3" json:"over_day,omitempty"`
	Reason               string               `protobuf:"bytes,13,opt,name=reason,proto3" json:"reason,omitempty"`
	Region               string               `protobuf:"bytes,14,opt,name=region,proto3" json:"region,omitempty"`
	Usrno                string               `protobuf:"bytes,15,opt,name=usrno,proto3" json:"usrno,omitempty"`
	Intime               *timestamp.Timestamp `protobuf:"bytes,16,opt,name=intime,proto3" json:"intime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AddLogRecordRequest) Reset()         { *m = AddLogRecordRequest{} }
func (m *AddLogRecordRequest) String() string { return proto.CompactTextString(m) }
func (*AddLogRecordRequest) ProtoMessage()    {}
func (*AddLogRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_addd2eb7948f2eba, []int{0}
}

func (m *AddLogRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddLogRecordRequest.Unmarshal(m, b)
}
func (m *AddLogRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddLogRecordRequest.Marshal(b, m, deterministic)
}
func (m *AddLogRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddLogRecordRequest.Merge(m, src)
}
func (m *AddLogRecordRequest) XXX_Size() int {
	return xxx_messageInfo_AddLogRecordRequest.Size(m)
}
func (m *AddLogRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddLogRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddLogRecordRequest proto.InternalMessageInfo

func (m *AddLogRecordRequest) GetType() LogType {
	if m != nil {
		return m.Type
	}
	return LogType_NO_THIS_Log_Type
}

func (m *AddLogRecordRequest) GetFactory() string {
	if m != nil {
		return m.Factory
	}
	return ""
}

func (m *AddLogRecordRequest) GetProductCate() string {
	if m != nil {
		return m.ProductCate
	}
	return ""
}

func (m *AddLogRecordRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *AddLogRecordRequest) GetMtrlId() string {
	if m != nil {
		return m.MtrlId
	}
	return ""
}

func (m *AddLogRecordRequest) GetLotId() string {
	if m != nil {
		return m.LotId
	}
	return ""
}

func (m *AddLogRecordRequest) GetEqptId() string {
	if m != nil {
		return m.EqptId
	}
	return ""
}

func (m *AddLogRecordRequest) GetOldState() string {
	if m != nil {
		return m.OldState
	}
	return ""
}

func (m *AddLogRecordRequest) GetNewState() string {
	if m != nil {
		return m.NewState
	}
	return ""
}

func (m *AddLogRecordRequest) GetOldLimdat() string {
	if m != nil {
		return m.OldLimdat
	}
	return ""
}

func (m *AddLogRecordRequest) GetNewLimdat() string {
	if m != nil {
		return m.NewLimdat
	}
	return ""
}

func (m *AddLogRecordRequest) GetOverDay() int32 {
	if m != nil {
		return m.OverDay
	}
	return 0
}

func (m *AddLogRecordRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *AddLogRecordRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AddLogRecordRequest) GetUsrno() string {
	if m != nil {
		return m.Usrno
	}
	return ""
}

func (m *AddLogRecordRequest) GetIntime() *timestamp.Timestamp {
	if m != nil {
		return m.Intime
	}
	return nil
}

type AddLogRecordReply struct {
	RowsAffect           int32    `protobuf:"varint,1,opt,name=rows_affect,json=rowsAffect,proto3" json:"rows_affect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddLogRecordReply) Reset()         { *m = AddLogRecordReply{} }
func (m *AddLogRecordReply) String() string { return proto.CompactTextString(m) }
func (*AddLogRecordReply) ProtoMessage()    {}
func (*AddLogRecordReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_addd2eb7948f2eba, []int{1}
}

func (m *AddLogRecordReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddLogRecordReply.Unmarshal(m, b)
}
func (m *AddLogRecordReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddLogRecordReply.Marshal(b, m, deterministic)
}
func (m *AddLogRecordReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddLogRecordReply.Merge(m, src)
}
func (m *AddLogRecordReply) XXX_Size() int {
	return xxx_messageInfo_AddLogRecordReply.Size(m)
}
func (m *AddLogRecordReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddLogRecordReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddLogRecordReply proto.InternalMessageInfo

func (m *AddLogRecordReply) GetRowsAffect() int32 {
	if m != nil {
		return m.RowsAffect
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.dm.mes.LogType", LogType_name, LogType_value)
	proto.RegisterType((*AddLogRecordRequest)(nil), "proto.dm.mes.AddLogRecordRequest")
	proto.RegisterType((*AddLogRecordReply)(nil), "proto.dm.mes.AddLogRecordReply")
}

func init() { proto.RegisterFile("add_log_record.proto", fileDescriptor_addd2eb7948f2eba) }

var fileDescriptor_addd2eb7948f2eba = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x51, 0x6f, 0xd3, 0x3c,
	0x18, 0x85, 0xbf, 0x6c, 0x4b, 0xda, 0xbe, 0xdd, 0x37, 0x8a, 0xe9, 0xc0, 0x0c, 0x4d, 0x2b, 0xbb,
	0x2a, 0x5c, 0x38, 0x52, 0x41, 0x08, 0x71, 0x37, 0x18, 0x88, 0x4a, 0x85, 0x4a, 0x59, 0x6f, 0xe0,
	0x26, 0x72, 0x63, 0x37, 0x8a, 0xb0, 0xf3, 0x66, 0x89, 0xbb, 0x2a, 0x7f, 0x9b, 0x5f, 0x80, 0x6c,
	0x27, 0x12, 0x5c, 0x55, 0xe7, 0x79, 0x8e, 0x2d, 0x57, 0x27, 0x30, 0xe5, 0x42, 0xa4, 0x0a, 0xf3,
	0xb4, 0x96, 0x19, 0xd6, 0x82, 0x55, 0x35, 0x1a, 0x24, 0xa7, 0xee, 0x87, 0x09, 0xcd, 0xb4, 0x6c,
	0x2e, 0xae, 0x72, 0xc4, 0x5c, 0xc9, 0xd8, 0xc1, 0xed, 0x7e, 0x17, 0x9b, 0x42, 0xcb, 0xc6, 0x70,
	0x5d, 0xf9, 0xfa, 0xf5, 0xef, 0x63, 0x78, 0x72, 0x23, 0xc4, 0x0a, 0xf3, 0xc4, 0xdd, 0x92, 0xc8,
	0xfb, 0xbd, 0x6c, 0x0c, 0x79, 0x05, 0x27, 0xa6, 0xad, 0x24, 0x0d, 0x66, 0xc1, 0xfc, 0x6c, 0x71,
	0xce, 0xfe, 0xbe, 0x95, 0xad, 0x30, 0xdf, 0xb4, 0x95, 0x4c, 0x5c, 0x85, 0x50, 0x18, 0xec, 0x78,
	0x66, 0xb0, 0x6e, 0xe9, 0xd1, 0x2c, 0x98, 0x8f, 0x92, 0x3e, 0x92, 0x97, 0x60, 0x5f, 0x23, 0xf6,
	0x99, 0x49, 0x33, 0x6e, 0x24, 0x3d, 0x76, 0x7a, 0xdc, 0xb1, 0x4f, 0xdc, 0x48, 0x72, 0x09, 0xd0,
	0x57, 0x0a, 0x41, 0x4f, 0x5c, 0x61, 0xd4, 0x91, 0xa5, 0x20, 0xcf, 0x60, 0xa0, 0x4d, 0xad, 0xac,
	0x0b, 0x9d, 0x8b, 0x6c, 0x5c, 0x0a, 0x72, 0x0e, 0x91, 0x42, 0x77, 0x26, 0x72, 0x3c, 0x54, 0xd8,
	0xf5, 0xe5, 0x7d, 0xe5, 0xf8, 0xc0, 0xf7, 0x6d, 0x5c, 0x0a, 0xf2, 0x02, 0x46, 0xa8, 0x44, 0xda,
	0x18, 0xfb, 0x8e, 0xa1, 0x53, 0x43, 0x54, 0xe2, 0xce, 0x66, 0x2b, 0x4b, 0x79, 0xe8, 0xe4, 0xc8,
	0xcb, 0x52, 0x1e, 0xbc, 0xbc, 0x04, 0xb0, 0x27, 0x55, 0xa1, 0x05, 0x37, 0x14, 0xfc, 0x0b, 0x51,
	0x89, 0x95, 0x03, 0x56, 0xdb, 0xb3, 0x9d, 0x1e, 0x7b, 0x5d, 0xca, 0x43, 0xa7, 0x9f, 0xc3, 0x10,
	0x1f, 0x64, 0x9d, 0x0a, 0xde, 0xd2, 0xd3, 0x59, 0x30, 0x0f, 0x93, 0x81, 0xcd, 0xb7, 0xbc, 0x25,
	0x4f, 0x21, 0xaa, 0x25, 0x6f, 0xb0, 0xa4, 0xff, 0xfb, 0xa7, 0xfa, 0xe4, 0x79, 0x5e, 0x60, 0x49,
	0xcf, 0x7a, 0x6e, 0x13, 0x99, 0x42, 0xb8, 0x6f, 0xea, 0x12, 0xe9, 0x23, 0xff, 0x8f, 0x5d, 0x20,
	0x0b, 0x88, 0x8a, 0xd2, 0xae, 0x4a, 0x27, 0xb3, 0x60, 0x3e, 0x5e, 0x5c, 0x30, 0x3f, 0x39, 0xeb,
	0x27, 0x67, 0x9b, 0x7e, 0xf2, 0xa4, 0x6b, 0x5e, 0xbf, 0x85, 0xc7, 0xff, 0x6e, 0x5e, 0xa9, 0x96,
	0x5c, 0xc1, 0xb8, 0xc6, 0x43, 0x93, 0xf2, 0xdd, 0x4e, 0x66, 0xc6, 0x0d, 0x1f, 0x26, 0x60, 0xd1,
	0x8d, 0x23, 0xaf, 0xdf, 0xc1, 0xa0, 0x1b, 0x9e, 0x4c, 0x61, 0xf2, 0x7d, 0x9d, 0x6e, 0xbe, 0x2e,
	0xef, 0xd2, 0x15, 0xe6, 0xa9, 0x65, 0x93, 0xff, 0xc8, 0x08, 0xc2, 0xcf, 0x49, 0xb2, 0x4e, 0x26,
	0x01, 0x01, 0x88, 0xbe, 0xad, 0x6f, 0x97, 0x5f, 0x7e, 0x4c, 0x8e, 0x3e, 0x7e, 0xf8, 0xf9, 0x3e,
	0x2f, 0x8c, 0xe2, 0x5b, 0xf6, 0x4b, 0x96, 0x82, 0xb3, 0x0c, 0x35, 0x33, 0x87, 0xd8, 0x85, 0x38,
	0x43, 0xad, 0xb1, 0x6c, 0xe2, 0x87, 0x85, 0xff, 0x46, 0xe3, 0x1c, 0x15, 0x2f, 0xf3, 0x58, 0xe8,
	0x58, 0xcb, 0x66, 0x1b, 0x39, 0xf8, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0xcd, 0x6d,
	0xd2, 0xec, 0x02, 0x00, 0x00,
}
