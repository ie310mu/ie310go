// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/placement_type.proto

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

// Possible placement types for a feed mapping.
type PlacementTypeEnum_PlacementType int32

const (
	// Not specified.
	PlacementTypeEnum_UNSPECIFIED PlacementTypeEnum_PlacementType = 0
	// Used for return value only. Represents value unknown in this version.
	PlacementTypeEnum_UNKNOWN PlacementTypeEnum_PlacementType = 1
	// Websites(e.g. 'www.flowers4sale.com').
	PlacementTypeEnum_WEBSITE PlacementTypeEnum_PlacementType = 2
	// Mobile application categories(e.g. 'Games').
	PlacementTypeEnum_MOBILE_APP_CATEGORY PlacementTypeEnum_PlacementType = 3
	// mobile applications(e.g. 'mobileapp::2-com.whatsthewordanswers').
	PlacementTypeEnum_MOBILE_APPLICATION PlacementTypeEnum_PlacementType = 4
	// YouTube videos(e.g. 'youtube.com/video/wtLJPvx7-ys').
	PlacementTypeEnum_YOUTUBE_VIDEO PlacementTypeEnum_PlacementType = 5
	// YouTube channels(e.g. 'youtube.com::L8ZULXASCc1I_oaOT0NaOQ').
	PlacementTypeEnum_YOUTUBE_CHANNEL PlacementTypeEnum_PlacementType = 6
)

var PlacementTypeEnum_PlacementType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "WEBSITE",
	3: "MOBILE_APP_CATEGORY",
	4: "MOBILE_APPLICATION",
	5: "YOUTUBE_VIDEO",
	6: "YOUTUBE_CHANNEL",
}

var PlacementTypeEnum_PlacementType_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"WEBSITE":             2,
	"MOBILE_APP_CATEGORY": 3,
	"MOBILE_APPLICATION":  4,
	"YOUTUBE_VIDEO":       5,
	"YOUTUBE_CHANNEL":     6,
}

func (x PlacementTypeEnum_PlacementType) String() string {
	return proto.EnumName(PlacementTypeEnum_PlacementType_name, int32(x))
}

func (PlacementTypeEnum_PlacementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b1a4b1f4d95c82a7, []int{0, 0}
}

// Container for enum describing possible placement types.
type PlacementTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlacementTypeEnum) Reset()         { *m = PlacementTypeEnum{} }
func (m *PlacementTypeEnum) String() string { return proto.CompactTextString(m) }
func (*PlacementTypeEnum) ProtoMessage()    {}
func (*PlacementTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a4b1f4d95c82a7, []int{0}
}

func (m *PlacementTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacementTypeEnum.Unmarshal(m, b)
}
func (m *PlacementTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacementTypeEnum.Marshal(b, m, deterministic)
}
func (m *PlacementTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementTypeEnum.Merge(m, src)
}
func (m *PlacementTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PlacementTypeEnum.Size(m)
}
func (m *PlacementTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.PlacementTypeEnum_PlacementType", PlacementTypeEnum_PlacementType_name, PlacementTypeEnum_PlacementType_value)
	proto.RegisterType((*PlacementTypeEnum)(nil), "google.ads.googleads.v1.enums.PlacementTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/placement_type.proto", fileDescriptor_b1a4b1f4d95c82a7)
}

var fileDescriptor_b1a4b1f4d95c82a7 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x6a, 0xa3, 0x40,
	0x14, 0x5e, 0xcd, 0x6e, 0x16, 0x26, 0x84, 0x98, 0x09, 0xec, 0xc2, 0xb2, 0xb9, 0x48, 0x1e, 0x60,
	0xc4, 0xf6, 0x6e, 0x7a, 0x35, 0x9a, 0x69, 0x2a, 0x4d, 0x55, 0x1a, 0x35, 0xa4, 0x08, 0x62, 0xa3,
	0x48, 0x20, 0xce, 0x48, 0xc6, 0x04, 0xf2, 0x1c, 0x7d, 0x83, 0xde, 0xb5, 0x8f, 0xd2, 0x17, 0x29,
	0xf4, 0x29, 0x8a, 0x5a, 0x53, 0x72, 0xd1, 0xde, 0x0c, 0xdf, 0x9c, 0xef, 0x87, 0x73, 0x3e, 0x70,
	0x96, 0x72, 0x9e, 0x6e, 0x12, 0x35, 0x8a, 0x85, 0x5a, 0xc3, 0x12, 0xed, 0x35, 0x35, 0x61, 0xbb,
	0x4c, 0xa8, 0xf9, 0x26, 0x5a, 0x25, 0x59, 0xc2, 0x8a, 0xb0, 0x38, 0xe4, 0x09, 0xca, 0xb7, 0xbc,
	0xe0, 0x70, 0x58, 0x0b, 0x51, 0x14, 0x0b, 0x74, 0xf4, 0xa0, 0xbd, 0x86, 0x2a, 0xcf, 0xbf, 0xff,
	0x4d, 0x64, 0xbe, 0x56, 0x23, 0xc6, 0x78, 0x11, 0x15, 0x6b, 0xce, 0x44, 0x6d, 0x1e, 0x3f, 0x49,
	0xa0, 0xef, 0x34, 0xa9, 0xee, 0x21, 0x4f, 0x28, 0xdb, 0x65, 0xe3, 0x07, 0x09, 0x74, 0x4f, 0xa6,
	0xb0, 0x07, 0x3a, 0x9e, 0x35, 0x77, 0xa8, 0x61, 0x5e, 0x9a, 0x74, 0xa2, 0xfc, 0x80, 0x1d, 0xf0,
	0xdb, 0xb3, 0xae, 0x2d, 0x7b, 0x61, 0x29, 0x52, 0xf9, 0x59, 0x50, 0x7d, 0x6e, 0xba, 0x54, 0x91,
	0xe1, 0x5f, 0x30, 0xb8, 0xb1, 0x75, 0x73, 0x46, 0x43, 0xe2, 0x38, 0xa1, 0x41, 0x5c, 0x3a, 0xb5,
	0x6f, 0x97, 0x4a, 0x0b, 0xfe, 0x01, 0xf0, 0x93, 0x98, 0x99, 0x06, 0x71, 0x4d, 0xdb, 0x52, 0x7e,
	0xc2, 0x3e, 0xe8, 0x2e, 0x6d, 0xcf, 0xf5, 0x74, 0x1a, 0xfa, 0xe6, 0x84, 0xda, 0xca, 0x2f, 0x38,
	0x00, 0xbd, 0x66, 0x64, 0x5c, 0x11, 0xcb, 0xa2, 0x33, 0xa5, 0xad, 0xbf, 0x4a, 0x60, 0xb4, 0xe2,
	0x19, 0xfa, 0xf6, 0x5e, 0x1d, 0x9e, 0x2c, 0xee, 0x94, 0x57, 0x3a, 0xd2, 0x9d, 0xfe, 0x61, 0x4a,
	0xf9, 0x26, 0x62, 0x29, 0xe2, 0xdb, 0x54, 0x4d, 0x13, 0x56, 0x75, 0xd0, 0x14, 0x9d, 0xaf, 0xc5,
	0x17, 0xbd, 0x5f, 0x54, 0xef, 0xa3, 0xdc, 0x9a, 0x12, 0xf2, 0x2c, 0x0f, 0xa7, 0x75, 0x14, 0x89,
	0x05, 0xaa, 0x61, 0x89, 0x7c, 0x0d, 0x95, 0xd5, 0x89, 0x97, 0x86, 0x0f, 0x48, 0x2c, 0x82, 0x23,
	0x1f, 0xf8, 0x5a, 0x50, 0xf1, 0x6f, 0xf2, 0xa8, 0x1e, 0x62, 0x4c, 0x62, 0x81, 0xf1, 0x51, 0x81,
	0xb1, 0xaf, 0x61, 0x5c, 0x69, 0xee, 0xdb, 0xd5, 0x62, 0xe7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x08, 0x71, 0xbc, 0x14, 0x0f, 0x02, 0x00, 0x00,
}
