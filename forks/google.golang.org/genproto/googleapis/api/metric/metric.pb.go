// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/metric.proto

package metric

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	duration "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/duration"
	api "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api"
	label "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api/label"
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

// The kind of measurement. It describes how the data is reported.
type MetricDescriptor_MetricKind int32

const (
	// Do not use this default value.
	MetricDescriptor_METRIC_KIND_UNSPECIFIED MetricDescriptor_MetricKind = 0
	// An instantaneous measurement of a value.
	MetricDescriptor_GAUGE MetricDescriptor_MetricKind = 1
	// The change in a value during a time interval.
	MetricDescriptor_DELTA MetricDescriptor_MetricKind = 2
	// A value accumulated over a time interval.  Cumulative
	// measurements in a time series should have the same start time
	// and increasing end times, until an event resets the cumulative
	// value to zero and sets a new start time for the following
	// points.
	MetricDescriptor_CUMULATIVE MetricDescriptor_MetricKind = 3
)

var MetricDescriptor_MetricKind_name = map[int32]string{
	0: "METRIC_KIND_UNSPECIFIED",
	1: "GAUGE",
	2: "DELTA",
	3: "CUMULATIVE",
}

var MetricDescriptor_MetricKind_value = map[string]int32{
	"METRIC_KIND_UNSPECIFIED": 0,
	"GAUGE":                   1,
	"DELTA":                   2,
	"CUMULATIVE":              3,
}

func (x MetricDescriptor_MetricKind) String() string {
	return proto.EnumName(MetricDescriptor_MetricKind_name, int32(x))
}

func (MetricDescriptor_MetricKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_927eaac1a24f8abb, []int{0, 0}
}

// The value type of a metric.
type MetricDescriptor_ValueType int32

const (
	// Do not use this default value.
	MetricDescriptor_VALUE_TYPE_UNSPECIFIED MetricDescriptor_ValueType = 0
	// The value is a boolean.
	// This value type can be used only if the metric kind is `GAUGE`.
	MetricDescriptor_BOOL MetricDescriptor_ValueType = 1
	// The value is a signed 64-bit integer.
	MetricDescriptor_INT64 MetricDescriptor_ValueType = 2
	// The value is a double precision floating point number.
	MetricDescriptor_DOUBLE MetricDescriptor_ValueType = 3
	// The value is a text string.
	// This value type can be used only if the metric kind is `GAUGE`.
	MetricDescriptor_STRING MetricDescriptor_ValueType = 4
	// The value is a [`Distribution`][google.api.Distribution].
	MetricDescriptor_DISTRIBUTION MetricDescriptor_ValueType = 5
	// The value is money.
	MetricDescriptor_MONEY MetricDescriptor_ValueType = 6
)

var MetricDescriptor_ValueType_name = map[int32]string{
	0: "VALUE_TYPE_UNSPECIFIED",
	1: "BOOL",
	2: "INT64",
	3: "DOUBLE",
	4: "STRING",
	5: "DISTRIBUTION",
	6: "MONEY",
}

var MetricDescriptor_ValueType_value = map[string]int32{
	"VALUE_TYPE_UNSPECIFIED": 0,
	"BOOL":                   1,
	"INT64":                  2,
	"DOUBLE":                 3,
	"STRING":                 4,
	"DISTRIBUTION":           5,
	"MONEY":                  6,
}

func (x MetricDescriptor_ValueType) String() string {
	return proto.EnumName(MetricDescriptor_ValueType_name, int32(x))
}

func (MetricDescriptor_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_927eaac1a24f8abb, []int{0, 1}
}

// Defines a metric type and its schema. Once a metric descriptor is created,
// deleting or altering it stops data collection and makes the metric type's
// existing data unusable.
type MetricDescriptor struct {
	// The resource name of the metric descriptor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The metric type, including its DNS name prefix. The type is not
	// URL-encoded.  All user-defined metric types have the DNS name
	// `custom.googleapis.com` or `external.googleapis.com`.  Metric types should
	// use a natural hierarchical grouping. For example:
	//
	//     "custom.googleapis.com/invoice/paid/amount"
	//     "external.googleapis.com/prometheus/up"
	//     "appengine.googleapis.com/http/server/response_latencies"
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	// The set of labels that can be used to describe a specific
	// instance of this metric type. For example, the
	// `appengine.googleapis.com/http/server/response_latencies` metric
	// type has a label for the HTTP response code, `response_code`, so
	// you can look at latencies for successful responses or just
	// for responses that failed.
	Labels []*label.LabelDescriptor `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	// Whether the metric records instantaneous values, changes to a value, etc.
	// Some combinations of `metric_kind` and `value_type` might not be supported.
	MetricKind MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,proto3,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of `metric_kind` and `value_type` might not be supported.
	ValueType MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The unit in which the metric value is reported. It is only applicable
	// if the `value_type` is `INT64`, `DOUBLE`, or `DISTRIBUTION`. The
	// supported units are a subset of [The Unified Code for Units of
	// Measure](http://unitsofmeasure.org/ucum.html) standard:
	//
	// **Basic units (UNIT)**
	//
	// * `bit`   bit
	// * `By`    byte
	// * `s`     second
	// * `min`   minute
	// * `h`     hour
	// * `d`     day
	//
	// **Prefixes (PREFIX)**
	//
	// * `k`     kilo    (10**3)
	// * `M`     mega    (10**6)
	// * `G`     giga    (10**9)
	// * `T`     tera    (10**12)
	// * `P`     peta    (10**15)
	// * `E`     exa     (10**18)
	// * `Z`     zetta   (10**21)
	// * `Y`     yotta   (10**24)
	// * `m`     milli   (10**-3)
	// * `u`     micro   (10**-6)
	// * `n`     nano    (10**-9)
	// * `p`     pico    (10**-12)
	// * `f`     femto   (10**-15)
	// * `a`     atto    (10**-18)
	// * `z`     zepto   (10**-21)
	// * `y`     yocto   (10**-24)
	// * `Ki`    kibi    (2**10)
	// * `Mi`    mebi    (2**20)
	// * `Gi`    gibi    (2**30)
	// * `Ti`    tebi    (2**40)
	//
	// **Grammar**
	//
	// The grammar also includes these connectors:
	//
	// * `/`    division (as an infix operator, e.g. `1/s`).
	// * `.`    multiplication (as an infix operator, e.g. `GBy.d`)
	//
	// The grammar for a unit is as follows:
	//
	//     Expression = Component { "." Component } { "/" Component } ;
	//
	//     Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ]
	//               | Annotation
	//               | "1"
	//               ;
	//
	//     Annotation = "{" NAME "}" ;
	//
	// Notes:
	//
	// * `Annotation` is just a comment if it follows a `UNIT` and is
	//    equivalent to `1` if it is used alone. For examples,
	//    `{requests}/s == 1/s`, `By{transmitted}/s == By/s`.
	// * `NAME` is a sequence of non-blank printable ASCII characters not
	//    containing '{' or '}'.
	// * `1` represents dimensionless value 1, such as in `1/s`.
	// * `%` represents dimensionless value 1/100, and annotates values giving
	//    a percentage.
	Unit string `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	// A detailed description of the metric, which can be used in documentation.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// A concise name for the metric, which can be displayed in user interfaces.
	// Use sentence case without an ending period, for example "Request count".
	// This field is optional but it is recommended to be set for any metrics
	// associated with user-visible concepts, such as Quota.
	DisplayName string `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. Metadata which can be used to guide usage of the metric.
	Metadata *MetricDescriptor_MetricDescriptorMetadata `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Optional. The launch stage of the metric definition.
	LaunchStage          api.LaunchStage `protobuf:"varint,12,opt,name=launch_stage,json=launchStage,proto3,enum=google.api.LaunchStage" json:"launch_stage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MetricDescriptor) Reset()         { *m = MetricDescriptor{} }
func (m *MetricDescriptor) String() string { return proto.CompactTextString(m) }
func (*MetricDescriptor) ProtoMessage()    {}
func (*MetricDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_927eaac1a24f8abb, []int{0}
}

func (m *MetricDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricDescriptor.Unmarshal(m, b)
}
func (m *MetricDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricDescriptor.Marshal(b, m, deterministic)
}
func (m *MetricDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricDescriptor.Merge(m, src)
}
func (m *MetricDescriptor) XXX_Size() int {
	return xxx_messageInfo_MetricDescriptor.Size(m)
}
func (m *MetricDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_MetricDescriptor proto.InternalMessageInfo

func (m *MetricDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricDescriptor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetricDescriptor) GetLabels() []*label.LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetricDescriptor) GetMetricKind() MetricDescriptor_MetricKind {
	if m != nil {
		return m.MetricKind
	}
	return MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

func (m *MetricDescriptor) GetValueType() MetricDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (m *MetricDescriptor) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *MetricDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MetricDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *MetricDescriptor) GetMetadata() *MetricDescriptor_MetricDescriptorMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MetricDescriptor) GetLaunchStage() api.LaunchStage {
	if m != nil {
		return m.LaunchStage
	}
	return api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
}

// Additional annotations that can be used to guide the usage of a metric.
type MetricDescriptor_MetricDescriptorMetadata struct {
	// Deprecated. Please use the MetricDescriptor.launch_stage instead.
	// The launch stage of the metric definition.
	LaunchStage api.LaunchStage `protobuf:"varint,1,opt,name=launch_stage,json=launchStage,proto3,enum=google.api.LaunchStage" json:"launch_stage,omitempty"` // Deprecated: Do not use.
	// The sampling period of metric data points. For metrics which are written
	// periodically, consecutive data points are stored at this time interval,
	// excluding data loss due to errors. Metrics with a higher granularity have
	// a smaller sampling period.
	SamplePeriod *duration.Duration `protobuf:"bytes,2,opt,name=sample_period,json=samplePeriod,proto3" json:"sample_period,omitempty"`
	// The delay of data points caused by ingestion. Data points older than this
	// age are guaranteed to be ingested and available to be read, excluding
	// data loss due to errors.
	IngestDelay          *duration.Duration `protobuf:"bytes,3,opt,name=ingest_delay,json=ingestDelay,proto3" json:"ingest_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetricDescriptor_MetricDescriptorMetadata) Reset() {
	*m = MetricDescriptor_MetricDescriptorMetadata{}
}
func (m *MetricDescriptor_MetricDescriptorMetadata) String() string { return proto.CompactTextString(m) }
func (*MetricDescriptor_MetricDescriptorMetadata) ProtoMessage()    {}
func (*MetricDescriptor_MetricDescriptorMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_927eaac1a24f8abb, []int{0, 0}
}

func (m *MetricDescriptor_MetricDescriptorMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricDescriptor_MetricDescriptorMetadata.Unmarshal(m, b)
}
func (m *MetricDescriptor_MetricDescriptorMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricDescriptor_MetricDescriptorMetadata.Marshal(b, m, deterministic)
}
func (m *MetricDescriptor_MetricDescriptorMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricDescriptor_MetricDescriptorMetadata.Merge(m, src)
}
func (m *MetricDescriptor_MetricDescriptorMetadata) XXX_Size() int {
	return xxx_messageInfo_MetricDescriptor_MetricDescriptorMetadata.Size(m)
}
func (m *MetricDescriptor_MetricDescriptorMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricDescriptor_MetricDescriptorMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MetricDescriptor_MetricDescriptorMetadata proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *MetricDescriptor_MetricDescriptorMetadata) GetLaunchStage() api.LaunchStage {
	if m != nil {
		return m.LaunchStage
	}
	return api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
}

func (m *MetricDescriptor_MetricDescriptorMetadata) GetSamplePeriod() *duration.Duration {
	if m != nil {
		return m.SamplePeriod
	}
	return nil
}

func (m *MetricDescriptor_MetricDescriptorMetadata) GetIngestDelay() *duration.Duration {
	if m != nil {
		return m.IngestDelay
	}
	return nil
}

// A specific metric, identified by specifying values for all of the
// labels of a [`MetricDescriptor`][google.api.MetricDescriptor].
type Metric struct {
	// An existing metric type, see [google.api.MetricDescriptor][google.api.MetricDescriptor].
	// For example, `custom.googleapis.com/invoice/paid/amount`.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// The set of label values that uniquely identify this metric. All
	// labels listed in the `MetricDescriptor` must be assigned values.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_927eaac1a24f8abb, []int{1}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Metric) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.api.MetricDescriptor_MetricKind", MetricDescriptor_MetricKind_name, MetricDescriptor_MetricKind_value)
	proto.RegisterEnum("google.api.MetricDescriptor_ValueType", MetricDescriptor_ValueType_name, MetricDescriptor_ValueType_value)
	proto.RegisterType((*MetricDescriptor)(nil), "google.api.MetricDescriptor")
	proto.RegisterType((*MetricDescriptor_MetricDescriptorMetadata)(nil), "google.api.MetricDescriptor.MetricDescriptorMetadata")
	proto.RegisterType((*Metric)(nil), "google.api.Metric")
	proto.RegisterMapType((map[string]string)(nil), "google.api.Metric.LabelsEntry")
}

func init() { proto.RegisterFile("google/api/metric.proto", fileDescriptor_927eaac1a24f8abb) }

var fileDescriptor_927eaac1a24f8abb = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0xe9, 0xcf, 0xd6, 0x93, 0x32, 0x45, 0x16, 0xda, 0x42, 0x27, 0xa6, 0xd2, 0x0b, 0xe8,
	0x55, 0x2b, 0x6d, 0x30, 0x60, 0xa0, 0x49, 0xed, 0x12, 0x4a, 0xb4, 0x36, 0x2d, 0x59, 0x32, 0x69,
	0xdc, 0x44, 0x5e, 0x63, 0x42, 0xb4, 0x34, 0x09, 0x49, 0x3a, 0xa9, 0x4f, 0xc0, 0x25, 0xef, 0xc0,
	0x53, 0xf1, 0x38, 0xc8, 0x76, 0xda, 0x66, 0x45, 0x54, 0x5c, 0xe5, 0xf8, 0xfb, 0xbe, 0xf3, 0xd9,
	0xc7, 0x3e, 0x27, 0x70, 0xe0, 0x45, 0x91, 0x17, 0x90, 0x2e, 0x8e, 0xfd, 0xee, 0x8c, 0x64, 0x89,
	0x3f, 0xed, 0xc4, 0x49, 0x94, 0x45, 0x08, 0x38, 0xd1, 0xc1, 0xb1, 0xdf, 0xd8, 0x2f, 0x88, 0x02,
	0x7c, 0x4b, 0x02, 0xae, 0x69, 0x3c, 0x7b, 0x80, 0xcf, 0xc3, 0xe9, 0x37, 0x27, 0xcd, 0xb0, 0x47,
	0x72, 0xfa, 0x28, 0xa7, 0xd9, 0xea, 0x76, 0xfe, 0xb5, 0xeb, 0xce, 0x13, 0x9c, 0xf9, 0x51, 0xc8,
	0xf9, 0xd6, 0x8f, 0x1d, 0x90, 0x47, 0x6c, 0x4f, 0x95, 0xa4, 0xd3, 0xc4, 0x8f, 0xb3, 0x28, 0x41,
	0x08, 0xca, 0x21, 0x9e, 0x11, 0x45, 0x68, 0x0a, 0xed, 0x9a, 0xc9, 0x62, 0x8a, 0x65, 0x8b, 0x98,
	0x28, 0xbb, 0x1c, 0xa3, 0x31, 0x3a, 0x81, 0x2a, 0x3b, 0x4a, 0xaa, 0x88, 0xcd, 0x52, 0x5b, 0x3a,
	0x3e, 0xec, 0xac, 0x0f, 0xdc, 0x19, 0x52, 0x66, 0x6d, 0x6a, 0xe6, 0x52, 0xf4, 0x09, 0x24, 0x5e,
	0xa4, 0x73, 0xe7, 0x87, 0xae, 0x52, 0x6a, 0x0a, 0xed, 0xbd, 0xe3, 0x97, 0xc5, 0xcc, 0xcd, 0xf3,
	0xe4, 0xc0, 0xa5, 0x1f, 0xba, 0x26, 0xcc, 0x56, 0x31, 0xd2, 0x00, 0xee, 0x71, 0x30, 0x27, 0x0e,
	0x3b, 0x58, 0x99, 0x19, 0xbd, 0xd8, 0x6a, 0x74, 0x4d, 0xe5, 0xd6, 0x22, 0x26, 0x66, 0xed, 0x7e,
	0x19, 0xd2, 0xca, 0xe6, 0xa1, 0x9f, 0x29, 0x15, 0x5e, 0x19, 0x8d, 0x51, 0x13, 0x24, 0x37, 0x4f,
	0xf3, 0xa3, 0x50, 0xa9, 0x32, 0xaa, 0x08, 0xa1, 0xe7, 0x50, 0x77, 0xfd, 0x34, 0x0e, 0xf0, 0xc2,
	0x61, 0x77, 0xb5, 0x93, 0x4b, 0x38, 0x66, 0xd0, 0x2b, 0xfb, 0x0c, 0xbb, 0x33, 0x92, 0x61, 0x17,
	0x67, 0x58, 0x81, 0xa6, 0xd0, 0x96, 0x8e, 0x5f, 0xff, 0x47, 0x99, 0x6b, 0x60, 0x94, 0x27, 0x9b,
	0x2b, 0x1b, 0x74, 0x06, 0xf5, 0xe2, 0x23, 0x2b, 0x75, 0x56, 0xf4, 0xc1, 0xc3, 0x7b, 0xa7, 0xfc,
	0x15, 0xa5, 0x4d, 0x29, 0x58, 0x2f, 0x1a, 0xbf, 0x05, 0x50, 0xfe, 0xb5, 0x05, 0x3a, 0xdf, 0x30,
	0x16, 0xb6, 0x1a, 0xf7, 0x45, 0x45, 0x78, 0x60, 0x8e, 0xce, 0xe1, 0x71, 0x8a, 0x67, 0x71, 0x40,
	0x9c, 0x98, 0x24, 0x7e, 0xe4, 0x2a, 0x22, 0x2b, 0xf8, 0xe9, 0xd2, 0x60, 0xd9, 0x7f, 0x1d, 0x35,
	0xef, 0x3f, 0xb3, 0xce, 0xf5, 0x13, 0x26, 0x47, 0x1f, 0xa0, 0xee, 0x87, 0x1e, 0x49, 0x33, 0xc7,
	0x25, 0x01, 0x5e, 0xb0, 0xb6, 0xd8, 0x9a, 0x2e, 0x71, 0xb9, 0x4a, 0xd5, 0xad, 0x31, 0xc0, 0xba,
	0x47, 0xd0, 0x21, 0x1c, 0x8c, 0x34, 0xcb, 0xd4, 0x2f, 0x9c, 0x4b, 0xdd, 0x50, 0x1d, 0xdb, 0xb8,
	0x9a, 0x68, 0x17, 0xfa, 0x47, 0x5d, 0x53, 0xe5, 0x47, 0xa8, 0x06, 0x95, 0x41, 0xcf, 0x1e, 0x68,
	0xb2, 0x40, 0x43, 0x55, 0x1b, 0x5a, 0x3d, 0x59, 0x44, 0x7b, 0x00, 0x17, 0xf6, 0xc8, 0x1e, 0xf6,
	0x2c, 0xfd, 0x5a, 0x93, 0x4b, 0xad, 0xef, 0x50, 0x5b, 0xf5, 0x0a, 0x6a, 0xc0, 0xfe, 0x75, 0x6f,
	0x68, 0x6b, 0x8e, 0x75, 0x33, 0xd1, 0x36, 0xec, 0x76, 0xa1, 0xdc, 0x1f, 0x8f, 0x87, 0xdc, 0x4d,
	0x37, 0xac, 0xd3, 0x57, 0xb2, 0x88, 0x00, 0xaa, 0xea, 0xd8, 0xee, 0x0f, 0x35, 0xb9, 0x44, 0xe3,
	0x2b, 0xcb, 0xd4, 0x8d, 0x81, 0x5c, 0x46, 0x32, 0xd4, 0x55, 0x9d, 0xae, 0xfa, 0xb6, 0xa5, 0x8f,
	0x0d, 0xb9, 0x42, 0x93, 0x46, 0x63, 0x43, 0xbb, 0x91, 0xab, 0xad, 0x9f, 0x02, 0x54, 0x79, 0x11,
	0xab, 0x59, 0x2b, 0x15, 0x66, 0xed, 0x74, 0x63, 0xd6, 0x8e, 0xfe, 0x6e, 0x25, 0x3e, 0x72, 0xa9,
	0x16, 0x66, 0xc9, 0x62, 0x39, 0x6e, 0x8d, 0x77, 0x20, 0x15, 0x60, 0x24, 0x43, 0xe9, 0x8e, 0x2c,
	0xf2, 0xc9, 0xa6, 0x21, 0x7a, 0x02, 0x15, 0x36, 0x0b, 0xec, 0xc5, 0x6a, 0x26, 0x5f, 0x9c, 0x89,
	0x6f, 0x85, 0xbe, 0x03, 0x7b, 0xd3, 0x68, 0x56, 0xd8, 0xa7, 0x2f, 0xf1, 0x8d, 0x26, 0xf4, 0x35,
	0x26, 0xc2, 0x97, 0x37, 0x39, 0xe5, 0x45, 0x01, 0x0e, 0xbd, 0x4e, 0x94, 0x78, 0x5d, 0x8f, 0x84,
	0xec, 0xad, 0xba, 0x9c, 0xc2, 0xb1, 0x9f, 0x16, 0xfe, 0x6b, 0xef, 0xf9, 0xe7, 0x97, 0x58, 0x1e,
	0xf4, 0x26, 0xfa, 0x6d, 0x95, 0x49, 0x4f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x9a, 0x6a,
	0xfb, 0x01, 0x05, 0x00, 0x00,
}
