// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/completion_service.proto

package talent

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	_ "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api/annotations"
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

// Enum to specify the scope of completion.
type CompleteQueryRequest_CompletionScope int32

const (
	// Default value.
	CompleteQueryRequest_COMPLETION_SCOPE_UNSPECIFIED CompleteQueryRequest_CompletionScope = 0
	// Suggestions are based only on the data provided by the client.
	CompleteQueryRequest_TENANT CompleteQueryRequest_CompletionScope = 1
	// Suggestions are based on all jobs data in the system that's visible to
	// the client
	CompleteQueryRequest_PUBLIC CompleteQueryRequest_CompletionScope = 2
)

var CompleteQueryRequest_CompletionScope_name = map[int32]string{
	0: "COMPLETION_SCOPE_UNSPECIFIED",
	1: "TENANT",
	2: "PUBLIC",
}

var CompleteQueryRequest_CompletionScope_value = map[string]int32{
	"COMPLETION_SCOPE_UNSPECIFIED": 0,
	"TENANT":                       1,
	"PUBLIC":                       2,
}

func (x CompleteQueryRequest_CompletionScope) String() string {
	return proto.EnumName(CompleteQueryRequest_CompletionScope_name, int32(x))
}

func (CompleteQueryRequest_CompletionScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6e2d15ed940916e, []int{0, 0}
}

// Enum to specify auto-completion topics.
type CompleteQueryRequest_CompletionType int32

const (
	// Default value.
	CompleteQueryRequest_COMPLETION_TYPE_UNSPECIFIED CompleteQueryRequest_CompletionType = 0
	// Only suggest job titles.
	CompleteQueryRequest_JOB_TITLE CompleteQueryRequest_CompletionType = 1
	// Only suggest company names.
	CompleteQueryRequest_COMPANY_NAME CompleteQueryRequest_CompletionType = 2
	// Suggest both job titles and company names.
	CompleteQueryRequest_COMBINED CompleteQueryRequest_CompletionType = 3
)

var CompleteQueryRequest_CompletionType_name = map[int32]string{
	0: "COMPLETION_TYPE_UNSPECIFIED",
	1: "JOB_TITLE",
	2: "COMPANY_NAME",
	3: "COMBINED",
}

var CompleteQueryRequest_CompletionType_value = map[string]int32{
	"COMPLETION_TYPE_UNSPECIFIED": 0,
	"JOB_TITLE":                   1,
	"COMPANY_NAME":                2,
	"COMBINED":                    3,
}

func (x CompleteQueryRequest_CompletionType) String() string {
	return proto.EnumName(CompleteQueryRequest_CompletionType_name, int32(x))
}

func (CompleteQueryRequest_CompletionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6e2d15ed940916e, []int{0, 1}
}

// Input only.
//
// Auto-complete parameters.
type CompleteQueryRequest struct {
	// Required.
	//
	// Resource name of tenant the completion is performed within.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/api-test-project/tenant/foo".
	//
	// Tenant id is optional and the default tenant is used if unspecified, for
	// example, "projects/api-test-project".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required.
	//
	// The query used to generate suggestions.
	//
	// The maximum number of allowed characters is 255.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// Optional.
	//
	// The list of languages of the query. This is
	// the BCP-47 language code, such as "en-US" or "sr-Latn".
	// For more information, see
	// [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47).
	//
	// For
	// [CompletionType.JOB_TITLE][google.cloud.talent.v4beta1.CompleteQueryRequest.CompletionType.JOB_TITLE]
	// type, only open jobs with the same
	// [language_codes][google.cloud.talent.v4beta1.CompleteQueryRequest.language_codes]
	// are returned.
	//
	// For
	// [CompletionType.COMPANY_NAME][google.cloud.talent.v4beta1.CompleteQueryRequest.CompletionType.COMPANY_NAME]
	// type, only companies having open jobs with the same
	// [language_codes][google.cloud.talent.v4beta1.CompleteQueryRequest.language_codes]
	// are returned.
	//
	// For
	// [CompletionType.COMBINED][google.cloud.talent.v4beta1.CompleteQueryRequest.CompletionType.COMBINED]
	// type, only open jobs with the same
	// [language_codes][google.cloud.talent.v4beta1.CompleteQueryRequest.language_codes]
	// or companies having open jobs with the same
	// [language_codes][google.cloud.talent.v4beta1.CompleteQueryRequest.language_codes]
	// are returned.
	//
	// The maximum number of allowed characters is 255.
	LanguageCodes []string `protobuf:"bytes,3,rep,name=language_codes,json=languageCodes,proto3" json:"language_codes,omitempty"`
	// Required.
	//
	// Completion result count.
	//
	// The maximum allowed page size is 10.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional.
	//
	// If provided, restricts completion to specified company.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/api-test-project/tenants/foo/companies/bar".
	//
	// Tenant id is optional and the default tenant is used if unspecified, for
	// example, "projects/api-test-project/companies/bar".
	Company string `protobuf:"bytes,5,opt,name=company,proto3" json:"company,omitempty"`
	// Optional.
	//
	// The scope of the completion. The defaults is
	// [CompletionScope.PUBLIC][google.cloud.talent.v4beta1.CompleteQueryRequest.CompletionScope.PUBLIC].
	Scope CompleteQueryRequest_CompletionScope `protobuf:"varint,6,opt,name=scope,proto3,enum=google.cloud.talent.v4beta1.CompleteQueryRequest_CompletionScope" json:"scope,omitempty"`
	// Optional.
	//
	// The completion topic. The default is
	// [CompletionType.COMBINED][google.cloud.talent.v4beta1.CompleteQueryRequest.CompletionType.COMBINED].
	Type                 CompleteQueryRequest_CompletionType `protobuf:"varint,7,opt,name=type,proto3,enum=google.cloud.talent.v4beta1.CompleteQueryRequest_CompletionType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *CompleteQueryRequest) Reset()         { *m = CompleteQueryRequest{} }
func (m *CompleteQueryRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteQueryRequest) ProtoMessage()    {}
func (*CompleteQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6e2d15ed940916e, []int{0}
}

func (m *CompleteQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteQueryRequest.Unmarshal(m, b)
}
func (m *CompleteQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteQueryRequest.Marshal(b, m, deterministic)
}
func (m *CompleteQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteQueryRequest.Merge(m, src)
}
func (m *CompleteQueryRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteQueryRequest.Size(m)
}
func (m *CompleteQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteQueryRequest proto.InternalMessageInfo

func (m *CompleteQueryRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CompleteQueryRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *CompleteQueryRequest) GetLanguageCodes() []string {
	if m != nil {
		return m.LanguageCodes
	}
	return nil
}

func (m *CompleteQueryRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *CompleteQueryRequest) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *CompleteQueryRequest) GetScope() CompleteQueryRequest_CompletionScope {
	if m != nil {
		return m.Scope
	}
	return CompleteQueryRequest_COMPLETION_SCOPE_UNSPECIFIED
}

func (m *CompleteQueryRequest) GetType() CompleteQueryRequest_CompletionType {
	if m != nil {
		return m.Type
	}
	return CompleteQueryRequest_COMPLETION_TYPE_UNSPECIFIED
}

// Output only.
//
// Response of auto-complete query.
type CompleteQueryResponse struct {
	// Results of the matching job/company candidates.
	CompletionResults []*CompleteQueryResponse_CompletionResult `protobuf:"bytes,1,rep,name=completion_results,json=completionResults,proto3" json:"completion_results,omitempty"`
	// Additional information for the API invocation, such as the request
	// tracking id.
	Metadata             *ResponseMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CompleteQueryResponse) Reset()         { *m = CompleteQueryResponse{} }
func (m *CompleteQueryResponse) String() string { return proto.CompactTextString(m) }
func (*CompleteQueryResponse) ProtoMessage()    {}
func (*CompleteQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6e2d15ed940916e, []int{1}
}

func (m *CompleteQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteQueryResponse.Unmarshal(m, b)
}
func (m *CompleteQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteQueryResponse.Marshal(b, m, deterministic)
}
func (m *CompleteQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteQueryResponse.Merge(m, src)
}
func (m *CompleteQueryResponse) XXX_Size() int {
	return xxx_messageInfo_CompleteQueryResponse.Size(m)
}
func (m *CompleteQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteQueryResponse proto.InternalMessageInfo

func (m *CompleteQueryResponse) GetCompletionResults() []*CompleteQueryResponse_CompletionResult {
	if m != nil {
		return m.CompletionResults
	}
	return nil
}

func (m *CompleteQueryResponse) GetMetadata() *ResponseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Output only.
//
// Resource that represents completion results.
type CompleteQueryResponse_CompletionResult struct {
	// The suggestion for the query.
	Suggestion string `protobuf:"bytes,1,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
	// The completion topic.
	Type CompleteQueryRequest_CompletionType `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.talent.v4beta1.CompleteQueryRequest_CompletionType" json:"type,omitempty"`
	// The URI of the company image for
	// [COMPANY_NAME][google.cloud.talent.v4beta1.CompleteQueryRequest.CompletionType.COMPANY_NAME].
	ImageUri             string   `protobuf:"bytes,3,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteQueryResponse_CompletionResult) Reset() {
	*m = CompleteQueryResponse_CompletionResult{}
}
func (m *CompleteQueryResponse_CompletionResult) String() string { return proto.CompactTextString(m) }
func (*CompleteQueryResponse_CompletionResult) ProtoMessage()    {}
func (*CompleteQueryResponse_CompletionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6e2d15ed940916e, []int{1, 0}
}

func (m *CompleteQueryResponse_CompletionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteQueryResponse_CompletionResult.Unmarshal(m, b)
}
func (m *CompleteQueryResponse_CompletionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteQueryResponse_CompletionResult.Marshal(b, m, deterministic)
}
func (m *CompleteQueryResponse_CompletionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteQueryResponse_CompletionResult.Merge(m, src)
}
func (m *CompleteQueryResponse_CompletionResult) XXX_Size() int {
	return xxx_messageInfo_CompleteQueryResponse_CompletionResult.Size(m)
}
func (m *CompleteQueryResponse_CompletionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteQueryResponse_CompletionResult.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteQueryResponse_CompletionResult proto.InternalMessageInfo

func (m *CompleteQueryResponse_CompletionResult) GetSuggestion() string {
	if m != nil {
		return m.Suggestion
	}
	return ""
}

func (m *CompleteQueryResponse_CompletionResult) GetType() CompleteQueryRequest_CompletionType {
	if m != nil {
		return m.Type
	}
	return CompleteQueryRequest_COMPLETION_TYPE_UNSPECIFIED
}

func (m *CompleteQueryResponse_CompletionResult) GetImageUri() string {
	if m != nil {
		return m.ImageUri
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.talent.v4beta1.CompleteQueryRequest_CompletionScope", CompleteQueryRequest_CompletionScope_name, CompleteQueryRequest_CompletionScope_value)
	proto.RegisterEnum("google.cloud.talent.v4beta1.CompleteQueryRequest_CompletionType", CompleteQueryRequest_CompletionType_name, CompleteQueryRequest_CompletionType_value)
	proto.RegisterType((*CompleteQueryRequest)(nil), "google.cloud.talent.v4beta1.CompleteQueryRequest")
	proto.RegisterType((*CompleteQueryResponse)(nil), "google.cloud.talent.v4beta1.CompleteQueryResponse")
	proto.RegisterType((*CompleteQueryResponse_CompletionResult)(nil), "google.cloud.talent.v4beta1.CompleteQueryResponse.CompletionResult")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/completion_service.proto", fileDescriptor_a6e2d15ed940916e)
}

var fileDescriptor_a6e2d15ed940916e = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x6f, 0x6b, 0xd3, 0x5e,
	0x14, 0xc7, 0x7f, 0x49, 0xd7, 0xae, 0x3d, 0xfb, 0xf3, 0x8b, 0x97, 0x39, 0x42, 0x37, 0x5c, 0x29,
	0x0c, 0x8b, 0x60, 0xc2, 0xea, 0x1e, 0x29, 0x82, 0x6d, 0x16, 0x21, 0xba, 0xa6, 0x35, 0xcd, 0x90,
	0xed, 0x49, 0x77, 0x97, 0x5d, 0x42, 0xa4, 0xcd, 0xcd, 0x72, 0x6f, 0x06, 0x9d, 0x08, 0x22, 0xbe,
	0x03, 0xdf, 0x81, 0xe0, 0x2b, 0xf2, 0x89, 0x2f, 0x40, 0xf0, 0x5d, 0x88, 0xdc, 0x24, 0xdd, 0xba,
	0x3a, 0x2a, 0x43, 0x9f, 0xe5, 0x7c, 0x73, 0xbe, 0x9f, 0x73, 0x73, 0x72, 0xee, 0x81, 0x5d, 0x9f,
	0x52, 0x7f, 0x48, 0x74, 0x6f, 0x48, 0x93, 0x53, 0x9d, 0xe3, 0x21, 0x09, 0xb9, 0x7e, 0xbe, 0x7b,
	0x42, 0x38, 0xde, 0xd1, 0x3d, 0x3a, 0x8a, 0x86, 0x84, 0x07, 0x34, 0x1c, 0x30, 0x12, 0x9f, 0x07,
	0x1e, 0xd1, 0xa2, 0x98, 0x72, 0x8a, 0x36, 0x32, 0x97, 0x96, 0xba, 0xb4, 0xcc, 0xa5, 0xe5, 0xae,
	0xea, 0x66, 0x8e, 0xc4, 0x51, 0xa0, 0xe3, 0x30, 0xa4, 0x1c, 0x0b, 0x02, 0xcb, 0xac, 0xd5, 0xc6,
	0x1f, 0x0a, 0x8e, 0x68, 0x98, 0x65, 0xd6, 0xdf, 0x2f, 0xc0, 0x9a, 0x91, 0x9d, 0x80, 0xbc, 0x4a,
	0x48, 0x3c, 0x76, 0xc8, 0x59, 0x42, 0x18, 0x47, 0xeb, 0x50, 0x8a, 0x70, 0x4c, 0x42, 0xae, 0x4a,
	0x35, 0xa9, 0x51, 0x71, 0xf2, 0x08, 0xad, 0x41, 0xf1, 0x4c, 0xe4, 0xa9, 0x72, 0x2a, 0x67, 0x01,
	0xda, 0x86, 0xd5, 0x21, 0x0e, 0xfd, 0x04, 0xfb, 0x64, 0xe0, 0xd1, 0x53, 0xc2, 0xd4, 0x42, 0xad,
	0xd0, 0xa8, 0x38, 0x2b, 0x13, 0xd5, 0x10, 0x22, 0xda, 0x80, 0x4a, 0x24, 0x52, 0x58, 0x70, 0x41,
	0xd4, 0x85, 0x9a, 0xd4, 0x28, 0x3a, 0x65, 0x21, 0xf4, 0x83, 0x0b, 0x82, 0x54, 0x58, 0x14, 0xbd,
	0xc0, 0xe1, 0x58, 0x2d, 0xa6, 0xec, 0x49, 0x88, 0x5e, 0x43, 0x91, 0x79, 0x34, 0x22, 0x6a, 0xa9,
	0x26, 0x35, 0x56, 0x9b, 0x2d, 0x6d, 0x4e, 0x67, 0xb4, 0x9b, 0xbe, 0x66, 0x22, 0x06, 0x34, 0xec,
	0x0b, 0x90, 0x93, 0xf1, 0x90, 0x0b, 0x0b, 0x7c, 0x1c, 0x11, 0x75, 0x31, 0xe5, 0x3e, 0xfb, 0x1b,
	0xae, 0x3b, 0x8e, 0x88, 0x93, 0xd2, 0xea, 0x2f, 0xe1, 0xff, 0x99, 0x7a, 0xa8, 0x06, 0x9b, 0x46,
	0xb7, 0xd3, 0xdb, 0x37, 0x5d, 0xab, 0x6b, 0x0f, 0xfa, 0x46, 0xb7, 0x67, 0x0e, 0x0e, 0xec, 0x7e,
	0xcf, 0x34, 0xac, 0xe7, 0x96, 0xb9, 0xa7, 0xfc, 0x87, 0x00, 0x4a, 0xae, 0x69, 0xb7, 0x6c, 0x57,
	0x91, 0xc4, 0x73, 0xef, 0xa0, 0xbd, 0x6f, 0x19, 0x8a, 0x5c, 0x3f, 0x86, 0xd5, 0xeb, 0x45, 0xd0,
	0x16, 0x6c, 0x4c, 0xb1, 0xdc, 0xc3, 0xdf, 0x50, 0x2b, 0x50, 0x79, 0xd1, 0x6d, 0x0f, 0x5c, 0xcb,
	0xdd, 0x37, 0x15, 0x09, 0x29, 0xb0, 0x2c, 0xf2, 0x5b, 0xf6, 0xe1, 0xc0, 0x6e, 0x75, 0x4c, 0x45,
	0x46, 0xcb, 0x50, 0x36, 0xba, 0x9d, 0xb6, 0x65, 0x9b, 0x7b, 0x4a, 0xa1, 0xfe, 0x53, 0x86, 0xbb,
	0x33, 0x1f, 0xc7, 0x22, 0x1a, 0x32, 0x82, 0x62, 0x40, 0x53, 0xd3, 0x19, 0x13, 0x96, 0x0c, 0x39,
	0x53, 0xa5, 0x5a, 0xa1, 0xb1, 0xd4, 0x34, 0x6e, 0xd3, 0xac, 0x8c, 0x37, 0xd5, 0x2d, 0x27, 0x65,
	0x39, 0x77, 0xbc, 0x19, 0x85, 0x21, 0x0b, 0xca, 0x23, 0xc2, 0xf1, 0x29, 0xe6, 0x38, 0x1d, 0xb1,
	0xa5, 0xe6, 0xc3, 0xb9, 0x95, 0x26, 0xf0, 0x4e, 0x6e, 0x72, 0x2e, 0xed, 0xd5, 0x2f, 0x12, 0x28,
	0xb3, 0x25, 0xd1, 0x3d, 0x00, 0x96, 0xf8, 0x3e, 0x61, 0x42, 0xcb, 0x67, 0x7b, 0x4a, 0xb9, 0x1c,
	0x09, 0xf9, 0x5f, 0x8e, 0x84, 0x18, 0xfc, 0x60, 0x24, 0x26, 0x3f, 0x89, 0x03, 0xb5, 0x90, 0x16,
	0x2d, 0xa7, 0xc2, 0x41, 0x1c, 0x34, 0x7f, 0x48, 0x00, 0x57, 0x2e, 0xf4, 0x4d, 0x82, 0x95, 0x6b,
	0x64, 0xb4, 0x73, 0xeb, 0x53, 0x54, 0x9b, 0xb7, 0xff, 0x3d, 0xf5, 0xe3, 0x0f, 0x5f, 0xbf, 0x7f,
	0x92, 0x8f, 0x90, 0x7e, 0xb9, 0x2a, 0xde, 0x66, 0x97, 0xfe, 0x69, 0x14, 0xd3, 0x37, 0xc4, 0xe3,
	0x4c, 0x7f, 0xa0, 0x73, 0x12, 0xe2, 0x50, 0x3c, 0xbd, 0x7b, 0x9c, 0xff, 0x3c, 0x72, 0x74, 0x1f,
	0x6d, 0xcf, 0xb1, 0x5c, 0x25, 0xb6, 0x3f, 0x4a, 0xb0, 0xe5, 0xd1, 0xd1, 0xbc, 0xb3, 0xb5, 0xd7,
	0xa7, 0xee, 0x4e, 0xb6, 0x0f, 0x7b, 0x62, 0x53, 0xf5, 0xa4, 0xa3, 0x56, 0x6e, 0xf3, 0xa9, 0xd8,
	0x2a, 0x1a, 0x8d, 0x7d, 0xdd, 0x27, 0x61, 0xba, 0xc7, 0xf4, 0xec, 0x15, 0x8e, 0x02, 0x76, 0xe3,
	0xd2, 0x7b, 0x92, 0x85, 0x9f, 0xe5, 0x82, 0xe1, 0xf6, 0x4f, 0x4a, 0xa9, 0xe7, 0xd1, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0xb9, 0x64, 0x84, 0x98, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompletionClient is the client API for Completion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/ie310mu/ie310go/forks/google.golang.org/grpc#ClientConn.NewStream.
type CompletionClient interface {
	// Completes the specified prefix with keyword suggestions.
	// Intended for use by a job search auto-complete search box.
	CompleteQuery(ctx context.Context, in *CompleteQueryRequest, opts ...grpc.CallOption) (*CompleteQueryResponse, error)
}

type completionClient struct {
	cc *grpc.ClientConn
}

func NewCompletionClient(cc *grpc.ClientConn) CompletionClient {
	return &completionClient{cc}
}

func (c *completionClient) CompleteQuery(ctx context.Context, in *CompleteQueryRequest, opts ...grpc.CallOption) (*CompleteQueryResponse, error) {
	out := new(CompleteQueryResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.Completion/CompleteQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompletionServer is the server API for Completion service.
type CompletionServer interface {
	// Completes the specified prefix with keyword suggestions.
	// Intended for use by a job search auto-complete search box.
	CompleteQuery(context.Context, *CompleteQueryRequest) (*CompleteQueryResponse, error)
}

func RegisterCompletionServer(s *grpc.Server, srv CompletionServer) {
	s.RegisterService(&_Completion_serviceDesc, srv)
}

func _Completion_CompleteQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServer).CompleteQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.Completion/CompleteQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServer).CompleteQuery(ctx, req.(*CompleteQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Completion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4beta1.Completion",
	HandlerType: (*CompletionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompleteQuery",
			Handler:    _Completion_CompleteQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4beta1/completion_service.proto",
}
