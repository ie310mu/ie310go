// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package envoy_api_v2

import proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/duration"
import wrappers "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import discovery "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/discovery"
import endpoint "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/endpoint/endpoint"
import percent "google.golang.org/grpc/balancer/xds/internal/proto/envoy/type/percent"
import _ "google.golang.org/grpc/balancer/xds/internal/proto/validate"

import (
	context "github.com/ie310mu/ie310go/forks/golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClusterLoadAssignment struct {
	ClusterName          string                          `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints            []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	NamedEndpoints       map[string]*endpoint.Endpoint   `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy               *ClusterLoadAssignment_Policy   `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_eds_1a80fb8e78974562, []int{0}
}
func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (dst *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(dst, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*endpoint.Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverloads          []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	OverprovisioningFactor *wrappers.UInt32Value                        `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	EndpointStaleAfter     *duration.Duration                           `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                                     `json:"-"`
	XXX_unrecognized       []byte                                       `json:"-"`
	XXX_sizecache          int32                                        `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eds_1a80fb8e78974562, []int{0, 1}
}
func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (dst *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(dst, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *duration.Duration {
	if m != nil {
		return m.EndpointStaleAfter
	}
	return nil
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	Category             string                     `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	DropPercentage       *percent.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_eds_1a80fb8e78974562, []int{0, 1, 0}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (dst *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(dst, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *percent.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*endpoint.Endpoint)(nil), "envoy.api.v2.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *discovery.DiscoveryRequest, opts ...grpc.CallOption) (*discovery.DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*discovery.DiscoveryRequest) error
	Recv() (*discovery.DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *discovery.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*discovery.DiscoveryResponse, error) {
	m := new(discovery.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *discovery.DiscoveryRequest, opts ...grpc.CallOption) (*discovery.DiscoveryResponse, error) {
	out := new(discovery.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	FetchEndpoints(context.Context, *discovery.DiscoveryRequest) (*discovery.DiscoveryResponse, error)
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*discovery.DiscoveryResponse) error
	Recv() (*discovery.DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *discovery.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*discovery.DiscoveryRequest, error) {
	m := new(discovery.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discovery.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*discovery.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor_eds_1a80fb8e78974562) }

var fileDescriptor_eds_1a80fb8e78974562 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6e, 0x13, 0x31,
	0x14, 0xc6, 0xeb, 0x49, 0x1a, 0x52, 0x37, 0xb4, 0x95, 0x81, 0x76, 0x18, 0x85, 0x36, 0x8a, 0x40,
	0xaa, 0x02, 0x9a, 0xa0, 0x54, 0xa8, 0xa8, 0xbb, 0x86, 0x36, 0x02, 0x54, 0x41, 0x34, 0x15, 0x08,
	0x36, 0x04, 0x67, 0xc6, 0x0d, 0x16, 0x13, 0xdb, 0xd8, 0xce, 0xc0, 0x2c, 0xd8, 0xb0, 0x62, 0xcf,
	0x2d, 0x7a, 0x04, 0x56, 0xac, 0xb8, 0x00, 0x57, 0x60, 0x83, 0xb8, 0x04, 0xf2, 0xfc, 0x6b, 0x87,
	0xb6, 0x12, 0x0b, 0x76, 0x9e, 0x79, 0xef, 0xfd, 0xfc, 0xf9, 0x7b, 0xcf, 0x86, 0xab, 0x84, 0x45,
	0x3c, 0xee, 0x62, 0x41, 0xbb, 0x51, 0xaf, 0x4b, 0x02, 0xe5, 0x0a, 0xc9, 0x35, 0x47, 0x8d, 0xe4,
	0xbf, 0x8b, 0x05, 0x75, 0xa3, 0x9e, 0xd3, 0x2c, 0x65, 0x05, 0x54, 0xf9, 0x3c, 0x22, 0x32, 0x4e,
	0x73, 0x9d, 0x9b, 0x65, 0x06, 0x0b, 0x04, 0xa7, 0x4c, 0x17, 0x8b, 0x2c, 0xcb, 0x4e, 0xb3, 0x74,
	0x2c, 0x48, 0x57, 0x10, 0xe9, 0x93, 0x22, 0xd2, 0x9c, 0x70, 0x3e, 0x09, 0x49, 0x02, 0xc0, 0x8c,
	0x71, 0x8d, 0x35, 0xe5, 0x2c, 0x53, 0xe2, 0xac, 0x45, 0x38, 0xa4, 0x01, 0xd6, 0xa4, 0x9b, 0x2f,
	0xb2, 0xc0, 0x7a, 0x56, 0x96, 0x7c, 0x8d, 0x67, 0x47, 0xdd, 0xf7, 0x12, 0x0b, 0x41, 0xa4, 0xba,
	0x28, 0x1e, 0xcc, 0x64, 0x42, 0x4e, 0xe3, 0xed, 0xef, 0x35, 0x78, 0xed, 0x41, 0x38, 0x53, 0x9a,
	0xc8, 0x03, 0x8e, 0x83, 0x5d, 0xa5, 0xe8, 0x84, 0x4d, 0x09, 0xd3, 0xe8, 0x0e, 0x6c, 0xf8, 0x69,
	0x60, 0xc4, 0xf0, 0x94, 0xd8, 0xa0, 0x05, 0x36, 0x17, 0xfa, 0x0b, 0x5f, 0x7f, 0x7d, 0xab, 0x54,
	0xa5, 0xd5, 0x02, 0xde, 0x62, 0x16, 0x7e, 0x82, 0xa7, 0x04, 0x3d, 0x84, 0x0b, 0xf9, 0x51, 0x95,
	0x6d, 0xb5, 0x2a, 0x9b, 0x8b, 0xbd, 0x8e, 0x7b, 0xda, 0x3e, 0xb7, 0x70, 0xe2, 0x80, 0xfb, 0x38,
	0xa4, 0x3a, 0x3e, 0x18, 0xef, 0xe7, 0x15, 0xde, 0x49, 0x31, 0x7a, 0x0d, 0x97, 0xcd, 0x7e, 0xc1,
	0xe8, 0x84, 0x37, 0x9f, 0xf0, 0xb6, 0xcb, 0xbc, 0x73, 0x55, 0xbb, 0x46, 0x4c, 0x50, 0x70, 0xf7,
	0x99, 0x96, 0xb1, 0xb7, 0xc4, 0x4a, 0x3f, 0x51, 0x1f, 0xd6, 0x04, 0x0f, 0xa9, 0x1f, 0xdb, 0xd5,
	0x16, 0x38, 0x2b, 0xf4, 0x7c, 0xf0, 0x30, 0xa9, 0xf0, 0xb2, 0x4a, 0x67, 0x0c, 0xaf, 0x9c, 0xb3,
	0x15, 0x5a, 0x81, 0x95, 0xb7, 0x24, 0x4e, 0xbd, 0xf2, 0xcc, 0x12, 0xdd, 0x83, 0xf3, 0x11, 0x0e,
	0x67, 0xc4, 0xb6, 0x92, 0xbd, 0x36, 0x2e, 0x30, 0x25, 0xe7, 0x78, 0x69, 0xf6, 0x8e, 0x75, 0x1f,
	0x38, 0xc7, 0x15, 0x58, 0x4b, 0xb7, 0x45, 0xaf, 0xe0, 0x52, 0x20, 0xb9, 0x18, 0x99, 0x89, 0x0b,
	0x39, 0x0e, 0x72, 0x8f, 0xb7, 0xff, 0x5d, 0xba, 0xbb, 0x27, 0xb9, 0x78, 0x9a, 0xd5, 0x7b, 0x97,
	0x83, 0x53, 0x5f, 0xc6, 0xf4, 0x35, 0x83, 0x16, 0x92, 0x47, 0x54, 0x51, 0xce, 0x28, 0x9b, 0x8c,
	0x8e, 0xb0, 0xaf, 0xb9, 0xb4, 0x2b, 0x89, 0xee, 0xa6, 0x9b, 0x0e, 0x92, 0x9b, 0x0f, 0x92, 0xfb,
	0xec, 0x11, 0xd3, 0x5b, 0xbd, 0xe7, 0x46, 0x6d, 0x36, 0x15, 0x1d, 0xab, 0x35, 0xe7, 0xad, 0xfe,
	0xcd, 0x19, 0x24, 0x18, 0xf4, 0x12, 0x5e, 0xcd, 0x0f, 0x3b, 0x52, 0x1a, 0x87, 0x64, 0x84, 0x8f,
	0x34, 0x91, 0x59, 0x0b, 0xae, 0x9f, 0xc1, 0xef, 0x65, 0x73, 0xda, 0x6f, 0x18, 0xf6, 0xa5, 0x63,
	0x50, 0xed, 0x58, 0xf5, 0x39, 0x0f, 0xe5, 0x90, 0x43, 0xc3, 0xd8, 0x35, 0x08, 0xe7, 0x23, 0x6c,
	0x9c, 0x3e, 0x1b, 0xba, 0x05, 0xeb, 0x3e, 0xd6, 0x64, 0xc2, 0x65, 0x7c, 0x76, 0x6a, 0x8b, 0x10,
	0x1a, 0xc0, 0xe5, 0xc4, 0xd3, 0xec, 0x1e, 0xe2, 0x49, 0xde, 0xa3, 0x1b, 0x99, 0xa9, 0xe6, 0x96,
	0xba, 0x03, 0x89, 0x7d, 0xa3, 0x03, 0x87, 0xc3, 0x34, 0xcf, 0x4b, 0x3a, 0x31, 0x2c, 0x8a, 0x1e,
	0x57, 0xeb, 0x60, 0xc5, 0xea, 0xfd, 0x06, 0xd0, 0xce, 0x9b, 0xb8, 0x97, 0xbf, 0x0d, 0x87, 0x44,
	0x46, 0xd4, 0x27, 0xe8, 0x05, 0x5c, 0x3e, 0xd4, 0x92, 0xe0, 0xe9, 0xc9, 0x10, 0xae, 0x97, 0x3b,
	0x57, 0x94, 0x78, 0xe4, 0xdd, 0x8c, 0x28, 0xed, 0x6c, 0x5c, 0x18, 0x57, 0x82, 0x33, 0x45, 0xda,
	0x73, 0x9b, 0xe0, 0x2e, 0x40, 0x33, 0xb8, 0x34, 0x20, 0xda, 0x7f, 0xf3, 0x1f, 0xc1, 0xed, 0x4f,
	0x3f, 0x7e, 0x7e, 0xb1, 0x9a, 0xed, 0xb5, 0xd2, 0x33, 0xb7, 0x53, 0x5c, 0xc7, 0x1d, 0xd0, 0xe9,
	0xdf, 0x86, 0x0e, 0xe5, 0x29, 0x48, 0x48, 0xfe, 0x21, 0x2e, 0x31, 0xfb, 0xf5, 0xfd, 0x40, 0x0d,
	0x4d, 0x23, 0x87, 0xe0, 0x33, 0x00, 0xe3, 0x5a, 0xd2, 0xd4, 0xad, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x26, 0x36, 0x57, 0x86, 0x67, 0x05, 0x00, 0x00,
}
