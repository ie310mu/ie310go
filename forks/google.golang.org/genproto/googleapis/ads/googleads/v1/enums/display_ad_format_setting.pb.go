// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/display_ad_format_setting.proto

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

// Enumerates display ad format settings.
type DisplayAdFormatSettingEnum_DisplayAdFormatSetting int32

const (
	// Not specified.
	DisplayAdFormatSettingEnum_UNSPECIFIED DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 0
	// The value is unknown in this version.
	DisplayAdFormatSettingEnum_UNKNOWN DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 1
	// Text, image and native formats.
	DisplayAdFormatSettingEnum_ALL_FORMATS DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 2
	// Text and image formats.
	DisplayAdFormatSettingEnum_NON_NATIVE DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 3
	// Native format, i.e. the format rendering is controlled by the publisher
	// and not by Google.
	DisplayAdFormatSettingEnum_NATIVE DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 4
)

var DisplayAdFormatSettingEnum_DisplayAdFormatSetting_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ALL_FORMATS",
	3: "NON_NATIVE",
	4: "NATIVE",
}

var DisplayAdFormatSettingEnum_DisplayAdFormatSetting_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ALL_FORMATS": 2,
	"NON_NATIVE":  3,
	"NATIVE":      4,
}

func (x DisplayAdFormatSettingEnum_DisplayAdFormatSetting) String() string {
	return proto.EnumName(DisplayAdFormatSettingEnum_DisplayAdFormatSetting_name, int32(x))
}

func (DisplayAdFormatSettingEnum_DisplayAdFormatSetting) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4acbddf81279131, []int{0, 0}
}

// Container for display ad format settings.
type DisplayAdFormatSettingEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisplayAdFormatSettingEnum) Reset()         { *m = DisplayAdFormatSettingEnum{} }
func (m *DisplayAdFormatSettingEnum) String() string { return proto.CompactTextString(m) }
func (*DisplayAdFormatSettingEnum) ProtoMessage()    {}
func (*DisplayAdFormatSettingEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4acbddf81279131, []int{0}
}

func (m *DisplayAdFormatSettingEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisplayAdFormatSettingEnum.Unmarshal(m, b)
}
func (m *DisplayAdFormatSettingEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisplayAdFormatSettingEnum.Marshal(b, m, deterministic)
}
func (m *DisplayAdFormatSettingEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisplayAdFormatSettingEnum.Merge(m, src)
}
func (m *DisplayAdFormatSettingEnum) XXX_Size() int {
	return xxx_messageInfo_DisplayAdFormatSettingEnum.Size(m)
}
func (m *DisplayAdFormatSettingEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DisplayAdFormatSettingEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DisplayAdFormatSettingEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.DisplayAdFormatSettingEnum_DisplayAdFormatSetting", DisplayAdFormatSettingEnum_DisplayAdFormatSetting_name, DisplayAdFormatSettingEnum_DisplayAdFormatSetting_value)
	proto.RegisterType((*DisplayAdFormatSettingEnum)(nil), "google.ads.googleads.v1.enums.DisplayAdFormatSettingEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/display_ad_format_setting.proto", fileDescriptor_c4acbddf81279131)
}

var fileDescriptor_c4acbddf81279131 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4b, 0xfb, 0x30,
	0x1c, 0xfd, 0xb7, 0xfb, 0x33, 0x21, 0x03, 0x2d, 0x3d, 0x78, 0x98, 0xee, 0xb0, 0x7d, 0x80, 0x94,
	0xe2, 0x2d, 0xe2, 0x21, 0x73, 0xdd, 0x18, 0xce, 0x6c, 0xb8, 0xad, 0x82, 0x14, 0x4a, 0x5c, 0x6a,
	0x28, 0xac, 0x49, 0x59, 0xb2, 0x81, 0x47, 0xbf, 0x8a, 0x47, 0x3f, 0x8a, 0x1f, 0xc5, 0x93, 0x1f,
	0x41, 0x9a, 0x6c, 0x3d, 0x4d, 0x2f, 0xe1, 0x91, 0xf7, 0x7b, 0xef, 0xf7, 0x7b, 0x0f, 0xdc, 0x70,
	0x29, 0xf9, 0x3a, 0x0b, 0x28, 0x53, 0x81, 0x85, 0x15, 0xda, 0x85, 0x41, 0x26, 0xb6, 0x85, 0x0a,
	0x58, 0xae, 0xca, 0x35, 0x7d, 0x4d, 0x29, 0x4b, 0x5f, 0xe4, 0xa6, 0xa0, 0x3a, 0x55, 0x99, 0xd6,
	0xb9, 0xe0, 0xb0, 0xdc, 0x48, 0x2d, 0xfd, 0x8e, 0xd5, 0x40, 0xca, 0x14, 0xac, 0xe5, 0x70, 0x17,
	0x42, 0x23, 0x6f, 0x5f, 0x1e, 0xdc, 0xcb, 0x3c, 0xa0, 0x42, 0x48, 0x4d, 0x75, 0x2e, 0x85, 0xb2,
	0xe2, 0xde, 0x9b, 0x03, 0xda, 0x03, 0xbb, 0x00, 0xb3, 0xa1, 0xb1, 0x9f, 0x5b, 0xf7, 0x48, 0x6c,
	0x8b, 0xde, 0x0a, 0x9c, 0x1f, 0x67, 0xfd, 0x33, 0xd0, 0x5a, 0x92, 0xf9, 0x2c, 0xba, 0x1d, 0x0f,
	0xc7, 0xd1, 0xc0, 0xfb, 0xe7, 0xb7, 0xc0, 0xc9, 0x92, 0xdc, 0x91, 0xe9, 0x23, 0xf1, 0x9c, 0x8a,
	0xc5, 0x93, 0x49, 0x3a, 0x9c, 0x3e, 0xdc, 0xe3, 0xc5, 0xdc, 0x73, 0xfd, 0x53, 0x00, 0xc8, 0x94,
	0xa4, 0x04, 0x2f, 0xc6, 0x71, 0xe4, 0x35, 0x7c, 0x00, 0x9a, 0x7b, 0xfc, 0xbf, 0xff, 0xed, 0x80,
	0xee, 0x4a, 0x16, 0xf0, 0xcf, 0x1c, 0xfd, 0x8b, 0xe3, 0x87, 0xcc, 0xaa, 0x18, 0x33, 0xe7, 0xa9,
	0xbf, 0x57, 0x73, 0xb9, 0xa6, 0x82, 0x43, 0xb9, 0xe1, 0x01, 0xcf, 0x84, 0x09, 0x79, 0x28, 0xb5,
	0xcc, 0xd5, 0x2f, 0x1d, 0x5f, 0x9b, 0xf7, 0xdd, 0x6d, 0x8c, 0x30, 0xfe, 0x70, 0x3b, 0x23, 0x6b,
	0x85, 0x99, 0x82, 0x16, 0x56, 0x28, 0x0e, 0x61, 0x55, 0x89, 0xfa, 0x3c, 0xf0, 0x09, 0x66, 0x2a,
	0xa9, 0xf9, 0x24, 0x0e, 0x13, 0xc3, 0x7f, 0xb9, 0x5d, 0xfb, 0x89, 0x10, 0x66, 0x0a, 0xa1, 0x7a,
	0x02, 0xa1, 0x38, 0x44, 0xc8, 0xcc, 0x3c, 0x37, 0xcd, 0x61, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0xeb, 0x40, 0x86, 0xfb, 0x01, 0x00, 0x00,
}
