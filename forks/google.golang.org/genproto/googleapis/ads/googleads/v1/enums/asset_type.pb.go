// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/asset_type.proto

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

// Enum describing possible types of asset.
type AssetTypeEnum_AssetType int32

const (
	// Not specified.
	AssetTypeEnum_UNSPECIFIED AssetTypeEnum_AssetType = 0
	// Used for return value only. Represents value unknown in this version.
	AssetTypeEnum_UNKNOWN AssetTypeEnum_AssetType = 1
	// YouTube video asset.
	AssetTypeEnum_YOUTUBE_VIDEO AssetTypeEnum_AssetType = 2
	// Media bundle asset.
	AssetTypeEnum_MEDIA_BUNDLE AssetTypeEnum_AssetType = 3
	// Image asset.
	AssetTypeEnum_IMAGE AssetTypeEnum_AssetType = 4
	// Text asset.
	AssetTypeEnum_TEXT AssetTypeEnum_AssetType = 5
)

var AssetTypeEnum_AssetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "YOUTUBE_VIDEO",
	3: "MEDIA_BUNDLE",
	4: "IMAGE",
	5: "TEXT",
}

var AssetTypeEnum_AssetType_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"YOUTUBE_VIDEO": 2,
	"MEDIA_BUNDLE":  3,
	"IMAGE":         4,
	"TEXT":          5,
}

func (x AssetTypeEnum_AssetType) String() string {
	return proto.EnumName(AssetTypeEnum_AssetType_name, int32(x))
}

func (AssetTypeEnum_AssetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_491e1f3963d1cbfd, []int{0, 0}
}

// Container for enum describing the types of asset.
type AssetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetTypeEnum) Reset()         { *m = AssetTypeEnum{} }
func (m *AssetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AssetTypeEnum) ProtoMessage()    {}
func (*AssetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_491e1f3963d1cbfd, []int{0}
}

func (m *AssetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetTypeEnum.Unmarshal(m, b)
}
func (m *AssetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetTypeEnum.Marshal(b, m, deterministic)
}
func (m *AssetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetTypeEnum.Merge(m, src)
}
func (m *AssetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AssetTypeEnum.Size(m)
}
func (m *AssetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AssetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.AssetTypeEnum_AssetType", AssetTypeEnum_AssetType_name, AssetTypeEnum_AssetType_value)
	proto.RegisterType((*AssetTypeEnum)(nil), "google.ads.googleads.v1.enums.AssetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/asset_type.proto", fileDescriptor_491e1f3963d1cbfd)
}

var fileDescriptor_491e1f3963d1cbfd = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4e, 0xf3, 0x30,
	0x14, 0xc4, 0xbf, 0xa6, 0xed, 0x07, 0x75, 0x29, 0x18, 0x2f, 0x11, 0x5d, 0xb4, 0x07, 0x70, 0x14,
	0xb1, 0x33, 0x2b, 0x87, 0x98, 0x2a, 0x82, 0xa6, 0x95, 0x48, 0xc2, 0x1f, 0x45, 0xaa, 0x42, 0x13,
	0x45, 0x95, 0x5a, 0x3b, 0xaa, 0xdd, 0x4a, 0xbd, 0x0e, 0x4b, 0x8e, 0xc2, 0x35, 0xd8, 0x71, 0x0a,
	0x14, 0x9b, 0x64, 0x07, 0x1b, 0x6b, 0xf4, 0xde, 0xef, 0x8d, 0xc6, 0x03, 0x70, 0x21, 0x44, 0xb1,
	0xce, 0xed, 0x34, 0x93, 0xb6, 0x91, 0x95, 0xda, 0x3b, 0x76, 0xce, 0x77, 0x1b, 0x69, 0xa7, 0x52,
	0xe6, 0x6a, 0xa1, 0x0e, 0x65, 0x8e, 0xcb, 0xad, 0x50, 0x02, 0x0d, 0x0d, 0x84, 0xd3, 0x4c, 0xe2,
	0x86, 0xc7, 0x7b, 0x07, 0x6b, 0xfe, 0xe2, 0xb2, 0xb6, 0x2b, 0x57, 0x76, 0xca, 0xb9, 0x50, 0xa9,
	0x5a, 0x09, 0x2e, 0xcd, 0xf1, 0x58, 0x81, 0x01, 0xad, 0x0c, 0xc3, 0x43, 0x99, 0x33, 0xbe, 0xdb,
	0x8c, 0x97, 0xa0, 0xd7, 0x0c, 0xd0, 0x19, 0xe8, 0x47, 0xc1, 0xc3, 0x9c, 0xdd, 0xf8, 0xb7, 0x3e,
	0xf3, 0xe0, 0x3f, 0xd4, 0x07, 0x47, 0x51, 0x70, 0x17, 0xcc, 0x1e, 0x03, 0xd8, 0x42, 0xe7, 0x60,
	0xf0, 0x3c, 0x8b, 0xc2, 0xc8, 0x65, 0x8b, 0xd8, 0xf7, 0xd8, 0x0c, 0x5a, 0x08, 0x82, 0x93, 0x29,
	0xf3, 0x7c, 0xba, 0x70, 0xa3, 0xc0, 0xbb, 0x67, 0xb0, 0x8d, 0x7a, 0xa0, 0xeb, 0x4f, 0xe9, 0x84,
	0xc1, 0x0e, 0x3a, 0x06, 0x9d, 0x90, 0x3d, 0x85, 0xb0, 0xeb, 0x7e, 0xb6, 0xc0, 0x68, 0x29, 0x36,
	0xf8, 0xcf, 0xe4, 0xee, 0x69, 0x13, 0x64, 0x5e, 0x65, 0x9d, 0xb7, 0x5e, 0xdc, 0x9f, 0x83, 0x42,
	0xac, 0x53, 0x5e, 0x60, 0xb1, 0x2d, 0xec, 0x22, 0xe7, 0xfa, 0x27, 0x75, 0x55, 0xe5, 0x4a, 0xfe,
	0xd2, 0xdc, 0xb5, 0x7e, 0xdf, 0xac, 0xf6, 0x84, 0xd2, 0x77, 0x6b, 0x38, 0x31, 0x56, 0x34, 0x93,
	0xd8, 0xc8, 0x4a, 0xc5, 0x0e, 0xae, 0x5a, 0x90, 0x1f, 0xf5, 0x3e, 0xa1, 0x99, 0x4c, 0x9a, 0x7d,
	0x12, 0x3b, 0x89, 0xde, 0x7f, 0x59, 0x23, 0x33, 0x24, 0x84, 0x66, 0x92, 0x90, 0x86, 0x20, 0x24,
	0x76, 0x08, 0xd1, 0xcc, 0xeb, 0x7f, 0x1d, 0xec, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x60, 0xe7,
	0xfb, 0x60, 0xd1, 0x01, 0x00, 0x00,
}
