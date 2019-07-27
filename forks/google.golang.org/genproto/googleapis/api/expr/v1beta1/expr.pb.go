// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/expr/v1beta1/expr.proto

package expr

import (
	fmt "fmt"
	math "math"

	proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
	_struct "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/struct"
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

// An expression together with source information as returned by the parser.
type ParsedExpr struct {
	// The parsed expression.
	Expr *Expr `protobuf:"bytes,2,opt,name=expr,proto3" json:"expr,omitempty"`
	// The source info derived from input that generated the parsed `expr`.
	SourceInfo *SourceInfo `protobuf:"bytes,3,opt,name=source_info,json=sourceInfo,proto3" json:"source_info,omitempty"`
	// The syntax version of the source, e.g. `cel1`.
	SyntaxVersion        string   `protobuf:"bytes,4,opt,name=syntax_version,json=syntaxVersion,proto3" json:"syntax_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParsedExpr) Reset()         { *m = ParsedExpr{} }
func (m *ParsedExpr) String() string { return proto.CompactTextString(m) }
func (*ParsedExpr) ProtoMessage()    {}
func (*ParsedExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{0}
}

func (m *ParsedExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParsedExpr.Unmarshal(m, b)
}
func (m *ParsedExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParsedExpr.Marshal(b, m, deterministic)
}
func (m *ParsedExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParsedExpr.Merge(m, src)
}
func (m *ParsedExpr) XXX_Size() int {
	return xxx_messageInfo_ParsedExpr.Size(m)
}
func (m *ParsedExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_ParsedExpr.DiscardUnknown(m)
}

var xxx_messageInfo_ParsedExpr proto.InternalMessageInfo

func (m *ParsedExpr) GetExpr() *Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (m *ParsedExpr) GetSourceInfo() *SourceInfo {
	if m != nil {
		return m.SourceInfo
	}
	return nil
}

func (m *ParsedExpr) GetSyntaxVersion() string {
	if m != nil {
		return m.SyntaxVersion
	}
	return ""
}

// An abstract representation of a common expression.
//
// Expressions are abstractly represented as a collection of identifiers,
// select statements, function calls, literals, and comprehensions. All
// operators with the exception of the '.' operator are modelled as function
// calls. This makes it easy to represent new operators into the existing AST.
//
// All references within expressions must resolve to a
// [Decl][google.api.expr.v1beta1.Decl] provided at type-check for an expression
// to be valid. A reference may either be a bare identifier `name` or a
// qualified identifier `google.api.name`. References may either refer to a
// value or a function declaration.
//
// For example, the expression `google.api.name.startsWith('expr')` references
// the declaration `google.api.name` within a
// [Expr.Select][google.api.expr.v1beta1.Expr.Select] expression, and the
// function declaration `startsWith`.
type Expr struct {
	// Required. An id assigned to this node by the parser which is unique in a
	// given expression tree. This is used to associate type information and other
	// attributes to a node in the parse tree.
	Id int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Variants of expressions.
	//
	// Types that are valid to be assigned to ExprKind:
	//	*Expr_LiteralExpr
	//	*Expr_IdentExpr
	//	*Expr_SelectExpr
	//	*Expr_CallExpr
	//	*Expr_ListExpr
	//	*Expr_StructExpr
	//	*Expr_ComprehensionExpr
	ExprKind             isExpr_ExprKind `protobuf_oneof:"expr_kind"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Expr) Reset()         { *m = Expr{} }
func (m *Expr) String() string { return proto.CompactTextString(m) }
func (*Expr) ProtoMessage()    {}
func (*Expr) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{1}
}

func (m *Expr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr.Unmarshal(m, b)
}
func (m *Expr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr.Marshal(b, m, deterministic)
}
func (m *Expr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr.Merge(m, src)
}
func (m *Expr) XXX_Size() int {
	return xxx_messageInfo_Expr.Size(m)
}
func (m *Expr) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr.DiscardUnknown(m)
}

var xxx_messageInfo_Expr proto.InternalMessageInfo

func (m *Expr) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isExpr_ExprKind interface {
	isExpr_ExprKind()
}

type Expr_LiteralExpr struct {
	LiteralExpr *Literal `protobuf:"bytes,3,opt,name=literal_expr,json=literalExpr,proto3,oneof"`
}

type Expr_IdentExpr struct {
	IdentExpr *Expr_Ident `protobuf:"bytes,4,opt,name=ident_expr,json=identExpr,proto3,oneof"`
}

type Expr_SelectExpr struct {
	SelectExpr *Expr_Select `protobuf:"bytes,5,opt,name=select_expr,json=selectExpr,proto3,oneof"`
}

type Expr_CallExpr struct {
	CallExpr *Expr_Call `protobuf:"bytes,6,opt,name=call_expr,json=callExpr,proto3,oneof"`
}

type Expr_ListExpr struct {
	ListExpr *Expr_CreateList `protobuf:"bytes,7,opt,name=list_expr,json=listExpr,proto3,oneof"`
}

type Expr_StructExpr struct {
	StructExpr *Expr_CreateStruct `protobuf:"bytes,8,opt,name=struct_expr,json=structExpr,proto3,oneof"`
}

type Expr_ComprehensionExpr struct {
	ComprehensionExpr *Expr_Comprehension `protobuf:"bytes,9,opt,name=comprehension_expr,json=comprehensionExpr,proto3,oneof"`
}

func (*Expr_LiteralExpr) isExpr_ExprKind() {}

func (*Expr_IdentExpr) isExpr_ExprKind() {}

func (*Expr_SelectExpr) isExpr_ExprKind() {}

func (*Expr_CallExpr) isExpr_ExprKind() {}

func (*Expr_ListExpr) isExpr_ExprKind() {}

func (*Expr_StructExpr) isExpr_ExprKind() {}

func (*Expr_ComprehensionExpr) isExpr_ExprKind() {}

func (m *Expr) GetExprKind() isExpr_ExprKind {
	if m != nil {
		return m.ExprKind
	}
	return nil
}

func (m *Expr) GetLiteralExpr() *Literal {
	if x, ok := m.GetExprKind().(*Expr_LiteralExpr); ok {
		return x.LiteralExpr
	}
	return nil
}

func (m *Expr) GetIdentExpr() *Expr_Ident {
	if x, ok := m.GetExprKind().(*Expr_IdentExpr); ok {
		return x.IdentExpr
	}
	return nil
}

func (m *Expr) GetSelectExpr() *Expr_Select {
	if x, ok := m.GetExprKind().(*Expr_SelectExpr); ok {
		return x.SelectExpr
	}
	return nil
}

func (m *Expr) GetCallExpr() *Expr_Call {
	if x, ok := m.GetExprKind().(*Expr_CallExpr); ok {
		return x.CallExpr
	}
	return nil
}

func (m *Expr) GetListExpr() *Expr_CreateList {
	if x, ok := m.GetExprKind().(*Expr_ListExpr); ok {
		return x.ListExpr
	}
	return nil
}

func (m *Expr) GetStructExpr() *Expr_CreateStruct {
	if x, ok := m.GetExprKind().(*Expr_StructExpr); ok {
		return x.StructExpr
	}
	return nil
}

func (m *Expr) GetComprehensionExpr() *Expr_Comprehension {
	if x, ok := m.GetExprKind().(*Expr_ComprehensionExpr); ok {
		return x.ComprehensionExpr
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Expr) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Expr_LiteralExpr)(nil),
		(*Expr_IdentExpr)(nil),
		(*Expr_SelectExpr)(nil),
		(*Expr_CallExpr)(nil),
		(*Expr_ListExpr)(nil),
		(*Expr_StructExpr)(nil),
		(*Expr_ComprehensionExpr)(nil),
	}
}

// An identifier expression. e.g. `request`.
type Expr_Ident struct {
	// Required. Holds a single, unqualified identifier, possibly preceded by a
	// '.'.
	//
	// Qualified names are represented by the
	// [Expr.Select][google.api.expr.v1beta1.Expr.Select] expression.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_Ident) Reset()         { *m = Expr_Ident{} }
func (m *Expr_Ident) String() string { return proto.CompactTextString(m) }
func (*Expr_Ident) ProtoMessage()    {}
func (*Expr_Ident) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{1, 0}
}

func (m *Expr_Ident) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Ident.Unmarshal(m, b)
}
func (m *Expr_Ident) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Ident.Marshal(b, m, deterministic)
}
func (m *Expr_Ident) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Ident.Merge(m, src)
}
func (m *Expr_Ident) XXX_Size() int {
	return xxx_messageInfo_Expr_Ident.Size(m)
}
func (m *Expr_Ident) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Ident.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Ident proto.InternalMessageInfo

func (m *Expr_Ident) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A field selection expression. e.g. `request.auth`.
type Expr_Select struct {
	// Required. The target of the selection expression.
	//
	// For example, in the select expression `request.auth`, the `request`
	// portion of the expression is the `operand`.
	Operand *Expr `protobuf:"bytes,1,opt,name=operand,proto3" json:"operand,omitempty"`
	// Required. The name of the field to select.
	//
	// For example, in the select expression `request.auth`, the `auth` portion
	// of the expression would be the `field`.
	Field string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	// Whether the select is to be interpreted as a field presence test.
	//
	// This results from the macro `has(request.auth)`.
	TestOnly             bool     `protobuf:"varint,3,opt,name=test_only,json=testOnly,proto3" json:"test_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_Select) Reset()         { *m = Expr_Select{} }
func (m *Expr_Select) String() string { return proto.CompactTextString(m) }
func (*Expr_Select) ProtoMessage()    {}
func (*Expr_Select) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{1, 1}
}

func (m *Expr_Select) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Select.Unmarshal(m, b)
}
func (m *Expr_Select) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Select.Marshal(b, m, deterministic)
}
func (m *Expr_Select) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Select.Merge(m, src)
}
func (m *Expr_Select) XXX_Size() int {
	return xxx_messageInfo_Expr_Select.Size(m)
}
func (m *Expr_Select) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Select.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Select proto.InternalMessageInfo

func (m *Expr_Select) GetOperand() *Expr {
	if m != nil {
		return m.Operand
	}
	return nil
}

func (m *Expr_Select) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Expr_Select) GetTestOnly() bool {
	if m != nil {
		return m.TestOnly
	}
	return false
}

// A call expression, including calls to predefined functions and operators.
//
// For example, `value == 10`, `size(map_value)`.
type Expr_Call struct {
	// The target of an method call-style expression. For example, `x` in
	// `x.f()`.
	Target *Expr `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Required. The name of the function or method being called.
	Function string `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
	// The arguments.
	Args                 []*Expr  `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_Call) Reset()         { *m = Expr_Call{} }
func (m *Expr_Call) String() string { return proto.CompactTextString(m) }
func (*Expr_Call) ProtoMessage()    {}
func (*Expr_Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{1, 2}
}

func (m *Expr_Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Call.Unmarshal(m, b)
}
func (m *Expr_Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Call.Marshal(b, m, deterministic)
}
func (m *Expr_Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Call.Merge(m, src)
}
func (m *Expr_Call) XXX_Size() int {
	return xxx_messageInfo_Expr_Call.Size(m)
}
func (m *Expr_Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Call proto.InternalMessageInfo

func (m *Expr_Call) GetTarget() *Expr {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Expr_Call) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Expr_Call) GetArgs() []*Expr {
	if m != nil {
		return m.Args
	}
	return nil
}

// A list creation expression.
//
// Lists may either be homogenous, e.g. `[1, 2, 3]`, or heterogenous, e.g.
// `dyn([1, 'hello', 2.0])`
type Expr_CreateList struct {
	// The elements part of the list.
	Elements             []*Expr  `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_CreateList) Reset()         { *m = Expr_CreateList{} }
func (m *Expr_CreateList) String() string { return proto.CompactTextString(m) }
func (*Expr_CreateList) ProtoMessage()    {}
func (*Expr_CreateList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{1, 3}
}

func (m *Expr_CreateList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_CreateList.Unmarshal(m, b)
}
func (m *Expr_CreateList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_CreateList.Marshal(b, m, deterministic)
}
func (m *Expr_CreateList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_CreateList.Merge(m, src)
}
func (m *Expr_CreateList) XXX_Size() int {
	return xxx_messageInfo_Expr_CreateList.Size(m)
}
func (m *Expr_CreateList) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_CreateList.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_CreateList proto.InternalMessageInfo

func (m *Expr_CreateList) GetElements() []*Expr {
	if m != nil {
		return m.Elements
	}
	return nil
}

// A map or message creation expression.
//
// Maps are constructed as `{'key_name': 'value'}`. Message construction is
// similar, but prefixed with a type name and composed of field ids:
// `types.MyType{field_id: 'value'}`.
type Expr_CreateStruct struct {
	// The type name of the message to be created, empty when creating map
	// literals.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The entries in the creation expression.
	Entries              []*Expr_CreateStruct_Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Expr_CreateStruct) Reset()         { *m = Expr_CreateStruct{} }
func (m *Expr_CreateStruct) String() string { return proto.CompactTextString(m) }
func (*Expr_CreateStruct) ProtoMessage()    {}
func (*Expr_CreateStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{1, 4}
}

func (m *Expr_CreateStruct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_CreateStruct.Unmarshal(m, b)
}
func (m *Expr_CreateStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_CreateStruct.Marshal(b, m, deterministic)
}
func (m *Expr_CreateStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_CreateStruct.Merge(m, src)
}
func (m *Expr_CreateStruct) XXX_Size() int {
	return xxx_messageInfo_Expr_CreateStruct.Size(m)
}
func (m *Expr_CreateStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_CreateStruct.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_CreateStruct proto.InternalMessageInfo

func (m *Expr_CreateStruct) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Expr_CreateStruct) GetEntries() []*Expr_CreateStruct_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Represents an entry.
type Expr_CreateStruct_Entry struct {
	// Required. An id assigned to this node by the parser which is unique
	// in a given expression tree. This is used to associate type
	// information and other attributes to the node.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The `Entry` key kinds.
	//
	// Types that are valid to be assigned to KeyKind:
	//	*Expr_CreateStruct_Entry_FieldKey
	//	*Expr_CreateStruct_Entry_MapKey
	KeyKind isExpr_CreateStruct_Entry_KeyKind `protobuf_oneof:"key_kind"`
	// Required. The value assigned to the key.
	Value                *Expr    `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_CreateStruct_Entry) Reset()         { *m = Expr_CreateStruct_Entry{} }
func (m *Expr_CreateStruct_Entry) String() string { return proto.CompactTextString(m) }
func (*Expr_CreateStruct_Entry) ProtoMessage()    {}
func (*Expr_CreateStruct_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{1, 4, 0}
}

func (m *Expr_CreateStruct_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_CreateStruct_Entry.Unmarshal(m, b)
}
func (m *Expr_CreateStruct_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_CreateStruct_Entry.Marshal(b, m, deterministic)
}
func (m *Expr_CreateStruct_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_CreateStruct_Entry.Merge(m, src)
}
func (m *Expr_CreateStruct_Entry) XXX_Size() int {
	return xxx_messageInfo_Expr_CreateStruct_Entry.Size(m)
}
func (m *Expr_CreateStruct_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_CreateStruct_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_CreateStruct_Entry proto.InternalMessageInfo

func (m *Expr_CreateStruct_Entry) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isExpr_CreateStruct_Entry_KeyKind interface {
	isExpr_CreateStruct_Entry_KeyKind()
}

type Expr_CreateStruct_Entry_FieldKey struct {
	FieldKey string `protobuf:"bytes,2,opt,name=field_key,json=fieldKey,proto3,oneof"`
}

type Expr_CreateStruct_Entry_MapKey struct {
	MapKey *Expr `protobuf:"bytes,3,opt,name=map_key,json=mapKey,proto3,oneof"`
}

func (*Expr_CreateStruct_Entry_FieldKey) isExpr_CreateStruct_Entry_KeyKind() {}

func (*Expr_CreateStruct_Entry_MapKey) isExpr_CreateStruct_Entry_KeyKind() {}

func (m *Expr_CreateStruct_Entry) GetKeyKind() isExpr_CreateStruct_Entry_KeyKind {
	if m != nil {
		return m.KeyKind
	}
	return nil
}

func (m *Expr_CreateStruct_Entry) GetFieldKey() string {
	if x, ok := m.GetKeyKind().(*Expr_CreateStruct_Entry_FieldKey); ok {
		return x.FieldKey
	}
	return ""
}

func (m *Expr_CreateStruct_Entry) GetMapKey() *Expr {
	if x, ok := m.GetKeyKind().(*Expr_CreateStruct_Entry_MapKey); ok {
		return x.MapKey
	}
	return nil
}

func (m *Expr_CreateStruct_Entry) GetValue() *Expr {
	if m != nil {
		return m.Value
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Expr_CreateStruct_Entry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Expr_CreateStruct_Entry_FieldKey)(nil),
		(*Expr_CreateStruct_Entry_MapKey)(nil),
	}
}

// A comprehension expression applied to a list or map.
//
// Comprehensions are not part of the core syntax, but enabled with macros.
// A macro matches a specific call signature within a parsed AST and replaces
// the call with an alternate AST block. Macro expansion happens at parse
// time.
//
// The following macros are supported within CEL:
//
// Aggregate type macros may be applied to all elements in a list or all keys
// in a map:
//
// *  `all`, `exists`, `exists_one` -  test a predicate expression against
//    the inputs and return `true` if the predicate is satisfied for all,
//    any, or only one value `list.all(x, x < 10)`.
// *  `filter` - test a predicate expression against the inputs and return
//    the subset of elements which satisfy the predicate:
//    `payments.filter(p, p > 1000)`.
// *  `map` - apply an expression to all elements in the input and return the
//    output aggregate type: `[1, 2, 3].map(i, i * i)`.
//
// The `has(m.x)` macro tests whether the property `x` is present in struct
// `m`. The semantics of this macro depend on the type of `m`. For proto2
// messages `has(m.x)` is defined as 'defined, but not set`. For proto3, the
// macro tests whether the property is set to its default. For map and struct
// types, the macro tests whether the property `x` is defined on `m`.
type Expr_Comprehension struct {
	// The name of the iteration variable.
	IterVar string `protobuf:"bytes,1,opt,name=iter_var,json=iterVar,proto3" json:"iter_var,omitempty"`
	// The range over which var iterates.
	IterRange *Expr `protobuf:"bytes,2,opt,name=iter_range,json=iterRange,proto3" json:"iter_range,omitempty"`
	// The name of the variable used for accumulation of the result.
	AccuVar string `protobuf:"bytes,3,opt,name=accu_var,json=accuVar,proto3" json:"accu_var,omitempty"`
	// The initial value of the accumulator.
	AccuInit *Expr `protobuf:"bytes,4,opt,name=accu_init,json=accuInit,proto3" json:"accu_init,omitempty"`
	// An expression which can contain iter_var and accu_var.
	//
	// Returns false when the result has been computed and may be used as
	// a hint to short-circuit the remainder of the comprehension.
	LoopCondition *Expr `protobuf:"bytes,5,opt,name=loop_condition,json=loopCondition,proto3" json:"loop_condition,omitempty"`
	// An expression which can contain iter_var and accu_var.
	//
	// Computes the next value of accu_var.
	LoopStep *Expr `protobuf:"bytes,6,opt,name=loop_step,json=loopStep,proto3" json:"loop_step,omitempty"`
	// An expression which can contain accu_var.
	//
	// Computes the result.
	Result               *Expr    `protobuf:"bytes,7,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expr_Comprehension) Reset()         { *m = Expr_Comprehension{} }
func (m *Expr_Comprehension) String() string { return proto.CompactTextString(m) }
func (*Expr_Comprehension) ProtoMessage()    {}
func (*Expr_Comprehension) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{1, 5}
}

func (m *Expr_Comprehension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expr_Comprehension.Unmarshal(m, b)
}
func (m *Expr_Comprehension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expr_Comprehension.Marshal(b, m, deterministic)
}
func (m *Expr_Comprehension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expr_Comprehension.Merge(m, src)
}
func (m *Expr_Comprehension) XXX_Size() int {
	return xxx_messageInfo_Expr_Comprehension.Size(m)
}
func (m *Expr_Comprehension) XXX_DiscardUnknown() {
	xxx_messageInfo_Expr_Comprehension.DiscardUnknown(m)
}

var xxx_messageInfo_Expr_Comprehension proto.InternalMessageInfo

func (m *Expr_Comprehension) GetIterVar() string {
	if m != nil {
		return m.IterVar
	}
	return ""
}

func (m *Expr_Comprehension) GetIterRange() *Expr {
	if m != nil {
		return m.IterRange
	}
	return nil
}

func (m *Expr_Comprehension) GetAccuVar() string {
	if m != nil {
		return m.AccuVar
	}
	return ""
}

func (m *Expr_Comprehension) GetAccuInit() *Expr {
	if m != nil {
		return m.AccuInit
	}
	return nil
}

func (m *Expr_Comprehension) GetLoopCondition() *Expr {
	if m != nil {
		return m.LoopCondition
	}
	return nil
}

func (m *Expr_Comprehension) GetLoopStep() *Expr {
	if m != nil {
		return m.LoopStep
	}
	return nil
}

func (m *Expr_Comprehension) GetResult() *Expr {
	if m != nil {
		return m.Result
	}
	return nil
}

// Represents a primitive literal.
//
// This is similar to the primitives supported in the well-known type
// `google.protobuf.Value`, but richer so it can represent CEL's full range of
// primitives.
//
// Lists and structs are not included as constants as these aggregate types may
// contain [Expr][google.api.expr.v1beta1.Expr] elements which require
// evaluation and are thus not constant.
//
// Examples of literals include: `"hello"`, `b'bytes'`, `1u`, `4.2`, `-2`,
// `true`, `null`.
type Literal struct {
	// Required. The valid constant kinds.
	//
	// Types that are valid to be assigned to ConstantKind:
	//	*Literal_NullValue
	//	*Literal_BoolValue
	//	*Literal_Int64Value
	//	*Literal_Uint64Value
	//	*Literal_DoubleValue
	//	*Literal_StringValue
	//	*Literal_BytesValue
	ConstantKind         isLiteral_ConstantKind `protobuf_oneof:"constant_kind"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Literal) Reset()         { *m = Literal{} }
func (m *Literal) String() string { return proto.CompactTextString(m) }
func (*Literal) ProtoMessage()    {}
func (*Literal) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb180304f2777248, []int{2}
}

func (m *Literal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Literal.Unmarshal(m, b)
}
func (m *Literal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Literal.Marshal(b, m, deterministic)
}
func (m *Literal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Literal.Merge(m, src)
}
func (m *Literal) XXX_Size() int {
	return xxx_messageInfo_Literal.Size(m)
}
func (m *Literal) XXX_DiscardUnknown() {
	xxx_messageInfo_Literal.DiscardUnknown(m)
}

var xxx_messageInfo_Literal proto.InternalMessageInfo

type isLiteral_ConstantKind interface {
	isLiteral_ConstantKind()
}

type Literal_NullValue struct {
	NullValue _struct.NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Literal_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Literal_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Literal_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,4,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type Literal_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Literal_StringValue struct {
	StringValue string `protobuf:"bytes,6,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Literal_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,7,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

func (*Literal_NullValue) isLiteral_ConstantKind() {}

func (*Literal_BoolValue) isLiteral_ConstantKind() {}

func (*Literal_Int64Value) isLiteral_ConstantKind() {}

func (*Literal_Uint64Value) isLiteral_ConstantKind() {}

func (*Literal_DoubleValue) isLiteral_ConstantKind() {}

func (*Literal_StringValue) isLiteral_ConstantKind() {}

func (*Literal_BytesValue) isLiteral_ConstantKind() {}

func (m *Literal) GetConstantKind() isLiteral_ConstantKind {
	if m != nil {
		return m.ConstantKind
	}
	return nil
}

func (m *Literal) GetNullValue() _struct.NullValue {
	if x, ok := m.GetConstantKind().(*Literal_NullValue); ok {
		return x.NullValue
	}
	return _struct.NullValue_NULL_VALUE
}

func (m *Literal) GetBoolValue() bool {
	if x, ok := m.GetConstantKind().(*Literal_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Literal) GetInt64Value() int64 {
	if x, ok := m.GetConstantKind().(*Literal_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Literal) GetUint64Value() uint64 {
	if x, ok := m.GetConstantKind().(*Literal_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *Literal) GetDoubleValue() float64 {
	if x, ok := m.GetConstantKind().(*Literal_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Literal) GetStringValue() string {
	if x, ok := m.GetConstantKind().(*Literal_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Literal) GetBytesValue() []byte {
	if x, ok := m.GetConstantKind().(*Literal_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Literal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Literal_NullValue)(nil),
		(*Literal_BoolValue)(nil),
		(*Literal_Int64Value)(nil),
		(*Literal_Uint64Value)(nil),
		(*Literal_DoubleValue)(nil),
		(*Literal_StringValue)(nil),
		(*Literal_BytesValue)(nil),
	}
}

func init() {
	proto.RegisterType((*ParsedExpr)(nil), "google.api.expr.v1beta1.ParsedExpr")
	proto.RegisterType((*Expr)(nil), "google.api.expr.v1beta1.Expr")
	proto.RegisterType((*Expr_Ident)(nil), "google.api.expr.v1beta1.Expr.Ident")
	proto.RegisterType((*Expr_Select)(nil), "google.api.expr.v1beta1.Expr.Select")
	proto.RegisterType((*Expr_Call)(nil), "google.api.expr.v1beta1.Expr.Call")
	proto.RegisterType((*Expr_CreateList)(nil), "google.api.expr.v1beta1.Expr.CreateList")
	proto.RegisterType((*Expr_CreateStruct)(nil), "google.api.expr.v1beta1.Expr.CreateStruct")
	proto.RegisterType((*Expr_CreateStruct_Entry)(nil), "google.api.expr.v1beta1.Expr.CreateStruct.Entry")
	proto.RegisterType((*Expr_Comprehension)(nil), "google.api.expr.v1beta1.Expr.Comprehension")
	proto.RegisterType((*Literal)(nil), "google.api.expr.v1beta1.Literal")
}

func init() { proto.RegisterFile("google/api/expr/v1beta1/expr.proto", fileDescriptor_fb180304f2777248) }

var fileDescriptor_fb180304f2777248 = []byte{
	// 936 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0x6f, 0x6f, 0x23, 0xb5,
	0x13, 0xc7, 0xbb, 0xf9, 0xbb, 0x3b, 0x69, 0xfb, 0xd3, 0xcf, 0x42, 0x22, 0x6c, 0x39, 0x51, 0x7a,
	0x87, 0x54, 0x81, 0x94, 0xd0, 0x3b, 0xfe, 0x97, 0x27, 0xf4, 0xae, 0xba, 0x14, 0x0e, 0xa8, 0xb6,
	0x52, 0x1f, 0x20, 0xa4, 0xc8, 0xd9, 0x38, 0x8b, 0xa9, 0x63, 0xaf, 0xbc, 0xde, 0xaa, 0x79, 0x0f,
	0xbc, 0x0a, 0x1e, 0xc3, 0x03, 0x5e, 0x00, 0xef, 0xeb, 0x1e, 0xa2, 0x19, 0x3b, 0xa1, 0x15, 0xca,
	0x25, 0xcf, 0xd6, 0xe3, 0xcf, 0x7c, 0x3d, 0x9e, 0x19, 0x4f, 0x02, 0x47, 0x85, 0x31, 0x85, 0x12,
	0x43, 0x5e, 0xca, 0xa1, 0xb8, 0x2b, 0xed, 0xf0, 0xf6, 0x64, 0x22, 0x1c, 0x3f, 0xa1, 0xc5, 0xa0,
	0xb4, 0xc6, 0x19, 0xf6, 0xb6, 0x67, 0x06, 0xbc, 0x94, 0x03, 0x32, 0x07, 0x26, 0x7d, 0xb2, 0xce,
	0xb9, 0x32, 0xb5, 0xcd, 0x85, 0x77, 0x4f, 0xdf, 0x0d, 0x14, 0xad, 0x26, 0xf5, 0x6c, 0x58, 0x39,
	0x5b, 0xe7, 0xce, 0xef, 0x1e, 0xfd, 0x11, 0x01, 0x5c, 0x72, 0x5b, 0x89, 0xe9, 0xf9, 0x5d, 0x69,
	0xd9, 0x09, 0xb4, 0x50, 0xa9, 0xdf, 0x38, 0x8c, 0x8e, 0x7b, 0x4f, 0x1f, 0x0d, 0xd6, 0x1c, 0x3d,
	0x40, 0x38, 0x23, 0x94, 0xbd, 0x80, 0x9e, 0x3f, 0x6f, 0x2c, 0xf5, 0xcc, 0xf4, 0x9b, 0xe4, 0xf9,
	0x78, 0xad, 0xe7, 0x15, 0xb1, 0x17, 0x7a, 0x66, 0x32, 0xa8, 0x56, 0xdf, 0xec, 0x03, 0xd8, 0xaf,
	0x16, 0xda, 0xf1, 0xbb, 0xf1, 0xad, 0xb0, 0x95, 0x34, 0xba, 0xdf, 0x3a, 0x8c, 0x8e, 0x93, 0x6c,
	0xcf, 0x5b, 0xaf, 0xbd, 0xf1, 0xe8, 0xef, 0x5d, 0x68, 0x51, 0xa0, 0xfb, 0xd0, 0x90, 0x53, 0x0a,
	0xb3, 0x9d, 0x35, 0xe4, 0x94, 0x9d, 0xc3, 0xae, 0x92, 0x4e, 0x58, 0xae, 0xc6, 0x74, 0x01, 0x1f,
	0xc6, 0xe1, 0xda, 0x30, 0x5e, 0x79, 0x78, 0xb4, 0x93, 0xf5, 0x82, 0xdf, 0xb9, 0xbf, 0x0c, 0xc8,
	0xa9, 0xd0, 0xce, 0x8b, 0xb4, 0x36, 0xdc, 0x05, 0x5d, 0x06, 0x17, 0xc8, 0x8f, 0x76, 0xb2, 0x84,
	0x1c, 0x49, 0xe5, 0x25, 0xf4, 0x2a, 0xa1, 0x44, 0x1e, 0x64, 0xda, 0x24, 0xf3, 0xe4, 0xcd, 0x32,
	0x57, 0xe4, 0x30, 0xda, 0xc9, 0xc0, 0xbb, 0x92, 0xd0, 0x37, 0x90, 0xe4, 0x5c, 0x85, 0x2b, 0x75,
	0x48, 0xe6, 0xe8, 0xcd, 0x32, 0xcf, 0xb9, 0xc2, 0x4b, 0xc5, 0xe8, 0x16, 0x62, 0x49, 0x94, 0xac,
	0x42, 0x24, 0x5d, 0x92, 0x38, 0xde, 0x20, 0x61, 0x05, 0x77, 0xe2, 0x95, 0xac, 0x30, 0x9a, 0x18,
	0x9d, 0x49, 0xe8, 0x7b, 0xe8, 0xf9, 0xce, 0xf1, 0x52, 0x31, 0x49, 0x7d, 0xb8, 0x8d, 0xd4, 0x15,
	0xb9, 0xd1, 0xd5, 0xe8, 0x8b, 0xe4, 0x7e, 0x06, 0x96, 0x9b, 0x79, 0x69, 0xc5, 0x2f, 0x42, 0x63,
	0x69, 0xbd, 0x6a, 0x42, 0xaa, 0x1f, 0x6d, 0x50, 0xbd, 0xef, 0x37, 0xda, 0xc9, 0xfe, 0xff, 0x40,
	0x08, 0x91, 0xf4, 0x00, 0xda, 0x54, 0x17, 0xc6, 0xa0, 0xa5, 0xf9, 0x5c, 0xf4, 0x23, 0xea, 0x26,
	0xfa, 0x4e, 0x1d, 0x74, 0x7c, 0xb6, 0xd9, 0xe7, 0xd0, 0x35, 0xa5, 0xb0, 0x5c, 0x4f, 0x09, 0xd8,
	0xd8, 0xf1, 0x4b, 0x9a, 0xbd, 0x05, 0xed, 0x99, 0x14, 0xca, 0x77, 0x60, 0x92, 0xf9, 0x05, 0x3b,
	0x80, 0xc4, 0x89, 0xca, 0x8d, 0x8d, 0x56, 0x0b, 0xea, 0xc0, 0x38, 0x8b, 0xd1, 0xf0, 0xa3, 0x56,
	0x8b, 0xf4, 0xb7, 0x08, 0x5a, 0x58, 0x1d, 0xf6, 0x29, 0x74, 0x1c, 0xb7, 0x85, 0x70, 0xdb, 0x9d,
	0x19, 0x60, 0x96, 0x42, 0x3c, 0xab, 0x75, 0xee, 0xf0, 0x6d, 0xf8, 0x53, 0x57, 0x6b, 0x7c, 0xb6,
	0xdc, 0x16, 0x55, 0xbf, 0x79, 0xd8, 0xdc, 0xe2, 0xd9, 0x22, 0x9a, 0xbe, 0x04, 0xf8, 0xb7, 0xd0,
	0xec, 0x4b, 0x88, 0x85, 0x12, 0x73, 0xa1, 0x5d, 0xd5, 0x8f, 0xb6, 0x11, 0x59, 0xe1, 0xe9, 0xef,
	0x0d, 0xd8, 0xbd, 0x5f, 0x67, 0x4c, 0xb9, 0x5b, 0x94, 0xab, 0x94, 0xe3, 0x37, 0xfb, 0x16, 0xba,
	0x42, 0x3b, 0x2b, 0x45, 0xd5, 0x6f, 0x90, 0xfc, 0xc7, 0xdb, 0x37, 0xce, 0xe0, 0x5c, 0x3b, 0xbb,
	0xc8, 0x96, 0x02, 0xe9, 0x5f, 0x11, 0xb4, 0xc9, 0x14, 0x86, 0x40, 0xb4, 0x1a, 0x02, 0x8f, 0x20,
	0xa1, 0x42, 0x8c, 0x6f, 0xc4, 0xc2, 0xe7, 0x08, 0x3b, 0x98, 0x4c, 0xdf, 0x89, 0x05, 0xfb, 0x02,
	0xba, 0x73, 0x5e, 0xd2, 0x66, 0x73, 0x8b, 0xcc, 0x8f, 0x76, 0xb2, 0xce, 0x9c, 0x97, 0xe8, 0xf9,
	0x0c, 0xda, 0xb7, 0x5c, 0xd5, 0x22, 0x4c, 0x84, 0x0d, 0xb9, 0xf1, 0xec, 0x19, 0x40, 0x7c, 0x23,
	0x16, 0xe3, 0x1b, 0xa9, 0xa7, 0xe9, 0xeb, 0x06, 0xec, 0x3d, 0x68, 0x5b, 0xf6, 0x0e, 0xc4, 0x38,
	0x77, 0xc6, 0xb7, 0xdc, 0x86, 0x4c, 0x75, 0x71, 0x7d, 0xcd, 0x2d, 0xfb, 0x1a, 0x80, 0xb6, 0x2c,
	0xd7, 0x85, 0xd8, 0x6e, 0x14, 0x27, 0xe8, 0x90, 0x21, 0x8f, 0xc2, 0x3c, 0xcf, 0x6b, 0x12, 0x6e,
	0x7a, 0x61, 0x5c, 0xa3, 0xf0, 0x57, 0x90, 0xd0, 0x96, 0xd4, 0xd2, 0x6d, 0x77, 0x15, 0x92, 0xba,
	0xd0, 0xd2, 0xb1, 0x17, 0xb0, 0xaf, 0x8c, 0x29, 0xc7, 0xb9, 0xd1, 0x53, 0x49, 0x4d, 0xd8, 0xde,
	0x46, 0x60, 0x0f, 0x9d, 0x9e, 0x2f, 0x7d, 0x30, 0x02, 0x52, 0xa9, 0x9c, 0x28, 0xc3, 0x40, 0xdb,
	0x14, 0x01, 0xf2, 0x57, 0x4e, 0x94, 0xf8, 0x6e, 0xac, 0xa8, 0x6a, 0xe5, 0xc2, 0x18, 0xdb, 0xf4,
	0x6e, 0x3c, 0x7c, 0xd6, 0x83, 0x04, 0x37, 0xa9, 0x0e, 0x47, 0x7f, 0x36, 0xa0, 0x1b, 0x46, 0x3f,
	0x3b, 0x05, 0xd0, 0xb5, 0x52, 0x63, 0x5f, 0x59, 0xac, 0xc1, 0xfe, 0xd3, 0x74, 0xa9, 0xb9, 0xfc,
	0xb5, 0x1c, 0xfc, 0x50, 0x2b, 0x75, 0x8d, 0x04, 0x8e, 0x78, 0xbd, 0x5c, 0xb0, 0xf7, 0x00, 0x26,
	0xc6, 0x2c, 0x9d, 0xb1, 0x46, 0x31, 0x02, 0x68, 0xf3, 0xc0, 0xfb, 0xd0, 0x93, 0xda, 0x7d, 0xf6,
	0x49, 0x20, 0xb0, 0x12, 0x4d, 0x1c, 0x81, 0x64, 0xf4, 0xc8, 0x63, 0xd8, 0xad, 0xef, 0x33, 0x58,
	0x91, 0x16, 0xfe, 0x22, 0xd5, 0x0f, 0xa1, 0xa9, 0xa9, 0x27, 0x4a, 0x04, 0x08, 0xb3, 0x1e, 0x21,
	0xe4, 0xad, 0x2b, 0xa8, 0x72, 0x56, 0xea, 0x22, 0x40, 0x9d, 0xd0, 0xfb, 0x3d, 0x6f, 0x5d, 0x45,
	0x34, 0x59, 0x38, 0x51, 0x05, 0x06, 0x93, 0xb8, 0x8b, 0x11, 0x91, 0x91, 0x90, 0xb3, 0xff, 0xc1,
	0x5e, 0x6e, 0x74, 0xe5, 0xb8, 0x76, 0x94, 0xaf, 0xb3, 0x5f, 0xe1, 0x20, 0x37, 0xf3, 0x75, 0x89,
	0x3e, 0x4b, 0x30, 0xd3, 0x97, 0x98, 0xaa, 0xcb, 0xe8, 0xa7, 0xd3, 0x40, 0x15, 0x46, 0x71, 0x5d,
	0x0c, 0x8c, 0x2d, 0x86, 0x85, 0xd0, 0x94, 0xc8, 0xa1, 0xdf, 0xe2, 0xa5, 0xac, 0xfe, 0xf3, 0x6f,
	0xe5, 0x14, 0x17, 0xaf, 0xa3, 0x68, 0xd2, 0x21, 0xf4, 0xd9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xcb, 0xa2, 0x13, 0x9a, 0x14, 0x09, 0x00, 0x00,
}
