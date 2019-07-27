// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datalabeling/v1beta1/annotation_spec_set.proto

package datalabeling

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

// AnnotationSpecSet is a collection of label definitions. For example, in
// image classification tasks, we define a set of labels, this set is called
// AnnotationSpecSet. AnnotationSpecSet is immutable upon creation.
type AnnotationSpecSet struct {
	// Output only.
	// AnnotationSpecSet resource name, format:
	// projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name for AnnotationSpecSet defined by user.
	// Maximum of 64 characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-provided description of the annotation specification set.
	// The description can be up to 10000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Required. The actual spec set defined by the users.
	AnnotationSpecs      []*AnnotationSpec `protobuf:"bytes,4,rep,name=annotation_specs,json=annotationSpecs,proto3" json:"annotation_specs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AnnotationSpecSet) Reset()         { *m = AnnotationSpecSet{} }
func (m *AnnotationSpecSet) String() string { return proto.CompactTextString(m) }
func (*AnnotationSpecSet) ProtoMessage()    {}
func (*AnnotationSpecSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee18bab2b343845e, []int{0}
}

func (m *AnnotationSpecSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationSpecSet.Unmarshal(m, b)
}
func (m *AnnotationSpecSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationSpecSet.Marshal(b, m, deterministic)
}
func (m *AnnotationSpecSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationSpecSet.Merge(m, src)
}
func (m *AnnotationSpecSet) XXX_Size() int {
	return xxx_messageInfo_AnnotationSpecSet.Size(m)
}
func (m *AnnotationSpecSet) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationSpecSet.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationSpecSet proto.InternalMessageInfo

func (m *AnnotationSpecSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnnotationSpecSet) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AnnotationSpecSet) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AnnotationSpecSet) GetAnnotationSpecs() []*AnnotationSpec {
	if m != nil {
		return m.AnnotationSpecs
	}
	return nil
}

// Container of information related to one annotation spec.
type AnnotationSpec struct {
	// Required. The display name of the AnnotationSpec. Maximum of 64 characters.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-provided description of the annotation specification.
	// The description can be up to 10000 characters long.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotationSpec) Reset()         { *m = AnnotationSpec{} }
func (m *AnnotationSpec) String() string { return proto.CompactTextString(m) }
func (*AnnotationSpec) ProtoMessage()    {}
func (*AnnotationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee18bab2b343845e, []int{1}
}

func (m *AnnotationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationSpec.Unmarshal(m, b)
}
func (m *AnnotationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationSpec.Marshal(b, m, deterministic)
}
func (m *AnnotationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationSpec.Merge(m, src)
}
func (m *AnnotationSpec) XXX_Size() int {
	return xxx_messageInfo_AnnotationSpec.Size(m)
}
func (m *AnnotationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationSpec proto.InternalMessageInfo

func (m *AnnotationSpec) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AnnotationSpec) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*AnnotationSpecSet)(nil), "google.cloud.datalabeling.v1beta1.AnnotationSpecSet")
	proto.RegisterType((*AnnotationSpec)(nil), "google.cloud.datalabeling.v1beta1.AnnotationSpec")
}

func init() {
	proto.RegisterFile("google/cloud/datalabeling/v1beta1/annotation_spec_set.proto", fileDescriptor_ee18bab2b343845e)
}

var fileDescriptor_ee18bab2b343845e = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xe9, 0x36, 0x04, 0x33, 0xf1, 0x4f, 0x9f, 0x8a, 0xf8, 0xd0, 0x0d, 0x84, 0x3d, 0x25,
	0x54, 0x1f, 0xf7, 0xa4, 0xef, 0x8a, 0x6c, 0xf8, 0x22, 0x42, 0xb9, 0x6d, 0x2f, 0x21, 0x90, 0xe5,
	0x86, 0x26, 0x8a, 0x7e, 0x45, 0x3f, 0x95, 0x2c, 0x0d, 0xda, 0x3a, 0xb0, 0x6f, 0xc9, 0xb9, 0xe7,
	0x1c, 0x7e, 0x37, 0x61, 0x6b, 0x49, 0x24, 0x35, 0x8a, 0x5a, 0xd3, 0x5b, 0x23, 0x1a, 0xf0, 0xa0,
	0xa1, 0x42, 0xad, 0x8c, 0x14, 0xef, 0x45, 0x85, 0x1e, 0x0a, 0x01, 0xc6, 0x90, 0x07, 0xaf, 0xc8,
	0x94, 0xce, 0x62, 0x5d, 0x3a, 0xf4, 0xdc, 0xb6, 0xe4, 0x29, 0x5d, 0x74, 0x61, 0x1e, 0xc2, 0xbc,
	0x1f, 0xe6, 0x31, 0x7c, 0x79, 0x15, 0xfb, 0xc1, 0xaa, 0x5e, 0x91, 0xeb, 0x0a, 0x96, 0x5f, 0x09,
	0xbb, 0xb8, 0xfb, 0x51, 0xb7, 0x16, 0xeb, 0x2d, 0xfa, 0x34, 0x65, 0x33, 0x03, 0x3b, 0xcc, 0x92,
	0x3c, 0x59, 0x1d, 0x6f, 0xc2, 0x39, 0x5d, 0xb0, 0x93, 0x46, 0x39, 0xab, 0xe1, 0xb3, 0x0c, 0xb3,
	0x49, 0x98, 0xcd, 0xa3, 0xf6, 0xb8, 0xb7, 0xe4, 0x6c, 0xde, 0xa0, 0xab, 0x5b, 0x65, 0xf7, 0x65,
	0xd9, 0x34, 0x3a, 0x7e, 0xa5, 0xf4, 0x95, 0x9d, 0xff, 0x59, 0xc6, 0x65, 0xb3, 0x7c, 0xba, 0x9a,
	0xdf, 0x14, 0x7c, 0x74, 0x15, 0x3e, 0x04, 0xdd, 0x9c, 0xc1, 0xe0, 0xee, 0x96, 0xcf, 0xec, 0x74,
	0x68, 0x39, 0x80, 0x4e, 0x46, 0xa1, 0x27, 0x07, 0xd0, 0xf7, 0x1f, 0xec, 0xba, 0xa6, 0xdd, 0x38,
	0xdf, 0x53, 0xf2, 0xf2, 0x10, 0x4d, 0x92, 0x34, 0x18, 0xc9, 0xa9, 0x95, 0x42, 0xa2, 0x09, 0x4f,
	0x2d, 0xba, 0x11, 0x58, 0xe5, 0xfe, 0xf9, 0xeb, 0x75, 0x5f, 0xac, 0x8e, 0x42, 0xf2, 0xf6, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0xca, 0xe4, 0xde, 0x24, 0x02, 0x00, 0x00,
}
