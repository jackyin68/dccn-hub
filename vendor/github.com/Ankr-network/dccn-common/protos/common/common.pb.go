// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common_proto

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

type Task struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string     `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string     `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Image                string     `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Replica              int32      `protobuf:"varint,6,opt,name=replica,proto3" json:"replica,omitempty"`
	DataCenter           string     `protobuf:"bytes,7,opt,name=data_center,json=dataCenter,proto3" json:"data_center,omitempty"`
	DataCenterId         int64      `protobuf:"varint,8,opt,name=data_center_id,json=dataCenterId,proto3" json:"data_center_id,omitempty"`
	Status               TaskStatus `protobuf:"varint,9,opt,name=status,proto3,enum=common.proto.TaskStatus" json:"status,omitempty"`
	UniqueName           string     `protobuf:"bytes,10,opt,name=unique_name,json=uniqueName,proto3" json:"unique_name,omitempty"`
	Url                  string     `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
	Hidden               bool       `protobuf:"varint,12,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Uptime               uint32     `protobuf:"varint,13,opt,name=uptime,proto3" json:"uptime,omitempty"`
	CreationDate         uint64     `protobuf:"varint,14,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	Extra                []byte     `protobuf:"bytes,15,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_0e3cd4e7d860622a, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Task) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Task) GetReplica() int32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *Task) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Task) GetDataCenterId() int64 {
	if m != nil {
		return m.DataCenterId
	}
	return 0
}

func (m *Task) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_START
}

func (m *Task) GetUniqueName() string {
	if m != nil {
		return m.UniqueName
	}
	return ""
}

func (m *Task) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Task) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *Task) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *Task) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *Task) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

type DataCenter struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lat                  string   `protobuf:"bytes,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  string   `protobuf:"bytes,4,opt,name=lng,proto3" json:"lng,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Metrics              string   `protobuf:"bytes,6,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Extra                string   `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`
	Report               string   `protobuf:"bytes,8,opt,name=report,proto3" json:"report,omitempty"`
	LastReportTime       string   `protobuf:"bytes,9,opt,name=last_report_time,json=lastReportTime,proto3" json:"last_report_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenter) Reset()         { *m = DataCenter{} }
func (m *DataCenter) String() string { return proto.CompactTextString(m) }
func (*DataCenter) ProtoMessage()    {}
func (*DataCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_0e3cd4e7d860622a, []int{1}
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

func (m *DataCenter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataCenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenter) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *DataCenter) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *DataCenter) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DataCenter) GetMetrics() string {
	if m != nil {
		return m.Metrics
	}
	return ""
}

func (m *DataCenter) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *DataCenter) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *DataCenter) GetLastReportTime() string {
	if m != nil {
		return m.LastReportTime
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "common.proto.Task")
	proto.RegisterType((*DataCenter)(nil), "common.proto.DataCenter")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_0e3cd4e7d860622a) }

var fileDescriptor_common_0e3cd4e7d860622a = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0x87, 0xe5, 0xa6, 0x4d, 0x37, 0xb3, 0x69, 0xa8, 0x2c, 0x04, 0x23, 0x2e, 0x44, 0x0b, 0x07,
	0x9f, 0x2a, 0x04, 0x8f, 0xc0, 0x5e, 0xf6, 0xc2, 0xc1, 0xec, 0x3d, 0x32, 0xf1, 0x68, 0xb1, 0x68,
	0xfe, 0xe0, 0x38, 0x12, 0xbc, 0x0c, 0x8f, 0xc6, 0xb3, 0x20, 0x8f, 0x13, 0xb5, 0xec, 0x6d, 0xbe,
	0xcf, 0x93, 0xd8, 0x33, 0x3f, 0x28, 0xdb, 0xa1, 0xeb, 0x86, 0xfe, 0x34, 0xfa, 0x21, 0x0c, 0xf2,
	0x3f, 0x7a, 0x03, 0xed, 0x60, 0x29, 0xd5, 0x77, 0x7f, 0x32, 0xd8, 0x3e, 0x9a, 0xe9, 0x87, 0xac,
	0x60, 0xe3, 0x2c, 0x8a, 0x5a, 0xa8, 0x42, 0x6f, 0x9c, 0x95, 0xaf, 0x61, 0x3f, 0x4f, 0xe4, 0x1b,
	0x67, 0x71, 0xc3, 0x32, 0x8f, 0xf8, 0x60, 0xa5, 0x84, 0x6d, 0x6f, 0x3a, 0xc2, 0x8c, 0x2d, 0xd7,
	0xd1, 0x85, 0xdf, 0x23, 0xe1, 0x36, 0xb9, 0x58, 0xcb, 0x97, 0xb0, 0x73, 0x9d, 0x79, 0x22, 0xdc,
	0xb1, 0x4c, 0x20, 0x11, 0xf6, 0x9e, 0xc6, 0xb3, 0x6b, 0x0d, 0xe6, 0xb5, 0x50, 0x3b, 0xbd, 0xa2,
	0x7c, 0x0b, 0xb7, 0xd6, 0x04, 0xd3, 0xb4, 0xd4, 0x07, 0xf2, 0xb8, 0xe7, 0xaf, 0x20, 0xaa, 0xcf,
	0x6c, 0xe4, 0x7b, 0xa8, 0xae, 0x1a, 0xe2, 0xc3, 0x6e, 0x6a, 0xa1, 0x32, 0x5d, 0x5e, 0x7a, 0x1e,
	0xac, 0xfc, 0x00, 0xf9, 0x14, 0x4c, 0x98, 0x27, 0x2c, 0x6a, 0xa1, 0xaa, 0x8f, 0x78, 0xba, 0x9e,
	0xfd, 0x14, 0x67, 0xfd, 0xca, 0xe7, 0x7a, 0xe9, 0x8b, 0x17, 0xcf, 0xbd, 0xfb, 0x39, 0x53, 0xc3,
	0x73, 0x41, 0xba, 0x38, 0xa9, 0x2f, 0x71, 0xba, 0x23, 0x64, 0xb3, 0x3f, 0xe3, 0x2d, 0x1f, 0xc4,
	0x52, 0xbe, 0x82, 0xfc, 0xbb, 0xb3, 0x96, 0x7a, 0x2c, 0x6b, 0xa1, 0x6e, 0xf4, 0x42, 0xd1, 0xcf,
	0x63, 0x70, 0x1d, 0xe1, 0xa1, 0x16, 0xea, 0xa0, 0x17, 0x92, 0xef, 0xe0, 0xd0, 0x7a, 0x32, 0xc1,
	0x0d, 0x7d, 0x63, 0x4d, 0x20, 0xac, 0x6a, 0xa1, 0xb6, 0xba, 0x5c, 0xe5, 0xbd, 0x09, 0xbc, 0x30,
	0xfa, 0x15, 0xbc, 0xc1, 0x17, 0xb5, 0x50, 0xa5, 0x4e, 0x70, 0xf7, 0x57, 0x00, 0xdc, 0x5f, 0x96,
	0xf0, 0x3c, 0xa6, 0x35, 0x8d, 0xcd, 0x55, 0x1a, 0x47, 0xc8, 0xce, 0x26, 0x2c, 0x01, 0xc5, 0x92,
	0x4d, 0xff, 0xb4, 0xc4, 0x13, 0xcb, 0xf8, 0xd2, 0x65, 0x4d, 0x29, 0x9e, 0x75, 0x19, 0x08, 0xfb,
	0x8e, 0x82, 0x77, 0xed, 0xc4, 0xf9, 0x14, 0x7a, 0xc5, 0xcb, 0xf3, 0x52, 0x32, 0x09, 0xe2, 0x7f,
	0x3c, 0x8d, 0x83, 0x0f, 0x1c, 0x46, 0xa1, 0x17, 0x92, 0x0a, 0x8e, 0x67, 0x33, 0x85, 0x26, 0x61,
	0xc3, 0x3b, 0x29, 0xb8, 0xa3, 0x8a, 0x5e, 0xb3, 0x7e, 0x74, 0x1d, 0x7d, 0xcb, 0x39, 0x98, 0x4f,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x97, 0x2e, 0x66, 0xc6, 0xb2, 0x02, 0x00, 0x00,
}