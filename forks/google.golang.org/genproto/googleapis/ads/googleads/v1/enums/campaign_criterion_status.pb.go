// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/campaign_criterion_status.proto

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

// The possible statuses of a CampaignCriterion.
type CampaignCriterionStatusEnum_CampaignCriterionStatus int32

const (
	// No value has been specified.
	CampaignCriterionStatusEnum_UNSPECIFIED CampaignCriterionStatusEnum_CampaignCriterionStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	CampaignCriterionStatusEnum_UNKNOWN CampaignCriterionStatusEnum_CampaignCriterionStatus = 1
	// The campaign criterion is enabled.
	CampaignCriterionStatusEnum_ENABLED CampaignCriterionStatusEnum_CampaignCriterionStatus = 2
	// The campaign criterion is paused.
	CampaignCriterionStatusEnum_PAUSED CampaignCriterionStatusEnum_CampaignCriterionStatus = 3
	// The campaign criterion is removed.
	CampaignCriterionStatusEnum_REMOVED CampaignCriterionStatusEnum_CampaignCriterionStatus = 4
)

var CampaignCriterionStatusEnum_CampaignCriterionStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var CampaignCriterionStatusEnum_CampaignCriterionStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x CampaignCriterionStatusEnum_CampaignCriterionStatus) String() string {
	return proto.EnumName(CampaignCriterionStatusEnum_CampaignCriterionStatus_name, int32(x))
}

func (CampaignCriterionStatusEnum_CampaignCriterionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d7780eca0a0bfd7, []int{0, 0}
}

// Message describing CampaignCriterion statuses.
type CampaignCriterionStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignCriterionStatusEnum) Reset()         { *m = CampaignCriterionStatusEnum{} }
func (m *CampaignCriterionStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionStatusEnum) ProtoMessage()    {}
func (*CampaignCriterionStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d7780eca0a0bfd7, []int{0}
}

func (m *CampaignCriterionStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionStatusEnum.Unmarshal(m, b)
}
func (m *CampaignCriterionStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionStatusEnum.Merge(m, src)
}
func (m *CampaignCriterionStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionStatusEnum.Size(m)
}
func (m *CampaignCriterionStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.CampaignCriterionStatusEnum_CampaignCriterionStatus", CampaignCriterionStatusEnum_CampaignCriterionStatus_name, CampaignCriterionStatusEnum_CampaignCriterionStatus_value)
	proto.RegisterType((*CampaignCriterionStatusEnum)(nil), "google.ads.googleads.v1.enums.CampaignCriterionStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/campaign_criterion_status.proto", fileDescriptor_0d7780eca0a0bfd7)
}

var fileDescriptor_0d7780eca0a0bfd7 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xb1, 0x6a, 0xf3, 0x30,
	0x18, 0xfc, 0xe3, 0xfc, 0xa4, 0xa0, 0x0c, 0x35, 0x5e, 0x0a, 0x6d, 0x32, 0x24, 0x0f, 0x20, 0x63,
	0xba, 0xa9, 0x74, 0x90, 0x63, 0x35, 0x84, 0xb6, 0x8e, 0x69, 0x88, 0x0b, 0xc5, 0x25, 0xa8, 0xb1,
	0x11, 0x82, 0x58, 0x32, 0x96, 0x92, 0xa9, 0x4f, 0xd3, 0xb1, 0x8f, 0xd2, 0x47, 0xe9, 0xd6, 0x37,
	0x28, 0x92, 0xe3, 0x6c, 0xee, 0x22, 0x4e, 0xdf, 0x7d, 0x77, 0x7c, 0x77, 0xe0, 0x96, 0x49, 0xc9,
	0x76, 0x85, 0x4f, 0x73, 0xe5, 0x37, 0xd0, 0xa0, 0x43, 0xe0, 0x17, 0x62, 0x5f, 0x2a, 0x7f, 0x4b,
	0xcb, 0x8a, 0x72, 0x26, 0x36, 0xdb, 0x9a, 0xeb, 0xa2, 0xe6, 0x52, 0x6c, 0x94, 0xa6, 0x7a, 0xaf,
	0x60, 0x55, 0x4b, 0x2d, 0xbd, 0x71, 0xa3, 0x81, 0x34, 0x57, 0xf0, 0x24, 0x87, 0x87, 0x00, 0x5a,
	0xf9, 0xe5, 0xa8, 0x75, 0xaf, 0xb8, 0x4f, 0x85, 0x90, 0x9a, 0x6a, 0x2e, 0xc5, 0x51, 0x3c, 0x7d,
	0x07, 0x57, 0xb3, 0xa3, 0xff, 0xac, 0xb5, 0x5f, 0x59, 0x77, 0x22, 0xf6, 0xe5, 0xf4, 0x15, 0x5c,
	0x74, 0xd0, 0xde, 0x39, 0x18, 0xae, 0xe3, 0x55, 0x42, 0x66, 0x8b, 0xbb, 0x05, 0x89, 0xdc, 0x7f,
	0xde, 0x10, 0x9c, 0xad, 0xe3, 0xfb, 0x78, 0xf9, 0x1c, 0xbb, 0x3d, 0xf3, 0x21, 0x31, 0x0e, 0x1f,
	0x48, 0xe4, 0x3a, 0x1e, 0x00, 0x83, 0x04, 0xaf, 0x57, 0x24, 0x72, 0xfb, 0x86, 0x78, 0x22, 0x8f,
	0xcb, 0x94, 0x44, 0xee, 0xff, 0xf0, 0xa7, 0x07, 0x26, 0x5b, 0x59, 0xc2, 0x3f, 0x13, 0x84, 0xa3,
	0x8e, 0x13, 0x12, 0x93, 0x20, 0xe9, 0xbd, 0x84, 0x47, 0x39, 0x93, 0x3b, 0x2a, 0x18, 0x94, 0x35,
	0xf3, 0x59, 0x21, 0x6c, 0xbe, 0xb6, 0xcf, 0x8a, 0xab, 0x8e, 0x7a, 0x6f, 0xec, 0xfb, 0xe1, 0xf4,
	0xe7, 0x18, 0x7f, 0x3a, 0xe3, 0x79, 0x63, 0x85, 0x73, 0x05, 0x1b, 0x68, 0x50, 0x1a, 0x40, 0xd3,
	0x86, 0xfa, 0x6a, 0xf9, 0x0c, 0xe7, 0x2a, 0x3b, 0xf1, 0x59, 0x1a, 0x64, 0x96, 0xff, 0x76, 0x26,
	0xcd, 0x10, 0x21, 0x9c, 0x2b, 0x84, 0x4e, 0x1b, 0x08, 0xa5, 0x01, 0x42, 0x76, 0xe7, 0x6d, 0x60,
	0x0f, 0xbb, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x21, 0x07, 0xad, 0xc8, 0xf6, 0x01, 0x00, 0x00,
}
