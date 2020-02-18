// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/sni_cluster/v2/sni_cluster.proto

package envoy_config_filter_network_sni_cluster_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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

type SniCluster struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SniCluster) Reset()         { *m = SniCluster{} }
func (m *SniCluster) String() string { return proto.CompactTextString(m) }
func (*SniCluster) ProtoMessage()    {}
func (*SniCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_c62c0a0da3b665be, []int{0}
}

func (m *SniCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SniCluster.Unmarshal(m, b)
}
func (m *SniCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SniCluster.Marshal(b, m, deterministic)
}
func (m *SniCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SniCluster.Merge(m, src)
}
func (m *SniCluster) XXX_Size() int {
	return xxx_messageInfo_SniCluster.Size(m)
}
func (m *SniCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_SniCluster.DiscardUnknown(m)
}

var xxx_messageInfo_SniCluster proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SniCluster)(nil), "envoy.config.filter.network.sni_cluster.v2.SniCluster")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/sni_cluster/v2/sni_cluster.proto", fileDescriptor_c62c0a0da3b665be)
}

var fileDescriptor_c62c0a0da3b665be = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0xce, 0xcb, 0x8c, 0x4f, 0xce, 0x29, 0x2d,
	0x06, 0x89, 0x95, 0x19, 0x21, 0x73, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb4, 0xc0, 0xba,
	0xf5, 0x20, 0xba, 0xf5, 0x20, 0xba, 0xf5, 0xa0, 0xba, 0xf5, 0x90, 0x95, 0x97, 0x19, 0x49, 0xc9,
	0x95, 0xa6, 0x14, 0x24, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15,
	0xeb, 0xe7, 0x66, 0xa6, 0x17, 0x25, 0x96, 0xa4, 0x42, 0xcc, 0x52, 0xe2, 0xe1, 0xe2, 0x0a, 0xce,
	0xcb, 0x74, 0x86, 0x68, 0x70, 0x6a, 0x61, 0xfc, 0x34, 0xe3, 0x5f, 0x3f, 0xab, 0xa1, 0x90, 0x3e,
	0xc4, 0x8a, 0xd4, 0x8a, 0x92, 0xd4, 0xbc, 0x62, 0x90, 0x36, 0xa8, 0x35, 0xc5, 0xd8, 0xed, 0x31,
	0xe6, 0xb2, 0xc8, 0xcc, 0xd7, 0x03, 0xeb, 0x29, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0xde, 0x85,
	0x4e, 0xfc, 0x08, 0xfb, 0x03, 0x40, 0x4e, 0x0a, 0x60, 0x4c, 0x62, 0x03, 0xbb, 0xcd, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xbe, 0x47, 0x8a, 0x1f, 0x27, 0x01, 0x00, 0x00,
}
