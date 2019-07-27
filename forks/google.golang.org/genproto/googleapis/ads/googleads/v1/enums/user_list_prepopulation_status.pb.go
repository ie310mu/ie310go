// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_list_prepopulation_status.proto

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

// Enum describing possible user list prepopulation status.
type UserListPrepopulationStatusEnum_UserListPrepopulationStatus int32

const (
	// Not specified.
	UserListPrepopulationStatusEnum_UNSPECIFIED UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 0
	// Used for return value only. Represents value unknown in this version.
	UserListPrepopulationStatusEnum_UNKNOWN UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 1
	// Prepopoulation is being requested.
	UserListPrepopulationStatusEnum_REQUESTED UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 2
	// Prepopulation is finished.
	UserListPrepopulationStatusEnum_FINISHED UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 3
	// Prepopulation failed.
	UserListPrepopulationStatusEnum_FAILED UserListPrepopulationStatusEnum_UserListPrepopulationStatus = 4
)

var UserListPrepopulationStatusEnum_UserListPrepopulationStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "REQUESTED",
	3: "FINISHED",
	4: "FAILED",
}

var UserListPrepopulationStatusEnum_UserListPrepopulationStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"REQUESTED":   2,
	"FINISHED":    3,
	"FAILED":      4,
}

func (x UserListPrepopulationStatusEnum_UserListPrepopulationStatus) String() string {
	return proto.EnumName(UserListPrepopulationStatusEnum_UserListPrepopulationStatus_name, int32(x))
}

func (UserListPrepopulationStatusEnum_UserListPrepopulationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f4f54f43cc394fe, []int{0, 0}
}

// Indicates status of prepopulation based on the rule.
type UserListPrepopulationStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListPrepopulationStatusEnum) Reset()         { *m = UserListPrepopulationStatusEnum{} }
func (m *UserListPrepopulationStatusEnum) String() string { return proto.CompactTextString(m) }
func (*UserListPrepopulationStatusEnum) ProtoMessage()    {}
func (*UserListPrepopulationStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4f54f43cc394fe, []int{0}
}

func (m *UserListPrepopulationStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListPrepopulationStatusEnum.Unmarshal(m, b)
}
func (m *UserListPrepopulationStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListPrepopulationStatusEnum.Marshal(b, m, deterministic)
}
func (m *UserListPrepopulationStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListPrepopulationStatusEnum.Merge(m, src)
}
func (m *UserListPrepopulationStatusEnum) XXX_Size() int {
	return xxx_messageInfo_UserListPrepopulationStatusEnum.Size(m)
}
func (m *UserListPrepopulationStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListPrepopulationStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListPrepopulationStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserListPrepopulationStatusEnum_UserListPrepopulationStatus", UserListPrepopulationStatusEnum_UserListPrepopulationStatus_name, UserListPrepopulationStatusEnum_UserListPrepopulationStatus_value)
	proto.RegisterType((*UserListPrepopulationStatusEnum)(nil), "google.ads.googleads.v1.enums.UserListPrepopulationStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_list_prepopulation_status.proto", fileDescriptor_8f4f54f43cc394fe)
}

var fileDescriptor_8f4f54f43cc394fe = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4e, 0xb3, 0x40,
	0x18, 0xfc, 0xa1, 0x7f, 0xaa, 0x6e, 0x35, 0x12, 0x8e, 0x6a, 0xa3, 0xed, 0x03, 0x2c, 0x21, 0xde,
	0xd6, 0x13, 0x95, 0x6d, 0x25, 0x36, 0x88, 0x22, 0x35, 0x31, 0x24, 0x0d, 0x0a, 0x21, 0x24, 0x74,
	0x77, 0xc3, 0xb7, 0xf4, 0x15, 0x7c, 0x0f, 0x8f, 0x3e, 0x8a, 0x8f, 0xe2, 0xdd, 0xbb, 0x61, 0x69,
	0x1b, 0x2f, 0x72, 0xd9, 0x4c, 0xf6, 0x9b, 0x6f, 0xe6, 0x9b, 0x41, 0x93, 0x9c, 0xf3, 0xbc, 0xcc,
	0xac, 0x24, 0x05, 0xab, 0x85, 0x0d, 0x5a, 0xdb, 0x56, 0xc6, 0xea, 0x15, 0x58, 0x35, 0x64, 0xd5,
	0xb2, 0x2c, 0x40, 0x2e, 0x45, 0x95, 0x09, 0x2e, 0xea, 0x32, 0x91, 0x05, 0x67, 0x4b, 0x90, 0x89,
	0xac, 0x01, 0x8b, 0x8a, 0x4b, 0x6e, 0x0e, 0xdb, 0x45, 0x9c, 0xa4, 0x80, 0x77, 0x1a, 0x78, 0x6d,
	0x63, 0xa5, 0x71, 0x72, 0xb6, 0xb5, 0x10, 0x85, 0x95, 0x30, 0xc6, 0xa5, 0x92, 0xd8, 0x2c, 0x8f,
	0xdf, 0x34, 0x74, 0x1e, 0x41, 0x56, 0xcd, 0x0b, 0x90, 0xc1, 0x6f, 0x8f, 0x50, 0x59, 0x50, 0x56,
	0xaf, 0xc6, 0x29, 0x3a, 0xed, 0xa0, 0x98, 0xc7, 0x68, 0x10, 0xf9, 0x61, 0x40, 0xaf, 0xbd, 0xa9,
	0x47, 0x5d, 0xe3, 0x9f, 0x39, 0x40, 0x7b, 0x91, 0x7f, 0xeb, 0xdf, 0x3d, 0xf9, 0x86, 0x66, 0x1e,
	0xa1, 0x83, 0x07, 0x7a, 0x1f, 0xd1, 0xf0, 0x91, 0xba, 0x86, 0x6e, 0x1e, 0xa2, 0xfd, 0xa9, 0xe7,
	0x7b, 0xe1, 0x0d, 0x75, 0x8d, 0x9e, 0x89, 0x50, 0x7f, 0xea, 0x78, 0x73, 0xea, 0x1a, 0xff, 0x27,
	0xdf, 0x1a, 0x1a, 0xbd, 0xf2, 0x15, 0xee, 0x4c, 0x33, 0xb9, 0xe8, 0xb8, 0x24, 0x68, 0x12, 0x05,
	0xda, 0xf3, 0xa6, 0x54, 0x9c, 0xf3, 0x32, 0x61, 0x39, 0xe6, 0x55, 0x6e, 0xe5, 0x19, 0x53, 0x79,
	0xb7, 0x25, 0x8b, 0x02, 0xfe, 0xe8, 0xfc, 0x4a, 0xbd, 0xef, 0x7a, 0x6f, 0xe6, 0x38, 0x1f, 0xfa,
	0x70, 0xd6, 0x4a, 0x39, 0x29, 0xe0, 0x16, 0x36, 0x68, 0x61, 0xe3, 0xa6, 0x18, 0xf8, 0xdc, 0xce,
	0x63, 0x27, 0x85, 0x78, 0x37, 0x8f, 0x17, 0x76, 0xac, 0xe6, 0x5f, 0xfa, 0xa8, 0xfd, 0x24, 0xc4,
	0x49, 0x81, 0x90, 0x1d, 0x83, 0x90, 0x85, 0x4d, 0x88, 0xe2, 0xbc, 0xf4, 0xd5, 0x61, 0x97, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0x85, 0x4f, 0x9a, 0x0b, 0x02, 0x00, 0x00,
}
