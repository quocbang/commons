// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package cdr

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type Product struct {
	Subno                string   `protobuf:"bytes,1,opt,name=subno,proto3" json:"subno,omitempty"`
	Itnbr                string   `protobuf:"bytes,2,opt,name=itnbr,proto3" json:"itnbr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetSubno() string {
	if m != nil {
		return m.Subno
	}
	return ""
}

func (m *Product) GetItnbr() string {
	if m != nil {
		return m.Itnbr
	}
	return ""
}

type ScrappingTireFee struct {
	Fee                  float32  `protobuf:"fixed32,1,opt,name=fee,proto3" json:"fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScrappingTireFee) Reset()         { *m = ScrappingTireFee{} }
func (m *ScrappingTireFee) String() string { return proto.CompactTextString(m) }
func (*ScrappingTireFee) ProtoMessage()    {}
func (*ScrappingTireFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ScrappingTireFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScrappingTireFee.Unmarshal(m, b)
}
func (m *ScrappingTireFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScrappingTireFee.Marshal(b, m, deterministic)
}
func (m *ScrappingTireFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScrappingTireFee.Merge(m, src)
}
func (m *ScrappingTireFee) XXX_Size() int {
	return xxx_messageInfo_ScrappingTireFee.Size(m)
}
func (m *ScrappingTireFee) XXX_DiscardUnknown() {
	xxx_messageInfo_ScrappingTireFee.DiscardUnknown(m)
}

var xxx_messageInfo_ScrappingTireFee proto.InternalMessageInfo

func (m *ScrappingTireFee) GetFee() float32 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type Cusno struct {
	Cusno                string   `protobuf:"bytes,1,opt,name=cusno,proto3" json:"cusno,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cusno) Reset()         { *m = Cusno{} }
func (m *Cusno) String() string { return proto.CompactTextString(m) }
func (*Cusno) ProtoMessage()    {}
func (*Cusno) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *Cusno) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cusno.Unmarshal(m, b)
}
func (m *Cusno) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cusno.Marshal(b, m, deterministic)
}
func (m *Cusno) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cusno.Merge(m, src)
}
func (m *Cusno) XXX_Size() int {
	return xxx_messageInfo_Cusno.Size(m)
}
func (m *Cusno) XXX_DiscardUnknown() {
	xxx_messageInfo_Cusno.DiscardUnknown(m)
}

var xxx_messageInfo_Cusno proto.InternalMessageInfo

func (m *Cusno) GetCusno() string {
	if m != nil {
		return m.Cusno
	}
	return ""
}

type CustomReport struct {
	CustomName           string   `protobuf:"bytes,1,opt,name=CustomName,proto3" json:"CustomName,omitempty"`
	Stdqty               int64    `protobuf:"varint,2,opt,name=Stdqty,proto3" json:"Stdqty,omitempty"`
	Stdbqty              int64    `protobuf:"varint,3,opt,name=Stdbqty,proto3" json:"Stdbqty,omitempty"`
	Cdrqty               int64    `protobuf:"varint,4,opt,name=Cdrqty,proto3" json:"Cdrqty,omitempty"`
	Shpqty               int64    `protobuf:"varint,5,opt,name=Shpqty,proto3" json:"Shpqty,omitempty"`
	Amount               int64    `protobuf:"varint,6,opt,name=Amount,proto3" json:"Amount,omitempty"`
	LtAmount             int64    `protobuf:"varint,7,opt,name=LtAmount,proto3" json:"LtAmount,omitempty"`
	Qty1                 int64    `protobuf:"varint,8,opt,name=Qty1,proto3" json:"Qty1,omitempty"`
	ResponseRate         float32  `protobuf:"fixed32,9,opt,name=ResponseRate,proto3" json:"ResponseRate,omitempty"`
	SixteenInchRate      float32  `protobuf:"fixed32,10,opt,name=SixteenInchRate,proto3" json:"SixteenInchRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomReport) Reset()         { *m = CustomReport{} }
func (m *CustomReport) String() string { return proto.CompactTextString(m) }
func (*CustomReport) ProtoMessage()    {}
func (*CustomReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *CustomReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomReport.Unmarshal(m, b)
}
func (m *CustomReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomReport.Marshal(b, m, deterministic)
}
func (m *CustomReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomReport.Merge(m, src)
}
func (m *CustomReport) XXX_Size() int {
	return xxx_messageInfo_CustomReport.Size(m)
}
func (m *CustomReport) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomReport.DiscardUnknown(m)
}

var xxx_messageInfo_CustomReport proto.InternalMessageInfo

func (m *CustomReport) GetCustomName() string {
	if m != nil {
		return m.CustomName
	}
	return ""
}

func (m *CustomReport) GetStdqty() int64 {
	if m != nil {
		return m.Stdqty
	}
	return 0
}

func (m *CustomReport) GetStdbqty() int64 {
	if m != nil {
		return m.Stdbqty
	}
	return 0
}

func (m *CustomReport) GetCdrqty() int64 {
	if m != nil {
		return m.Cdrqty
	}
	return 0
}

func (m *CustomReport) GetShpqty() int64 {
	if m != nil {
		return m.Shpqty
	}
	return 0
}

func (m *CustomReport) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CustomReport) GetLtAmount() int64 {
	if m != nil {
		return m.LtAmount
	}
	return 0
}

func (m *CustomReport) GetQty1() int64 {
	if m != nil {
		return m.Qty1
	}
	return 0
}

func (m *CustomReport) GetResponseRate() float32 {
	if m != nil {
		return m.ResponseRate
	}
	return 0
}

func (m *CustomReport) GetSixteenInchRate() float32 {
	if m != nil {
		return m.SixteenInchRate
	}
	return 0
}

type CustomReports struct {
	Dataset              []*CustomReport `protobuf:"bytes,1,rep,name=dataset,proto3" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CustomReports) Reset()         { *m = CustomReports{} }
func (m *CustomReports) String() string { return proto.CompactTextString(m) }
func (*CustomReports) ProtoMessage()    {}
func (*CustomReports) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *CustomReports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomReports.Unmarshal(m, b)
}
func (m *CustomReports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomReports.Marshal(b, m, deterministic)
}
func (m *CustomReports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomReports.Merge(m, src)
}
func (m *CustomReports) XXX_Size() int {
	return xxx_messageInfo_CustomReports.Size(m)
}
func (m *CustomReports) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomReports.DiscardUnknown(m)
}

var xxx_messageInfo_CustomReports proto.InternalMessageInfo

func (m *CustomReports) GetDataset() []*CustomReport {
	if m != nil {
		return m.Dataset
	}
	return nil
}

//數量
type CountReply struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountReply) Reset()         { *m = CountReply{} }
func (m *CountReply) String() string { return proto.CompactTextString(m) }
func (*CountReply) ProtoMessage()    {}
func (*CountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *CountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountReply.Unmarshal(m, b)
}
func (m *CountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountReply.Marshal(b, m, deterministic)
}
func (m *CountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountReply.Merge(m, src)
}
func (m *CountReply) XXX_Size() int {
	return xxx_messageInfo_CountReply.Size(m)
}
func (m *CountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CountReply.DiscardUnknown(m)
}

var xxx_messageInfo_CountReply proto.InternalMessageInfo

func (m *CountReply) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ListPackingItnbrRequest struct {
	Subno                string   `protobuf:"bytes,1,opt,name=subno,proto3" json:"subno,omitempty"`
	PackingNo            string   `protobuf:"bytes,2,opt,name=packing_no,json=packingNo,proto3" json:"packing_no,omitempty"`
	BoxNumber            int32    `protobuf:"varint,3,opt,name=box_number,json=boxNumber,proto3" json:"box_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPackingItnbrRequest) Reset()         { *m = ListPackingItnbrRequest{} }
func (m *ListPackingItnbrRequest) String() string { return proto.CompactTextString(m) }
func (*ListPackingItnbrRequest) ProtoMessage()    {}
func (*ListPackingItnbrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *ListPackingItnbrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPackingItnbrRequest.Unmarshal(m, b)
}
func (m *ListPackingItnbrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPackingItnbrRequest.Marshal(b, m, deterministic)
}
func (m *ListPackingItnbrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPackingItnbrRequest.Merge(m, src)
}
func (m *ListPackingItnbrRequest) XXX_Size() int {
	return xxx_messageInfo_ListPackingItnbrRequest.Size(m)
}
func (m *ListPackingItnbrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPackingItnbrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPackingItnbrRequest proto.InternalMessageInfo

func (m *ListPackingItnbrRequest) GetSubno() string {
	if m != nil {
		return m.Subno
	}
	return ""
}

func (m *ListPackingItnbrRequest) GetPackingNo() string {
	if m != nil {
		return m.PackingNo
	}
	return ""
}

func (m *ListPackingItnbrRequest) GetBoxNumber() int32 {
	if m != nil {
		return m.BoxNumber
	}
	return 0
}

type ListPackingItnbrReply struct {
	Itnbrs               []*ListPackingItnbrReply_Itnbr `protobuf:"bytes,1,rep,name=itnbrs,proto3" json:"itnbrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ListPackingItnbrReply) Reset()         { *m = ListPackingItnbrReply{} }
func (m *ListPackingItnbrReply) String() string { return proto.CompactTextString(m) }
func (*ListPackingItnbrReply) ProtoMessage()    {}
func (*ListPackingItnbrReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *ListPackingItnbrReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPackingItnbrReply.Unmarshal(m, b)
}
func (m *ListPackingItnbrReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPackingItnbrReply.Marshal(b, m, deterministic)
}
func (m *ListPackingItnbrReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPackingItnbrReply.Merge(m, src)
}
func (m *ListPackingItnbrReply) XXX_Size() int {
	return xxx_messageInfo_ListPackingItnbrReply.Size(m)
}
func (m *ListPackingItnbrReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPackingItnbrReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListPackingItnbrReply proto.InternalMessageInfo

func (m *ListPackingItnbrReply) GetItnbrs() []*ListPackingItnbrReply_Itnbr {
	if m != nil {
		return m.Itnbrs
	}
	return nil
}

type ListPackingItnbrReply_Itnbr struct {
	Itnbr                string   `protobuf:"bytes,1,opt,name=itnbr,proto3" json:"itnbr,omitempty"`
	Quantity             int64    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPackingItnbrReply_Itnbr) Reset()         { *m = ListPackingItnbrReply_Itnbr{} }
func (m *ListPackingItnbrReply_Itnbr) String() string { return proto.CompactTextString(m) }
func (*ListPackingItnbrReply_Itnbr) ProtoMessage()    {}
func (*ListPackingItnbrReply_Itnbr) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7, 0}
}

func (m *ListPackingItnbrReply_Itnbr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPackingItnbrReply_Itnbr.Unmarshal(m, b)
}
func (m *ListPackingItnbrReply_Itnbr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPackingItnbrReply_Itnbr.Marshal(b, m, deterministic)
}
func (m *ListPackingItnbrReply_Itnbr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPackingItnbrReply_Itnbr.Merge(m, src)
}
func (m *ListPackingItnbrReply_Itnbr) XXX_Size() int {
	return xxx_messageInfo_ListPackingItnbrReply_Itnbr.Size(m)
}
func (m *ListPackingItnbrReply_Itnbr) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPackingItnbrReply_Itnbr.DiscardUnknown(m)
}

var xxx_messageInfo_ListPackingItnbrReply_Itnbr proto.InternalMessageInfo

func (m *ListPackingItnbrReply_Itnbr) GetItnbr() string {
	if m != nil {
		return m.Itnbr
	}
	return ""
}

func (m *ListPackingItnbrReply_Itnbr) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func init() {
	proto.RegisterType((*Product)(nil), "proto.dm.cdr.Product")
	proto.RegisterType((*ScrappingTireFee)(nil), "proto.dm.cdr.ScrappingTireFee")
	proto.RegisterType((*Cusno)(nil), "proto.dm.cdr.Cusno")
	proto.RegisterType((*CustomReport)(nil), "proto.dm.cdr.CustomReport")
	proto.RegisterType((*CustomReports)(nil), "proto.dm.cdr.CustomReports")
	proto.RegisterType((*CountReply)(nil), "proto.dm.cdr.CountReply")
	proto.RegisterType((*ListPackingItnbrRequest)(nil), "proto.dm.cdr.ListPackingItnbrRequest")
	proto.RegisterType((*ListPackingItnbrReply)(nil), "proto.dm.cdr.ListPackingItnbrReply")
	proto.RegisterType((*ListPackingItnbrReply_Itnbr)(nil), "proto.dm.cdr.ListPackingItnbrReply.Itnbr")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x4e, 0x13, 0x4f,
	0x14, 0xce, 0xb6, 0xb4, 0xa5, 0xe7, 0x07, 0xf9, 0x91, 0x41, 0xa0, 0x2e, 0x42, 0x70, 0xc4, 0x04,
	0xbd, 0xd8, 0x8d, 0xa8, 0x89, 0x72, 0x87, 0x15, 0x0d, 0x09, 0x21, 0xb8, 0xf5, 0x4a, 0x2f, 0xc8,
	0xfe, 0x19, 0x96, 0x0d, 0xdd, 0x99, 0x65, 0x66, 0x16, 0xdb, 0x4b, 0x7d, 0x05, 0xe3, 0x2b, 0xf8,
	0x1a, 0x3e, 0x84, 0xaf, 0xe0, 0x83, 0x98, 0x39, 0xb3, 0x85, 0xfe, 0x91, 0xe8, 0x45, 0xd3, 0xf9,
	0xbe, 0xef, 0x9c, 0xef, 0x9c, 0xdd, 0x73, 0x66, 0x61, 0x51, 0x31, 0x79, 0x95, 0xc5, 0xcc, 0x2b,
	0xa4, 0xd0, 0x82, 0x2c, 0xe0, 0x9f, 0x97, 0xe4, 0x5e, 0x9c, 0x48, 0xf7, 0x5e, 0x2a, 0x44, 0xda,
	0x67, 0x7e, 0x58, 0x64, 0x7e, 0xc8, 0xb9, 0xd0, 0xa1, 0xce, 0x04, 0x57, 0x36, 0xd6, 0x5d, 0xaf,
	0x54, 0x44, 0x51, 0x79, 0xe6, 0xb3, 0xbc, 0xd0, 0x43, 0x2b, 0xd2, 0xe7, 0xd0, 0x3a, 0x91, 0x22,
	0x29, 0x63, 0x4d, 0xee, 0x40, 0x43, 0x95, 0x11, 0x17, 0x1d, 0x67, 0xcb, 0xd9, 0x69, 0x07, 0x16,
	0x18, 0x36, 0xd3, 0x3c, 0x92, 0x9d, 0x9a, 0x65, 0x11, 0xd0, 0x6d, 0x58, 0xea, 0xc5, 0x32, 0x2c,
	0x8a, 0x8c, 0xa7, 0xef, 0x33, 0xc9, 0xde, 0x30, 0x46, 0x96, 0xa0, 0x7e, 0xc6, 0x18, 0x66, 0xd7,
	0x02, 0x73, 0xa4, 0x1b, 0xd0, 0xe8, 0x96, 0xca, 0x9a, 0xc4, 0xe6, 0x30, 0xb2, 0x46, 0x40, 0xbf,
	0xd7, 0x60, 0xa1, 0x5b, 0x2a, 0x2d, 0xf2, 0x80, 0x15, 0x42, 0x6a, 0xb2, 0x09, 0x60, 0xf1, 0x71,
	0x98, 0xb3, 0x2a, 0x76, 0x8c, 0x21, 0xab, 0xd0, 0xec, 0xe9, 0xe4, 0x52, 0x0f, 0xb1, 0x99, 0x7a,
	0x50, 0x21, 0xd2, 0x81, 0x56, 0x4f, 0x27, 0x91, 0x11, 0xea, 0x28, 0x8c, 0xa0, 0xc9, 0xe8, 0x26,
	0xd2, 0x08, 0x73, 0x36, 0xc3, 0x22, 0x74, 0x3a, 0x2f, 0x0c, 0xdf, 0xa8, 0x9c, 0x10, 0x19, 0x7e,
	0x3f, 0x17, 0x25, 0xd7, 0x9d, 0xa6, 0xe5, 0x2d, 0x22, 0x2e, 0xcc, 0x1f, 0xe9, 0x4a, 0x69, 0xa1,
	0x72, 0x8d, 0x09, 0x81, 0xb9, 0x77, 0x7a, 0xf8, 0xa4, 0x33, 0x8f, 0x3c, 0x9e, 0x09, 0x85, 0x85,
	0x80, 0xa9, 0x42, 0x70, 0xc5, 0x82, 0x50, 0xb3, 0x4e, 0x1b, 0x5f, 0xca, 0x04, 0x47, 0x76, 0xe0,
	0xff, 0x5e, 0x36, 0xd0, 0x8c, 0xf1, 0x43, 0x1e, 0x9f, 0x63, 0x18, 0x60, 0xd8, 0x34, 0x4d, 0x0f,
	0x60, 0x71, 0xfc, 0x3d, 0x29, 0xf2, 0x0c, 0x5a, 0x49, 0xa8, 0x43, 0xc5, 0x74, 0xc7, 0xd9, 0xaa,
	0xef, 0xfc, 0xb7, 0xeb, 0x7a, 0xe3, 0x0b, 0xe1, 0x8d, 0x47, 0x07, 0xa3, 0x50, 0x4a, 0x01, 0xba,
	0xa6, 0xe3, 0x80, 0x15, 0xfd, 0x21, 0xce, 0x04, 0x9f, 0xc7, 0xc1, 0xbe, 0x2d, 0xa0, 0x39, 0xac,
	0x1d, 0x65, 0x4a, 0x9f, 0x84, 0xf1, 0x45, 0xc6, 0xd3, 0x43, 0x33, 0xec, 0x80, 0x5d, 0x96, 0x4c,
	0xdd, 0xb6, 0x1f, 0x1b, 0x00, 0x85, 0x0d, 0x3e, 0xe5, 0xa2, 0x5a, 0x92, 0x76, 0xc5, 0x1c, 0xa3,
	0x1c, 0x89, 0xc1, 0x29, 0x2f, 0xf3, 0x88, 0x49, 0x9c, 0x4e, 0x23, 0x68, 0x47, 0x62, 0x70, 0x8c,
	0x04, 0xfd, 0xe6, 0xc0, 0xca, 0x6c, 0x3d, 0xd3, 0xde, 0x3e, 0x34, 0x71, 0xd5, 0x54, 0xf5, 0x84,
	0x8f, 0x26, 0x9f, 0xf0, 0x8f, 0x49, 0x9e, 0x3d, 0x56, 0x89, 0xee, 0x4b, 0x68, 0x20, 0x71, 0xb3,
	0xc3, 0xce, 0xd8, 0x0e, 0x9b, 0x99, 0x5e, 0x96, 0x21, 0xd7, 0xd9, 0xf5, 0x3e, 0x5d, 0xe3, 0xdd,
	0x1f, 0x75, 0xa8, 0x77, 0x5f, 0x07, 0x24, 0x86, 0xc5, 0x6e, 0x22, 0x8b, 0x24, 0x3e, 0x18, 0x84,
	0x79, 0xd1, 0x67, 0x64, 0xd5, 0xb3, 0xb7, 0xc9, 0x1b, 0xdd, 0x26, 0xef, 0xc0, 0xdc, 0x26, 0x77,
	0xfd, 0xf6, 0x01, 0x28, 0xba, 0xf1, 0xe5, 0xe7, 0xaf, 0xaf, 0xb5, 0x35, 0x4a, 0xfc, 0x38, 0x91,
	0xe6, 0x77, 0x63, 0xb8, 0xe7, 0x3c, 0x26, 0x1f, 0xb1, 0x48, 0x2c, 0xd5, 0xa8, 0xc8, 0xf2, 0x8c,
	0x19, 0x17, 0xee, 0x2d, 0x95, 0x67, 0xcd, 0x6f, 0x8c, 0x8c, 0xb9, 0x84, 0xe5, 0xb7, 0x4c, 0xcf,
	0x5c, 0xd6, 0x95, 0xc9, 0x12, 0xd5, 0x37, 0xc0, 0xdd, 0x9c, 0xa4, 0xa7, 0xd3, 0xe8, 0x36, 0x16,
	0xdb, 0xa4, 0x77, 0xb1, 0x98, 0x9a, 0x92, 0xfd, 0x94, 0x69, 0x53, 0xf3, 0xb3, 0x03, 0x4b, 0xd3,
	0x03, 0x22, 0x0f, 0xff, 0x36, 0x40, 0xdc, 0x32, 0xf7, 0xc1, 0x3f, 0xcc, 0x99, 0xde, 0xc7, 0x36,
	0xd6, 0xe9, 0x2a, 0xb6, 0x51, 0x8c, 0xe9, 0x7e, 0x3f, 0x53, 0xa6, 0x87, 0x57, 0x7b, 0x1f, 0x5e,
	0xa4, 0x99, 0xee, 0x87, 0x91, 0x77, 0xc1, 0x78, 0x12, 0x7a, 0xb1, 0xc8, 0x3d, 0xfd, 0xc9, 0x47,
	0xe0, 0xc7, 0x22, 0xcf, 0x05, 0x57, 0xfe, 0xd5, 0xae, 0xfd, 0x2a, 0xfa, 0xa9, 0xe8, 0x87, 0x3c,
	0xf5, 0x93, 0xdc, 0x18, 0x46, 0x4d, 0x24, 0x9f, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x21, 0x23,
	0x8d, 0xa2, 0x75, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CDRClient is the client API for CDR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CDRClient interface {
	// cdrpdc.4gl 範例
	CdrpdcExample(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CustomReports, error)
	// cdrcrs.4gl 範例
	CdrcrsExample(ctx context.Context, in *Cusno, opts ...grpc.CallOption) (*empty.Empty, error)
	//廢胎費用
	GetScrappingTireFee(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ScrappingTireFee, error)
	// 查詢包裝成代
	//
	// 依據 packing 編號和箱號起訖, 查詢對應成代與數量.
	ListPackingItnbr(ctx context.Context, in *ListPackingItnbrRequest, opts ...grpc.CallOption) (*ListPackingItnbrReply, error)
}

type cDRClient struct {
	cc *grpc.ClientConn
}

func NewCDRClient(cc *grpc.ClientConn) CDRClient {
	return &cDRClient{cc}
}

func (c *cDRClient) CdrpdcExample(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CustomReports, error) {
	out := new(CustomReports)
	err := c.cc.Invoke(ctx, "/proto.dm.cdr.CDR/CdrpdcExample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDRClient) CdrcrsExample(ctx context.Context, in *Cusno, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.dm.cdr.CDR/CdrcrsExample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDRClient) GetScrappingTireFee(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ScrappingTireFee, error) {
	out := new(ScrappingTireFee)
	err := c.cc.Invoke(ctx, "/proto.dm.cdr.CDR/GetScrappingTireFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDRClient) ListPackingItnbr(ctx context.Context, in *ListPackingItnbrRequest, opts ...grpc.CallOption) (*ListPackingItnbrReply, error) {
	out := new(ListPackingItnbrReply)
	err := c.cc.Invoke(ctx, "/proto.dm.cdr.CDR/ListPackingItnbr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CDRServer is the server API for CDR service.
type CDRServer interface {
	// cdrpdc.4gl 範例
	CdrpdcExample(context.Context, *empty.Empty) (*CustomReports, error)
	// cdrcrs.4gl 範例
	CdrcrsExample(context.Context, *Cusno) (*empty.Empty, error)
	//廢胎費用
	GetScrappingTireFee(context.Context, *Product) (*ScrappingTireFee, error)
	// 查詢包裝成代
	//
	// 依據 packing 編號和箱號起訖, 查詢對應成代與數量.
	ListPackingItnbr(context.Context, *ListPackingItnbrRequest) (*ListPackingItnbrReply, error)
}

func RegisterCDRServer(s *grpc.Server, srv CDRServer) {
	s.RegisterService(&_CDR_serviceDesc, srv)
}

func _CDR_CdrpdcExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDRServer).CdrpdcExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.cdr.CDR/CdrpdcExample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDRServer).CdrpdcExample(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDR_CdrcrsExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cusno)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDRServer).CdrcrsExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.cdr.CDR/CdrcrsExample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDRServer).CdrcrsExample(ctx, req.(*Cusno))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDR_GetScrappingTireFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDRServer).GetScrappingTireFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.cdr.CDR/GetScrappingTireFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDRServer).GetScrappingTireFee(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDR_ListPackingItnbr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPackingItnbrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDRServer).ListPackingItnbr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dm.cdr.CDR/ListPackingItnbr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDRServer).ListPackingItnbr(ctx, req.(*ListPackingItnbrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CDR_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.dm.cdr.CDR",
	HandlerType: (*CDRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CdrpdcExample",
			Handler:    _CDR_CdrpdcExample_Handler,
		},
		{
			MethodName: "CdrcrsExample",
			Handler:    _CDR_CdrcrsExample_Handler,
		},
		{
			MethodName: "GetScrappingTireFee",
			Handler:    _CDR_GetScrappingTireFee_Handler,
		},
		{
			MethodName: "ListPackingItnbr",
			Handler:    _CDR_ListPackingItnbr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}