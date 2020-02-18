// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/squash/v2/squash.proto

package envoy_config_filter_http_squash_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// [#next-free-field: 6]
type Squash struct {
	// The name of the cluster that hosts the Squash server.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// When the filter requests the Squash server to create a DebugAttachment, it will use this
	// structure as template for the body of the request. It can contain reference to environment
	// variables in the form of '{{ ENV_VAR_NAME }}'. These can be used to provide the Squash server
	// with more information to find the process to attach the debugger to. For example, in a
	// Istio/k8s environment, this will contain information on the pod:
	//
	// .. code-block:: json
	//
	//  {
	//    "spec": {
	//      "attachment": {
	//        "pod": "{{ POD_NAME }}",
	//        "namespace": "{{ POD_NAMESPACE }}"
	//      },
	//      "match_request": true
	//    }
	//  }
	//
	// (where POD_NAME, POD_NAMESPACE are configured in the pod via the Downward API)
	AttachmentTemplate *_struct.Struct `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate,proto3" json:"attachment_template,omitempty"`
	// The timeout for individual requests sent to the Squash cluster. Defaults to 1 second.
	RequestTimeout *duration.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// The total timeout Squash will delay a request and wait for it to be attached. Defaults to 60
	// seconds.
	AttachmentTimeout *duration.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,proto3" json:"attachment_timeout,omitempty"`
	// Amount of time to poll for the status of the attachment object in the Squash server
	// (to check if has been attached). Defaults to 1 second.
	AttachmentPollPeriod *duration.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,proto3" json:"attachment_poll_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Squash) Reset()         { *m = Squash{} }
func (m *Squash) String() string { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()    {}
func (*Squash) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fc8434388b1e13, []int{0}
}

func (m *Squash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Squash.Unmarshal(m, b)
}
func (m *Squash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Squash.Marshal(b, m, deterministic)
}
func (m *Squash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Squash.Merge(m, src)
}
func (m *Squash) XXX_Size() int {
	return xxx_messageInfo_Squash.Size(m)
}
func (m *Squash) XXX_DiscardUnknown() {
	xxx_messageInfo_Squash.DiscardUnknown(m)
}

var xxx_messageInfo_Squash proto.InternalMessageInfo

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *_struct.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *duration.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *duration.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *duration.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.config.filter.http.squash.v2.Squash")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/squash/v2/squash.proto", fileDescriptor_63fc8434388b1e13)
}

var fileDescriptor_63fc8434388b1e13 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0xef, 0xd2, 0x30,
	0x1c, 0x86, 0x33, 0xfc, 0x03, 0xb1, 0x24, 0x1a, 0xab, 0x91, 0x49, 0x0c, 0x41, 0x2e, 0xe2, 0xa5,
	0x35, 0xf0, 0x0d, 0x16, 0x0f, 0xdc, 0x5c, 0x80, 0x3b, 0x29, 0x5b, 0x19, 0x4d, 0xba, 0xb5, 0xb4,
	0xbf, 0x2e, 0xf0, 0x29, 0xbc, 0xfa, 0xc1, 0xfc, 0x24, 0x1e, 0x3d, 0x18, 0x43, 0x5b, 0x02, 0xea,
	0x81, 0x5b, 0xb3, 0xb7, 0xcf, 0x9b, 0x27, 0x7b, 0x8b, 0x28, 0x6f, 0x5a, 0x75, 0xa6, 0x85, 0x6a,
	0xf6, 0xa2, 0xa2, 0x7b, 0x21, 0x81, 0x1b, 0x7a, 0x00, 0xd0, 0xd4, 0x1e, 0x1d, 0xb3, 0x07, 0xda,
	0xce, 0xe3, 0x89, 0x68, 0xa3, 0x40, 0xe1, 0xa9, 0x07, 0x48, 0x00, 0x48, 0x00, 0xc8, 0x05, 0x20,
	0xf1, 0x5a, 0x3b, 0x1f, 0x8d, 0x2b, 0xa5, 0x2a, 0xc9, 0xa9, 0x27, 0x76, 0x6e, 0x4f, 0x4b, 0x67,
	0x18, 0x08, 0xd5, 0x84, 0x8e, 0xd1, 0xfb, 0x7f, 0x73, 0x0b, 0xc6, 0x15, 0x10, 0xd3, 0xb1, 0x2b,
	0x35, 0xa3, 0xac, 0x69, 0x14, 0x78, 0xc8, 0xd2, 0x5a, 0x54, 0x86, 0x01, 0x8f, 0xf9, 0xb0, 0x65,
	0x52, 0x94, 0x0c, 0x38, 0xbd, 0x1e, 0x42, 0x30, 0xfd, 0xd1, 0x41, 0xbd, 0xb5, 0x97, 0xc0, 0x1f,
	0x50, 0xbf, 0x90, 0xce, 0x02, 0x37, 0x69, 0x32, 0x49, 0x66, 0xcf, 0xb3, 0xfe, 0xaf, 0xec, 0xc9,
	0x74, 0x26, 0xc9, 0xea, 0xfa, 0x1d, 0x2f, 0xd1, 0x6b, 0x06, 0xc0, 0x8a, 0x43, 0xcd, 0x1b, 0xd8,
	0x02, 0xaf, 0xb5, 0x64, 0xc0, 0xd3, 0xce, 0x24, 0x99, 0x0d, 0xe6, 0x43, 0x12, 0x14, 0xc9, 0x55,
	0x91, 0xac, 0xbd, 0xe2, 0x0a, 0xdf, 0x98, 0x4d, 0x44, 0x70, 0x86, 0x5e, 0x1a, 0x7e, 0x74, 0xdc,
	0xc2, 0x16, 0x44, 0xcd, 0x95, 0x83, 0xf4, 0x99, 0x6f, 0x79, 0xf7, 0x5f, 0xcb, 0x97, 0xf8, 0x23,
	0x56, 0x2f, 0x22, 0xb1, 0x09, 0x00, 0x5e, 0x22, 0x7c, 0x6f, 0x13, 0x6b, 0x9e, 0x1e, 0xd5, 0xbc,
	0xba, 0xd3, 0x89, 0x4d, 0x5f, 0xd1, 0xdb, 0xbb, 0x26, 0xad, 0xa4, 0xdc, 0x6a, 0x6e, 0x84, 0x2a,
	0xd3, 0xee, 0xa3, 0xb6, 0x37, 0x37, 0x30, 0x57, 0x52, 0xe6, 0x1e, 0xcb, 0xf4, 0xcf, 0xef, 0xbf,
	0xbf, 0x75, 0x3f, 0xe1, 0x8f, 0x61, 0x79, 0x7e, 0x02, 0xde, 0xd8, 0xcb, 0x2e, 0x71, 0x7d, 0xfb,
	0xf7, 0xfc, 0x0b, 0xf4, 0x59, 0x28, 0xe2, 0xef, 0x6a, 0xa3, 0x4e, 0x67, 0xf2, 0xf8, 0xc1, 0x64,
	0x83, 0x30, 0x5b, 0x7e, 0x31, 0xca, 0x93, 0x5d, 0xcf, 0xab, 0x2d, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x97, 0x76, 0x9f, 0xfe, 0x9d, 0x02, 0x00, 0x00,
}
