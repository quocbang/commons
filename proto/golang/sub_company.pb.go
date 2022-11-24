// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sub_company.proto

package golang

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

//集團廠別子公司SubCompany
type SubCompany int32

const (
	SubCompany_SUB_COMPANY_UNSPECIFIED SubCompany = 0
	//生產廠
	//KD
	SubCompany_KD SubCompany = 100
	//KS
	SubCompany_KS SubCompany = 101
	//KC
	SubCompany_KC SubCompany = 102
	//KV
	SubCompany_KV SubCompany = 103
	//KT
	SubCompany_KT SubCompany = 104
	//KI
	SubCompany_KI SubCompany = 105
	//KH
	SubCompany_KH SubCompany = 106
	//KZ
	SubCompany_KZ SubCompany = 107
	//CKT
	SubCompany_CKT SubCompany = 108
	//KJ
	SubCompany_KJ SubCompany = 109
	//境外公司
	//KIC
	SubCompany_KIC SubCompany = 301
	//KGH
	SubCompany_KGH SubCompany = 302
	//KGI
	SubCompany_KGI SubCompany = 303
	//KGCI
	SubCompany_KGCI SubCompany = 304
	//KHK
	SubCompany_KHK SubCompany = 305
	//KCE
	SubCompany_KCE SubCompany = 306
	//銷售公司
	//KF
	SubCompany_KF SubCompany = 501
	//KA
	SubCompany_KA SubCompany = 502
	//ADI
	SubCompany_ADI SubCompany = 503
	//KE
	SubCompany_KE SubCompany = 504
	//STARCO
	SubCompany_STARCO SubCompany = 505
)

var SubCompany_name = map[int32]string{
	0:   "SUB_COMPANY_UNSPECIFIED",
	100: "KD",
	101: "KS",
	102: "KC",
	103: "KV",
	104: "KT",
	105: "KI",
	106: "KH",
	107: "KZ",
	108: "CKT",
	109: "KJ",
	301: "KIC",
	302: "KGH",
	303: "KGI",
	304: "KGCI",
	305: "KHK",
	306: "KCE",
	501: "KF",
	502: "KA",
	503: "ADI",
	504: "KE",
	505: "STARCO",
}

var SubCompany_value = map[string]int32{
	"SUB_COMPANY_UNSPECIFIED": 0,
	"KD":                      100,
	"KS":                      101,
	"KC":                      102,
	"KV":                      103,
	"KT":                      104,
	"KI":                      105,
	"KH":                      106,
	"KZ":                      107,
	"CKT":                     108,
	"KJ":                      109,
	"KIC":                     301,
	"KGH":                     302,
	"KGI":                     303,
	"KGCI":                    304,
	"KHK":                     305,
	"KCE":                     306,
	"KF":                      501,
	"KA":                      502,
	"ADI":                     503,
	"KE":                      504,
	"STARCO":                  505,
}

func (x SubCompany) String() string {
	return proto.EnumName(SubCompany_name, int32(x))
}

func (SubCompany) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_741749f4976fe480, []int{0}
}

func init() {
	proto.RegisterEnum("proto.SubCompany", SubCompany_name, SubCompany_value)
}

func init() { proto.RegisterFile("sub_company.proto", fileDescriptor_741749f4976fe480) }

var fileDescriptor_741749f4976fe480 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xd0, 0xcd, 0x4e, 0x83, 0x40,
	0x14, 0x05, 0x60, 0x65, 0xb4, 0xd5, 0x71, 0x73, 0x9c, 0x8d, 0x0b, 0xdf, 0xc0, 0x05, 0x44, 0xfb,
	0x04, 0x74, 0xa0, 0x65, 0x9c, 0xd8, 0x36, 0x42, 0x4d, 0xec, 0x86, 0x00, 0xad, 0x58, 0x0b, 0x4c,
	0x63, 0xa9, 0xc6, 0x87, 0xf2, 0xf7, 0xad, 0x7c, 0x00, 0xff, 0x76, 0xc6, 0x8b, 0xab, 0xef, 0xdc,
	0x73, 0x76, 0x97, 0xef, 0xaf, 0xd6, 0x69, 0x9c, 0x99, 0x72, 0x99, 0x54, 0x0f, 0xf6, 0xf2, 0xd6,
	0xd4, 0x46, 0x6c, 0x13, 0x47, 0xef, 0x9b, 0x9c, 0x87, 0xeb, 0x54, 0x36, 0x9b, 0x38, 0xe4, 0x07,
	0xe1, 0xb8, 0x1b, 0xcb, 0xe1, 0xd9, 0xc8, 0x1d, 0x5c, 0xc6, 0xe3, 0x41, 0x38, 0xf2, 0xa5, 0xea,
	0x29, 0xdf, 0xc3, 0x86, 0x68, 0x71, 0x4b, 0x7b, 0x98, 0x92, 0x21, 0x66, 0xa4, 0xc4, 0x15, 0x79,
	0x81, 0x9c, 0x8c, 0x70, 0x4d, 0x2a, 0xcc, 0xc9, 0x00, 0x37, 0xe4, 0x04, 0x0b, 0xd1, 0xe6, 0x4c,
	0xea, 0x08, 0x05, 0x15, 0xa7, 0x28, 0xc5, 0x0e, 0x67, 0x5a, 0x49, 0x3c, 0x5a, 0x94, 0xfa, 0x01,
	0x9e, 0xfe, 0x93, 0xc2, 0xb3, 0x25, 0x76, 0xf9, 0x96, 0xee, 0x4b, 0x85, 0x97, 0xa6, 0x0c, 0x34,
	0x5e, 0x9b, 0x24, 0x7d, 0xbc, 0x59, 0xa2, 0xcd, 0x2d, 0xdd, 0xc3, 0x07, 0xa3, 0xe0, 0xe2, 0x93,
	0xfd, 0x6d, 0xae, 0xa7, 0xf0, 0xd5, 0x54, 0x3e, 0xbe, 0x99, 0xd8, 0xe3, 0xad, 0x30, 0x72, 0xcf,
	0xe5, 0x10, 0x3f, 0xac, 0xdb, 0x99, 0x1c, 0xe7, 0xf3, 0xba, 0x48, 0x52, 0x7b, 0x31, 0xab, 0xa6,
	0x89, 0x9d, 0x99, 0xd2, 0xae, 0xef, 0x1d, 0x3a, 0x9c, 0xcc, 0x94, 0xa5, 0xa9, 0x56, 0xce, 0xdd,
	0x89, 0x43, 0x9f, 0x71, 0x72, 0x53, 0x24, 0x55, 0x9e, 0xb6, 0xe8, 0xea, 0xfc, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x6a, 0x41, 0x9f, 0x43, 0x01, 0x00, 0x00,
}