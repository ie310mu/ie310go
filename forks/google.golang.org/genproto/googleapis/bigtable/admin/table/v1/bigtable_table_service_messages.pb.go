// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/bigtable/admin/table/v1/bigtable_table_service_messages.proto

package table

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
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

type CreateTableRequest struct {
	// The unique name of the cluster in which to create the new table.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name by which the new table should be referred to within the cluster,
	// e.g. "foobar" rather than "<cluster_name>/tables/foobar".
	TableId string `protobuf:"bytes,2,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	// The Table to create. The `name` field of the Table and all of its
	// ColumnFamilies must be left blank, and will be populated in the response.
	Table *Table `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	// The optional list of row keys that will be used to initially split the
	// table into several tablets (Tablets are similar to HBase regions).
	// Given two split keys, "s1" and "s2", three tablets will be created,
	// spanning the key ranges: [, s1), [s1, s2), [s2, ).
	//
	// Example:
	//  * Row keys := ["a", "apple", "custom", "customer_1", "customer_2",
	//                 "other", "zz"]
	//  * initial_split_keys := ["apple", "customer_1", "customer_2", "other"]
	//  * Key assignment:
	//    - Tablet 1 [, apple)                => {"a"}.
	//    - Tablet 2 [apple, customer_1)      => {"apple", "custom"}.
	//    - Tablet 3 [customer_1, customer_2) => {"customer_1"}.
	//    - Tablet 4 [customer_2, other)      => {"customer_2"}.
	//    - Tablet 5 [other, )                => {"other", "zz"}.
	InitialSplitKeys     []string `protobuf:"bytes,4,rep,name=initial_split_keys,json=initialSplitKeys,proto3" json:"initial_split_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTableRequest) Reset()         { *m = CreateTableRequest{} }
func (m *CreateTableRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTableRequest) ProtoMessage()    {}
func (*CreateTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{0}
}

func (m *CreateTableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTableRequest.Unmarshal(m, b)
}
func (m *CreateTableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTableRequest.Marshal(b, m, deterministic)
}
func (m *CreateTableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTableRequest.Merge(m, src)
}
func (m *CreateTableRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTableRequest.Size(m)
}
func (m *CreateTableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTableRequest proto.InternalMessageInfo

func (m *CreateTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTableRequest) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

func (m *CreateTableRequest) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *CreateTableRequest) GetInitialSplitKeys() []string {
	if m != nil {
		return m.InitialSplitKeys
	}
	return nil
}

type ListTablesRequest struct {
	// The unique name of the cluster for which tables should be listed.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTablesRequest) Reset()         { *m = ListTablesRequest{} }
func (m *ListTablesRequest) String() string { return proto.CompactTextString(m) }
func (*ListTablesRequest) ProtoMessage()    {}
func (*ListTablesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{1}
}

func (m *ListTablesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTablesRequest.Unmarshal(m, b)
}
func (m *ListTablesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTablesRequest.Marshal(b, m, deterministic)
}
func (m *ListTablesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTablesRequest.Merge(m, src)
}
func (m *ListTablesRequest) XXX_Size() int {
	return xxx_messageInfo_ListTablesRequest.Size(m)
}
func (m *ListTablesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTablesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTablesRequest proto.InternalMessageInfo

func (m *ListTablesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListTablesResponse struct {
	// The tables present in the requested cluster.
	// At present, only the names of the tables are populated.
	Tables               []*Table `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTablesResponse) Reset()         { *m = ListTablesResponse{} }
func (m *ListTablesResponse) String() string { return proto.CompactTextString(m) }
func (*ListTablesResponse) ProtoMessage()    {}
func (*ListTablesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{2}
}

func (m *ListTablesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTablesResponse.Unmarshal(m, b)
}
func (m *ListTablesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTablesResponse.Marshal(b, m, deterministic)
}
func (m *ListTablesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTablesResponse.Merge(m, src)
}
func (m *ListTablesResponse) XXX_Size() int {
	return xxx_messageInfo_ListTablesResponse.Size(m)
}
func (m *ListTablesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTablesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTablesResponse proto.InternalMessageInfo

func (m *ListTablesResponse) GetTables() []*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

type GetTableRequest struct {
	// The unique name of the requested table.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTableRequest) Reset()         { *m = GetTableRequest{} }
func (m *GetTableRequest) String() string { return proto.CompactTextString(m) }
func (*GetTableRequest) ProtoMessage()    {}
func (*GetTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{3}
}

func (m *GetTableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTableRequest.Unmarshal(m, b)
}
func (m *GetTableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTableRequest.Marshal(b, m, deterministic)
}
func (m *GetTableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTableRequest.Merge(m, src)
}
func (m *GetTableRequest) XXX_Size() int {
	return xxx_messageInfo_GetTableRequest.Size(m)
}
func (m *GetTableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTableRequest proto.InternalMessageInfo

func (m *GetTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteTableRequest struct {
	// The unique name of the table to be deleted.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTableRequest) Reset()         { *m = DeleteTableRequest{} }
func (m *DeleteTableRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTableRequest) ProtoMessage()    {}
func (*DeleteTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{4}
}

func (m *DeleteTableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTableRequest.Unmarshal(m, b)
}
func (m *DeleteTableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTableRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTableRequest.Merge(m, src)
}
func (m *DeleteTableRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTableRequest.Size(m)
}
func (m *DeleteTableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTableRequest proto.InternalMessageInfo

func (m *DeleteTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RenameTableRequest struct {
	// The current unique name of the table.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The new name by which the table should be referred to within its containing
	// cluster, e.g. "foobar" rather than "<cluster_name>/tables/foobar".
	NewId                string   `protobuf:"bytes,2,opt,name=new_id,json=newId,proto3" json:"new_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameTableRequest) Reset()         { *m = RenameTableRequest{} }
func (m *RenameTableRequest) String() string { return proto.CompactTextString(m) }
func (*RenameTableRequest) ProtoMessage()    {}
func (*RenameTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{5}
}

func (m *RenameTableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameTableRequest.Unmarshal(m, b)
}
func (m *RenameTableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameTableRequest.Marshal(b, m, deterministic)
}
func (m *RenameTableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameTableRequest.Merge(m, src)
}
func (m *RenameTableRequest) XXX_Size() int {
	return xxx_messageInfo_RenameTableRequest.Size(m)
}
func (m *RenameTableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameTableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenameTableRequest proto.InternalMessageInfo

func (m *RenameTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RenameTableRequest) GetNewId() string {
	if m != nil {
		return m.NewId
	}
	return ""
}

type CreateColumnFamilyRequest struct {
	// The unique name of the table in which to create the new column family.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name by which the new column family should be referred to within the
	// table, e.g. "foobar" rather than "<table_name>/columnFamilies/foobar".
	ColumnFamilyId string `protobuf:"bytes,2,opt,name=column_family_id,json=columnFamilyId,proto3" json:"column_family_id,omitempty"`
	// The column family to create. The `name` field must be left blank.
	ColumnFamily         *ColumnFamily `protobuf:"bytes,3,opt,name=column_family,json=columnFamily,proto3" json:"column_family,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateColumnFamilyRequest) Reset()         { *m = CreateColumnFamilyRequest{} }
func (m *CreateColumnFamilyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateColumnFamilyRequest) ProtoMessage()    {}
func (*CreateColumnFamilyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{6}
}

func (m *CreateColumnFamilyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateColumnFamilyRequest.Unmarshal(m, b)
}
func (m *CreateColumnFamilyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateColumnFamilyRequest.Marshal(b, m, deterministic)
}
func (m *CreateColumnFamilyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateColumnFamilyRequest.Merge(m, src)
}
func (m *CreateColumnFamilyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateColumnFamilyRequest.Size(m)
}
func (m *CreateColumnFamilyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateColumnFamilyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateColumnFamilyRequest proto.InternalMessageInfo

func (m *CreateColumnFamilyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateColumnFamilyRequest) GetColumnFamilyId() string {
	if m != nil {
		return m.ColumnFamilyId
	}
	return ""
}

func (m *CreateColumnFamilyRequest) GetColumnFamily() *ColumnFamily {
	if m != nil {
		return m.ColumnFamily
	}
	return nil
}

type DeleteColumnFamilyRequest struct {
	// The unique name of the column family to be deleted.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteColumnFamilyRequest) Reset()         { *m = DeleteColumnFamilyRequest{} }
func (m *DeleteColumnFamilyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteColumnFamilyRequest) ProtoMessage()    {}
func (*DeleteColumnFamilyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{7}
}

func (m *DeleteColumnFamilyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteColumnFamilyRequest.Unmarshal(m, b)
}
func (m *DeleteColumnFamilyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteColumnFamilyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteColumnFamilyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteColumnFamilyRequest.Merge(m, src)
}
func (m *DeleteColumnFamilyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteColumnFamilyRequest.Size(m)
}
func (m *DeleteColumnFamilyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteColumnFamilyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteColumnFamilyRequest proto.InternalMessageInfo

func (m *DeleteColumnFamilyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BulkDeleteRowsRequest struct {
	// The unique name of the table on which to perform the bulk delete
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// Types that are valid to be assigned to Target:
	//	*BulkDeleteRowsRequest_RowKeyPrefix
	//	*BulkDeleteRowsRequest_DeleteAllDataFromTable
	Target               isBulkDeleteRowsRequest_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *BulkDeleteRowsRequest) Reset()         { *m = BulkDeleteRowsRequest{} }
func (m *BulkDeleteRowsRequest) String() string { return proto.CompactTextString(m) }
func (*BulkDeleteRowsRequest) ProtoMessage()    {}
func (*BulkDeleteRowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a159d72e7e8b0be6, []int{8}
}

func (m *BulkDeleteRowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkDeleteRowsRequest.Unmarshal(m, b)
}
func (m *BulkDeleteRowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkDeleteRowsRequest.Marshal(b, m, deterministic)
}
func (m *BulkDeleteRowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkDeleteRowsRequest.Merge(m, src)
}
func (m *BulkDeleteRowsRequest) XXX_Size() int {
	return xxx_messageInfo_BulkDeleteRowsRequest.Size(m)
}
func (m *BulkDeleteRowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkDeleteRowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkDeleteRowsRequest proto.InternalMessageInfo

func (m *BulkDeleteRowsRequest) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

type isBulkDeleteRowsRequest_Target interface {
	isBulkDeleteRowsRequest_Target()
}

type BulkDeleteRowsRequest_RowKeyPrefix struct {
	RowKeyPrefix []byte `protobuf:"bytes,2,opt,name=row_key_prefix,json=rowKeyPrefix,proto3,oneof"`
}

type BulkDeleteRowsRequest_DeleteAllDataFromTable struct {
	DeleteAllDataFromTable bool `protobuf:"varint,3,opt,name=delete_all_data_from_table,json=deleteAllDataFromTable,proto3,oneof"`
}

func (*BulkDeleteRowsRequest_RowKeyPrefix) isBulkDeleteRowsRequest_Target() {}

func (*BulkDeleteRowsRequest_DeleteAllDataFromTable) isBulkDeleteRowsRequest_Target() {}

func (m *BulkDeleteRowsRequest) GetTarget() isBulkDeleteRowsRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *BulkDeleteRowsRequest) GetRowKeyPrefix() []byte {
	if x, ok := m.GetTarget().(*BulkDeleteRowsRequest_RowKeyPrefix); ok {
		return x.RowKeyPrefix
	}
	return nil
}

func (m *BulkDeleteRowsRequest) GetDeleteAllDataFromTable() bool {
	if x, ok := m.GetTarget().(*BulkDeleteRowsRequest_DeleteAllDataFromTable); ok {
		return x.DeleteAllDataFromTable
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BulkDeleteRowsRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BulkDeleteRowsRequest_RowKeyPrefix)(nil),
		(*BulkDeleteRowsRequest_DeleteAllDataFromTable)(nil),
	}
}

func init() {
	proto.RegisterType((*CreateTableRequest)(nil), "google.bigtable.admin.table.v1.CreateTableRequest")
	proto.RegisterType((*ListTablesRequest)(nil), "google.bigtable.admin.table.v1.ListTablesRequest")
	proto.RegisterType((*ListTablesResponse)(nil), "google.bigtable.admin.table.v1.ListTablesResponse")
	proto.RegisterType((*GetTableRequest)(nil), "google.bigtable.admin.table.v1.GetTableRequest")
	proto.RegisterType((*DeleteTableRequest)(nil), "google.bigtable.admin.table.v1.DeleteTableRequest")
	proto.RegisterType((*RenameTableRequest)(nil), "google.bigtable.admin.table.v1.RenameTableRequest")
	proto.RegisterType((*CreateColumnFamilyRequest)(nil), "google.bigtable.admin.table.v1.CreateColumnFamilyRequest")
	proto.RegisterType((*DeleteColumnFamilyRequest)(nil), "google.bigtable.admin.table.v1.DeleteColumnFamilyRequest")
	proto.RegisterType((*BulkDeleteRowsRequest)(nil), "google.bigtable.admin.table.v1.BulkDeleteRowsRequest")
}

func init() {
	proto.RegisterFile("google/bigtable/admin/table/v1/bigtable_table_service_messages.proto", fileDescriptor_a159d72e7e8b0be6)
}

var fileDescriptor_a159d72e7e8b0be6 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x49, 0x1b, 0x92, 0x21, 0x94, 0xb2, 0x52, 0x51, 0x52, 0x09, 0x14, 0x56, 0x2a, 0xe4,
	0x50, 0xd9, 0x2a, 0x5c, 0x90, 0x0a, 0x42, 0x24, 0x51, 0x69, 0x54, 0x40, 0xc1, 0xe1, 0xc4, 0xc5,
	0xda, 0xc4, 0x13, 0x6b, 0xd5, 0xb5, 0x37, 0xec, 0x6e, 0x12, 0xf2, 0x13, 0x7c, 0x06, 0x27, 0xc4,
	0x37, 0x22, 0xef, 0x9a, 0x26, 0x3d, 0x10, 0x97, 0x8b, 0x35, 0x9e, 0x79, 0xf3, 0x66, 0xf6, 0xcd,
	0x0c, 0xf4, 0x13, 0x29, 0x13, 0x81, 0xc1, 0x98, 0x27, 0x86, 0x8d, 0x05, 0x06, 0x2c, 0x4e, 0x79,
	0x16, 0x38, 0x7b, 0x71, 0x7a, 0xed, 0x8f, 0xdc, 0x57, 0xa3, 0x5a, 0xf0, 0x09, 0x46, 0x29, 0x6a,
	0xcd, 0x12, 0xd4, 0xfe, 0x4c, 0x49, 0x23, 0xc9, 0x13, 0xc7, 0xe2, 0xff, 0x45, 0xfb, 0x96, 0xc5,
	0x77, 0xf6, 0xe2, 0xf4, 0xe8, 0xd5, 0xff, 0x55, 0x89, 0x99, 0x61, 0x8e, 0x99, 0xfe, 0xf6, 0x80,
	0xf4, 0x14, 0x32, 0x83, 0x5f, 0xf2, 0x50, 0x88, 0xdf, 0xe6, 0xa8, 0x0d, 0x21, 0xb0, 0x9b, 0xb1,
	0x14, 0x9b, 0x5e, 0xdb, 0xeb, 0xd4, 0x43, 0x6b, 0x93, 0x16, 0xd4, 0x5c, 0x3a, 0x8f, 0x9b, 0x77,
	0xac, 0xff, 0xae, 0xfd, 0x1f, 0xc4, 0xe4, 0x0c, 0xf6, 0xac, 0xd9, 0xac, 0xb4, 0xbd, 0xce, 0xbd,
	0x17, 0xc7, 0xfe, 0xf6, 0x7e, 0x7d, 0x57, 0xcb, 0xe5, 0x90, 0x13, 0x20, 0x3c, 0xe3, 0x86, 0x33,
	0x11, 0xe9, 0x99, 0xe0, 0x26, 0xba, 0xc2, 0x95, 0x6e, 0xee, 0xb6, 0x2b, 0x9d, 0x7a, 0x78, 0x50,
	0x44, 0x46, 0x79, 0xe0, 0x12, 0x57, 0x9a, 0x3e, 0x87, 0x87, 0x1f, 0xb8, 0x36, 0x96, 0x41, 0x6f,
	0x69, 0x97, 0x8e, 0x80, 0x6c, 0x02, 0xf5, 0x4c, 0x66, 0x1a, 0xc9, 0x1b, 0xa8, 0xda, 0xaa, 0xba,
	0xe9, 0xb5, 0x2b, 0xb7, 0x6f, 0xb5, 0x48, 0xa2, 0xc7, 0xf0, 0xe0, 0x3d, 0x9a, 0x32, 0xa9, 0x68,
	0x07, 0x48, 0x1f, 0x05, 0x96, 0x8b, 0x4a, 0xdf, 0x02, 0x09, 0x31, 0xb7, 0x4a, 0xe5, 0x3f, 0x84,
	0x6a, 0x86, 0xcb, 0xb5, 0xf8, 0x7b, 0x19, 0x2e, 0x07, 0x31, 0xfd, 0xe5, 0x41, 0xcb, 0x0d, 0xb0,
	0x27, 0xc5, 0x3c, 0xcd, 0xce, 0x59, 0xca, 0xc5, 0x6a, 0x1b, 0x51, 0x07, 0x0e, 0x26, 0x16, 0x1a,
	0x4d, 0x2d, 0x76, 0x4d, 0xb9, 0x3f, 0xd9, 0xa0, 0x18, 0xc4, 0xe4, 0x33, 0xdc, 0xbf, 0x81, 0x2c,
	0xc6, 0x7b, 0x52, 0xa6, 0xd9, 0x8d, 0x4e, 0x1a, 0x9b, 0xa4, 0x34, 0x80, 0x96, 0x53, 0xe6, 0x96,
	0xdd, 0xd2, 0x9f, 0x1e, 0x1c, 0x76, 0xe7, 0xe2, 0xca, 0x65, 0x85, 0x72, 0x79, 0x3d, 0xf4, 0xc7,
	0x00, 0x6e, 0x1f, 0x37, 0x72, 0xea, 0xd6, 0xf3, 0x29, 0x7f, 0xe6, 0x33, 0xd8, 0x57, 0x72, 0x99,
	0x2f, 0x53, 0x34, 0x53, 0x38, 0xe5, 0xdf, 0xed, 0x23, 0x1b, 0x17, 0x3b, 0x61, 0x43, 0xc9, 0xe5,
	0x25, 0xae, 0x86, 0xd6, 0x4b, 0x5e, 0xc3, 0x51, 0x6c, 0xb9, 0x23, 0x26, 0x84, 0x3d, 0x8d, 0x68,
	0xaa, 0x64, 0x1a, 0xad, 0x17, 0xba, 0x76, 0xb1, 0x13, 0x3e, 0x72, 0x98, 0x77, 0x42, 0xf4, 0x99,
	0x61, 0xe7, 0x4a, 0xa6, 0x76, 0x60, 0xdd, 0x5a, 0xbe, 0x4f, 0x2a, 0x41, 0xd3, 0xfd, 0xe1, 0x01,
	0x9d, 0xc8, 0xb4, 0x44, 0x9b, 0xee, 0xd3, 0x6e, 0x11, 0xb0, 0xf9, 0x23, 0x77, 0xef, 0x1f, 0x8b,
	0x73, 0x1f, 0xe6, 0x37, 0x39, 0xf4, 0xbe, 0xf6, 0x0a, 0x92, 0x44, 0x0a, 0x96, 0x25, 0xbe, 0x54,
	0x49, 0x90, 0x60, 0x66, 0x2f, 0x36, 0x70, 0x21, 0x36, 0xe3, 0xfa, 0x5f, 0xe7, 0x7e, 0x66, 0x8d,
	0x71, 0xd5, 0xe2, 0x5f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x08, 0x29, 0x16, 0x83, 0x04,
	0x00, 0x00,
}
