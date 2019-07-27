// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/header_error.proto

package errors

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

// Enum describing possible header errors.
type HeaderErrorEnum_HeaderError int32

const (
	// Enum unspecified.
	HeaderErrorEnum_UNSPECIFIED HeaderErrorEnum_HeaderError = 0
	// The received error code is not known in this version.
	HeaderErrorEnum_UNKNOWN HeaderErrorEnum_HeaderError = 1
	// The login customer id could not be validated.
	HeaderErrorEnum_INVALID_LOGIN_CUSTOMER_ID HeaderErrorEnum_HeaderError = 3
)

var HeaderErrorEnum_HeaderError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "INVALID_LOGIN_CUSTOMER_ID",
}

var HeaderErrorEnum_HeaderError_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"INVALID_LOGIN_CUSTOMER_ID": 3,
}

func (x HeaderErrorEnum_HeaderError) String() string {
	return proto.EnumName(HeaderErrorEnum_HeaderError_name, int32(x))
}

func (HeaderErrorEnum_HeaderError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17dbd7ec3240a65f, []int{0, 0}
}

// Container for enum describing possible header errors.
type HeaderErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderErrorEnum) Reset()         { *m = HeaderErrorEnum{} }
func (m *HeaderErrorEnum) String() string { return proto.CompactTextString(m) }
func (*HeaderErrorEnum) ProtoMessage()    {}
func (*HeaderErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_17dbd7ec3240a65f, []int{0}
}

func (m *HeaderErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderErrorEnum.Unmarshal(m, b)
}
func (m *HeaderErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderErrorEnum.Marshal(b, m, deterministic)
}
func (m *HeaderErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderErrorEnum.Merge(m, src)
}
func (m *HeaderErrorEnum) XXX_Size() int {
	return xxx_messageInfo_HeaderErrorEnum.Size(m)
}
func (m *HeaderErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.HeaderErrorEnum_HeaderError", HeaderErrorEnum_HeaderError_name, HeaderErrorEnum_HeaderError_value)
	proto.RegisterType((*HeaderErrorEnum)(nil), "google.ads.googleads.v1.errors.HeaderErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/header_error.proto", fileDescriptor_17dbd7ec3240a65f)
}

var fileDescriptor_17dbd7ec3240a65f = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0xdd, 0x06, 0x0a, 0xd9, 0x61, 0xa5, 0x37, 0x45, 0x77, 0xe8, 0x03, 0x24, 0x04, 0x6f,
	0xf1, 0x94, 0xad, 0x75, 0x46, 0x67, 0x36, 0x9c, 0xab, 0x20, 0x95, 0x12, 0x4d, 0x89, 0x83, 0x2d,
	0x19, 0x49, 0xdd, 0x03, 0x79, 0xf4, 0x51, 0x7c, 0x14, 0x7d, 0x09, 0x69, 0x63, 0xcb, 0x2e, 0x7a,
	0xca, 0x97, 0x3f, 0xbf, 0xef, 0xfb, 0x7f, 0xfc, 0x01, 0x56, 0xc6, 0xa8, 0x75, 0x81, 0x84, 0x74,
	0xc8, 0xcb, 0x4a, 0xed, 0x30, 0x2a, 0xac, 0x35, 0xd6, 0xa1, 0xd7, 0x42, 0xc8, 0xc2, 0xe6, 0xf5,
	0x0f, 0x6e, 0xad, 0x29, 0x4d, 0x38, 0xf4, 0x1c, 0x14, 0xd2, 0xc1, 0xd6, 0x02, 0x77, 0x18, 0x7a,
	0xcb, 0xc9, 0x69, 0x13, 0xb9, 0x5d, 0x21, 0xa1, 0xb5, 0x29, 0x45, 0xb9, 0x32, 0xda, 0x79, 0x77,
	0xf4, 0x04, 0x06, 0x57, 0x75, 0x66, 0x52, 0xd1, 0x89, 0x7e, 0xdb, 0x44, 0xd7, 0xa0, 0xbf, 0x37,
	0x0a, 0x07, 0xa0, 0xbf, 0xe4, 0x8b, 0x79, 0x32, 0x66, 0x97, 0x2c, 0x89, 0x83, 0x83, 0xb0, 0x0f,
	0x8e, 0x96, 0xfc, 0x86, 0xcf, 0x1e, 0x78, 0xd0, 0x09, 0xcf, 0xc0, 0x31, 0xe3, 0x29, 0x9d, 0xb2,
	0x38, 0x9f, 0xce, 0x26, 0x8c, 0xe7, 0xe3, 0xe5, 0xe2, 0x7e, 0x76, 0x9b, 0xdc, 0xe5, 0x2c, 0x0e,
	0x7a, 0xa3, 0xef, 0x0e, 0x88, 0x5e, 0xcc, 0x06, 0xfe, 0xdf, 0x71, 0x14, 0xec, 0x2d, 0x9c, 0x57,
	0xbd, 0xe6, 0x9d, 0xc7, 0xf8, 0xd7, 0xa3, 0xcc, 0x5a, 0x68, 0x05, 0x8d, 0x55, 0x48, 0x15, 0xba,
	0x6e, 0xdd, 0x9c, 0x66, 0xbb, 0x72, 0x7f, 0x5d, 0xea, 0xc2, 0x3f, 0xef, 0xdd, 0xde, 0x84, 0xd2,
	0x8f, 0xee, 0x70, 0xe2, 0xc3, 0xa8, 0x74, 0xd0, 0xcb, 0x4a, 0xa5, 0x18, 0xd6, 0x2b, 0xdd, 0x67,
	0x03, 0x64, 0x54, 0xba, 0xac, 0x05, 0xb2, 0x14, 0x67, 0x1e, 0xf8, 0xea, 0x46, 0x7e, 0x4a, 0x08,
	0x95, 0x8e, 0x90, 0x16, 0x21, 0x24, 0xc5, 0x84, 0x78, 0xe8, 0xf9, 0xb0, 0x6e, 0x77, 0xfe, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xab, 0x94, 0x78, 0x3f, 0xc6, 0x01, 0x00, 0x00,
}
