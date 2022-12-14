// Code generated by protoc-gen-go. DO NOT EDIT.
// source: table.proto

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

type Table struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alias                string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{0}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Table) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

type Column struct {
	TableAlias           string   `protobuf:"bytes,1,opt,name=table_alias,json=tableAlias,proto3" json:"table_alias,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{1}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetTableAlias() string {
	if m != nil {
		return m.TableAlias
	}
	return ""
}

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Function struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Parameters           []*Function_Param `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	ExpectedType         DataType          `protobuf:"varint,3,opt,name=expected_type,json=expectedType,proto3,enum=proto.dm.kendaql.DataType" json:"expected_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{2}
}

func (m *Function) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function.Unmarshal(m, b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function.Marshal(b, m, deterministic)
}
func (m *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(m, src)
}
func (m *Function) XXX_Size() int {
	return xxx_messageInfo_Function.Size(m)
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetParameters() []*Function_Param {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Function) GetExpectedType() DataType {
	if m != nil {
		return m.ExpectedType
	}
	return DataType_DATA_TYPE_UNSPECIFIED
}

type Function_Param struct {
	// Types that are valid to be assigned to Value:
	//	*Function_Param_StringValue
	//	*Function_Param_ColumnValue
	Value                isFunction_Param_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Function_Param) Reset()         { *m = Function_Param{} }
func (m *Function_Param) String() string { return proto.CompactTextString(m) }
func (*Function_Param) ProtoMessage()    {}
func (*Function_Param) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{2, 0}
}

func (m *Function_Param) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function_Param.Unmarshal(m, b)
}
func (m *Function_Param) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function_Param.Marshal(b, m, deterministic)
}
func (m *Function_Param) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function_Param.Merge(m, src)
}
func (m *Function_Param) XXX_Size() int {
	return xxx_messageInfo_Function_Param.Size(m)
}
func (m *Function_Param) XXX_DiscardUnknown() {
	xxx_messageInfo_Function_Param.DiscardUnknown(m)
}

var xxx_messageInfo_Function_Param proto.InternalMessageInfo

type isFunction_Param_Value interface {
	isFunction_Param_Value()
}

type Function_Param_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Function_Param_ColumnValue struct {
	ColumnValue *Column `protobuf:"bytes,2,opt,name=column_value,json=columnValue,proto3,oneof"`
}

func (*Function_Param_StringValue) isFunction_Param_Value() {}

func (*Function_Param_ColumnValue) isFunction_Param_Value() {}

func (m *Function_Param) GetValue() isFunction_Param_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Function_Param) GetStringValue() string {
	if x, ok := m.GetValue().(*Function_Param_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Function_Param) GetColumnValue() *Column {
	if x, ok := m.GetValue().(*Function_Param_ColumnValue); ok {
		return x.ColumnValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Function_Param) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Function_Param_StringValue)(nil),
		(*Function_Param_ColumnValue)(nil),
	}
}

type Col struct {
	// Types that are valid to be assigned to Col:
	//	*Col_Column
	//	*Col_Function
	Col                  isCol_Col `protobuf_oneof:"col"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Col) Reset()         { *m = Col{} }
func (m *Col) String() string { return proto.CompactTextString(m) }
func (*Col) ProtoMessage()    {}
func (*Col) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{3}
}

func (m *Col) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Col.Unmarshal(m, b)
}
func (m *Col) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Col.Marshal(b, m, deterministic)
}
func (m *Col) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Col.Merge(m, src)
}
func (m *Col) XXX_Size() int {
	return xxx_messageInfo_Col.Size(m)
}
func (m *Col) XXX_DiscardUnknown() {
	xxx_messageInfo_Col.DiscardUnknown(m)
}

var xxx_messageInfo_Col proto.InternalMessageInfo

type isCol_Col interface {
	isCol_Col()
}

type Col_Column struct {
	Column *Column `protobuf:"bytes,1,opt,name=column,proto3,oneof"`
}

type Col_Function struct {
	Function *Function `protobuf:"bytes,2,opt,name=function,proto3,oneof"`
}

func (*Col_Column) isCol_Col() {}

func (*Col_Function) isCol_Col() {}

func (m *Col) GetCol() isCol_Col {
	if m != nil {
		return m.Col
	}
	return nil
}

func (m *Col) GetColumn() *Column {
	if x, ok := m.GetCol().(*Col_Column); ok {
		return x.Column
	}
	return nil
}

func (m *Col) GetFunction() *Function {
	if x, ok := m.GetCol().(*Col_Function); ok {
		return x.Function
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Col) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Col_Column)(nil),
		(*Col_Function)(nil),
	}
}

func init() {
	proto.RegisterType((*Table)(nil), "proto.dm.kendaql.Table")
	proto.RegisterType((*Column)(nil), "proto.dm.kendaql.Column")
	proto.RegisterType((*Function)(nil), "proto.dm.kendaql.Function")
	proto.RegisterType((*Function_Param)(nil), "proto.dm.kendaql.Function.Param")
	proto.RegisterType((*Col)(nil), "proto.dm.kendaql.Col")
}

func init() { proto.RegisterFile("table.proto", fileDescriptor_448a2743262f7a00) }

var fileDescriptor_448a2743262f7a00 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcb, 0x6e, 0xda, 0x40,
	0x14, 0x05, 0x5c, 0x53, 0x7a, 0x4d, 0xab, 0x6a, 0xd4, 0x85, 0xe5, 0x4d, 0x2d, 0x77, 0xc3, 0x6a,
	0xac, 0xba, 0x9b, 0x2e, 0x4a, 0x93, 0x40, 0x14, 0xb1, 0x8c, 0x2c, 0x94, 0x45, 0x36, 0x68, 0xb0,
	0x27, 0x96, 0x95, 0x79, 0x38, 0x66, 0x20, 0x21, 0x3f, 0x92, 0xdf, 0x8d, 0xe6, 0x01, 0x41, 0x21,
	0xc9, 0xca, 0x9e, 0x33, 0xe7, 0x31, 0xf7, 0xe8, 0x42, 0xa0, 0xc8, 0x92, 0x51, 0xdc, 0xb4, 0x52,
	0x49, 0xf4, 0xdd, 0x7c, 0x70, 0xc9, 0xf1, 0x2d, 0x15, 0x25, 0xb9, 0x63, 0x51, 0x6c, 0x90, 0xb4,
	0xe4, 0xa9, 0x43, 0xd2, 0x92, 0x28, 0xb2, 0x50, 0xdb, 0xc6, 0x69, 0x92, 0xdf, 0xe0, 0xcf, 0xb5,
	0x05, 0x42, 0xf0, 0x49, 0x10, 0x4e, 0xc3, 0x6e, 0xdc, 0x1d, 0x7d, 0xc9, 0xcd, 0x3f, 0xfa, 0x01,
	0x3e, 0x61, 0x35, 0x59, 0x85, 0x3d, 0x03, 0xda, 0x43, 0x32, 0x86, 0xfe, 0x54, 0xb2, 0x35, 0x17,
	0xe8, 0xa7, 0xcb, 0x5f, 0x58, 0x96, 0x95, 0x82, 0x81, 0xce, 0x34, 0xb2, 0x37, 0xed, 0xbd, 0x98,
	0x26, 0x4f, 0x3d, 0x18, 0x5c, 0xac, 0x45, 0xa1, 0x6a, 0x29, 0xde, 0x4c, 0x3d, 0x05, 0x68, 0x48,
	0x4b, 0x38, 0x55, 0xb4, 0xd5, 0xd1, 0xde, 0x28, 0xc8, 0x62, 0xfc, 0x7a, 0x36, 0xbc, 0xf3, 0xc0,
	0x97, 0x9a, 0x9c, 0x1f, 0x68, 0xd0, 0x09, 0x7c, 0xa5, 0x0f, 0x0d, 0x2d, 0x14, 0x2d, 0xcd, 0xac,
	0xa1, 0x17, 0x77, 0x47, 0xdf, 0xb2, 0xe8, 0xd8, 0xe4, 0x9c, 0x28, 0x32, 0xdf, 0x36, 0x34, 0x1f,
	0xee, 0x04, 0xfa, 0x14, 0x29, 0xf0, 0x8d, 0x2b, 0xfa, 0x05, 0xc3, 0x95, 0x6a, 0x6b, 0x51, 0x2d,
	0x36, 0x84, 0xad, 0xdd, 0x3b, 0x67, 0x9d, 0x3c, 0xb0, 0xe8, 0x95, 0x06, 0xd1, 0x18, 0x86, 0x85,
	0x29, 0xc4, 0x91, 0xf4, 0xb4, 0x41, 0x16, 0x1e, 0xa7, 0xd9, 0xda, 0xb4, 0xdc, 0xf2, 0x8d, 0x7c,
	0xf2, 0x19, 0x7c, 0xa3, 0x4b, 0x1e, 0xc1, 0x9b, 0x4a, 0x86, 0x32, 0xe8, 0xdb, 0x6b, 0x93, 0xf6,
	0xb1, 0x91, 0x63, 0xa2, 0xbf, 0x30, 0xb8, 0x71, 0x7d, 0xb8, 0xf8, 0xe8, 0xfd, 0xc6, 0x66, 0x9d,
	0x7c, 0xcf, 0x9e, 0xf8, 0xe0, 0x15, 0x92, 0x4d, 0xfe, 0x5f, 0xff, 0xab, 0x6a, 0xc5, 0xc8, 0xd2,
	0xb2, 0x71, 0x21, 0x39, 0x56, 0xf7, 0x76, 0x6d, 0xd2, 0x42, 0x72, 0x2e, 0xc5, 0x2a, 0xdd, 0x64,
	0xa9, 0x5d, 0xa8, 0x4a, 0x32, 0x22, 0xaa, 0x83, 0xbd, 0x5a, 0xf6, 0xcd, 0xc5, 0x9f, 0xe7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0x27, 0x2a, 0x55, 0x91, 0x02, 0x00, 0x00,
}
