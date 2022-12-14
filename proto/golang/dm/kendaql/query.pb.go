// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package kendaql

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

type JoinType int32

const (
	JoinType_JOIN        JoinType = 0
	JoinType_LEFT        JoinType = 1
	JoinType_RIGHT       JoinType = 2
	JoinType_INNER       JoinType = 3
	JoinType_LEFT_OUTER  JoinType = 4
	JoinType_RIGHT_OUTER JoinType = 5
	JoinType_FULL_OUTER  JoinType = 6
	JoinType_CROSS       JoinType = 7
)

var JoinType_name = map[int32]string{
	0: "JOIN",
	1: "LEFT",
	2: "RIGHT",
	3: "INNER",
	4: "LEFT_OUTER",
	5: "RIGHT_OUTER",
	6: "FULL_OUTER",
	7: "CROSS",
}

var JoinType_value = map[string]int32{
	"JOIN":        0,
	"LEFT":        1,
	"RIGHT":       2,
	"INNER":       3,
	"LEFT_OUTER":  4,
	"RIGHT_OUTER": 5,
	"FULL_OUTER":  6,
	"CROSS":       7,
}

func (x JoinType) String() string {
	return proto.EnumName(JoinType_name, int32(x))
}

func (JoinType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}

type Order int32

const (
	Order_ASCENDING  Order = 0
	Order_DESCENDING Order = 1
)

var Order_name = map[int32]string{
	0: "ASCENDING",
	1: "DESCENDING",
}

var Order_value = map[string]int32{
	"ASCENDING":  0,
	"DESCENDING": 1,
}

func (x Order) String() string {
	return proto.EnumName(Order_name, int32(x))
}

func (Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}

type QueryRequest struct {
	Db                   *Database `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	Sql                  *Query    `protobuf:"bytes,2,opt,name=sql,proto3" json:"sql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetDb() *Database {
	if m != nil {
		return m.Db
	}
	return nil
}

func (m *QueryRequest) GetSql() *Query {
	if m != nil {
		return m.Sql
	}
	return nil
}

type Query struct {
	Selects              []*Col        `protobuf:"bytes,1,rep,name=selects,proto3" json:"selects,omitempty"`
	From                 *Table        `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Joins                []*Join       `protobuf:"bytes,3,rep,name=joins,proto3" json:"joins,omitempty"`
	Where                *LogicalGroup `protobuf:"bytes,4,opt,name=where,proto3" json:"where,omitempty"`
	GroupBys             []*Column     `protobuf:"bytes,5,rep,name=group_bys,json=groupBys,proto3" json:"group_bys,omitempty"`
	Having               *LogicalGroup `protobuf:"bytes,6,opt,name=having,proto3" json:"having,omitempty"`
	OrderBys             []*OrderBy    `protobuf:"bytes,7,rep,name=order_bys,json=orderBys,proto3" json:"order_bys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetSelects() []*Col {
	if m != nil {
		return m.Selects
	}
	return nil
}

func (m *Query) GetFrom() *Table {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Query) GetJoins() []*Join {
	if m != nil {
		return m.Joins
	}
	return nil
}

func (m *Query) GetWhere() *LogicalGroup {
	if m != nil {
		return m.Where
	}
	return nil
}

func (m *Query) GetGroupBys() []*Column {
	if m != nil {
		return m.GroupBys
	}
	return nil
}

func (m *Query) GetHaving() *LogicalGroup {
	if m != nil {
		return m.Having
	}
	return nil
}

func (m *Query) GetOrderBys() []*OrderBy {
	if m != nil {
		return m.OrderBys
	}
	return nil
}

type Join struct {
	Type                 JoinType  `protobuf:"varint,1,opt,name=type,proto3,enum=proto.dm.kendaql.JoinType" json:"type,omitempty"`
	Table                *Table    `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Ons                  []*JoinOn `protobuf:"bytes,3,rep,name=ons,proto3" json:"ons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Join) Reset()         { *m = Join{} }
func (m *Join) String() string { return proto.CompactTextString(m) }
func (*Join) ProtoMessage()    {}
func (*Join) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}

func (m *Join) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Join.Unmarshal(m, b)
}
func (m *Join) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Join.Marshal(b, m, deterministic)
}
func (m *Join) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Join.Merge(m, src)
}
func (m *Join) XXX_Size() int {
	return xxx_messageInfo_Join.Size(m)
}
func (m *Join) XXX_DiscardUnknown() {
	xxx_messageInfo_Join.DiscardUnknown(m)
}

var xxx_messageInfo_Join proto.InternalMessageInfo

func (m *Join) GetType() JoinType {
	if m != nil {
		return m.Type
	}
	return JoinType_JOIN
}

func (m *Join) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *Join) GetOns() []*JoinOn {
	if m != nil {
		return m.Ons
	}
	return nil
}

type JoinOn struct {
	From                 *Column  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Column  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinOn) Reset()         { *m = JoinOn{} }
func (m *JoinOn) String() string { return proto.CompactTextString(m) }
func (*JoinOn) ProtoMessage()    {}
func (*JoinOn) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{3}
}

func (m *JoinOn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinOn.Unmarshal(m, b)
}
func (m *JoinOn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinOn.Marshal(b, m, deterministic)
}
func (m *JoinOn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinOn.Merge(m, src)
}
func (m *JoinOn) XXX_Size() int {
	return xxx_messageInfo_JoinOn.Size(m)
}
func (m *JoinOn) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinOn.DiscardUnknown(m)
}

var xxx_messageInfo_JoinOn proto.InternalMessageInfo

func (m *JoinOn) GetFrom() *Column {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *JoinOn) GetTo() *Column {
	if m != nil {
		return m.To
	}
	return nil
}

type OrderBy struct {
	Type                 Order    `protobuf:"varint,1,opt,name=type,proto3,enum=proto.dm.kendaql.Order" json:"type,omitempty"`
	Column               *Column  `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderBy) Reset()         { *m = OrderBy{} }
func (m *OrderBy) String() string { return proto.CompactTextString(m) }
func (*OrderBy) ProtoMessage()    {}
func (*OrderBy) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{4}
}

func (m *OrderBy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderBy.Unmarshal(m, b)
}
func (m *OrderBy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderBy.Marshal(b, m, deterministic)
}
func (m *OrderBy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBy.Merge(m, src)
}
func (m *OrderBy) XXX_Size() int {
	return xxx_messageInfo_OrderBy.Size(m)
}
func (m *OrderBy) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBy.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBy proto.InternalMessageInfo

func (m *OrderBy) GetType() Order {
	if m != nil {
		return m.Type
	}
	return Order_ASCENDING
}

func (m *OrderBy) GetColumn() *Column {
	if m != nil {
		return m.Column
	}
	return nil
}

type QueryReply struct {
	Headers              []*ColumnProperty `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	Rows                 []*Row            `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}
func (*QueryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{5}
}

func (m *QueryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReply.Unmarshal(m, b)
}
func (m *QueryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReply.Marshal(b, m, deterministic)
}
func (m *QueryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReply.Merge(m, src)
}
func (m *QueryReply) XXX_Size() int {
	return xxx_messageInfo_QueryReply.Size(m)
}
func (m *QueryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReply proto.InternalMessageInfo

func (m *QueryReply) GetHeaders() []*ColumnProperty {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *QueryReply) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type ColumnProperty struct {
	DataType             DataType `protobuf:"varint,1,opt,name=data_type,json=dataType,proto3,enum=proto.dm.kendaql.DataType" json:"data_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColumnProperty) Reset()         { *m = ColumnProperty{} }
func (m *ColumnProperty) String() string { return proto.CompactTextString(m) }
func (*ColumnProperty) ProtoMessage()    {}
func (*ColumnProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{6}
}

func (m *ColumnProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnProperty.Unmarshal(m, b)
}
func (m *ColumnProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnProperty.Marshal(b, m, deterministic)
}
func (m *ColumnProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnProperty.Merge(m, src)
}
func (m *ColumnProperty) XXX_Size() int {
	return xxx_messageInfo_ColumnProperty.Size(m)
}
func (m *ColumnProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnProperty.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnProperty proto.InternalMessageInfo

func (m *ColumnProperty) GetDataType() DataType {
	if m != nil {
		return m.DataType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

type Row struct {
	Values               [][]byte `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{7}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.dm.kendaql.JoinType", JoinType_name, JoinType_value)
	proto.RegisterEnum("proto.dm.kendaql.Order", Order_name, Order_value)
	proto.RegisterType((*QueryRequest)(nil), "proto.dm.kendaql.QueryRequest")
	proto.RegisterType((*Query)(nil), "proto.dm.kendaql.Query")
	proto.RegisterType((*Join)(nil), "proto.dm.kendaql.Join")
	proto.RegisterType((*JoinOn)(nil), "proto.dm.kendaql.JoinOn")
	proto.RegisterType((*OrderBy)(nil), "proto.dm.kendaql.OrderBy")
	proto.RegisterType((*QueryReply)(nil), "proto.dm.kendaql.QueryReply")
	proto.RegisterType((*ColumnProperty)(nil), "proto.dm.kendaql.ColumnProperty")
	proto.RegisterType((*Row)(nil), "proto.dm.kendaql.Row")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0xc7, 0xd7, 0xb4, 0x49, 0xdb, 0xd3, 0x3d, 0x7b, 0x22, 0x4b, 0x6c, 0xd9, 0x04, 0xa8, 0xca,
	0x05, 0xda, 0xca, 0x48, 0xd0, 0x80, 0x21, 0x21, 0x84, 0xc4, 0xb6, 0x6e, 0x74, 0xaa, 0x5a, 0xf0,
	0xba, 0x1b, 0x6e, 0x86, 0xd3, 0x78, 0x6d, 0x21, 0x89, 0xd3, 0x24, 0x5d, 0x95, 0x8f, 0xc1, 0x37,
	0xe5, 0x23, 0x20, 0xbf, 0x74, 0x1a, 0x64, 0xd1, 0xb8, 0x8a, 0x4f, 0xfc, 0x3b, 0xff, 0xe3, 0xf3,
	0x62, 0x43, 0x6b, 0xbe, 0xa0, 0x49, 0xee, 0xc4, 0x09, 0xcb, 0x18, 0x32, 0xc5, 0xc7, 0xf1, 0x43,
	0xe7, 0x07, 0x8d, 0x7c, 0x32, 0x0f, 0x76, 0xda, 0xe2, 0x8f, 0xeb, 0x87, 0xae, 0xfa, 0xe3, 0xfa,
	0x24, 0x23, 0x57, 0x59, 0x1e, 0x53, 0xe9, 0xb3, 0xf3, 0xb8, 0x40, 0x5c, 0xcf, 0x68, 0xe0, 0xab,
	0xdd, 0xed, 0xa2, 0xbf, 0x57, 0xea, 0x98, 0x11, 0x2f, 0x50, 0xb2, 0x36, 0x85, 0xf5, 0x2f, 0xfc,
	0x64, 0x98, 0xce, 0x17, 0x34, 0xcd, 0x50, 0x07, 0x34, 0xdf, 0xb3, 0x2a, 0xed, 0xca, 0x6e, 0xeb,
	0x60, 0xc7, 0xf9, 0xfb, 0x9c, 0xce, 0x09, 0xc9, 0x88, 0x47, 0x52, 0x8a, 0x35, 0xdf, 0x43, 0x7b,
	0x50, 0x4d, 0xe7, 0x81, 0xa5, 0x09, 0x78, 0xab, 0x08, 0x4b, 0x61, 0xce, 0xd8, 0xbf, 0x34, 0xd0,
	0x85, 0x89, 0x5c, 0xa8, 0xa7, 0x34, 0xa0, 0xe3, 0x2c, 0xb5, 0x2a, 0xed, 0xea, 0x6e, 0xeb, 0xe0,
	0x51, 0xd1, 0xf1, 0x98, 0x05, 0x78, 0x45, 0xa1, 0xe7, 0x50, 0xbb, 0x4e, 0x58, 0x58, 0x1e, 0x66,
	0xc4, 0xd3, 0xc1, 0x02, 0x42, 0xfb, 0xa0, 0x7f, 0x67, 0xb3, 0x28, 0xb5, 0xaa, 0x42, 0x7b, 0xb3,
	0x48, 0x9f, 0xb3, 0x59, 0x84, 0x25, 0x84, 0x5e, 0x83, 0xbe, 0x9c, 0xd2, 0x84, 0x5a, 0x35, 0xa1,
	0xfd, 0xb4, 0x48, 0xf7, 0xd9, 0x64, 0x36, 0x26, 0xc1, 0x59, 0xc2, 0x16, 0x31, 0x96, 0x30, 0x7a,
	0x03, 0xcd, 0x09, 0xb7, 0xaf, 0xbc, 0x3c, 0xb5, 0x74, 0x11, 0xc7, 0xba, 0x37, 0x87, 0x45, 0x18,
	0xe1, 0x86, 0x40, 0x8f, 0xf2, 0x14, 0x1d, 0x82, 0x31, 0x25, 0x37, 0xb3, 0x68, 0x62, 0x19, 0xff,
	0x14, 0x4d, 0xd1, 0xe8, 0x10, 0x9a, 0x2c, 0xf1, 0x69, 0x22, 0xc2, 0xd5, 0x45, 0xb8, 0xed, 0xa2,
	0xeb, 0x90, 0x23, 0x47, 0x39, 0x6e, 0x30, 0xb9, 0x48, 0xed, 0x9f, 0x15, 0xa8, 0xf1, 0x64, 0x91,
	0x03, 0x35, 0x3e, 0x47, 0xa2, 0xa9, 0x1b, 0xf7, 0x35, 0x95, 0x53, 0xa3, 0x3c, 0xa6, 0x58, 0x70,
	0xe8, 0x05, 0xe8, 0x62, 0x42, 0x1e, 0xaa, 0xb8, 0xa4, 0x50, 0x07, 0xaa, 0xec, 0xb6, 0xe0, 0xd6,
	0xfd, 0xea, 0xc3, 0x08, 0x73, 0xc8, 0xfe, 0x06, 0x86, 0x34, 0xd1, 0xbe, 0xea, 0xaa, 0x9c, 0xb4,
	0xf2, 0xfa, 0xc9, 0xb6, 0xee, 0x82, 0x96, 0x31, 0x75, 0x9e, 0x72, 0x56, 0xcb, 0x98, 0x3d, 0x85,
	0xba, 0x2a, 0x05, 0x1f, 0x9c, 0x3b, 0x79, 0x6f, 0x95, 0xd4, 0x4c, 0x25, 0xfd, 0x12, 0x8c, 0xb1,
	0x50, 0x79, 0x30, 0x8a, 0xe2, 0xec, 0x14, 0x40, 0xdd, 0x9c, 0x38, 0xc8, 0xd1, 0x3b, 0xa8, 0x4f,
	0x29, 0xf1, 0x69, 0xb2, 0x1a, 0xeb, 0x76, 0x99, 0xc0, 0xe7, 0x84, 0xc5, 0x34, 0xc9, 0x72, 0xbc,
	0x72, 0x40, 0x7b, 0x50, 0x4b, 0xd8, 0x32, 0xb5, 0xb4, 0xb2, 0xfb, 0x80, 0xd9, 0x12, 0x0b, 0xc4,
	0xee, 0xc1, 0xc6, 0x9f, 0x2a, 0xe8, 0x2d, 0x34, 0x6f, 0x9f, 0x8a, 0xf2, 0x16, 0xf3, 0x7b, 0x2b,
	0x5a, 0xdc, 0xf0, 0xd5, 0xca, 0x7e, 0x02, 0x55, 0xcc, 0x96, 0x68, 0x13, 0x8c, 0x1b, 0x12, 0x2c,
	0xa8, 0x3c, 0xf7, 0x3a, 0x56, 0x56, 0x27, 0x86, 0xc6, 0x6a, 0x2e, 0x50, 0x03, 0x6a, 0xe7, 0xc3,
	0xde, 0xc0, 0x5c, 0xe3, 0xab, 0x7e, 0xf7, 0x74, 0x64, 0x56, 0x50, 0x13, 0x74, 0xdc, 0x3b, 0xfb,
	0x34, 0x32, 0x35, 0xbe, 0xec, 0x0d, 0x06, 0x5d, 0x6c, 0x56, 0xd1, 0x06, 0x00, 0xdf, 0xbf, 0x1a,
	0x5e, 0x8e, 0xba, 0xd8, 0xac, 0xa1, 0xff, 0xa1, 0x25, 0x28, 0xf5, 0x43, 0xe7, 0xc0, 0xe9, 0x65,
	0xbf, 0xaf, 0x6c, 0x83, 0xfb, 0x1e, 0xe3, 0xe1, 0xc5, 0x85, 0x59, 0xef, 0x3c, 0x03, 0x5d, 0x74,
	0x04, 0xfd, 0x07, 0xcd, 0x8f, 0x17, 0xc7, 0xdd, 0xc1, 0x49, 0x6f, 0x70, 0x66, 0xae, 0x71, 0x97,
	0x93, 0xee, 0xad, 0x5d, 0x39, 0xfa, 0xf0, 0xf5, 0xfd, 0x64, 0x96, 0x05, 0xc4, 0x93, 0xd9, 0x39,
	0x63, 0x16, 0x3a, 0xd9, 0x52, 0xbe, 0x6e, 0xee, 0x98, 0x85, 0x21, 0x8b, 0x52, 0xf7, 0xe6, 0xc0,
	0x95, 0xef, 0xde, 0x84, 0x05, 0x24, 0x9a, 0xdc, 0x79, 0xfe, 0x3c, 0x43, 0x6c, 0xbc, 0xfa, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0x75, 0x59, 0x27, 0x93, 0x05, 0x00, 0x00,
}
