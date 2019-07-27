// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_interest_taxonomy_type.proto

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

// Enum containing the possible UserInterestTaxonomyTypes.
type UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType int32

const (
	// Not specified.
	UserInterestTaxonomyTypeEnum_UNSPECIFIED UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 0
	// Used for return value only. Represents value unknown in this version.
	UserInterestTaxonomyTypeEnum_UNKNOWN UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 1
	// The affinity for this user interest.
	UserInterestTaxonomyTypeEnum_AFFINITY UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 2
	// The market for this user interest.
	UserInterestTaxonomyTypeEnum_IN_MARKET UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 3
	// Users known to have installed applications in the specified categories.
	UserInterestTaxonomyTypeEnum_MOBILE_APP_INSTALL_USER UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 4
	// The geographical location of the interest-based vertical.
	UserInterestTaxonomyTypeEnum_VERTICAL_GEO UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 5
	// User interest criteria for new smart phone users.
	UserInterestTaxonomyTypeEnum_NEW_SMART_PHONE_USER UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType = 6
)

var UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AFFINITY",
	3: "IN_MARKET",
	4: "MOBILE_APP_INSTALL_USER",
	5: "VERTICAL_GEO",
	6: "NEW_SMART_PHONE_USER",
}

var UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"AFFINITY":                2,
	"IN_MARKET":               3,
	"MOBILE_APP_INSTALL_USER": 4,
	"VERTICAL_GEO":            5,
	"NEW_SMART_PHONE_USER":    6,
}

func (x UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) String() string {
	return proto.EnumName(UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_name, int32(x))
}

func (UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb8bc4f7b2be793c, []int{0, 0}
}

// Message describing a UserInterestTaxonomyType.
type UserInterestTaxonomyTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInterestTaxonomyTypeEnum) Reset()         { *m = UserInterestTaxonomyTypeEnum{} }
func (m *UserInterestTaxonomyTypeEnum) String() string { return proto.CompactTextString(m) }
func (*UserInterestTaxonomyTypeEnum) ProtoMessage()    {}
func (*UserInterestTaxonomyTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb8bc4f7b2be793c, []int{0}
}

func (m *UserInterestTaxonomyTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInterestTaxonomyTypeEnum.Unmarshal(m, b)
}
func (m *UserInterestTaxonomyTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInterestTaxonomyTypeEnum.Marshal(b, m, deterministic)
}
func (m *UserInterestTaxonomyTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInterestTaxonomyTypeEnum.Merge(m, src)
}
func (m *UserInterestTaxonomyTypeEnum) XXX_Size() int {
	return xxx_messageInfo_UserInterestTaxonomyTypeEnum.Size(m)
}
func (m *UserInterestTaxonomyTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInterestTaxonomyTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserInterestTaxonomyTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType", UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_name, UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType_value)
	proto.RegisterType((*UserInterestTaxonomyTypeEnum)(nil), "google.ads.googleads.v1.enums.UserInterestTaxonomyTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_interest_taxonomy_type.proto", fileDescriptor_bb8bc4f7b2be793c)
}

var fileDescriptor_bb8bc4f7b2be793c = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcf, 0x8a, 0xd4, 0x30,
	0x1c, 0xb6, 0xb3, 0xba, 0x6a, 0x76, 0xc5, 0x12, 0x04, 0x17, 0xdd, 0x39, 0xec, 0x3e, 0x40, 0x4a,
	0xf1, 0x16, 0x0f, 0x92, 0xae, 0x99, 0x31, 0x6c, 0x27, 0x2d, 0xfd, 0xb7, 0x28, 0x85, 0x50, 0x6d,
	0x28, 0x85, 0x6d, 0x52, 0x9a, 0xce, 0xe0, 0xbc, 0x8c, 0x07, 0x8f, 0xbe, 0x84, 0x77, 0x1f, 0xc5,
	0xa3, 0x4f, 0x20, 0x6d, 0x67, 0xe6, 0x36, 0x7b, 0x09, 0x1f, 0xf9, 0xfe, 0xfc, 0x92, 0xef, 0x07,
	0x3e, 0x54, 0x5a, 0x57, 0xf7, 0xd2, 0x29, 0x4a, 0xe3, 0x4c, 0x70, 0x40, 0x1b, 0xd7, 0x91, 0x6a,
	0xdd, 0x18, 0x67, 0x6d, 0x64, 0x27, 0x6a, 0xd5, 0xcb, 0x4e, 0x9a, 0x5e, 0xf4, 0xc5, 0x77, 0xad,
	0x74, 0xb3, 0x15, 0xfd, 0xb6, 0x95, 0xa8, 0xed, 0x74, 0xaf, 0xe1, 0x7c, 0x72, 0xa1, 0xa2, 0x34,
	0xe8, 0x10, 0x80, 0x36, 0x2e, 0x1a, 0x03, 0xde, 0x5c, 0xee, 0xf3, 0xdb, 0xda, 0x29, 0x94, 0xd2,
	0x7d, 0xd1, 0xd7, 0x5a, 0x99, 0xc9, 0x7c, 0xfd, 0xdb, 0x02, 0x97, 0xa9, 0x91, 0x1d, 0xdb, 0x4d,
	0x48, 0x76, 0x03, 0x92, 0x6d, 0x2b, 0xa9, 0x5a, 0x37, 0xd7, 0x3f, 0x2c, 0x70, 0x71, 0x4c, 0x00,
	0x5f, 0x82, 0xb3, 0x94, 0xc7, 0x21, 0xbd, 0x61, 0x0b, 0x46, 0x3f, 0xda, 0x8f, 0xe0, 0x19, 0x78,
	0x9a, 0xf2, 0x5b, 0x1e, 0xdc, 0x71, 0xdb, 0x82, 0xe7, 0xe0, 0x19, 0x59, 0x2c, 0x18, 0x67, 0xc9,
	0x67, 0x7b, 0x06, 0x5f, 0x80, 0xe7, 0x8c, 0x8b, 0x15, 0x89, 0x6e, 0x69, 0x62, 0x9f, 0xc0, 0xb7,
	0xe0, 0xf5, 0x2a, 0xf0, 0x98, 0x4f, 0x05, 0x09, 0x43, 0xc1, 0x78, 0x9c, 0x10, 0xdf, 0x17, 0x69,
	0x4c, 0x23, 0xfb, 0x31, 0xb4, 0xc1, 0x79, 0x46, 0xa3, 0x84, 0xdd, 0x10, 0x5f, 0x2c, 0x69, 0x60,
	0x3f, 0x81, 0x17, 0xe0, 0x15, 0xa7, 0x77, 0x22, 0x5e, 0x91, 0x28, 0x11, 0xe1, 0xa7, 0x80, 0xd3,
	0x49, 0x7b, 0xea, 0xfd, 0xb3, 0xc0, 0xd5, 0x37, 0xdd, 0xa0, 0x07, 0x5b, 0xf0, 0xe6, 0xc7, 0xfe,
	0x10, 0x0e, 0x35, 0x84, 0xd6, 0x17, 0x6f, 0xe7, 0xaf, 0xf4, 0x7d, 0xa1, 0x2a, 0xa4, 0xbb, 0xca,
	0xa9, 0xa4, 0x1a, 0x4b, 0xda, 0xaf, 0xa5, 0xad, 0xcd, 0x91, 0x2d, 0xbd, 0x1f, 0xcf, 0x9f, 0xb3,
	0x93, 0x25, 0x21, 0xbf, 0x66, 0xf3, 0xe5, 0x14, 0x45, 0x4a, 0x83, 0x26, 0x38, 0xa0, 0xcc, 0x45,
	0x43, 0xa1, 0xe6, 0xcf, 0x9e, 0xcf, 0x49, 0x69, 0xf2, 0x03, 0x9f, 0x67, 0x6e, 0x3e, 0xf2, 0x7f,
	0x67, 0x57, 0xd3, 0x25, 0xc6, 0xa4, 0x34, 0x18, 0x1f, 0x14, 0x18, 0x67, 0x2e, 0xc6, 0xa3, 0xe6,
	0xeb, 0xe9, 0xf8, 0xb0, 0x77, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x2c, 0x3f, 0x4e, 0x3d,
	0x02, 0x00, 0x00,
}
