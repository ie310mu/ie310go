// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/frequency_cap_level.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	_ "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api/annotations"
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

// The level on which the cap is to be applied (e.g ad group ad, ad group).
// Cap is applied to all the resources of this level.
type FrequencyCapLevelEnum_FrequencyCapLevel int32

const (
	// Not specified.
	FrequencyCapLevelEnum_UNSPECIFIED FrequencyCapLevelEnum_FrequencyCapLevel = 0
	// Used for return value only. Represents value unknown in this version.
	FrequencyCapLevelEnum_UNKNOWN FrequencyCapLevelEnum_FrequencyCapLevel = 1
	// The cap is applied at the ad group ad level.
	FrequencyCapLevelEnum_AD_GROUP_AD FrequencyCapLevelEnum_FrequencyCapLevel = 2
	// The cap is applied at the ad group level.
	FrequencyCapLevelEnum_AD_GROUP FrequencyCapLevelEnum_FrequencyCapLevel = 3
	// The cap is applied at the campaign level.
	FrequencyCapLevelEnum_CAMPAIGN FrequencyCapLevelEnum_FrequencyCapLevel = 4
)

var FrequencyCapLevelEnum_FrequencyCapLevel_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_GROUP_AD",
	3: "AD_GROUP",
	4: "CAMPAIGN",
}

var FrequencyCapLevelEnum_FrequencyCapLevel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"AD_GROUP_AD": 2,
	"AD_GROUP":    3,
	"CAMPAIGN":    4,
}

func (x FrequencyCapLevelEnum_FrequencyCapLevel) String() string {
	return proto.EnumName(FrequencyCapLevelEnum_FrequencyCapLevel_name, int32(x))
}

func (FrequencyCapLevelEnum_FrequencyCapLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f3b8ac2167d6463c, []int{0, 0}
}

// Container for enum describing the level on which the cap is to be applied.
type FrequencyCapLevelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrequencyCapLevelEnum) Reset()         { *m = FrequencyCapLevelEnum{} }
func (m *FrequencyCapLevelEnum) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapLevelEnum) ProtoMessage()    {}
func (*FrequencyCapLevelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3b8ac2167d6463c, []int{0}
}

func (m *FrequencyCapLevelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapLevelEnum.Unmarshal(m, b)
}
func (m *FrequencyCapLevelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapLevelEnum.Marshal(b, m, deterministic)
}
func (m *FrequencyCapLevelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapLevelEnum.Merge(m, src)
}
func (m *FrequencyCapLevelEnum) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapLevelEnum.Size(m)
}
func (m *FrequencyCapLevelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapLevelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapLevelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.FrequencyCapLevelEnum_FrequencyCapLevel", FrequencyCapLevelEnum_FrequencyCapLevel_name, FrequencyCapLevelEnum_FrequencyCapLevel_value)
	proto.RegisterType((*FrequencyCapLevelEnum)(nil), "google.ads.googleads.v1.enums.FrequencyCapLevelEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/frequency_cap_level.proto", fileDescriptor_f3b8ac2167d6463c)
}

var fileDescriptor_f3b8ac2167d6463c = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4a, 0xfb, 0x30,
	0x1c, 0xfd, 0xaf, 0xfb, 0xa3, 0x92, 0x09, 0xd6, 0x82, 0x5e, 0x88, 0xbb, 0xd8, 0x1e, 0x20, 0xa1,
	0x78, 0x21, 0xc4, 0xab, 0xec, 0xab, 0x0c, 0xb5, 0x2b, 0xca, 0x26, 0x48, 0xb1, 0xc4, 0x35, 0x86,
	0x41, 0x97, 0xd4, 0xa5, 0xad, 0xf8, 0x3a, 0x5e, 0xfa, 0x28, 0x3e, 0x8a, 0xbe, 0x84, 0x24, 0xb1,
	0xbd, 0x19, 0x7a, 0x53, 0x4e, 0x7f, 0xe7, 0x83, 0x93, 0x03, 0xce, 0xb9, 0x94, 0x3c, 0x63, 0x88,
	0xa6, 0x0a, 0x59, 0xa8, 0x51, 0xe5, 0x23, 0x26, 0xca, 0xb5, 0x42, 0x4f, 0x1b, 0xf6, 0x5c, 0x32,
	0xb1, 0x7c, 0x4d, 0x96, 0x34, 0x4f, 0x32, 0x56, 0xb1, 0x0c, 0xe6, 0x1b, 0x59, 0x48, 0xaf, 0x6b,
	0xd5, 0x90, 0xa6, 0x0a, 0x36, 0x46, 0x58, 0xf9, 0xd0, 0x18, 0x4f, 0x4e, 0xeb, 0xdc, 0x7c, 0x85,
	0xa8, 0x10, 0xb2, 0xa0, 0xc5, 0x4a, 0x0a, 0x65, 0xcd, 0xfd, 0x17, 0x70, 0x34, 0xa9, 0x93, 0x87,
	0x34, 0xbf, 0xd2, 0xb9, 0x63, 0x51, 0xae, 0xfb, 0x0f, 0xe0, 0x70, 0x8b, 0xf0, 0x0e, 0x40, 0x67,
	0x1e, 0xde, 0x46, 0xe3, 0xe1, 0x74, 0x32, 0x1d, 0x8f, 0xdc, 0x7f, 0x5e, 0x07, 0xec, 0xce, 0xc3,
	0xcb, 0x70, 0x76, 0x17, 0xba, 0x2d, 0xcd, 0x92, 0x51, 0x12, 0xdc, 0xcc, 0xe6, 0x51, 0x42, 0x46,
	0xae, 0xe3, 0xed, 0x83, 0xbd, 0xfa, 0xe0, 0xb6, 0xf5, 0xdf, 0x90, 0x5c, 0x47, 0x64, 0x1a, 0x84,
	0xee, 0xff, 0xc1, 0x57, 0x0b, 0xf4, 0x96, 0x72, 0x0d, 0xff, 0x2c, 0x3f, 0x38, 0xde, 0xea, 0x10,
	0xe9, 0xda, 0x51, 0xeb, 0x7e, 0xf0, 0x63, 0xe4, 0x32, 0xa3, 0x82, 0x43, 0xb9, 0xe1, 0x88, 0x33,
	0x61, 0x1e, 0x55, 0xcf, 0x97, 0xaf, 0xd4, 0x2f, 0x6b, 0x5e, 0x98, 0xef, 0x9b, 0xd3, 0x0e, 0x08,
	0x79, 0x77, 0xba, 0x81, 0x8d, 0x22, 0xa9, 0x82, 0x16, 0x6a, 0xb4, 0xf0, 0xa1, 0x1e, 0x42, 0x7d,
	0xd4, 0x7c, 0x4c, 0x52, 0x15, 0x37, 0x7c, 0xbc, 0xf0, 0x63, 0xc3, 0x7f, 0x3a, 0x3d, 0x7b, 0xc4,
	0x98, 0xa4, 0x0a, 0xe3, 0x46, 0x81, 0xf1, 0xc2, 0xc7, 0xd8, 0x68, 0x1e, 0x77, 0x4c, 0xb1, 0xb3,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xa7, 0xd0, 0xfb, 0xe5, 0x01, 0x00, 0x00,
}
