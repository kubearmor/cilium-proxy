// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/runtime/v3/rtds.proto

package envoy_service_runtime_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/service/discovery/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type RtdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RtdsDummy) Reset()         { *m = RtdsDummy{} }
func (m *RtdsDummy) String() string { return proto.CompactTextString(m) }
func (*RtdsDummy) ProtoMessage()    {}
func (*RtdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_40cfda5ab878f66c, []int{0}
}

func (m *RtdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RtdsDummy.Unmarshal(m, b)
}
func (m *RtdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RtdsDummy.Marshal(b, m, deterministic)
}
func (m *RtdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RtdsDummy.Merge(m, src)
}
func (m *RtdsDummy) XXX_Size() int {
	return xxx_messageInfo_RtdsDummy.Size(m)
}
func (m *RtdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_RtdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_RtdsDummy proto.InternalMessageInfo

// RTDS resource type. This describes a layer in the runtime virtual filesystem.
type Runtime struct {
	// Runtime resource name. This makes the Runtime a self-describing xDS
	// resource.
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Layer                *_struct.Struct `protobuf:"bytes,2,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_40cfda5ab878f66c, []int{1}
}

func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (m *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(m, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Runtime) GetLayer() *_struct.Struct {
	if m != nil {
		return m.Layer
	}
	return nil
}

func init() {
	proto.RegisterType((*RtdsDummy)(nil), "envoy.service.runtime.v3.RtdsDummy")
	proto.RegisterType((*Runtime)(nil), "envoy.service.runtime.v3.Runtime")
}

func init() {
	proto.RegisterFile("envoy/service/runtime/v3/rtds.proto", fileDescriptor_40cfda5ab878f66c)
}

var fileDescriptor_40cfda5ab878f66c = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0x71, 0xef, 0xa0, 0x3a, 0x73, 0x2c, 0x91, 0x50, 0xab, 0x82, 0x50, 0x08, 0xa7, 0xaa,
	0x1c, 0x9c, 0x0d, 0xed, 0x72, 0xca, 0x18, 0x55, 0xcc, 0xa7, 0xf4, 0x13, 0xf8, 0x92, 0x47, 0xb1,
	0x94, 0xd8, 0xc1, 0x76, 0x2c, 0x22, 0x16, 0xc4, 0x74, 0x62, 0x3d, 0x89, 0x81, 0x99, 0xaf, 0xc2,
	0xc2, 0xca, 0x57, 0xe0, 0x1b, 0xb0, 0x31, 0xa1, 0xc4, 0x49, 0xb9, 0x80, 0x5a, 0xc1, 0x72, 0x9b,
	0xdd, 0xf7, 0xfb, 0xbf, 0xf7, 0xff, 0x37, 0xcf, 0xf8, 0x11, 0x08, 0x2b, 0x2b, 0xaa, 0x41, 0x59,
	0x9e, 0x00, 0x55, 0xa5, 0x30, 0x3c, 0x07, 0x6a, 0x17, 0x54, 0x99, 0x54, 0x93, 0x42, 0x49, 0x23,
	0xbd, 0x71, 0x03, 0x91, 0x16, 0x22, 0x2d, 0x44, 0xec, 0x62, 0x72, 0xdc, 0x97, 0xa7, 0x5c, 0x27,
	0xd2, 0x82, 0xaa, 0xea, 0x06, 0x9b, 0x8b, 0xeb, 0x32, 0xb9, 0xbf, 0x96, 0x72, 0x9d, 0x01, 0x65,
	0x05, 0xa7, 0x4c, 0x08, 0x69, 0x98, 0xe1, 0x52, 0xe8, 0x3f, 0xaa, 0xcd, 0xed, 0xbc, 0x7c, 0x49,
	0xb5, 0x51, 0x65, 0x62, 0xda, 0xea, 0xc3, 0x32, 0x2d, 0xd8, 0x55, 0x15, 0xb5, 0xa0, 0x34, 0x97,
	0x82, 0x8b, 0x75, 0x8b, 0xf8, 0xce, 0xca, 0x55, 0x46, 0x81, 0x96, 0xa5, 0x4a, 0xa0, 0x25, 0x46,
	0x96, 0x65, 0x3c, 0x65, 0x06, 0x68, 0x77, 0x70, 0x85, 0xe0, 0x14, 0x1f, 0xc4, 0x26, 0xd5, 0xcb,
	0x32, 0xcf, 0xab, 0xf0, 0xc9, 0xa7, 0x2f, 0x17, 0x0f, 0xa6, 0xf8, 0xa8, 0x9f, 0xf9, 0x77, 0x18,
	0x3b, 0x27, 0x1b, 0x38, 0x78, 0x87, 0xf0, 0x30, 0x76, 0x7f, 0x87, 0x77, 0x0f, 0xef, 0x0b, 0x96,
	0xc3, 0x18, 0xf9, 0x68, 0x76, 0x10, 0x0d, 0x7f, 0x46, 0xfb, 0x6a, 0xe0, 0xa3, 0xb8, 0xf9, 0xd1,
	0x3b, 0xc1, 0x37, 0x33, 0x56, 0x81, 0x1a, 0x0f, 0x7c, 0x34, 0xbb, 0x3d, 0x1f, 0x11, 0x17, 0x97,
	0x74, 0x71, 0xc9, 0xaa, 0x89, 0x1b, 0x3b, 0x2a, 0x7c, 0x5c, 0x9b, 0x38, 0xc2, 0xc1, 0x2e, 0x13,
	0x6e, 0xec, 0xfc, 0xeb, 0x1e, 0x1e, 0xb5, 0xe7, 0x65, 0x57, 0x5f, 0x39, 0x81, 0xa7, 0xf0, 0x9d,
	0x95, 0x51, 0xc0, 0xf2, 0xce, 0xe3, 0x53, 0xb2, 0xb5, 0xe3, 0x82, 0x6c, 0xe4, 0x31, 0xbc, 0x2e,
	0x41, 0x9b, 0xc9, 0xc9, 0x3f, 0xd2, 0xba, 0x90, 0x42, 0x43, 0x70, 0x63, 0x86, 0x9e, 0x21, 0xef,
	0x2d, 0x3e, 0x5c, 0x42, 0x66, 0x58, 0x37, 0xf2, 0xf9, 0xce, 0x26, 0x35, 0xf9, 0xd7, 0xdc, 0xf9,
	0xff, 0x48, 0x7a, 0xc3, 0x2f, 0x11, 0x3e, 0x7c, 0x01, 0x26, 0x79, 0x75, 0x2d, 0x81, 0xa7, 0xef,
	0xbf, 0x7d, 0xbf, 0x1c, 0x8c, 0x82, 0xbb, 0xbd, 0x75, 0x0f, 0xdb, 0xf7, 0xd1, 0x14, 0xf7, 0x42,
	0x74, 0x3c, 0x99, 0x7d, 0xf8, 0xfc, 0xf1, 0xc7, 0x30, 0xc0, 0xfe, 0xb6, 0x77, 0xd4, 0x7d, 0xcc,
	0xe8, 0x14, 0x4f, 0xb9, 0x74, 0x26, 0x0a, 0x25, 0xdf, 0x54, 0x64, 0x9b, 0x22, 0x6a, 0x36, 0xf6,
	0xac, 0xde, 0x9e, 0x33, 0x74, 0x81, 0xd0, 0xf9, 0xad, 0x66, 0x93, 0x16, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x5f, 0xc5, 0x06, 0x8a, 0xd1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RuntimeDiscoveryServiceClient is the client API for RuntimeDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeDiscoveryServiceClient interface {
	StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error)
	DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error)
	FetchRuntime(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type runtimeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeDiscoveryServiceClient(cc *grpc.ClientConn) RuntimeDiscoveryServiceClient {
	return &runtimeDiscoveryServiceClient{cc}
}

func (c *runtimeDiscoveryServiceClient) StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[0], "/envoy.service.runtime.v3.RuntimeDiscoveryService/StreamRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceStreamRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_StreamRuntimeClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceStreamRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[1], "/envoy.service.runtime.v3.RuntimeDiscoveryService/DeltaRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceDeltaRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_DeltaRuntimeClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceDeltaRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) FetchRuntime(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.runtime.v3.RuntimeDiscoveryService/FetchRuntime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeDiscoveryServiceServer is the server API for RuntimeDiscoveryService service.
type RuntimeDiscoveryServiceServer interface {
	StreamRuntime(RuntimeDiscoveryService_StreamRuntimeServer) error
	DeltaRuntime(RuntimeDiscoveryService_DeltaRuntimeServer) error
	FetchRuntime(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedRuntimeDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRuntimeDiscoveryServiceServer struct {
}

func (*UnimplementedRuntimeDiscoveryServiceServer) StreamRuntime(srv RuntimeDiscoveryService_StreamRuntimeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRuntime not implemented")
}
func (*UnimplementedRuntimeDiscoveryServiceServer) DeltaRuntime(srv RuntimeDiscoveryService_DeltaRuntimeServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaRuntime not implemented")
}
func (*UnimplementedRuntimeDiscoveryServiceServer) FetchRuntime(ctx context.Context, req *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRuntime not implemented")
}

func RegisterRuntimeDiscoveryServiceServer(s *grpc.Server, srv RuntimeDiscoveryServiceServer) {
	s.RegisterService(&_RuntimeDiscoveryService_serviceDesc, srv)
}

func _RuntimeDiscoveryService_StreamRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).StreamRuntime(&runtimeDiscoveryServiceStreamRuntimeServer{stream})
}

type RuntimeDiscoveryService_StreamRuntimeServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceStreamRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_DeltaRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).DeltaRuntime(&runtimeDiscoveryServiceDeltaRuntimeServer{stream})
}

type RuntimeDiscoveryService_DeltaRuntimeServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceDeltaRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_FetchRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.runtime.v3.RuntimeDiscoveryService/FetchRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RuntimeDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.runtime.v3.RuntimeDiscoveryService",
	HandlerType: (*RuntimeDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRuntime",
			Handler:    _RuntimeDiscoveryService_FetchRuntime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRuntime",
			Handler:       _RuntimeDiscoveryService_StreamRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRuntime",
			Handler:       _RuntimeDiscoveryService_DeltaRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/runtime/v3/rtds.proto",
}
