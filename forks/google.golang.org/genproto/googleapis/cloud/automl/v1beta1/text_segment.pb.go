// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/text_segment.proto

package automl

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

// A contiguous part of a text (string), assuming it has an UTF-8 NFC encoding.
// .
type TextSegment struct {
	// Output only. The content of the TextSegment.
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Required. Zero-based character index of the first character of the text
	// segment (counting characters from the beginning of the text).
	StartOffset int64 `protobuf:"varint,1,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	// Required. Zero-based character index of the first character past the end of
	// the text segment (counting character from the beginning of the text).
	// The character at the end_offset is NOT included in the text segment.
	EndOffset            int64    `protobuf:"varint,2,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSegment) Reset()         { *m = TextSegment{} }
func (m *TextSegment) String() string { return proto.CompactTextString(m) }
func (*TextSegment) ProtoMessage()    {}
func (*TextSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbd0421c131b44e6, []int{0}
}

func (m *TextSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSegment.Unmarshal(m, b)
}
func (m *TextSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSegment.Marshal(b, m, deterministic)
}
func (m *TextSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSegment.Merge(m, src)
}
func (m *TextSegment) XXX_Size() int {
	return xxx_messageInfo_TextSegment.Size(m)
}
func (m *TextSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSegment.DiscardUnknown(m)
}

var xxx_messageInfo_TextSegment proto.InternalMessageInfo

func (m *TextSegment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TextSegment) GetStartOffset() int64 {
	if m != nil {
		return m.StartOffset
	}
	return 0
}

func (m *TextSegment) GetEndOffset() int64 {
	if m != nil {
		return m.EndOffset
	}
	return 0
}

func init() {
	proto.RegisterType((*TextSegment)(nil), "google.cloud.automl.v1beta1.TextSegment")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/text_segment.proto", fileDescriptor_bbd0421c131b44e6)
}

var fileDescriptor_bbd0421c131b44e6 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0xd2, 0x82, 0xd2, 0xd4, 0x83, 0xec, 0x69, 0xb1, 0x8a, 0xd5, 0x53, 0x4f, 0x09,
	0xc5, 0xa3, 0xa7, 0xd6, 0x83, 0x27, 0xb1, 0x54, 0xf1, 0x20, 0x0b, 0x25, 0xdd, 0x9d, 0x86, 0x40,
	0x76, 0x66, 0xd9, 0xcc, 0x4a, 0x9f, 0xc4, 0x87, 0xf2, 0xa9, 0xa4, 0xc9, 0x16, 0x3c, 0x48, 0x8f,
	0x99, 0xf9, 0xe6, 0xe7, 0x27, 0x52, 0x59, 0x22, 0xeb, 0x41, 0x97, 0x9e, 0xba, 0x4a, 0x9b, 0x8e,
	0xa9, 0xf6, 0xfa, 0x6b, 0xbe, 0x05, 0x36, 0x73, 0xcd, 0xb0, 0xe7, 0x4d, 0x00, 0x5b, 0x03, 0xb2,
	0x6a, 0x5a, 0x62, 0xca, 0x26, 0xc9, 0xab, 0xe8, 0x55, 0xf2, 0xaa, 0xf7, 0x57, 0xd7, 0x7d, 0x98,
	0x69, 0x9c, 0x36, 0x88, 0xc4, 0x86, 0x1d, 0x61, 0x48, 0xa7, 0xf7, 0x4e, 0x8e, 0xdf, 0x61, 0xcf,
	0x6f, 0x29, 0x2f, 0xcb, 0xe5, 0x79, 0x49, 0xc8, 0x80, 0x9c, 0x0f, 0xa7, 0x62, 0x36, 0x5a, 0x1f,
	0x9f, 0xd9, 0x9d, 0xbc, 0x08, 0x6c, 0x5a, 0xde, 0xd0, 0x6e, 0x17, 0x80, 0x73, 0x31, 0x15, 0xb3,
	0xe1, 0x7a, 0x1c, 0x67, 0xaf, 0x71, 0x94, 0xdd, 0x48, 0x09, 0x58, 0x1d, 0xc1, 0x20, 0x82, 0x11,
	0x60, 0x95, 0xd6, 0xcb, 0x6f, 0x21, 0x6f, 0x4b, 0xaa, 0xd5, 0x89, 0xb2, 0xcb, 0xcb, 0x3f, 0x65,
	0x56, 0x87, 0x82, 0x2b, 0xf1, 0xb9, 0xe8, 0x0f, 0x2c, 0x79, 0x83, 0x56, 0x51, 0x6b, 0xb5, 0x05,
	0x8c, 0xf5, 0x75, 0x5a, 0x99, 0xc6, 0x85, 0x7f, 0x3f, 0xeb, 0x31, 0x3d, 0x7f, 0x06, 0x93, 0xe7,
	0x08, 0x8b, 0xa7, 0x03, 0x2a, 0x16, 0x1d, 0xd3, 0x8b, 0x2f, 0x3e, 0x12, 0xda, 0x9e, 0xc5, 0xac,
	0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x39, 0x59, 0x73, 0x77, 0x01, 0x00, 0x00,
}
