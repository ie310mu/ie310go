// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/advertising_channel_type.proto

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

// Enum describing the various advertising channel types.
type AdvertisingChannelTypeEnum_AdvertisingChannelType int32

const (
	// Not specified.
	AdvertisingChannelTypeEnum_UNSPECIFIED AdvertisingChannelTypeEnum_AdvertisingChannelType = 0
	// Used for return value only. Represents value unknown in this version.
	AdvertisingChannelTypeEnum_UNKNOWN AdvertisingChannelTypeEnum_AdvertisingChannelType = 1
	// Search Network. Includes display bundled, and Search+ campaigns.
	AdvertisingChannelTypeEnum_SEARCH AdvertisingChannelTypeEnum_AdvertisingChannelType = 2
	// Google Display Network only.
	AdvertisingChannelTypeEnum_DISPLAY AdvertisingChannelTypeEnum_AdvertisingChannelType = 3
	// Shopping campaigns serve on the shopping property
	// and on google.com search results.
	AdvertisingChannelTypeEnum_SHOPPING AdvertisingChannelTypeEnum_AdvertisingChannelType = 4
	// Hotel Ads campaigns.
	AdvertisingChannelTypeEnum_HOTEL AdvertisingChannelTypeEnum_AdvertisingChannelType = 5
	// Video campaigns.
	AdvertisingChannelTypeEnum_VIDEO AdvertisingChannelTypeEnum_AdvertisingChannelType = 6
	// App Campaigns, and App Campaigns for Engagement, that run
	// across multiple channels.
	AdvertisingChannelTypeEnum_MULTI_CHANNEL AdvertisingChannelTypeEnum_AdvertisingChannelType = 7
)

var AdvertisingChannelTypeEnum_AdvertisingChannelType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SEARCH",
	3: "DISPLAY",
	4: "SHOPPING",
	5: "HOTEL",
	6: "VIDEO",
	7: "MULTI_CHANNEL",
}

var AdvertisingChannelTypeEnum_AdvertisingChannelType_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"SEARCH":        2,
	"DISPLAY":       3,
	"SHOPPING":      4,
	"HOTEL":         5,
	"VIDEO":         6,
	"MULTI_CHANNEL": 7,
}

func (x AdvertisingChannelTypeEnum_AdvertisingChannelType) String() string {
	return proto.EnumName(AdvertisingChannelTypeEnum_AdvertisingChannelType_name, int32(x))
}

func (AdvertisingChannelTypeEnum_AdvertisingChannelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d496394fd88a2a26, []int{0, 0}
}

// The channel type a campaign may target to serve on.
type AdvertisingChannelTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdvertisingChannelTypeEnum) Reset()         { *m = AdvertisingChannelTypeEnum{} }
func (m *AdvertisingChannelTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdvertisingChannelTypeEnum) ProtoMessage()    {}
func (*AdvertisingChannelTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d496394fd88a2a26, []int{0}
}

func (m *AdvertisingChannelTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdvertisingChannelTypeEnum.Unmarshal(m, b)
}
func (m *AdvertisingChannelTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdvertisingChannelTypeEnum.Marshal(b, m, deterministic)
}
func (m *AdvertisingChannelTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertisingChannelTypeEnum.Merge(m, src)
}
func (m *AdvertisingChannelTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdvertisingChannelTypeEnum.Size(m)
}
func (m *AdvertisingChannelTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertisingChannelTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertisingChannelTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdvertisingChannelTypeEnum_AdvertisingChannelType", AdvertisingChannelTypeEnum_AdvertisingChannelType_name, AdvertisingChannelTypeEnum_AdvertisingChannelType_value)
	proto.RegisterType((*AdvertisingChannelTypeEnum)(nil), "google.ads.googleads.v1.enums.AdvertisingChannelTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/advertising_channel_type.proto", fileDescriptor_d496394fd88a2a26)
}

var fileDescriptor_d496394fd88a2a26 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4e, 0xbb, 0x40,
	0x18, 0xc5, 0xff, 0xd0, 0x7f, 0x5b, 0x9d, 0x6a, 0x1c, 0x59, 0xb8, 0xa8, 0x76, 0xd1, 0x1e, 0x60,
	0x08, 0x71, 0x37, 0xba, 0x99, 0x52, 0x2c, 0x44, 0x04, 0x22, 0x2d, 0x46, 0x43, 0xd2, 0x60, 0x21,
	0x48, 0xd2, 0xce, 0x90, 0x0e, 0x6d, 0xd2, 0x0b, 0x78, 0x10, 0x13, 0x37, 0x1e, 0xc5, 0xa3, 0xb8,
	0xf2, 0x08, 0x06, 0xc6, 0xd6, 0x4d, 0x75, 0x43, 0x5e, 0x78, 0xdf, 0xef, 0xe5, 0xfb, 0xde, 0x80,
	0xcb, 0x94, 0xb1, 0x74, 0x96, 0xa8, 0x51, 0xcc, 0x55, 0x21, 0x4b, 0xb5, 0xd2, 0xd4, 0x84, 0x2e,
	0xe7, 0x5c, 0x8d, 0xe2, 0x55, 0xb2, 0x28, 0x32, 0x9e, 0xd1, 0x74, 0x32, 0x7d, 0x8a, 0x28, 0x4d,
	0x66, 0x93, 0x62, 0x9d, 0x27, 0x28, 0x5f, 0xb0, 0x82, 0x29, 0x1d, 0x81, 0xa0, 0x28, 0xe6, 0x68,
	0x4b, 0xa3, 0x95, 0x86, 0x2a, 0xba, 0x7d, 0xb6, 0x09, 0xcf, 0x33, 0x35, 0xa2, 0x94, 0x15, 0x51,
	0x91, 0x31, 0xca, 0x05, 0xdc, 0x7b, 0x95, 0x40, 0x9b, 0xfc, 0xe4, 0xeb, 0x22, 0x7e, 0xb4, 0xce,
	0x13, 0x83, 0x2e, 0xe7, 0xbd, 0x67, 0x09, 0x9c, 0xec, 0xb6, 0x95, 0x23, 0xd0, 0x1a, 0x3b, 0xbe,
	0x67, 0xe8, 0xd6, 0x95, 0x65, 0x0c, 0xe0, 0x3f, 0xa5, 0x05, 0x9a, 0x63, 0xe7, 0xda, 0x71, 0xef,
	0x1c, 0x28, 0x29, 0x00, 0x34, 0x7c, 0x83, 0xdc, 0xea, 0x26, 0x94, 0x4b, 0x63, 0x60, 0xf9, 0x9e,
	0x4d, 0xee, 0x61, 0x4d, 0x39, 0x00, 0x7b, 0xbe, 0xe9, 0x7a, 0x9e, 0xe5, 0x0c, 0xe1, 0x7f, 0x65,
	0x1f, 0xd4, 0x4d, 0x77, 0x64, 0xd8, 0xb0, 0x5e, 0xca, 0xc0, 0x1a, 0x18, 0x2e, 0x6c, 0x28, 0xc7,
	0xe0, 0xf0, 0x66, 0x6c, 0x8f, 0xac, 0x89, 0x6e, 0x12, 0xc7, 0x31, 0x6c, 0xd8, 0xec, 0x7f, 0x4a,
	0xa0, 0x3b, 0x65, 0x73, 0xf4, 0xe7, 0xad, 0xfd, 0xd3, 0xdd, 0xbb, 0x7a, 0xe5, 0xa9, 0x9e, 0xf4,
	0xd0, 0xff, 0xa6, 0x53, 0x36, 0x8b, 0x68, 0x8a, 0xd8, 0x22, 0x55, 0xd3, 0x84, 0x56, 0x45, 0x6c,
	0x7a, 0xcf, 0x33, 0xfe, 0xcb, 0x33, 0x5c, 0x54, 0xdf, 0x17, 0xb9, 0x36, 0x24, 0xe4, 0x4d, 0xee,
	0x0c, 0x45, 0x14, 0x89, 0x39, 0x12, 0xb2, 0x54, 0x81, 0x86, 0xca, 0xda, 0xf8, 0xfb, 0xc6, 0x0f,
	0x49, 0xcc, 0xc3, 0xad, 0x1f, 0x06, 0x5a, 0x58, 0xf9, 0x1f, 0x72, 0x57, 0xfc, 0xc4, 0x98, 0xc4,
	0x1c, 0xe3, 0xed, 0x04, 0xc6, 0x81, 0x86, 0x71, 0x35, 0xf3, 0xd8, 0xa8, 0x16, 0x3b, 0xff, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0x32, 0x41, 0x6f, 0x1e, 0x02, 0x00, 0x00,
}
