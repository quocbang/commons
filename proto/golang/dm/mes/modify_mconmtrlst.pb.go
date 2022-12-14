// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modify_mconmtrlst.proto

package mes

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

// Update Meqptconmtrlst Info
type UpdateMconmtrlstRequest struct {
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	MtrlLotId            string   `protobuf:"bytes,2,opt,name=mtrl_lot_id,json=mtrlLotId,proto3" json:"mtrl_lot_id,omitempty"`
	NewMtrlLotId         string   `protobuf:"bytes,3,opt,name=new_mtrl_lot_id,json=newMtrlLotId,proto3" json:"new_mtrl_lot_id,omitempty"`
	ActionOrder          string   `protobuf:"bytes,4,opt,name=action_order,json=actionOrder,proto3" json:"action_order,omitempty"`
	NewActionType        string   `protobuf:"bytes,5,opt,name=new_action_type,json=newActionType,proto3" json:"new_action_type,omitempty"`
	UpdateBy             string   `protobuf:"bytes,6,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	UpdateAt             string   `protobuf:"bytes,7,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMconmtrlstRequest) Reset()         { *m = UpdateMconmtrlstRequest{} }
func (m *UpdateMconmtrlstRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMconmtrlstRequest) ProtoMessage()    {}
func (*UpdateMconmtrlstRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fad1d815e9ecba4d, []int{0}
}

func (m *UpdateMconmtrlstRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMconmtrlstRequest.Unmarshal(m, b)
}
func (m *UpdateMconmtrlstRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMconmtrlstRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMconmtrlstRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMconmtrlstRequest.Merge(m, src)
}
func (m *UpdateMconmtrlstRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMconmtrlstRequest.Size(m)
}
func (m *UpdateMconmtrlstRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMconmtrlstRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMconmtrlstRequest proto.InternalMessageInfo

func (m *UpdateMconmtrlstRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *UpdateMconmtrlstRequest) GetMtrlLotId() string {
	if m != nil {
		return m.MtrlLotId
	}
	return ""
}

func (m *UpdateMconmtrlstRequest) GetNewMtrlLotId() string {
	if m != nil {
		return m.NewMtrlLotId
	}
	return ""
}

func (m *UpdateMconmtrlstRequest) GetActionOrder() string {
	if m != nil {
		return m.ActionOrder
	}
	return ""
}

func (m *UpdateMconmtrlstRequest) GetNewActionType() string {
	if m != nil {
		return m.NewActionType
	}
	return ""
}

func (m *UpdateMconmtrlstRequest) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *UpdateMconmtrlstRequest) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

type UpdateMconmtrlstReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMconmtrlstReply) Reset()         { *m = UpdateMconmtrlstReply{} }
func (m *UpdateMconmtrlstReply) String() string { return proto.CompactTextString(m) }
func (*UpdateMconmtrlstReply) ProtoMessage()    {}
func (*UpdateMconmtrlstReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_fad1d815e9ecba4d, []int{1}
}

func (m *UpdateMconmtrlstReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMconmtrlstReply.Unmarshal(m, b)
}
func (m *UpdateMconmtrlstReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMconmtrlstReply.Marshal(b, m, deterministic)
}
func (m *UpdateMconmtrlstReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMconmtrlstReply.Merge(m, src)
}
func (m *UpdateMconmtrlstReply) XXX_Size() int {
	return xxx_messageInfo_UpdateMconmtrlstReply.Size(m)
}
func (m *UpdateMconmtrlstReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMconmtrlstReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMconmtrlstReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateMconmtrlstRequest)(nil), "proto.dm.mes.UpdateMconmtrlstRequest")
	proto.RegisterType((*UpdateMconmtrlstReply)(nil), "proto.dm.mes.UpdateMconmtrlstReply")
}

func init() { proto.RegisterFile("modify_mconmtrlst.proto", fileDescriptor_fad1d815e9ecba4d) }

var fileDescriptor_fad1d815e9ecba4d = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x86, 0xd9, 0xbe, 0xcf, 0xe9, 0xb2, 0x8a, 0x10, 0x90, 0x16, 0x04, 0x71, 0x03, 0xc5, 0xab,
	0x04, 0xf4, 0x46, 0xbc, 0xdb, 0xee, 0x06, 0x0e, 0x61, 0xe8, 0x8d, 0x37, 0x21, 0x6d, 0x8e, 0xa5,
	0xd8, 0xe4, 0xd4, 0xf6, 0xcc, 0x92, 0x7f, 0xe0, 0xcf, 0x96, 0xa6, 0x52, 0x0a, 0x5e, 0x85, 0x3c,
	0xef, 0x93, 0x1c, 0xce, 0xcb, 0x62, 0x8b, 0xa6, 0x78, 0xf7, 0xca, 0x66, 0xe8, 0x2c, 0xd5, 0x65,
	0x43, 0xa2, 0xaa, 0x91, 0x90, 0x47, 0xe1, 0x10, 0xc6, 0x0a, 0x0b, 0xcd, 0xea, 0x7b, 0xca, 0xe2,
	0xd7, 0xca, 0x68, 0x82, 0xdd, 0x20, 0xee, 0xe1, 0xf3, 0x00, 0x0d, 0xf1, 0x25, 0x8b, 0x32, 0x74,
	0xa4, 0x0b, 0x07, 0xb5, 0x2a, 0x4c, 0x32, 0xb9, 0x9a, 0xdc, 0xce, 0xf7, 0x8b, 0x81, 0x6d, 0x0d,
	0xbf, 0x64, 0x8b, 0xee, 0x8d, 0x2a, 0x91, 0x3a, 0x63, 0x1a, 0x8c, 0x79, 0x87, 0x9e, 0x90, 0xb6,
	0x86, 0x5f, 0xb3, 0x33, 0x07, 0xad, 0x1a, 0x3b, 0xff, 0x82, 0x13, 0x39, 0x68, 0x77, 0x83, 0xb6,
	0x64, 0x91, 0xce, 0xa8, 0x40, 0xa7, 0xb0, 0x36, 0x50, 0x27, 0xff, 0xfb, 0x49, 0x3d, 0x7b, 0xee,
	0x10, 0xbf, 0xe9, 0x7f, 0xfa, 0xd5, 0xc8, 0x57, 0x90, 0x1c, 0x05, 0xeb, 0xd4, 0x41, 0xbb, 0x0e,
	0xf4, 0xc5, 0x57, 0xc0, 0x2f, 0xd8, 0xfc, 0x10, 0xf6, 0x51, 0xa9, 0x4f, 0x66, 0xc1, 0x38, 0xe9,
	0xc1, 0xc6, 0x8f, 0x42, 0x4d, 0xc9, 0xf1, 0x38, 0x5c, 0xd3, 0x2a, 0x66, 0xe7, 0x7f, 0x9b, 0xa8,
	0x4a, 0xbf, 0x79, 0x7c, 0x7b, 0xc8, 0x0b, 0x2a, 0x75, 0x2a, 0x3e, 0xc0, 0x19, 0x2d, 0x32, 0xb4,
	0x82, 0x5a, 0x19, 0x2e, 0x32, 0x43, 0x6b, 0xd1, 0x35, 0xf2, 0xeb, 0x4e, 0x86, 0x62, 0x65, 0x8e,
	0xa5, 0x76, 0xb9, 0x34, 0x56, 0x5a, 0x68, 0xd2, 0x59, 0x80, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x57, 0x1c, 0x27, 0x84, 0x8f, 0x01, 0x00, 0x00,
}
