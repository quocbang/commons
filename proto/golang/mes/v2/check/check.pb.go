// Code generated by protoc-gen-go. DO NOT EDIT.
// source: check/check.proto

package check

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

type ActionMode int32

const (
	ActionMode_ACTION_AUTO     ActionMode = 0
	ActionMode_ACTION_FORCE    ActionMode = 1
	ActionMode_ACTION_PRECHECK ActionMode = 2
)

var ActionMode_name = map[int32]string{
	0: "ACTION_AUTO",
	1: "ACTION_FORCE",
	2: "ACTION_PRECHECK",
}

var ActionMode_value = map[string]int32{
	"ACTION_AUTO":     0,
	"ACTION_FORCE":    1,
	"ACTION_PRECHECK": 2,
}

func (x ActionMode) String() string {
	return proto.EnumName(ActionMode_name, int32(x))
}

func (ActionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9059152c7fbd1876, []int{0}
}

type ResourceStatus int32

const (
	ResourceStatus_RESOURCE_STATUS_NONE      ResourceStatus = 0
	ResourceStatus_RESOURCE_STATUS_BOUND     ResourceStatus = 1
	ResourceStatus_RESOURCE_STATUS_PRESERVED ResourceStatus = 2
	ResourceStatus_RESOURCE_STATUS_READY     ResourceStatus = 3
	ResourceStatus_RESOURCE_STATUS_FED       ResourceStatus = 4
)

var ResourceStatus_name = map[int32]string{
	0: "RESOURCE_STATUS_NONE",
	1: "RESOURCE_STATUS_BOUND",
	2: "RESOURCE_STATUS_PRESERVED",
	3: "RESOURCE_STATUS_READY",
	4: "RESOURCE_STATUS_FED",
}

var ResourceStatus_value = map[string]int32{
	"RESOURCE_STATUS_NONE":      0,
	"RESOURCE_STATUS_BOUND":     1,
	"RESOURCE_STATUS_PRESERVED": 2,
	"RESOURCE_STATUS_READY":     3,
	"RESOURCE_STATUS_FED":       4,
}

func (x ResourceStatus) String() string {
	return proto.EnumName(ResourceStatus_name, int32(x))
}

func (ResourceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9059152c7fbd1876, []int{1}
}

// Checks if a resource action is ok.
// If workOrder is empty, simply check if the action is possible & allowed, e.g., material is enough and not expired.
// Otherwise check if the action is expected for the work order, e.g., material or tool is required by the work order.
type ResourceAction struct {
	ActionMode           ActionMode                `protobuf:"varint,1,opt,name=actionMode,proto3,enum=proto.mes.v2.check.ActionMode" json:"actionMode,omitempty"`
	WorkOrder            string                    `protobuf:"bytes,2,opt,name=workOrder,proto3" json:"workOrder,omitempty"`
	Batch                uint32                    `protobuf:"varint,3,opt,name=batch,proto3" json:"batch,omitempty"`
	GroupRequirements    map[string]ResourceStatus `protobuf:"bytes,4,rep,name=groupRequirements,proto3" json:"groupRequirements,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=proto.mes.v2.check.ResourceStatus"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ResourceAction) Reset()         { *m = ResourceAction{} }
func (m *ResourceAction) String() string { return proto.CompactTextString(m) }
func (*ResourceAction) ProtoMessage()    {}
func (*ResourceAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_9059152c7fbd1876, []int{0}
}

func (m *ResourceAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceAction.Unmarshal(m, b)
}
func (m *ResourceAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceAction.Marshal(b, m, deterministic)
}
func (m *ResourceAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceAction.Merge(m, src)
}
func (m *ResourceAction) XXX_Size() int {
	return xxx_messageInfo_ResourceAction.Size(m)
}
func (m *ResourceAction) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceAction.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceAction proto.InternalMessageInfo

func (m *ResourceAction) GetActionMode() ActionMode {
	if m != nil {
		return m.ActionMode
	}
	return ActionMode_ACTION_AUTO
}

func (m *ResourceAction) GetWorkOrder() string {
	if m != nil {
		return m.WorkOrder
	}
	return ""
}

func (m *ResourceAction) GetBatch() uint32 {
	if m != nil {
		return m.Batch
	}
	return 0
}

func (m *ResourceAction) GetGroupRequirements() map[string]ResourceStatus {
	if m != nil {
		return m.GroupRequirements
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.mes.v2.check.ActionMode", ActionMode_name, ActionMode_value)
	proto.RegisterEnum("proto.mes.v2.check.ResourceStatus", ResourceStatus_name, ResourceStatus_value)
	proto.RegisterType((*ResourceAction)(nil), "proto.mes.v2.check.ResourceAction")
	proto.RegisterMapType((map[string]ResourceStatus)(nil), "proto.mes.v2.check.ResourceAction.GroupRequirementsEntry")
}

func init() { proto.RegisterFile("check/check.proto", fileDescriptor_9059152c7fbd1876) }

var fileDescriptor_9059152c7fbd1876 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x9e, 0xd3, 0x0d, 0xa9, 0x37, 0xe8, 0x3c, 0x6f, 0x40, 0x86, 0x00, 0x55, 0x7b, 0xaa, 0xf6,
	0xe0, 0x48, 0xe1, 0x65, 0xf0, 0x30, 0x91, 0x25, 0x1e, 0x20, 0x44, 0x32, 0x39, 0x09, 0x12, 0xbc,
	0x54, 0x69, 0x6a, 0xa5, 0x55, 0x9a, 0x78, 0x38, 0x4e, 0xa7, 0xfd, 0x13, 0xfe, 0x19, 0x7f, 0x07,
	0xe1, 0x4e, 0x14, 0xd2, 0x6a, 0x2f, 0xf6, 0xdd, 0xf7, 0xdd, 0xe7, 0xbb, 0xfb, 0x0c, 0x87, 0xf9,
	0x4c, 0xe4, 0xa5, 0x63, 0x4e, 0x7a, 0xa3, 0xa4, 0x96, 0x84, 0x98, 0x8b, 0x56, 0xa2, 0xa1, 0x4b,
	0x97, 0x1a, 0xe6, 0xf4, 0x97, 0x05, 0x03, 0x2e, 0x1a, 0xd9, 0xaa, 0x5c, 0x78, 0xb9, 0x9e, 0xcb,
	0x9a, 0x5c, 0x00, 0x64, 0x26, 0xfa, 0x22, 0xa7, 0xc2, 0x46, 0x43, 0x34, 0x1a, 0xb8, 0xaf, 0xe9,
	0xa6, 0x96, 0x7a, 0x7f, 0xab, 0xf8, 0x3f, 0x0a, 0xf2, 0x12, 0xfa, 0xb7, 0x52, 0x95, 0x91, 0x9a,
	0x0a, 0x65, 0x5b, 0x43, 0x34, 0xea, 0xf3, 0x35, 0x40, 0x8e, 0x61, 0x6f, 0x92, 0xe9, 0x7c, 0x66,
	0xf7, 0x86, 0x68, 0xf4, 0x84, 0xaf, 0x12, 0x52, 0xc0, 0x61, 0xa1, 0x64, 0x7b, 0xc3, 0xc5, 0x8f,
	0x76, 0xae, 0x44, 0x25, 0x6a, 0xdd, 0xd8, 0xbb, 0xc3, 0xde, 0x68, 0xdf, 0x7d, 0xbb, 0xad, 0xf5,
	0xff, 0x23, 0xd3, 0x0f, 0x5d, 0x2d, 0xab, 0xb5, 0xba, 0xe3, 0x9b, 0x6f, 0xbe, 0x98, 0xc1, 0xb3,
	0xed, 0xc5, 0x04, 0x43, 0xaf, 0x14, 0x77, 0x66, 0xdf, 0x3e, 0xff, 0x13, 0x92, 0x73, 0xd8, 0x5b,
	0x66, 0x8b, 0x56, 0x98, 0x25, 0x06, 0xee, 0xe9, 0x43, 0x83, 0xc4, 0x3a, 0xd3, 0x6d, 0xc3, 0x57,
	0x82, 0x77, 0xd6, 0x39, 0x3a, 0x0b, 0x00, 0xd6, 0x06, 0x91, 0x03, 0xd8, 0xf7, 0xfc, 0xe4, 0x53,
	0x14, 0x8e, 0xbd, 0x34, 0x89, 0xf0, 0x0e, 0xc1, 0xf0, 0xf8, 0x1e, 0xb8, 0x8a, 0xb8, 0xcf, 0x30,
	0x22, 0x47, 0x70, 0x70, 0x8f, 0x5c, 0x73, 0xe6, 0x7f, 0x64, 0xfe, 0x67, 0x6c, 0x9d, 0xfd, 0x44,
	0xeb, 0xff, 0x59, 0xf5, 0x20, 0x36, 0x1c, 0x73, 0x16, 0x47, 0x29, 0xf7, 0xd9, 0x38, 0x4e, 0xbc,
	0x24, 0x8d, 0xc7, 0x61, 0x14, 0x32, 0xbc, 0x43, 0x4e, 0xe0, 0x69, 0x97, 0xb9, 0x8c, 0xd2, 0x30,
	0xc0, 0x88, 0xbc, 0x82, 0x93, 0x2e, 0x75, 0xcd, 0x59, 0xcc, 0xf8, 0x57, 0x16, 0x60, 0x6b, 0x9b,
	0x92, 0x33, 0x2f, 0xf8, 0x86, 0x7b, 0xe4, 0x39, 0x1c, 0x75, 0xa9, 0x2b, 0x16, 0xe0, 0xdd, 0xcb,
	0xf7, 0xdf, 0x2f, 0x8a, 0xb9, 0x5e, 0x64, 0x13, 0x5a, 0x8a, 0x7a, 0x9a, 0xd1, 0x5c, 0x56, 0x54,
	0xdf, 0x3a, 0x26, 0x71, 0x72, 0x59, 0x55, 0xb2, 0x6e, 0x9c, 0xa5, 0xeb, 0x18, 0xd7, 0x9c, 0x42,
	0x2e, 0xb2, 0xba, 0x70, 0x2a, 0x61, 0x30, 0x63, 0xde, 0xe4, 0x91, 0xa1, 0xde, 0xfc, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x85, 0x74, 0x53, 0xe1, 0xac, 0x02, 0x00, 0x00,
}
