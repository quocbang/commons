// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission.proto

package utility

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

type Permission int32

const (
	Permission_NO_PERMISSION               Permission = 0
	Permission_FLEXIBLE_HOUR               Permission = 1
	Permission_PCR_DEALER_PERFORMANCE_DATA Permission = 2
	Permission_ACCOUNTING_ENTRY            Permission = 3
	Permission_MES                         Permission = 4
	Permission_PROJECT_WINDOW_OPEN         Permission = 5
	Permission_PROJECT_QUERY               Permission = 6
	Permission_PROJECT_FILE_UPLOAD         Permission = 7
	Permission_PROJECT_DETAIL              Permission = 8
)

var Permission_name = map[int32]string{
	0: "NO_PERMISSION",
	1: "FLEXIBLE_HOUR",
	2: "PCR_DEALER_PERFORMANCE_DATA",
	3: "ACCOUNTING_ENTRY",
	4: "MES",
	5: "PROJECT_WINDOW_OPEN",
	6: "PROJECT_QUERY",
	7: "PROJECT_FILE_UPLOAD",
	8: "PROJECT_DETAIL",
}

var Permission_value = map[string]int32{
	"NO_PERMISSION":               0,
	"FLEXIBLE_HOUR":               1,
	"PCR_DEALER_PERFORMANCE_DATA": 2,
	"ACCOUNTING_ENTRY":            3,
	"MES":                         4,
	"PROJECT_WINDOW_OPEN":         5,
	"PROJECT_QUERY":               6,
	"PROJECT_FILE_UPLOAD":         7,
	"PROJECT_DETAIL":              8,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}

type Permissions struct {
	Dataset              []Permission `protobuf:"varint,1,rep,packed,name=dataset,proto3,enum=proto.dm.utility.Permission" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}

func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetDataset() []Permission {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type UserPermission struct {
	UserId               string               `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PermissionId         Permission           `protobuf:"varint,2,opt,name=permission_id,json=permissionId,proto3,enum=proto.dm.utility.Permission" json:"permission_id,omitempty"`
	ValidTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=valid_time,json=validTime,proto3" json:"valid_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserPermission) Reset()         { *m = UserPermission{} }
func (m *UserPermission) String() string { return proto.CompactTextString(m) }
func (*UserPermission) ProtoMessage()    {}
func (*UserPermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{1}
}

func (m *UserPermission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPermission.Unmarshal(m, b)
}
func (m *UserPermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPermission.Marshal(b, m, deterministic)
}
func (m *UserPermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPermission.Merge(m, src)
}
func (m *UserPermission) XXX_Size() int {
	return xxx_messageInfo_UserPermission.Size(m)
}
func (m *UserPermission) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPermission.DiscardUnknown(m)
}

var xxx_messageInfo_UserPermission proto.InternalMessageInfo

func (m *UserPermission) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserPermission) GetPermissionId() Permission {
	if m != nil {
		return m.PermissionId
	}
	return Permission_NO_PERMISSION
}

func (m *UserPermission) GetValidTime() *timestamp.Timestamp {
	if m != nil {
		return m.ValidTime
	}
	return nil
}

type MutationPermissionRequest struct {
	Id                   Permission                          `protobuf:"varint,1,opt,name=id,proto3,enum=proto.dm.utility.Permission" json:"id,omitempty"`
	Mutation             *MutationPermissionRequest_Mutation `protobuf:"bytes,2,opt,name=mutation,proto3" json:"mutation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *MutationPermissionRequest) Reset()         { *m = MutationPermissionRequest{} }
func (m *MutationPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*MutationPermissionRequest) ProtoMessage()    {}
func (*MutationPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{2}
}

func (m *MutationPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutationPermissionRequest.Unmarshal(m, b)
}
func (m *MutationPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutationPermissionRequest.Marshal(b, m, deterministic)
}
func (m *MutationPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutationPermissionRequest.Merge(m, src)
}
func (m *MutationPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_MutationPermissionRequest.Size(m)
}
func (m *MutationPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutationPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutationPermissionRequest proto.InternalMessageInfo

func (m *MutationPermissionRequest) GetId() Permission {
	if m != nil {
		return m.Id
	}
	return Permission_NO_PERMISSION
}

func (m *MutationPermissionRequest) GetMutation() *MutationPermissionRequest_Mutation {
	if m != nil {
		return m.Mutation
	}
	return nil
}

type MutationPermissionRequest_Mutation struct {
	IncludedReadableColumns map[string]*MutationPermissionRequest_Mutation_Column `protobuf:"bytes,1,rep,name=included_readable_columns,json=includedReadableColumns,proto3" json:"included_readable_columns,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExcludedReadableColumns map[string]*MutationPermissionRequest_Mutation_Column `protobuf:"bytes,2,rep,name=excluded_readable_columns,json=excludedReadableColumns,proto3" json:"excluded_readable_columns,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IncludedWritableColumns map[string]*MutationPermissionRequest_Mutation_Column `protobuf:"bytes,3,rep,name=included_writable_columns,json=includedWritableColumns,proto3" json:"included_writable_columns,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExcludedWritableColumns map[string]*MutationPermissionRequest_Mutation_Column `protobuf:"bytes,4,rep,name=excluded_writable_columns,json=excludedWritableColumns,proto3" json:"excluded_writable_columns,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral    struct{}                                              `json:"-"`
	XXX_unrecognized        []byte                                                `json:"-"`
	XXX_sizecache           int32                                                 `json:"-"`
}

func (m *MutationPermissionRequest_Mutation) Reset()         { *m = MutationPermissionRequest_Mutation{} }
func (m *MutationPermissionRequest_Mutation) String() string { return proto.CompactTextString(m) }
func (*MutationPermissionRequest_Mutation) ProtoMessage()    {}
func (*MutationPermissionRequest_Mutation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{2, 0}
}

func (m *MutationPermissionRequest_Mutation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutationPermissionRequest_Mutation.Unmarshal(m, b)
}
func (m *MutationPermissionRequest_Mutation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutationPermissionRequest_Mutation.Marshal(b, m, deterministic)
}
func (m *MutationPermissionRequest_Mutation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutationPermissionRequest_Mutation.Merge(m, src)
}
func (m *MutationPermissionRequest_Mutation) XXX_Size() int {
	return xxx_messageInfo_MutationPermissionRequest_Mutation.Size(m)
}
func (m *MutationPermissionRequest_Mutation) XXX_DiscardUnknown() {
	xxx_messageInfo_MutationPermissionRequest_Mutation.DiscardUnknown(m)
}

var xxx_messageInfo_MutationPermissionRequest_Mutation proto.InternalMessageInfo

func (m *MutationPermissionRequest_Mutation) GetIncludedReadableColumns() map[string]*MutationPermissionRequest_Mutation_Column {
	if m != nil {
		return m.IncludedReadableColumns
	}
	return nil
}

func (m *MutationPermissionRequest_Mutation) GetExcludedReadableColumns() map[string]*MutationPermissionRequest_Mutation_Column {
	if m != nil {
		return m.ExcludedReadableColumns
	}
	return nil
}

func (m *MutationPermissionRequest_Mutation) GetIncludedWritableColumns() map[string]*MutationPermissionRequest_Mutation_Column {
	if m != nil {
		return m.IncludedWritableColumns
	}
	return nil
}

func (m *MutationPermissionRequest_Mutation) GetExcludedWritableColumns() map[string]*MutationPermissionRequest_Mutation_Column {
	if m != nil {
		return m.ExcludedWritableColumns
	}
	return nil
}

type MutationPermissionRequest_Mutation_Column struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutationPermissionRequest_Mutation_Column) Reset() {
	*m = MutationPermissionRequest_Mutation_Column{}
}
func (m *MutationPermissionRequest_Mutation_Column) String() string {
	return proto.CompactTextString(m)
}
func (*MutationPermissionRequest_Mutation_Column) ProtoMessage() {}
func (*MutationPermissionRequest_Mutation_Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{2, 0, 0}
}

func (m *MutationPermissionRequest_Mutation_Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutationPermissionRequest_Mutation_Column.Unmarshal(m, b)
}
func (m *MutationPermissionRequest_Mutation_Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutationPermissionRequest_Mutation_Column.Marshal(b, m, deterministic)
}
func (m *MutationPermissionRequest_Mutation_Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutationPermissionRequest_Mutation_Column.Merge(m, src)
}
func (m *MutationPermissionRequest_Mutation_Column) XXX_Size() int {
	return xxx_messageInfo_MutationPermissionRequest_Mutation_Column.Size(m)
}
func (m *MutationPermissionRequest_Mutation_Column) XXX_DiscardUnknown() {
	xxx_messageInfo_MutationPermissionRequest_Mutation_Column.DiscardUnknown(m)
}

var xxx_messageInfo_MutationPermissionRequest_Mutation_Column proto.InternalMessageInfo

func (m *MutationPermissionRequest_Mutation_Column) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type PermissionRequest struct {
	Id                   Permission `protobuf:"varint,1,opt,name=id,proto3,enum=proto.dm.utility.Permission" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PermissionRequest) Reset()         { *m = PermissionRequest{} }
func (m *PermissionRequest) String() string { return proto.CompactTextString(m) }
func (*PermissionRequest) ProtoMessage()    {}
func (*PermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{3}
}

func (m *PermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionRequest.Unmarshal(m, b)
}
func (m *PermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionRequest.Marshal(b, m, deterministic)
}
func (m *PermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionRequest.Merge(m, src)
}
func (m *PermissionRequest) XXX_Size() int {
	return xxx_messageInfo_PermissionRequest.Size(m)
}
func (m *PermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionRequest proto.InternalMessageInfo

func (m *PermissionRequest) GetId() Permission {
	if m != nil {
		return m.Id
	}
	return Permission_NO_PERMISSION
}

type PermissionDetailReply struct {
	Id                   Permission                     `protobuf:"varint,1,opt,name=id,proto3,enum=proto.dm.utility.Permission" json:"id,omitempty"`
	ReadableTables       []*PermissionDetailReply_Table `protobuf:"bytes,2,rep,name=readable_tables,json=readableTables,proto3" json:"readable_tables,omitempty"`
	WritableTables       []*PermissionDetailReply_Table `protobuf:"bytes,3,rep,name=writable_tables,json=writableTables,proto3" json:"writable_tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *PermissionDetailReply) Reset()         { *m = PermissionDetailReply{} }
func (m *PermissionDetailReply) String() string { return proto.CompactTextString(m) }
func (*PermissionDetailReply) ProtoMessage()    {}
func (*PermissionDetailReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{4}
}

func (m *PermissionDetailReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionDetailReply.Unmarshal(m, b)
}
func (m *PermissionDetailReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionDetailReply.Marshal(b, m, deterministic)
}
func (m *PermissionDetailReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionDetailReply.Merge(m, src)
}
func (m *PermissionDetailReply) XXX_Size() int {
	return xxx_messageInfo_PermissionDetailReply.Size(m)
}
func (m *PermissionDetailReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionDetailReply.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionDetailReply proto.InternalMessageInfo

func (m *PermissionDetailReply) GetId() Permission {
	if m != nil {
		return m.Id
	}
	return Permission_NO_PERMISSION
}

func (m *PermissionDetailReply) GetReadableTables() []*PermissionDetailReply_Table {
	if m != nil {
		return m.ReadableTables
	}
	return nil
}

func (m *PermissionDetailReply) GetWritableTables() []*PermissionDetailReply_Table {
	if m != nil {
		return m.WritableTables
	}
	return nil
}

type PermissionDetailReply_Table struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExcludedColumns      []string `protobuf:"bytes,2,rep,name=excluded_columns,json=excludedColumns,proto3" json:"excluded_columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermissionDetailReply_Table) Reset()         { *m = PermissionDetailReply_Table{} }
func (m *PermissionDetailReply_Table) String() string { return proto.CompactTextString(m) }
func (*PermissionDetailReply_Table) ProtoMessage()    {}
func (*PermissionDetailReply_Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{4, 0}
}

func (m *PermissionDetailReply_Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionDetailReply_Table.Unmarshal(m, b)
}
func (m *PermissionDetailReply_Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionDetailReply_Table.Marshal(b, m, deterministic)
}
func (m *PermissionDetailReply_Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionDetailReply_Table.Merge(m, src)
}
func (m *PermissionDetailReply_Table) XXX_Size() int {
	return xxx_messageInfo_PermissionDetailReply_Table.Size(m)
}
func (m *PermissionDetailReply_Table) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionDetailReply_Table.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionDetailReply_Table proto.InternalMessageInfo

func (m *PermissionDetailReply_Table) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PermissionDetailReply_Table) GetExcludedColumns() []string {
	if m != nil {
		return m.ExcludedColumns
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.dm.utility.Permission", Permission_name, Permission_value)
	proto.RegisterType((*Permissions)(nil), "proto.dm.utility.Permissions")
	proto.RegisterType((*UserPermission)(nil), "proto.dm.utility.UserPermission")
	proto.RegisterType((*MutationPermissionRequest)(nil), "proto.dm.utility.MutationPermissionRequest")
	proto.RegisterType((*MutationPermissionRequest_Mutation)(nil), "proto.dm.utility.MutationPermissionRequest.Mutation")
	proto.RegisterMapType((map[string]*MutationPermissionRequest_Mutation_Column)(nil), "proto.dm.utility.MutationPermissionRequest.Mutation.ExcludedReadableColumnsEntry")
	proto.RegisterMapType((map[string]*MutationPermissionRequest_Mutation_Column)(nil), "proto.dm.utility.MutationPermissionRequest.Mutation.ExcludedWritableColumnsEntry")
	proto.RegisterMapType((map[string]*MutationPermissionRequest_Mutation_Column)(nil), "proto.dm.utility.MutationPermissionRequest.Mutation.IncludedReadableColumnsEntry")
	proto.RegisterMapType((map[string]*MutationPermissionRequest_Mutation_Column)(nil), "proto.dm.utility.MutationPermissionRequest.Mutation.IncludedWritableColumnsEntry")
	proto.RegisterType((*MutationPermissionRequest_Mutation_Column)(nil), "proto.dm.utility.MutationPermissionRequest.Mutation.Column")
	proto.RegisterType((*PermissionRequest)(nil), "proto.dm.utility.PermissionRequest")
	proto.RegisterType((*PermissionDetailReply)(nil), "proto.dm.utility.PermissionDetailReply")
	proto.RegisterType((*PermissionDetailReply_Table)(nil), "proto.dm.utility.PermissionDetailReply.Table")
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }

var fileDescriptor_c837ef01cbda0ad8 = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x6e, 0xda, 0x48,
	0x14, 0x5e, 0x43, 0x80, 0x70, 0xd8, 0x10, 0x67, 0x36, 0xab, 0x24, 0x6c, 0xb4, 0x41, 0x5c, 0x65,
	0x57, 0xbb, 0xb6, 0x44, 0xab, 0xaa, 0x7f, 0xaa, 0xe4, 0xc0, 0xd0, 0xba, 0x02, 0x9b, 0x4c, 0xa0,
	0x69, 0x7a, 0x63, 0x19, 0x3c, 0x45, 0x56, 0xfc, 0x43, 0xf1, 0x38, 0x09, 0x4f, 0xd0, 0xab, 0xde,
	0xf4, 0x25, 0x7a, 0xdb, 0x17, 0xe8, 0x13, 0xf4, 0xa5, 0x2a, 0xff, 0x01, 0x49, 0x08, 0x51, 0x93,
	0xa8, 0xb9, 0x62, 0x66, 0xbe, 0x73, 0xbe, 0xf3, 0x9d, 0x73, 0x66, 0x0e, 0x06, 0x7e, 0x48, 0x47,
	0xb6, 0xe9, 0x79, 0xa6, 0xeb, 0x08, 0xc3, 0x91, 0xcb, 0x5c, 0xc4, 0x87, 0x3f, 0x82, 0x61, 0x0b,
	0x3e, 0x33, 0x2d, 0x93, 0x8d, 0x4b, 0x3b, 0x03, 0xd7, 0x1d, 0x58, 0x54, 0x0c, 0x81, 0x9e, 0xff,
	0x5e, 0x64, 0xa6, 0x4d, 0x3d, 0xa6, 0xdb, 0xc3, 0xc8, 0xa5, 0x82, 0xa1, 0xd0, 0x9e, 0xd0, 0x78,
	0xe8, 0x11, 0xe4, 0x0c, 0x9d, 0xe9, 0x1e, 0x65, 0x9b, 0x5c, 0x39, 0xbd, 0x5b, 0xac, 0x6e, 0x0b,
	0x17, 0x39, 0x85, 0xa9, 0x3d, 0x49, 0x8c, 0x2b, 0x5f, 0x38, 0x28, 0x76, 0x3d, 0x3a, 0x9a, 0x62,
	0x68, 0x03, 0x72, 0xbe, 0x47, 0x47, 0x9a, 0x69, 0x6c, 0x72, 0x65, 0x6e, 0x37, 0x4f, 0xb2, 0xc1,
	0x56, 0x36, 0x90, 0x04, 0x2b, 0x53, 0xe5, 0x01, 0x9c, 0x2a, 0x73, 0xd7, 0x46, 0xfa, 0x7d, 0xea,
	0x22, 0x1b, 0xe8, 0x09, 0xc0, 0x89, 0x6e, 0x99, 0x86, 0x16, 0xa4, 0xb3, 0x99, 0x2e, 0x73, 0xbb,
	0x85, 0x6a, 0x49, 0x88, 0x72, 0x15, 0x92, 0x5c, 0x85, 0x4e, 0x92, 0x2b, 0xc9, 0x87, 0xd6, 0xc1,
	0xbe, 0xf2, 0x15, 0x60, 0xab, 0xe5, 0x33, 0x9d, 0x99, 0xae, 0x33, 0xc3, 0x4f, 0x3f, 0xf8, 0xd4,
	0x63, 0xe8, 0x3f, 0x48, 0xc5, 0x7a, 0xaf, 0x13, 0x94, 0x32, 0x0d, 0xd4, 0x86, 0x65, 0x3b, 0xa6,
	0x0a, 0x93, 0x28, 0x54, 0x1f, 0x5e, 0xf6, 0xb9, 0x32, 0xd8, 0x04, 0x21, 0x13, 0x96, 0xd2, 0xe7,
	0x3c, 0x2c, 0x27, 0xc7, 0xe8, 0x13, 0x07, 0x5b, 0xa6, 0xd3, 0xb7, 0x7c, 0x83, 0x1a, 0xda, 0x88,
	0xea, 0x86, 0xde, 0xb3, 0xa8, 0xd6, 0x77, 0x2d, 0xdf, 0x76, 0xbc, 0xb0, 0x3f, 0x85, 0xea, 0xfe,
	0x4d, 0x02, 0x0a, 0x72, 0xcc, 0x4a, 0x62, 0xd2, 0x5a, 0xc4, 0x89, 0x1d, 0x36, 0x1a, 0x93, 0x0d,
	0x73, 0x3e, 0x1a, 0xea, 0xa1, 0x67, 0x57, 0xe9, 0x49, 0xdd, 0x42, 0x0f, 0x3e, 0x5b, 0xa4, 0x87,
	0x9e, 0x5d, 0xad, 0x67, 0x52, 0x9f, 0xd3, 0x91, 0xc9, 0xce, 0xe9, 0x49, 0xdf, 0x41, 0x7d, 0x0e,
	0x63, 0xd2, 0xf9, 0xf5, 0xb9, 0x80, 0x9e, 0xaf, 0xcf, 0x25, 0x3d, 0x4b, 0x77, 0x50, 0x9f, 0xf9,
	0x7a, 0xe8, 0x7c, 0xb4, 0xf4, 0x37, 0x64, 0xa3, 0x25, 0x5a, 0x87, 0x8c, 0xa3, 0xdb, 0x34, 0xba,
	0x34, 0x79, 0x12, 0x6d, 0x4a, 0x1f, 0x39, 0xd8, 0x5e, 0x74, 0x13, 0x10, 0x0f, 0xe9, 0x63, 0x3a,
	0x8e, 0x9f, 0x6f, 0xb0, 0x44, 0xfb, 0x90, 0x39, 0xd1, 0x2d, 0x9f, 0xc6, 0xd7, 0xfd, 0xd9, 0x8d,
	0xb2, 0x89, 0x62, 0x90, 0x88, 0xe9, 0x69, 0xea, 0x31, 0x17, 0x2a, 0x59, 0x74, 0x07, 0x7e, 0xad,
	0x92, 0x45, 0xdd, 0xbf, 0x9f, 0x9a, 0xdc, 0xab, 0x92, 0x8a, 0x04, 0x6b, 0xb7, 0x9c, 0x94, 0x95,
	0x6f, 0x29, 0xf8, 0x73, 0x7a, 0x54, 0xa7, 0x4c, 0x37, 0x2d, 0x42, 0x87, 0xd6, 0xf8, 0x27, 0x27,
	0xee, 0x1b, 0x58, 0x9d, 0x0c, 0x9e, 0xb0, 0x22, 0xc9, 0xdc, 0xf9, 0x7f, 0x91, 0xeb, 0x4c, 0x3c,
	0xa1, 0x13, 0x78, 0x91, 0x62, 0xc2, 0x12, 0x6e, 0xbd, 0x80, 0x77, 0xf2, 0x60, 0x63, 0xde, 0xf4,
	0x8d, 0x78, 0x13, 0x96, 0x88, 0xb7, 0xd4, 0x80, 0x4c, 0xb8, 0x42, 0x08, 0x96, 0x82, 0x47, 0x17,
	0x77, 0x2b, 0x5c, 0xa3, 0x7f, 0x80, 0x9f, 0x8c, 0x8b, 0xd9, 0x29, 0x9a, 0x27, 0xab, 0xc9, 0x79,
	0xdc, 0xf0, 0x7f, 0xbf, 0x73, 0x00, 0x33, 0xff, 0xad, 0x6b, 0xb0, 0xa2, 0xa8, 0x5a, 0x1b, 0x93,
	0x96, 0x7c, 0x70, 0x20, 0xab, 0x0a, 0xff, 0x5b, 0x70, 0xd4, 0x68, 0xe2, 0xb7, 0xf2, 0x5e, 0x13,
	0x6b, 0xaf, 0xd4, 0x2e, 0xe1, 0x39, 0xb4, 0x03, 0x7f, 0xb5, 0x6b, 0x44, 0xab, 0x63, 0xa9, 0x89,
	0x49, 0x60, 0xdd, 0x50, 0x49, 0x4b, 0x52, 0x6a, 0x58, 0xab, 0x4b, 0x1d, 0x89, 0x4f, 0xa1, 0x75,
	0xe0, 0xa5, 0x5a, 0x4d, 0xed, 0x2a, 0x1d, 0x59, 0x79, 0xa9, 0x61, 0xa5, 0x43, 0x8e, 0xf8, 0x34,
	0xca, 0x41, 0xba, 0x85, 0x0f, 0xf8, 0x25, 0xb4, 0x01, 0x7f, 0xb4, 0x89, 0xfa, 0x1a, 0xd7, 0x3a,
	0xda, 0xa1, 0xac, 0xd4, 0xd5, 0x43, 0x4d, 0x6d, 0x63, 0x85, 0xcf, 0x04, 0xb1, 0x12, 0x60, 0xbf,
	0x8b, 0xc9, 0x11, 0x9f, 0x9d, 0xb5, 0x6d, 0xc8, 0x4d, 0xac, 0x75, 0xdb, 0x4d, 0x55, 0xaa, 0xf3,
	0x39, 0x84, 0xa0, 0x98, 0x00, 0x75, 0xdc, 0x91, 0xe4, 0x26, 0xbf, 0xbc, 0xf7, 0xe2, 0xdd, 0xf3,
	0x81, 0xc9, 0x2c, 0xbd, 0x27, 0x1c, 0x53, 0xc7, 0xd0, 0x85, 0xbe, 0x6b, 0x0b, 0xec, 0x54, 0x0c,
	0x37, 0x62, 0xdf, 0xb5, 0x6d, 0xd7, 0xf1, 0xc4, 0x93, 0x6a, 0xf4, 0xd5, 0x22, 0x0e, 0x5c, 0x4b,
	0x77, 0x06, 0xa2, 0x61, 0x8b, 0x71, 0x07, 0x7a, 0xd9, 0x10, 0x78, 0xf0, 0x23, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0x08, 0x06, 0x7c, 0x02, 0x09, 0x00, 0x00,
}
