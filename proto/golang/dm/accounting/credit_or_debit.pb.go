// Code generated by protoc-gen-go. DO NOT EDIT.
// source: credit_or_debit.proto

package accounting

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

//借貸別Cod
type CreditOrDebit int32

const (
	CreditOrDebit_CREDIT_OR_DEBIT_UNSPECIFIED CreditOrDebit = 0
	//借方-Debit
	CreditOrDebit_DEBIT CreditOrDebit = 1
	//貸方-Credit
	CreditOrDebit_CREDIT CreditOrDebit = 2
)

var CreditOrDebit_name = map[int32]string{
	0: "CREDIT_OR_DEBIT_UNSPECIFIED",
	1: "DEBIT",
	2: "CREDIT",
}

var CreditOrDebit_value = map[string]int32{
	"CREDIT_OR_DEBIT_UNSPECIFIED": 0,
	"DEBIT":                       1,
	"CREDIT":                      2,
}

func (x CreditOrDebit) String() string {
	return proto.EnumName(CreditOrDebit_name, int32(x))
}

func (CreditOrDebit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_48d5a29b3c6c92d8, []int{0}
}

func init() {
	proto.RegisterEnum("proto.dm.accounting.CreditOrDebit", CreditOrDebit_name, CreditOrDebit_value)
}

func init() { proto.RegisterFile("credit_or_debit.proto", fileDescriptor_48d5a29b3c6c92d8) }

var fileDescriptor_48d5a29b3c6c92d8 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2e, 0x4a, 0x4d,
	0xc9, 0x2c, 0x89, 0xcf, 0x2f, 0x8a, 0x4f, 0x49, 0x4d, 0xca, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x06, 0x53, 0x7a, 0x29, 0xb9, 0x7a, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25,
	0x99, 0x79, 0xe9, 0x5a, 0xee, 0x5c, 0xbc, 0xce, 0x60, 0xd5, 0xfe, 0x45, 0x2e, 0x20, 0xb5, 0x42,
	0xf2, 0x5c, 0xd2, 0xce, 0x41, 0xae, 0x2e, 0x9e, 0x21, 0xf1, 0xfe, 0x41, 0xf1, 0x2e, 0xae, 0x4e,
	0x9e, 0x21, 0xf1, 0xa1, 0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x0c,
	0x42, 0x9c, 0x5c, 0xac, 0x60, 0x61, 0x01, 0x46, 0x21, 0x2e, 0x2e, 0x36, 0x88, 0x5a, 0x01, 0x26,
	0x27, 0xc7, 0x28, 0xfb, 0xf4, 0xcc, 0x92, 0x9c, 0xc4, 0x24, 0xbd, 0xec, 0xd4, 0xbc, 0x94, 0x44,
	0xbd, 0xe4, 0xfc, 0x5c, 0xbd, 0x92, 0x72, 0x7d, 0x30, 0x47, 0x3f, 0x39, 0x3f, 0x37, 0x37, 0x3f,
	0xaf, 0x58, 0xbf, 0xcc, 0x48, 0x1f, 0xec, 0x08, 0xfd, 0xf4, 0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0xfd,
	0x94, 0x5c, 0x7d, 0x84, 0x5b, 0x92, 0xd8, 0xc0, 0x72, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0x25, 0x03, 0x7a, 0xc0, 0x00, 0x00, 0x00,
}