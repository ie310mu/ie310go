// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/promotion_placeholder_field.proto

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

// Possible values for Promotion placeholder fields.
type PromotionPlaceholderFieldEnum_PromotionPlaceholderField int32

const (
	// Not specified.
	PromotionPlaceholderFieldEnum_UNSPECIFIED PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	PromotionPlaceholderFieldEnum_UNKNOWN PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 1
	// Data Type: STRING. The text that appears on the ad when the extension is
	// shown.
	PromotionPlaceholderFieldEnum_PROMOTION_TARGET PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 2
	// Data Type: STRING. Allows you to add "up to" phrase to the promotion,
	// in case you have variable promotion rates.
	PromotionPlaceholderFieldEnum_DISCOUNT_MODIFIER PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 3
	// Data Type: INT64. Takes a value in micros, where 1 million micros
	// represents 1%, and is shown as a percentage when rendered.
	PromotionPlaceholderFieldEnum_PERCENT_OFF PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 4
	// Data Type: MONEY. Requires a currency and an amount of money.
	PromotionPlaceholderFieldEnum_MONEY_AMOUNT_OFF PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 5
	// Data Type: STRING. A string that the user enters to get the discount.
	PromotionPlaceholderFieldEnum_PROMOTION_CODE PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 6
	// Data Type: MONEY. A minimum spend before the user qualifies for the
	// promotion.
	PromotionPlaceholderFieldEnum_ORDERS_OVER_AMOUNT PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 7
	// Data Type: DATE. The start date of the promotion.
	PromotionPlaceholderFieldEnum_PROMOTION_START PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 8
	// Data Type: DATE. The end date of the promotion.
	PromotionPlaceholderFieldEnum_PROMOTION_END PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 9
	// Data Type: STRING. Describes the associated event for the promotion using
	// one of the PromotionExtensionOccasion enum values, for example NEW_YEARS.
	PromotionPlaceholderFieldEnum_OCCASION PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 10
	// Data Type: URL_LIST. Final URLs to be used in the ad when using Upgraded
	// URLs.
	PromotionPlaceholderFieldEnum_FINAL_URLS PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 11
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	PromotionPlaceholderFieldEnum_FINAL_MOBILE_URLS PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 12
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	PromotionPlaceholderFieldEnum_TRACKING_URL PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 13
	// Data Type: STRING. A string represented by a language code for the
	// promotion.
	PromotionPlaceholderFieldEnum_LANGUAGE PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 14
	// Data Type: STRING. Final URL suffix for the ad when using parallel
	// tracking.
	PromotionPlaceholderFieldEnum_FINAL_URL_SUFFIX PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 15
)

var PromotionPlaceholderFieldEnum_PromotionPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "PROMOTION_TARGET",
	3:  "DISCOUNT_MODIFIER",
	4:  "PERCENT_OFF",
	5:  "MONEY_AMOUNT_OFF",
	6:  "PROMOTION_CODE",
	7:  "ORDERS_OVER_AMOUNT",
	8:  "PROMOTION_START",
	9:  "PROMOTION_END",
	10: "OCCASION",
	11: "FINAL_URLS",
	12: "FINAL_MOBILE_URLS",
	13: "TRACKING_URL",
	14: "LANGUAGE",
	15: "FINAL_URL_SUFFIX",
}

var PromotionPlaceholderFieldEnum_PromotionPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"PROMOTION_TARGET":   2,
	"DISCOUNT_MODIFIER":  3,
	"PERCENT_OFF":        4,
	"MONEY_AMOUNT_OFF":   5,
	"PROMOTION_CODE":     6,
	"ORDERS_OVER_AMOUNT": 7,
	"PROMOTION_START":    8,
	"PROMOTION_END":      9,
	"OCCASION":           10,
	"FINAL_URLS":         11,
	"FINAL_MOBILE_URLS":  12,
	"TRACKING_URL":       13,
	"LANGUAGE":           14,
	"FINAL_URL_SUFFIX":   15,
}

func (x PromotionPlaceholderFieldEnum_PromotionPlaceholderField) String() string {
	return proto.EnumName(PromotionPlaceholderFieldEnum_PromotionPlaceholderField_name, int32(x))
}

func (PromotionPlaceholderFieldEnum_PromotionPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_82cfe61f98d75179, []int{0, 0}
}

// Values for Promotion placeholder fields.
type PromotionPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromotionPlaceholderFieldEnum) Reset()         { *m = PromotionPlaceholderFieldEnum{} }
func (m *PromotionPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*PromotionPlaceholderFieldEnum) ProtoMessage()    {}
func (*PromotionPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_82cfe61f98d75179, []int{0}
}

func (m *PromotionPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromotionPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *PromotionPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromotionPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *PromotionPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromotionPlaceholderFieldEnum.Merge(m, src)
}
func (m *PromotionPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_PromotionPlaceholderFieldEnum.Size(m)
}
func (m *PromotionPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PromotionPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PromotionPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.PromotionPlaceholderFieldEnum_PromotionPlaceholderField", PromotionPlaceholderFieldEnum_PromotionPlaceholderField_name, PromotionPlaceholderFieldEnum_PromotionPlaceholderField_value)
	proto.RegisterType((*PromotionPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.PromotionPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/promotion_placeholder_field.proto", fileDescriptor_82cfe61f98d75179)
}

var fileDescriptor_82cfe61f98d75179 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0x9b, 0x4e,
	0x10, 0xc7, 0x7f, 0x26, 0xbf, 0x26, 0xe9, 0xda, 0xb1, 0x37, 0xdb, 0x3f, 0x52, 0xab, 0xba, 0x52,
	0xf2, 0x00, 0x20, 0xd4, 0x1b, 0x3d, 0x54, 0x6b, 0x58, 0x10, 0x8a, 0xd9, 0x45, 0x0b, 0xb8, 0x7f,
	0x64, 0x69, 0x45, 0x03, 0xa5, 0x96, 0x30, 0x8b, 0xbc, 0x4e, 0x1e, 0xa8, 0xc7, 0x3e, 0x4a, 0xd5,
	0x27, 0xe9, 0xa1, 0x97, 0xbe, 0x40, 0xb5, 0x10, 0xdb, 0x27, 0xf7, 0x82, 0x86, 0xef, 0x7c, 0xe7,
	0x33, 0xa3, 0x9d, 0x01, 0xef, 0x2a, 0x29, 0xab, 0xba, 0xb4, 0xf2, 0x42, 0x59, 0x7d, 0xa8, 0xa3,
	0x7b, 0xdb, 0x2a, 0x9b, 0xbb, 0xb5, 0xb2, 0xda, 0x8d, 0x5c, 0xcb, 0xed, 0x4a, 0x36, 0xa2, 0xad,
	0xf3, 0xdb, 0xf2, 0xab, 0xac, 0x8b, 0x72, 0x23, 0xbe, 0xac, 0xca, 0xba, 0x30, 0xdb, 0x8d, 0xdc,
	0x4a, 0x34, 0xed, 0xab, 0xcc, 0xbc, 0x50, 0xe6, 0x1e, 0x60, 0xde, 0xdb, 0x66, 0x07, 0x78, 0xf9,
	0x6a, 0xc7, 0x6f, 0x57, 0x56, 0xde, 0x34, 0x72, 0x9b, 0x6b, 0x9a, 0xea, 0x8b, 0xaf, 0x7f, 0x1b,
	0x60, 0x1a, 0xef, 0x5a, 0xc4, 0x87, 0x0e, 0xbe, 0x6e, 0x40, 0x9a, 0xbb, 0xf5, 0xf5, 0x4f, 0x03,
	0xbc, 0x38, 0xea, 0x40, 0x13, 0x30, 0xcc, 0x68, 0x12, 0x13, 0x37, 0xf4, 0x43, 0xe2, 0xc1, 0xff,
	0xd0, 0x10, 0x9c, 0x65, 0xf4, 0x86, 0xb2, 0xf7, 0x14, 0x0e, 0xd0, 0x53, 0x00, 0x63, 0xce, 0x22,
	0x96, 0x86, 0x8c, 0x8a, 0x14, 0xf3, 0x80, 0xa4, 0xd0, 0x40, 0xcf, 0xc0, 0xa5, 0x17, 0x26, 0x2e,
	0xcb, 0x68, 0x2a, 0x22, 0xe6, 0xe9, 0x4a, 0x0e, 0x4f, 0x34, 0x2a, 0x26, 0xdc, 0x25, 0x34, 0x15,
	0xcc, 0xf7, 0xe1, 0xff, 0xba, 0x3a, 0x62, 0x94, 0x7c, 0x14, 0x38, 0xea, 0xbc, 0x5a, 0x7d, 0x84,
	0x10, 0x18, 0x1f, 0x98, 0x2e, 0xf3, 0x08, 0x3c, 0x45, 0xcf, 0x01, 0x62, 0xdc, 0x23, 0x3c, 0x11,
	0x6c, 0x41, 0xf8, 0x83, 0x1f, 0x9e, 0xa1, 0x27, 0x60, 0x72, 0xf0, 0x26, 0x29, 0xe6, 0x29, 0x3c,
	0x47, 0x97, 0xe0, 0xe2, 0x20, 0x12, 0xea, 0xc1, 0xc7, 0x68, 0x04, 0xce, 0x99, 0xeb, 0xe2, 0x24,
	0x64, 0x14, 0x02, 0x34, 0x06, 0xc0, 0x0f, 0x29, 0x9e, 0x8b, 0x8c, 0xcf, 0x13, 0x38, 0xd4, 0xf3,
	0xf6, 0xff, 0x11, 0x9b, 0x85, 0x73, 0xd2, 0xcb, 0x23, 0x04, 0xc1, 0x28, 0xe5, 0xd8, 0xbd, 0x09,
	0x69, 0xa0, 0x25, 0x78, 0xa1, 0x31, 0x73, 0x4c, 0x83, 0x0c, 0x07, 0x04, 0x8e, 0xf5, 0xf8, 0x7b,
	0x8c, 0x48, 0x32, 0xdf, 0x0f, 0x3f, 0xc0, 0xc9, 0xec, 0xcf, 0x00, 0x5c, 0xdd, 0xca, 0xb5, 0xf9,
	0xcf, 0xa5, 0xcd, 0x5e, 0x1f, 0x7d, 0xf1, 0x58, 0xaf, 0x2d, 0x1e, 0x7c, 0x9a, 0x3d, 0x00, 0x2a,
	0x59, 0xe7, 0x4d, 0x65, 0xca, 0x4d, 0x65, 0x55, 0x65, 0xd3, 0x2d, 0x75, 0x77, 0x46, 0xed, 0x4a,
	0x1d, 0xb9, 0xaa, 0xb7, 0xdd, 0xf7, 0x9b, 0x71, 0x12, 0x60, 0xfc, 0xdd, 0x98, 0x06, 0x3d, 0x0a,
	0x17, 0xca, 0xec, 0x43, 0x1d, 0x2d, 0x6c, 0x53, 0xef, 0x5f, 0xfd, 0xd8, 0xe5, 0x97, 0xb8, 0x50,
	0xcb, 0x7d, 0x7e, 0xb9, 0xb0, 0x97, 0x5d, 0xfe, 0x97, 0x71, 0xd5, 0x8b, 0x8e, 0x83, 0x0b, 0xe5,
	0x38, 0x7b, 0x87, 0xe3, 0x2c, 0x6c, 0xc7, 0xe9, 0x3c, 0x9f, 0x4f, 0xbb, 0xc1, 0xde, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xf6, 0x79, 0x56, 0x03, 0xed, 0x02, 0x00, 0x00,
}
