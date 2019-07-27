// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/billing_setup_status.proto

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

// The possible statuses of a BillingSetup.
type BillingSetupStatusEnum_BillingSetupStatus int32

const (
	// Not specified.
	BillingSetupStatusEnum_UNSPECIFIED BillingSetupStatusEnum_BillingSetupStatus = 0
	// Used for return value only. Represents value unknown in this version.
	BillingSetupStatusEnum_UNKNOWN BillingSetupStatusEnum_BillingSetupStatus = 1
	// The billing setup is pending approval.
	BillingSetupStatusEnum_PENDING BillingSetupStatusEnum_BillingSetupStatus = 2
	// The billing setup has been approved but the corresponding first budget
	// has not.  This can only occur for billing setups configured for monthly
	// invoicing.
	BillingSetupStatusEnum_APPROVED_HELD BillingSetupStatusEnum_BillingSetupStatus = 3
	// The billing setup has been approved.
	BillingSetupStatusEnum_APPROVED BillingSetupStatusEnum_BillingSetupStatus = 4
	// The billing setup was cancelled by the user prior to approval.
	BillingSetupStatusEnum_CANCELLED BillingSetupStatusEnum_BillingSetupStatus = 5
)

var BillingSetupStatusEnum_BillingSetupStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "APPROVED_HELD",
	4: "APPROVED",
	5: "CANCELLED",
}

var BillingSetupStatusEnum_BillingSetupStatus_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"PENDING":       2,
	"APPROVED_HELD": 3,
	"APPROVED":      4,
	"CANCELLED":     5,
}

func (x BillingSetupStatusEnum_BillingSetupStatus) String() string {
	return proto.EnumName(BillingSetupStatusEnum_BillingSetupStatus_name, int32(x))
}

func (BillingSetupStatusEnum_BillingSetupStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1a153c8a5f7bc531, []int{0, 0}
}

// Message describing BillingSetup statuses.
type BillingSetupStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillingSetupStatusEnum) Reset()         { *m = BillingSetupStatusEnum{} }
func (m *BillingSetupStatusEnum) String() string { return proto.CompactTextString(m) }
func (*BillingSetupStatusEnum) ProtoMessage()    {}
func (*BillingSetupStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a153c8a5f7bc531, []int{0}
}

func (m *BillingSetupStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetupStatusEnum.Unmarshal(m, b)
}
func (m *BillingSetupStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetupStatusEnum.Marshal(b, m, deterministic)
}
func (m *BillingSetupStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetupStatusEnum.Merge(m, src)
}
func (m *BillingSetupStatusEnum) XXX_Size() int {
	return xxx_messageInfo_BillingSetupStatusEnum.Size(m)
}
func (m *BillingSetupStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetupStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetupStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.BillingSetupStatusEnum_BillingSetupStatus", BillingSetupStatusEnum_BillingSetupStatus_name, BillingSetupStatusEnum_BillingSetupStatus_value)
	proto.RegisterType((*BillingSetupStatusEnum)(nil), "google.ads.googleads.v1.enums.BillingSetupStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/billing_setup_status.proto", fileDescriptor_1a153c8a5f7bc531)
}

var fileDescriptor_1a153c8a5f7bc531 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xfb, 0x30,
	0x18, 0xfd, 0xb5, 0xfb, 0xf9, 0x2f, 0x73, 0x58, 0x73, 0xa1, 0x20, 0xee, 0x62, 0x7b, 0x80, 0x94,
	0xe2, 0x8d, 0xc4, 0xab, 0x74, 0x8d, 0x73, 0x38, 0xb2, 0xe2, 0x58, 0x05, 0x29, 0x8c, 0xce, 0x96,
	0x50, 0xe8, 0x92, 0xb2, 0xb4, 0x7b, 0x07, 0x5f, 0xc3, 0x4b, 0x1f, 0xc5, 0x47, 0x11, 0x1f, 0x42,
	0x9a, 0xae, 0xbd, 0x19, 0x7a, 0x13, 0x4e, 0xbe, 0xf3, 0x9d, 0xc3, 0x77, 0x0e, 0xb8, 0xe5, 0x52,
	0xf2, 0x2c, 0xb1, 0xa3, 0x58, 0xd9, 0x35, 0xac, 0xd0, 0xd6, 0xb1, 0x13, 0x51, 0xae, 0x95, 0xbd,
	0x4a, 0xb3, 0x2c, 0x15, 0x7c, 0xa9, 0x92, 0xa2, 0xcc, 0x97, 0xaa, 0x88, 0x8a, 0x52, 0xa1, 0x7c,
	0x23, 0x0b, 0x09, 0xfb, 0xf5, 0x3a, 0x8a, 0x62, 0x85, 0x5a, 0x25, 0xda, 0x3a, 0x48, 0x2b, 0xaf,
	0xae, 0x1b, 0xe3, 0x3c, 0xb5, 0x23, 0x21, 0x64, 0x11, 0x15, 0xa9, 0x14, 0x3b, 0xf1, 0xf0, 0xcd,
	0x00, 0x17, 0x6e, 0xed, 0x3d, 0xaf, 0xac, 0xe7, 0xda, 0x99, 0x8a, 0x72, 0x3d, 0x94, 0x00, 0xee,
	0x33, 0xf0, 0x0c, 0x74, 0x17, 0x6c, 0xee, 0xd3, 0xd1, 0xe4, 0x7e, 0x42, 0x3d, 0xeb, 0x1f, 0xec,
	0x82, 0xa3, 0x05, 0x7b, 0x64, 0xb3, 0x67, 0x66, 0x19, 0xd5, 0xc7, 0xa7, 0xcc, 0x9b, 0xb0, 0xb1,
	0x65, 0xc2, 0x73, 0xd0, 0x23, 0xbe, 0xff, 0x34, 0x0b, 0xa8, 0xb7, 0x7c, 0xa0, 0x53, 0xcf, 0xea,
	0xc0, 0x53, 0x70, 0xdc, 0x8c, 0xac, 0xff, 0xb0, 0x07, 0x4e, 0x46, 0x84, 0x8d, 0xe8, 0x74, 0x4a,
	0x3d, 0xeb, 0xc0, 0xfd, 0x36, 0xc0, 0xe0, 0x55, 0xae, 0xd1, 0x9f, 0x79, 0xdc, 0xcb, 0xfd, 0xa3,
	0xfc, 0x2a, 0x8a, 0x6f, 0xbc, 0xb8, 0x3b, 0x25, 0x97, 0x59, 0x24, 0x38, 0x92, 0x1b, 0x6e, 0xf3,
	0x44, 0xe8, 0xa0, 0x4d, 0xa7, 0x79, 0xaa, 0x7e, 0xa9, 0xf8, 0x4e, 0xbf, 0xef, 0x66, 0x67, 0x4c,
	0xc8, 0x87, 0xd9, 0x1f, 0xd7, 0x56, 0x24, 0x56, 0xa8, 0x86, 0x15, 0x0a, 0x1c, 0x54, 0x55, 0xa3,
	0x3e, 0x1b, 0x3e, 0x24, 0xb1, 0x0a, 0x5b, 0x3e, 0x0c, 0x9c, 0x50, 0xf3, 0x5f, 0xe6, 0xa0, 0x1e,
	0x62, 0x4c, 0x62, 0x85, 0x71, 0xbb, 0x81, 0x71, 0xe0, 0x60, 0xac, 0x77, 0x56, 0x87, 0xfa, 0xb0,
	0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x95, 0x7e, 0x6c, 0xfa, 0x01, 0x00, 0x00,
}
