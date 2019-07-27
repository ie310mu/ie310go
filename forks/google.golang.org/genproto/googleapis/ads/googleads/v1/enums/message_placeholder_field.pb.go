// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/message_placeholder_field.proto

package enums

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

// Possible values for Message placeholder fields.
type MessagePlaceholderFieldEnum_MessagePlaceholderField int32

const (
	// Not specified.
	MessagePlaceholderFieldEnum_UNSPECIFIED MessagePlaceholderFieldEnum_MessagePlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	MessagePlaceholderFieldEnum_UNKNOWN MessagePlaceholderFieldEnum_MessagePlaceholderField = 1
	// Data Type: STRING. The name of your business.
	MessagePlaceholderFieldEnum_BUSINESS_NAME MessagePlaceholderFieldEnum_MessagePlaceholderField = 2
	// Data Type: STRING. Country code of phone number.
	MessagePlaceholderFieldEnum_COUNTRY_CODE MessagePlaceholderFieldEnum_MessagePlaceholderField = 3
	// Data Type: STRING. A phone number that's capable of sending and receiving
	// text messages.
	MessagePlaceholderFieldEnum_PHONE_NUMBER MessagePlaceholderFieldEnum_MessagePlaceholderField = 4
	// Data Type: STRING. The text that will go in your click-to-message ad.
	MessagePlaceholderFieldEnum_MESSAGE_EXTENSION_TEXT MessagePlaceholderFieldEnum_MessagePlaceholderField = 5
	// Data Type: STRING. The message text automatically shows in people's
	// messaging apps when they tap to send you a message.
	MessagePlaceholderFieldEnum_MESSAGE_TEXT MessagePlaceholderFieldEnum_MessagePlaceholderField = 6
)

var MessagePlaceholderFieldEnum_MessagePlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BUSINESS_NAME",
	3: "COUNTRY_CODE",
	4: "PHONE_NUMBER",
	5: "MESSAGE_EXTENSION_TEXT",
	6: "MESSAGE_TEXT",
}

var MessagePlaceholderFieldEnum_MessagePlaceholderField_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"BUSINESS_NAME":          2,
	"COUNTRY_CODE":           3,
	"PHONE_NUMBER":           4,
	"MESSAGE_EXTENSION_TEXT": 5,
	"MESSAGE_TEXT":           6,
}

func (x MessagePlaceholderFieldEnum_MessagePlaceholderField) String() string {
	return proto.EnumName(MessagePlaceholderFieldEnum_MessagePlaceholderField_name, int32(x))
}

func (MessagePlaceholderFieldEnum_MessagePlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_854ce0923e797876, []int{0, 0}
}

// Values for Message placeholder fields.
type MessagePlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagePlaceholderFieldEnum) Reset()         { *m = MessagePlaceholderFieldEnum{} }
func (m *MessagePlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*MessagePlaceholderFieldEnum) ProtoMessage()    {}
func (*MessagePlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_854ce0923e797876, []int{0}
}

func (m *MessagePlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagePlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *MessagePlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagePlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *MessagePlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagePlaceholderFieldEnum.Merge(m, src)
}
func (m *MessagePlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_MessagePlaceholderFieldEnum.Size(m)
}
func (m *MessagePlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagePlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MessagePlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.MessagePlaceholderFieldEnum_MessagePlaceholderField", MessagePlaceholderFieldEnum_MessagePlaceholderField_name, MessagePlaceholderFieldEnum_MessagePlaceholderField_value)
	proto.RegisterType((*MessagePlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.MessagePlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/message_placeholder_field.proto", fileDescriptor_854ce0923e797876)
}

var fileDescriptor_854ce0923e797876 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4f, 0x4e, 0xb3, 0x40,
	0x1c, 0xfd, 0xa0, 0x9f, 0x35, 0x99, 0x6a, 0x44, 0x16, 0x9a, 0xd4, 0x76, 0xd1, 0x1e, 0x60, 0x08,
	0x71, 0x37, 0xc6, 0x05, 0xb4, 0xd3, 0x4a, 0x0c, 0x03, 0x29, 0x50, 0xab, 0x21, 0x21, 0x58, 0x46,
	0x24, 0x01, 0x86, 0x74, 0xda, 0x9e, 0xc6, 0x95, 0x4b, 0xcf, 0xe0, 0x09, 0x3c, 0x8a, 0x3b, 0x6f,
	0x60, 0x00, 0xa9, 0xab, 0xba, 0x99, 0xbc, 0xfc, 0xde, 0x9f, 0x99, 0x79, 0x3f, 0x70, 0x1d, 0x33,
	0x16, 0xa7, 0x54, 0x09, 0x23, 0xae, 0xd4, 0xb0, 0x44, 0x5b, 0x55, 0xa1, 0xf9, 0x26, 0xe3, 0x4a,
	0x46, 0x39, 0x0f, 0x63, 0x1a, 0x14, 0x69, 0xb8, 0xa4, 0xcf, 0x2c, 0x8d, 0xe8, 0x2a, 0x78, 0x4a,
	0x68, 0x1a, 0xc1, 0x62, 0xc5, 0xd6, 0x4c, 0xee, 0xd7, 0x1e, 0x18, 0x46, 0x1c, 0xee, 0xec, 0x70,
	0xab, 0xc2, 0xca, 0xde, 0xed, 0x35, 0xe9, 0x45, 0xa2, 0x84, 0x79, 0xce, 0xd6, 0xe1, 0x3a, 0x61,
	0x39, 0xaf, 0xcd, 0xc3, 0x77, 0x01, 0x5c, 0x98, 0xf5, 0x05, 0xf6, 0x6f, 0xfe, 0xa4, 0x8c, 0xc7,
	0xf9, 0x26, 0x1b, 0xbe, 0x08, 0xe0, 0x7c, 0x0f, 0x2f, 0x9f, 0x80, 0x8e, 0x47, 0x1c, 0x1b, 0x8f,
	0x8c, 0x89, 0x81, 0xc7, 0xd2, 0x3f, 0xb9, 0x03, 0x0e, 0x3d, 0x72, 0x4b, 0xac, 0x3b, 0x22, 0x09,
	0xf2, 0x29, 0x38, 0xd6, 0x3d, 0xc7, 0x20, 0xd8, 0x71, 0x02, 0xa2, 0x99, 0x58, 0x12, 0x65, 0x09,
	0x1c, 0x8d, 0x2c, 0x8f, 0xb8, 0xb3, 0xfb, 0x60, 0x64, 0x8d, 0xb1, 0xd4, 0x2a, 0x27, 0xf6, 0x8d,
	0x45, 0x70, 0x40, 0x3c, 0x53, 0xc7, 0x33, 0xe9, 0xbf, 0xdc, 0x05, 0x67, 0x26, 0x76, 0x1c, 0x6d,
	0x8a, 0x03, 0xbc, 0x70, 0x31, 0x71, 0x0c, 0x8b, 0x04, 0x2e, 0x5e, 0xb8, 0xd2, 0x41, 0xa9, 0x6e,
	0xb8, 0x6a, 0xd2, 0xd6, 0xbf, 0x04, 0x30, 0x58, 0xb2, 0x0c, 0xfe, 0x59, 0x81, 0xde, 0xdb, 0xf3,
	0x03, 0xbb, 0xac, 0xc0, 0x16, 0x1e, 0xf4, 0x1f, 0x7b, 0xcc, 0xd2, 0x30, 0x8f, 0x21, 0x5b, 0xc5,
	0x4a, 0x4c, 0xf3, 0xaa, 0xa0, 0x66, 0x21, 0x45, 0xc2, 0xf7, 0xec, 0xe7, 0xaa, 0x3a, 0x5f, 0xc5,
	0xd6, 0x54, 0xd3, 0xde, 0xc4, 0xfe, 0xb4, 0x8e, 0xd2, 0x22, 0x0e, 0x6b, 0x58, 0xa2, 0xb9, 0x0a,
	0xcb, 0x36, 0xf9, 0x47, 0xc3, 0xfb, 0x5a, 0xc4, 0xfd, 0x1d, 0xef, 0xcf, 0x55, 0xbf, 0xe2, 0x3f,
	0xc5, 0x41, 0x3d, 0x44, 0x48, 0x8b, 0x38, 0x42, 0x3b, 0x05, 0x42, 0x73, 0x15, 0xa1, 0x4a, 0xf3,
	0xd8, 0xae, 0x1e, 0x76, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xea, 0xa9, 0x43, 0xed, 0x37, 0x02,
	0x00, 0x00,
}
