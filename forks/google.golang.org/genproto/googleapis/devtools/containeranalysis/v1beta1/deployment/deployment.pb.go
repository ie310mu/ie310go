// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1beta1/deployment/deployment.proto

package deployment

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	timestamp "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/timestamp"
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

// Types of platforms.
type Deployment_Platform int32

const (
	// Unknown.
	Deployment_PLATFORM_UNSPECIFIED Deployment_Platform = 0
	// Google Container Engine.
	Deployment_GKE Deployment_Platform = 1
	// Google App Engine: Flexible Environment.
	Deployment_FLEX Deployment_Platform = 2
	// Custom user-defined platform.
	Deployment_CUSTOM Deployment_Platform = 3
)

var Deployment_Platform_name = map[int32]string{
	0: "PLATFORM_UNSPECIFIED",
	1: "GKE",
	2: "FLEX",
	3: "CUSTOM",
}

var Deployment_Platform_value = map[string]int32{
	"PLATFORM_UNSPECIFIED": 0,
	"GKE":                  1,
	"FLEX":                 2,
	"CUSTOM":               3,
}

func (x Deployment_Platform) String() string {
	return proto.EnumName(Deployment_Platform_name, int32(x))
}

func (Deployment_Platform) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_27c2fb509818da47, []int{2, 0}
}

// An artifact that can be deployed in some runtime.
type Deployable struct {
	// Resource URI for the artifact being deployed.
	ResourceUri          []string `protobuf:"bytes,1,rep,name=resource_uri,json=resourceUri,proto3" json:"resource_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deployable) Reset()         { *m = Deployable{} }
func (m *Deployable) String() string { return proto.CompactTextString(m) }
func (*Deployable) ProtoMessage()    {}
func (*Deployable) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c2fb509818da47, []int{0}
}

func (m *Deployable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployable.Unmarshal(m, b)
}
func (m *Deployable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployable.Marshal(b, m, deterministic)
}
func (m *Deployable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployable.Merge(m, src)
}
func (m *Deployable) XXX_Size() int {
	return xxx_messageInfo_Deployable.Size(m)
}
func (m *Deployable) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployable.DiscardUnknown(m)
}

var xxx_messageInfo_Deployable proto.InternalMessageInfo

func (m *Deployable) GetResourceUri() []string {
	if m != nil {
		return m.ResourceUri
	}
	return nil
}

// Details of a deployment occurrence.
type Details struct {
	// Deployment history for the resource.
	Deployment           *Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c2fb509818da47, []int{1}
}

func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (m *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(m, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

// The period during which some deployable was active in a runtime.
type Deployment struct {
	// Identity of the user that triggered this deployment.
	UserEmail string `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// Beginning of the lifetime of this deployment.
	DeployTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=deploy_time,json=deployTime,proto3" json:"deploy_time,omitempty"`
	// End of the lifetime of this deployment.
	UndeployTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=undeploy_time,json=undeployTime,proto3" json:"undeploy_time,omitempty"`
	// Configuration used to create this deployment.
	Config string `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// Address of the runtime element hosting this deployment.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Output only. Resource URI for the artifact being deployed taken from
	// the deployable field with the same name.
	ResourceUri []string `protobuf:"bytes,6,rep,name=resource_uri,json=resourceUri,proto3" json:"resource_uri,omitempty"`
	// Platform hosting this deployment.
	Platform             Deployment_Platform `protobuf:"varint,7,opt,name=platform,proto3,enum=grafeas.v1beta1.deployment.Deployment_Platform" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c2fb509818da47, []int{2}
}

func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *Deployment) GetDeployTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeployTime
	}
	return nil
}

func (m *Deployment) GetUndeployTime() *timestamp.Timestamp {
	if m != nil {
		return m.UndeployTime
	}
	return nil
}

func (m *Deployment) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *Deployment) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Deployment) GetResourceUri() []string {
	if m != nil {
		return m.ResourceUri
	}
	return nil
}

func (m *Deployment) GetPlatform() Deployment_Platform {
	if m != nil {
		return m.Platform
	}
	return Deployment_PLATFORM_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("grafeas.v1beta1.deployment.Deployment_Platform", Deployment_Platform_name, Deployment_Platform_value)
	proto.RegisterType((*Deployable)(nil), "grafeas.v1beta1.deployment.Deployable")
	proto.RegisterType((*Details)(nil), "grafeas.v1beta1.deployment.Details")
	proto.RegisterType((*Deployment)(nil), "grafeas.v1beta1.deployment.Deployment")
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1beta1/deployment/deployment.proto", fileDescriptor_27c2fb509818da47)
}

var fileDescriptor_27c2fb509818da47 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xad, 0xac, 0xd4, 0x1f, 0xe3, 0xb4, 0x98, 0xa5, 0x14, 0x61, 0x08, 0x75, 0x7d, 0x28, 0x3e,
	0xed, 0x92, 0xf4, 0x98, 0x43, 0x49, 0x6d, 0x39, 0x84, 0x24, 0x8d, 0xab, 0xd8, 0x50, 0x7a, 0xa8,
	0x59, 0xdb, 0x63, 0xb1, 0xb0, 0xd2, 0x8a, 0xdd, 0x55, 0x20, 0xbf, 0xa0, 0xd7, 0xfe, 0x86, 0xfe,
	0xd2, 0xa2, 0x95, 0x14, 0x1b, 0x4c, 0xbf, 0x6e, 0x33, 0x6f, 0xe7, 0xcd, 0xbc, 0xf7, 0x58, 0x08,
	0x63, 0xa5, 0x62, 0x89, 0x6c, 0x83, 0x0f, 0x56, 0x29, 0x69, 0xd8, 0x5a, 0xa5, 0x96, 0x8b, 0x14,
	0x35, 0x4f, 0xb9, 0x7c, 0x34, 0xc2, 0xb0, 0x87, 0xd3, 0x15, 0x5a, 0x7e, 0xca, 0x36, 0x98, 0x49,
	0xf5, 0x98, 0x60, 0x6a, 0xf7, 0x4a, 0x9a, 0x69, 0x65, 0x15, 0xe9, 0xc7, 0x9a, 0x6f, 0x91, 0x1b,
	0x5a, 0x0d, 0xd3, 0xdd, 0x44, 0xff, 0x4d, 0x75, 0xc2, 0x4d, 0xae, 0xf2, 0x2d, 0xb3, 0x22, 0x41,
	0x63, 0x79, 0x92, 0x95, 0xe4, 0x21, 0x03, 0x98, 0xb8, 0x71, 0xbe, 0x92, 0x48, 0xde, 0xc2, 0xb1,
	0x46, 0xa3, 0x72, 0xbd, 0xc6, 0x65, 0xae, 0x45, 0xe0, 0x0d, 0xfc, 0x51, 0x27, 0xea, 0xd6, 0xd8,
	0x42, 0x8b, 0xe1, 0x67, 0x68, 0x4d, 0xd0, 0x72, 0x21, 0x0d, 0x99, 0x02, 0xec, 0x4e, 0x05, 0xde,
	0xc0, 0x1b, 0x75, 0xcf, 0xde, 0xd1, 0xdf, 0xab, 0xa1, 0x93, 0xa7, 0x32, 0xda, 0x63, 0x0e, 0x7f,
	0xf8, 0xb5, 0x88, 0xa2, 0x25, 0x27, 0x00, 0xb9, 0x41, 0xbd, 0xc4, 0x84, 0x0b, 0xe9, 0xd6, 0x76,
	0xa2, 0x4e, 0x81, 0x84, 0x05, 0x40, 0xce, 0xa1, 0x5b, 0x72, 0x97, 0x85, 0x97, 0xa0, 0xe1, 0xce,
	0xf6, 0x69, 0x69, 0x94, 0xd6, 0x46, 0xe9, 0xbc, 0x36, 0x5a, 0x9f, 0x2a, 0x00, 0xf2, 0x01, 0x5e,
	0xe4, 0xe9, 0x3e, 0xdd, 0xff, 0x2b, 0xfd, 0xb8, 0x26, 0xb8, 0x05, 0xaf, 0xa1, 0xb9, 0x56, 0xe9,
	0x56, 0xc4, 0xc1, 0x91, 0x13, 0x56, 0x75, 0x24, 0x80, 0x16, 0xdf, 0x6c, 0x34, 0x1a, 0x13, 0x3c,
	0x77, 0x0f, 0x75, 0x7b, 0x90, 0x69, 0xf3, 0x20, 0x53, 0x72, 0x0d, 0xed, 0x4c, 0x72, 0xbb, 0x55,
	0x3a, 0x09, 0x5a, 0x03, 0x6f, 0xf4, 0xf2, 0x8c, 0xfd, 0x5b, 0x8c, 0x74, 0x56, 0xd1, 0xa2, 0xa7,
	0x05, 0xc3, 0x31, 0xb4, 0x6b, 0x94, 0x04, 0xf0, 0x6a, 0x76, 0x73, 0x31, 0x9f, 0xde, 0x45, 0xb7,
	0xcb, 0xc5, 0xa7, 0xfb, 0x59, 0x38, 0xbe, 0x9a, 0x5e, 0x85, 0x93, 0xde, 0x33, 0xd2, 0x02, 0xff,
	0xf2, 0x3a, 0xec, 0x79, 0xa4, 0x0d, 0x47, 0xd3, 0x9b, 0xf0, 0x4b, 0xaf, 0x41, 0x00, 0x9a, 0xe3,
	0xc5, 0xfd, 0xfc, 0xee, 0xb6, 0xe7, 0x7f, 0xfc, 0xee, 0xc1, 0x89, 0x50, 0x7f, 0x10, 0x31, 0xf3,
	0xbe, 0x7e, 0xab, 0x32, 0x8b, 0x95, 0xe4, 0x69, 0x4c, 0x95, 0x8e, 0x59, 0x8c, 0xa9, 0x4b, 0x90,
	0x95, 0x4f, 0x3c, 0x13, 0xe6, 0xff, 0x7e, 0xf7, 0xf9, 0xae, 0xfc, 0xd9, 0xf0, 0x2f, 0xa3, 0x8b,
	0x55, 0xd3, 0x2d, 0x7c, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x59, 0xd5, 0x23, 0x2d, 0x03,
	0x00, 0x00,
}
