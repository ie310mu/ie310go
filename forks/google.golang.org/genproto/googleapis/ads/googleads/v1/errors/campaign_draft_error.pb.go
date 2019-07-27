// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/campaign_draft_error.proto

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

// Enum describing possible campaign draft errors.
type CampaignDraftErrorEnum_CampaignDraftError int32

const (
	// Enum unspecified.
	CampaignDraftErrorEnum_UNSPECIFIED CampaignDraftErrorEnum_CampaignDraftError = 0
	// The received error code is not known in this version.
	CampaignDraftErrorEnum_UNKNOWN CampaignDraftErrorEnum_CampaignDraftError = 1
	// A draft with this name already exists for this campaign.
	CampaignDraftErrorEnum_DUPLICATE_DRAFT_NAME CampaignDraftErrorEnum_CampaignDraftError = 2
	// The draft is removed and cannot be transitioned to another status.
	CampaignDraftErrorEnum_INVALID_STATUS_TRANSITION_FROM_REMOVED CampaignDraftErrorEnum_CampaignDraftError = 3
	// The draft has been promoted and cannot be transitioned to the specified
	// status.
	CampaignDraftErrorEnum_INVALID_STATUS_TRANSITION_FROM_PROMOTED CampaignDraftErrorEnum_CampaignDraftError = 4
	// The draft has failed to be promoted and cannot be transitioned to the
	// specified status.
	CampaignDraftErrorEnum_INVALID_STATUS_TRANSITION_FROM_PROMOTE_FAILED CampaignDraftErrorEnum_CampaignDraftError = 5
	// This customer is not allowed to create drafts.
	CampaignDraftErrorEnum_CUSTOMER_CANNOT_CREATE_DRAFT CampaignDraftErrorEnum_CampaignDraftError = 6
	// This campaign is not allowed to create drafts.
	CampaignDraftErrorEnum_CAMPAIGN_CANNOT_CREATE_DRAFT CampaignDraftErrorEnum_CampaignDraftError = 7
	// This modification cannot be made on a draft.
	CampaignDraftErrorEnum_INVALID_DRAFT_CHANGE CampaignDraftErrorEnum_CampaignDraftError = 8
	// The draft cannot be transitioned to the specified status from its
	// current status.
	CampaignDraftErrorEnum_INVALID_STATUS_TRANSITION CampaignDraftErrorEnum_CampaignDraftError = 9
	// The campaign has reached the maximum number of drafts that can be created
	// for a campaign throughout its lifetime. No additional drafts can be
	// created for this campaign. Removed drafts also count towards this limit.
	CampaignDraftErrorEnum_MAX_NUMBER_OF_DRAFTS_PER_CAMPAIGN_REACHED CampaignDraftErrorEnum_CampaignDraftError = 10
	// ListAsyncErrors was called without first promoting the draft.
	CampaignDraftErrorEnum_LIST_ERRORS_FOR_PROMOTED_DRAFT_ONLY CampaignDraftErrorEnum_CampaignDraftError = 11
)

var CampaignDraftErrorEnum_CampaignDraftError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DUPLICATE_DRAFT_NAME",
	3:  "INVALID_STATUS_TRANSITION_FROM_REMOVED",
	4:  "INVALID_STATUS_TRANSITION_FROM_PROMOTED",
	5:  "INVALID_STATUS_TRANSITION_FROM_PROMOTE_FAILED",
	6:  "CUSTOMER_CANNOT_CREATE_DRAFT",
	7:  "CAMPAIGN_CANNOT_CREATE_DRAFT",
	8:  "INVALID_DRAFT_CHANGE",
	9:  "INVALID_STATUS_TRANSITION",
	10: "MAX_NUMBER_OF_DRAFTS_PER_CAMPAIGN_REACHED",
	11: "LIST_ERRORS_FOR_PROMOTED_DRAFT_ONLY",
}

var CampaignDraftErrorEnum_CampaignDraftError_value = map[string]int32{
	"UNSPECIFIED":                                   0,
	"UNKNOWN":                                       1,
	"DUPLICATE_DRAFT_NAME":                          2,
	"INVALID_STATUS_TRANSITION_FROM_REMOVED":        3,
	"INVALID_STATUS_TRANSITION_FROM_PROMOTED":       4,
	"INVALID_STATUS_TRANSITION_FROM_PROMOTE_FAILED": 5,
	"CUSTOMER_CANNOT_CREATE_DRAFT":                  6,
	"CAMPAIGN_CANNOT_CREATE_DRAFT":                  7,
	"INVALID_DRAFT_CHANGE":                          8,
	"INVALID_STATUS_TRANSITION":                     9,
	"MAX_NUMBER_OF_DRAFTS_PER_CAMPAIGN_REACHED":     10,
	"LIST_ERRORS_FOR_PROMOTED_DRAFT_ONLY":           11,
}

func (x CampaignDraftErrorEnum_CampaignDraftError) String() string {
	return proto.EnumName(CampaignDraftErrorEnum_CampaignDraftError_name, int32(x))
}

func (CampaignDraftErrorEnum_CampaignDraftError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f9d5ffa518a63998, []int{0, 0}
}

// Container for enum describing possible campaign draft errors.
type CampaignDraftErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignDraftErrorEnum) Reset()         { *m = CampaignDraftErrorEnum{} }
func (m *CampaignDraftErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignDraftErrorEnum) ProtoMessage()    {}
func (*CampaignDraftErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9d5ffa518a63998, []int{0}
}

func (m *CampaignDraftErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignDraftErrorEnum.Unmarshal(m, b)
}
func (m *CampaignDraftErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignDraftErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignDraftErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignDraftErrorEnum.Merge(m, src)
}
func (m *CampaignDraftErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignDraftErrorEnum.Size(m)
}
func (m *CampaignDraftErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignDraftErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignDraftErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.CampaignDraftErrorEnum_CampaignDraftError", CampaignDraftErrorEnum_CampaignDraftError_name, CampaignDraftErrorEnum_CampaignDraftError_value)
	proto.RegisterType((*CampaignDraftErrorEnum)(nil), "google.ads.googleads.v1.errors.CampaignDraftErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/campaign_draft_error.proto", fileDescriptor_f9d5ffa518a63998)
}

var fileDescriptor_f9d5ffa518a63998 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xdd, 0x56, 0x77, 0x75, 0x7a, 0xe1, 0x30, 0x88, 0xff, 0x58, 0x17, 0xa9, 0xe0, 0xa2,
	0xb2, 0x09, 0xc1, 0x2b, 0xe3, 0xd5, 0x34, 0x33, 0xe9, 0x0e, 0x36, 0x33, 0x61, 0x32, 0x89, 0x7f,
	0x28, 0x0c, 0x71, 0x53, 0x43, 0x61, 0x9b, 0x29, 0x49, 0xdd, 0x07, 0xf2, 0x52, 0xf0, 0x3d, 0xc4,
	0x47, 0xf1, 0xd2, 0x27, 0x90, 0x64, 0xda, 0xdc, 0xac, 0xab, 0x5e, 0xe5, 0x70, 0xf2, 0xfb, 0xce,
	0xf9, 0x0e, 0xf3, 0x81, 0x57, 0xa5, 0x31, 0xe5, 0xf9, 0xc2, 0xcd, 0x8b, 0xc6, 0xb5, 0x65, 0x5b,
	0x5d, 0x78, 0xee, 0xa2, 0xae, 0x4d, 0xdd, 0xb8, 0x67, 0xf9, 0x6a, 0x9d, 0x2f, 0xcb, 0x4a, 0x17,
	0x75, 0xfe, 0x69, 0xa3, 0xbb, 0xae, 0xb3, 0xae, 0xcd, 0xc6, 0xa0, 0x23, 0xcb, 0x3b, 0x79, 0xd1,
	0x38, 0xbd, 0xd4, 0xb9, 0xf0, 0x1c, 0x2b, 0x7d, 0x78, 0xb8, 0x1b, 0xbd, 0x5e, 0xba, 0x79, 0x55,
	0x99, 0x4d, 0xbe, 0x59, 0x9a, 0xaa, 0xb1, 0xea, 0xf1, 0xf7, 0x21, 0xb8, 0x1b, 0x6c, 0x87, 0x93,
	0x76, 0x36, 0x6d, 0x55, 0xb4, 0xfa, 0xbc, 0x1a, 0x7f, 0x1b, 0x02, 0x74, 0xf9, 0x17, 0xba, 0x0d,
	0x46, 0x29, 0x4f, 0x62, 0x1a, 0xb0, 0x90, 0x51, 0x02, 0xaf, 0xa1, 0x11, 0x38, 0x48, 0xf9, 0x1b,
	0x2e, 0xde, 0x72, 0xb8, 0x87, 0xee, 0x83, 0x3b, 0x24, 0x8d, 0x67, 0x2c, 0xc0, 0x8a, 0x6a, 0x22,
	0x71, 0xa8, 0x34, 0xc7, 0x11, 0x85, 0x03, 0xf4, 0x1c, 0x3c, 0x65, 0x3c, 0xc3, 0x33, 0x46, 0x74,
	0xa2, 0xb0, 0x4a, 0x13, 0xad, 0x24, 0xe6, 0x09, 0x53, 0x4c, 0x70, 0x1d, 0x4a, 0x11, 0x69, 0x49,
	0x23, 0x91, 0x51, 0x02, 0x87, 0xe8, 0x05, 0x38, 0xfe, 0x07, 0x1b, 0x4b, 0x11, 0x09, 0x45, 0x09,
	0xbc, 0x8e, 0x3c, 0x70, 0xf2, 0x7f, 0xb0, 0x0e, 0x31, 0x9b, 0x51, 0x02, 0x6f, 0xa0, 0xc7, 0xe0,
	0x30, 0x48, 0x13, 0x25, 0x22, 0x2a, 0x75, 0x80, 0x39, 0x17, 0x4a, 0x07, 0x92, 0xf6, 0x96, 0xe1,
	0x7e, 0x47, 0xe0, 0x28, 0xc6, 0x6c, 0xca, 0xff, 0x48, 0x1c, 0xb4, 0x97, 0xee, 0xd6, 0xda, 0x3b,
	0x83, 0x53, 0xcc, 0xa7, 0x14, 0xde, 0x44, 0x8f, 0xc0, 0x83, 0x2b, 0x0d, 0xc1, 0x5b, 0xe8, 0x04,
	0x3c, 0x8b, 0xf0, 0x3b, 0xcd, 0xd3, 0x68, 0x42, 0xa5, 0x16, 0xa1, 0x95, 0x27, 0x3a, 0xee, 0xcc,
	0x6c, 0x77, 0x4a, 0x8a, 0x83, 0x53, 0x4a, 0x20, 0x40, 0xc7, 0xe0, 0xc9, 0x8c, 0x25, 0x4a, 0x53,
	0x29, 0x85, 0x4c, 0x74, 0x28, 0x64, 0x7f, 0xfc, 0x76, 0xb1, 0xe0, 0xb3, 0xf7, 0x70, 0x34, 0xf9,
	0xb5, 0x07, 0xc6, 0x67, 0x66, 0xe5, 0xfc, 0x3d, 0x0f, 0x93, 0x7b, 0x97, 0xdf, 0x34, 0x6e, 0xa3,
	0x10, 0xef, 0x7d, 0x20, 0x5b, 0x69, 0x69, 0xce, 0xf3, 0xaa, 0x74, 0x4c, 0x5d, 0xba, 0xe5, 0xa2,
	0xea, 0x82, 0xb2, 0x4b, 0xe5, 0x7a, 0xd9, 0x5c, 0x15, 0xd2, 0xd7, 0xf6, 0xf3, 0x65, 0x30, 0x9c,
	0x62, 0xfc, 0x75, 0x70, 0x34, 0xb5, 0xc3, 0x70, 0xd1, 0x38, 0xb6, 0x6c, 0xab, 0xcc, 0x73, 0xba,
	0x95, 0xcd, 0x8f, 0x1d, 0x30, 0xc7, 0x45, 0x33, 0xef, 0x81, 0x79, 0xe6, 0xcd, 0x2d, 0xf0, 0x73,
	0x30, 0xb6, 0x5d, 0xdf, 0xc7, 0x45, 0xe3, 0xfb, 0x3d, 0xe2, 0xfb, 0x99, 0xe7, 0xfb, 0x16, 0xfa,
	0xb8, 0xdf, 0xb9, 0x7b, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x6b, 0x9e, 0x4f, 0x41, 0x03,
	0x00, 0x00,
}
