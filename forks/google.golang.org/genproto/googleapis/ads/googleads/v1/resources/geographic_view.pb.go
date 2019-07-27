// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/geographic_view.proto

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

// A geographic view.
//
// Geographic View includes all metrics aggregated at the country level,
// one row per country. It reports metrics at either actual physical location of
// the user or an area of interest. If other segment fields are used, you may
// get more than one row per country.
type GeographicView struct {
	// The resource name of the geographic view.
	// Geographic view resource names have the form:
	//
	//
	// `customers/{customer_id}/geographicViews/{country_criterion_id}~{location_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// CriterionId for the geo target for a country.
	CountryGeoTargetConstant *wrappers.StringValue `protobuf:"bytes,2,opt,name=country_geo_target_constant,json=countryGeoTargetConstant,proto3" json:"country_geo_target_constant,omitempty"`
	// Type of the geo targeting of the campaign.
	LocationType         enums.GeoTargetingTypeEnum_GeoTargetingType `protobuf:"varint,3,opt,name=location_type,json=locationType,proto3,enum=google.ads.googleads.v1.enums.GeoTargetingTypeEnum_GeoTargetingType" json:"location_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *GeographicView) Reset()         { *m = GeographicView{} }
func (m *GeographicView) String() string { return proto.CompactTextString(m) }
func (*GeographicView) ProtoMessage()    {}
func (*GeographicView) Descriptor() ([]byte, []int) {
	return fileDescriptor_45a03e49a5e672c5, []int{0}
}

func (m *GeographicView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeographicView.Unmarshal(m, b)
}
func (m *GeographicView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeographicView.Marshal(b, m, deterministic)
}
func (m *GeographicView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeographicView.Merge(m, src)
}
func (m *GeographicView) XXX_Size() int {
	return xxx_messageInfo_GeographicView.Size(m)
}
func (m *GeographicView) XXX_DiscardUnknown() {
	xxx_messageInfo_GeographicView.DiscardUnknown(m)
}

var xxx_messageInfo_GeographicView proto.InternalMessageInfo

func (m *GeographicView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GeographicView) GetCountryGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.CountryGeoTargetConstant
	}
	return nil
}

func (m *GeographicView) GetLocationType() enums.GeoTargetingTypeEnum_GeoTargetingType {
	if m != nil {
		return m.LocationType
	}
	return enums.GeoTargetingTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*GeographicView)(nil), "google.ads.googleads.v1.resources.GeographicView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/geographic_view.proto", fileDescriptor_45a03e49a5e672c5)
}

var fileDescriptor_45a03e49a5e672c5 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x8b, 0xd4, 0x30,
	0x18, 0xa5, 0x5d, 0x10, 0xac, 0xbb, 0x7b, 0xa8, 0x97, 0xb2, 0x2e, 0x32, 0xab, 0x2c, 0xcc, 0x29,
	0xa5, 0x2b, 0x28, 0xc4, 0x53, 0x57, 0xa5, 0xe0, 0x41, 0x96, 0x3a, 0xf4, 0xa0, 0x85, 0x92, 0x69,
	0x3f, 0x63, 0xa0, 0x4d, 0x42, 0x92, 0xce, 0x30, 0x37, 0x7f, 0x8b, 0x47, 0x7f, 0x8a, 0x3f, 0xc5,
	0xff, 0x20, 0x48, 0xd3, 0x26, 0x22, 0x32, 0x7a, 0x7b, 0xc9, 0xf7, 0xde, 0xeb, 0x7b, 0x5f, 0x1a,
	0xbd, 0xa0, 0x42, 0xd0, 0x1e, 0x52, 0xd2, 0xe9, 0x74, 0x86, 0x13, 0xda, 0x65, 0xa9, 0x02, 0x2d,
	0x46, 0xd5, 0x82, 0x4e, 0x29, 0x08, 0xaa, 0x88, 0xfc, 0xcc, 0xda, 0x66, 0xc7, 0x60, 0x8f, 0xa4,
	0x12, 0x46, 0xc4, 0x57, 0x33, 0x1b, 0x91, 0x4e, 0x23, 0x2f, 0x44, 0xbb, 0x0c, 0x79, 0xe1, 0xc5,
	0xf3, 0x63, 0xde, 0xc0, 0xc7, 0xc1, 0xfa, 0x36, 0x86, 0x28, 0x0a, 0x86, 0x71, 0xda, 0x98, 0x83,
	0x84, 0xd9, 0xfa, 0xe2, 0xd2, 0xe9, 0x24, 0x4b, 0x09, 0xe7, 0xc2, 0x10, 0xc3, 0x04, 0xd7, 0xcb,
	0xf4, 0xf1, 0x32, 0xb5, 0xa7, 0xed, 0xf8, 0x29, 0xdd, 0x2b, 0x22, 0x25, 0xa8, 0x65, 0xfe, 0xe4,
	0x67, 0x10, 0x9d, 0x17, 0x3e, 0x72, 0xc5, 0x60, 0x1f, 0x3f, 0x8d, 0xce, 0x5c, 0xaa, 0x86, 0x93,
	0x01, 0x92, 0x60, 0x15, 0xac, 0xef, 0x97, 0xa7, 0xee, 0xf2, 0x1d, 0x19, 0x20, 0xfe, 0x18, 0x3d,
	0x6a, 0xc5, 0xc8, 0x8d, 0x3a, 0x34, 0xbf, 0x93, 0x35, 0xad, 0xe0, 0xda, 0x10, 0x6e, 0x92, 0x70,
	0x15, 0xac, 0x1f, 0xdc, 0x5c, 0x2e, 0x5d, 0x91, 0xfb, 0x3a, 0x7a, 0x6f, 0x14, 0xe3, 0xb4, 0x22,
	0xfd, 0x08, 0x65, 0xb2, 0x18, 0x14, 0x20, 0x36, 0x56, 0xfe, 0x6a, 0x51, 0xc7, 0x2c, 0x3a, 0xeb,
	0x45, 0x6b, 0x7b, 0xd8, 0xa6, 0xc9, 0xc9, 0x2a, 0x58, 0x9f, 0xdf, 0xbc, 0x46, 0xc7, 0xb6, 0x68,
	0x57, 0x84, 0xbc, 0x11, 0xe3, 0x74, 0x73, 0x90, 0xf0, 0x86, 0x8f, 0xc3, 0x5f, 0x97, 0xe5, 0xa9,
	0xb3, 0x9e, 0x4e, 0xb7, 0x5f, 0xc2, 0xe8, 0xba, 0x15, 0x03, 0xfa, 0xef, 0xfb, 0xdc, 0x3e, 0xfc,
	0x73, 0x4d, 0x77, 0x53, 0xa5, 0xbb, 0xe0, 0xc3, 0xdb, 0x45, 0x49, 0x45, 0x4f, 0x38, 0x45, 0x42,
	0xd1, 0x94, 0x02, 0xb7, 0x85, 0xdd, 0x33, 0x4a, 0xa6, 0xff, 0xf1, 0xc7, 0xbc, 0xf4, 0xe8, 0x6b,
	0x78, 0x52, 0xe4, 0xf9, 0xb7, 0xf0, 0xaa, 0x98, 0x2d, 0xf3, 0x4e, 0xa3, 0x19, 0x4e, 0xa8, 0xca,
	0x50, 0xe9, 0x98, 0xdf, 0x1d, 0xa7, 0xce, 0x3b, 0x5d, 0x7b, 0x4e, 0x5d, 0x65, 0xb5, 0xe7, 0xfc,
	0x08, 0xaf, 0xe7, 0x01, 0xc6, 0x79, 0xa7, 0x31, 0xf6, 0x2c, 0x8c, 0xab, 0x0c, 0x63, 0xcf, 0xdb,
	0xde, 0xb3, 0x61, 0x9f, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x46, 0xa1, 0xd2, 0xdd, 0x02,
	0x00, 0x00,
}
