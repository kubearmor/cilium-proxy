// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/ratelimit/v2/rls.proto

package envoy_service_ratelimit_v2

import (
	context "context"
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	ratelimit "github.com/cilium/proxy/go/envoy/api/v2/ratelimit"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type RateLimitResponse_Code int32

const (
	// The response code is not known.
	RateLimitResponse_UNKNOWN RateLimitResponse_Code = 0
	// The response code to notify that the number of requests are under limit.
	RateLimitResponse_OK RateLimitResponse_Code = 1
	// The response code to notify that the number of requests are over limit.
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}

var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}

func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	// The time unit is not known.
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	// The time unit representing a second.
	RateLimitResponse_RateLimit_SECOND RateLimitResponse_RateLimit_Unit = 1
	// The time unit representing a minute.
	RateLimitResponse_RateLimit_MINUTE RateLimitResponse_RateLimit_Unit = 2
	// The time unit representing an hour.
	RateLimitResponse_RateLimit_HOUR RateLimitResponse_RateLimit_Unit = 3
	// The time unit representing a day.
	RateLimitResponse_RateLimit_DAY RateLimitResponse_RateLimit_Unit = 4
)

var RateLimitResponse_RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}

var RateLimitResponse_RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return proto.EnumName(RateLimitResponse_RateLimit_Unit_name, int32(x))
}

func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0, 0}
}

// Main message for a rate limit request. The rate limit service is designed to be fully generic
// in the sense that it can operate on arbitrary hierarchical key/value pairs. The loaded
// configuration will parse the request and find the most specific limit to apply. In addition,
// a RateLimitRequest can contain multiple "descriptors" to limit on. When multiple descriptors
// are provided, the server will limit on *ALL* of them and return an OVER_LIMIT response if any
// of them are over limit. This enables more complex application level rate limiting scenarios
// if desired.
type RateLimitRequest struct {
	// All rate limit requests must specify a domain. This enables the configuration to be per
	// application without fear of overlap. E.g., "envoy".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// All rate limit requests must specify at least one RateLimitDescriptor. Each descriptor is
	// processed by the service (see below). If any of the descriptors are over limit, the entire
	// request is considered to be over limit.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// Rate limit requests can optionally specify the number of hits a request adds to the matched
	// limit. If the value is not set in the message, a request increases the matched limit by 1.
	HitsAddend           uint32   `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{0}
}

func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRequest.Unmarshal(m, b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
}
func (m *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(m, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitRequest.Size(m)
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitRequest) GetHitsAddend() uint32 {
	if m != nil {
		return m.HitsAddend
	}
	return 0
}

// A response from a ShouldRateLimit call.
type RateLimitResponse struct {
	// The overall response code which takes into account all of the descriptors that were passed
	// in the RateLimitRequest message.
	OverallCode RateLimitResponse_Code `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"overall_code,omitempty"`
	// A list of DescriptorStatus messages which matches the length of the descriptor list passed
	// in the RateLimitRequest. This can be used by the caller to determine which individual
	// descriptors failed and/or what the currently configured limits are for all of them.
	Statuses []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	// A list of headers to add to the response
	Headers []*core.HeaderValue `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	// A list of headers to add to the request when forwarded
	RequestHeadersToAdd  []*core.HeaderValue `protobuf:"bytes,4,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1}
}

func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse.Unmarshal(m, b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(m, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse.Size(m)
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (m *RateLimitResponse) GetHeaders() []*core.HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *RateLimitResponse) GetRequestHeadersToAdd() []*core.HeaderValue {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

// Defines an actual rate limit in terms of requests per unit of time and the unit itself.
type RateLimitResponse_RateLimit struct {
	// The number of requests per unit of time.
	RequestsPerUnit uint32 `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	// The unit of time.
	Unit                 RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitResponse_RateLimit) Reset()         { *m = RateLimitResponse_RateLimit{} }
func (m *RateLimitResponse_RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_RateLimit) ProtoMessage()    {}
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0}
}

func (m *RateLimitResponse_RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Unmarshal(m, b)
}
func (m *RateLimitResponse_RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_RateLimit.Merge(m, src)
}
func (m *RateLimitResponse_RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Size(m)
}
func (m *RateLimitResponse_RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_RateLimit proto.InternalMessageInfo

func (m *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	// The response code for an individual descriptor.
	Code RateLimitResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"code,omitempty"`
	// The current limit as configured by the server. Useful for debugging, etc.
	CurrentLimit *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	// The limit remaining in the current time unit.
	LimitRemaining       uint32   `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 1}
}

func (m *RateLimitResponse_DescriptorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Unmarshal(m, b)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.Merge(m, src)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Size(m)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_DescriptorStatus proto.InternalMessageInfo

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit", RateLimitResponse_RateLimit_Unit_name, RateLimitResponse_RateLimit_Unit_value)
	proto.RegisterType((*RateLimitRequest)(nil), "envoy.service.ratelimit.v2.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "envoy.service.ratelimit.v2.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_RateLimit)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.RateLimit")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus")
}

func init() {
	proto.RegisterFile("envoy/service/ratelimit/v2/rls.proto", fileDescriptor_1de95711edb19ee8)
}

var fileDescriptor_1de95711edb19ee8 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0x9b, 0xdd, 0xb8, 0x6d, 0xcf, 0xf6, 0x23, 0x1d, 0xa1, 0x5d, 0x17, 0x69, 0xcb, 0x22,
	0xba, 0x58, 0x4d, 0x20, 0x5e, 0x88, 0x20, 0x85, 0x7e, 0x49, 0x4b, 0xdb, 0xdd, 0x65, 0xb6, 0x5b,
	0xa9, 0x08, 0x61, 0xba, 0x33, 0xb4, 0x03, 0x69, 0x26, 0xce, 0x4c, 0x82, 0xbd, 0xf7, 0xc2, 0x3b,
	0x6f, 0xc5, 0x37, 0xf2, 0xda, 0x27, 0xf0, 0x15, 0x7c, 0x00, 0x91, 0x4c, 0xb2, 0x1f, 0xad, 0x28,
	0xad, 0xde, 0x65, 0xce, 0x39, 0xff, 0x5f, 0xe6, 0x7f, 0xce, 0xcc, 0xc0, 0x03, 0x16, 0xa5, 0xe2,
	0xd2, 0x53, 0x4c, 0xa6, 0xbc, 0xcf, 0x3c, 0x49, 0x34, 0x0b, 0xf9, 0x05, 0xd7, 0x5e, 0xea, 0x7b,
	0x32, 0x54, 0x6e, 0x2c, 0x85, 0x16, 0xa8, 0x6e, 0xaa, 0xdc, 0xa2, 0xca, 0x1d, 0x56, 0xb9, 0xa9,
	0x5f, 0xbf, 0x9f, 0x13, 0x48, 0xcc, 0x33, 0x4d, 0x5f, 0x48, 0xe6, 0x9d, 0x12, 0xc5, 0x72, 0x65,
	0xfd, 0xe1, 0x95, 0xec, 0x08, 0x3f, 0x42, 0xe4, 0x75, 0xcb, 0x09, 0x8d, 0x89, 0x47, 0xa2, 0x48,
	0x68, 0xa2, 0xb9, 0x88, 0x94, 0x77, 0xc1, 0xcf, 0xb2, 0xa2, 0x22, 0xbf, 0x94, 0x92, 0x90, 0x53,
	0xa2, 0x99, 0x37, 0xf8, 0xc8, 0x13, 0x8d, 0x2f, 0x16, 0x38, 0x98, 0x68, 0x76, 0x90, 0xc1, 0x30,
	0x7b, 0x97, 0x30, 0xa5, 0xd1, 0x22, 0x54, 0xa8, 0xb8, 0x20, 0x3c, 0xaa, 0x59, 0xab, 0x56, 0x73,
	0x1a, 0x17, 0x2b, 0x74, 0x08, 0x55, 0xca, 0x54, 0x5f, 0xf2, 0x58, 0x0b, 0xa9, 0x6a, 0xa5, 0xd5,
	0x72, 0xb3, 0xea, 0xaf, 0xb9, 0xb9, 0x3b, 0x12, 0x73, 0x37, 0xf5, 0xc7, 0xcc, 0x0d, 0xb1, 0xdb,
	0x43, 0x0d, 0x1e, 0xd7, 0xa3, 0x15, 0xa8, 0x9e, 0x73, 0xad, 0x02, 0x42, 0x29, 0x8b, 0x68, 0xad,
	0xbc, 0x6a, 0x35, 0x67, 0x31, 0x64, 0xa1, 0x0d, 0x13, 0x69, 0x7c, 0xab, 0xc0, 0xc2, 0xd8, 0xe6,
	0x54, 0x2c, 0x22, 0xc5, 0x50, 0x0f, 0x66, 0x44, 0xca, 0x24, 0x09, 0xc3, 0xa0, 0x2f, 0x28, 0x33,
	0x7b, 0x9c, 0xf3, 0x7d, 0xf7, 0xcf, 0x4d, 0x76, 0x7f, 0x83, 0xb8, 0x5b, 0x82, 0x32, 0x5c, 0x2d,
	0x38, 0xd9, 0x02, 0xbd, 0x81, 0x29, 0xa5, 0x89, 0x4e, 0x14, 0x1b, 0x38, 0x5b, 0xbf, 0x1d, 0x72,
	0x64, 0xb3, 0x6b, 0x38, 0x78, 0xc8, 0x43, 0x27, 0x30, 0x79, 0xce, 0x08, 0x65, 0x52, 0xd5, 0xca,
	0x06, 0xbd, 0x7c, 0xb5, 0x69, 0xd9, 0xd8, 0xdd, 0x5d, 0x53, 0x71, 0x4c, 0xc2, 0x84, 0x6d, 0xae,
	0xfc, 0xf8, 0xfc, 0xf3, 0xd3, 0x9d, 0x7b, 0xb0, 0x24, 0x0b, 0x7a, 0x50, 0xe8, 0x03, 0x2d, 0xb2,
	0x7e, 0xe1, 0x01, 0x0f, 0x75, 0x61, 0x51, 0xe6, 0x63, 0xbb, 0x56, 0x52, 0xb3, 0x6f, 0xf2, 0x27,
	0x7c, 0xb7, 0x50, 0xe7, 0x31, 0x75, 0x24, 0x36, 0x28, 0xad, 0x7f, 0xb5, 0x60, 0x7a, 0x68, 0x10,
	0x3d, 0x86, 0x85, 0xa2, 0x48, 0x05, 0x31, 0x93, 0x41, 0x12, 0x71, 0x6d, 0xba, 0x3e, 0x8b, 0xe7,
	0x07, 0x89, 0x0e, 0x93, 0xbd, 0x88, 0x6b, 0xd4, 0x01, 0xdb, 0xa4, 0x4b, 0x66, 0x28, 0x2f, 0x6f,
	0xd7, 0xc1, 0x61, 0xc4, 0xcd, 0x58, 0xd8, 0x90, 0x1a, 0xeb, 0x60, 0x1b, 0x72, 0x15, 0x26, 0x7b,
	0xad, 0xfd, 0x56, 0xfb, 0x75, 0xcb, 0x99, 0x40, 0x00, 0x95, 0xee, 0xce, 0x56, 0xbb, 0xb5, 0xed,
	0x58, 0xd9, 0xf7, 0xe1, 0x5e, 0xab, 0x77, 0xb4, 0xe3, 0x94, 0xd0, 0x14, 0xd8, 0xbb, 0xed, 0x1e,
	0x76, 0xca, 0x68, 0x12, 0xca, 0xdb, 0x1b, 0x27, 0x8e, 0x5d, 0xff, 0x6e, 0x81, 0x73, 0x7d, 0x34,
	0xe8, 0x15, 0xd8, 0xff, 0x79, 0x76, 0x8c, 0x1e, 0xbd, 0x85, 0xd9, 0x7e, 0x22, 0x25, 0x8b, 0x74,
	0x60, 0x04, 0xc6, 0x77, 0xd5, 0x7f, 0xfe, 0x8f, 0xbe, 0xf1, 0x4c, 0x41, 0xcb, 0x1b, 0xff, 0x08,
	0xe6, 0x8d, 0x2a, 0x90, 0x2c, 0xbb, 0x7f, 0x3c, 0x3a, 0x2b, 0x2e, 0xc9, 0x5c, 0x98, 0xeb, 0x8b,
	0x68, 0x63, 0x0d, 0x6c, 0x73, 0x86, 0xaf, 0xf4, 0xa8, 0x02, 0xa5, 0xf6, 0xbe, 0x63, 0xa1, 0x39,
	0x80, 0xf6, 0xf1, 0x0e, 0x0e, 0x0e, 0xf6, 0x0e, 0xf7, 0x8e, 0x9c, 0x92, 0xff, 0x61, 0xfc, 0xca,
	0x77, 0xf3, 0x1d, 0xa2, 0x18, 0xe6, 0xbb, 0xe7, 0x22, 0x09, 0xe9, 0x68, 0xec, 0x4f, 0x6e, 0x68,
	0xc2, 0x1c, 0x80, 0xfa, 0xd3, 0x5b, 0x59, 0x6e, 0x4c, 0x6c, 0xbe, 0x80, 0x26, 0x17, 0xb9, 0x28,
	0x96, 0xe2, 0xfd, 0xe5, 0x5f, 0xf4, 0x9b, 0x53, 0x38, 0x54, 0x9d, 0xec, 0xbd, 0xea, 0x58, 0x1f,
	0x2d, 0xeb, 0xb4, 0x62, 0xde, 0xae, 0x67, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xff, 0xc2,
	0x79, 0x7e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

// UnimplementedRateLimitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitServiceServer struct {
}

func (*UnimplementedRateLimitServiceServer) ShouldRateLimit(ctx context.Context, req *RateLimitRequest) (*RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShouldRateLimit not implemented")
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v2.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v2/rls.proto",
}
