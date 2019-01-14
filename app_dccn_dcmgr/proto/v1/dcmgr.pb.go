// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/dcmgr.proto

package go_micro_srv_v1_dcmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_e0725f6679b5c964, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DataCenter struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               int32    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenter) Reset()         { *m = DataCenter{} }
func (m *DataCenter) String() string { return proto.CompactTextString(m) }
func (*DataCenter) ProtoMessage()    {}
func (*DataCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_e0725f6679b5c964, []int{1}
}
func (m *DataCenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenter.Unmarshal(m, b)
}
func (m *DataCenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenter.Marshal(b, m, deterministic)
}
func (dst *DataCenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenter.Merge(dst, src)
}
func (m *DataCenter) XXX_Size() int {
	return xxx_messageInfo_DataCenter.Size(m)
}
func (m *DataCenter) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenter.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenter proto.InternalMessageInfo

func (m *DataCenter) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DataCenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenter) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_e0725f6679b5c964, []int{2}
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

func init() {
	proto.RegisterType((*ID)(nil), "go.micro.srv.v1.dcmgr.ID")
	proto.RegisterType((*DataCenter)(nil), "go.micro.srv.v1.dcmgr.DataCenter")
	proto.RegisterType((*Response)(nil), "go.micro.srv.v1.dcmgr.Response")
}

func init() { proto.RegisterFile("proto/v1/dcmgr.proto", fileDescriptor_dcmgr_e0725f6679b5c964) }

var fileDescriptor_dcmgr_e0725f6679b5c964 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4e, 0x87, 0x30,
	0x10, 0x87, 0xff, 0x2d, 0x42, 0xf4, 0x06, 0x87, 0x0b, 0x1a, 0x64, 0x11, 0x3b, 0x31, 0x95, 0xa0,
	0x8f, 0x40, 0x8d, 0x32, 0xe8, 0xd0, 0x37, 0xa8, 0x70, 0x21, 0x0c, 0x50, 0xd2, 0x56, 0x1e, 0xc7,
	0x67, 0x35, 0x69, 0x34, 0x26, 0x46, 0xe2, 0x76, 0xbf, 0x7c, 0xdf, 0xf0, 0xe5, 0x20, 0xdf, 0x9c,
	0x0d, 0xb6, 0xd9, 0xdb, 0x66, 0x1c, 0x96, 0xc9, 0xc9, 0x38, 0xf1, 0x6a, 0xb2, 0x72, 0x99, 0x07,
	0x67, 0xa5, 0x77, 0xbb, 0xdc, 0x5b, 0x19, 0xa1, 0xc8, 0x81, 0xf7, 0x0a, 0x2f, 0x81, 0xcf, 0x63,
	0xc1, 0x2a, 0x56, 0x27, 0x9a, 0xcf, 0xa3, 0x78, 0x06, 0x50, 0x26, 0x98, 0x8e, 0xd6, 0x40, 0xee,
	0x37, 0x45, 0x84, 0xb3, 0xd5, 0x2c, 0x54, 0xf0, 0x8a, 0xd5, 0x17, 0x3a, 0xde, 0x78, 0x0d, 0x99,
	0x0f, 0x26, 0xbc, 0xfb, 0x22, 0xa9, 0x58, 0x9d, 0xea, 0xaf, 0x25, 0x00, 0xce, 0x35, 0xf9, 0xcd,
	0xae, 0x9e, 0xee, 0x3f, 0x18, 0xa4, 0x6a, 0x78, 0x99, 0x1c, 0x3e, 0x42, 0xf2, 0x44, 0x01, 0x6f,
	0xe4, 0x9f, 0x51, 0xb2, 0x57, 0xe5, 0xdd, 0x01, 0xfa, 0xc9, 0x12, 0x27, 0x7c, 0x85, 0xac, 0x73,
	0x64, 0x02, 0xe1, 0xff, 0x7a, 0x79, 0x7b, 0xa0, 0x7c, 0xe7, 0x89, 0xd3, 0x5b, 0x16, 0x5f, 0xf5,
	0xf0, 0x19, 0x00, 0x00, 0xff, 0xff, 0x52, 0xc4, 0xd0, 0x18, 0x42, 0x01, 0x00, 0x00,
}
