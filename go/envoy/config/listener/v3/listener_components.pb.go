// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/listener/v3/listener_components.proto

package envoy_config_listener_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
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

type FilterChainMatch_ConnectionSourceType int32

const (
	// Any connection source matches.
	FilterChainMatch_ANY FilterChainMatch_ConnectionSourceType = 0
	// Match a connection originating from the same host.
	FilterChainMatch_SAME_IP_OR_LOOPBACK FilterChainMatch_ConnectionSourceType = 1
	// Match a connection originating from a different host.
	FilterChainMatch_EXTERNAL FilterChainMatch_ConnectionSourceType = 2
)

var FilterChainMatch_ConnectionSourceType_name = map[int32]string{
	0: "ANY",
	1: "SAME_IP_OR_LOOPBACK",
	2: "EXTERNAL",
}

var FilterChainMatch_ConnectionSourceType_value = map[string]int32{
	"ANY":                 0,
	"SAME_IP_OR_LOOPBACK": 1,
	"EXTERNAL":            2,
}

func (x FilterChainMatch_ConnectionSourceType) String() string {
	return proto.EnumName(FilterChainMatch_ConnectionSourceType_name, int32(x))
}

func (FilterChainMatch_ConnectionSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{1, 0}
}

type Filter struct {
	// The name of the filter to instantiate. The name must match a
	// :ref:`supported filter <config_network_filters>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*Filter_TypedConfig
	ConfigType           isFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Filter_TypedConfig) isFilter_ConfigType() {}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Filter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Filter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Filter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Filter_TypedConfig)(nil),
	}
}

// Specifies the match criteria for selecting a specific filter chain for a
// listener.
//
// In order for a filter chain to be selected, *ALL* of its criteria must be
// fulfilled by the incoming connection, properties of which are set by the
// networking stack and/or listener filters.
//
// The following order applies:
//
// 1. Destination port.
// 2. Destination IP address.
// 3. Server name (e.g. SNI for TLS protocol),
// 4. Transport protocol.
// 5. Application protocols (e.g. ALPN for TLS protocol).
// 6. Source type (e.g. any, local or external network).
// 7. Source IP address.
// 8. Source port.
//
// For criteria that allow ranges or wildcards, the most specific value in any
// of the configured filter chains that matches the incoming connection is going
// to be used (e.g. for SNI ``www.example.com`` the most specific match would be
// ``www.example.com``, then ``*.example.com``, then ``*.com``, then any filter
// chain without ``server_names`` requirements).
//
// [#comment: Implemented rules are kept in the preference order, with deprecated fields
// listed at the end, because that's how we want to list them in the docs.
//
// [#comment:TODO(PiotrSikora): Add support for configurable precedence of the rules]
// [#next-free-field: 13]
type FilterChainMatch struct {
	// Optional destination port to consider when use_original_dst is set on the
	// listener in determining a filter chain match.
	DestinationPort *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	// If non-empty, an IP address and prefix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	PrefixRanges []*v3.CidrRange `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	// If non-empty, an IP address and suffix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	// [#not-implemented-hide:]
	AddressSuffix string `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix,proto3" json:"address_suffix,omitempty"`
	// [#not-implemented-hide:]
	SuffixLen *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen,proto3" json:"suffix_len,omitempty"`
	// Specifies the connection source IP match type. Can be any, local or external network.
	SourceType FilterChainMatch_ConnectionSourceType `protobuf:"varint,12,opt,name=source_type,json=sourceType,proto3,enum=envoy.config.listener.v3.FilterChainMatch_ConnectionSourceType" json:"source_type,omitempty"`
	// The criteria is satisfied if the source IP address of the downstream
	// connection is contained in at least one of the specified subnets. If the
	// parameter is not specified or the list is empty, the source IP address is
	// ignored.
	SourcePrefixRanges []*v3.CidrRange `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	// The criteria is satisfied if the source port of the downstream connection
	// is contained in at least one of the specified ports. If the parameter is
	// not specified, the source port is ignored.
	SourcePorts []uint32 `protobuf:"varint,7,rep,packed,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	// If non-empty, a list of server names (e.g. SNI for TLS protocol) to consider when determining
	// a filter chain match. Those values will be compared against the server names of a new
	// connection, when detected by one of the listener filters.
	//
	// The server name will be matched against all wildcard domains, i.e. ``www.example.com``
	// will be first matched against ``www.example.com``, then ``*.example.com``, then ``*.com``.
	//
	// Note that partial wildcards are not supported, and values like ``*w.example.com`` are invalid.
	//
	// .. attention::
	//
	//   See the :ref:`FAQ entry <faq_how_to_setup_sni>` on how to configure SNI for more
	//   information.
	ServerNames []string `protobuf:"bytes,11,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	// If non-empty, a transport protocol to consider when determining a filter chain match.
	// This value will be compared against the transport protocol of a new connection, when
	// it's detected by one of the listener filters.
	//
	// Suggested values include:
	//
	// * ``raw_buffer`` - default, used when no transport protocol is detected,
	// * ``tls`` - set by :ref:`envoy.listener.tls_inspector <config_listener_filters_tls_inspector>`
	//   when TLS protocol is detected.
	TransportProtocol string `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	// If non-empty, a list of application protocols (e.g. ALPN for TLS protocol) to consider when
	// determining a filter chain match. Those values will be compared against the application
	// protocols of a new connection, when detected by one of the listener filters.
	//
	// Suggested values include:
	//
	// * ``http/1.1`` - set by :ref:`envoy.listener.tls_inspector
	//   <config_listener_filters_tls_inspector>`,
	// * ``h2`` - set by :ref:`envoy.listener.tls_inspector <config_listener_filters_tls_inspector>`
	//
	// .. attention::
	//
	//   Currently, only :ref:`TLS Inspector <config_listener_filters_tls_inspector>` provides
	//   application protocol detection based on the requested
	//   `ALPN <https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation>`_ values.
	//
	//   However, the use of ALPN is pretty much limited to the HTTP/2 traffic on the Internet,
	//   and matching on values other than ``h2`` is going to lead to a lot of false negatives,
	//   unless all connecting clients are known to use ALPN.
	ApplicationProtocols []string `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterChainMatch) Reset()         { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()    {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{1}
}

func (m *FilterChainMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChainMatch.Unmarshal(m, b)
}
func (m *FilterChainMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChainMatch.Marshal(b, m, deterministic)
}
func (m *FilterChainMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChainMatch.Merge(m, src)
}
func (m *FilterChainMatch) XXX_Size() int {
	return xxx_messageInfo_FilterChainMatch.Size(m)
}
func (m *FilterChainMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChainMatch.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChainMatch proto.InternalMessageInfo

func (m *FilterChainMatch) GetDestinationPort() *wrappers.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*v3.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourceType() FilterChainMatch_ConnectionSourceType {
	if m != nil {
		return m.SourceType
	}
	return FilterChainMatch_ANY
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*v3.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []uint32 {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetServerNames() []string {
	if m != nil {
		return m.ServerNames
	}
	return nil
}

func (m *FilterChainMatch) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *FilterChainMatch) GetApplicationProtocols() []string {
	if m != nil {
		return m.ApplicationProtocols
	}
	return nil
}

// A filter chain wraps a set of match criteria, an option TLS context, a set of filters, and
// various other parameters.
// [#next-free-field: 8]
type FilterChain struct {
	// The criteria to use when matching a connection to this filter chain.
	FilterChainMatch *FilterChainMatch `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	// A list of individual network filters that make up the filter chain for
	// connections established with the listener. Order matters as the filters are
	// processed sequentially as connection events happen. Note: If the filter
	// list is empty, the connection will close by default.
	Filters []*Filter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// Whether the listener should expect a PROXY protocol V1 header on new
	// connections. If this option is enabled, the listener will assume that that
	// remote address of the connection is the one specified in the header. Some
	// load balancers including the AWS ELB support this option. If the option is
	// absent or set to false, Envoy will use the physical peer address of the
	// connection as the remote address.
	UseProxyProto *wrappers.BoolValue `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	// [#not-implemented-hide:] filter chain metadata.
	Metadata *v3.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Optional custom transport socket implementation to use for downstream connections.
	// To setup TLS, set a transport socket with name `tls` and
	// :ref:`DownstreamTlsContext <envoy_api_msg_extensions.transport_sockets.tls.v3.DownstreamTlsContext>` in the `typed_config`.
	// If no transport socket configuration is specified, new connections
	// will be set up with plaintext.
	TransportSocket *v3.TransportSocket `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	// [#not-implemented-hide:] The unique name (or empty) by which this filter chain is known. If no
	// name is provided, Envoy will allocate an internal UUID for the filter chain. If the filter
	// chain is to be dynamically updated or removed via FCDS a unique name must be provided.
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterChain) Reset()         { *m = FilterChain{} }
func (m *FilterChain) String() string { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()    {}
func (*FilterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{2}
}

func (m *FilterChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChain.Unmarshal(m, b)
}
func (m *FilterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChain.Marshal(b, m, deterministic)
}
func (m *FilterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChain.Merge(m, src)
}
func (m *FilterChain) XXX_Size() int {
	return xxx_messageInfo_FilterChain.Size(m)
}
func (m *FilterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChain.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChain proto.InternalMessageInfo

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *wrappers.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *v3.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *v3.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func (m *FilterChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListenerFilter struct {
	// The name of the filter to instantiate. The name must match a
	// :ref:`supported filter <config_listener_filters>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being instantiated.
	// See the supported filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*ListenerFilter_TypedConfig
	ConfigType           isListenerFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListenerFilter) Reset()         { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()    {}
func (*ListenerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{3}
}

func (m *ListenerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilter.Unmarshal(m, b)
}
func (m *ListenerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilter.Marshal(b, m, deterministic)
}
func (m *ListenerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilter.Merge(m, src)
}
func (m *ListenerFilter) XXX_Size() int {
	return xxx_messageInfo_ListenerFilter.Size(m)
}
func (m *ListenerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilter proto.InternalMessageInfo

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListenerFilter_ConfigType interface {
	isListenerFilter_ConfigType()
}

type ListenerFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ListenerFilter_TypedConfig) isListenerFilter_ConfigType() {}

func (m *ListenerFilter) GetConfigType() isListenerFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ListenerFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ListenerFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilter_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.config.listener.v3.FilterChainMatch_ConnectionSourceType", FilterChainMatch_ConnectionSourceType_name, FilterChainMatch_ConnectionSourceType_value)
	proto.RegisterType((*Filter)(nil), "envoy.config.listener.v3.Filter")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.config.listener.v3.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.config.listener.v3.FilterChain")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.config.listener.v3.ListenerFilter")
}

func init() {
	proto.RegisterFile("envoy/config/listener/v3/listener_components.proto", fileDescriptor_87f255d2eccc91b5)
}

var fileDescriptor_87f255d2eccc91b5 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x72, 0xe3, 0x44,
	0x10, 0x5e, 0x59, 0x5e, 0xff, 0xb4, 0xec, 0x8d, 0x18, 0x42, 0x45, 0x84, 0x54, 0xd6, 0xf1, 0xb2,
	0x29, 0xd7, 0x52, 0x2b, 0x57, 0xd9, 0x27, 0xcc, 0x81, 0xb2, 0x4c, 0xb6, 0xd8, 0x90, 0x1f, 0xa3,
	0x04, 0x6a, 0x39, 0xa9, 0x26, 0xf2, 0xd8, 0x2b, 0x50, 0x66, 0x54, 0x33, 0x63, 0x13, 0x5f, 0x39,
	0x51, 0x1c, 0x39, 0xf2, 0x04, 0x9c, 0x79, 0x07, 0x78, 0xa8, 0x5c, 0x96, 0xd2, 0x68, 0xe4, 0xd8,
	0x89, 0x97, 0xa4, 0xb8, 0x69, 0xba, 0xbf, 0xee, 0xe9, 0xf9, 0xfa, 0xeb, 0x16, 0x74, 0x08, 0x9d,
	0xb1, 0x79, 0x3b, 0x64, 0x74, 0x1c, 0x4d, 0xda, 0x71, 0x24, 0x24, 0xa1, 0x84, 0xb7, 0x67, 0xdd,
	0xc5, 0x77, 0x10, 0xb2, 0xcb, 0x84, 0x51, 0x42, 0xa5, 0x70, 0x13, 0xce, 0x24, 0x43, 0x8e, 0x8a,
	0x71, 0xb3, 0x18, 0x37, 0xc7, 0xb9, 0xb3, 0xee, 0x76, 0x73, 0x25, 0x5b, 0xc8, 0x38, 0x49, 0x33,
	0xe1, 0xd1, 0x88, 0x13, 0xa1, 0xa3, 0xb7, 0x9f, 0xae, 0xc5, 0x5c, 0x60, 0x41, 0x34, 0xe0, 0xe3,
	0x09, 0x63, 0x93, 0x98, 0xb4, 0xd5, 0xe9, 0x62, 0x3a, 0x6e, 0x63, 0x3a, 0xd7, 0xae, 0x9d, 0xdb,
	0x2e, 0x21, 0xf9, 0x34, 0x94, 0xda, 0xbb, 0x7b, 0xdb, 0xfb, 0x33, 0xc7, 0x49, 0x42, 0x78, 0x7e,
	0xf3, 0xde, 0x74, 0x94, 0xe0, 0x36, 0xa6, 0x94, 0x49, 0x2c, 0x23, 0x46, 0x45, 0x7b, 0x46, 0xb8,
	0x88, 0x18, 0x8d, 0xe8, 0x44, 0x43, 0xb6, 0x66, 0x38, 0x8e, 0x46, 0x58, 0x92, 0x76, 0xfe, 0x91,
	0x39, 0x9a, 0x7f, 0x1a, 0x50, 0x7a, 0x15, 0xc5, 0x92, 0x70, 0xf4, 0x09, 0x14, 0x29, 0xbe, 0x24,
	0x8e, 0xd1, 0x30, 0x5a, 0x55, 0xaf, 0x7c, 0xed, 0x15, 0x79, 0xa1, 0x61, 0xf8, 0xca, 0x88, 0x3e,
	0x87, 0x9a, 0x9c, 0x27, 0x64, 0x14, 0x64, 0xef, 0x73, 0x8a, 0x0d, 0xa3, 0x65, 0x75, 0x36, 0xdd,
	0xac, 0x34, 0x37, 0x2f, 0xcd, 0xed, 0xd3, 0xf9, 0xd7, 0x8f, 0x7c, 0x4b, 0x61, 0x07, 0x0a, 0xda,
	0x7b, 0xf6, 0xc7, 0xdf, 0xbf, 0xee, 0xee, 0xc2, 0x4e, 0xc6, 0x2e, 0x4e, 0x22, 0x77, 0xd6, 0xb9,
	0x61, 0x37, 0xbb, 0xdc, 0xab, 0x83, 0x95, 0x65, 0x0e, 0xd2, 0xd0, 0xc3, 0x62, 0xc5, 0xb4, 0x8b,
	0x87, 0xc5, 0x4a, 0xc1, 0x36, 0xfd, 0x52, 0xe6, 0x68, 0xfe, 0x56, 0x02, 0x3b, 0x43, 0x0f, 0xde,
	0xe2, 0x88, 0x1e, 0x63, 0x19, 0xbe, 0x45, 0xe7, 0x60, 0x8f, 0x88, 0x90, 0x11, 0x55, 0x2f, 0x0f,
	0x12, 0xc6, 0xa5, 0x53, 0x51, 0xb5, 0xed, 0xdc, 0xa9, 0xed, 0xbb, 0xd7, 0x54, 0x76, 0x3b, 0xdf,
	0xe3, 0x78, 0x4a, 0x3c, 0xeb, 0xda, 0xab, 0xbc, 0x28, 0x39, 0xef, 0xde, 0x99, 0x2d, 0xc3, 0xdf,
	0x58, 0x4a, 0x31, 0x64, 0x5c, 0xa2, 0xaf, 0xa0, 0x9e, 0x70, 0x32, 0x8e, 0xae, 0x02, 0x8e, 0xe9,
	0x84, 0x08, 0xc7, 0x6c, 0x98, 0x2d, 0xab, 0xf3, 0xd4, 0x5d, 0x51, 0x48, 0xda, 0x63, 0x77, 0xd6,
	0x75, 0x07, 0xd1, 0x88, 0xfb, 0x29, 0xce, 0xaf, 0x65, 0x51, 0xea, 0x20, 0xd0, 0x73, 0x78, 0xa2,
	0x25, 0x12, 0x88, 0xe9, 0x78, 0x1c, 0x5d, 0x29, 0xd6, 0xaa, 0x7e, 0x5d, 0x5b, 0xcf, 0x94, 0x11,
	0x7d, 0x01, 0x90, 0xb9, 0x83, 0x98, 0x50, 0xe7, 0xf1, 0xfd, 0xc5, 0xfb, 0xd5, 0x0c, 0x7f, 0x44,
	0x28, 0xfa, 0x11, 0x2c, 0xc1, 0xa6, 0x3c, 0x24, 0x8a, 0x37, 0xa7, 0xd6, 0x30, 0x5a, 0x4f, 0x3a,
	0x5f, 0xba, 0xef, 0x53, 0xb2, 0x7b, 0x9b, 0x40, 0x77, 0xc0, 0x28, 0x25, 0x61, 0xfa, 0xf2, 0x33,
	0x95, 0xe7, 0x7c, 0x9e, 0x10, 0xaf, 0x72, 0xed, 0x3d, 0xfe, 0xc5, 0x28, 0xd8, 0x86, 0x0f, 0x62,
	0x61, 0x45, 0xdf, 0xc2, 0xa6, 0xbe, 0x6b, 0x95, 0x9c, 0xd2, 0xc3, 0xc8, 0x41, 0x59, 0xf0, 0x70,
	0x99, 0xa2, 0x2e, 0xd4, 0xf2, 0x94, 0x8c, 0x4b, 0xe1, 0x94, 0x1b, 0x66, 0xab, 0xee, 0xd9, 0xd7,
	0x5e, 0xfd, 0x77, 0x03, 0x9a, 0x37, 0x1d, 0xd2, 0x8f, 0x4c, 0x9b, 0x23, 0xd0, 0x1e, 0xd4, 0x04,
	0xe1, 0x33, 0xc2, 0x83, 0x54, 0x9a, 0xc2, 0xb1, 0x1a, 0x66, 0xab, 0xea, 0x5b, 0x99, 0xed, 0x24,
	0x35, 0xa1, 0x97, 0x80, 0x24, 0xc7, 0x54, 0xa4, 0x59, 0x03, 0xc5, 0x61, 0xc8, 0x62, 0xa7, 0xaa,
	0xe8, 0xff, 0x60, 0xe1, 0x19, 0x6a, 0x07, 0xea, 0xc2, 0x47, 0x38, 0x49, 0xe2, 0x28, 0xd4, 0x2a,
	0xd2, 0x76, 0xe1, 0x80, 0x4a, 0xbd, 0xb9, 0xe4, 0xcc, 0x63, 0x44, 0xf3, 0x15, 0x6c, 0xae, 0x23,
	0x0f, 0x95, 0xc1, 0xec, 0x9f, 0xfc, 0x60, 0x3f, 0x42, 0x5b, 0xf0, 0xe1, 0x59, 0xff, 0xf8, 0x20,
	0x78, 0x3d, 0x0c, 0x4e, 0xfd, 0xe0, 0xe8, 0xf4, 0x74, 0xe8, 0xf5, 0x07, 0xdf, 0xd8, 0x06, 0xaa,
	0x41, 0xe5, 0xe0, 0xcd, 0xf9, 0x81, 0x7f, 0xd2, 0x3f, 0xb2, 0x0b, 0xbd, 0x97, 0xe9, 0x7c, 0xb4,
	0x60, 0xff, 0xbf, 0xe6, 0xe3, 0xa6, 0x61, 0x87, 0xc5, 0x8a, 0x61, 0x17, 0x9a, 0xff, 0x98, 0x60,
	0x2d, 0xb9, 0xd0, 0x1b, 0x40, 0x63, 0x75, 0x0c, 0xc2, 0xf4, 0x1c, 0x5c, 0xa6, 0x58, 0x35, 0xca,
	0x56, 0xe7, 0xc5, 0xc3, 0xe5, 0xe0, 0xdb, 0xe3, 0xdb, 0x13, 0xd6, 0x83, 0x72, 0x66, 0xcb, 0xa7,
	0xa0, 0x71, 0x5f, 0x3a, 0x3f, 0x0f, 0x40, 0x1e, 0x6c, 0x4c, 0x45, 0x2a, 0x17, 0x76, 0x35, 0xcf,
	0x58, 0xd5, 0x8b, 0x63, 0xfb, 0x8e, 0xbe, 0x3d, 0xc6, 0xe2, 0x4c, 0xdd, 0xf5, 0xa9, 0x20, 0xc3,
	0x34, 0x42, 0x51, 0x8d, 0x7a, 0x50, 0xb9, 0x24, 0x12, 0x8f, 0xb0, 0xc4, 0x7a, 0x38, 0x76, 0xd7,
	0x2b, 0xed, 0x58, 0xa3, 0xfc, 0x05, 0x1e, 0x0d, 0xc1, 0xbe, 0x91, 0x81, 0x60, 0xe1, 0x4f, 0x44,
	0x3a, 0x25, 0x95, 0xe3, 0xf9, 0xfa, 0x1c, 0xe7, 0x39, 0xfa, 0x4c, 0x81, 0xfd, 0x0d, 0xb9, 0x6a,
	0x40, 0x48, 0x2f, 0xc9, 0xb2, 0x92, 0x92, 0xfa, 0xee, 0xb5, 0xd2, 0x06, 0x3e, 0x83, 0xbd, 0x7b,
	0x1b, 0xa8, 0x17, 0x9a, 0x25, 0x63, 0x91, 0xee, 0x51, 0x49, 0xae, 0x64, 0xf3, 0x2f, 0x03, 0x9e,
	0x1c, 0x69, 0xec, 0xff, 0x59, 0xc4, 0xe6, 0xc3, 0x17, 0xf1, 0x67, 0x69, 0x9d, 0xfb, 0xf0, 0xe9,
	0xfa, 0x3a, 0x57, 0x8b, 0xb8, 0xbb, 0x90, 0x97, 0x56, 0xb1, 0x37, 0x80, 0xfd, 0x88, 0x65, 0x0c,
	0xaa, 0xde, 0xbe, 0x57, 0x11, 0xde, 0x56, 0x9e, 0x76, 0xb0, 0xf8, 0xdb, 0xaa, 0xb6, 0x0e, 0x8d,
	0x8b, 0x92, 0xaa, 0xb5, 0xfb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x6e, 0x30, 0xc1, 0xab,
	0x07, 0x00, 0x00,
}
