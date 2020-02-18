// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto

package envoy_config_filter_network_thrift_proxy_v2alpha1

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
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

type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Types that are valid to be assigned to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	// Inverts whatever matching is done in the :ref:`method_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.method_name>` or
	// :ref:`service_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.service_name>` fields.
	// Cannot be combined with wildcard matching as that would result in routes never being matched.
	//
	// .. note::
	//
	//   This does not invert matching done as part of the :ref:`headers field
	//   <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteMatch.headers>` field. To
	//   invert header matching, see :ref:`invert_match
	//   <envoy_api_field_route.HeaderMatcher.invert_match>`.
	Invert bool `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config). Note that this only applies for Thrift transports and/or
	// protocols that support headers.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (m *RouteMatch) GetMethodName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (m *RouteMatch) GetServiceName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (m *RouteMatch) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
}

// [#next-free-field: 7]
type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	//	*RouteAction_ClusterHeader
	ClusterSpecifier isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field will be considered.
	// Note that this will be merged with what's provided in :ref:`WeightedCluster.metadata_match
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight.metadata_match>`,
	// with values there taking precedence. Keys and values should be provided under the "envoy.lb"
	// metadata key.
	MetadataMatch *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	// Specifies a set of rate limit configurations that could be applied to the route.
	// N.B. Thrift service or method name matching can be achieved by specifying a RequestHeaders
	// action with the header name ":method-name".
	RateLimits []*route.RateLimit `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Strip the service prefix from the method name, if there's a prefix. For
	// example, the method call Service:method would end up being just method.
	StripServiceName     bool     `protobuf:"varint,5,opt,name=strip_service_name,json=stripServiceName,proto3" json:"strip_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

type RouteAction_ClusterHeader struct {
	ClusterHeader string `protobuf:"bytes,6,opt,name=cluster_header,json=clusterHeader,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_ClusterHeader) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *RouteAction) GetClusterHeader() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_ClusterHeader); ok {
		return x.ClusterHeader
	}
	return ""
}

func (m *RouteAction) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *RouteAction) GetRateLimits() []*route.RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *RouteAction) GetStripServiceName() bool {
	if m != nil {
		return m.StripServiceName
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
		(*RouteAction_ClusterHeader)(nil),
	}
}

// Allows for specification of multiple upstream clusters along with weights that indicate the
// percentage of traffic to be forwarded to each cluster. The router selects an upstream cluster
// based on these weights.
type WeightedCluster struct {
	// Specifies one or more upstream clusters associated with the route.
	Clusters             []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{4}
}

func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(m, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	// Name of the upstream cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When a request matches the route, the choice of an upstream cluster is determined by its
	// weight. The sum of weights across all entries in the clusters array determines the total
	// weight.
	Weight *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field, combined with what's
	// provided in :ref:`RouteAction's metadata_match
	// <envoy_api_field_config.filter.network.thrift_proxy.v2alpha1.RouteAction.metadata_match>`,
	// will be considered. Values here will take precedence. Keys and values should be provided
	// under the "envoy.lb" metadata key.
	MetadataMatch        *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WeightedCluster_ClusterWeight) Reset()         { *m = WeightedCluster_ClusterWeight{} }
func (m *WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_3de6ac0eae6369f5, []int{4, 0}
}

func (m *WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Size(m)
}
func (m *WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeight) GetWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction")
	proto.RegisterType((*WeightedCluster)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto", fileDescriptor_3de6ac0eae6369f5)
}

var fileDescriptor_3de6ac0eae6369f5 = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xee, 0x38, 0x4d, 0x9a, 0x7b, 0x72, 0xfb, 0x37, 0x8b, 0xde, 0x28, 0xed, 0x45, 0x69, 0xba,
	0x09, 0x08, 0x8d, 0xdb, 0x74, 0x83, 0x84, 0x5a, 0x84, 0xbb, 0x09, 0x12, 0x45, 0x91, 0x11, 0xb0,
	0x02, 0x6b, 0xea, 0x4c, 0x92, 0x11, 0x89, 0xc7, 0x8c, 0x27, 0x4e, 0xcb, 0x0b, 0x20, 0xb1, 0x41,
	0x62, 0xc5, 0x9e, 0x07, 0xe0, 0x01, 0x58, 0xb1, 0xe1, 0x81, 0x58, 0xa1, 0x2e, 0x10, 0xf2, 0xcc,
	0x38, 0x4d, 0x2a, 0x58, 0xb4, 0xb0, 0x49, 0xec, 0x33, 0xe7, 0xfb, 0xbe, 0x73, 0xbe, 0x33, 0x3e,
	0x70, 0xc0, 0xa2, 0x54, 0x9c, 0xb9, 0xa1, 0x88, 0x7a, 0xbc, 0xef, 0xf6, 0xf8, 0x50, 0x31, 0xe9,
	0x46, 0x4c, 0x4d, 0x84, 0x7c, 0xe9, 0xaa, 0x81, 0xe4, 0x3d, 0x15, 0xc4, 0x52, 0x9c, 0x9e, 0xb9,
	0x69, 0x8b, 0x0e, 0xe3, 0x01, 0xdd, 0x73, 0xa5, 0x18, 0x2b, 0x46, 0x62, 0x29, 0x94, 0xc0, 0x7b,
	0x1a, 0x4e, 0x0c, 0x9c, 0x18, 0x38, 0xb1, 0x70, 0x32, 0x0b, 0x27, 0x39, 0xbc, 0xb6, 0x65, 0x14,
	0x69, 0xcc, 0xdd, 0xb4, 0xe5, 0x86, 0x42, 0x32, 0xf7, 0x84, 0x26, 0x96, 0xb0, 0x76, 0x73, 0xee,
	0x54, 0x4b, 0x99, 0xdf, 0x20, 0x14, 0xa3, 0x58, 0x44, 0x2c, 0x52, 0x89, 0x4d, 0xbd, 0xd1, 0x17,
	0xa2, 0x3f, 0x64, 0xae, 0x7e, 0x3b, 0x19, 0xf7, 0xdc, 0x89, 0xa4, 0x71, 0xcc, 0xe4, 0xf4, 0x7c,
	0xdc, 0x8d, 0xa9, 0x4b, 0xa3, 0x48, 0x28, 0xaa, 0xb8, 0x88, 0x12, 0x77, 0xc4, 0xfb, 0x92, 0xe6,
	0xb5, 0xd7, 0xfe, 0x4b, 0xe9, 0x90, 0x77, 0xa9, 0x62, 0x6e, 0xfe, 0x60, 0x0e, 0x1a, 0xaf, 0x01,
	0xfb, 0x99, 0xe4, 0x91, 0xee, 0x6a, 0x2c, 0x35, 0x1a, 0x63, 0x58, 0x8c, 0xe8, 0x88, 0x55, 0x51,
	0x1d, 0x35, 0xff, 0xf1, 0xf5, 0x33, 0xee, 0x40, 0x49, 0x17, 0x97, 0x54, 0x9d, 0x7a, 0xa1, 0x59,
	0x69, 0xdd, 0x21, 0x57, 0xf6, 0x83, 0x68, 0x29, 0xdf, 0xf2, 0x34, 0xbe, 0x22, 0x28, 0xea, 0x08,
	0x7e, 0x0e, 0xc5, 0x11, 0x55, 0xe1, 0x40, 0x0b, 0x56, 0x5a, 0x07, 0xd7, 0xa5, 0x3e, 0xce, 0x48,
	0xbc, 0xf2, 0xb9, 0x57, 0x7c, 0x8b, 0x9c, 0x35, 0xe4, 0x1b, 0x56, 0xfc, 0x02, 0x8a, 0x5a, 0xb2,
	0xea, 0x68, 0xfa, 0xc3, 0xeb, 0xd2, 0xdf, 0x0f, 0x33, 0x77, 0x66, 0xf9, 0x35, 0x6d, 0xe3, 0x0b,
	0x02, 0xb8, 0xd0, 0xc7, 0xdb, 0x50, 0x19, 0x31, 0x35, 0x10, 0xdd, 0xe0, 0xc2, 0xc4, 0xf6, 0x82,
	0x0f, 0x26, 0xf8, 0x28, 0x33, 0x73, 0x07, 0xfe, 0x4d, 0x98, 0x4c, 0x79, 0xc8, 0x4c, 0x8e, 0x63,
	0x73, 0x2a, 0x36, 0xaa, 0x93, 0x36, 0xa0, 0xc4, 0xa3, 0x94, 0x49, 0x55, 0x2d, 0xd4, 0x51, 0xb3,
	0xec, 0xdb, 0x37, 0x7c, 0x17, 0x96, 0x06, 0x8c, 0x76, 0x99, 0x4c, 0xaa, 0x8b, 0x7a, 0x14, 0xdb,
	0xb6, 0x21, 0x1a, 0x73, 0x92, 0xb6, 0x88, 0xb9, 0xb4, 0x6d, 0x9d, 0xa2, 0x2b, 0x62, 0xd2, 0xcf,
	0x11, 0xde, 0x06, 0xac, 0x6a, 0x53, 0x82, 0x24, 0x66, 0x21, 0xef, 0x71, 0x26, 0x71, 0xe1, 0xbb,
	0x87, 0x1a, 0x1f, 0x0b, 0x50, 0x99, 0x69, 0x12, 0xef, 0xc0, 0x52, 0x38, 0x1c, 0x27, 0x8a, 0x49,
	0xd3, 0x80, 0xb7, 0x74, 0xee, 0x2d, 0x4a, 0xa7, 0x8e, 0xda, 0x0b, 0x7e, 0x7e, 0x82, 0x5f, 0xc1,
	0xfa, 0x84, 0xf1, 0xfe, 0x40, 0xb1, 0x6e, 0x60, 0x63, 0x89, 0x35, 0xd9, 0xbb, 0x86, 0xc9, 0xcf,
	0x2c, 0xd7, 0x91, 0xa1, 0x6a, 0x2f, 0xf8, 0x6b, 0x93, 0xf9, 0x50, 0x82, 0x77, 0x61, 0xc5, 0x2a,
	0x05, 0xa6, 0xa5, 0x6a, 0xe9, 0x72, 0x79, 0xcb, 0x36, 0xc1, 0x38, 0x80, 0x3d, 0x58, 0x19, 0x31,
	0x45, 0xbb, 0x54, 0xd1, 0xc0, 0xdc, 0xb2, 0x82, 0xae, 0x70, 0x73, 0xde, 0xb5, 0xec, 0xeb, 0x24,
	0xc7, 0x36, 0xd1, 0x5f, 0xce, 0x21, 0x66, 0xa4, 0x87, 0x50, 0xc9, 0xbe, 0xa6, 0x60, 0xc8, 0x47,
	0x5c, 0xe5, 0xb6, 0xff, 0xff, 0x2b, 0xdb, 0x7d, 0xaa, 0xd8, 0xc3, 0x2c, 0xcb, 0x07, 0x99, 0x3f,
	0x26, 0xf8, 0x36, 0xe0, 0x44, 0x49, 0x1e, 0x07, 0x73, 0x53, 0x2f, 0xea, 0xb1, 0xae, 0xe9, 0x93,
	0xc7, 0x17, 0x83, 0xf7, 0xaa, 0xb0, 0x9e, 0xf7, 0x78, 0x69, 0x4a, 0x9f, 0x1d, 0x58, 0xbd, 0xe4,
	0x12, 0x4e, 0xa1, 0x3c, 0xf5, 0x1e, 0xe9, 0xc2, 0x3a, 0x7f, 0xee, 0x3d, 0xb1, 0xff, 0x26, 0xac,
	0xaf, 0xfc, 0x7b, 0xe4, 0x94, 0x91, 0x3f, 0xd5, 0xaa, 0x7d, 0x42, 0xb0, 0x3c, 0x97, 0x85, 0x37,
	0x67, 0xd7, 0xc6, 0x74, 0x22, 0x76, 0x7f, 0x1c, 0x40, 0xc9, 0x0c, 0xd3, 0x5e, 0x90, 0x2d, 0x62,
	0x76, 0x1a, 0xc9, 0x77, 0x1a, 0x79, 0xf2, 0x20, 0x52, 0xfb, 0xad, 0xa7, 0x74, 0x38, 0x66, 0x1a,
	0x7c, 0xcb, 0x69, 0x22, 0xdf, 0x82, 0xfe, 0xc6, 0x14, 0xbd, 0x37, 0xe8, 0xdb, 0x87, 0x1f, 0xef,
	0x8a, 0x2d, 0xbc, 0x6b, 0x30, 0xec, 0x54, 0xb1, 0x28, 0xc9, 0xd6, 0xa5, 0xf5, 0x28, 0xf9, 0x8d,
	0x49, 0xfb, 0x70, 0x8f, 0x0b, 0x23, 0x64, 0x22, 0x57, 0xf6, 0xd7, 0x33, 0x0b, 0xa2, 0x93, 0xf5,
	0xda, 0x41, 0x27, 0x25, 0xdd, 0xf4, 0xfe, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0xf9, 0x9b,
	0x4f, 0x93, 0x06, 0x00, 0x00,
}
