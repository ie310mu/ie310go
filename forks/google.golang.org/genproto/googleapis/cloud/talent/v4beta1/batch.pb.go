// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/batch.proto

package talent

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	timestamp "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/api/annotations"
	status "github.com/ie310mu/ie310go/forks/google.golang.org/genproto/googleapis/rpc/status"
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

type BatchOperationMetadata_State int32

const (
	// Default value.
	BatchOperationMetadata_STATE_UNSPECIFIED BatchOperationMetadata_State = 0
	// The batch operation is being prepared for processing.
	BatchOperationMetadata_INITIALIZING BatchOperationMetadata_State = 1
	// The batch operation is actively being processed.
	BatchOperationMetadata_PROCESSING BatchOperationMetadata_State = 2
	// The batch operation is processed, and at least one item has been
	// successfully processed.
	BatchOperationMetadata_SUCCEEDED BatchOperationMetadata_State = 3
	// The batch operation is done and no item has been successfully processed.
	BatchOperationMetadata_FAILED BatchOperationMetadata_State = 4
	// The batch operation is in the process of cancelling after
	// [google.longrunning.Operation.CancelOperation] is called.
	BatchOperationMetadata_CANCELLING BatchOperationMetadata_State = 5
	// The batch operation is done after
	// [google.longrunning.Operation.CancelOperation] is called. Any items
	// processed before cancelling are returned in the response.
	BatchOperationMetadata_CANCELLED BatchOperationMetadata_State = 6
)

var BatchOperationMetadata_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "INITIALIZING",
	2: "PROCESSING",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "CANCELLING",
	6: "CANCELLED",
}

var BatchOperationMetadata_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"INITIALIZING":      1,
	"PROCESSING":        2,
	"SUCCEEDED":         3,
	"FAILED":            4,
	"CANCELLING":        5,
	"CANCELLED":         6,
}

func (x BatchOperationMetadata_State) String() string {
	return proto.EnumName(BatchOperationMetadata_State_name, int32(x))
}

func (BatchOperationMetadata_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b52f05e5c2b02efe, []int{0, 0}
}

// Output only.
//
// Metadata used for long running operations returned by CTS batch APIs.
// It's used to replace
// [google.longrunning.Operation.metadata][google.longrunning.Operation.metadata].
type BatchOperationMetadata struct {
	// The state of a long running operation.
	State BatchOperationMetadata_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.talent.v4beta1.BatchOperationMetadata_State" json:"state,omitempty"`
	// More detailed information about operation state.
	StateDescription string `protobuf:"bytes,2,opt,name=state_description,json=stateDescription,proto3" json:"state_description,omitempty"`
	// Count of successful item(s) inside an operation.
	SuccessCount int32 `protobuf:"varint,3,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	// Count of failed item(s) inside an operation.
	FailureCount int32 `protobuf:"varint,4,opt,name=failure_count,json=failureCount,proto3" json:"failure_count,omitempty"`
	// Count of total item(s) inside an operation.
	TotalCount int32 `protobuf:"varint,5,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// The time when the batch operation is created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time when the batch operation status is updated. The metadata and the
	// [update_time][google.cloud.talent.v4beta1.BatchOperationMetadata.update_time]
	// is refreshed every minute otherwise cached data is returned.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The time when the batch operation is finished and
	// [google.longrunning.Operation.done][google.longrunning.Operation.done] is
	// set to `true`.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BatchOperationMetadata) Reset()         { *m = BatchOperationMetadata{} }
func (m *BatchOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*BatchOperationMetadata) ProtoMessage()    {}
func (*BatchOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52f05e5c2b02efe, []int{0}
}

func (m *BatchOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchOperationMetadata.Unmarshal(m, b)
}
func (m *BatchOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchOperationMetadata.Marshal(b, m, deterministic)
}
func (m *BatchOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchOperationMetadata.Merge(m, src)
}
func (m *BatchOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_BatchOperationMetadata.Size(m)
}
func (m *BatchOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BatchOperationMetadata proto.InternalMessageInfo

func (m *BatchOperationMetadata) GetState() BatchOperationMetadata_State {
	if m != nil {
		return m.State
	}
	return BatchOperationMetadata_STATE_UNSPECIFIED
}

func (m *BatchOperationMetadata) GetStateDescription() string {
	if m != nil {
		return m.StateDescription
	}
	return ""
}

func (m *BatchOperationMetadata) GetSuccessCount() int32 {
	if m != nil {
		return m.SuccessCount
	}
	return 0
}

func (m *BatchOperationMetadata) GetFailureCount() int32 {
	if m != nil {
		return m.FailureCount
	}
	return 0
}

func (m *BatchOperationMetadata) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *BatchOperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *BatchOperationMetadata) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *BatchOperationMetadata) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// Output only.
//
// The result of [BatchCreateJobs][] or [BatchUpdateJobs][] APIs. It's used to
// replace
// [google.longrunning.Operation.response][google.longrunning.Operation.response]
// in case of success.
type JobOperationResult struct {
	// List of job mutation results from a batch mutate operation. It can change
	// until operation status is FINISHED, FAILED or CANCELLED.
	JobResults           []*JobOperationResult_JobResult `protobuf:"bytes,1,rep,name=job_results,json=jobResults,proto3" json:"job_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *JobOperationResult) Reset()         { *m = JobOperationResult{} }
func (m *JobOperationResult) String() string { return proto.CompactTextString(m) }
func (*JobOperationResult) ProtoMessage()    {}
func (*JobOperationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52f05e5c2b02efe, []int{1}
}

func (m *JobOperationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobOperationResult.Unmarshal(m, b)
}
func (m *JobOperationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobOperationResult.Marshal(b, m, deterministic)
}
func (m *JobOperationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobOperationResult.Merge(m, src)
}
func (m *JobOperationResult) XXX_Size() int {
	return xxx_messageInfo_JobOperationResult.Size(m)
}
func (m *JobOperationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_JobOperationResult.DiscardUnknown(m)
}

var xxx_messageInfo_JobOperationResult proto.InternalMessageInfo

func (m *JobOperationResult) GetJobResults() []*JobOperationResult_JobResult {
	if m != nil {
		return m.JobResults
	}
	return nil
}

// Mutation result of a job.
type JobOperationResult_JobResult struct {
	// Here [Job][google.cloud.talent.v4beta1.Job] only contains basic
	// information including [name][google.cloud.talent.v4beta1.Job.name],
	// [company][google.cloud.talent.v4beta1.Job.company],
	// [language_code][google.cloud.talent.v4beta1.Job.language_code] and
	// [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id], use
	// getJob method to retrieve detailed information of the created/updated
	// job.
	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	// The status of the job processed. This field is populated if the
	// processing of the
	// [job][google.cloud.talent.v4beta1.JobOperationResult.JobResult.job]
	// fails.
	Status               *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JobOperationResult_JobResult) Reset()         { *m = JobOperationResult_JobResult{} }
func (m *JobOperationResult_JobResult) String() string { return proto.CompactTextString(m) }
func (*JobOperationResult_JobResult) ProtoMessage()    {}
func (*JobOperationResult_JobResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52f05e5c2b02efe, []int{1, 0}
}

func (m *JobOperationResult_JobResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobOperationResult_JobResult.Unmarshal(m, b)
}
func (m *JobOperationResult_JobResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobOperationResult_JobResult.Marshal(b, m, deterministic)
}
func (m *JobOperationResult_JobResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobOperationResult_JobResult.Merge(m, src)
}
func (m *JobOperationResult_JobResult) XXX_Size() int {
	return xxx_messageInfo_JobOperationResult_JobResult.Size(m)
}
func (m *JobOperationResult_JobResult) XXX_DiscardUnknown() {
	xxx_messageInfo_JobOperationResult_JobResult.DiscardUnknown(m)
}

var xxx_messageInfo_JobOperationResult_JobResult proto.InternalMessageInfo

func (m *JobOperationResult_JobResult) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *JobOperationResult_JobResult) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.talent.v4beta1.BatchOperationMetadata_State", BatchOperationMetadata_State_name, BatchOperationMetadata_State_value)
	proto.RegisterType((*BatchOperationMetadata)(nil), "google.cloud.talent.v4beta1.BatchOperationMetadata")
	proto.RegisterType((*JobOperationResult)(nil), "google.cloud.talent.v4beta1.JobOperationResult")
	proto.RegisterType((*JobOperationResult_JobResult)(nil), "google.cloud.talent.v4beta1.JobOperationResult.JobResult")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/batch.proto", fileDescriptor_b52f05e5c2b02efe)
}

var fileDescriptor_b52f05e5c2b02efe = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0xd3, 0xa4, 0xed, 0xb8, 0xad, 0xdc, 0x95, 0x80, 0x28, 0x20, 0xd5, 0x2a, 0x42,
	0x44, 0x20, 0xd9, 0x22, 0xc0, 0x01, 0xf5, 0x94, 0xda, 0x2e, 0x32, 0x0a, 0x69, 0x64, 0xa7, 0x97,
	0x5e, 0xa2, 0xb5, 0xbd, 0x35, 0x0e, 0x8e, 0xd7, 0xf2, 0xae, 0x11, 0xe2, 0x71, 0x78, 0x00, 0xde,
	0x88, 0x77, 0x41, 0xfb, 0x27, 0xe1, 0x40, 0x95, 0xdc, 0x76, 0xbe, 0xfd, 0x7d, 0x33, 0xde, 0x99,
	0x31, 0xbc, 0xca, 0x29, 0xcd, 0x4b, 0xe2, 0xa6, 0x25, 0x6d, 0x33, 0x97, 0xe3, 0x92, 0x54, 0xdc,
	0xfd, 0xfe, 0x3e, 0x21, 0x1c, 0xbf, 0x75, 0x13, 0xcc, 0xd3, 0xaf, 0x4e, 0xdd, 0x50, 0x4e, 0xd1,
	0x33, 0x05, 0x3a, 0x12, 0x74, 0x14, 0xe8, 0x68, 0x70, 0xf0, 0x5c, 0x67, 0xc1, 0x75, 0xe1, 0xe2,
	0xaa, 0xa2, 0x1c, 0xf3, 0x82, 0x56, 0x4c, 0x59, 0x07, 0x2f, 0xb7, 0xd5, 0x58, 0xd2, 0x44, 0x63,
	0xe7, 0x1a, 0x93, 0x51, 0xd2, 0xde, 0xbb, 0xbc, 0x58, 0x11, 0xc6, 0xf1, 0xaa, 0xd6, 0xc0, 0x53,
	0x0d, 0x34, 0x75, 0xea, 0x32, 0x8e, 0x79, 0xab, 0x0b, 0x5c, 0xfc, 0xde, 0x87, 0x27, 0x57, 0xe2,
	0x5b, 0x6f, 0x6a, 0xd2, 0xc8, 0xd2, 0x5f, 0x08, 0xc7, 0x19, 0xe6, 0x18, 0xdd, 0x40, 0x57, 0xa0,
	0xa4, 0x6f, 0xd8, 0xc6, 0xf0, 0x74, 0xf4, 0xd1, 0xd9, 0xf2, 0x0c, 0xe7, 0xe1, 0x1c, 0x4e, 0x2c,
	0x12, 0x44, 0x2a, 0x0f, 0x7a, 0x03, 0x67, 0xf2, 0xb0, 0xc8, 0x08, 0x4b, 0x9b, 0xa2, 0x16, 0x64,
	0x7f, 0xcf, 0x36, 0x86, 0x47, 0x91, 0x25, 0x2f, 0xfc, 0x7f, 0x3a, 0x7a, 0x01, 0x27, 0xac, 0x4d,
	0x53, 0xc2, 0xd8, 0x22, 0xa5, 0x6d, 0xc5, 0xfb, 0x1d, 0xdb, 0x18, 0x76, 0xa3, 0x63, 0x2d, 0x7a,
	0x42, 0x13, 0xd0, 0x3d, 0x2e, 0xca, 0xb6, 0x21, 0x1a, 0xda, 0x57, 0x90, 0x16, 0x15, 0x74, 0x0e,
	0x26, 0xa7, 0x1c, 0x97, 0x1a, 0xe9, 0x4a, 0x04, 0xa4, 0xa4, 0x80, 0x4b, 0x30, 0xd3, 0x86, 0x88,
	0x0f, 0x13, 0x6d, 0xeb, 0xf7, 0x6c, 0x63, 0x68, 0x8e, 0x06, 0xeb, 0xe7, 0xae, 0x7b, 0xea, 0xcc,
	0xd7, 0x3d, 0x8d, 0x40, 0xe1, 0x42, 0x10, 0xe6, 0xb6, 0xce, 0x36, 0xe6, 0x83, 0xdd, 0x66, 0x85,
	0x4b, 0xf3, 0x07, 0x38, 0x24, 0x55, 0xa6, 0x9c, 0x87, 0x3b, 0x9d, 0x07, 0xa4, 0xca, 0x44, 0x74,
	0xf1, 0x13, 0xba, 0xb2, 0xb1, 0xe8, 0x31, 0x9c, 0xc5, 0xf3, 0xf1, 0x3c, 0x58, 0xdc, 0x4e, 0xe3,
	0x59, 0xe0, 0x85, 0xd7, 0x61, 0xe0, 0x5b, 0x8f, 0x90, 0x05, 0xc7, 0xe1, 0x34, 0x9c, 0x87, 0xe3,
	0x49, 0x78, 0x17, 0x4e, 0x3f, 0x59, 0x06, 0x3a, 0x05, 0x98, 0x45, 0x37, 0x5e, 0x10, 0xc7, 0x22,
	0xde, 0x43, 0x27, 0x70, 0x14, 0xdf, 0x7a, 0x5e, 0x10, 0xf8, 0x81, 0x6f, 0x75, 0x10, 0x40, 0xef,
	0x7a, 0x1c, 0x4e, 0x02, 0xdf, 0xda, 0x17, 0xa8, 0x37, 0x9e, 0x7a, 0xc1, 0x64, 0x22, 0xd0, 0xae,
	0x40, 0x75, 0x1c, 0xf8, 0x56, 0xef, 0xe2, 0x8f, 0x01, 0xe8, 0x33, 0x4d, 0x36, 0xa3, 0x8e, 0x08,
	0x6b, 0x4b, 0x8e, 0xee, 0xc0, 0x5c, 0xd2, 0x64, 0xd1, 0xc8, 0x88, 0xf5, 0x0d, 0xbb, 0x33, 0x34,
	0x77, 0xac, 0xcc, 0xff, 0x59, 0x84, 0xa4, 0x4e, 0x11, 0x2c, 0xd7, 0x47, 0x36, 0xf8, 0x06, 0x47,
	0x9b, 0x0b, 0x34, 0x82, 0xce, 0x92, 0x26, 0x72, 0x27, 0xcd, 0x91, 0xbd, 0xab, 0x40, 0x24, 0x60,
	0xf4, 0x1a, 0x7a, 0x6a, 0xe9, 0xe5, 0xb6, 0x99, 0x23, 0xb4, 0xb6, 0x35, 0x75, 0x2a, 0x57, 0xb4,
	0x65, 0x91, 0x26, 0xae, 0x7e, 0xc0, 0x79, 0x4a, 0x57, 0xdb, 0xf2, 0x5e, 0x81, 0x5c, 0xf6, 0x99,
	0x18, 0xd0, 0xcc, 0xb8, 0x1b, 0x6b, 0x34, 0xa7, 0x25, 0xae, 0x72, 0x87, 0x36, 0xb9, 0x9b, 0x93,
	0x4a, 0x8e, 0xcf, 0x55, 0x57, 0xb8, 0x2e, 0xd8, 0x83, 0x7f, 0xf0, 0xa5, 0x0a, 0x7f, 0xed, 0x75,
	0xbc, 0x79, 0x9c, 0xf4, 0xa4, 0xe7, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x95, 0x8d,
	0xe8, 0x58, 0x04, 0x00, 0x00,
}
