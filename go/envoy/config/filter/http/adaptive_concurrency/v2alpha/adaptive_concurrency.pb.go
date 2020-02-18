// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/adaptive_concurrency/v2alpha/adaptive_concurrency.proto

package envoy_config_filter_http_adaptive_concurrency_v2alpha

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_type "github.com/cilium/proxy/go/envoy/type"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Configuration parameters for the gradient controller.
type GradientControllerConfig struct {
	// The percentile to use when summarizing aggregated samples. Defaults to p50.
	SampleAggregatePercentile *_type.Percent                                              `protobuf:"bytes,1,opt,name=sample_aggregate_percentile,json=sampleAggregatePercentile,proto3" json:"sample_aggregate_percentile,omitempty"`
	ConcurrencyLimitParams    *GradientControllerConfig_ConcurrencyLimitCalculationParams `protobuf:"bytes,2,opt,name=concurrency_limit_params,json=concurrencyLimitParams,proto3" json:"concurrency_limit_params,omitempty"`
	MinRttCalcParams          *GradientControllerConfig_MinimumRTTCalculationParams       `protobuf:"bytes,3,opt,name=min_rtt_calc_params,json=minRttCalcParams,proto3" json:"min_rtt_calc_params,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                                                    `json:"-"`
	XXX_unrecognized          []byte                                                      `json:"-"`
	XXX_sizecache             int32                                                       `json:"-"`
}

func (m *GradientControllerConfig) Reset()         { *m = GradientControllerConfig{} }
func (m *GradientControllerConfig) String() string { return proto.CompactTextString(m) }
func (*GradientControllerConfig) ProtoMessage()    {}
func (*GradientControllerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c58a0beecb0ec580, []int{0}
}

func (m *GradientControllerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig.Unmarshal(m, b)
}
func (m *GradientControllerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig.Merge(m, src)
}
func (m *GradientControllerConfig) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig.Size(m)
}
func (m *GradientControllerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig proto.InternalMessageInfo

func (m *GradientControllerConfig) GetSampleAggregatePercentile() *_type.Percent {
	if m != nil {
		return m.SampleAggregatePercentile
	}
	return nil
}

func (m *GradientControllerConfig) GetConcurrencyLimitParams() *GradientControllerConfig_ConcurrencyLimitCalculationParams {
	if m != nil {
		return m.ConcurrencyLimitParams
	}
	return nil
}

func (m *GradientControllerConfig) GetMinRttCalcParams() *GradientControllerConfig_MinimumRTTCalculationParams {
	if m != nil {
		return m.MinRttCalcParams
	}
	return nil
}

// Parameters controlling the periodic recalculation of the concurrency limit from sampled request
// latencies.
type GradientControllerConfig_ConcurrencyLimitCalculationParams struct {
	// The allowed upper-bound on the calculated concurrency limit. Defaults to 1000.
	MaxConcurrencyLimit *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrency_limit,json=maxConcurrencyLimit,proto3" json:"max_concurrency_limit,omitempty"`
	// The period of time samples are taken to recalculate the concurrency limit.
	ConcurrencyUpdateInterval *duration.Duration `protobuf:"bytes,3,opt,name=concurrency_update_interval,json=concurrencyUpdateInterval,proto3" json:"concurrency_update_interval,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) Reset() {
	*m = GradientControllerConfig_ConcurrencyLimitCalculationParams{}
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) String() string {
	return proto.CompactTextString(m)
}
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) ProtoMessage() {}
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c58a0beecb0ec580, []int{0, 0}
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Unmarshal(m, b)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Merge(m, src)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Size(m)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams proto.InternalMessageInfo

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetMaxConcurrencyLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrencyLimit
	}
	return nil
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetConcurrencyUpdateInterval() *duration.Duration {
	if m != nil {
		return m.ConcurrencyUpdateInterval
	}
	return nil
}

// Parameters controlling the periodic minRTT recalculation.
// [#next-free-field: 6]
type GradientControllerConfig_MinimumRTTCalculationParams struct {
	// The time interval between recalculating the minimum request round-trip time.
	Interval *duration.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// The number of requests to aggregate/sample during the minRTT recalculation window before
	// updating. Defaults to 50.
	RequestCount *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	// Randomized time delta that will be introduced to the start of the minRTT calculation window.
	// This is represented as a percentage of the interval duration. Defaults to 15%.
	//
	// Example: If the interval is 10s and the jitter is 15%, the next window will begin
	// somewhere in the range (10s - 11.5s).
	Jitter *_type.Percent `protobuf:"bytes,3,opt,name=jitter,proto3" json:"jitter,omitempty"`
	// The concurrency limit set while measuring the minRTT. Defaults to 3.
	MinConcurrency *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=min_concurrency,json=minConcurrency,proto3" json:"min_concurrency,omitempty"`
	// Amount added to the measured minRTT to add stability to the concurrency limit during natural
	// variability in latency. This is expressed as a percentage of the measured value and can be
	// adjusted to allow more or less tolerance to the sampled latency values.
	//
	// Defaults to 25%.
	Buffer               *_type.Percent `protobuf:"bytes,5,opt,name=buffer,proto3" json:"buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) Reset() {
	*m = GradientControllerConfig_MinimumRTTCalculationParams{}
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) String() string {
	return proto.CompactTextString(m)
}
func (*GradientControllerConfig_MinimumRTTCalculationParams) ProtoMessage() {}
func (*GradientControllerConfig_MinimumRTTCalculationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c58a0beecb0ec580, []int{0, 1}
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Unmarshal(m, b)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Merge(m, src)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Size(m)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams proto.InternalMessageInfo

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetRequestCount() *wrappers.UInt32Value {
	if m != nil {
		return m.RequestCount
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetJitter() *_type.Percent {
	if m != nil {
		return m.Jitter
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetMinConcurrency() *wrappers.UInt32Value {
	if m != nil {
		return m.MinConcurrency
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetBuffer() *_type.Percent {
	if m != nil {
		return m.Buffer
	}
	return nil
}

type AdaptiveConcurrency struct {
	// Types that are valid to be assigned to ConcurrencyControllerConfig:
	//	*AdaptiveConcurrency_GradientControllerConfig
	ConcurrencyControllerConfig isAdaptiveConcurrency_ConcurrencyControllerConfig `protobuf_oneof:"concurrency_controller_config"`
	// If set to false, the adaptive concurrency filter will operate as a pass-through filter. If the
	// message is unspecified, the filter will be enabled.
	Enabled              *core.RuntimeFeatureFlag `protobuf:"bytes,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AdaptiveConcurrency) Reset()         { *m = AdaptiveConcurrency{} }
func (m *AdaptiveConcurrency) String() string { return proto.CompactTextString(m) }
func (*AdaptiveConcurrency) ProtoMessage()    {}
func (*AdaptiveConcurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_c58a0beecb0ec580, []int{1}
}

func (m *AdaptiveConcurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdaptiveConcurrency.Unmarshal(m, b)
}
func (m *AdaptiveConcurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdaptiveConcurrency.Marshal(b, m, deterministic)
}
func (m *AdaptiveConcurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdaptiveConcurrency.Merge(m, src)
}
func (m *AdaptiveConcurrency) XXX_Size() int {
	return xxx_messageInfo_AdaptiveConcurrency.Size(m)
}
func (m *AdaptiveConcurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_AdaptiveConcurrency.DiscardUnknown(m)
}

var xxx_messageInfo_AdaptiveConcurrency proto.InternalMessageInfo

type isAdaptiveConcurrency_ConcurrencyControllerConfig interface {
	isAdaptiveConcurrency_ConcurrencyControllerConfig()
}

type AdaptiveConcurrency_GradientControllerConfig struct {
	GradientControllerConfig *GradientControllerConfig `protobuf:"bytes,1,opt,name=gradient_controller_config,json=gradientControllerConfig,proto3,oneof"`
}

func (*AdaptiveConcurrency_GradientControllerConfig) isAdaptiveConcurrency_ConcurrencyControllerConfig() {
}

func (m *AdaptiveConcurrency) GetConcurrencyControllerConfig() isAdaptiveConcurrency_ConcurrencyControllerConfig {
	if m != nil {
		return m.ConcurrencyControllerConfig
	}
	return nil
}

func (m *AdaptiveConcurrency) GetGradientControllerConfig() *GradientControllerConfig {
	if x, ok := m.GetConcurrencyControllerConfig().(*AdaptiveConcurrency_GradientControllerConfig); ok {
		return x.GradientControllerConfig
	}
	return nil
}

func (m *AdaptiveConcurrency) GetEnabled() *core.RuntimeFeatureFlag {
	if m != nil {
		return m.Enabled
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdaptiveConcurrency) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdaptiveConcurrency_GradientControllerConfig)(nil),
	}
}

func init() {
	proto.RegisterType((*GradientControllerConfig)(nil), "envoy.config.filter.http.adaptive_concurrency.v2alpha.GradientControllerConfig")
	proto.RegisterType((*GradientControllerConfig_ConcurrencyLimitCalculationParams)(nil), "envoy.config.filter.http.adaptive_concurrency.v2alpha.GradientControllerConfig.ConcurrencyLimitCalculationParams")
	proto.RegisterType((*GradientControllerConfig_MinimumRTTCalculationParams)(nil), "envoy.config.filter.http.adaptive_concurrency.v2alpha.GradientControllerConfig.MinimumRTTCalculationParams")
	proto.RegisterType((*AdaptiveConcurrency)(nil), "envoy.config.filter.http.adaptive_concurrency.v2alpha.AdaptiveConcurrency")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/adaptive_concurrency/v2alpha/adaptive_concurrency.proto", fileDescriptor_c58a0beecb0ec580)
}

var fileDescriptor_c58a0beecb0ec580 = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0xd3, 0xbf, 0x68, 0xbe, 0xbf, 0xca, 0xd1, 0x07, 0x6e, 0x5a, 0x50, 0xa9, 0x40, 0x42,
	0x45, 0x1a, 0x4b, 0xa9, 0x10, 0x4b, 0x54, 0x07, 0x15, 0x8a, 0xf8, 0x89, 0x4c, 0x8b, 0xc4, 0xca,
	0x9a, 0x38, 0x37, 0xee, 0xc0, 0x78, 0x66, 0x3a, 0x19, 0x87, 0xe4, 0x15, 0xd8, 0x74, 0x5b, 0x36,
	0xac, 0x58, 0xb1, 0xe7, 0x55, 0x78, 0x03, 0x5e, 0x80, 0x15, 0xea, 0x02, 0x21, 0x7b, 0x26, 0xa9,
	0xd5, 0xa6, 0x41, 0xad, 0xba, 0x6b, 0x75, 0xef, 0x39, 0xf7, 0x9e, 0x73, 0x8f, 0x27, 0xa8, 0x05,
	0xbc, 0x2f, 0x86, 0x7e, 0x2c, 0x78, 0x97, 0x26, 0x7e, 0x97, 0x32, 0x0d, 0xca, 0xdf, 0xd7, 0x5a,
	0xfa, 0xa4, 0x43, 0xa4, 0xa6, 0x7d, 0x88, 0x62, 0xc1, 0xe3, 0x4c, 0x29, 0xe0, 0xf1, 0xd0, 0xef,
	0x37, 0x08, 0x93, 0xfb, 0x64, 0x62, 0x11, 0x4b, 0x25, 0xb4, 0x70, 0xef, 0x17, 0x8c, 0xd8, 0x30,
	0x62, 0xc3, 0x88, 0x73, 0x46, 0x3c, 0x11, 0x64, 0x19, 0xeb, 0xab, 0x66, 0x11, 0x22, 0xa9, 0xdf,
	0x6f, 0xf8, 0xb1, 0x50, 0xe0, 0xb7, 0x49, 0x0f, 0x0c, 0x69, 0xdd, 0x33, 0x55, 0x3d, 0x94, 0xe0,
	0x4b, 0x50, 0x31, 0x70, 0x6d, 0x2b, 0xab, 0x89, 0x10, 0x09, 0x83, 0x02, 0x48, 0x38, 0x17, 0x9a,
	0x68, 0x2a, 0x78, 0xcf, 0x56, 0x6f, 0xda, 0x6a, 0xf1, 0x5f, 0x3b, 0xeb, 0xfa, 0x9d, 0x4c, 0x15,
	0x0d, 0xe7, 0xd5, 0xdf, 0x2b, 0x22, 0x25, 0xa8, 0x31, 0x3e, 0xeb, 0x48, 0x52, 0xe6, 0xf5, 0x53,
	0x9a, 0x28, 0xa2, 0x47, 0x7b, 0x5d, 0xef, 0x13, 0x46, 0x3b, 0x44, 0x83, 0x3f, 0xfa, 0xc3, 0x14,
	0xd6, 0x0f, 0xab, 0xc8, 0x7b, 0xac, 0x48, 0x87, 0x02, 0xd7, 0x4d, 0xc1, 0xb5, 0x12, 0x8c, 0x81,
	0x6a, 0x16, 0xa6, 0xb8, 0xaf, 0xd0, 0x4a, 0x8f, 0xa4, 0x92, 0x41, 0x44, 0x92, 0x44, 0x41, 0x42,
	0x34, 0x44, 0x56, 0x15, 0x65, 0xe0, 0x39, 0x6b, 0xce, 0xdd, 0xbf, 0x1a, 0x35, 0x6c, 0x8c, 0xcc,
	0x35, 0xe3, 0x96, 0xa9, 0x86, 0xcb, 0x06, 0xb7, 0x35, 0x82, 0xb5, 0xc6, 0x28, 0xf7, 0xab, 0x83,
	0xbc, 0x92, 0xb1, 0x11, 0xa3, 0x29, 0xd5, 0x91, 0x24, 0x8a, 0xa4, 0x3d, 0xaf, 0x52, 0x50, 0x1e,
	0xe0, 0x4b, 0xdd, 0x06, 0x9f, 0x27, 0x04, 0x37, 0x4f, 0x9a, 0x9f, 0xe5, 0xe3, 0x9a, 0x84, 0xc5,
	0x19, 0x2b, 0x9c, 0x6a, 0x15, 0x83, 0x83, 0xea, 0x71, 0x30, 0xff, 0xc1, 0xa9, 0x2c, 0x39, 0xe1,
	0xb5, 0xf8, 0x54, 0xb3, 0xe9, 0x70, 0x3f, 0x3b, 0xa8, 0x96, 0x52, 0x1e, 0x29, 0xad, 0xa3, 0x98,
	0xb0, 0x78, 0xb4, 0xf2, 0x6c, 0xb1, 0xf2, 0xbb, 0xab, 0x5e, 0xf9, 0x39, 0xe5, 0x34, 0xcd, 0xd2,
	0x70, 0x77, 0x77, 0xda, 0xb2, 0x4b, 0x29, 0xe5, 0xa1, 0x2e, 0xf4, 0x98, 0x5a, 0xfd, 0xbb, 0x83,
	0x6e, 0xfd, 0x51, 0xae, 0xfb, 0x06, 0xfd, 0x9f, 0x92, 0x41, 0x74, 0xe6, 0x0e, 0xf6, 0x00, 0xab,
	0xd8, 0xe4, 0x0d, 0x8f, 0xf2, 0x86, 0xf7, 0x76, 0xb8, 0xde, 0x6c, 0xbc, 0x26, 0x2c, 0x83, 0x60,
	0xf1, 0x38, 0x98, 0xdb, 0xa8, 0xac, 0xcd, 0x84, 0xb5, 0x94, 0x0c, 0x4e, 0xcf, 0x72, 0x01, 0xad,
	0x94, 0x69, 0x33, 0x99, 0xa7, 0x2d, 0xa2, 0x5c, 0x83, 0xea, 0x13, 0x66, 0xed, 0x5a, 0x3e, 0x33,
	0xe0, 0x91, 0x0d, 0x7c, 0x80, 0x8e, 0x83, 0xc5, 0x2f, 0xce, 0x5c, 0xd5, 0xd9, 0x98, 0x09, 0x97,
	0x4b, 0x4c, 0x7b, 0x05, 0xd1, 0x8e, 0xe5, 0xa9, 0x7f, 0xab, 0xa0, 0x95, 0x29, 0x1e, 0xb9, 0x5b,
	0xa8, 0x3a, 0x9e, 0xe9, 0x5c, 0x64, 0xe6, 0x18, 0xe6, 0x3e, 0x45, 0xff, 0x28, 0x38, 0xc8, 0xa0,
	0xa7, 0xa3, 0x58, 0x64, 0xfc, 0x82, 0xe6, 0xfc, 0x6d, 0xb1, 0xcd, 0x1c, 0xea, 0xde, 0x43, 0x0b,
	0x6f, 0xa9, 0xd6, 0xa0, 0xac, 0x01, 0x13, 0xbf, 0x1a, 0xdb, 0xe2, 0xbe, 0x40, 0xff, 0xe5, 0x49,
	0x2b, 0x89, 0xf7, 0xe6, 0x2e, 0x32, 0xfa, 0xdf, 0x94, 0xf2, 0xd2, 0x5d, 0xf2, 0xe1, 0xed, 0xac,
	0xdb, 0x05, 0xe5, 0xcd, 0x4f, 0x19, 0x6e, 0x5a, 0xd6, 0x8f, 0x2a, 0xa8, 0xb6, 0x65, 0x23, 0x5b,
	0x26, 0xf9, 0xe8, 0xa0, 0x7a, 0x62, 0xd3, 0x9a, 0xaf, 0x66, 0xe3, 0x1a, 0x99, 0xc4, 0x5b, 0x8f,
	0x5f, 0x5e, 0xf1, 0x67, 0x70, 0x12, 0xf5, 0x27, 0x33, 0xa1, 0x97, 0x9c, 0xf7, 0x50, 0x3d, 0x44,
	0x8b, 0xc0, 0x49, 0x9b, 0x41, 0xc7, 0xde, 0xe8, 0x8e, 0xdd, 0x83, 0x48, 0x8a, 0xfb, 0x0d, 0x9c,
	0x3f, 0xd3, 0x38, 0xcc, 0xb8, 0xa6, 0x29, 0x6c, 0x03, 0xd1, 0x99, 0x82, 0x6d, 0x46, 0x92, 0x70,
	0x84, 0x0a, 0x6e, 0xa3, 0x1b, 0xe5, 0xd0, 0x9e, 0x91, 0xe7, 0xce, 0xfe, 0x0c, 0x9c, 0xe0, 0x93,
	0xf3, 0xe3, 0xe8, 0xd7, 0xe1, 0xfc, 0x83, 0xd1, 0x6f, 0x07, 0x0c, 0x34, 0xf0, 0x5e, 0xf1, 0x8c,
	0x1b, 0xa5, 0xbd, 0x69, 0x52, 0x37, 0x51, 0x93, 0x0a, 0xb3, 0x97, 0x54, 0x62, 0x30, 0xbc, 0x9c,
	0x55, 0x81, 0x37, 0xe1, 0x38, 0xad, 0x3c, 0x0e, 0x2d, 0xa7, 0xbd, 0x50, 0xe4, 0x62, 0xf3, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x34, 0xa9, 0xe3, 0x2e, 0x07, 0x00, 0x00,
}
