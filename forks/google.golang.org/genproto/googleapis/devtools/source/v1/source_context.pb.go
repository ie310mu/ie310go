// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/source/v1/source_context.proto

package source

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

// The type of an Alias.
type AliasContext_Kind int32

const (
	// Do not use.
	AliasContext_ANY AliasContext_Kind = 0
	// Git tag
	AliasContext_FIXED AliasContext_Kind = 1
	// Git branch
	AliasContext_MOVABLE AliasContext_Kind = 2
	// OTHER is used to specify non-standard aliases, those not of the kinds
	// above. For example, if a Git repo has a ref named "refs/foo/bar", it
	// is considered to be of kind OTHER.
	AliasContext_OTHER AliasContext_Kind = 4
)

var AliasContext_Kind_name = map[int32]string{
	0: "ANY",
	1: "FIXED",
	2: "MOVABLE",
	4: "OTHER",
}

var AliasContext_Kind_value = map[string]int32{
	"ANY":     0,
	"FIXED":   1,
	"MOVABLE": 2,
	"OTHER":   4,
}

func (x AliasContext_Kind) String() string {
	return proto.EnumName(AliasContext_Kind_name, int32(x))
}

func (AliasContext_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{2, 0}
}

// A SourceContext is a reference to a tree of files. A SourceContext together
// with a path point to a unique revision of a single file or directory.
type SourceContext struct {
	// A SourceContext can refer any one of the following types of repositories.
	//
	// Types that are valid to be assigned to Context:
	//	*SourceContext_CloudRepo
	//	*SourceContext_CloudWorkspace
	//	*SourceContext_Gerrit
	//	*SourceContext_Git
	Context              isSourceContext_Context `protobuf_oneof:"context"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SourceContext) Reset()         { *m = SourceContext{} }
func (m *SourceContext) String() string { return proto.CompactTextString(m) }
func (*SourceContext) ProtoMessage()    {}
func (*SourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{0}
}

func (m *SourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceContext.Unmarshal(m, b)
}
func (m *SourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceContext.Marshal(b, m, deterministic)
}
func (m *SourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceContext.Merge(m, src)
}
func (m *SourceContext) XXX_Size() int {
	return xxx_messageInfo_SourceContext.Size(m)
}
func (m *SourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_SourceContext proto.InternalMessageInfo

type isSourceContext_Context interface {
	isSourceContext_Context()
}

type SourceContext_CloudRepo struct {
	CloudRepo *CloudRepoSourceContext `protobuf:"bytes,1,opt,name=cloud_repo,json=cloudRepo,proto3,oneof"`
}

type SourceContext_CloudWorkspace struct {
	CloudWorkspace *CloudWorkspaceSourceContext `protobuf:"bytes,2,opt,name=cloud_workspace,json=cloudWorkspace,proto3,oneof"`
}

type SourceContext_Gerrit struct {
	Gerrit *GerritSourceContext `protobuf:"bytes,3,opt,name=gerrit,proto3,oneof"`
}

type SourceContext_Git struct {
	Git *GitSourceContext `protobuf:"bytes,6,opt,name=git,proto3,oneof"`
}

func (*SourceContext_CloudRepo) isSourceContext_Context() {}

func (*SourceContext_CloudWorkspace) isSourceContext_Context() {}

func (*SourceContext_Gerrit) isSourceContext_Context() {}

func (*SourceContext_Git) isSourceContext_Context() {}

func (m *SourceContext) GetContext() isSourceContext_Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *SourceContext) GetCloudRepo() *CloudRepoSourceContext {
	if x, ok := m.GetContext().(*SourceContext_CloudRepo); ok {
		return x.CloudRepo
	}
	return nil
}

func (m *SourceContext) GetCloudWorkspace() *CloudWorkspaceSourceContext {
	if x, ok := m.GetContext().(*SourceContext_CloudWorkspace); ok {
		return x.CloudWorkspace
	}
	return nil
}

func (m *SourceContext) GetGerrit() *GerritSourceContext {
	if x, ok := m.GetContext().(*SourceContext_Gerrit); ok {
		return x.Gerrit
	}
	return nil
}

func (m *SourceContext) GetGit() *GitSourceContext {
	if x, ok := m.GetContext().(*SourceContext_Git); ok {
		return x.Git
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SourceContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SourceContext_CloudRepo)(nil),
		(*SourceContext_CloudWorkspace)(nil),
		(*SourceContext_Gerrit)(nil),
		(*SourceContext_Git)(nil),
	}
}

// An ExtendedSourceContext is a SourceContext combined with additional
// details describing the context.
type ExtendedSourceContext struct {
	// Any source context.
	Context *SourceContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	// Labels with user defined metadata.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExtendedSourceContext) Reset()         { *m = ExtendedSourceContext{} }
func (m *ExtendedSourceContext) String() string { return proto.CompactTextString(m) }
func (*ExtendedSourceContext) ProtoMessage()    {}
func (*ExtendedSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{1}
}

func (m *ExtendedSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtendedSourceContext.Unmarshal(m, b)
}
func (m *ExtendedSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtendedSourceContext.Marshal(b, m, deterministic)
}
func (m *ExtendedSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedSourceContext.Merge(m, src)
}
func (m *ExtendedSourceContext) XXX_Size() int {
	return xxx_messageInfo_ExtendedSourceContext.Size(m)
}
func (m *ExtendedSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedSourceContext proto.InternalMessageInfo

func (m *ExtendedSourceContext) GetContext() *SourceContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *ExtendedSourceContext) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// An alias to a repo revision.
type AliasContext struct {
	// The alias kind.
	Kind AliasContext_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=google.devtools.source.v1.AliasContext_Kind" json:"kind,omitempty"`
	// The alias name.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AliasContext) Reset()         { *m = AliasContext{} }
func (m *AliasContext) String() string { return proto.CompactTextString(m) }
func (*AliasContext) ProtoMessage()    {}
func (*AliasContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{2}
}

func (m *AliasContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AliasContext.Unmarshal(m, b)
}
func (m *AliasContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AliasContext.Marshal(b, m, deterministic)
}
func (m *AliasContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AliasContext.Merge(m, src)
}
func (m *AliasContext) XXX_Size() int {
	return xxx_messageInfo_AliasContext.Size(m)
}
func (m *AliasContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AliasContext.DiscardUnknown(m)
}

var xxx_messageInfo_AliasContext proto.InternalMessageInfo

func (m *AliasContext) GetKind() AliasContext_Kind {
	if m != nil {
		return m.Kind
	}
	return AliasContext_ANY
}

func (m *AliasContext) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A CloudRepoSourceContext denotes a particular revision in a cloud
// repo (a repo hosted by the Google Cloud Platform).
type CloudRepoSourceContext struct {
	// The ID of the repo.
	RepoId *RepoId `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	// A revision in a cloud repository can be identified by either its revision
	// ID or its Alias.
	//
	// Types that are valid to be assigned to Revision:
	//	*CloudRepoSourceContext_RevisionId
	//	*CloudRepoSourceContext_AliasName
	//	*CloudRepoSourceContext_AliasContext
	Revision             isCloudRepoSourceContext_Revision `protobuf_oneof:"revision"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CloudRepoSourceContext) Reset()         { *m = CloudRepoSourceContext{} }
func (m *CloudRepoSourceContext) String() string { return proto.CompactTextString(m) }
func (*CloudRepoSourceContext) ProtoMessage()    {}
func (*CloudRepoSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{3}
}

func (m *CloudRepoSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudRepoSourceContext.Unmarshal(m, b)
}
func (m *CloudRepoSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudRepoSourceContext.Marshal(b, m, deterministic)
}
func (m *CloudRepoSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudRepoSourceContext.Merge(m, src)
}
func (m *CloudRepoSourceContext) XXX_Size() int {
	return xxx_messageInfo_CloudRepoSourceContext.Size(m)
}
func (m *CloudRepoSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudRepoSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_CloudRepoSourceContext proto.InternalMessageInfo

func (m *CloudRepoSourceContext) GetRepoId() *RepoId {
	if m != nil {
		return m.RepoId
	}
	return nil
}

type isCloudRepoSourceContext_Revision interface {
	isCloudRepoSourceContext_Revision()
}

type CloudRepoSourceContext_RevisionId struct {
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3,oneof"`
}

type CloudRepoSourceContext_AliasName struct {
	AliasName string `protobuf:"bytes,3,opt,name=alias_name,json=aliasName,proto3,oneof"`
}

type CloudRepoSourceContext_AliasContext struct {
	AliasContext *AliasContext `protobuf:"bytes,4,opt,name=alias_context,json=aliasContext,proto3,oneof"`
}

func (*CloudRepoSourceContext_RevisionId) isCloudRepoSourceContext_Revision() {}

func (*CloudRepoSourceContext_AliasName) isCloudRepoSourceContext_Revision() {}

func (*CloudRepoSourceContext_AliasContext) isCloudRepoSourceContext_Revision() {}

func (m *CloudRepoSourceContext) GetRevision() isCloudRepoSourceContext_Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *CloudRepoSourceContext) GetRevisionId() string {
	if x, ok := m.GetRevision().(*CloudRepoSourceContext_RevisionId); ok {
		return x.RevisionId
	}
	return ""
}

// Deprecated: Do not use.
func (m *CloudRepoSourceContext) GetAliasName() string {
	if x, ok := m.GetRevision().(*CloudRepoSourceContext_AliasName); ok {
		return x.AliasName
	}
	return ""
}

func (m *CloudRepoSourceContext) GetAliasContext() *AliasContext {
	if x, ok := m.GetRevision().(*CloudRepoSourceContext_AliasContext); ok {
		return x.AliasContext
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CloudRepoSourceContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CloudRepoSourceContext_RevisionId)(nil),
		(*CloudRepoSourceContext_AliasName)(nil),
		(*CloudRepoSourceContext_AliasContext)(nil),
	}
}

// A CloudWorkspaceSourceContext denotes a workspace at a particular snapshot.
type CloudWorkspaceSourceContext struct {
	// The ID of the workspace.
	WorkspaceId *CloudWorkspaceId `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// The ID of the snapshot.
	// An empty snapshot_id refers to the most recent snapshot.
	SnapshotId           string   `protobuf:"bytes,2,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudWorkspaceSourceContext) Reset()         { *m = CloudWorkspaceSourceContext{} }
func (m *CloudWorkspaceSourceContext) String() string { return proto.CompactTextString(m) }
func (*CloudWorkspaceSourceContext) ProtoMessage()    {}
func (*CloudWorkspaceSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{4}
}

func (m *CloudWorkspaceSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudWorkspaceSourceContext.Unmarshal(m, b)
}
func (m *CloudWorkspaceSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudWorkspaceSourceContext.Marshal(b, m, deterministic)
}
func (m *CloudWorkspaceSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudWorkspaceSourceContext.Merge(m, src)
}
func (m *CloudWorkspaceSourceContext) XXX_Size() int {
	return xxx_messageInfo_CloudWorkspaceSourceContext.Size(m)
}
func (m *CloudWorkspaceSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudWorkspaceSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_CloudWorkspaceSourceContext proto.InternalMessageInfo

func (m *CloudWorkspaceSourceContext) GetWorkspaceId() *CloudWorkspaceId {
	if m != nil {
		return m.WorkspaceId
	}
	return nil
}

func (m *CloudWorkspaceSourceContext) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

// A SourceContext referring to a Gerrit project.
type GerritSourceContext struct {
	// The URI of a running Gerrit instance.
	HostUri string `protobuf:"bytes,1,opt,name=host_uri,json=hostUri,proto3" json:"host_uri,omitempty"`
	// The full project name within the host. Projects may be nested, so
	// "project/subproject" is a valid project name.
	// The "repo name" is hostURI/project.
	GerritProject string `protobuf:"bytes,2,opt,name=gerrit_project,json=gerritProject,proto3" json:"gerrit_project,omitempty"`
	// A revision in a Gerrit project can be identified by either its revision ID
	// or its alias.
	//
	// Types that are valid to be assigned to Revision:
	//	*GerritSourceContext_RevisionId
	//	*GerritSourceContext_AliasName
	//	*GerritSourceContext_AliasContext
	Revision             isGerritSourceContext_Revision `protobuf_oneof:"revision"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GerritSourceContext) Reset()         { *m = GerritSourceContext{} }
func (m *GerritSourceContext) String() string { return proto.CompactTextString(m) }
func (*GerritSourceContext) ProtoMessage()    {}
func (*GerritSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{5}
}

func (m *GerritSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GerritSourceContext.Unmarshal(m, b)
}
func (m *GerritSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GerritSourceContext.Marshal(b, m, deterministic)
}
func (m *GerritSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GerritSourceContext.Merge(m, src)
}
func (m *GerritSourceContext) XXX_Size() int {
	return xxx_messageInfo_GerritSourceContext.Size(m)
}
func (m *GerritSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_GerritSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_GerritSourceContext proto.InternalMessageInfo

func (m *GerritSourceContext) GetHostUri() string {
	if m != nil {
		return m.HostUri
	}
	return ""
}

func (m *GerritSourceContext) GetGerritProject() string {
	if m != nil {
		return m.GerritProject
	}
	return ""
}

type isGerritSourceContext_Revision interface {
	isGerritSourceContext_Revision()
}

type GerritSourceContext_RevisionId struct {
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3,oneof"`
}

type GerritSourceContext_AliasName struct {
	AliasName string `protobuf:"bytes,4,opt,name=alias_name,json=aliasName,proto3,oneof"`
}

type GerritSourceContext_AliasContext struct {
	AliasContext *AliasContext `protobuf:"bytes,5,opt,name=alias_context,json=aliasContext,proto3,oneof"`
}

func (*GerritSourceContext_RevisionId) isGerritSourceContext_Revision() {}

func (*GerritSourceContext_AliasName) isGerritSourceContext_Revision() {}

func (*GerritSourceContext_AliasContext) isGerritSourceContext_Revision() {}

func (m *GerritSourceContext) GetRevision() isGerritSourceContext_Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *GerritSourceContext) GetRevisionId() string {
	if x, ok := m.GetRevision().(*GerritSourceContext_RevisionId); ok {
		return x.RevisionId
	}
	return ""
}

// Deprecated: Do not use.
func (m *GerritSourceContext) GetAliasName() string {
	if x, ok := m.GetRevision().(*GerritSourceContext_AliasName); ok {
		return x.AliasName
	}
	return ""
}

func (m *GerritSourceContext) GetAliasContext() *AliasContext {
	if x, ok := m.GetRevision().(*GerritSourceContext_AliasContext); ok {
		return x.AliasContext
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GerritSourceContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GerritSourceContext_RevisionId)(nil),
		(*GerritSourceContext_AliasName)(nil),
		(*GerritSourceContext_AliasContext)(nil),
	}
}

// A GitSourceContext denotes a particular revision in a third party Git
// repository (e.g. GitHub).
type GitSourceContext struct {
	// Git repository URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Git commit hash.
	// required.
	RevisionId           string   `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitSourceContext) Reset()         { *m = GitSourceContext{} }
func (m *GitSourceContext) String() string { return proto.CompactTextString(m) }
func (*GitSourceContext) ProtoMessage()    {}
func (*GitSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{6}
}

func (m *GitSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitSourceContext.Unmarshal(m, b)
}
func (m *GitSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitSourceContext.Marshal(b, m, deterministic)
}
func (m *GitSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitSourceContext.Merge(m, src)
}
func (m *GitSourceContext) XXX_Size() int {
	return xxx_messageInfo_GitSourceContext.Size(m)
}
func (m *GitSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_GitSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_GitSourceContext proto.InternalMessageInfo

func (m *GitSourceContext) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *GitSourceContext) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

// A unique identifier for a cloud repo.
type RepoId struct {
	// A cloud repository can be identified by either its project ID and
	// repository name combination, or its globally unique identifier.
	//
	// Types that are valid to be assigned to Id:
	//	*RepoId_ProjectRepoId
	//	*RepoId_Uid
	Id                   isRepoId_Id `protobuf_oneof:"id"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RepoId) Reset()         { *m = RepoId{} }
func (m *RepoId) String() string { return proto.CompactTextString(m) }
func (*RepoId) ProtoMessage()    {}
func (*RepoId) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{7}
}

func (m *RepoId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepoId.Unmarshal(m, b)
}
func (m *RepoId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepoId.Marshal(b, m, deterministic)
}
func (m *RepoId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoId.Merge(m, src)
}
func (m *RepoId) XXX_Size() int {
	return xxx_messageInfo_RepoId.Size(m)
}
func (m *RepoId) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoId.DiscardUnknown(m)
}

var xxx_messageInfo_RepoId proto.InternalMessageInfo

type isRepoId_Id interface {
	isRepoId_Id()
}

type RepoId_ProjectRepoId struct {
	ProjectRepoId *ProjectRepoId `protobuf:"bytes,1,opt,name=project_repo_id,json=projectRepoId,proto3,oneof"`
}

type RepoId_Uid struct {
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3,oneof"`
}

func (*RepoId_ProjectRepoId) isRepoId_Id() {}

func (*RepoId_Uid) isRepoId_Id() {}

func (m *RepoId) GetId() isRepoId_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RepoId) GetProjectRepoId() *ProjectRepoId {
	if x, ok := m.GetId().(*RepoId_ProjectRepoId); ok {
		return x.ProjectRepoId
	}
	return nil
}

func (m *RepoId) GetUid() string {
	if x, ok := m.GetId().(*RepoId_Uid); ok {
		return x.Uid
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RepoId) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RepoId_ProjectRepoId)(nil),
		(*RepoId_Uid)(nil),
	}
}

// Selects a repo using a Google Cloud Platform project ID
// (e.g. winged-cargo-31) and a repo name within that project.
type ProjectRepoId struct {
	// The ID of the project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The name of the repo. Leave empty for the default repo.
	RepoName             string   `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectRepoId) Reset()         { *m = ProjectRepoId{} }
func (m *ProjectRepoId) String() string { return proto.CompactTextString(m) }
func (*ProjectRepoId) ProtoMessage()    {}
func (*ProjectRepoId) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{8}
}

func (m *ProjectRepoId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectRepoId.Unmarshal(m, b)
}
func (m *ProjectRepoId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectRepoId.Marshal(b, m, deterministic)
}
func (m *ProjectRepoId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectRepoId.Merge(m, src)
}
func (m *ProjectRepoId) XXX_Size() int {
	return xxx_messageInfo_ProjectRepoId.Size(m)
}
func (m *ProjectRepoId) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectRepoId.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectRepoId proto.InternalMessageInfo

func (m *ProjectRepoId) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ProjectRepoId) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

// A CloudWorkspaceId is a unique identifier for a cloud workspace.
// A cloud workspace is a place associated with a repo where modified files
// can be stored before they are committed.
type CloudWorkspaceId struct {
	// The ID of the repo containing the workspace.
	RepoId *RepoId `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	// The unique name of the workspace within the repo.  This is the name
	// chosen by the client in the Source API's CreateWorkspace method.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudWorkspaceId) Reset()         { *m = CloudWorkspaceId{} }
func (m *CloudWorkspaceId) String() string { return proto.CompactTextString(m) }
func (*CloudWorkspaceId) ProtoMessage()    {}
func (*CloudWorkspaceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_abda43b4d0b03743, []int{9}
}

func (m *CloudWorkspaceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudWorkspaceId.Unmarshal(m, b)
}
func (m *CloudWorkspaceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudWorkspaceId.Marshal(b, m, deterministic)
}
func (m *CloudWorkspaceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudWorkspaceId.Merge(m, src)
}
func (m *CloudWorkspaceId) XXX_Size() int {
	return xxx_messageInfo_CloudWorkspaceId.Size(m)
}
func (m *CloudWorkspaceId) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudWorkspaceId.DiscardUnknown(m)
}

var xxx_messageInfo_CloudWorkspaceId proto.InternalMessageInfo

func (m *CloudWorkspaceId) GetRepoId() *RepoId {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *CloudWorkspaceId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.devtools.source.v1.AliasContext_Kind", AliasContext_Kind_name, AliasContext_Kind_value)
	proto.RegisterType((*SourceContext)(nil), "google.devtools.source.v1.SourceContext")
	proto.RegisterType((*ExtendedSourceContext)(nil), "google.devtools.source.v1.ExtendedSourceContext")
	proto.RegisterMapType((map[string]string)(nil), "google.devtools.source.v1.ExtendedSourceContext.LabelsEntry")
	proto.RegisterType((*AliasContext)(nil), "google.devtools.source.v1.AliasContext")
	proto.RegisterType((*CloudRepoSourceContext)(nil), "google.devtools.source.v1.CloudRepoSourceContext")
	proto.RegisterType((*CloudWorkspaceSourceContext)(nil), "google.devtools.source.v1.CloudWorkspaceSourceContext")
	proto.RegisterType((*GerritSourceContext)(nil), "google.devtools.source.v1.GerritSourceContext")
	proto.RegisterType((*GitSourceContext)(nil), "google.devtools.source.v1.GitSourceContext")
	proto.RegisterType((*RepoId)(nil), "google.devtools.source.v1.RepoId")
	proto.RegisterType((*ProjectRepoId)(nil), "google.devtools.source.v1.ProjectRepoId")
	proto.RegisterType((*CloudWorkspaceId)(nil), "google.devtools.source.v1.CloudWorkspaceId")
}

func init() {
	proto.RegisterFile("google/devtools/source/v1/source_context.proto", fileDescriptor_abda43b4d0b03743)
}

var fileDescriptor_abda43b4d0b03743 = []byte{
	// 809 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x4e, 0xdb, 0x4a,
	0x14, 0x8e, 0x9d, 0x90, 0xe0, 0x13, 0x02, 0xd1, 0xdc, 0x1f, 0x05, 0xb8, 0x08, 0xf0, 0xd5, 0xd5,
	0x45, 0xa2, 0x72, 0x94, 0x54, 0xaa, 0x5a, 0x5a, 0x89, 0x62, 0x48, 0x49, 0x04, 0x0d, 0x68, 0x4a,
	0xe9, 0x8f, 0x22, 0x45, 0xc6, 0x1e, 0x19, 0x17, 0xe3, 0xb1, 0x6c, 0x27, 0xc0, 0x26, 0xfa, 0xdc,
	0x35, 0x74, 0x0f, 0xdd, 0x40, 0xd7, 0xd0, 0x25, 0x74, 0x01, 0x48, 0x7d, 0xa9, 0x66, 0xc6, 0x86,
	0x24, 0x04, 0x83, 0xd4, 0x3e, 0x79, 0xe6, 0xf8, 0xfb, 0xbe, 0x73, 0xe6, 0xfc, 0xcc, 0x80, 0x66,
	0x53, 0x6a, 0xbb, 0xa4, 0x6a, 0x91, 0x7e, 0x44, 0xa9, 0x1b, 0x56, 0x43, 0xda, 0x0b, 0x4c, 0x52,
	0xed, 0xd7, 0xe2, 0x55, 0xd7, 0xa4, 0x5e, 0x44, 0xce, 0x23, 0xcd, 0x0f, 0x68, 0x44, 0xd1, 0xac,
	0xc0, 0x6b, 0x09, 0x5e, 0x13, 0x28, 0xad, 0x5f, 0x9b, 0xfb, 0x27, 0x96, 0x32, 0x7c, 0xa7, 0x6a,
	0x78, 0x1e, 0x8d, 0x8c, 0xc8, 0xa1, 0x5e, 0x28, 0x88, 0xea, 0x37, 0x19, 0x4a, 0xaf, 0x38, 0x76,
	0x53, 0x08, 0x22, 0x0c, 0x60, 0xba, 0xb4, 0x67, 0x75, 0x03, 0xe2, 0xd3, 0x8a, 0xb4, 0x24, 0xad,
	0x14, 0xeb, 0x35, 0xed, 0x56, 0x7d, 0x6d, 0x93, 0x81, 0x31, 0xf1, 0xe9, 0x90, 0x4c, 0x33, 0x83,
	0x15, 0x33, 0xf9, 0x83, 0x0c, 0x98, 0x11, 0x9a, 0x67, 0x34, 0x38, 0x09, 0x7d, 0xc3, 0x24, 0x15,
	0x99, 0x0b, 0x3f, 0xba, 0x4b, 0xf8, 0x4d, 0x42, 0x18, 0x55, 0x9f, 0x36, 0x87, 0x7e, 0xa3, 0x26,
	0xe4, 0x6d, 0x12, 0x04, 0x4e, 0x54, 0xc9, 0x72, 0x65, 0x2d, 0x45, 0x79, 0x9b, 0x03, 0x47, 0x15,
	0x63, 0x3e, 0x5a, 0x87, 0xac, 0xed, 0x44, 0x95, 0x3c, 0x97, 0x59, 0x4d, 0x93, 0xb9, 0xa9, 0xc1,
	0x98, 0xba, 0x02, 0x85, 0xb8, 0x3a, 0xea, 0x77, 0x09, 0xfe, 0x6a, 0x9c, 0x47, 0xc4, 0xb3, 0x88,
	0x35, 0x9c, 0x66, 0xfd, 0x0a, 0x14, 0xe7, 0x78, 0x25, 0xc5, 0xd3, 0x10, 0x15, 0x27, 0x44, 0x74,
	0x00, 0x79, 0xd7, 0x38, 0x22, 0x6e, 0x58, 0x91, 0x97, 0xb2, 0x2b, 0xc5, 0xfa, 0xb3, 0x14, 0x89,
	0xb1, 0x51, 0x68, 0xbb, 0x9c, 0xde, 0xf0, 0xa2, 0xe0, 0x02, 0xc7, 0x5a, 0x73, 0x4f, 0xa0, 0x38,
	0x60, 0x46, 0x65, 0xc8, 0x9e, 0x90, 0x0b, 0x1e, 0xa4, 0x82, 0xd9, 0x12, 0xfd, 0x09, 0x13, 0x7d,
	0xc3, 0xed, 0x89, 0x1a, 0x2a, 0x58, 0x6c, 0xd6, 0xe4, 0xc7, 0x92, 0xfa, 0x49, 0x82, 0xa9, 0x0d,
	0xd7, 0x31, 0xc2, 0xe4, 0x94, 0xcf, 0x21, 0x77, 0xe2, 0x78, 0x16, 0x67, 0x4f, 0xd7, 0x1f, 0xa4,
	0xc4, 0x37, 0x48, 0xd3, 0x76, 0x1c, 0xcf, 0xc2, 0x9c, 0x89, 0x10, 0xe4, 0x3c, 0xe3, 0x34, 0xf1,
	0xc5, 0xd7, 0x6a, 0x1d, 0x72, 0x0c, 0x81, 0x0a, 0x90, 0xdd, 0x68, 0xbf, 0x2b, 0x67, 0x90, 0x02,
	0x13, 0x2f, 0x5a, 0x6f, 0x1b, 0x5b, 0x65, 0x09, 0x15, 0xa1, 0xf0, 0x72, 0xef, 0x70, 0x43, 0xdf,
	0x6d, 0x94, 0x65, 0x66, 0xdf, 0x3b, 0x68, 0x36, 0x70, 0x39, 0xa7, 0x5e, 0x4a, 0xf0, 0xf7, 0xf8,
	0x56, 0x45, 0x6b, 0x50, 0x60, 0xbd, 0xde, 0x75, 0xac, 0xb8, 0x14, 0xcb, 0x29, 0x71, 0x32, 0x7a,
	0xcb, 0xc2, 0xf9, 0x80, 0x7f, 0xd1, 0x32, 0x14, 0x03, 0xd2, 0x77, 0x42, 0x87, 0x7a, 0x8c, 0xcf,
	0xa3, 0x6c, 0x66, 0x30, 0x24, 0xc6, 0x96, 0x85, 0xfe, 0x05, 0x30, 0xd8, 0xe1, 0xba, 0xfc, 0x1c,
	0xac, 0x3b, 0x15, 0x5d, 0xae, 0x48, 0x6c, 0x42, 0xb8, 0xbd, 0x6d, 0x9c, 0x12, 0xd4, 0x86, 0x92,
	0x00, 0x25, 0x4d, 0x91, 0xe3, 0x91, 0xfc, 0x7f, 0xcf, 0x8c, 0x35, 0x33, 0x78, 0xca, 0x18, 0xd8,
	0xeb, 0x00, 0x93, 0x49, 0x08, 0xea, 0x47, 0x09, 0xe6, 0x53, 0x86, 0x09, 0xb5, 0x61, 0xea, 0x6a,
	0x2e, 0xaf, 0x93, 0xb0, 0x7a, 0xef, 0xd1, 0x6c, 0x59, 0xb8, 0x78, 0x76, 0xbd, 0x41, 0x8b, 0x50,
	0x0c, 0x3d, 0xc3, 0x0f, 0x8f, 0x69, 0x74, 0x95, 0x13, 0x0c, 0x89, 0xa9, 0x65, 0xa9, 0x3f, 0x24,
	0xf8, 0x63, 0xcc, 0x0c, 0xa2, 0x59, 0x98, 0x3c, 0xa6, 0x61, 0xd4, 0xed, 0x05, 0x4e, 0xdc, 0x6f,
	0x05, 0xb6, 0x7f, 0x1d, 0x38, 0xe8, 0x3f, 0x98, 0x16, 0xe3, 0xd9, 0xf5, 0x03, 0xfa, 0x81, 0x98,
	0x51, 0x2c, 0x5b, 0x12, 0xd6, 0x7d, 0x61, 0x1c, 0x2d, 0x47, 0xf6, 0xce, 0x72, 0xe4, 0xee, 0x59,
	0x8e, 0x89, 0xdf, 0x57, 0x8e, 0x06, 0x94, 0x47, 0x6f, 0x0e, 0x36, 0x64, 0xbd, 0xc0, 0x4d, 0x86,
	0xac, 0x17, 0xb8, 0x2c, 0x89, 0x37, 0x1a, 0x6b, 0xf0, 0x1c, 0x6a, 0x1f, 0xf2, 0xa2, 0x17, 0x11,
	0x86, 0x99, 0x38, 0x29, 0xdd, 0xe1, 0x3e, 0x4e, 0xbb, 0x52, 0xe2, 0x8c, 0x09, 0x89, 0x66, 0x06,
	0x97, 0xfc, 0x41, 0x03, 0x42, 0x90, 0xed, 0x0d, 0xf4, 0x33, 0xdb, 0xe8, 0x39, 0x90, 0x1d, 0x4b,
	0xdd, 0x81, 0xd2, 0x10, 0x17, 0x2d, 0x00, 0x24, 0xee, 0x63, 0xcf, 0x0a, 0x56, 0x62, 0x4b, 0xcb,
	0x42, 0xf3, 0xa0, 0xf0, 0xa8, 0x06, 0xa6, 0x78, 0x92, 0x19, 0x58, 0x9e, 0xd5, 0x23, 0x28, 0x8f,
	0xf6, 0xd2, 0x2f, 0x8d, 0xe3, 0x98, 0xdb, 0x42, 0xff, 0x22, 0xc1, 0x82, 0x49, 0x4f, 0x6f, 0x17,
	0xd1, 0xd1, 0x50, 0x31, 0xf6, 0xd9, 0xc3, 0xb8, 0x2f, 0xbd, 0x5f, 0x8f, 0x09, 0x36, 0x75, 0x0d,
	0xcf, 0xd6, 0x68, 0x60, 0x57, 0x6d, 0xe2, 0xf1, 0x67, 0xb3, 0x2a, 0x7e, 0x19, 0xbe, 0x13, 0x8e,
	0x79, 0xa2, 0x9f, 0x8a, 0xd5, 0xa5, 0x24, 0x7d, 0x96, 0x17, 0xb7, 0x85, 0x08, 0x3f, 0xa5, 0xb6,
	0x45, 0xfa, 0x07, 0xdc, 0xb7, 0x70, 0xa8, 0x1d, 0xd6, 0xbe, 0x26, 0x88, 0x0e, 0x47, 0x74, 0x12,
	0x44, 0x47, 0x20, 0x3a, 0x87, 0xb5, 0xa3, 0x3c, 0x77, 0xf9, 0xf0, 0x67, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x98, 0x86, 0xf5, 0x6f, 0x14, 0x08, 0x00, 0x00,
}
