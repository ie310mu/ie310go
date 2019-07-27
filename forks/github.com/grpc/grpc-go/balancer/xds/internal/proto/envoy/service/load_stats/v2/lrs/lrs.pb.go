// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/load_stats/v2/lrs.proto

package v2

import proto "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/ie310mu/ie310go/forks/github.com/golang/protobuf/ptypes/duration"
import base "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/core/base"
import load_report "google.golang.org/grpc/balancer/xds/internal/proto/envoy/api/v2/endpoint/load_report"
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

type LoadStatsRequest struct {
	Node                 *base.Node                  `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	ClusterStats         []*load_report.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lrs_0f8f2c1d40a1b9f7, []int{0}
}
func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsRequest.Unmarshal(m, b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
}
func (dst *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(dst, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return xxx_messageInfo_LoadStatsRequest.Size(m)
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *base.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*load_report.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

type LoadStatsResponse struct {
	Clusters                  []string           `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	LoadReportingInterval     *duration.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	ReportEndpointGranularity bool               `protobuf:"varint,3,opt,name=report_endpoint_granularity,json=reportEndpointGranularity,proto3" json:"report_endpoint_granularity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lrs_0f8f2c1d40a1b9f7, []int{1}
}
func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsResponse.Unmarshal(m, b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
}
func (dst *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(dst, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return xxx_messageInfo_LoadStatsResponse.Size(m)
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func (m *LoadStatsResponse) GetReportEndpointGranularity() bool {
	if m != nil {
		return m.ReportEndpointGranularity
	}
	return false
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v2.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v2.LoadStatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v2.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v2.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v2/lrs.proto",
}

func init() {
	proto.RegisterFile("envoy/service/load_stats/v2/lrs.proto", fileDescriptor_lrs_0f8f2c1d40a1b9f7)
}

var fileDescriptor_lrs_0f8f2c1d40a1b9f7 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x71, 0x0b, 0xa8, 0x78, 0x40, 0x80, 0x05, 0x6a, 0xa6, 0x83, 0x50, 0x55, 0x04, 0x04,
	0x21, 0x6c, 0x14, 0xf6, 0xb3, 0x28, 0x20, 0x40, 0xaa, 0xd0, 0x90, 0xee, 0xd8, 0x54, 0x6e, 0xf2,
	0x88, 0x2c, 0x05, 0xbf, 0x60, 0x3b, 0x16, 0xbd, 0x01, 0x6c, 0x58, 0x70, 0x1c, 0x56, 0x9c, 0x80,
	0x7b, 0x70, 0x0b, 0x94, 0x38, 0x99, 0xc9, 0xb0, 0xa8, 0x66, 0x17, 0xeb, 0x7d, 0xff, 0xf3, 0xff,
	0xff, 0x31, 0x7d, 0x08, 0xda, 0xe3, 0x4e, 0x58, 0x30, 0x5e, 0x65, 0x20, 0x4a, 0x94, 0xf9, 0xc6,
	0x3a, 0xe9, 0xac, 0xf0, 0x89, 0x28, 0x8d, 0xe5, 0x95, 0x41, 0x87, 0xec, 0xa8, 0xc5, 0x78, 0x87,
	0xf1, 0x33, 0x8c, 0xfb, 0x64, 0x76, 0x2f, 0xec, 0x90, 0x95, 0x6a, 0x44, 0x19, 0x1a, 0x10, 0x5b,
	0x69, 0x21, 0x48, 0x67, 0x8f, 0xcf, 0x4d, 0x41, 0xe7, 0x15, 0x2a, 0xed, 0xc2, 0x4d, 0x06, 0x2a,
	0x34, 0xae, 0x03, 0xef, 0x17, 0x88, 0x45, 0x09, 0xa2, 0x3d, 0x6d, 0xeb, 0x4f, 0x22, 0xaf, 0x8d,
	0x74, 0x0a, 0x75, 0x37, 0x9f, 0x7a, 0x59, 0xaa, 0x5c, 0x3a, 0x10, 0xfd, 0x47, 0x18, 0x2c, 0xbe,
	0x13, 0x7a, 0x6b, 0x85, 0x32, 0x5f, 0x37, 0x86, 0x52, 0xf8, 0x52, 0x83, 0x75, 0xec, 0x29, 0xbd,
	0xac, 0x31, 0x87, 0x88, 0xcc, 0x49, 0x7c, 0x90, 0x4c, 0x79, 0x08, 0x20, 0x2b, 0xc5, 0x7d, 0xc2,
	0x1b, 0x8f, 0xfc, 0x3d, 0xe6, 0x90, 0xb6, 0x10, 0x7b, 0x4b, 0x6f, 0x64, 0x65, 0x6d, 0x1d, 0x98,
	0x90, 0x2a, 0x1a, 0xcd, 0xc7, 0xf1, 0x41, 0xf2, 0xe0, 0xbc, 0xaa, 0xf7, 0xce, 0x5f, 0x06, 0x36,
	0xdc, 0x77, 0x3d, 0x1b, 0x9c, 0x16, 0x7f, 0x08, 0xbd, 0x3d, 0xf0, 0x62, 0x2b, 0xd4, 0x16, 0xd8,
	0x23, 0x3a, 0xe9, 0x28, 0x1b, 0x91, 0xf9, 0x38, 0xbe, 0xb6, 0xa4, 0xbf, 0xfe, 0xfe, 0x1e, 0x5f,
	0xf9, 0x49, 0x46, 0x13, 0x92, 0x9e, 0xce, 0xd8, 0x07, 0x3a, 0x1d, 0xf4, 0xa2, 0x74, 0xb1, 0x51,
	0xda, 0x81, 0xf1, 0xb2, 0x8c, 0x46, 0x6d, 0x8e, 0x43, 0x1e, 0x4a, 0xe2, 0x7d, 0x49, 0xfc, 0x55,
	0x57, 0x52, 0x7a, 0xb7, 0x51, 0xa6, 0xbd, 0xf0, 0x5d, 0xa7, 0x63, 0xc7, 0xf4, 0x28, 0x6c, 0xdb,
	0xf4, 0xf6, 0x37, 0x85, 0x91, 0xba, 0x2e, 0xa5, 0x51, 0x6e, 0x17, 0x8d, 0xe7, 0x24, 0x9e, 0xa4,
	0x87, 0x01, 0x79, 0xdd, 0x11, 0x6f, 0xce, 0x80, 0xe4, 0x07, 0xa1, 0x77, 0x56, 0xc3, 0xcd, 0xeb,
	0xf0, 0x06, 0x98, 0xa7, 0x37, 0xd7, 0xce, 0x80, 0xfc, 0x7c, 0x1a, 0x97, 0x3d, 0xe3, 0x7b, 0x9e,
	0x09, 0xff, 0xff, 0x17, 0xcd, 0xf8, 0x45, 0xf1, 0xd0, 0xe2, 0xe2, 0x52, 0x4c, 0x9e, 0x93, 0xe5,
	0x31, 0x7d, 0xa2, 0x30, 0x28, 0x2b, 0x83, 0x5f, 0x77, 0xfb, 0x96, 0x2c, 0x27, 0x2b, 0x63, 0x4f,
	0x9a, 0xaa, 0x4e, 0xc8, 0xc7, 0x91, 0x4f, 0xbe, 0x11, 0xb2, 0xbd, 0xda, 0x56, 0xf7, 0xe2, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0xa2, 0xf5, 0xb3, 0xfa, 0x02, 0x00, 0x00,
}
