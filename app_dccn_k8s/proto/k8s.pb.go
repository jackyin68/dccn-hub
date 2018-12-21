// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/k8s.proto

package k8s

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Event struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b74d577d0586ba, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type Task struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Taskid               int64    `protobuf:"varint,2,opt,name=taskid,proto3" json:"taskid,omitempty"`
	TaskType             string   `protobuf:"bytes,3,opt,name=taskType,proto3" json:"taskType,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Replica              int64    `protobuf:"varint,6,opt,name=replica,proto3" json:"replica,omitempty"`
	Extra                string   `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b74d577d0586ba, []int{1}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Task) GetTaskid() int64 {
	if m != nil {
		return m.Taskid
	}
	return 0
}

func (m *Task) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Task) GetReplica() int64 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *Task) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

type K8SMessage struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Taskid               int64    `protobuf:"varint,2,opt,name=taskid,proto3" json:"taskid,omitempty"`
	Taskname             string   `protobuf:"bytes,3,opt,name=taskname,proto3" json:"taskname,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Datacenter           string   `protobuf:"bytes,5,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	Url                  string   `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Report               string   `protobuf:"bytes,7,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K8SMessage) Reset()         { *m = K8SMessage{} }
func (m *K8SMessage) String() string { return proto.CompactTextString(m) }
func (*K8SMessage) ProtoMessage()    {}
func (*K8SMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b74d577d0586ba, []int{2}
}

func (m *K8SMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SMessage.Unmarshal(m, b)
}
func (m *K8SMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SMessage.Marshal(b, m, deterministic)
}
func (m *K8SMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SMessage.Merge(m, src)
}
func (m *K8SMessage) XXX_Size() int {
	return xxx_messageInfo_K8SMessage.Size(m)
}
func (m *K8SMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SMessage.DiscardUnknown(m)
}

var xxx_messageInfo_K8SMessage proto.InternalMessageInfo

func (m *K8SMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *K8SMessage) GetTaskid() int64 {
	if m != nil {
		return m.Taskid
	}
	return 0
}

func (m *K8SMessage) GetTaskname() string {
	if m != nil {
		return m.Taskname
	}
	return ""
}

func (m *K8SMessage) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *K8SMessage) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

func (m *K8SMessage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *K8SMessage) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "k8s.Event")
	proto.RegisterType((*Task)(nil), "k8s.Task")
	proto.RegisterType((*K8SMessage)(nil), "k8s.K8SMessage")
}

func init() { proto.RegisterFile("proto/k8s.proto", fileDescriptor_15b74d577d0586ba) }

var fileDescriptor_15b74d577d0586ba = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe5, 0xa6, 0x4d, 0xc9, 0x1b, 0x28, 0xb2, 0xaa, 0xca, 0xaa, 0x10, 0xaa, 0x32, 0x55,
	0x0c, 0x0d, 0x82, 0x25, 0x62, 0x60, 0xeb, 0x14, 0xb1, 0x84, 0x5e, 0xe0, 0xd1, 0x5a, 0x51, 0x94,
	0xd4, 0x8e, 0x6c, 0x17, 0xc1, 0xca, 0x15, 0xb8, 0x04, 0x27, 0xe0, 0x22, 0x5c, 0x81, 0x83, 0x20,
	0xbf, 0x18, 0xe8, 0xcc, 0xf6, 0xff, 0xcf, 0xbf, 0x7f, 0x7d, 0xf6, 0x83, 0x49, 0x67, 0xb4, 0xd3,
	0x59, 0x93, 0xdb, 0x15, 0x29, 0x1e, 0x35, 0xb9, 0x9d, 0x9f, 0x57, 0x5a, 0x57, 0xad, 0xcc, 0xb0,
	0xab, 0x33, 0x54, 0x4a, 0x3b, 0x74, 0xb5, 0x56, 0x21, 0x92, 0x8e, 0x61, 0xb4, 0x7e, 0x92, 0xca,
	0xa5, 0xef, 0x0c, 0x86, 0x1b, 0xb4, 0x0d, 0xe7, 0x30, 0xdc, 0xbc, 0x74, 0x52, 0xb0, 0x05, 0x5b,
	0x26, 0x25, 0x69, 0x3e, 0x83, 0xd8, 0xa1, 0x6d, 0xea, 0x9d, 0x18, 0x2c, 0xd8, 0x32, 0x2a, 0x83,
	0xe3, 0x73, 0x38, 0xf1, 0x8a, 0xf2, 0x11, 0xe5, 0x7f, 0xbd, 0xef, 0x51, 0xb8, 0x97, 0x62, 0xd8,
	0xf7, 0x78, 0xcd, 0xa7, 0x30, 0xaa, 0xf7, 0x58, 0x49, 0x31, 0xa2, 0x61, 0x6f, 0xb8, 0x80, 0xb1,
	0x91, 0x5d, 0x5b, 0x6f, 0x51, 0xc4, 0x54, 0xff, 0x63, 0x7d, 0x5e, 0x3e, 0x3b, 0x83, 0x62, 0xdc,
	0xe7, 0xc9, 0xa4, 0x1f, 0x0c, 0xa0, 0xc8, 0x1f, 0xee, 0xa5, 0xb5, 0xfe, 0xfa, 0x3f, 0x80, 0x09,
	0xec, 0x08, 0x98, 0xe0, 0x66, 0x10, 0x5b, 0x87, 0xee, 0x60, 0x03, 0x72, 0x70, 0xfc, 0x02, 0x60,
	0x87, 0x0e, 0xb7, 0x52, 0x39, 0x69, 0x02, 0xf9, 0xd1, 0x84, 0x9f, 0x41, 0x74, 0x30, 0x2d, 0xa1,
	0x27, 0xa5, 0x97, 0xbe, 0xc9, 0xc8, 0x4e, 0x1b, 0x17, 0xb8, 0x83, 0xbb, 0x5e, 0x43, 0x54, 0xe4,
	0x96, 0xdf, 0x41, 0x5c, 0xe4, 0xf4, 0xd7, 0x93, 0x95, 0x5f, 0xd6, 0xdf, 0x5b, 0xe6, 0x09, 0x0d,
	0xfc, 0x59, 0x3a, 0x7d, 0xfd, 0xfc, 0x7a, 0x1b, 0x9c, 0xa6, 0x89, 0x5f, 0x68, 0xe6, 0x29, 0x6f,
	0xd9, 0xe5, 0x92, 0x5d, 0xb1, 0xc7, 0x98, 0x56, 0x77, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x1f,
	0xd8, 0x2a, 0x3e, 0xf0, 0x01, 0x00, 0x00,
}