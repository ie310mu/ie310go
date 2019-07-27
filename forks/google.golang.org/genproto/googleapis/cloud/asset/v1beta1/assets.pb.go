// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/asset/v1beta1/assets.proto

package asset

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	_ "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/any"
	_struct "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api/annotations"
	v1 "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/iam/v1"
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

// Temporal asset. In addition to the asset, the temporal asset includes the
// status of the asset and valid from and to time of it.
type TemporalAsset struct {
	// The time window when the asset data and state was observed.
	Window *TimeWindow `protobuf:"bytes,1,opt,name=window,proto3" json:"window,omitempty"`
	// If the asset is deleted or not.
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Asset.
	Asset                *Asset   `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemporalAsset) Reset()         { *m = TemporalAsset{} }
func (m *TemporalAsset) String() string { return proto.CompactTextString(m) }
func (*TemporalAsset) ProtoMessage()    {}
func (*TemporalAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0a79b6d8a101c1, []int{0}
}

func (m *TemporalAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemporalAsset.Unmarshal(m, b)
}
func (m *TemporalAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemporalAsset.Marshal(b, m, deterministic)
}
func (m *TemporalAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemporalAsset.Merge(m, src)
}
func (m *TemporalAsset) XXX_Size() int {
	return xxx_messageInfo_TemporalAsset.Size(m)
}
func (m *TemporalAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_TemporalAsset.DiscardUnknown(m)
}

var xxx_messageInfo_TemporalAsset proto.InternalMessageInfo

func (m *TemporalAsset) GetWindow() *TimeWindow {
	if m != nil {
		return m.Window
	}
	return nil
}

func (m *TemporalAsset) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *TemporalAsset) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

// A time window of (start_time, end_time].
type TimeWindow struct {
	// Start time of the time window (exclusive).
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time of the time window (inclusive).
	// Current timestamp if not specified.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeWindow) Reset()         { *m = TimeWindow{} }
func (m *TimeWindow) String() string { return proto.CompactTextString(m) }
func (*TimeWindow) ProtoMessage()    {}
func (*TimeWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0a79b6d8a101c1, []int{1}
}

func (m *TimeWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeWindow.Unmarshal(m, b)
}
func (m *TimeWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeWindow.Marshal(b, m, deterministic)
}
func (m *TimeWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeWindow.Merge(m, src)
}
func (m *TimeWindow) XXX_Size() int {
	return xxx_messageInfo_TimeWindow.Size(m)
}
func (m *TimeWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeWindow.DiscardUnknown(m)
}

var xxx_messageInfo_TimeWindow proto.InternalMessageInfo

func (m *TimeWindow) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeWindow) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// Cloud asset. This includes all Google Cloud Platform resources,
// Cloud IAM policies, and other non-GCP assets.
type Asset struct {
	// The full name of the asset. For example:
	// `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`.
	// See [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more information.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the asset. Example: "google.compute.Disk".
	AssetType string `protobuf:"bytes,2,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	// Representation of the resource.
	Resource *Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// Representation of the actual Cloud IAM policy set on a cloud resource. For
	// each resource, there must be at most one Cloud IAM policy set on it.
	IamPolicy            *v1.Policy `protobuf:"bytes,4,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0a79b6d8a101c1, []int{2}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *Asset) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Asset) GetIamPolicy() *v1.Policy {
	if m != nil {
		return m.IamPolicy
	}
	return nil
}

// Representation of a cloud resource.
type Resource struct {
	// The API version. Example: "v1".
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The URL of the discovery document containing the resource's JSON schema.
	// For example:
	// `"https://www.googleapis.com/discovery/v1/apis/compute/v1/rest"`.
	// It will be left unspecified for resources without a discovery-based API,
	// such as Cloud Bigtable.
	DiscoveryDocumentUri string `protobuf:"bytes,2,opt,name=discovery_document_uri,json=discoveryDocumentUri,proto3" json:"discovery_document_uri,omitempty"`
	// The JSON schema name listed in the discovery document.
	// Example: "Project". It will be left unspecified for resources (such as
	// Cloud Bigtable) without a discovery-based API.
	DiscoveryName string `protobuf:"bytes,3,opt,name=discovery_name,json=discoveryName,proto3" json:"discovery_name,omitempty"`
	// The REST URL for accessing the resource. An HTTP GET operation using this
	// URL returns the resource itself.
	// Example:
	// `https://cloudresourcemanager.googleapis.com/v1/projects/my-project-123`.
	// It will be left unspecified for resources without a REST API.
	ResourceUrl string `protobuf:"bytes,4,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// The full name of the immediate parent of this resource. See
	// [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more information.
	//
	// For GCP assets, it is the parent resource defined in the [Cloud IAM policy
	// hierarchy](https://cloud.google.com/iam/docs/overview#policy_hierarchy).
	// For example:
	// `"//cloudresourcemanager.googleapis.com/projects/my_project_123"`.
	//
	// For third-party assets, it is up to the users to define.
	Parent string `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	// The content of the resource, in which some sensitive fields are scrubbed
	// away and may not be present.
	Data                 *_struct.Struct `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0a79b6d8a101c1, []int{3}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Resource) GetDiscoveryDocumentUri() string {
	if m != nil {
		return m.DiscoveryDocumentUri
	}
	return ""
}

func (m *Resource) GetDiscoveryName() string {
	if m != nil {
		return m.DiscoveryName
	}
	return ""
}

func (m *Resource) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *Resource) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Resource) GetData() *_struct.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TemporalAsset)(nil), "google.cloud.asset.v1beta1.TemporalAsset")
	proto.RegisterType((*TimeWindow)(nil), "google.cloud.asset.v1beta1.TimeWindow")
	proto.RegisterType((*Asset)(nil), "google.cloud.asset.v1beta1.Asset")
	proto.RegisterType((*Resource)(nil), "google.cloud.asset.v1beta1.Resource")
}

func init() {
	proto.RegisterFile("google/cloud/asset/v1beta1/assets.proto", fileDescriptor_2e0a79b6d8a101c1)
}

var fileDescriptor_2e0a79b6d8a101c1 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdb, 0x6a, 0xd4, 0x40,
	0x18, 0x26, 0xed, 0xee, 0x76, 0xf3, 0xd7, 0x7a, 0x31, 0x68, 0x8d, 0x4b, 0xd5, 0x76, 0xf1, 0x50,
	0x10, 0x12, 0x56, 0x2b, 0x22, 0x82, 0x87, 0x55, 0xf0, 0x4e, 0x96, 0x71, 0x5b, 0x41, 0x16, 0xc2,
	0x6c, 0x32, 0x86, 0x81, 0x64, 0x26, 0x4c, 0x26, 0x29, 0x7b, 0xe3, 0xc3, 0x78, 0xa7, 0x97, 0x3e,
	0x86, 0x6f, 0xe2, 0x5b, 0x48, 0xfe, 0x99, 0x6c, 0xa1, 0xea, 0x7a, 0x97, 0x6f, 0xbe, 0xc3, 0xfc,
	0xdf, 0xcc, 0x04, 0x1e, 0x64, 0x4a, 0x65, 0x39, 0x8f, 0x92, 0x5c, 0xd5, 0x69, 0xc4, 0xaa, 0x8a,
	0x9b, 0xa8, 0x99, 0x2c, 0xb9, 0x61, 0x13, 0x8b, 0xaa, 0xb0, 0xd4, 0xca, 0x28, 0x32, 0xb2, 0xc2,
	0x10, 0x85, 0x21, 0x52, 0xa1, 0x13, 0x8e, 0x0e, 0x5c, 0x08, 0x2b, 0x45, 0xc4, 0xa4, 0x54, 0x86,
	0x19, 0xa1, 0xa4, 0x73, 0x8e, 0x9c, 0x33, 0x12, 0xac, 0x88, 0x9a, 0x49, 0x54, 0xaa, 0x5c, 0x24,
	0x2b, 0xc7, 0xdd, 0x74, 0x1c, 0xa2, 0x65, 0xfd, 0x39, 0x62, 0xb2, 0xa3, 0x0e, 0x2e, 0x53, 0x95,
	0xd1, 0x75, 0x62, 0x1c, 0x7b, 0xe7, 0x32, 0x6b, 0x44, 0xc1, 0x2b, 0xc3, 0x8a, 0xd2, 0x0a, 0xc6,
	0x5f, 0x3d, 0xd8, 0x9b, 0xf3, 0xa2, 0x54, 0x9a, 0xe5, 0xaf, 0xdb, 0x69, 0xc9, 0x0b, 0x18, 0x9c,
	0x0b, 0x99, 0xaa, 0xf3, 0xc0, 0x3b, 0xf4, 0x8e, 0x77, 0x1f, 0xdd, 0x0f, 0xff, 0x5d, 0x29, 0x9c,
	0x8b, 0x82, 0x7f, 0x44, 0x35, 0x75, 0x2e, 0x12, 0xc0, 0x4e, 0xca, 0x73, 0x6e, 0x78, 0x1a, 0x6c,
	0x1d, 0x7a, 0xc7, 0x43, 0xda, 0x41, 0xf2, 0x14, 0xfa, 0xe8, 0x0e, 0xb6, 0x31, 0xf8, 0x68, 0x53,
	0x30, 0xce, 0x42, 0xad, 0x7e, 0xfc, 0x05, 0xe0, 0x62, 0x23, 0xf2, 0x0c, 0xa0, 0x32, 0x4c, 0x9b,
	0xb8, 0xed, 0xe2, 0x86, 0x1c, 0x75, 0x59, 0x5d, 0x51, 0x9c, 0x0c, 0x8b, 0x52, 0x1f, 0xd5, 0x2d,
	0x26, 0x4f, 0x60, 0xc8, 0x65, 0x6a, 0x8d, 0x5b, 0xff, 0x35, 0xee, 0x70, 0x99, 0xb6, 0x68, 0xfc,
	0xc3, 0x83, 0xbe, 0x3d, 0x1c, 0x02, 0x3d, 0xc9, 0xdc, 0xae, 0x3e, 0xc5, 0x6f, 0x72, 0x0b, 0x00,
	0xc7, 0x8c, 0xcd, 0xaa, 0xb4, 0xb1, 0x3e, 0xf5, 0x71, 0x65, 0xbe, 0x2a, 0x39, 0x79, 0x05, 0x43,
	0xcd, 0x2b, 0x55, 0xeb, 0x84, 0xbb, 0xe2, 0x77, 0x37, 0x15, 0xa7, 0x4e, 0x4b, 0xd7, 0x2e, 0x72,
	0x02, 0x20, 0x58, 0x11, 0xdb, 0x17, 0x11, 0xf4, 0x30, 0xe3, 0x7a, 0x97, 0x21, 0x58, 0x11, 0x36,
	0x93, 0x70, 0x86, 0x24, 0xf5, 0x05, 0x2b, 0xec, 0xe7, 0xf8, 0x97, 0x07, 0xc3, 0x2e, 0xac, 0xbd,
	0x94, 0x86, 0xeb, 0x4a, 0x28, 0xe9, 0x46, 0xef, 0x20, 0x39, 0x81, 0xfd, 0x54, 0x54, 0x89, 0x6a,
	0xb8, 0x5e, 0xc5, 0xa9, 0x4a, 0xea, 0x82, 0x4b, 0x13, 0xd7, 0x5a, 0xb8, 0x26, 0xd7, 0xd6, 0xec,
	0x5b, 0x47, 0x9e, 0x6a, 0x41, 0xee, 0xc1, 0xd5, 0x0b, 0x17, 0x9e, 0xc8, 0x36, 0xaa, 0xf7, 0xd6,
	0xab, 0xef, 0xdb, 0xa3, 0x39, 0x82, 0x2b, 0x5d, 0x8b, 0xb8, 0xd6, 0x39, 0xce, 0xee, 0xd3, 0xdd,
	0x6e, 0xed, 0x54, 0xe7, 0x64, 0x1f, 0x06, 0x25, 0xd3, 0x5c, 0x9a, 0xa0, 0x8f, 0xa4, 0x43, 0xe4,
	0x21, 0xf4, 0x52, 0x66, 0x58, 0x30, 0xc0, 0xba, 0x37, 0xfe, 0xb8, 0xa6, 0x0f, 0xf8, 0xcc, 0x29,
	0x8a, 0xa6, 0xdf, 0x3c, 0xb8, 0x9d, 0xa8, 0x62, 0xc3, 0xb9, 0x4e, 0x01, 0x2f, 0x70, 0xd6, 0xda,
	0x67, 0xde, 0xa7, 0x97, 0x4e, 0x99, 0xa9, 0x9c, 0xc9, 0x2c, 0x54, 0x3a, 0x8b, 0x32, 0x2e, 0x31,
	0x3c, 0xb2, 0x14, 0x2b, 0x45, 0xf5, 0xb7, 0xdf, 0xfd, 0x39, 0xa2, 0xef, 0x5b, 0xa3, 0x77, 0x36,
	0xe1, 0x0d, 0xee, 0x85, 0xe1, 0xe1, 0xd9, 0x64, 0xda, 0x4a, 0x7e, 0x76, 0xe4, 0x02, 0xc9, 0x05,
	0x92, 0x8b, 0x33, 0xeb, 0x5f, 0x0e, 0x70, 0x97, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x07,
	0x7f, 0x2e, 0x7b, 0x53, 0x04, 0x00, 0x00,
}
