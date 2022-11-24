// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recipe/recipe.proto

package recipe

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commons "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/v2/commons"
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

type DecimalParam struct {
	Value                *commons.Decimal `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Min                  *commons.Decimal `protobuf:"bytes,2,opt,name=min,proto3" json:"min,omitempty"`
	Max                  *commons.Decimal `protobuf:"bytes,3,opt,name=max,proto3" json:"max,omitempty"`
	Unit                 string           `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DecimalParam) Reset()         { *m = DecimalParam{} }
func (m *DecimalParam) String() string { return proto.CompactTextString(m) }
func (*DecimalParam) ProtoMessage()    {}
func (*DecimalParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{0}
}

func (m *DecimalParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecimalParam.Unmarshal(m, b)
}
func (m *DecimalParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecimalParam.Marshal(b, m, deterministic)
}
func (m *DecimalParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecimalParam.Merge(m, src)
}
func (m *DecimalParam) XXX_Size() int {
	return xxx_messageInfo_DecimalParam.Size(m)
}
func (m *DecimalParam) XXX_DiscardUnknown() {
	xxx_messageInfo_DecimalParam.DiscardUnknown(m)
}

var xxx_messageInfo_DecimalParam proto.InternalMessageInfo

func (m *DecimalParam) GetValue() *commons.Decimal {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DecimalParam) GetMin() *commons.Decimal {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *DecimalParam) GetMax() *commons.Decimal {
	if m != nil {
		return m.Max
	}
	return nil
}

func (m *DecimalParam) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type Freshness struct {
	Resting              *commons.Time    `protobuf:"bytes,1,opt,name=resting,proto3" json:"resting,omitempty"`
	Expiry               *commons.Time    `protobuf:"bytes,2,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Extension            *commons.Decimal `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	MaxExtensionTimes    uint32           `protobuf:"varint,4,opt,name=maxExtensionTimes,proto3" json:"maxExtensionTimes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Freshness) Reset()         { *m = Freshness{} }
func (m *Freshness) String() string { return proto.CompactTextString(m) }
func (*Freshness) ProtoMessage()    {}
func (*Freshness) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{1}
}

func (m *Freshness) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Freshness.Unmarshal(m, b)
}
func (m *Freshness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Freshness.Marshal(b, m, deterministic)
}
func (m *Freshness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Freshness.Merge(m, src)
}
func (m *Freshness) XXX_Size() int {
	return xxx_messageInfo_Freshness.Size(m)
}
func (m *Freshness) XXX_DiscardUnknown() {
	xxx_messageInfo_Freshness.DiscardUnknown(m)
}

var xxx_messageInfo_Freshness proto.InternalMessageInfo

func (m *Freshness) GetResting() *commons.Time {
	if m != nil {
		return m.Resting
	}
	return nil
}

func (m *Freshness) GetExpiry() *commons.Time {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func (m *Freshness) GetExtension() *commons.Decimal {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *Freshness) GetMaxExtensionTimes() uint32 {
	if m != nil {
		return m.MaxExtensionTimes
	}
	return 0
}

type ResourceRequirement struct {
	Name                 string                               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Quantity             *DecimalParam                        `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Grade                string                               `protobuf:"bytes,3,opt,name=grade,proto3" json:"grade,omitempty"`
	Freshness            *Freshness                           `protobuf:"bytes,4,opt,name=freshness,proto3" json:"freshness,omitempty"`
	Site                 *ResourceRequirement_SiteRequirement `protobuf:"bytes,5,opt,name=site,proto3" json:"site,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ResourceRequirement) Reset()         { *m = ResourceRequirement{} }
func (m *ResourceRequirement) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirement) ProtoMessage()    {}
func (*ResourceRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{2}
}

func (m *ResourceRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirement.Unmarshal(m, b)
}
func (m *ResourceRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirement.Marshal(b, m, deterministic)
}
func (m *ResourceRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirement.Merge(m, src)
}
func (m *ResourceRequirement) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirement.Size(m)
}
func (m *ResourceRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirement proto.InternalMessageInfo

func (m *ResourceRequirement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceRequirement) GetQuantity() *DecimalParam {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *ResourceRequirement) GetGrade() string {
	if m != nil {
		return m.Grade
	}
	return ""
}

func (m *ResourceRequirement) GetFreshness() *Freshness {
	if m != nil {
		return m.Freshness
	}
	return nil
}

func (m *ResourceRequirement) GetSite() *ResourceRequirement_SiteRequirement {
	if m != nil {
		return m.Site
	}
	return nil
}

type ResourceRequirement_SiteRequirement struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Requirement:
	//	*ResourceRequirement_SiteRequirement_Unique
	//	*ResourceRequirement_SiteRequirement_Any
	//	*ResourceRequirement_SiteRequirement_Index
	Requirement          isResourceRequirement_SiteRequirement_Requirement `protobuf_oneof:"requirement"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *ResourceRequirement_SiteRequirement) Reset()         { *m = ResourceRequirement_SiteRequirement{} }
func (m *ResourceRequirement_SiteRequirement) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirement_SiteRequirement) ProtoMessage()    {}
func (*ResourceRequirement_SiteRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{2, 0}
}

func (m *ResourceRequirement_SiteRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirement_SiteRequirement.Unmarshal(m, b)
}
func (m *ResourceRequirement_SiteRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirement_SiteRequirement.Marshal(b, m, deterministic)
}
func (m *ResourceRequirement_SiteRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirement_SiteRequirement.Merge(m, src)
}
func (m *ResourceRequirement_SiteRequirement) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirement_SiteRequirement.Size(m)
}
func (m *ResourceRequirement_SiteRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirement_SiteRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirement_SiteRequirement proto.InternalMessageInfo

func (m *ResourceRequirement_SiteRequirement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isResourceRequirement_SiteRequirement_Requirement interface {
	isResourceRequirement_SiteRequirement_Requirement()
}

type ResourceRequirement_SiteRequirement_Unique struct {
	Unique bool `protobuf:"varint,2,opt,name=unique,proto3,oneof"`
}

type ResourceRequirement_SiteRequirement_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type ResourceRequirement_SiteRequirement_Index struct {
	Index uint32 `protobuf:"varint,4,opt,name=index,proto3,oneof"`
}

func (*ResourceRequirement_SiteRequirement_Unique) isResourceRequirement_SiteRequirement_Requirement() {
}

func (*ResourceRequirement_SiteRequirement_Any) isResourceRequirement_SiteRequirement_Requirement() {}

func (*ResourceRequirement_SiteRequirement_Index) isResourceRequirement_SiteRequirement_Requirement() {
}

func (m *ResourceRequirement_SiteRequirement) GetRequirement() isResourceRequirement_SiteRequirement_Requirement {
	if m != nil {
		return m.Requirement
	}
	return nil
}

func (m *ResourceRequirement_SiteRequirement) GetUnique() bool {
	if x, ok := m.GetRequirement().(*ResourceRequirement_SiteRequirement_Unique); ok {
		return x.Unique
	}
	return false
}

func (m *ResourceRequirement_SiteRequirement) GetAny() bool {
	if x, ok := m.GetRequirement().(*ResourceRequirement_SiteRequirement_Any); ok {
		return x.Any
	}
	return false
}

func (m *ResourceRequirement_SiteRequirement) GetIndex() uint32 {
	if x, ok := m.GetRequirement().(*ResourceRequirement_SiteRequirement_Index); ok {
		return x.Index
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResourceRequirement_SiteRequirement) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResourceRequirement_SiteRequirement_Unique)(nil),
		(*ResourceRequirement_SiteRequirement_Any)(nil),
		(*ResourceRequirement_SiteRequirement_Index)(nil),
	}
}

type ResourceRequirements struct {
	Requirements         []*ResourceRequirement `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{3}
}

func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (m *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(m, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetRequirements() []*ResourceRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

type Params struct {
	Params               []*Params_Param `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{4}
}

func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetParams() []*Params_Param {
	if m != nil {
		return m.Params
	}
	return nil
}

type Params_Param struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DecimalValue         *DecimalParam `protobuf:"bytes,2,opt,name=decimalValue,proto3" json:"decimalValue,omitempty"`
	StringValue          string        `protobuf:"bytes,3,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Params_Param) Reset()         { *m = Params_Param{} }
func (m *Params_Param) String() string { return proto.CompactTextString(m) }
func (*Params_Param) ProtoMessage()    {}
func (*Params_Param) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{4, 0}
}

func (m *Params_Param) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params_Param.Unmarshal(m, b)
}
func (m *Params_Param) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params_Param.Marshal(b, m, deterministic)
}
func (m *Params_Param) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params_Param.Merge(m, src)
}
func (m *Params_Param) XXX_Size() int {
	return xxx_messageInfo_Params_Param.Size(m)
}
func (m *Params_Param) XXX_DiscardUnknown() {
	xxx_messageInfo_Params_Param.DiscardUnknown(m)
}

var xxx_messageInfo_Params_Param proto.InternalMessageInfo

func (m *Params_Param) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Params_Param) GetDecimalValue() *DecimalParam {
	if m != nil {
		return m.DecimalValue
	}
	return nil
}

func (m *Params_Param) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

type RecipeDetails struct {
	GroupedResources     map[string]*ResourceRequirements `protobuf:"bytes,1,rep,name=groupedResources,proto3" json:"groupedResources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	GroupedParams        map[string]*Params               `protobuf:"bytes,2,rep,name=groupedParams,proto3" json:"groupedParams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RecipeDetails) Reset()         { *m = RecipeDetails{} }
func (m *RecipeDetails) String() string { return proto.CompactTextString(m) }
func (*RecipeDetails) ProtoMessage()    {}
func (*RecipeDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{5}
}

func (m *RecipeDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecipeDetails.Unmarshal(m, b)
}
func (m *RecipeDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecipeDetails.Marshal(b, m, deterministic)
}
func (m *RecipeDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecipeDetails.Merge(m, src)
}
func (m *RecipeDetails) XXX_Size() int {
	return xxx_messageInfo_RecipeDetails.Size(m)
}
func (m *RecipeDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_RecipeDetails.DiscardUnknown(m)
}

var xxx_messageInfo_RecipeDetails proto.InternalMessageInfo

func (m *RecipeDetails) GetGroupedResources() map[string]*ResourceRequirements {
	if m != nil {
		return m.GroupedResources
	}
	return nil
}

func (m *RecipeDetails) GetGroupedParams() map[string]*Params {
	if m != nil {
		return m.GroupedParams
	}
	return nil
}

type Recipe struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string           `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	BatchSize            *commons.Decimal `protobuf:"bytes,5,opt,name=batchSize,proto3" json:"batchSize,omitempty"`
	ReleasedAt           *commons.Time    `protobuf:"bytes,4,opt,name=releasedAt,proto3" json:"releasedAt,omitempty"`
	Details              *RecipeDetails   `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Recipe) Reset()         { *m = Recipe{} }
func (m *Recipe) String() string { return proto.CompactTextString(m) }
func (*Recipe) ProtoMessage()    {}
func (*Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_1098dce7f2864922, []int{6}
}

func (m *Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipe.Unmarshal(m, b)
}
func (m *Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipe.Marshal(b, m, deterministic)
}
func (m *Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipe.Merge(m, src)
}
func (m *Recipe) XXX_Size() int {
	return xxx_messageInfo_Recipe.Size(m)
}
func (m *Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_Recipe proto.InternalMessageInfo

func (m *Recipe) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Recipe) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Recipe) GetBatchSize() *commons.Decimal {
	if m != nil {
		return m.BatchSize
	}
	return nil
}

func (m *Recipe) GetReleasedAt() *commons.Time {
	if m != nil {
		return m.ReleasedAt
	}
	return nil
}

func (m *Recipe) GetDetails() *RecipeDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*DecimalParam)(nil), "proto.mes.v2.recipe.DecimalParam")
	proto.RegisterType((*Freshness)(nil), "proto.mes.v2.recipe.Freshness")
	proto.RegisterType((*ResourceRequirement)(nil), "proto.mes.v2.recipe.ResourceRequirement")
	proto.RegisterType((*ResourceRequirement_SiteRequirement)(nil), "proto.mes.v2.recipe.ResourceRequirement.SiteRequirement")
	proto.RegisterType((*ResourceRequirements)(nil), "proto.mes.v2.recipe.ResourceRequirements")
	proto.RegisterType((*Params)(nil), "proto.mes.v2.recipe.Params")
	proto.RegisterType((*Params_Param)(nil), "proto.mes.v2.recipe.Params.Param")
	proto.RegisterType((*RecipeDetails)(nil), "proto.mes.v2.recipe.RecipeDetails")
	proto.RegisterMapType((map[string]*Params)(nil), "proto.mes.v2.recipe.RecipeDetails.GroupedParamsEntry")
	proto.RegisterMapType((map[string]*ResourceRequirements)(nil), "proto.mes.v2.recipe.RecipeDetails.GroupedResourcesEntry")
	proto.RegisterType((*Recipe)(nil), "proto.mes.v2.recipe.Recipe")
}

func init() { proto.RegisterFile("recipe/recipe.proto", fileDescriptor_1098dce7f2864922) }

var fileDescriptor_1098dce7f2864922 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x4e, 0x1b, 0x3b,
	0x14, 0x66, 0x12, 0x12, 0xc8, 0x09, 0xd1, 0xe5, 0x1a, 0xee, 0xd5, 0x28, 0x57, 0xb7, 0x4a, 0x67,
	0x95, 0x4a, 0xd5, 0x44, 0x0d, 0xad, 0x44, 0x29, 0x15, 0x02, 0x41, 0xcb, 0xa2, 0x8b, 0xca, 0x54,
	0x5d, 0xb4, 0xea, 0xc2, 0x64, 0x4e, 0x83, 0x45, 0xc6, 0x13, 0x6c, 0x4f, 0x9a, 0xb0, 0xae, 0xfa,
	0x1e, 0x7d, 0x8c, 0xbe, 0x42, 0x9f, 0xa2, 0x2f, 0xd0, 0x77, 0xa8, 0xc6, 0x9e, 0x09, 0x13, 0x18,
	0x60, 0x56, 0xf6, 0xb1, 0xcf, 0xf7, 0x9d, 0x5f, 0x1f, 0xc3, 0x86, 0xc4, 0x01, 0x1f, 0x63, 0xcf,
	0x2e, 0xfe, 0x58, 0x46, 0x3a, 0x22, 0x1b, 0x66, 0xf1, 0x43, 0x54, 0xfe, 0xa4, 0xef, 0xdb, 0xab,
	0xb6, 0x67, 0x0e, 0x7b, 0x21, 0xaa, 0xde, 0xa4, 0xdf, 0x1b, 0x44, 0x61, 0x18, 0x09, 0x95, 0xad,
	0x16, 0xe8, 0xfd, 0x70, 0x60, 0xed, 0x10, 0x07, 0x3c, 0x64, 0xa3, 0xb7, 0x4c, 0xb2, 0x90, 0x6c,
	0x41, 0x6d, 0xc2, 0x46, 0x31, 0xba, 0x4e, 0xc7, 0xe9, 0x36, 0xfb, 0xff, 0xfb, 0x0b, 0xcc, 0x19,
	0x38, 0x85, 0x50, 0xab, 0x4b, 0x7a, 0x50, 0x0d, 0xb9, 0x70, 0x2b, 0x65, 0x20, 0x89, 0xa6, 0x01,
	0xb0, 0xa9, 0x5b, 0x2d, 0x07, 0x60, 0x53, 0x42, 0x60, 0x39, 0x16, 0x5c, 0xbb, 0xcb, 0x1d, 0xa7,
	0xdb, 0xa0, 0x66, 0xef, 0xfd, 0x72, 0xa0, 0xf1, 0x4a, 0xa2, 0x3a, 0x13, 0xa8, 0x14, 0x79, 0x0a,
	0x2b, 0x12, 0x95, 0xe6, 0x62, 0x98, 0xba, 0xde, 0x2e, 0xa6, 0x7d, 0xc7, 0x43, 0xa4, 0x99, 0x2a,
	0xe9, 0x43, 0x1d, 0xa7, 0x63, 0x2e, 0x67, 0xa9, 0xf3, 0x77, 0x81, 0x52, 0x4d, 0xf2, 0x02, 0x1a,
	0x38, 0xd5, 0x28, 0x14, 0x8f, 0x44, 0xb9, 0x10, 0xae, 0xf4, 0xc9, 0x63, 0xf8, 0x3b, 0x64, 0xd3,
	0xa3, 0x4c, 0x4e, 0x88, 0x95, 0x89, 0xaa, 0x45, 0x6f, 0x5e, 0x78, 0xdf, 0xaa, 0xb0, 0x41, 0x51,
	0x45, 0xb1, 0x1c, 0x20, 0xc5, 0x8b, 0x98, 0x4b, 0x0c, 0x51, 0xe8, 0x24, 0x1d, 0x82, 0x85, 0xb6,
	0x48, 0x0d, 0x6a, 0xf6, 0xe4, 0x25, 0xac, 0x5e, 0xc4, 0x4c, 0x68, 0xae, 0xb3, 0x60, 0x1e, 0xfa,
	0x05, 0x6d, 0xe1, 0xe7, 0xcb, 0x4d, 0xe7, 0x10, 0xb2, 0x09, 0xb5, 0xa1, 0x64, 0x01, 0x9a, 0x88,
	0x1a, 0xd4, 0x0a, 0x64, 0x17, 0x1a, 0x9f, 0xb3, 0x14, 0x1b, 0x37, 0x9b, 0xfd, 0x07, 0x85, 0xac,
	0xf3, 0x42, 0xd0, 0x2b, 0x00, 0x79, 0x03, 0xcb, 0x8a, 0x6b, 0x74, 0x6b, 0x06, 0xb8, 0x5d, 0x08,
	0x2c, 0x08, 0xcf, 0x3f, 0xe1, 0x3a, 0x2f, 0x53, 0xc3, 0xd2, 0xbe, 0x84, 0xbf, 0xae, 0x5d, 0x14,
	0xe6, 0xc1, 0x85, 0x7a, 0x2c, 0xf8, 0x45, 0x8c, 0x26, 0x0b, 0xab, 0xc7, 0x4b, 0x34, 0x95, 0x09,
	0x81, 0x2a, 0x13, 0x33, 0x13, 0x60, 0x72, 0x9c, 0x08, 0xe4, 0x5f, 0xa8, 0x71, 0x11, 0xe0, 0xd4,
	0xd6, 0xe0, 0x78, 0x89, 0x5a, 0xf1, 0xa0, 0x05, 0x4d, 0x79, 0x65, 0xc8, 0x0b, 0x60, 0xb3, 0xc0,
	0xd1, 0x24, 0xc2, 0xb5, 0x9c, 0x9a, 0x72, 0x9d, 0x4e, 0xb5, 0xdb, 0xec, 0x77, 0xcb, 0x46, 0x4a,
	0x17, 0xd0, 0xde, 0x4f, 0x07, 0xea, 0xa6, 0x2e, 0x8a, 0x3c, 0x87, 0xfa, 0xd8, 0xec, 0x52, 0xca,
	0xe2, 0x5a, 0x5a, 0x65, 0xbb, 0xd0, 0x14, 0xd0, 0xfe, 0xea, 0x40, 0xcd, 0x3e, 0xe6, 0xa2, 0xf4,
	0x1c, 0xc1, 0x5a, 0x60, 0x3b, 0xe0, 0xbd, 0x79, 0xe7, 0xa5, 0x5b, 0x65, 0x01, 0x46, 0x3a, 0xd0,
	0x54, 0x5a, 0x72, 0x31, 0xb4, 0x2c, 0xb6, 0x69, 0xf2, 0x47, 0xde, 0xf7, 0x2a, 0xb4, 0xa8, 0xe1,
	0x39, 0x44, 0xcd, 0xf8, 0x48, 0x91, 0x00, 0xd6, 0x87, 0x32, 0x8a, 0xc7, 0x18, 0x64, 0xa9, 0xc8,
	0xa2, 0xbb, 0xad, 0x35, 0x72, 0x68, 0xff, 0xf5, 0x35, 0xe8, 0x91, 0xd0, 0x72, 0x46, 0x6f, 0x30,
	0x92, 0x8f, 0xd0, 0x4a, 0xcf, 0x6c, 0x76, 0xdc, 0x8a, 0x31, 0xf1, 0xac, 0xbc, 0x09, 0x8b, 0xb3,
	0xfc, 0x8b, 0x5c, 0x6d, 0x01, 0xff, 0x14, 0xfa, 0x41, 0xd6, 0xa1, 0x7a, 0x8e, 0xb3, 0x34, 0xd3,
	0xc9, 0x96, 0xec, 0x65, 0x93, 0xd4, 0x66, 0xf8, 0x51, 0xd9, 0x9e, 0x50, 0xe9, 0x54, 0xdd, 0xa9,
	0x6c, 0x3b, 0xed, 0x4f, 0x40, 0x6e, 0x3a, 0x55, 0x60, 0xec, 0xc9, 0xa2, 0xb1, 0xff, 0xee, 0xe8,
	0x96, 0x1c, 0xbd, 0xf7, 0xdb, 0x81, 0xba, 0x4d, 0xc1, 0x2d, 0x4f, 0x69, 0x65, 0x82, 0xd2, 0xcc,
	0xb9, 0x8a, 0x39, 0xce, 0xc4, 0x64, 0x06, 0x9e, 0x32, 0x3d, 0x38, 0x3b, 0xe1, 0x97, 0xd9, 0xf3,
	0xbe, 0x6f, 0x06, 0xce, 0xf5, 0xc9, 0x0e, 0x80, 0xc4, 0x11, 0x32, 0x85, 0xc1, 0xbe, 0x4e, 0xa7,
	0xca, 0x5d, 0x83, 0x37, 0xa7, 0x4d, 0x76, 0x61, 0x25, 0xb0, 0xd5, 0x4a, 0x47, 0xaf, 0x77, 0x7f,
	0x5d, 0x69, 0x06, 0x39, 0xd8, 0xff, 0xb0, 0x37, 0xe4, 0x7a, 0xc4, 0x4e, 0xfd, 0x73, 0x14, 0x01,
	0x4b, 0xcc, 0xf8, 0xfa, 0x4b, 0xcf, 0x08, 0xf3, 0x1f, 0x72, 0xd2, 0xef, 0xd9, 0x9f, 0x73, 0x18,
	0x8d, 0x98, 0x18, 0x66, 0x1f, 0xa8, 0x65, 0x3e, 0xad, 0x9b, 0xbb, 0xad, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xd8, 0xea, 0xa6, 0xa7, 0x88, 0x07, 0x00, 0x00,
}
