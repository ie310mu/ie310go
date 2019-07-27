// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/tenant_service.proto

package talent

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	empty "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/empty"
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

// The Request of the CreateTenant method.
type CreateTenantRequest struct {
	// Required.
	//
	// Resource name of the project under which the tenant is created.
	//
	// The format is "projects/{project_id}", for example,
	// "projects/api-test-project".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required.
	//
	// The tenant to be created.
	Tenant               *Tenant  `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTenantRequest) Reset()         { *m = CreateTenantRequest{} }
func (m *CreateTenantRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTenantRequest) ProtoMessage()    {}
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df49c13fa187082, []int{0}
}

func (m *CreateTenantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTenantRequest.Unmarshal(m, b)
}
func (m *CreateTenantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTenantRequest.Marshal(b, m, deterministic)
}
func (m *CreateTenantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTenantRequest.Merge(m, src)
}
func (m *CreateTenantRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTenantRequest.Size(m)
}
func (m *CreateTenantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTenantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTenantRequest proto.InternalMessageInfo

func (m *CreateTenantRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateTenantRequest) GetTenant() *Tenant {
	if m != nil {
		return m.Tenant
	}
	return nil
}

// Request for getting a tenant by name.
type GetTenantRequest struct {
	// Required.
	//
	// The resource name of the tenant to be retrieved.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/api-test-project/tenants/foo".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTenantRequest) Reset()         { *m = GetTenantRequest{} }
func (m *GetTenantRequest) String() string { return proto.CompactTextString(m) }
func (*GetTenantRequest) ProtoMessage()    {}
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df49c13fa187082, []int{1}
}

func (m *GetTenantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTenantRequest.Unmarshal(m, b)
}
func (m *GetTenantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTenantRequest.Marshal(b, m, deterministic)
}
func (m *GetTenantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTenantRequest.Merge(m, src)
}
func (m *GetTenantRequest) XXX_Size() int {
	return xxx_messageInfo_GetTenantRequest.Size(m)
}
func (m *GetTenantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTenantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTenantRequest proto.InternalMessageInfo

func (m *GetTenantRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for updating a specified tenant.
type UpdateTenantRequest struct {
	// Required.
	//
	// The tenant resource to replace the current resource in the system.
	Tenant *Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// Optional but strongly recommended for the best service
	// experience.
	//
	// If [update_mask][google.cloud.talent.v4beta1.UpdateTenantRequest.update_mask] is provided, only the specified fields in
	// [tenant][google.cloud.talent.v4beta1.UpdateTenantRequest.tenant] are updated. Otherwise all the fields are updated.
	//
	// A field mask to specify the tenant fields to be updated. Only
	// top level fields of [Tenant][google.cloud.talent.v4beta1.Tenant] are supported.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateTenantRequest) Reset()         { *m = UpdateTenantRequest{} }
func (m *UpdateTenantRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTenantRequest) ProtoMessage()    {}
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df49c13fa187082, []int{2}
}

func (m *UpdateTenantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTenantRequest.Unmarshal(m, b)
}
func (m *UpdateTenantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTenantRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTenantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTenantRequest.Merge(m, src)
}
func (m *UpdateTenantRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTenantRequest.Size(m)
}
func (m *UpdateTenantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTenantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTenantRequest proto.InternalMessageInfo

func (m *UpdateTenantRequest) GetTenant() *Tenant {
	if m != nil {
		return m.Tenant
	}
	return nil
}

func (m *UpdateTenantRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request to delete a tenant.
type DeleteTenantRequest struct {
	// Required.
	//
	// The resource name of the tenant to be deleted.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/api-test-project/tenants/foo".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTenantRequest) Reset()         { *m = DeleteTenantRequest{} }
func (m *DeleteTenantRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTenantRequest) ProtoMessage()    {}
func (*DeleteTenantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df49c13fa187082, []int{3}
}

func (m *DeleteTenantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTenantRequest.Unmarshal(m, b)
}
func (m *DeleteTenantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTenantRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTenantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTenantRequest.Merge(m, src)
}
func (m *DeleteTenantRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTenantRequest.Size(m)
}
func (m *DeleteTenantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTenantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTenantRequest proto.InternalMessageInfo

func (m *DeleteTenantRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// List tenants for which the client has ACL visibility.
type ListTenantsRequest struct {
	// Required.
	//
	// Resource name of the project under which the tenant is created.
	//
	// The format is "projects/{project_id}", for example,
	// "projects/api-test-project".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional.
	//
	// The starting indicator from which to return results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional.
	//
	// The maximum number of tenants to be returned, at most 100.
	// Default is 100 if a non-positive number is provided.
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTenantsRequest) Reset()         { *m = ListTenantsRequest{} }
func (m *ListTenantsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTenantsRequest) ProtoMessage()    {}
func (*ListTenantsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df49c13fa187082, []int{4}
}

func (m *ListTenantsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTenantsRequest.Unmarshal(m, b)
}
func (m *ListTenantsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTenantsRequest.Marshal(b, m, deterministic)
}
func (m *ListTenantsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTenantsRequest.Merge(m, src)
}
func (m *ListTenantsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTenantsRequest.Size(m)
}
func (m *ListTenantsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTenantsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTenantsRequest proto.InternalMessageInfo

func (m *ListTenantsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListTenantsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListTenantsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Output only.
//
// The List tenants response object.
type ListTenantsResponse struct {
	// Tenants for the current client.
	Tenants []*Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	// A token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Additional information for the API invocation, such as the request
	// tracking id.
	Metadata             *ResponseMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListTenantsResponse) Reset()         { *m = ListTenantsResponse{} }
func (m *ListTenantsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTenantsResponse) ProtoMessage()    {}
func (*ListTenantsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df49c13fa187082, []int{5}
}

func (m *ListTenantsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTenantsResponse.Unmarshal(m, b)
}
func (m *ListTenantsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTenantsResponse.Marshal(b, m, deterministic)
}
func (m *ListTenantsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTenantsResponse.Merge(m, src)
}
func (m *ListTenantsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTenantsResponse.Size(m)
}
func (m *ListTenantsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTenantsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTenantsResponse proto.InternalMessageInfo

func (m *ListTenantsResponse) GetTenants() []*Tenant {
	if m != nil {
		return m.Tenants
	}
	return nil
}

func (m *ListTenantsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListTenantsResponse) GetMetadata() *ResponseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTenantRequest)(nil), "google.cloud.talent.v4beta1.CreateTenantRequest")
	proto.RegisterType((*GetTenantRequest)(nil), "google.cloud.talent.v4beta1.GetTenantRequest")
	proto.RegisterType((*UpdateTenantRequest)(nil), "google.cloud.talent.v4beta1.UpdateTenantRequest")
	proto.RegisterType((*DeleteTenantRequest)(nil), "google.cloud.talent.v4beta1.DeleteTenantRequest")
	proto.RegisterType((*ListTenantsRequest)(nil), "google.cloud.talent.v4beta1.ListTenantsRequest")
	proto.RegisterType((*ListTenantsResponse)(nil), "google.cloud.talent.v4beta1.ListTenantsResponse")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/tenant_service.proto", fileDescriptor_2df49c13fa187082)
}

var fileDescriptor_2df49c13fa187082 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x4f, 0x4f, 0x13, 0x41,
	0x18, 0xc6, 0x33, 0xa0, 0x95, 0xbe, 0x85, 0x68, 0xa6, 0x09, 0x69, 0x8a, 0xc6, 0x66, 0x35, 0xa4,
	0x54, 0xdd, 0xa9, 0xd5, 0x78, 0x90, 0x70, 0x10, 0xfc, 0x13, 0x13, 0x49, 0x48, 0xc1, 0x8b, 0x97,
	0x66, 0x68, 0x5f, 0xd6, 0x85, 0xee, 0xcc, 0xda, 0x99, 0x12, 0xc5, 0x90, 0x18, 0x0e, 0xde, 0x8c,
	0x07, 0x13, 0x6f, 0x7a, 0xf0, 0xcb, 0xf8, 0x01, 0xfc, 0x0a, 0x7e, 0x10, 0xb3, 0x33, 0xd3, 0x52,
	0x68, 0xd9, 0xae, 0xb7, 0xce, 0xec, 0x3c, 0xef, 0xfb, 0x9b, 0x67, 0xdf, 0x67, 0x0b, 0xf5, 0x40,
	0xca, 0xa0, 0x8b, 0xac, 0xdd, 0x95, 0xfd, 0x0e, 0xd3, 0xbc, 0x8b, 0x42, 0xb3, 0xc3, 0x87, 0xbb,
	0xa8, 0xf9, 0x7d, 0xa6, 0x51, 0x70, 0xa1, 0x5b, 0x0a, 0x7b, 0x87, 0x61, 0x1b, 0xfd, 0xb8, 0x27,
	0xb5, 0xa4, 0x4b, 0x56, 0xe1, 0x1b, 0x85, 0x6f, 0x15, 0xbe, 0x53, 0x94, 0xaf, 0xbb, 0x72, 0x3c,
	0x0e, 0x19, 0x17, 0x42, 0x6a, 0xae, 0x43, 0x29, 0x94, 0x95, 0x96, 0xab, 0x69, 0xcd, 0xda, 0x32,
	0x8a, 0xa4, 0xc8, 0x72, 0xd2, 0x62, 0xb9, 0x93, 0x0e, 0x87, 0x99, 0xd5, 0x6e, 0x7f, 0x8f, 0x61,
	0x14, 0xeb, 0x0f, 0xee, 0x61, 0xe5, 0xfc, 0xc3, 0xbd, 0x10, 0xbb, 0x9d, 0x56, 0xc4, 0xd5, 0x81,
	0x3d, 0xe1, 0xed, 0x43, 0x71, 0xa3, 0x87, 0x5c, 0xe3, 0x8e, 0x29, 0xda, 0xc4, 0x77, 0x7d, 0x54,
	0x9a, 0x2e, 0x42, 0x2e, 0xe6, 0x3d, 0x14, 0xba, 0x44, 0x2a, 0xa4, 0x9a, 0x6f, 0xba, 0x15, 0x5d,
	0x85, 0x9c, 0xed, 0x5e, 0x9a, 0xa9, 0x90, 0x6a, 0xa1, 0x71, 0xcb, 0x4f, 0x71, 0xc3, 0x77, 0x35,
	0x9d, 0xc4, 0x5b, 0x86, 0x6b, 0x2f, 0x50, 0x9f, 0x6d, 0x44, 0xe1, 0x92, 0xe0, 0x11, 0xba, 0x36,
	0xe6, 0xb7, 0xf7, 0x95, 0x40, 0xf1, 0x75, 0xdc, 0x19, 0x83, 0x3a, 0x6d, 0x4e, 0xfe, 0xbb, 0x39,
	0x5d, 0x85, 0x42, 0xdf, 0xd4, 0x34, 0xb7, 0x77, 0xf8, 0xe5, 0x41, 0x85, 0x81, 0x41, 0xfe, 0xf3,
	0xc4, 0xa0, 0x4d, 0xae, 0x0e, 0x9a, 0x60, 0x8f, 0x27, 0xbf, 0xbd, 0x15, 0x28, 0x3e, 0xc5, 0x2e,
	0x9e, 0x07, 0x9a, 0x04, 0xff, 0x16, 0xe8, 0xab, 0x50, 0xb9, 0x5b, 0xaa, 0x69, 0x7e, 0xde, 0x00,
	0x88, 0x79, 0x80, 0x2d, 0x2d, 0x0f, 0x50, 0x18, 0xa8, 0x7c, 0x33, 0x9f, 0xec, 0xec, 0x24, 0x1b,
	0x74, 0x09, 0xcc, 0xa2, 0xa5, 0xc2, 0x23, 0x2c, 0xcd, 0x56, 0x48, 0xf5, 0x72, 0x73, 0x2e, 0xd9,
	0xd8, 0x0e, 0x8f, 0xd0, 0xfb, 0x4d, 0xa0, 0x78, 0xa6, 0x95, 0x8a, 0xa5, 0x50, 0x48, 0xd7, 0xe0,
	0x8a, 0xbd, 0xb3, 0x2a, 0x91, 0xca, 0x6c, 0x56, 0x9f, 0x06, 0x1a, 0xba, 0x0c, 0x57, 0x05, 0xbe,
	0xd7, 0xad, 0x31, 0xae, 0x85, 0x64, 0x7b, 0x6b, 0xc8, 0xf6, 0x12, 0xe6, 0x22, 0xd4, 0xbc, 0xc3,
	0x35, 0x37, 0x68, 0x85, 0xc6, 0xbd, 0xd4, 0x3e, 0x03, 0xbe, 0x4d, 0x27, 0x6a, 0x0e, 0xe5, 0x8d,
	0x93, 0x1c, 0x2c, 0x58, 0x8c, 0x6d, 0x1b, 0x35, 0xfa, 0x9d, 0xc0, 0xfc, 0xe8, 0x5c, 0xd2, 0x7a,
	0x6a, 0xed, 0x09, 0x23, 0x5c, 0xce, 0x72, 0x6b, 0x8f, 0x9d, 0xfc, 0xf9, 0xfb, 0x6d, 0x66, 0xc5,
	0xbb, 0x3d, 0x0c, 0xd7, 0x47, 0xfb, 0x66, 0xd6, 0xe2, 0x9e, 0xdc, 0xc7, 0xb6, 0x56, 0xac, 0x76,
	0xec, 0x02, 0xa7, 0x1e, 0x93, 0x1a, 0xfd, 0x42, 0x20, 0x3f, 0x1c, 0x62, 0x9a, 0x7e, 0xe3, 0xf3,
	0xc3, 0x9e, 0x0d, 0xe9, 0xae, 0x41, 0x5a, 0xa6, 0x23, 0x48, 0xc9, 0x60, 0x8d, 0x00, 0x0d, 0x78,
	0x58, 0xed, 0x98, 0xfe, 0x20, 0x30, 0x3f, 0x9a, 0x95, 0x29, 0x46, 0x4d, 0x88, 0x55, 0x36, 0xaa,
	0x47, 0x86, 0xaa, 0xde, 0xb8, 0x73, 0x4a, 0xe5, 0x3e, 0x43, 0x17, 0xc3, 0x25, 0x7e, 0x7d, 0x26,
	0x30, 0x3f, 0x1a, 0x9d, 0x29, 0x7c, 0x13, 0x52, 0x56, 0x5e, 0x1c, 0x0b, 0xe9, 0xb3, 0xe4, 0x13,
	0x37, 0x30, 0xaa, 0x96, 0xcd, 0xa8, 0x9f, 0x04, 0x0a, 0x23, 0x69, 0xa1, 0x2c, 0x95, 0x63, 0x3c,
	0xc2, 0xe5, 0x7a, 0x76, 0x81, 0x1d, 0xf4, 0x49, 0x6f, 0xf2, 0xe2, 0xe1, 0x5a, 0xff, 0x44, 0xe0,
	0x66, 0x5b, 0x46, 0x69, 0x5d, 0xd6, 0xe9, 0x99, 0x94, 0x6c, 0x25, 0x7e, 0x6c, 0x91, 0x37, 0x4f,
	0x9c, 0x24, 0x90, 0x5d, 0x2e, 0x02, 0x5f, 0xf6, 0x02, 0x16, 0xa0, 0x30, 0x6e, 0x31, 0xfb, 0x88,
	0xc7, 0xa1, 0x9a, 0xf8, 0x5f, 0xb2, 0x6a, 0x97, 0xbf, 0x66, 0x66, 0x37, 0x76, 0xb6, 0x77, 0x73,
	0x46, 0xf3, 0xe0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x45, 0x0b, 0xcb, 0x15, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/ie310mu/ie310go/forks/google.golang.org/grpc#ClientConn.NewStream.
type TenantServiceClient interface {
	// Creates a new tenant entity.
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// Retrieves specified tenant.
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// Updates specified tenant.
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// Deletes specified tenant.
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all tenants associated with the project.
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
}

type tenantServiceClient struct {
	cc *grpc.ClientConn
}

func NewTenantServiceClient(cc *grpc.ClientConn) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.TenantService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.TenantService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.TenantService/UpdateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.TenantService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.TenantService/ListTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
type TenantServiceServer interface {
	// Creates a new tenant entity.
	CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error)
	// Retrieves specified tenant.
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	// Updates specified tenant.
	UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error)
	// Deletes specified tenant.
	DeleteTenant(context.Context, *DeleteTenantRequest) (*empty.Empty, error)
	// Lists all tenants associated with the project.
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
}

func RegisterTenantServiceServer(s *grpc.Server, srv TenantServiceServer) {
	s.RegisterService(&_TenantService_serviceDesc, srv)
}

func _TenantService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.TenantService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.TenantService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.TenantService/UpdateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.TenantService/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.TenantService/ListTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4beta1.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTenant",
			Handler:    _TenantService_CreateTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantService_GetTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _TenantService_UpdateTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantService_DeleteTenant_Handler,
		},
		{
			MethodName: "ListTenants",
			Handler:    _TenantService_ListTenants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4beta1/tenant_service.proto",
}
