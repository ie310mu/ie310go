// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_list_access_status.proto

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

// Enum containing possible user list access statuses.
type UserListAccessStatusEnum_UserListAccessStatus int32

const (
	// Not specified.
	UserListAccessStatusEnum_UNSPECIFIED UserListAccessStatusEnum_UserListAccessStatus = 0
	// Used for return value only. Represents value unknown in this version.
	UserListAccessStatusEnum_UNKNOWN UserListAccessStatusEnum_UserListAccessStatus = 1
	// The access is enabled.
	UserListAccessStatusEnum_ENABLED UserListAccessStatusEnum_UserListAccessStatus = 2
	// The access is disabled.
	UserListAccessStatusEnum_DISABLED UserListAccessStatusEnum_UserListAccessStatus = 3
)

var UserListAccessStatusEnum_UserListAccessStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "DISABLED",
}

var UserListAccessStatusEnum_UserListAccessStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"DISABLED":    3,
}

func (x UserListAccessStatusEnum_UserListAccessStatus) String() string {
	return proto.EnumName(UserListAccessStatusEnum_UserListAccessStatus_name, int32(x))
}

func (UserListAccessStatusEnum_UserListAccessStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4bc74499510e4946, []int{0, 0}
}

// Indicates if this client still has access to the list.
type UserListAccessStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListAccessStatusEnum) Reset()         { *m = UserListAccessStatusEnum{} }
func (m *UserListAccessStatusEnum) String() string { return proto.CompactTextString(m) }
func (*UserListAccessStatusEnum) ProtoMessage()    {}
func (*UserListAccessStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bc74499510e4946, []int{0}
}

func (m *UserListAccessStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListAccessStatusEnum.Unmarshal(m, b)
}
func (m *UserListAccessStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListAccessStatusEnum.Marshal(b, m, deterministic)
}
func (m *UserListAccessStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListAccessStatusEnum.Merge(m, src)
}
func (m *UserListAccessStatusEnum) XXX_Size() int {
	return xxx_messageInfo_UserListAccessStatusEnum.Size(m)
}
func (m *UserListAccessStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListAccessStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListAccessStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserListAccessStatusEnum_UserListAccessStatus", UserListAccessStatusEnum_UserListAccessStatus_name, UserListAccessStatusEnum_UserListAccessStatus_value)
	proto.RegisterType((*UserListAccessStatusEnum)(nil), "google.ads.googleads.v1.enums.UserListAccessStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_list_access_status.proto", fileDescriptor_4bc74499510e4946)
}

var fileDescriptor_4bc74499510e4946 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x6a, 0xeb, 0x30,
	0x10, 0xfd, 0x71, 0xe0, 0xb7, 0x28, 0x85, 0x06, 0xd3, 0x45, 0x5b, 0x9a, 0x45, 0x72, 0x00, 0x09,
	0xd3, 0x9d, 0xb2, 0x92, 0x1b, 0x37, 0x84, 0x06, 0x27, 0x10, 0x92, 0x42, 0x31, 0x04, 0x35, 0x16,
	0xc2, 0x34, 0x96, 0x8c, 0x47, 0xce, 0x81, 0xba, 0xec, 0x51, 0x7a, 0x94, 0x2e, 0x7a, 0x86, 0x62,
	0x29, 0xf6, 0x2a, 0xed, 0x46, 0xbc, 0xd1, 0x9b, 0xf7, 0xe6, 0xcd, 0xa0, 0xb1, 0xd4, 0x5a, 0xee,
	0x05, 0xe1, 0x29, 0x10, 0x07, 0x6b, 0x74, 0x08, 0x88, 0x50, 0x55, 0x0e, 0xa4, 0x02, 0x51, 0x6e,
	0xf7, 0x19, 0x98, 0x2d, 0xdf, 0xed, 0x04, 0xc0, 0x16, 0x0c, 0x37, 0x15, 0xe0, 0xa2, 0xd4, 0x46,
	0xfb, 0x03, 0xa7, 0xc0, 0x3c, 0x05, 0xdc, 0x8a, 0xf1, 0x21, 0xc0, 0x56, 0x7c, 0x7b, 0xd7, 0x78,
	0x17, 0x19, 0xe1, 0x4a, 0x69, 0xc3, 0x4d, 0xa6, 0xd5, 0x51, 0x3c, 0x7a, 0x43, 0xd7, 0x6b, 0x10,
	0xe5, 0x3c, 0x03, 0xc3, 0xac, 0xf7, 0xca, 0x5a, 0x47, 0xaa, 0xca, 0x47, 0x0b, 0x74, 0x75, 0x8a,
	0xf3, 0x2f, 0x51, 0x6f, 0x1d, 0xaf, 0x96, 0xd1, 0xc3, 0xec, 0x71, 0x16, 0x4d, 0xfa, 0xff, 0xfc,
	0x1e, 0x3a, 0x5b, 0xc7, 0x4f, 0xf1, 0xe2, 0x39, 0xee, 0x77, 0xea, 0x22, 0x8a, 0x59, 0x38, 0x8f,
	0x26, 0x7d, 0xcf, 0xbf, 0x40, 0xe7, 0x93, 0xd9, 0xca, 0x55, 0xdd, 0xf0, 0xbb, 0x83, 0x86, 0x3b,
	0x9d, 0xe3, 0x3f, 0x03, 0x87, 0x37, 0xa7, 0x86, 0x2e, 0xeb, 0xb4, 0xcb, 0xce, 0x4b, 0x78, 0xd4,
	0x4a, 0xbd, 0xe7, 0x4a, 0x62, 0x5d, 0x4a, 0x22, 0x85, 0xb2, 0xbb, 0x34, 0x97, 0x2b, 0x32, 0xf8,
	0xe5, 0x90, 0x63, 0xfb, 0xbe, 0x7b, 0xdd, 0x29, 0x63, 0x1f, 0xde, 0x60, 0xea, 0xac, 0x58, 0x0a,
	0xd8, 0xc1, 0x1a, 0x6d, 0x02, 0x5c, 0x2f, 0x0f, 0x9f, 0x0d, 0x9f, 0xb0, 0x14, 0x92, 0x96, 0x4f,
	0x36, 0x41, 0x62, 0xf9, 0x2f, 0x6f, 0xe8, 0x3e, 0x29, 0x65, 0x29, 0x50, 0xda, 0x76, 0x50, 0xba,
	0x09, 0x28, 0xb5, 0x3d, 0xaf, 0xff, 0x6d, 0xb0, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99,
	0x4b, 0x85, 0x3b, 0xe0, 0x01, 0x00, 0x00,
}
