// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query_mconmtrlst.proto

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

type GetMconmtrlstRequest struct {
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	MtrlLotId            string   `protobuf:"bytes,2,opt,name=mtrl_lot_id,json=mtrlLotId,proto3" json:"mtrl_lot_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMconmtrlstRequest) Reset()         { *m = GetMconmtrlstRequest{} }
func (m *GetMconmtrlstRequest) String() string { return proto.CompactTextString(m) }
func (*GetMconmtrlstRequest) ProtoMessage()    {}
func (*GetMconmtrlstRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d9d7a06040131d, []int{0}
}

func (m *GetMconmtrlstRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconmtrlstRequest.Unmarshal(m, b)
}
func (m *GetMconmtrlstRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconmtrlstRequest.Marshal(b, m, deterministic)
}
func (m *GetMconmtrlstRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconmtrlstRequest.Merge(m, src)
}
func (m *GetMconmtrlstRequest) XXX_Size() int {
	return xxx_messageInfo_GetMconmtrlstRequest.Size(m)
}
func (m *GetMconmtrlstRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconmtrlstRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconmtrlstRequest proto.InternalMessageInfo

func (m *GetMconmtrlstRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *GetMconmtrlstRequest) GetMtrlLotId() string {
	if m != nil {
		return m.MtrlLotId
	}
	return ""
}

type GetMconmtrlstReply struct {
	Info                 []*GetMconmtrlstReply_Mconmtrlst `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetMconmtrlstReply) Reset()         { *m = GetMconmtrlstReply{} }
func (m *GetMconmtrlstReply) String() string { return proto.CompactTextString(m) }
func (*GetMconmtrlstReply) ProtoMessage()    {}
func (*GetMconmtrlstReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d9d7a06040131d, []int{1}
}

func (m *GetMconmtrlstReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconmtrlstReply.Unmarshal(m, b)
}
func (m *GetMconmtrlstReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconmtrlstReply.Marshal(b, m, deterministic)
}
func (m *GetMconmtrlstReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconmtrlstReply.Merge(m, src)
}
func (m *GetMconmtrlstReply) XXX_Size() int {
	return xxx_messageInfo_GetMconmtrlstReply.Size(m)
}
func (m *GetMconmtrlstReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconmtrlstReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconmtrlstReply proto.InternalMessageInfo

func (m *GetMconmtrlstReply) GetInfo() []*GetMconmtrlstReply_Mconmtrlst {
	if m != nil {
		return m.Info
	}
	return nil
}

type GetMconmtrlstReply_Mconmtrlst struct {
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	MtrlLotId            string   `protobuf:"bytes,2,opt,name=mtrl_lot_id,json=mtrlLotId,proto3" json:"mtrl_lot_id,omitempty"`
	ActionType           string   `protobuf:"bytes,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	UpdateBy             string   `protobuf:"bytes,4,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	UpdateAt             string   `protobuf:"bytes,5,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMconmtrlstReply_Mconmtrlst) Reset()         { *m = GetMconmtrlstReply_Mconmtrlst{} }
func (m *GetMconmtrlstReply_Mconmtrlst) String() string { return proto.CompactTextString(m) }
func (*GetMconmtrlstReply_Mconmtrlst) ProtoMessage()    {}
func (*GetMconmtrlstReply_Mconmtrlst) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d9d7a06040131d, []int{1, 0}
}

func (m *GetMconmtrlstReply_Mconmtrlst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconmtrlstReply_Mconmtrlst.Unmarshal(m, b)
}
func (m *GetMconmtrlstReply_Mconmtrlst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconmtrlstReply_Mconmtrlst.Marshal(b, m, deterministic)
}
func (m *GetMconmtrlstReply_Mconmtrlst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconmtrlstReply_Mconmtrlst.Merge(m, src)
}
func (m *GetMconmtrlstReply_Mconmtrlst) XXX_Size() int {
	return xxx_messageInfo_GetMconmtrlstReply_Mconmtrlst.Size(m)
}
func (m *GetMconmtrlstReply_Mconmtrlst) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconmtrlstReply_Mconmtrlst.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconmtrlstReply_Mconmtrlst proto.InternalMessageInfo

func (m *GetMconmtrlstReply_Mconmtrlst) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *GetMconmtrlstReply_Mconmtrlst) GetMtrlLotId() string {
	if m != nil {
		return m.MtrlLotId
	}
	return ""
}

func (m *GetMconmtrlstReply_Mconmtrlst) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *GetMconmtrlstReply_Mconmtrlst) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *GetMconmtrlstReply_Mconmtrlst) GetUpdateAt() string {
	if m != nil {
		return m.UpdateAt
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMconmtrlstRequest)(nil), "proto.dm.mes.GetMconmtrlstRequest")
	proto.RegisterType((*GetMconmtrlstReply)(nil), "proto.dm.mes.GetMconmtrlstReply")
	proto.RegisterType((*GetMconmtrlstReply_Mconmtrlst)(nil), "proto.dm.mes.GetMconmtrlstReply.Mconmtrlst")
}

func init() { proto.RegisterFile("query_mconmtrlst.proto", fileDescriptor_18d9d7a06040131d) }

var fileDescriptor_18d9d7a06040131d = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x49, 0xdb, 0xef, 0xc3, 0x4e, 0x7a, 0x5a, 0x44, 0x82, 0x82, 0xd6, 0x9e, 0x0a, 0xc2,
	0x2e, 0xd4, 0x8b, 0x78, 0x11, 0x7b, 0x91, 0x82, 0x5e, 0x8a, 0x17, 0xbd, 0x84, 0x6d, 0x76, 0x0c,
	0xc1, 0xec, 0x4e, 0x9a, 0x4c, 0x94, 0xbd, 0xfa, 0x73, 0xfc, 0x95, 0x92, 0x2d, 0xd4, 0xa0, 0x57,
	0x4f, 0xcb, 0x3c, 0xef, 0x33, 0xcb, 0xcb, 0xc0, 0xd1, 0xb6, 0xc5, 0xda, 0xa7, 0x36, 0x23, 0x67,
	0xb9, 0x2e, 0x1b, 0x96, 0x55, 0x4d, 0x4c, 0x62, 0x12, 0x1e, 0x69, 0xac, 0xb4, 0xd8, 0xcc, 0x9e,
	0xe0, 0xf0, 0x0e, 0xf9, 0x61, 0x2f, 0xad, 0x71, 0xdb, 0x62, 0xc3, 0xe2, 0x1c, 0x26, 0x19, 0x39,
	0xd6, 0x85, 0xc3, 0x3a, 0x2d, 0x4c, 0x12, 0x4d, 0xa3, 0xf9, 0x78, 0x1d, 0xef, 0xd9, 0xca, 0x88,
	0x53, 0x88, 0xbb, 0x9d, 0xb4, 0x24, 0xee, 0x8c, 0x41, 0x30, 0xc6, 0x1d, 0xba, 0x27, 0x5e, 0x99,
	0xd9, 0xc7, 0x00, 0xc4, 0x8f, 0xbf, 0xab, 0xd2, 0x8b, 0x1b, 0x18, 0x15, 0xee, 0x85, 0x92, 0x68,
	0x3a, 0x9c, 0xc7, 0x8b, 0x0b, 0xd9, 0xaf, 0x23, 0x7f, 0xfb, 0xb2, 0x37, 0x87, 0xc5, 0xe3, 0xcf,
	0x08, 0xe0, 0x1b, 0xfe, 0x41, 0x53, 0x71, 0x06, 0xb1, 0xce, 0xb8, 0x20, 0x97, 0xb2, 0xaf, 0x30,
	0x19, 0x86, 0x1c, 0x76, 0xe8, 0xd1, 0x57, 0x28, 0x4e, 0x60, 0xdc, 0x56, 0x46, 0x33, 0xa6, 0x1b,
	0x9f, 0x8c, 0x42, 0x7c, 0xb0, 0x03, 0x4b, 0xdf, 0x0b, 0x35, 0x27, 0xff, 0xfa, 0xe1, 0x2d, 0x2f,
	0xaf, 0x9f, 0xaf, 0xf2, 0x82, 0x4b, 0xbd, 0x91, 0xaf, 0xe8, 0x8c, 0x96, 0x19, 0x59, 0xc9, 0xef,
	0x2a, 0x0c, 0x2a, 0x23, 0x6b, 0xc9, 0x35, 0xea, 0x6d, 0xa1, 0xc2, 0x15, 0x54, 0x4e, 0xa5, 0x76,
	0xb9, 0x32, 0x56, 0x59, 0x6c, 0x36, 0xff, 0x03, 0xbc, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x5f,
	0xad, 0xec, 0x1a, 0xca, 0x01, 0x00, 0x00,
}
