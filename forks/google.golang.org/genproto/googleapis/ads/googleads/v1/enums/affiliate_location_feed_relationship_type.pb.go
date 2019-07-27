// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/affiliate_location_feed_relationship_type.proto

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

// Possible values for a relationship type for an affiliate location feed.
type AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType int32

const (
	// Not specified.
	AffiliateLocationFeedRelationshipTypeEnum_UNSPECIFIED AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType = 0
	// Used for return value only. Represents value unknown in this version.
	AffiliateLocationFeedRelationshipTypeEnum_UNKNOWN AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType = 1
	// General retailer relationship.
	AffiliateLocationFeedRelationshipTypeEnum_GENERAL_RETAILER AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType = 2
)

var AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "GENERAL_RETAILER",
}

var AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_value = map[string]int32{
	"UNSPECIFIED":      0,
	"UNKNOWN":          1,
	"GENERAL_RETAILER": 2,
}

func (x AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType) String() string {
	return proto.EnumName(AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_name, int32(x))
}

func (AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7f857a321b243c6c, []int{0, 0}
}

// Container for enum describing possible values for a relationship type for
// an affiliate location feed.
type AffiliateLocationFeedRelationshipTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AffiliateLocationFeedRelationshipTypeEnum) Reset() {
	*m = AffiliateLocationFeedRelationshipTypeEnum{}
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AffiliateLocationFeedRelationshipTypeEnum) ProtoMessage()    {}
func (*AffiliateLocationFeedRelationshipTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f857a321b243c6c, []int{0}
}

func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.Unmarshal(m, b)
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.Marshal(b, m, deterministic)
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.Merge(m, src)
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.Size(m)
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType", AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_name, AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_value)
	proto.RegisterType((*AffiliateLocationFeedRelationshipTypeEnum)(nil), "google.ads.googleads.v1.enums.AffiliateLocationFeedRelationshipTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/affiliate_location_feed_relationship_type.proto", fileDescriptor_7f857a321b243c6c)
}

var fileDescriptor_7f857a321b243c6c = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x6d, 0x05, 0x85, 0xec, 0x60, 0x29, 0x9e, 0xc4, 0x1d, 0x36, 0xf0, 0xa0, 0x87, 0x94,
	0xe2, 0x2d, 0x9e, 0x32, 0xcd, 0xc6, 0x70, 0xd6, 0x51, 0xb7, 0x09, 0x5a, 0x28, 0x71, 0xc9, 0x6a,
	0xa1, 0x4b, 0xc2, 0xd2, 0x0d, 0xf6, 0x06, 0x3e, 0x87, 0x47, 0x1f, 0xc5, 0x47, 0xf1, 0x25, 0x94,
	0x26, 0xeb, 0xf0, 0xa2, 0xec, 0x52, 0xfe, 0xcd, 0xf7, 0xe5, 0xf7, 0xff, 0xfe, 0x5f, 0xc0, 0x5d,
	0x26, 0x65, 0x56, 0xf0, 0x80, 0x32, 0x1d, 0x58, 0x59, 0xa9, 0x55, 0x18, 0x70, 0xb1, 0x9c, 0xeb,
	0x80, 0xce, 0x66, 0x79, 0x91, 0xd3, 0x92, 0xa7, 0x85, 0x9c, 0xd2, 0x32, 0x97, 0x22, 0x9d, 0x71,
	0xce, 0xd2, 0x05, 0x2f, 0xcc, 0x9f, 0x7e, 0xcd, 0x55, 0x5a, 0xae, 0x15, 0x87, 0x6a, 0x21, 0x4b,
	0xe9, 0x37, 0x2d, 0x03, 0x52, 0xa6, 0xe1, 0x16, 0x07, 0x57, 0x21, 0x34, 0xb8, 0x93, 0xd3, 0xda,
	0x4d, 0xe5, 0x01, 0x15, 0x42, 0x96, 0x16, 0x62, 0x2f, 0xb7, 0xdf, 0x1c, 0x70, 0x8e, 0x6b, 0xc3,
	0xc1, 0xc6, 0xaf, 0xcb, 0x39, 0x8b, 0x7f, 0xb9, 0x8d, 0xd6, 0x8a, 0x13, 0xb1, 0x9c, 0xb7, 0x9f,
	0xc1, 0xd9, 0x4e, 0xcd, 0xfe, 0x11, 0x68, 0x8c, 0xa3, 0x87, 0x21, 0xb9, 0xee, 0x77, 0xfb, 0xe4,
	0xc6, 0xdb, 0xf3, 0x1b, 0xe0, 0x70, 0x1c, 0xdd, 0x46, 0xf7, 0x8f, 0x91, 0xe7, 0xf8, 0xc7, 0xc0,
	0xeb, 0x91, 0x88, 0xc4, 0x78, 0x90, 0xc6, 0x64, 0x84, 0xfb, 0x03, 0x12, 0x7b, 0x6e, 0xe7, 0xdb,
	0x01, 0xad, 0xa9, 0x9c, 0xc3, 0x7f, 0xe3, 0x74, 0x2e, 0x76, 0x1a, 0x60, 0x58, 0x85, 0x1b, 0x3a,
	0x4f, 0x9d, 0x0d, 0x2c, 0x93, 0x05, 0x15, 0x19, 0x94, 0x8b, 0x2c, 0xc8, 0xb8, 0x30, 0xd1, 0xeb,
	0xd5, 0xab, 0x5c, 0xff, 0xf1, 0x12, 0x57, 0xe6, 0xfb, 0xee, 0xee, 0xf7, 0x30, 0xfe, 0x70, 0x9b,
	0x3d, 0x8b, 0xc2, 0x4c, 0x43, 0x2b, 0x2b, 0x35, 0x09, 0x61, 0xb5, 0x19, 0xfd, 0x59, 0xd7, 0x13,
	0xcc, 0x74, 0xb2, 0xad, 0x27, 0x93, 0x30, 0x31, 0xf5, 0x2f, 0xb7, 0x65, 0x0f, 0x11, 0xc2, 0x4c,
	0x23, 0xb4, 0xed, 0x40, 0x68, 0x12, 0x22, 0x64, 0x7a, 0x5e, 0x0e, 0xcc, 0x60, 0x97, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x67, 0xb4, 0x2f, 0x6b, 0x21, 0x02, 0x00, 0x00,
}
