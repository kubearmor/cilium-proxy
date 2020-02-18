// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/fault/v2/fault.proto

package envoy_config_filter_http_fault_v2

import (
	fmt "fmt"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	v2 "github.com/cilium/proxy/go/envoy/config/filter/fault/v2"
	_type "github.com/cilium/proxy/go/envoy/type"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	// The percentage of requests/operations/connections that will be aborted with the error code
	// provided.
	Percentage           *_type.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{0}
}

func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

// [#next-free-field: 14]
type HTTPFault struct {
	// If specified, the filter will inject delays based on the values in the
	// object.
	Delay *v2.FaultDelay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percentage
	// <envoy_api_field_config.filter.http.fault.v2.FaultAbort.percentage>` field.
	// The filter will check the request's headers against all the specified
	// headers in the filter config. A match will happen if all the headers in the
	// config are present in the request with the same values (or based on
	// presence if the *value* field is not in the config).
	Headers []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	// The maximum number of faults that can be active at a single time via the configured fault
	// filter. Note that because this setting can be overridden at the route level, it's possible
	// for the number of active faults to be greater than this value (if injected via a different
	// route). If not specified, defaults to unlimited. This setting can be overridden via
	// `runtime <config_http_filters_fault_injection_runtime>` and any faults that are not injected
	// due to overflow will be indicated via the `faults_overflow
	// <config_http_filters_fault_injection_stats>` stat.
	//
	// .. attention::
	//   Like other :ref:`circuit breakers <arch_overview_circuit_break>` in Envoy, this is a fuzzy
	//   limit. It's possible for the number of active faults to rise slightly above the configured
	//   amount due to the implementation details.
	MaxActiveFaults *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	// The response rate limit to be applied to the response body of the stream. When configured,
	// the percentage can be overridden by the :ref:`fault.http.rate_limit.response_percent
	// <config_http_filters_fault_injection_runtime>` runtime key.
	//
	// .. attention::
	//  This is a per-stream limit versus a connection level limit. This means that concurrent streams
	//  will each get an independent limit.
	ResponseRateLimit *v2.FaultRateLimit `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_delay_percent
	DelayPercentRuntime string `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.abort_percent
	AbortPercentRuntime string `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_duration_ms
	DelayDurationRuntime string `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.http_status
	AbortHttpStatusRuntime string `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.max_active_faults
	MaxActiveFaultsRuntime string `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.rate_limit.response_percent
	ResponseRateLimitPercentRuntime string   `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{1}
}

func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v2.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func (m *HTTPFault) GetDelayPercentRuntime() string {
	if m != nil {
		return m.DelayPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortPercentRuntime() string {
	if m != nil {
		return m.AbortPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetDelayDurationRuntime() string {
	if m != nil {
		return m.DelayDurationRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortHttpStatusRuntime() string {
	if m != nil {
		return m.AbortHttpStatusRuntime
	}
	return ""
}

func (m *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if m != nil {
		return m.MaxActiveFaultsRuntime
	}
	return ""
}

func (m *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if m != nil {
		return m.ResponseRateLimitPercentRuntime
	}
	return ""
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.config.filter.http.fault.v2.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.config.filter.http.fault.v2.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/fault/v2/fault.proto", fileDescriptor_26070db6b6576d5c)
}

var fileDescriptor_26070db6b6576d5c = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xd3, 0x4a,
	0x10, 0xc6, 0x8f, 0x9b, 0x26, 0x6d, 0x37, 0xa7, 0x6a, 0xeb, 0x9e, 0x53, 0x4c, 0x05, 0x25, 0x2d,
	0x12, 0x4a, 0x11, 0x5d, 0x4b, 0x29, 0x37, 0x08, 0x81, 0xd4, 0xb4, 0xaa, 0x02, 0x2a, 0x28, 0x32,
	0x85, 0x2b, 0x24, 0x6b, 0x1b, 0x4f, 0x52, 0x4b, 0xf6, 0xae, 0xb5, 0xbb, 0x4e, 0x93, 0xa7, 0xe0,
	0x86, 0x0b, 0x5e, 0x82, 0xf7, 0xe0, 0x11, 0x78, 0x0e, 0xae, 0x10, 0x17, 0x08, 0xed, 0xac, 0x9d,
	0xfe, 0x49, 0xa4, 0x72, 0x13, 0x39, 0x3b, 0xdf, 0x6f, 0x66, 0x67, 0xe6, 0xb3, 0xc9, 0x1e, 0xf0,
	0xa1, 0x18, 0xfb, 0x3d, 0xc1, 0xfb, 0xf1, 0xc0, 0xef, 0xc7, 0x89, 0x06, 0xe9, 0x9f, 0x6b, 0x9d,
	0xf9, 0x7d, 0x96, 0x27, 0xda, 0x1f, 0xb6, 0xec, 0x03, 0xcd, 0xa4, 0xd0, 0xc2, 0xdd, 0x46, 0x39,
	0xb5, 0x72, 0x6a, 0xe5, 0xd4, 0xc8, 0xa9, 0x55, 0x0d, 0x5b, 0x9b, 0xbb, 0x36, 0x23, 0xcb, 0x62,
	0x03, 0x4b, 0x91, 0x6b, 0xb0, 0xbf, 0x61, 0x4f, 0xa4, 0x99, 0xe0, 0xc0, 0xb5, 0xb2, 0xd9, 0x36,
	0x9b, 0xb3, 0x8a, 0xcf, 0xaa, 0xbb, 0xe9, 0x59, 0xa5, 0x1e, 0x67, 0xe0, 0x67, 0x20, 0x7b, 0xc0,
	0xcb, 0xc8, 0xd6, 0x40, 0x88, 0x41, 0x02, 0x3e, 0xfe, 0x3b, 0xcb, 0xfb, 0xfe, 0x85, 0x64, 0x59,
	0x06, 0xb2, 0xac, 0xb1, 0x95, 0x47, 0x19, 0xf3, 0x19, 0xe7, 0x42, 0x33, 0x1d, 0x0b, 0xae, 0xfc,
	0x34, 0x1e, 0x48, 0xa6, 0xa1, 0x88, 0xdf, 0x19, 0xb2, 0x24, 0x8e, 0x98, 0x06, 0xbf, 0x7c, 0xb0,
	0x81, 0x9d, 0xcf, 0x0e, 0x21, 0xc7, 0xe6, 0x0a, 0x07, 0x67, 0x42, 0x6a, 0x97, 0x92, 0xba, 0xe9,
	0x33, 0x54, 0x9a, 0xe9, 0x5c, 0x79, 0x73, 0x0d, 0xa7, 0xb9, 0xdc, 0xae, 0xff, 0x6a, 0x2f, 0x3e,
	0xae, 0xad, 0x7e, 0x9f, 0x6f, 0x7e, 0x73, 0x3a, 0xff, 0x04, 0xc4, 0x28, 0xde, 0xa1, 0xc0, 0x7d,
	0x41, 0x48, 0x71, 0x51, 0x36, 0x00, 0xaf, 0xd2, 0x70, 0x9a, 0xf5, 0xd6, 0x7d, 0x6a, 0xc7, 0x67,
	0xda, 0xa0, 0xc7, 0x92, 0xf5, 0xcc, 0x85, 0x58, 0xd2, 0xb5, 0xba, 0xe0, 0x0a, 0xd0, 0x5e, 0x23,
	0x04, 0xa4, 0x14, 0x32, 0x34, 0x5a, 0xb7, 0xf2, 0xb3, 0xed, 0xbc, 0x9e, 0x5f, 0x74, 0x56, 0xe7,
	0x76, 0xbe, 0xd6, 0xc8, 0x52, 0xe7, 0xf4, 0xb4, 0x8b, 0x57, 0x73, 0x5f, 0x92, 0x6a, 0x04, 0x09,
	0x1b, 0x7b, 0x0e, 0x16, 0x68, 0xd2, 0x59, 0xfb, 0x29, 0x57, 0x43, 0x91, 0x39, 0x32, 0xfa, 0xc0,
	0x62, 0xee, 0x21, 0xa9, 0x32, 0xd3, 0x1e, 0xf6, 0x53, 0x6f, 0xed, 0xd1, 0x5b, 0xf7, 0x4b, 0x2f,
	0x67, 0x12, 0x58, 0xd6, 0xdd, 0x25, 0xab, 0x79, 0xa6, 0xb4, 0x04, 0x96, 0x86, 0xbd, 0x24, 0x57,
	0x1a, 0x24, 0x36, 0xbc, 0x14, 0xac, 0x94, 0xe7, 0x87, 0xf6, 0xd8, 0x7d, 0x4e, 0x16, 0xce, 0x81,
	0x45, 0x20, 0x95, 0x37, 0xdf, 0xa8, 0x34, 0xeb, 0xad, 0xed, 0xa2, 0x22, 0xcb, 0x62, 0x93, 0x1c,
	0x8d, 0x42, 0x3b, 0x28, 0x79, 0xc3, 0x74, 0xef, 0x1c, 0x64, 0x50, 0x12, 0xa6, 0x4e, 0x24, 0x2e,
	0x78, 0x51, 0x89, 0x8b, 0x08, 0x94, 0x57, 0x6d, 0x54, 0x4c, 0x9d, 0xcb, 0xf3, 0xb7, 0xe6, 0xd8,
	0xed, 0x90, 0xb5, 0x94, 0x8d, 0x42, 0x33, 0xe1, 0x21, 0x84, 0x78, 0x77, 0xe5, 0xd5, 0xb0, 0xc7,
	0x7b, 0xd4, 0x3a, 0x86, 0x96, 0x8e, 0xa1, 0xef, 0x5f, 0x71, 0xbd, 0xdf, 0xfa, 0xc0, 0x92, 0x1c,
	0x82, 0x95, 0x94, 0x8d, 0x0e, 0x90, 0xc2, 0x3e, 0x95, 0xfb, 0x91, 0xac, 0x4b, 0x50, 0x99, 0xe0,
	0x0a, 0x42, 0x63, 0x9b, 0x30, 0x89, 0xd3, 0x58, 0x7b, 0x0b, 0x98, 0xeb, 0xc9, 0x5f, 0xcc, 0x3b,
	0x60, 0x1a, 0x4e, 0x0c, 0x13, 0xac, 0x95, 0x89, 0x26, 0x47, 0x6e, 0x8b, 0xfc, 0x8f, 0x8b, 0x08,
	0x8b, 0xd5, 0x87, 0x32, 0xe7, 0x3a, 0x4e, 0xc1, 0x5b, 0xc4, 0xf9, 0xad, 0x63, 0xb0, 0xf4, 0x87,
	0x0d, 0x19, 0x06, 0xe7, 0x3e, 0xc5, 0x2c, 0x59, 0x06, 0x83, 0x37, 0x98, 0xa7, 0x64, 0xc3, 0xd6,
	0x89, 0x72, 0x89, 0xaf, 0xc1, 0x04, 0x22, 0x08, 0xfd, 0x87, 0xd1, 0xa3, 0x22, 0x58, 0x52, 0xcf,
	0xc8, 0x5d, 0x5b, 0xe9, 0x8a, 0xf3, 0x27, 0x60, 0x1d, 0xc1, 0x0d, 0x14, 0x74, 0x26, 0xbe, 0xbf,
	0x82, 0x4e, 0x2d, 0x60, 0x82, 0xfe, 0x6b, 0xd1, 0x1b, 0xa3, 0x2e, 0xd1, 0x13, 0xf2, 0x70, 0xc6,
	0xc4, 0xa7, 0xba, 0x5d, 0xc6, 0x24, 0x0f, 0xa6, 0x66, 0x7a, 0xbd, 0xf3, 0x76, 0xfa, 0xe3, 0xcb,
	0xef, 0x4f, 0xd5, 0xa6, 0xfb, 0xc8, 0x6e, 0x0a, 0x46, 0x1a, 0xb8, 0x32, 0xdf, 0x81, 0x62, 0x5b,
	0xea, 0x9a, 0xbd, 0xf7, 0x89, 0x1f, 0x0b, 0xbb, 0xd4, 0x4c, 0x8a, 0xd1, 0xf8, 0xf6, 0xf7, 0xa1,
	0x6d, 0x3f, 0x12, 0x5d, 0x63, 0xa6, 0xae, 0x73, 0x56, 0x43, 0x57, 0xed, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x17, 0x2c, 0xe0, 0x7b, 0x58, 0x05, 0x00, 0x00,
}
