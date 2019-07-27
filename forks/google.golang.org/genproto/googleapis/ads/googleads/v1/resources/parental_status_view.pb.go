// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/parental_status_view.proto

package resources

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

// A parental status view.
type ParentalStatusView struct {
	// The resource name of the parental status view.
	// Parental Status view resource names have the form:
	//
	// `customers/{customer_id}/parentalStatusViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParentalStatusView) Reset()         { *m = ParentalStatusView{} }
func (m *ParentalStatusView) String() string { return proto.CompactTextString(m) }
func (*ParentalStatusView) ProtoMessage()    {}
func (*ParentalStatusView) Descriptor() ([]byte, []int) {
	return fileDescriptor_701a359e48c28a79, []int{0}
}

func (m *ParentalStatusView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentalStatusView.Unmarshal(m, b)
}
func (m *ParentalStatusView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentalStatusView.Marshal(b, m, deterministic)
}
func (m *ParentalStatusView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentalStatusView.Merge(m, src)
}
func (m *ParentalStatusView) XXX_Size() int {
	return xxx_messageInfo_ParentalStatusView.Size(m)
}
func (m *ParentalStatusView) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentalStatusView.DiscardUnknown(m)
}

var xxx_messageInfo_ParentalStatusView proto.InternalMessageInfo

func (m *ParentalStatusView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ParentalStatusView)(nil), "google.ads.googleads.v1.resources.ParentalStatusView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/parental_status_view.proto", fileDescriptor_701a359e48c28a79)
}

var fileDescriptor_701a359e48c28a79 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0xec, 0x30,
	0x14, 0xc6, 0x69, 0x2f, 0x5c, 0xb0, 0xe8, 0xa6, 0x1b, 0x45, 0x5c, 0x38, 0xca, 0x80, 0xab, 0x84,
	0xe2, 0xca, 0xe8, 0xa6, 0xb3, 0x19, 0x70, 0x21, 0x65, 0x84, 0x2e, 0xa4, 0x50, 0x8e, 0xd3, 0x10,
	0x0a, 0x6d, 0x4e, 0xc9, 0xc9, 0x74, 0x5e, 0xc0, 0x27, 0x71, 0xe9, 0xa3, 0xf8, 0x28, 0x3e, 0x85,
	0xb4, 0x99, 0x64, 0x23, 0xe8, 0xee, 0x23, 0xf9, 0x7d, 0x7f, 0x38, 0xc9, 0x83, 0x42, 0x54, 0x9d,
	0xe4, 0xd0, 0x10, 0x77, 0x72, 0x52, 0x63, 0xc6, 0x8d, 0x24, 0xdc, 0x99, 0xad, 0x24, 0x3e, 0x80,
	0x91, 0xda, 0x42, 0x57, 0x93, 0x05, 0xbb, 0xa3, 0x7a, 0x6c, 0xe5, 0x9e, 0x0d, 0x06, 0x2d, 0xa6,
	0x0b, 0x67, 0x61, 0xd0, 0x10, 0x0b, 0x6e, 0x36, 0x66, 0x2c, 0xb8, 0xcf, 0x2f, 0x7c, 0xc1, 0xd0,
	0x72, 0xd0, 0x1a, 0x2d, 0xd8, 0x16, 0x35, 0xb9, 0x80, 0xab, 0xbb, 0x24, 0x2d, 0x0e, 0xf1, 0xcf,
	0x73, 0x7a, 0xd9, 0xca, 0x7d, 0x7a, 0x9d, 0x9c, 0xf8, 0x80, 0x5a, 0x43, 0x2f, 0xcf, 0xa2, 0xcb,
	0xe8, 0xe6, 0x68, 0x73, 0xec, 0x1f, 0x9f, 0xa0, 0x97, 0xab, 0xb7, 0x38, 0x59, 0x6e, 0xb1, 0x67,
	0x7f, 0x4e, 0x58, 0x9d, 0xfe, 0xac, 0x28, 0xa6, 0xf6, 0x22, 0x7a, 0x79, 0x3c, 0xb8, 0x15, 0x76,
	0xa0, 0x15, 0x43, 0xa3, 0xb8, 0x92, 0x7a, 0xde, 0xe6, 0xcf, 0x31, 0xb4, 0xf4, 0xcb, 0x75, 0xee,
	0x83, 0x7a, 0x8f, 0xff, 0xad, 0xf3, 0xfc, 0x23, 0x5e, 0xac, 0x5d, 0x64, 0xde, 0x10, 0x73, 0x72,
	0x52, 0x65, 0xc6, 0x36, 0x9e, 0xfc, 0xf4, 0x4c, 0x95, 0x37, 0x54, 0x05, 0xa6, 0x2a, 0xb3, 0x2a,
	0x30, 0x5f, 0xf1, 0xd2, 0x7d, 0x08, 0x91, 0x37, 0x24, 0x44, 0xa0, 0x84, 0x28, 0x33, 0x21, 0x02,
	0xf7, 0xfa, 0x7f, 0x1e, 0x7b, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x98, 0xeb, 0x23, 0x15, 0xc9,
	0x01, 0x00, 0x00,
}
