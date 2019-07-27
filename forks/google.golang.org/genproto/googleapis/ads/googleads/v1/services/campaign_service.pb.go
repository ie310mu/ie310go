// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	_ "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/wrappers"
	resources "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api/annotations"
	status "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/rpc/status"
	field_mask "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/protobuf/field_mask"
	grpc "github.com/ie310mu/ie310go/forks/google.golang.org/grpc"
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

// Request message for
// [CampaignService.GetCampaign][google.ads.googleads.v1.services.CampaignService.GetCampaign].
type GetCampaignRequest struct {
	// The resource name of the campaign to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignRequest) Reset()         { *m = GetCampaignRequest{} }
func (m *GetCampaignRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignRequest) ProtoMessage()    {}
func (*GetCampaignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57837f1a83a55878, []int{0}
}

func (m *GetCampaignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignRequest.Unmarshal(m, b)
}
func (m *GetCampaignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignRequest.Merge(m, src)
}
func (m *GetCampaignRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignRequest.Size(m)
}
func (m *GetCampaignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignRequest proto.InternalMessageInfo

func (m *GetCampaignRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CampaignService.MutateCampaigns][google.ads.googleads.v1.services.CampaignService.MutateCampaigns].
type MutateCampaignsRequest struct {
	// The ID of the customer whose campaigns are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual campaigns.
	Operations []*CampaignOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignsRequest) Reset()         { *m = MutateCampaignsRequest{} }
func (m *MutateCampaignsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignsRequest) ProtoMessage()    {}
func (*MutateCampaignsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57837f1a83a55878, []int{1}
}

func (m *MutateCampaignsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignsRequest.Merge(m, src)
}
func (m *MutateCampaignsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignsRequest.Size(m)
}
func (m *MutateCampaignsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignsRequest proto.InternalMessageInfo

func (m *MutateCampaignsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignsRequest) GetOperations() []*CampaignOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign.
type CampaignOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignOperation_Create
	//	*CampaignOperation_Update
	//	*CampaignOperation_Remove
	Operation            isCampaignOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CampaignOperation) Reset()         { *m = CampaignOperation{} }
func (m *CampaignOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignOperation) ProtoMessage()    {}
func (*CampaignOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_57837f1a83a55878, []int{2}
}

func (m *CampaignOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignOperation.Unmarshal(m, b)
}
func (m *CampaignOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignOperation.Marshal(b, m, deterministic)
}
func (m *CampaignOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignOperation.Merge(m, src)
}
func (m *CampaignOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignOperation.Size(m)
}
func (m *CampaignOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignOperation proto.InternalMessageInfo

func (m *CampaignOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignOperation_Operation interface {
	isCampaignOperation_Operation()
}

type CampaignOperation_Create struct {
	Create *resources.Campaign `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignOperation_Update struct {
	Update *resources.Campaign `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignOperation_Create) isCampaignOperation_Operation() {}

func (*CampaignOperation_Update) isCampaignOperation_Operation() {}

func (*CampaignOperation_Remove) isCampaignOperation_Operation() {}

func (m *CampaignOperation) GetOperation() isCampaignOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignOperation) GetCreate() *resources.Campaign {
	if x, ok := m.GetOperation().(*CampaignOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignOperation) GetUpdate() *resources.Campaign {
	if x, ok := m.GetOperation().(*CampaignOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignOperation_Create)(nil),
		(*CampaignOperation_Update)(nil),
		(*CampaignOperation_Remove)(nil),
	}
}

// Response message for campaign mutate.
type MutateCampaignsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MutateCampaignsResponse) Reset()         { *m = MutateCampaignsResponse{} }
func (m *MutateCampaignsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignsResponse) ProtoMessage()    {}
func (*MutateCampaignsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57837f1a83a55878, []int{3}
}

func (m *MutateCampaignsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignsResponse.Merge(m, src)
}
func (m *MutateCampaignsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignsResponse.Size(m)
}
func (m *MutateCampaignsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignsResponse proto.InternalMessageInfo

func (m *MutateCampaignsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignsResponse) GetResults() []*MutateCampaignResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign mutate.
type MutateCampaignResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignResult) Reset()         { *m = MutateCampaignResult{} }
func (m *MutateCampaignResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignResult) ProtoMessage()    {}
func (*MutateCampaignResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_57837f1a83a55878, []int{4}
}

func (m *MutateCampaignResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignResult.Unmarshal(m, b)
}
func (m *MutateCampaignResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignResult.Merge(m, src)
}
func (m *MutateCampaignResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignResult.Size(m)
}
func (m *MutateCampaignResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignResult proto.InternalMessageInfo

func (m *MutateCampaignResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignRequest)(nil), "google.ads.googleads.v1.services.GetCampaignRequest")
	proto.RegisterType((*MutateCampaignsRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignsRequest")
	proto.RegisterType((*CampaignOperation)(nil), "google.ads.googleads.v1.services.CampaignOperation")
	proto.RegisterType((*MutateCampaignsResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignsResponse")
	proto.RegisterType((*MutateCampaignResult)(nil), "google.ads.googleads.v1.services.MutateCampaignResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_service.proto", fileDescriptor_57837f1a83a55878)
}

var fileDescriptor_57837f1a83a55878 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6b, 0xd4, 0x4e,
	0x18, 0xfe, 0x25, 0xfd, 0x51, 0xed, 0xa4, 0x5a, 0x1c, 0xab, 0x5d, 0x16, 0xd1, 0x25, 0x16, 0x2c,
	0x5b, 0x9c, 0xb8, 0xbb, 0xa2, 0x36, 0xa5, 0x87, 0xad, 0xf4, 0x8f, 0x87, 0xda, 0x92, 0x42, 0x0f,
	0xb2, 0x10, 0xa6, 0xc9, 0x34, 0x84, 0x26, 0x99, 0x38, 0x33, 0x59, 0x29, 0xa5, 0x17, 0xc1, 0x4f,
	0xe0, 0x37, 0x10, 0xbc, 0x78, 0xf5, 0x13, 0x78, 0xf5, 0xea, 0xd5, 0xa3, 0x27, 0xbf, 0x82, 0x08,
	0x92, 0x4c, 0x26, 0xdd, 0x6e, 0x5d, 0xd6, 0xf6, 0x36, 0x79, 0xf3, 0x3c, 0xcf, 0xfb, 0xcc, 0xfb,
	0x67, 0xc0, 0xd3, 0x80, 0xd2, 0x20, 0x22, 0x16, 0xf6, 0xb9, 0x25, 0x8f, 0xf9, 0xa9, 0xdf, 0xb2,
	0x38, 0x61, 0xfd, 0xd0, 0x23, 0xdc, 0xf2, 0x70, 0x9c, 0xe2, 0x30, 0x48, 0xdc, 0x32, 0x82, 0x52,
	0x46, 0x05, 0x85, 0x0d, 0x89, 0x46, 0xd8, 0xe7, 0xa8, 0x22, 0xa2, 0x7e, 0x0b, 0x29, 0x62, 0xfd,
	0xd1, 0x28, 0x69, 0x46, 0x38, 0xcd, 0xd8, 0xa0, 0xb6, 0xd4, 0xac, 0xdf, 0x51, 0x8c, 0x34, 0xb4,
	0x70, 0x92, 0x50, 0x81, 0x45, 0x48, 0x13, 0x5e, 0xfe, 0x2d, 0x33, 0x5a, 0xc5, 0xd7, 0x7e, 0x76,
	0x60, 0x1d, 0x84, 0x24, 0xf2, 0xdd, 0x18, 0xf3, 0xc3, 0x12, 0x71, 0x77, 0x18, 0xf1, 0x86, 0xe1,
	0x34, 0x25, 0x4c, 0x29, 0xcc, 0x95, 0xff, 0x59, 0xea, 0x59, 0x5c, 0x60, 0x91, 0x95, 0x3f, 0xcc,
	0x25, 0x00, 0x37, 0x88, 0x78, 0x5e, 0xba, 0x71, 0xc8, 0xeb, 0x8c, 0x70, 0x01, 0xef, 0x83, 0x6b,
	0xca, 0xaa, 0x9b, 0xe0, 0x98, 0xd4, 0xb4, 0x86, 0xb6, 0x30, 0xe5, 0x4c, 0xab, 0xe0, 0x4b, 0x1c,
	0x13, 0xf3, 0xbb, 0x06, 0x6e, 0x6f, 0x65, 0x02, 0x0b, 0xa2, 0xe8, 0x5c, 0xf1, 0xef, 0x01, 0xc3,
	0xcb, 0xb8, 0xa0, 0x31, 0x61, 0x6e, 0xe8, 0x97, 0x6c, 0xa0, 0x42, 0x2f, 0x7c, 0xb8, 0x0b, 0x00,
	0x4d, 0x09, 0x93, 0xb7, 0xac, 0xe9, 0x8d, 0x89, 0x05, 0xa3, 0xdd, 0x41, 0xe3, 0x0a, 0x8b, 0x54,
	0xa2, 0x6d, 0xc5, 0x75, 0x06, 0x64, 0xe0, 0x03, 0x30, 0x93, 0x62, 0x26, 0x42, 0x1c, 0xb9, 0x07,
	0x38, 0x8c, 0x32, 0x46, 0x6a, 0x13, 0x0d, 0x6d, 0xe1, 0xaa, 0x73, 0xbd, 0x0c, 0xaf, 0xcb, 0x68,
	0x7e, 0xbd, 0x3e, 0x8e, 0x42, 0x1f, 0x0b, 0xe2, 0xd2, 0x24, 0x3a, 0xaa, 0xfd, 0x5f, 0xc0, 0xa6,
	0x55, 0x70, 0x3b, 0x89, 0x8e, 0xcc, 0x77, 0x3a, 0xb8, 0x71, 0x2e, 0x1f, 0x5c, 0x06, 0x46, 0x96,
	0x16, 0xc4, 0xbc, 0xfa, 0x05, 0xd1, 0x68, 0xd7, 0x95, 0x73, 0x55, 0x7e, 0xb4, 0x9e, 0x37, 0x68,
	0x0b, 0xf3, 0x43, 0x07, 0x48, 0x78, 0x7e, 0x86, 0x6b, 0x60, 0xd2, 0x63, 0x04, 0x0b, 0x59, 0x4f,
	0xa3, 0xbd, 0x38, 0xf2, 0xc6, 0xd5, 0xa0, 0x54, 0x57, 0xde, 0xfc, 0xcf, 0x29, 0xc9, 0xb9, 0x8c,
	0x14, 0xad, 0xe9, 0x97, 0x92, 0x91, 0x64, 0x58, 0x03, 0x93, 0x8c, 0xc4, 0xb4, 0x2f, 0xab, 0x34,
	0x95, 0xff, 0x91, 0xdf, 0xab, 0x06, 0x98, 0xaa, 0xca, 0x6a, 0x7e, 0xd6, 0xc0, 0xdc, 0xb9, 0x36,
	0xf3, 0x94, 0x26, 0x9c, 0xc0, 0x75, 0x70, 0x6b, 0xa8, 0xe2, 0x2e, 0x61, 0x8c, 0xb2, 0x42, 0xd1,
	0x68, 0x43, 0x65, 0x8c, 0xa5, 0x1e, 0xda, 0x2d, 0xc6, 0xce, 0xb9, 0x79, 0xb6, 0x17, 0x6b, 0x39,
	0x1c, 0xee, 0x80, 0x2b, 0x8c, 0xf0, 0x2c, 0x12, 0x6a, 0x16, 0x9e, 0x8c, 0x9f, 0x85, 0xb3, 0x9e,
	0x9c, 0x82, 0xee, 0x28, 0x19, 0x73, 0x19, 0xcc, 0xfe, 0x0d, 0xf0, 0x4f, 0x93, 0xdd, 0xfe, 0xad,
	0x83, 0x19, 0xc5, 0xdb, 0x95, 0xf9, 0xe0, 0x47, 0x0d, 0x18, 0x03, 0x9b, 0x02, 0x1f, 0x8f, 0x77,
	0x78, 0x7e, 0xb1, 0xea, 0x17, 0x69, 0x95, 0xd9, 0x79, 0xfb, 0xed, 0xc7, 0x7b, 0xfd, 0x21, 0x5c,
	0xcc, 0x9f, 0x8e, 0xe3, 0x33, 0xb6, 0x57, 0xd4, 0x2e, 0x71, 0xab, 0x59, 0xbd, 0x25, 0xdc, 0x6a,
	0x9e, 0xc0, 0x2f, 0x1a, 0x98, 0x19, 0x6a, 0x17, 0x7c, 0x76, 0xd1, 0x6a, 0xaa, 0x45, 0xae, 0x2f,
	0x5d, 0x82, 0x29, 0x67, 0xc3, 0x5c, 0x2a, 0xdc, 0x77, 0x4c, 0x94, 0xbb, 0x3f, 0xb5, 0x7b, 0x3c,
	0xf0, 0x30, 0xac, 0x34, 0x4f, 0x4e, 0xcd, 0xdb, 0x71, 0x21, 0x64, 0x6b, 0xcd, 0xd5, 0x5f, 0x1a,
	0x98, 0xf7, 0x68, 0x3c, 0x36, 0xf7, 0xea, 0xec, 0x50, 0x97, 0x76, 0xf2, 0xfd, 0xdb, 0xd1, 0x5e,
	0x6d, 0x96, 0xcc, 0x80, 0x46, 0x38, 0x09, 0x10, 0x65, 0x81, 0x15, 0x90, 0xa4, 0xd8, 0x4e, 0xf5,
	0x20, 0xa7, 0x21, 0x1f, 0xfd, 0xf4, 0x2f, 0xab, 0xc3, 0x07, 0x7d, 0x62, 0xa3, 0xdb, 0xfd, 0xa4,
	0x37, 0x36, 0xa4, 0x60, 0xd7, 0xe7, 0x48, 0x1e, 0xf3, 0xd3, 0x5e, 0x0b, 0x95, 0x89, 0xf9, 0x57,
	0x05, 0xe9, 0x75, 0x7d, 0xde, 0xab, 0x20, 0xbd, 0xbd, 0x56, 0x4f, 0x41, 0x7e, 0xea, 0xf3, 0x32,
	0x6e, 0xdb, 0x5d, 0x9f, 0xdb, 0x76, 0x05, 0xb2, 0xed, 0xbd, 0x96, 0x6d, 0x2b, 0xd8, 0xfe, 0x64,
	0xe1, 0xb3, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x56, 0xbf, 0x71, 0xcc, 0xa1, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignServiceClient is the client API for CampaignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/ie310mu/ie310go/forks/google.golang.org/grpc#ClientConn.NewStream.
type CampaignServiceClient interface {
	// Returns the requested campaign in full detail.
	GetCampaign(ctx context.Context, in *GetCampaignRequest, opts ...grpc.CallOption) (*resources.Campaign, error)
	// Creates, updates, or removes campaigns. Operation statuses are returned.
	MutateCampaigns(ctx context.Context, in *MutateCampaignsRequest, opts ...grpc.CallOption) (*MutateCampaignsResponse, error)
}

type campaignServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignServiceClient(cc *grpc.ClientConn) CampaignServiceClient {
	return &campaignServiceClient{cc}
}

func (c *campaignServiceClient) GetCampaign(ctx context.Context, in *GetCampaignRequest, opts ...grpc.CallOption) (*resources.Campaign, error) {
	out := new(resources.Campaign)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignService/GetCampaign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignServiceClient) MutateCampaigns(ctx context.Context, in *MutateCampaignsRequest, opts ...grpc.CallOption) (*MutateCampaignsResponse, error) {
	out := new(MutateCampaignsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignService/MutateCampaigns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignServiceServer is the server API for CampaignService service.
type CampaignServiceServer interface {
	// Returns the requested campaign in full detail.
	GetCampaign(context.Context, *GetCampaignRequest) (*resources.Campaign, error)
	// Creates, updates, or removes campaigns. Operation statuses are returned.
	MutateCampaigns(context.Context, *MutateCampaignsRequest) (*MutateCampaignsResponse, error)
}

func RegisterCampaignServiceServer(s *grpc.Server, srv CampaignServiceServer) {
	s.RegisterService(&_CampaignService_serviceDesc, srv)
}

func _CampaignService_GetCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).GetCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignService/GetCampaign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).GetCampaign(ctx, req.(*GetCampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignService_MutateCampaigns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).MutateCampaigns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignService/MutateCampaigns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).MutateCampaigns(ctx, req.(*MutateCampaignsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignService",
	HandlerType: (*CampaignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaign",
			Handler:    _CampaignService_GetCampaign_Handler,
		},
		{
			MethodName: "MutateCampaigns",
			Handler:    _CampaignService_MutateCampaigns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_service.proto",
}
