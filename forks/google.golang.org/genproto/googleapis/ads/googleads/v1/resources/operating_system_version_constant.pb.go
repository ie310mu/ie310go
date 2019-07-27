// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/operating_system_version_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	wrappers "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/wrappers"
	enums "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A mobile operating system version or a range of versions, depending on
// 'operator_type'. The complete list of available mobile platforms is available
// <a
//
// href="https://developers.google.com/adwords/api/docs/appendix/codes-formats#mobile-platforms>
// here</a>.
type OperatingSystemVersionConstant struct {
	// The resource name of the operating system version constant.
	// Operating system version constant resource names have the form:
	//
	// `operatingSystemVersionConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the operating system version.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the operating system.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The OS Major Version number.
	OsMajorVersion *wrappers.Int32Value `protobuf:"bytes,4,opt,name=os_major_version,json=osMajorVersion,proto3" json:"os_major_version,omitempty"`
	// The OS Minor Version number.
	OsMinorVersion *wrappers.Int32Value `protobuf:"bytes,5,opt,name=os_minor_version,json=osMinorVersion,proto3" json:"os_minor_version,omitempty"`
	// Determines whether this constant represents a single version or a range of
	// versions.
	OperatorType         enums.OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType `protobuf:"varint,6,opt,name=operator_type,json=operatorType,proto3,enum=google.ads.googleads.v1.enums.OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType" json:"operator_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                        `json:"-"`
	XXX_unrecognized     []byte                                                                          `json:"-"`
	XXX_sizecache        int32                                                                           `json:"-"`
}

func (m *OperatingSystemVersionConstant) Reset()         { *m = OperatingSystemVersionConstant{} }
func (m *OperatingSystemVersionConstant) String() string { return proto.CompactTextString(m) }
func (*OperatingSystemVersionConstant) ProtoMessage()    {}
func (*OperatingSystemVersionConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_255222b6c86bb69b, []int{0}
}

func (m *OperatingSystemVersionConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatingSystemVersionConstant.Unmarshal(m, b)
}
func (m *OperatingSystemVersionConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatingSystemVersionConstant.Marshal(b, m, deterministic)
}
func (m *OperatingSystemVersionConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatingSystemVersionConstant.Merge(m, src)
}
func (m *OperatingSystemVersionConstant) XXX_Size() int {
	return xxx_messageInfo_OperatingSystemVersionConstant.Size(m)
}
func (m *OperatingSystemVersionConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatingSystemVersionConstant.DiscardUnknown(m)
}

var xxx_messageInfo_OperatingSystemVersionConstant proto.InternalMessageInfo

func (m *OperatingSystemVersionConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *OperatingSystemVersionConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *OperatingSystemVersionConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OperatingSystemVersionConstant) GetOsMajorVersion() *wrappers.Int32Value {
	if m != nil {
		return m.OsMajorVersion
	}
	return nil
}

func (m *OperatingSystemVersionConstant) GetOsMinorVersion() *wrappers.Int32Value {
	if m != nil {
		return m.OsMinorVersion
	}
	return nil
}

func (m *OperatingSystemVersionConstant) GetOperatorType() enums.OperatingSystemVersionOperatorTypeEnum_OperatingSystemVersionOperatorType {
	if m != nil {
		return m.OperatorType
	}
	return enums.OperatingSystemVersionOperatorTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*OperatingSystemVersionConstant)(nil), "google.ads.googleads.v1.resources.OperatingSystemVersionConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/operating_system_version_constant.proto", fileDescriptor_255222b6c86bb69b)
}

var fileDescriptor_255222b6c86bb69b = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x49, 0xb6, 0x16, 0x8c, 0x6d, 0x91, 0x5c, 0x2d, 0xb5, 0x94, 0xad, 0xa5, 0xb0, 0x20,
	0x4c, 0xcc, 0x56, 0xbc, 0x88, 0x57, 0xa9, 0x94, 0x52, 0x41, 0x5b, 0xb6, 0x92, 0x0b, 0x09, 0x84,
	0xe9, 0x66, 0x1c, 0x47, 0x36, 0x73, 0x86, 0x99, 0xc9, 0xca, 0x3e, 0x84, 0xe0, 0x33, 0x78, 0xe9,
	0x63, 0x78, 0xe9, 0xa3, 0xf8, 0x14, 0x92, 0xf9, 0x13, 0x2a, 0xba, 0xdb, 0xbd, 0x3b, 0x93, 0xf3,
	0x3b, 0xdf, 0x37, 0x67, 0xce, 0x49, 0x74, 0x49, 0x01, 0xe8, 0x9c, 0x24, 0xb8, 0x56, 0x89, 0x0d,
	0xbb, 0x68, 0x91, 0x26, 0x92, 0x28, 0x68, 0xe5, 0x8c, 0xa8, 0x04, 0x04, 0x91, 0x58, 0x33, 0x4e,
	0x2b, 0xb5, 0x54, 0x9a, 0x34, 0xd5, 0x82, 0x48, 0xc5, 0x80, 0x57, 0x33, 0xe0, 0x4a, 0x63, 0xae,
	0x91, 0x90, 0xa0, 0x21, 0x3e, 0xb2, 0xf5, 0x08, 0xd7, 0x0a, 0xf5, 0x52, 0x68, 0x91, 0xa2, 0x5e,
	0x6a, 0xff, 0xcd, 0x2a, 0x37, 0xc2, 0xdb, 0x66, 0x8d, 0x93, 0x4d, 0x80, 0xac, 0xf4, 0x52, 0x10,
	0x6b, 0xb7, 0x7f, 0xe0, 0xb5, 0x04, 0x4b, 0x30, 0xe7, 0xa0, 0xb1, 0x66, 0xc0, 0x95, 0xcb, 0x1e,
	0xba, 0xac, 0x39, 0xdd, 0xb6, 0x1f, 0x93, 0x2f, 0x12, 0x0b, 0x41, 0xa4, 0xcb, 0x3f, 0xfd, 0x39,
	0x88, 0x0e, 0xaf, 0xbc, 0xdd, 0x8d, 0x71, 0x2b, 0xac, 0xd9, 0x6b, 0xd7, 0x55, 0x7c, 0x1c, 0xed,
	0xfa, 0x9b, 0x57, 0x1c, 0x37, 0x64, 0x18, 0x8c, 0x82, 0xf1, 0xc3, 0xe9, 0x8e, 0xff, 0xf8, 0x0e,
	0x37, 0x24, 0x7e, 0x16, 0x85, 0xac, 0x1e, 0x86, 0xa3, 0x60, 0xfc, 0x68, 0xf2, 0xc4, 0xb5, 0x8d,
	0xbc, 0x29, 0xba, 0xe4, 0xfa, 0xe5, 0x8b, 0x02, 0xcf, 0x5b, 0x32, 0x0d, 0x59, 0x1d, 0x3f, 0x8f,
	0xb6, 0x8c, 0xd0, 0xc0, 0xe0, 0x07, 0xff, 0xe0, 0x37, 0x5a, 0x32, 0x4e, 0x2d, 0x6f, 0xc8, 0xf8,
	0x3c, 0x7a, 0x0c, 0xaa, 0x6a, 0xf0, 0x67, 0x90, 0xfe, 0x31, 0x86, 0x5b, 0xab, 0xcd, 0x4e, 0x27,
	0xb6, 0x78, 0x0f, 0xd4, 0xdb, 0xae, 0xc6, 0xb5, 0xe4, 0x65, 0x18, 0xbf, 0x23, 0xf3, 0x60, 0x33,
	0x99, 0xae, 0xc6, 0xcb, 0x7c, 0x0d, 0xa2, 0xdd, 0xbf, 0x46, 0x31, 0xdc, 0x1e, 0x05, 0xe3, 0xbd,
	0xc9, 0x27, 0xb4, 0x6a, 0xf4, 0x66, 0xae, 0xe8, 0xff, 0x0f, 0x7d, 0xe5, 0x94, 0xde, 0x2f, 0x05,
	0x39, 0xe7, 0x6d, 0xb3, 0x01, 0x36, 0xdd, 0x81, 0x3b, 0xa7, 0xb3, 0x6f, 0x61, 0x74, 0x32, 0x83,
	0x06, 0xdd, 0xbb, 0x78, 0x67, 0xc7, 0xeb, 0x67, 0x7d, 0xdd, 0x35, 0x7f, 0x1d, 0x7c, 0x70, 0xfb,
	0x89, 0x28, 0xcc, 0x31, 0xa7, 0x08, 0x24, 0x4d, 0x28, 0xe1, 0xe6, 0x69, 0xfc, 0xbe, 0x0a, 0xa6,
	0xd6, 0xfc, 0x2c, 0xaf, 0xfa, 0xe8, 0x7b, 0x38, 0xb8, 0xc8, 0xf3, 0x1f, 0xe1, 0xd1, 0x85, 0x95,
	0xcc, 0x6b, 0x85, 0x6c, 0xd8, 0x45, 0x45, 0x8a, 0xa6, 0x9e, 0xfc, 0xe5, 0x99, 0x32, 0xaf, 0x55,
	0xd9, 0x33, 0x65, 0x91, 0x96, 0x3d, 0xf3, 0x3b, 0x3c, 0xb1, 0x89, 0x2c, 0xcb, 0x6b, 0x95, 0x65,
	0x3d, 0x95, 0x65, 0x45, 0x9a, 0x65, 0x3d, 0x77, 0xbb, 0x6d, 0x2e, 0x7b, 0xfa, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0x9a, 0x43, 0x8c, 0xd8, 0x03, 0x00, 0x00,
}
