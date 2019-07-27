// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/month_of_year.proto

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

// Enumerates months of the year, e.g., "January".
type MonthOfYearEnum_MonthOfYear int32

const (
	// Not specified.
	MonthOfYearEnum_UNSPECIFIED MonthOfYearEnum_MonthOfYear = 0
	// The value is unknown in this version.
	MonthOfYearEnum_UNKNOWN MonthOfYearEnum_MonthOfYear = 1
	// January.
	MonthOfYearEnum_JANUARY MonthOfYearEnum_MonthOfYear = 2
	// February.
	MonthOfYearEnum_FEBRUARY MonthOfYearEnum_MonthOfYear = 3
	// March.
	MonthOfYearEnum_MARCH MonthOfYearEnum_MonthOfYear = 4
	// April.
	MonthOfYearEnum_APRIL MonthOfYearEnum_MonthOfYear = 5
	// May.
	MonthOfYearEnum_MAY MonthOfYearEnum_MonthOfYear = 6
	// June.
	MonthOfYearEnum_JUNE MonthOfYearEnum_MonthOfYear = 7
	// July.
	MonthOfYearEnum_JULY MonthOfYearEnum_MonthOfYear = 8
	// August.
	MonthOfYearEnum_AUGUST MonthOfYearEnum_MonthOfYear = 9
	// September.
	MonthOfYearEnum_SEPTEMBER MonthOfYearEnum_MonthOfYear = 10
	// October.
	MonthOfYearEnum_OCTOBER MonthOfYearEnum_MonthOfYear = 11
	// November.
	MonthOfYearEnum_NOVEMBER MonthOfYearEnum_MonthOfYear = 12
	// December.
	MonthOfYearEnum_DECEMBER MonthOfYearEnum_MonthOfYear = 13
)

var MonthOfYearEnum_MonthOfYear_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "JANUARY",
	3:  "FEBRUARY",
	4:  "MARCH",
	5:  "APRIL",
	6:  "MAY",
	7:  "JUNE",
	8:  "JULY",
	9:  "AUGUST",
	10: "SEPTEMBER",
	11: "OCTOBER",
	12: "NOVEMBER",
	13: "DECEMBER",
}

var MonthOfYearEnum_MonthOfYear_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"JANUARY":     2,
	"FEBRUARY":    3,
	"MARCH":       4,
	"APRIL":       5,
	"MAY":         6,
	"JUNE":        7,
	"JULY":        8,
	"AUGUST":      9,
	"SEPTEMBER":   10,
	"OCTOBER":     11,
	"NOVEMBER":    12,
	"DECEMBER":    13,
}

func (x MonthOfYearEnum_MonthOfYear) String() string {
	return proto.EnumName(MonthOfYearEnum_MonthOfYear_name, int32(x))
}

func (MonthOfYearEnum_MonthOfYear) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5a6459dd75e57391, []int{0, 0}
}

// Container for enumeration of months of the year, e.g., "January".
type MonthOfYearEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonthOfYearEnum) Reset()         { *m = MonthOfYearEnum{} }
func (m *MonthOfYearEnum) String() string { return proto.CompactTextString(m) }
func (*MonthOfYearEnum) ProtoMessage()    {}
func (*MonthOfYearEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a6459dd75e57391, []int{0}
}

func (m *MonthOfYearEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonthOfYearEnum.Unmarshal(m, b)
}
func (m *MonthOfYearEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonthOfYearEnum.Marshal(b, m, deterministic)
}
func (m *MonthOfYearEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonthOfYearEnum.Merge(m, src)
}
func (m *MonthOfYearEnum) XXX_Size() int {
	return xxx_messageInfo_MonthOfYearEnum.Size(m)
}
func (m *MonthOfYearEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MonthOfYearEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MonthOfYearEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.MonthOfYearEnum_MonthOfYear", MonthOfYearEnum_MonthOfYear_name, MonthOfYearEnum_MonthOfYear_value)
	proto.RegisterType((*MonthOfYearEnum)(nil), "google.ads.googleads.v1.enums.MonthOfYearEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/month_of_year.proto", fileDescriptor_5a6459dd75e57391)
}

var fileDescriptor_5a6459dd75e57391 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x25, 0xe9, 0x4c, 0x1f, 0xce, 0x8c, 0xc6, 0xf2, 0x12, 0x31, 0x8b, 0x99, 0x0f, 0x70, 0x14,
	0xb1, 0x33, 0x2b, 0x27, 0xe3, 0x29, 0x2d, 0xcd, 0x43, 0x69, 0x13, 0x14, 0x14, 0xa9, 0x0a, 0x24,
	0x0d, 0x95, 0x1a, 0xbb, 0x8a, 0xd3, 0x4a, 0xfc, 0x0e, 0x4b, 0xfe, 0x81, 0x1f, 0x60, 0xc7, 0x47,
	0xb0, 0xe1, 0x2b, 0x90, 0x63, 0x5a, 0x75, 0x03, 0x9b, 0xe8, 0x9c, 0x7b, 0xee, 0xb9, 0x37, 0x3e,
	0x17, 0x38, 0xb5, 0x10, 0xf5, 0xae, 0xb2, 0x8b, 0x52, 0xda, 0x1a, 0x2a, 0x74, 0x74, 0xec, 0x8a,
	0x1f, 0x1a, 0x69, 0x37, 0x82, 0x77, 0x9f, 0xd7, 0x62, 0xb3, 0xfe, 0x52, 0x15, 0x2d, 0xde, 0xb7,
	0xa2, 0x13, 0xe8, 0x5e, 0xf7, 0xe1, 0xa2, 0x94, 0xf8, 0x6c, 0xc1, 0x47, 0x07, 0xf7, 0x96, 0x97,
	0xaf, 0x4e, 0x13, 0xf7, 0x5b, 0xbb, 0xe0, 0x5c, 0x74, 0x45, 0xb7, 0x15, 0x5c, 0x6a, 0xf3, 0xe3,
	0x4f, 0x03, 0xdc, 0xf9, 0x6a, 0x68, 0xb8, 0xc9, 0xaa, 0xa2, 0x65, 0xfc, 0xd0, 0x3c, 0x7e, 0x37,
	0x80, 0x75, 0x51, 0x43, 0x77, 0xc0, 0x4a, 0x82, 0x65, 0xc4, 0xbc, 0xd9, 0xf3, 0x8c, 0x3d, 0xc1,
	0x17, 0xc8, 0x02, 0xa3, 0x24, 0x78, 0x17, 0x84, 0xef, 0x03, 0x68, 0x28, 0x32, 0xa7, 0x41, 0x42,
	0xe3, 0x0c, 0x9a, 0xe8, 0x06, 0x8c, 0x9f, 0x99, 0x1b, 0xf7, 0x6c, 0x80, 0x26, 0xe0, 0xda, 0xa7,
	0xb1, 0xf7, 0x16, 0x5e, 0x29, 0x48, 0xa3, 0x78, 0xb6, 0x80, 0xd7, 0x68, 0x04, 0x06, 0x3e, 0xcd,
	0xe0, 0x10, 0x8d, 0xc1, 0xd5, 0x3c, 0x09, 0x18, 0x1c, 0x69, 0xb4, 0xc8, 0xe0, 0x18, 0x01, 0x30,
	0xa4, 0xc9, 0x34, 0x59, 0xae, 0xe0, 0x04, 0xdd, 0x82, 0xc9, 0x92, 0x45, 0x2b, 0xe6, 0xbb, 0x2c,
	0x86, 0x40, 0x2d, 0x0a, 0xbd, 0x55, 0xa8, 0x88, 0xa5, 0x16, 0x05, 0x61, 0xaa, 0xa5, 0x1b, 0xc5,
	0x9e, 0x98, 0xa7, 0xd9, 0xad, 0xfb, 0xcb, 0x00, 0x0f, 0x9f, 0x44, 0x83, 0xff, 0x9b, 0x8b, 0x0b,
	0x2f, 0x9e, 0x18, 0xa9, 0x2c, 0x22, 0xe3, 0x83, 0xfb, 0xd7, 0x52, 0x8b, 0x5d, 0xc1, 0x6b, 0x2c,
	0xda, 0xda, 0xae, 0x2b, 0xde, 0x27, 0x75, 0xba, 0xc6, 0x7e, 0x2b, 0xff, 0x71, 0x9c, 0x37, 0xfd,
	0xf7, 0xab, 0x39, 0x98, 0x52, 0xfa, 0xcd, 0xbc, 0x9f, 0xea, 0x51, 0xb4, 0x94, 0x58, 0x43, 0x85,
	0x52, 0x07, 0xab, 0x88, 0xe5, 0x8f, 0x93, 0x9e, 0xd3, 0x52, 0xe6, 0x67, 0x3d, 0x4f, 0x9d, 0xbc,
	0xd7, 0x7f, 0x9b, 0x0f, 0xba, 0x48, 0x08, 0x2d, 0x25, 0x21, 0xe7, 0x0e, 0x42, 0x52, 0x87, 0x90,
	0xbe, 0xe7, 0xe3, 0xb0, 0xff, 0xb1, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xa9, 0x63,
	0xff, 0x34, 0x02, 0x00, 0x00,
}
