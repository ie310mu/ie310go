// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/expanded_landing_page_view.proto

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

// A landing page view with metrics aggregated at the expanded final URL
// level.
type ExpandedLandingPageView struct {
	// The resource name of the expanded landing page view.
	// Expanded landing page view resource names have the form:
	//
	//
	// `customers/{customer_id}/expandedLandingPageViews/{expanded_final_url_fingerprint}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The final URL that clicks are directed to.
	ExpandedFinalUrl     *wrappers.StringValue `protobuf:"bytes,2,opt,name=expanded_final_url,json=expandedFinalUrl,proto3" json:"expanded_final_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExpandedLandingPageView) Reset()         { *m = ExpandedLandingPageView{} }
func (m *ExpandedLandingPageView) String() string { return proto.CompactTextString(m) }
func (*ExpandedLandingPageView) ProtoMessage()    {}
func (*ExpandedLandingPageView) Descriptor() ([]byte, []int) {
	return fileDescriptor_b526ad7687a265f6, []int{0}
}

func (m *ExpandedLandingPageView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpandedLandingPageView.Unmarshal(m, b)
}
func (m *ExpandedLandingPageView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpandedLandingPageView.Marshal(b, m, deterministic)
}
func (m *ExpandedLandingPageView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpandedLandingPageView.Merge(m, src)
}
func (m *ExpandedLandingPageView) XXX_Size() int {
	return xxx_messageInfo_ExpandedLandingPageView.Size(m)
}
func (m *ExpandedLandingPageView) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpandedLandingPageView.DiscardUnknown(m)
}

var xxx_messageInfo_ExpandedLandingPageView proto.InternalMessageInfo

func (m *ExpandedLandingPageView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ExpandedLandingPageView) GetExpandedFinalUrl() *wrappers.StringValue {
	if m != nil {
		return m.ExpandedFinalUrl
	}
	return nil
}

func init() {
	proto.RegisterType((*ExpandedLandingPageView)(nil), "google.ads.googleads.v1.resources.ExpandedLandingPageView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/expanded_landing_page_view.proto", fileDescriptor_b526ad7687a265f6)
}

var fileDescriptor_b526ad7687a265f6 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0xb9, 0x70, 0xe1, 0xe6, 0x2a, 0x48, 0x36, 0x96, 0x52, 0xa4, 0x55, 0x0a, 0x5d, 0x4d,
	0x88, 0xee, 0xc6, 0x55, 0x0a, 0x5a, 0x28, 0x22, 0xa5, 0x62, 0x16, 0x12, 0x08, 0xa7, 0x9d, 0xd3,
	0x61, 0x20, 0x9d, 0x09, 0x33, 0x49, 0xeb, 0x2b, 0xe8, 0x63, 0xb8, 0xf4, 0x51, 0x7c, 0x14, 0x9f,
	0x42, 0xd2, 0x64, 0x66, 0xa7, 0xee, 0xbe, 0x99, 0xf3, 0xfd, 0x1d, 0x4e, 0x30, 0xe5, 0x4a, 0xf1,
	0x02, 0x23, 0x60, 0x26, 0x6a, 0x61, 0x83, 0x76, 0x71, 0xa4, 0xd1, 0xa8, 0x5a, 0xaf, 0xd1, 0x44,
	0xf8, 0x5c, 0x82, 0x64, 0xc8, 0xf2, 0x02, 0x24, 0x13, 0x92, 0xe7, 0x25, 0x70, 0xcc, 0x77, 0x02,
	0xf7, 0xa4, 0xd4, 0xaa, 0x52, 0xe1, 0xa8, 0x15, 0x12, 0x60, 0x86, 0x38, 0x0f, 0xb2, 0x8b, 0x89,
	0xf3, 0xe8, 0x0f, 0x6c, 0x4c, 0x29, 0x22, 0x90, 0x52, 0x55, 0x50, 0x09, 0x25, 0x4d, 0x6b, 0xd0,
	0x3f, 0xeb, 0xa6, 0x87, 0xd7, 0xaa, 0xde, 0x44, 0x7b, 0x0d, 0x65, 0x89, 0xba, 0x9b, 0x9f, 0xbf,
	0x7a, 0xc1, 0xe9, 0x4d, 0xd7, 0xe2, 0xae, 0x2d, 0xb1, 0x00, 0x8e, 0xa9, 0xc0, 0x7d, 0x78, 0x11,
	0x1c, 0xdb, 0x98, 0x5c, 0xc2, 0x16, 0x7b, 0xde, 0xd0, 0x9b, 0xfc, 0x5b, 0x1e, 0xd9, 0xcf, 0x7b,
	0xd8, 0x62, 0x38, 0x0f, 0x42, 0xb7, 0xc5, 0x46, 0x48, 0x28, 0xf2, 0x5a, 0x17, 0x3d, 0x7f, 0xe8,
	0x4d, 0xfe, 0x5f, 0x0e, 0xba, 0xce, 0xc4, 0xa6, 0x93, 0x87, 0x4a, 0x0b, 0xc9, 0x53, 0x28, 0x6a,
	0x5c, 0x9e, 0x58, 0xdd, 0x6d, 0x23, 0x7b, 0xd4, 0xc5, 0xf4, 0xc5, 0x0f, 0xc6, 0x6b, 0xb5, 0x25,
	0xbf, 0x2e, 0x3d, 0x1d, 0x7c, 0xd3, 0x79, 0xd1, 0x04, 0x2d, 0xbc, 0xa7, 0x79, 0x67, 0xc1, 0x55,
	0x01, 0x92, 0x13, 0xa5, 0x79, 0xc4, 0x51, 0x1e, 0x6a, 0xd8, 0x5b, 0x94, 0xc2, 0xfc, 0x70, 0x9a,
	0x6b, 0x87, 0xde, 0xfc, 0x3f, 0xb3, 0x24, 0x79, 0xf7, 0x47, 0xb3, 0xd6, 0x32, 0x61, 0x86, 0xb4,
	0xb0, 0x41, 0x69, 0x4c, 0x96, 0x96, 0xf9, 0x61, 0x39, 0x59, 0xc2, 0x4c, 0xe6, 0x38, 0x59, 0x1a,
	0x67, 0x8e, 0xf3, 0xe9, 0x8f, 0xdb, 0x01, 0xa5, 0x09, 0x33, 0x94, 0x3a, 0x16, 0xa5, 0x69, 0x4c,
	0xa9, 0xe3, 0xad, 0xfe, 0x1e, 0xca, 0x5e, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x16, 0x32, 0xb4,
	0x08, 0x46, 0x02, 0x00, 0x00,
}
