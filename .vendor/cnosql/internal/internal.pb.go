// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/internal.proto

package cnosql

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Metrics struct {
	Items                []*Metric `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_41ca0a4a9dd77d9e, []int{0}
}
func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (m *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(m, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetItems() []*Metric {
	if m != nil {
		return m.Items
	}
	return nil
}

type Metric struct {
	Database             *string  `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	TimeToLive           *string  `protobuf:"bytes,2,opt,name=TimeToLive" json:"TimeToLive,omitempty"`
	Name                 *string  `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex                *string  `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget             *bool    `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_41ca0a4a9dd77d9e, []int{1}
}
func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *Metric) GetTimeToLive() string {
	if m != nil && m.TimeToLive != nil {
		return *m.TimeToLive
	}
	return ""
}

func (m *Metric) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Metric) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *Metric) GetIsTarget() bool {
	if m != nil && m.IsTarget != nil {
		return *m.IsTarget
	}
	return false
}

func init() {
	proto.RegisterType((*Metrics)(nil), "cnosql.Metrics")
	proto.RegisterType((*Metric)(nil), "cnosql.Metric")
}

func init() { proto.RegisterFile("internal/internal.proto", fileDescriptor_41ca0a4a9dd77d9e) }

var fileDescriptor_41ca0a4a9dd77d9e = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0x92,
	0xf3, 0xf2, 0x8b, 0x0b, 0x73, 0x94, 0xf4, 0xb9, 0xd8, 0x7d, 0x53, 0x4b, 0x8a, 0x32, 0x93, 0x8b,
	0x85, 0x54, 0xb8, 0x58, 0x3d, 0x4b, 0x52, 0x73, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d,
	0xf8, 0xf4, 0x20, 0x4a, 0xf4, 0x20, 0xf2, 0x41, 0x10, 0x49, 0xa5, 0x2e, 0x46, 0x2e, 0x36, 0x88,
	0x88, 0x90, 0x14, 0x17, 0x87, 0x4b, 0x62, 0x49, 0x62, 0x52, 0x62, 0x71, 0xaa, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f, 0x24, 0xc7, 0xc5, 0x15, 0x92, 0x99, 0x9b, 0x1a, 0x92, 0xef,
	0x93, 0x59, 0x96, 0x2a, 0xc1, 0x04, 0x96, 0x45, 0x12, 0x11, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x06, 0xcb, 0x80, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x41, 0xa9, 0xe9, 0xa9, 0x15,
	0x12, 0x2c, 0x60, 0x41, 0x08, 0x07, 0x64, 0x8b, 0x67, 0x71, 0x48, 0x62, 0x51, 0x7a, 0x6a, 0x89,
	0x04, 0xab, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x9c, 0x0f, 0x08, 0x00, 0x00, 0xff, 0xff, 0x52, 0x9a,
	0xda, 0x9c, 0xdf, 0x00, 0x00, 0x00,
}
