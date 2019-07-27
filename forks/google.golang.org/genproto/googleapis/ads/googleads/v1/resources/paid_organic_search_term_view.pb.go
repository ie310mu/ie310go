// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/paid_organic_search_term_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	wrappers "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/wrappers"
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

// A paid organic search term view providing a view of search stats across
// ads and organic listings aggregated by search term at the ad group level.
type PaidOrganicSearchTermView struct {
	// The resource name of the search term view.
	// Search term view resource names have the form:
	//
	// `customers/{customer_id}/paidOrganicSearchTermViews/{campaign_id}~
	// {ad_group_id}~{URL-base64 search term}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The search term.
	SearchTerm           *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PaidOrganicSearchTermView) Reset()         { *m = PaidOrganicSearchTermView{} }
func (m *PaidOrganicSearchTermView) String() string { return proto.CompactTextString(m) }
func (*PaidOrganicSearchTermView) ProtoMessage()    {}
func (*PaidOrganicSearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e209ceb6c96a605, []int{0}
}

func (m *PaidOrganicSearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaidOrganicSearchTermView.Unmarshal(m, b)
}
func (m *PaidOrganicSearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaidOrganicSearchTermView.Marshal(b, m, deterministic)
}
func (m *PaidOrganicSearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaidOrganicSearchTermView.Merge(m, src)
}
func (m *PaidOrganicSearchTermView) XXX_Size() int {
	return xxx_messageInfo_PaidOrganicSearchTermView.Size(m)
}
func (m *PaidOrganicSearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_PaidOrganicSearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_PaidOrganicSearchTermView proto.InternalMessageInfo

func (m *PaidOrganicSearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *PaidOrganicSearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func init() {
	proto.RegisterType((*PaidOrganicSearchTermView)(nil), "google.ads.googleads.v1.resources.PaidOrganicSearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/paid_organic_search_term_view.proto", fileDescriptor_8e209ceb6c96a605)
}

var fileDescriptor_8e209ceb6c96a605 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xf9, 0xe1, 0x07, 0xa7, 0xba, 0xe9, 0xaa, 0x96, 0x52, 0x5a, 0xa5, 0xd0, 0x55,
	0xc2, 0xe8, 0x2e, 0xe2, 0x62, 0x0a, 0x52, 0x70, 0xa1, 0xa5, 0x95, 0x59, 0xc8, 0xc0, 0x90, 0x4e,
	0xae, 0x31, 0xd0, 0x49, 0x86, 0x64, 0xda, 0x2e, 0x7d, 0x00, 0xdf, 0xc2, 0xa5, 0x8f, 0xe2, 0xa3,
	0xf8, 0x14, 0xd2, 0xc9, 0x24, 0xb8, 0x51, 0x77, 0x27, 0xb9, 0x27, 0xe7, 0x7c, 0xe1, 0x46, 0x37,
	0x5c, 0x29, 0xbe, 0x01, 0x4c, 0x99, 0xc1, 0x56, 0x1e, 0xd4, 0x2e, 0xc6, 0x1a, 0x8c, 0xda, 0xea,
	0x02, 0x0c, 0xae, 0xa8, 0x60, 0xb9, 0xd2, 0x9c, 0x4a, 0x51, 0xe4, 0x06, 0xa8, 0x2e, 0x9e, 0xf3,
	0x1a, 0x74, 0x99, 0xef, 0x04, 0xec, 0x51, 0xa5, 0x55, 0xad, 0xba, 0x63, 0xfb, 0x16, 0x51, 0x66,
	0x90, 0x8f, 0x41, 0xbb, 0x18, 0xf9, 0x98, 0xfe, 0xc0, 0x35, 0x55, 0x02, 0x53, 0x29, 0x55, 0x4d,
	0x6b, 0xa1, 0xa4, 0xb1, 0x01, 0xfd, 0x61, 0x3b, 0x6d, 0x4e, 0xeb, 0xed, 0x13, 0xde, 0x6b, 0x5a,
	0x55, 0xa0, 0xdb, 0xf9, 0xd9, 0x4b, 0x74, 0xba, 0xa0, 0x82, 0xdd, 0x5b, 0x8c, 0x55, 0x43, 0xf1,
	0x00, 0xba, 0x4c, 0x05, 0xec, 0xbb, 0xe7, 0xd1, 0x89, 0xeb, 0xc9, 0x25, 0x2d, 0xa1, 0x17, 0x8c,
	0x82, 0xe9, 0xd1, 0xf2, 0xd8, 0x5d, 0xde, 0xd1, 0x12, 0xba, 0xd7, 0x51, 0xe7, 0x1b, 0x7c, 0x2f,
	0x1c, 0x05, 0xd3, 0xce, 0xc5, 0xa0, 0xa5, 0x45, 0xae, 0x17, 0xad, 0x6a, 0x2d, 0x24, 0x4f, 0xe9,
	0x66, 0x0b, 0xcb, 0xc8, 0xf8, 0x9e, 0xd9, 0x6b, 0x18, 0x4d, 0x0a, 0x55, 0xa2, 0x3f, 0x3f, 0x3a,
	0x1b, 0xfe, 0x08, 0xba, 0x38, 0x94, 0x2c, 0x82, 0xc7, 0xdb, 0x36, 0x84, 0xab, 0x0d, 0x95, 0x1c,
	0x29, 0xcd, 0x31, 0x07, 0xd9, 0x20, 0xb8, 0x25, 0x54, 0xc2, 0xfc, 0xb2, 0x93, 0x2b, 0xaf, 0xde,
	0xc2, 0x7f, 0xf3, 0x24, 0x79, 0x0f, 0xc7, 0x73, 0x1b, 0x99, 0x30, 0x83, 0xac, 0x3c, 0xa8, 0x34,
	0x46, 0x4b, 0xe7, 0xfc, 0x70, 0x9e, 0x2c, 0x61, 0x26, 0xf3, 0x9e, 0x2c, 0x8d, 0x33, 0xef, 0xf9,
	0x0c, 0x27, 0x76, 0x40, 0x48, 0xc2, 0x0c, 0x21, 0xde, 0x45, 0x48, 0x1a, 0x13, 0xe2, 0x7d, 0xeb,
	0xff, 0x0d, 0xec, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x85, 0xa5, 0x68, 0x3f, 0x02,
	0x00, 0x00,
}
