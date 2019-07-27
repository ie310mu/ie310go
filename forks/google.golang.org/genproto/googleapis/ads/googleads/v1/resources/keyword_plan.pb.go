// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/keyword_plan.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	wrappers "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// A Keyword Planner plan.
// Max number of saved keyword plans: 10000.
// It's possible to remove plans if limit is reached.
type KeywordPlan struct {
	// The resource name of the Keyword Planner plan.
	// KeywordPlan resource names have the form:
	//
	// `customers/{customer_id}/keywordPlans/{kp_plan_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the keyword plan.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the keyword plan.
	//
	// This field is required and should not be empty when creating new keyword
	// plans.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The date period used for forecasting the plan.
	ForecastPeriod       *KeywordPlanForecastPeriod `protobuf:"bytes,4,opt,name=forecast_period,json=forecastPeriod,proto3" json:"forecast_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *KeywordPlan) Reset()         { *m = KeywordPlan{} }
func (m *KeywordPlan) String() string { return proto.CompactTextString(m) }
func (*KeywordPlan) ProtoMessage()    {}
func (*KeywordPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_b229f482c8618309, []int{0}
}

func (m *KeywordPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlan.Unmarshal(m, b)
}
func (m *KeywordPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlan.Marshal(b, m, deterministic)
}
func (m *KeywordPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlan.Merge(m, src)
}
func (m *KeywordPlan) XXX_Size() int {
	return xxx_messageInfo_KeywordPlan.Size(m)
}
func (m *KeywordPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlan.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlan proto.InternalMessageInfo

func (m *KeywordPlan) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlan) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlan) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlan) GetForecastPeriod() *KeywordPlanForecastPeriod {
	if m != nil {
		return m.ForecastPeriod
	}
	return nil
}

// The forecasting period associated with the keyword plan.
type KeywordPlanForecastPeriod struct {
	// Required. The date used for forecasting the Plan.
	//
	// Types that are valid to be assigned to Interval:
	//	*KeywordPlanForecastPeriod_DateInterval
	//	*KeywordPlanForecastPeriod_DateRange
	Interval             isKeywordPlanForecastPeriod_Interval `protobuf_oneof:"interval"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *KeywordPlanForecastPeriod) Reset()         { *m = KeywordPlanForecastPeriod{} }
func (m *KeywordPlanForecastPeriod) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanForecastPeriod) ProtoMessage()    {}
func (*KeywordPlanForecastPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_b229f482c8618309, []int{1}
}

func (m *KeywordPlanForecastPeriod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Unmarshal(m, b)
}
func (m *KeywordPlanForecastPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Marshal(b, m, deterministic)
}
func (m *KeywordPlanForecastPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanForecastPeriod.Merge(m, src)
}
func (m *KeywordPlanForecastPeriod) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Size(m)
}
func (m *KeywordPlanForecastPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanForecastPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanForecastPeriod proto.InternalMessageInfo

type isKeywordPlanForecastPeriod_Interval interface {
	isKeywordPlanForecastPeriod_Interval()
}

type KeywordPlanForecastPeriod_DateInterval struct {
	DateInterval enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval `protobuf:"varint,1,opt,name=date_interval,json=dateInterval,proto3,enum=google.ads.googleads.v1.enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval,oneof"`
}

type KeywordPlanForecastPeriod_DateRange struct {
	DateRange *common.DateRange `protobuf:"bytes,2,opt,name=date_range,json=dateRange,proto3,oneof"`
}

func (*KeywordPlanForecastPeriod_DateInterval) isKeywordPlanForecastPeriod_Interval() {}

func (*KeywordPlanForecastPeriod_DateRange) isKeywordPlanForecastPeriod_Interval() {}

func (m *KeywordPlanForecastPeriod) GetInterval() isKeywordPlanForecastPeriod_Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *KeywordPlanForecastPeriod) GetDateInterval() enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval {
	if x, ok := m.GetInterval().(*KeywordPlanForecastPeriod_DateInterval); ok {
		return x.DateInterval
	}
	return enums.KeywordPlanForecastIntervalEnum_UNSPECIFIED
}

func (m *KeywordPlanForecastPeriod) GetDateRange() *common.DateRange {
	if x, ok := m.GetInterval().(*KeywordPlanForecastPeriod_DateRange); ok {
		return x.DateRange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeywordPlanForecastPeriod) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeywordPlanForecastPeriod_DateInterval)(nil),
		(*KeywordPlanForecastPeriod_DateRange)(nil),
	}
}

func init() {
	proto.RegisterType((*KeywordPlan)(nil), "google.ads.googleads.v1.resources.KeywordPlan")
	proto.RegisterType((*KeywordPlanForecastPeriod)(nil), "google.ads.googleads.v1.resources.KeywordPlanForecastPeriod")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/keyword_plan.proto", fileDescriptor_b229f482c8618309)
}

var fileDescriptor_b229f482c8618309 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0x9b, 0xb4, 0x88, 0x9d, 0x7e, 0x28, 0xb9, 0x5a, 0x6b, 0x91, 0xb6, 0x52, 0xa8, 0x0a,
	0x13, 0x53, 0x8b, 0x17, 0xd1, 0x9b, 0x2c, 0x6a, 0x3f, 0x04, 0x59, 0x22, 0xec, 0x45, 0x59, 0x58,
	0xa6, 0x3b, 0x67, 0x43, 0x30, 0x99, 0x09, 0x33, 0x93, 0x2d, 0x5e, 0xfa, 0x2a, 0x5e, 0xfa, 0x28,
	0x3e, 0x8a, 0x2f, 0xa0, 0x37, 0x82, 0x64, 0xbe, 0x68, 0xb1, 0x69, 0xef, 0xce, 0xc9, 0xf9, 0x9d,
	0xff, 0xf9, 0x9a, 0xa0, 0xa3, 0x82, 0xf3, 0xa2, 0x82, 0x98, 0x50, 0x19, 0x1b, 0xb3, 0xb3, 0x16,
	0x49, 0x2c, 0x40, 0xf2, 0x56, 0xcc, 0x40, 0xc6, 0x5f, 0xe0, 0xeb, 0x25, 0x17, 0x74, 0xda, 0x54,
	0x84, 0xe1, 0x46, 0x70, 0xc5, 0xa3, 0x5d, 0x83, 0x62, 0x42, 0x25, 0xf6, 0x59, 0x78, 0x91, 0x60,
	0x9f, 0xb5, 0xf5, 0xbc, 0x4f, 0x78, 0xc6, 0xeb, 0x9a, 0xb3, 0x98, 0x12, 0x05, 0xd2, 0xc8, 0x6d,
	0x0d, 0xfb, 0x58, 0x60, 0x6d, 0x7d, 0xbd, 0x81, 0xe9, 0x9c, 0x0b, 0x98, 0x11, 0xa9, 0xa6, 0x25,
	0x53, 0x20, 0x16, 0xa4, 0xb2, 0x1a, 0xdb, 0x4e, 0xa3, 0x29, 0x63, 0xc2, 0x18, 0x57, 0x44, 0x95,
	0x9c, 0xb9, 0x0a, 0x4f, 0x6c, 0x54, 0x7b, 0x17, 0xed, 0x3c, 0xbe, 0x14, 0xa4, 0x69, 0x40, 0xd8,
	0xf8, 0xde, 0x9f, 0x00, 0xad, 0x7d, 0x34, 0x65, 0x46, 0x15, 0x61, 0xd1, 0x53, 0xb4, 0xe1, 0x46,
	0x99, 0x32, 0x52, 0xc3, 0x20, 0xd8, 0x09, 0x0e, 0x56, 0xf3, 0x75, 0xf7, 0xf1, 0x13, 0xa9, 0x21,
	0x7a, 0x81, 0xc2, 0x92, 0x0e, 0xc2, 0x9d, 0xe0, 0x60, 0xed, 0xf0, 0xb1, 0xdd, 0x03, 0x76, 0x15,
	0xf0, 0x29, 0x53, 0xaf, 0x8f, 0xc6, 0xa4, 0x6a, 0x21, 0x0f, 0x4b, 0x1a, 0xbd, 0x44, 0x2b, 0x5a,
	0x68, 0x59, 0xe3, 0xdb, 0xff, 0xe1, 0x9f, 0x95, 0x28, 0x59, 0x61, 0x78, 0x4d, 0x46, 0x80, 0x1e,
	0xf8, 0x61, 0x1b, 0x10, 0x25, 0xa7, 0x83, 0x15, 0x9d, 0xfc, 0x16, 0xdf, 0xb9, 0x7e, 0x7c, 0x65,
	0x98, 0x0f, 0x56, 0x64, 0xa4, 0x35, 0xf2, 0xcd, 0xf9, 0x35, 0x7f, 0xef, 0x77, 0x80, 0x1e, 0xf5,
	0xd2, 0xd1, 0xb7, 0x00, 0x6d, 0x74, 0xa7, 0xf2, 0xeb, 0xd6, 0x9b, 0xd8, 0x3c, 0x3c, 0xef, 0xed,
	0x41, 0xdf, 0xec, 0xa6, 0xfa, 0xa7, 0x56, 0xe1, 0x3d, 0x6b, 0xeb, 0xdb, 0xe2, 0x27, 0x4b, 0xf9,
	0x7a, 0x57, 0xd2, 0xf9, 0xd1, 0x19, 0x42, 0xba, 0x05, 0x41, 0x58, 0x01, 0x76, 0xdf, 0xcf, 0x7a,
	0xeb, 0x9b, 0xf7, 0x85, 0xdf, 0x11, 0x05, 0x79, 0x97, 0x70, 0xb2, 0x94, 0xaf, 0x52, 0xe7, 0x0c,
	0x11, 0xba, 0xef, 0x26, 0x19, 0xfe, 0x0d, 0xd0, 0xfe, 0x8c, 0xd7, 0x77, 0x6f, 0x73, 0xf8, 0xf0,
	0x4a, 0xbb, 0xa3, 0xee, 0x62, 0xa3, 0xe0, 0xfc, 0xcc, 0xa6, 0x15, 0xbc, 0x22, 0xac, 0xc0, 0x5c,
	0x14, 0x71, 0x01, 0x4c, 0xdf, 0xd3, 0x3d, 0xe2, 0xa6, 0x94, 0xb7, 0xfc, 0x58, 0x6f, 0xbc, 0xf5,
	0x3d, 0x5c, 0x3e, 0xce, 0xb2, 0x1f, 0xe1, 0xee, 0xb1, 0x91, 0xcc, 0xa8, 0xc4, 0xc6, 0xec, 0xac,
	0x71, 0x82, 0x73, 0x47, 0xfe, 0x74, 0xcc, 0x24, 0xa3, 0x72, 0xe2, 0x99, 0xc9, 0x38, 0x99, 0x78,
	0xe6, 0x57, 0xb8, 0x6f, 0x02, 0x69, 0x9a, 0x51, 0x99, 0xa6, 0x9e, 0x4a, 0xd3, 0x71, 0x92, 0xa6,
	0x9e, 0xbb, 0xb8, 0xa7, 0x9b, 0x7d, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x93, 0x7a, 0x15, 0xc4,
	0x04, 0x04, 0x00, 0x00,
}
