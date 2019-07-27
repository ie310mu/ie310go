// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/feed_item_validation_status.proto

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

// The possible validation statuses of a feed item.
type FeedItemValidationStatusEnum_FeedItemValidationStatus int32

const (
	// No value has been specified.
	FeedItemValidationStatusEnum_UNSPECIFIED FeedItemValidationStatusEnum_FeedItemValidationStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemValidationStatusEnum_UNKNOWN FeedItemValidationStatusEnum_FeedItemValidationStatus = 1
	// Validation pending.
	FeedItemValidationStatusEnum_PENDING FeedItemValidationStatusEnum_FeedItemValidationStatus = 2
	// An error was found.
	FeedItemValidationStatusEnum_INVALID FeedItemValidationStatusEnum_FeedItemValidationStatus = 3
	// Feed item is semantically well-formed.
	FeedItemValidationStatusEnum_VALID FeedItemValidationStatusEnum_FeedItemValidationStatus = 4
)

var FeedItemValidationStatusEnum_FeedItemValidationStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "INVALID",
	4: "VALID",
}

var FeedItemValidationStatusEnum_FeedItemValidationStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PENDING":     2,
	"INVALID":     3,
	"VALID":       4,
}

func (x FeedItemValidationStatusEnum_FeedItemValidationStatus) String() string {
	return proto.EnumName(FeedItemValidationStatusEnum_FeedItemValidationStatus_name, int32(x))
}

func (FeedItemValidationStatusEnum_FeedItemValidationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3f43edd0f85e01a, []int{0, 0}
}

// Container for enum describing possible validation statuses of a feed item.
type FeedItemValidationStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemValidationStatusEnum) Reset()         { *m = FeedItemValidationStatusEnum{} }
func (m *FeedItemValidationStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemValidationStatusEnum) ProtoMessage()    {}
func (*FeedItemValidationStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3f43edd0f85e01a, []int{0}
}

func (m *FeedItemValidationStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemValidationStatusEnum.Unmarshal(m, b)
}
func (m *FeedItemValidationStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemValidationStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemValidationStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemValidationStatusEnum.Merge(m, src)
}
func (m *FeedItemValidationStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemValidationStatusEnum.Size(m)
}
func (m *FeedItemValidationStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemValidationStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemValidationStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.FeedItemValidationStatusEnum_FeedItemValidationStatus", FeedItemValidationStatusEnum_FeedItemValidationStatus_name, FeedItemValidationStatusEnum_FeedItemValidationStatus_value)
	proto.RegisterType((*FeedItemValidationStatusEnum)(nil), "google.ads.googleads.v1.enums.FeedItemValidationStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/feed_item_validation_status.proto", fileDescriptor_e3f43edd0f85e01a)
}

var fileDescriptor_e3f43edd0f85e01a = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4a, 0x03, 0x31,
	0x14, 0xc6, 0xed, 0xd4, 0x3f, 0x98, 0x2e, 0x1c, 0x66, 0x25, 0xd2, 0x2e, 0xda, 0x03, 0x64, 0x18,
	0xdc, 0xc5, 0x85, 0x4c, 0xed, 0xb4, 0x04, 0x25, 0x0e, 0x94, 0x8e, 0x20, 0x23, 0x25, 0x9a, 0x18,
	0x06, 0x3a, 0x49, 0x69, 0xd2, 0xee, 0xbc, 0x8c, 0x4b, 0x8f, 0xe2, 0x51, 0x5c, 0x7a, 0x02, 0x49,
	0xd2, 0xce, 0x6e, 0xdc, 0x84, 0xef, 0xe5, 0x7b, 0xef, 0xc7, 0x7b, 0x1f, 0xb8, 0x15, 0x4a, 0x89,
	0x15, 0x8f, 0x29, 0xd3, 0xb1, 0x97, 0x56, 0xed, 0x92, 0x98, 0xcb, 0x6d, 0xad, 0xe3, 0x77, 0xce,
	0xd9, 0xb2, 0x32, 0xbc, 0x5e, 0xee, 0xe8, 0xaa, 0x62, 0xd4, 0x54, 0x4a, 0x2e, 0xb5, 0xa1, 0x66,
	0xab, 0xe1, 0x7a, 0xa3, 0x8c, 0x8a, 0x06, 0x7e, 0x0a, 0x52, 0xa6, 0x61, 0x03, 0x80, 0xbb, 0x04,
	0x3a, 0xc0, 0x55, 0xff, 0xc0, 0x5f, 0x57, 0x31, 0x95, 0x52, 0x19, 0x87, 0xd8, 0x0f, 0x8f, 0x3e,
	0x40, 0x7f, 0xca, 0x39, 0xc3, 0x86, 0xd7, 0x45, 0xc3, 0x9f, 0x3b, 0x7c, 0x26, 0xb7, 0xf5, 0xe8,
	0x05, 0x5c, 0xb6, 0xf9, 0xd1, 0x05, 0xe8, 0x2d, 0xc8, 0x3c, 0xcf, 0xee, 0xf0, 0x14, 0x67, 0x93,
	0xf0, 0x28, 0xea, 0x81, 0xb3, 0x05, 0xb9, 0x27, 0x8f, 0x4f, 0x24, 0xec, 0xd8, 0x22, 0xcf, 0xc8,
	0x04, 0x93, 0x59, 0x18, 0xd8, 0x02, 0x93, 0x22, 0x7d, 0xc0, 0x93, 0xb0, 0x1b, 0x9d, 0x83, 0x13,
	0x2f, 0x8f, 0xc7, 0xbf, 0x1d, 0x30, 0x7c, 0x53, 0x35, 0xfc, 0xf7, 0x84, 0xf1, 0xa0, 0x6d, 0x85,
	0xdc, 0xde, 0x90, 0x77, 0x9e, 0xc7, 0xfb, 0x79, 0xa1, 0x56, 0x54, 0x0a, 0xa8, 0x36, 0x22, 0x16,
	0x5c, 0xba, 0x0b, 0x0f, 0x99, 0xae, 0x2b, 0xdd, 0x12, 0xf1, 0x8d, 0x7b, 0x3f, 0x83, 0xee, 0x2c,
	0x4d, 0xbf, 0x82, 0xc1, 0xcc, 0xa3, 0x52, 0xa6, 0xa1, 0x97, 0x56, 0x15, 0x09, 0xb4, 0x71, 0xe8,
	0xef, 0x83, 0x5f, 0xa6, 0x4c, 0x97, 0x8d, 0x5f, 0x16, 0x49, 0xe9, 0xfc, 0x9f, 0x60, 0xe8, 0x3f,
	0x11, 0x4a, 0x99, 0x46, 0xa8, 0xe9, 0x40, 0xa8, 0x48, 0x10, 0x72, 0x3d, 0xaf, 0xa7, 0x6e, 0xb1,
	0xeb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x79, 0x89, 0xf3, 0xfa, 0x01, 0x00, 0x00,
}
