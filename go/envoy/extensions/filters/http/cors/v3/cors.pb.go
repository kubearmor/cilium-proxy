// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/cors/v3/cors.proto

package envoy_extensions_filters_http_cors_v3

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

// Cors filter config.
type Cors struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cors) Reset()         { *m = Cors{} }
func (m *Cors) String() string { return proto.CompactTextString(m) }
func (*Cors) ProtoMessage()    {}
func (*Cors) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14462ebc0c27ba0, []int{0}
}

func (m *Cors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cors.Unmarshal(m, b)
}
func (m *Cors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cors.Marshal(b, m, deterministic)
}
func (m *Cors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cors.Merge(m, src)
}
func (m *Cors) XXX_Size() int {
	return xxx_messageInfo_Cors.Size(m)
}
func (m *Cors) XXX_DiscardUnknown() {
	xxx_messageInfo_Cors.DiscardUnknown(m)
}

var xxx_messageInfo_Cors proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Cors)(nil), "envoy.extensions.filters.http.cors.v3.Cors")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/cors/v3/cors.proto", fileDescriptor_b14462ebc0c27ba0)
}

var fileDescriptor_b14462ebc0c27ba0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0x31, 0x0a, 0x02, 0x31,
	0x10, 0x45, 0x11, 0x44, 0x30, 0xa5, 0xa5, 0x85, 0xa0, 0x20, 0x36, 0x32, 0x11, 0x63, 0x65, 0xb9,
	0x5e, 0x60, 0xaf, 0xb0, 0xae, 0xd9, 0x35, 0x20, 0x33, 0x21, 0x19, 0xc3, 0xee, 0x0d, 0x3c, 0x83,
	0xf7, 0xf1, 0x5e, 0x92, 0x71, 0x41, 0x4b, 0xab, 0x69, 0xde, 0x7b, 0xf3, 0xd5, 0xce, 0x62, 0xa2,
	0x5e, 0xdb, 0x8e, 0x2d, 0x46, 0x47, 0x18, 0x75, 0xe3, 0x6e, 0x6c, 0x43, 0xd4, 0x57, 0x66, 0xaf,
	0x6b, 0x0a, 0x51, 0x27, 0x23, 0x17, 0x7c, 0x20, 0xa6, 0xd9, 0x5a, 0x0c, 0xf8, 0x1a, 0x30, 0x18,
	0x90, 0x0d, 0x10, 0x32, 0x99, 0xf9, 0xf2, 0x7e, 0xf1, 0x95, 0xae, 0x10, 0x89, 0x2b, 0x96, 0x70,
	0xb2, 0x21, 0xf3, 0x0e, 0xdb, 0x4f, 0x69, 0x75, 0x50, 0xe3, 0x13, 0x85, 0x78, 0xdc, 0x3e, 0x5f,
	0x8f, 0xc5, 0x46, 0x0d, 0xe1, 0x9a, 0xb0, 0x71, 0xed, 0x10, 0xfd, 0x6d, 0xee, 0x21, 0xd3, 0x45,
	0xa1, 0x8c, 0x23, 0x10, 0xd6, 0x07, 0xea, 0x7a, 0xf8, 0x6b, 0x4f, 0x31, 0xcd, 0x72, 0x99, 0xff,
	0x96, 0xa3, 0xf3, 0x44, 0x06, 0x98, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xaa, 0x95, 0x94,
	0xfe, 0x00, 0x00, 0x00,
}
