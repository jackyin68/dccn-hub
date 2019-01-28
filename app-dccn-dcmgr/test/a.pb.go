// Code generated by protoc-gen-go. DO NOT EDIT.
// source: a.email

package a

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the email package it is being compiled against.
// A compilation error at this line likely means your copy of the
// email package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the email package

type Request struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_a216544577412dac, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Response struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a_a216544577412dac, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("a.email", fileDescriptor_a_a216544577412dac) }

var fileDescriptor_a_a216544577412dac = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0x4f, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe1, 0x62, 0x4d, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x70,
	0x94, 0x14, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xb1, 0xab, 0x30, 0x0a,
	0xe5, 0x62, 0x75, 0x49, 0xf6, 0x4d, 0x2f, 0x12, 0x52, 0xe6, 0x62, 0x0b, 0x2e, 0x29, 0x4a, 0x4d,
	0xcc, 0x15, 0xe2, 0xd0, 0x83, 0x1a, 0x2a, 0xc5, 0xa9, 0x07, 0xd3, 0xad, 0xc4, 0xa0, 0xc1, 0x68,
	0xc0, 0x28, 0xa4, 0xce, 0xc5, 0x13, 0x9c, 0x5a, 0x54, 0x96, 0x5a, 0x84, 0x57, 0xa9, 0x01, 0x63,
	0x12, 0x1b, 0xd8, 0x81, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0xe3, 0xcd, 0x4a, 0xab,
	0x00, 0x00, 0x00,
}