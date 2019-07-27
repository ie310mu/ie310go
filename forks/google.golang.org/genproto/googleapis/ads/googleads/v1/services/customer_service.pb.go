// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/customer_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	_ "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/wrappers"
	resources "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api/annotations"
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
// [CustomerService.GetCustomer][google.ads.googleads.v1.services.CustomerService.GetCustomer].
type GetCustomerRequest struct {
	// The resource name of the customer to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerRequest) Reset()         { *m = GetCustomerRequest{} }
func (m *GetCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerRequest) ProtoMessage()    {}
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{0}
}

func (m *GetCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerRequest.Unmarshal(m, b)
}
func (m *GetCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerRequest.Merge(m, src)
}
func (m *GetCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerRequest.Size(m)
}
func (m *GetCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerRequest proto.InternalMessageInfo

func (m *GetCustomerRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CustomerService.MutateCustomer][google.ads.googleads.v1.services.CustomerService.MutateCustomer].
type MutateCustomerRequest struct {
	// The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform on the customer
	Operation *CustomerOperation `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,5,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerRequest) Reset()         { *m = MutateCustomerRequest{} }
func (m *MutateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerRequest) ProtoMessage()    {}
func (*MutateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{1}
}

func (m *MutateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerRequest.Unmarshal(m, b)
}
func (m *MutateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerRequest.Merge(m, src)
}
func (m *MutateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerRequest.Size(m)
}
func (m *MutateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerRequest proto.InternalMessageInfo

func (m *MutateCustomerRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerRequest) GetOperation() *CustomerOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MutateCustomerRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// Request message for
// [CustomerService.CreateCustomerClient][google.ads.googleads.v1.services.CustomerService.CreateCustomerClient].
type CreateCustomerClientRequest struct {
	// The ID of the Manager under whom client customer is being created.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The new client customer to create. The resource name on this customer
	// will be ignored.
	CustomerClient       *resources.Customer `protobuf:"bytes,2,opt,name=customer_client,json=customerClient,proto3" json:"customer_client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateCustomerClientRequest) Reset()         { *m = CreateCustomerClientRequest{} }
func (m *CreateCustomerClientRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerClientRequest) ProtoMessage()    {}
func (*CreateCustomerClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{2}
}

func (m *CreateCustomerClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerClientRequest.Unmarshal(m, b)
}
func (m *CreateCustomerClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerClientRequest.Marshal(b, m, deterministic)
}
func (m *CreateCustomerClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerClientRequest.Merge(m, src)
}
func (m *CreateCustomerClientRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerClientRequest.Size(m)
}
func (m *CreateCustomerClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerClientRequest proto.InternalMessageInfo

func (m *CreateCustomerClientRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *CreateCustomerClientRequest) GetCustomerClient() *resources.Customer {
	if m != nil {
		return m.CustomerClient
	}
	return nil
}

// A single update on a customer.
type CustomerOperation struct {
	// Mutate operation. Only updates are supported for customer.
	Update *resources.Customer `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomerOperation) Reset()         { *m = CustomerOperation{} }
func (m *CustomerOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerOperation) ProtoMessage()    {}
func (*CustomerOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{3}
}

func (m *CustomerOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerOperation.Unmarshal(m, b)
}
func (m *CustomerOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerOperation.Marshal(b, m, deterministic)
}
func (m *CustomerOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerOperation.Merge(m, src)
}
func (m *CustomerOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerOperation.Size(m)
}
func (m *CustomerOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerOperation proto.InternalMessageInfo

func (m *CustomerOperation) GetUpdate() *resources.Customer {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *CustomerOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Response message for CreateCustomerClient mutate.
type CreateCustomerClientResponse struct {
	// The resource name of the newly created customer client.
	ResourceName         string   `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCustomerClientResponse) Reset()         { *m = CreateCustomerClientResponse{} }
func (m *CreateCustomerClientResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerClientResponse) ProtoMessage()    {}
func (*CreateCustomerClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{4}
}

func (m *CreateCustomerClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerClientResponse.Unmarshal(m, b)
}
func (m *CreateCustomerClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerClientResponse.Marshal(b, m, deterministic)
}
func (m *CreateCustomerClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerClientResponse.Merge(m, src)
}
func (m *CreateCustomerClientResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerClientResponse.Size(m)
}
func (m *CreateCustomerClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerClientResponse proto.InternalMessageInfo

func (m *CreateCustomerClientResponse) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Response message for customer mutate.
type MutateCustomerResponse struct {
	// Result for the mutate.
	Result               *MutateCustomerResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MutateCustomerResponse) Reset()         { *m = MutateCustomerResponse{} }
func (m *MutateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerResponse) ProtoMessage()    {}
func (*MutateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{5}
}

func (m *MutateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerResponse.Unmarshal(m, b)
}
func (m *MutateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerResponse.Merge(m, src)
}
func (m *MutateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerResponse.Size(m)
}
func (m *MutateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerResponse proto.InternalMessageInfo

func (m *MutateCustomerResponse) GetResult() *MutateCustomerResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for the customer mutate.
type MutateCustomerResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerResult) Reset()         { *m = MutateCustomerResult{} }
func (m *MutateCustomerResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerResult) ProtoMessage()    {}
func (*MutateCustomerResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{6}
}

func (m *MutateCustomerResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerResult.Unmarshal(m, b)
}
func (m *MutateCustomerResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerResult.Merge(m, src)
}
func (m *MutateCustomerResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerResult.Size(m)
}
func (m *MutateCustomerResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerResult proto.InternalMessageInfo

func (m *MutateCustomerResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CustomerService.ListAccessibleCustomers][google.ads.googleads.v1.services.CustomerService.ListAccessibleCustomers].
type ListAccessibleCustomersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessibleCustomersRequest) Reset()         { *m = ListAccessibleCustomersRequest{} }
func (m *ListAccessibleCustomersRequest) String() string { return proto.CompactTextString(m) }
func (*ListAccessibleCustomersRequest) ProtoMessage()    {}
func (*ListAccessibleCustomersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{7}
}

func (m *ListAccessibleCustomersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Unmarshal(m, b)
}
func (m *ListAccessibleCustomersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Marshal(b, m, deterministic)
}
func (m *ListAccessibleCustomersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessibleCustomersRequest.Merge(m, src)
}
func (m *ListAccessibleCustomersRequest) XXX_Size() int {
	return xxx_messageInfo_ListAccessibleCustomersRequest.Size(m)
}
func (m *ListAccessibleCustomersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessibleCustomersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessibleCustomersRequest proto.InternalMessageInfo

// Response message for
// [CustomerService.ListAccessibleCustomers][google.ads.googleads.v1.services.CustomerService.ListAccessibleCustomers].
type ListAccessibleCustomersResponse struct {
	// Resource name of customers directly accessible by the
	// user authenticating the call.
	ResourceNames        []string `protobuf:"bytes,1,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccessibleCustomersResponse) Reset()         { *m = ListAccessibleCustomersResponse{} }
func (m *ListAccessibleCustomersResponse) String() string { return proto.CompactTextString(m) }
func (*ListAccessibleCustomersResponse) ProtoMessage()    {}
func (*ListAccessibleCustomersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35dc615eb25453ea, []int{8}
}

func (m *ListAccessibleCustomersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Unmarshal(m, b)
}
func (m *ListAccessibleCustomersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Marshal(b, m, deterministic)
}
func (m *ListAccessibleCustomersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccessibleCustomersResponse.Merge(m, src)
}
func (m *ListAccessibleCustomersResponse) XXX_Size() int {
	return xxx_messageInfo_ListAccessibleCustomersResponse.Size(m)
}
func (m *ListAccessibleCustomersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccessibleCustomersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccessibleCustomersResponse proto.InternalMessageInfo

func (m *ListAccessibleCustomersResponse) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCustomerRequest)(nil), "google.ads.googleads.v1.services.GetCustomerRequest")
	proto.RegisterType((*MutateCustomerRequest)(nil), "google.ads.googleads.v1.services.MutateCustomerRequest")
	proto.RegisterType((*CreateCustomerClientRequest)(nil), "google.ads.googleads.v1.services.CreateCustomerClientRequest")
	proto.RegisterType((*CustomerOperation)(nil), "google.ads.googleads.v1.services.CustomerOperation")
	proto.RegisterType((*CreateCustomerClientResponse)(nil), "google.ads.googleads.v1.services.CreateCustomerClientResponse")
	proto.RegisterType((*MutateCustomerResponse)(nil), "google.ads.googleads.v1.services.MutateCustomerResponse")
	proto.RegisterType((*MutateCustomerResult)(nil), "google.ads.googleads.v1.services.MutateCustomerResult")
	proto.RegisterType((*ListAccessibleCustomersRequest)(nil), "google.ads.googleads.v1.services.ListAccessibleCustomersRequest")
	proto.RegisterType((*ListAccessibleCustomersResponse)(nil), "google.ads.googleads.v1.services.ListAccessibleCustomersResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/customer_service.proto", fileDescriptor_35dc615eb25453ea)
}

var fileDescriptor_35dc615eb25453ea = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xd5, 0xa4, 0xdf, 0x17, 0xd1, 0x49, 0x7f, 0xc4, 0xa8, 0x40, 0x14, 0xaa, 0x36, 0x98, 0x56,
	0x0d, 0x41, 0xd8, 0xa4, 0x45, 0x14, 0x5c, 0x05, 0xe1, 0x46, 0xa2, 0x45, 0xa2, 0x3f, 0x18, 0xd4,
	0x05, 0x8a, 0x14, 0xb9, 0xf6, 0x34, 0x58, 0x75, 0x3c, 0xc6, 0x63, 0x07, 0x55, 0x55, 0x37, 0x6c,
	0x59, 0x02, 0x0b, 0xb6, 0x2c, 0x59, 0xb1, 0xe0, 0x29, 0x90, 0x58, 0xc1, 0x23, 0xb0, 0xe2, 0x15,
	0xd8, 0x20, 0x8f, 0x67, 0xdc, 0xfc, 0xb9, 0x69, 0xcb, 0xee, 0x66, 0x72, 0xcf, 0xb9, 0xe7, 0xde,
	0x39, 0x73, 0x0d, 0x97, 0x9b, 0x84, 0x34, 0x1d, 0xac, 0x18, 0x16, 0x55, 0xe2, 0x30, 0x8a, 0xda,
	0x15, 0x85, 0x62, 0xbf, 0x6d, 0x9b, 0x98, 0x2a, 0x66, 0x48, 0x03, 0xd2, 0xc2, 0x7e, 0x83, 0x9f,
	0xc8, 0x9e, 0x4f, 0x02, 0x82, 0x8a, 0x71, 0xb6, 0x6c, 0x58, 0x54, 0x4e, 0x80, 0x72, 0xbb, 0x22,
	0x0b, 0x60, 0xe1, 0x76, 0x1a, 0xb5, 0x8f, 0x29, 0x09, 0xfd, 0x4e, 0xee, 0x98, 0xb3, 0x30, 0x2d,
	0x10, 0x9e, 0xad, 0x18, 0xae, 0x4b, 0x02, 0x23, 0xb0, 0x89, 0x4b, 0xf9, 0xbf, 0xbc, 0xa2, 0xc2,
	0x7e, 0xed, 0x86, 0x7b, 0xca, 0x9e, 0x8d, 0x1d, 0xab, 0xd1, 0x32, 0xe8, 0x3e, 0xcf, 0x98, 0xe9,
	0xcd, 0x78, 0xed, 0x1b, 0x9e, 0x87, 0x7d, 0xce, 0x20, 0xdd, 0x87, 0x68, 0x0d, 0x07, 0x35, 0x5e,
	0x54, 0xc7, 0xaf, 0x42, 0x4c, 0x03, 0x74, 0x1d, 0x8e, 0x0b, 0x45, 0x0d, 0xd7, 0x68, 0xe1, 0x3c,
	0x28, 0x82, 0xd2, 0xa8, 0x3e, 0x26, 0x0e, 0x37, 0x8d, 0x16, 0x96, 0xbe, 0x00, 0x78, 0x69, 0x23,
	0x0c, 0x8c, 0x00, 0xf7, 0xc2, 0x67, 0x61, 0x2e, 0x19, 0x91, 0x6d, 0x71, 0x30, 0x14, 0x47, 0x8f,
	0x2d, 0xf4, 0x14, 0x8e, 0x12, 0x0f, 0xfb, 0xac, 0x97, 0xfc, 0x7f, 0x45, 0x50, 0xca, 0x2d, 0x2e,
	0xc9, 0xc3, 0xa6, 0x27, 0x8b, 0x32, 0x5b, 0x02, 0xaa, 0x1f, 0xb3, 0x44, 0x92, 0xdb, 0x86, 0x63,
	0x5b, 0x46, 0x80, 0x1b, 0xc4, 0x75, 0x0e, 0xf2, 0xff, 0x17, 0x41, 0xe9, 0x82, 0x3e, 0x26, 0x0e,
	0xb7, 0x5c, 0xe7, 0x40, 0x7a, 0x0f, 0xe0, 0xd5, 0x9a, 0x8f, 0x3b, 0x24, 0xd7, 0x1c, 0x1b, 0xbb,
	0xc1, 0xa9, 0x85, 0x3f, 0x87, 0x93, 0x49, 0x82, 0xc9, 0xa0, 0xf9, 0x0c, 0x93, 0x7f, 0x33, 0x55,
	0x7e, 0x72, 0xb5, 0x89, 0x7e, 0x7d, 0xc2, 0xec, 0xaa, 0x2e, 0x7d, 0x00, 0xf0, 0x62, 0x5f, 0x73,
	0xa8, 0x06, 0xb3, 0xa1, 0x17, 0x49, 0x67, 0x3a, 0xce, 0x58, 0x82, 0x43, 0xd1, 0x0a, 0xcc, 0xc5,
	0x11, 0x33, 0x05, 0x17, 0x5b, 0x10, 0x4c, 0xc2, 0x15, 0xf2, 0xa3, 0xc8, 0x37, 0x1b, 0x06, 0xdd,
	0xd7, 0x61, 0x9c, 0x1e, 0xc5, 0x52, 0x0d, 0x4e, 0x0f, 0x9e, 0x16, 0xf5, 0x88, 0x4b, 0x71, 0xbf,
	0x4d, 0x32, 0x03, 0x6c, 0xf2, 0x12, 0x5e, 0xee, 0x75, 0x09, 0x87, 0x6f, 0xc2, 0xac, 0x8f, 0x69,
	0xe8, 0x88, 0x19, 0xde, 0x1d, 0x6e, 0x81, 0x3e, 0xa6, 0xd0, 0x09, 0x74, 0xce, 0x22, 0xad, 0xc0,
	0xa9, 0x41, 0xff, 0x9f, 0xce, 0xcd, 0x45, 0x38, 0xf3, 0xc4, 0xa6, 0x81, 0x66, 0x9a, 0x98, 0x52,
	0x7b, 0xd7, 0x49, 0x48, 0x28, 0x37, 0x87, 0xb4, 0x0e, 0x67, 0x53, 0x33, 0x78, 0x47, 0xf3, 0x70,
	0xa2, 0xab, 0x12, 0xcd, 0x83, 0xe2, 0x48, 0x69, 0x54, 0x1f, 0xef, 0x2c, 0x45, 0x17, 0xdf, 0x66,
	0xe1, 0xa4, 0x00, 0x3f, 0x8b, 0x5b, 0x43, 0x1f, 0x01, 0xcc, 0x75, 0xbc, 0x44, 0x74, 0x67, 0xf8,
	0x30, 0xfa, 0x1f, 0x6e, 0xe1, 0x2c, 0x1e, 0x91, 0x16, 0xde, 0xfc, 0xf8, 0xf5, 0x2e, 0x73, 0x0d,
	0xcd, 0x46, 0x1b, 0xe8, 0xb0, 0x4b, 0x78, 0x55, 0x38, 0x94, 0x2a, 0xe5, 0x23, 0xf4, 0x15, 0xc0,
	0x89, 0xee, 0xc9, 0xa2, 0xe5, 0xb3, 0xdf, 0x55, 0xac, 0xf0, 0xde, 0x39, 0x2e, 0x99, 0x0d, 0x57,
	0x52, 0x98, 0xdc, 0x1b, 0xd2, 0x5c, 0x24, 0xf7, 0x58, 0xdf, 0x61, 0xc7, 0x8b, 0xad, 0x96, 0x8f,
	0xd4, 0x16, 0x43, 0xab, 0xa0, 0x8c, 0xbe, 0x03, 0x78, 0x25, 0xe5, 0xc6, 0xd0, 0xc3, 0xe1, 0x32,
	0x4e, 0xb6, 0x43, 0x41, 0xfb, 0x07, 0x06, 0xde, 0xd1, 0x2d, 0xd6, 0xd1, 0x02, 0x9a, 0xef, 0xea,
	0x48, 0x75, 0x52, 0x34, 0xff, 0x04, 0x70, 0x6a, 0xd0, 0x7b, 0x44, 0xd5, 0x53, 0xec, 0xce, 0xf4,
	0xad, 0x57, 0x78, 0x70, 0x5e, 0x38, 0x6f, 0xa3, 0xca, 0xda, 0x58, 0x96, 0x16, 0x4f, 0xbe, 0x18,
	0x73, 0x00, 0x87, 0x0a, 0xca, 0xab, 0x7f, 0x00, 0x9c, 0x33, 0x49, 0x6b, 0xa8, 0x88, 0xd5, 0xa9,
	0x9e, 0x37, 0xb3, 0x1d, 0x6d, 0xaf, 0x6d, 0xf0, 0x62, 0x9d, 0x23, 0x9b, 0xc4, 0x31, 0xdc, 0xa6,
	0x4c, 0xfc, 0xa6, 0xd2, 0xc4, 0x2e, 0xdb, 0x6d, 0xe2, 0x2b, 0xeb, 0xd9, 0x34, 0xfd, 0x7b, 0xbe,
	0x22, 0x82, 0x4f, 0x99, 0x91, 0x35, 0x4d, 0xfb, 0x9c, 0x29, 0xae, 0xc5, 0x84, 0x9a, 0x45, 0xe5,
	0x38, 0x8c, 0xa2, 0x9d, 0x8a, 0xcc, 0x0b, 0xd3, 0x6f, 0x22, 0xa5, 0xae, 0x59, 0xb4, 0x9e, 0xa4,
	0xd4, 0x77, 0x2a, 0x75, 0x91, 0xf2, 0x3b, 0x33, 0x17, 0x9f, 0xab, 0xaa, 0x66, 0x51, 0x55, 0x4d,
	0x92, 0x54, 0x75, 0xa7, 0xa2, 0xaa, 0x22, 0x6d, 0x37, 0xcb, 0x74, 0x2e, 0xfd, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x0a, 0xed, 0x23, 0x76, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/ie310mu/ie310go/forks/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	// Returns the requested customer in full detail.
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*resources.Customer, error)
	// Updates a customer. Operation statuses are returned.
	MutateCustomer(ctx context.Context, in *MutateCustomerRequest, opts ...grpc.CallOption) (*MutateCustomerResponse, error)
	// Returns resource names of customers directly accessible by the
	// user authenticating the call.
	ListAccessibleCustomers(ctx context.Context, in *ListAccessibleCustomersRequest, opts ...grpc.CallOption) (*ListAccessibleCustomersResponse, error)
	// Creates a new client under manager. The new client customer is returned.
	CreateCustomerClient(ctx context.Context, in *CreateCustomerClientRequest, opts ...grpc.CallOption) (*CreateCustomerClientResponse, error)
}

type customerServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerServiceClient(cc *grpc.ClientConn) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*resources.Customer, error) {
	out := new(resources.Customer)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) MutateCustomer(ctx context.Context, in *MutateCustomerRequest, opts ...grpc.CallOption) (*MutateCustomerResponse, error) {
	out := new(MutateCustomerResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerService/MutateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) ListAccessibleCustomers(ctx context.Context, in *ListAccessibleCustomersRequest, opts ...grpc.CallOption) (*ListAccessibleCustomersResponse, error) {
	out := new(ListAccessibleCustomersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerService/ListAccessibleCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateCustomerClient(ctx context.Context, in *CreateCustomerClientRequest, opts ...grpc.CallOption) (*CreateCustomerClientResponse, error) {
	out := new(CreateCustomerClientResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerService/CreateCustomerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	// Returns the requested customer in full detail.
	GetCustomer(context.Context, *GetCustomerRequest) (*resources.Customer, error)
	// Updates a customer. Operation statuses are returned.
	MutateCustomer(context.Context, *MutateCustomerRequest) (*MutateCustomerResponse, error)
	// Returns resource names of customers directly accessible by the
	// user authenticating the call.
	ListAccessibleCustomers(context.Context, *ListAccessibleCustomersRequest) (*ListAccessibleCustomersResponse, error)
	// Creates a new client under manager. The new client customer is returned.
	CreateCustomerClient(context.Context, *CreateCustomerClientRequest) (*CreateCustomerClientResponse, error)
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_MutateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).MutateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerService/MutateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).MutateCustomer(ctx, req.(*MutateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_ListAccessibleCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccessibleCustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).ListAccessibleCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerService/ListAccessibleCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).ListAccessibleCustomers(ctx, req.(*ListAccessibleCustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerService/CreateCustomerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomerClient(ctx, req.(*CreateCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
		{
			MethodName: "MutateCustomer",
			Handler:    _CustomerService_MutateCustomer_Handler,
		},
		{
			MethodName: "ListAccessibleCustomers",
			Handler:    _CustomerService_ListAccessibleCustomers_Handler,
		},
		{
			MethodName: "CreateCustomerClient",
			Handler:    _CustomerService_CreateCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/customer_service.proto",
}
