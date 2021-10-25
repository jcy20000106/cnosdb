// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/data.proto

package internal

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

type WriteShardRequest struct {
	ShardID              *uint64  `protobuf:"varint,1,req,name=ShardID" json:"ShardID,omitempty"`
	Points               [][]byte `protobuf:"bytes,2,rep,name=Points" json:"Points,omitempty"`
	Database             *string  `protobuf:"bytes,3,opt,name=Database" json:"Database,omitempty"`
	TimeToLive           *string  `protobuf:"bytes,4,opt,name=TimeToLive" json:"TimeToLive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteShardRequest) Reset()         { *m = WriteShardRequest{} }
func (m *WriteShardRequest) String() string { return proto.CompactTextString(m) }
func (*WriteShardRequest) ProtoMessage()    {}
func (*WriteShardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7438786364df21e1, []int{0}
}
func (m *WriteShardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteShardRequest.Unmarshal(m, b)
}
func (m *WriteShardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteShardRequest.Marshal(b, m, deterministic)
}
func (m *WriteShardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteShardRequest.Merge(m, src)
}
func (m *WriteShardRequest) XXX_Size() int {
	return xxx_messageInfo_WriteShardRequest.Size(m)
}
func (m *WriteShardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteShardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteShardRequest proto.InternalMessageInfo

func (m *WriteShardRequest) GetShardID() uint64 {
	if m != nil && m.ShardID != nil {
		return *m.ShardID
	}
	return 0
}

func (m *WriteShardRequest) GetPoints() [][]byte {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *WriteShardRequest) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *WriteShardRequest) GetTimeToLive() string {
	if m != nil && m.TimeToLive != nil {
		return *m.TimeToLive
	}
	return ""
}

type WriteShardResponse struct {
	Code                 *int32   `protobuf:"varint,1,req,name=Code" json:"Code,omitempty"`
	Message              *string  `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteShardResponse) Reset()         { *m = WriteShardResponse{} }
func (m *WriteShardResponse) String() string { return proto.CompactTextString(m) }
func (*WriteShardResponse) ProtoMessage()    {}
func (*WriteShardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7438786364df21e1, []int{1}
}
func (m *WriteShardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteShardResponse.Unmarshal(m, b)
}
func (m *WriteShardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteShardResponse.Marshal(b, m, deterministic)
}
func (m *WriteShardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteShardResponse.Merge(m, src)
}
func (m *WriteShardResponse) XXX_Size() int {
	return xxx_messageInfo_WriteShardResponse.Size(m)
}
func (m *WriteShardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteShardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteShardResponse proto.InternalMessageInfo

func (m *WriteShardResponse) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *WriteShardResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ExecuteStatementRequest struct {
	Statement            *string  `protobuf:"bytes,1,req,name=Statement" json:"Statement,omitempty"`
	Database             *string  `protobuf:"bytes,2,req,name=Database" json:"Database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteStatementRequest) Reset()         { *m = ExecuteStatementRequest{} }
func (m *ExecuteStatementRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteStatementRequest) ProtoMessage()    {}
func (*ExecuteStatementRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7438786364df21e1, []int{2}
}
func (m *ExecuteStatementRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStatementRequest.Unmarshal(m, b)
}
func (m *ExecuteStatementRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStatementRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteStatementRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStatementRequest.Merge(m, src)
}
func (m *ExecuteStatementRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteStatementRequest.Size(m)
}
func (m *ExecuteStatementRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStatementRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStatementRequest proto.InternalMessageInfo

func (m *ExecuteStatementRequest) GetStatement() string {
	if m != nil && m.Statement != nil {
		return *m.Statement
	}
	return ""
}

func (m *ExecuteStatementRequest) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

type ExecuteStatementResponse struct {
	Code                 *int32   `protobuf:"varint,1,req,name=Code" json:"Code,omitempty"`
	Message              *string  `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteStatementResponse) Reset()         { *m = ExecuteStatementResponse{} }
func (m *ExecuteStatementResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteStatementResponse) ProtoMessage()    {}
func (*ExecuteStatementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7438786364df21e1, []int{3}
}
func (m *ExecuteStatementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStatementResponse.Unmarshal(m, b)
}
func (m *ExecuteStatementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStatementResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteStatementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStatementResponse.Merge(m, src)
}
func (m *ExecuteStatementResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteStatementResponse.Size(m)
}
func (m *ExecuteStatementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStatementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStatementResponse proto.InternalMessageInfo

func (m *ExecuteStatementResponse) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ExecuteStatementResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type CreateIteratorRequest struct {
	ShardIDs             []uint64 `protobuf:"varint,1,rep,name=ShardIDs" json:"ShardIDs,omitempty"`
	Opt                  []byte   `protobuf:"bytes,2,req,name=Opt" json:"Opt,omitempty"`
	Database             []byte   `protobuf:"bytes,3,req,name=Database" json:"Database,omitempty"`
	TimeToLive           []byte   `protobuf:"bytes,4,req,name=TimeToLive" json:"TimeToLive,omitempty"`
	MetricName           []byte   `protobuf:"bytes,5,req,name=MetricName" json:"MetricName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateIteratorRequest) Reset()         { *m = CreateIteratorRequest{} }
func (m *CreateIteratorRequest) String() string { return proto.CompactTextString(m) }
func (*CreateIteratorRequest) ProtoMessage()    {}
func (*CreateIteratorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7438786364df21e1, []int{4}
}
func (m *CreateIteratorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateIteratorRequest.Unmarshal(m, b)
}
func (m *CreateIteratorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateIteratorRequest.Marshal(b, m, deterministic)
}
func (m *CreateIteratorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIteratorRequest.Merge(m, src)
}
func (m *CreateIteratorRequest) XXX_Size() int {
	return xxx_messageInfo_CreateIteratorRequest.Size(m)
}
func (m *CreateIteratorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIteratorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIteratorRequest proto.InternalMessageInfo

func (m *CreateIteratorRequest) GetShardIDs() []uint64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

func (m *CreateIteratorRequest) GetOpt() []byte {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *CreateIteratorRequest) GetDatabase() []byte {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *CreateIteratorRequest) GetTimeToLive() []byte {
	if m != nil {
		return m.TimeToLive
	}
	return nil
}

func (m *CreateIteratorRequest) GetMetricName() []byte {
	if m != nil {
		return m.MetricName
	}
	return nil
}

type CreateIteratorResponse struct {
	Err                  *string  `protobuf:"bytes,1,opt,name=Err" json:"Err,omitempty"`
	DataType             *int32   `protobuf:"varint,2,opt,name=DataType" json:"DataType,omitempty"`
	SeriesN              *int32   `protobuf:"varint,3,opt,name=SeriesN" json:"SeriesN,omitempty"`
	PointN               *int32   `protobuf:"varint,4,opt,name=PointN" json:"PointN,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateIteratorResponse) Reset()         { *m = CreateIteratorResponse{} }
func (m *CreateIteratorResponse) String() string { return proto.CompactTextString(m) }
func (*CreateIteratorResponse) ProtoMessage()    {}
func (*CreateIteratorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7438786364df21e1, []int{5}
}
func (m *CreateIteratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateIteratorResponse.Unmarshal(m, b)
}
func (m *CreateIteratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateIteratorResponse.Marshal(b, m, deterministic)
}
func (m *CreateIteratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIteratorResponse.Merge(m, src)
}
func (m *CreateIteratorResponse) XXX_Size() int {
	return xxx_messageInfo_CreateIteratorResponse.Size(m)
}
func (m *CreateIteratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIteratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIteratorResponse proto.InternalMessageInfo

func (m *CreateIteratorResponse) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

func (m *CreateIteratorResponse) GetDataType() int32 {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return 0
}

func (m *CreateIteratorResponse) GetSeriesN() int32 {
	if m != nil && m.SeriesN != nil {
		return *m.SeriesN
	}
	return 0
}

func (m *CreateIteratorResponse) GetPointN() int32 {
	if m != nil && m.PointN != nil {
		return *m.PointN
	}
	return 0
}

type FieldDimensionsRequest struct {
	ShardIDs             []uint64 `protobuf:"varint,1,rep,name=ShardIDs" json:"ShardIDs,omitempty"`
	Metric               []byte   `protobuf:"bytes,2,req,name=Metric" json:"Metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldDimensionsRequest) Reset()         { *m = FieldDimensionsRequest{} }
func (m *FieldDimensionsRequest) String() string { return proto.CompactTextString(m) }
func (*FieldDimensionsRequest) ProtoMessage()    {}
func (*FieldDimensionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7438786364df21e1, []int{6}
}
func (m *FieldDimensionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldDimensionsRequest.Unmarshal(m, b)
}
func (m *FieldDimensionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldDimensionsRequest.Marshal(b, m, deterministic)
}
func (m *FieldDimensionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldDimensionsRequest.Merge(m, src)
}
func (m *FieldDimensionsRequest) XXX_Size() int {
	return xxx_messageInfo_FieldDimensionsRequest.Size(m)
}
func (m *FieldDimensionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldDimensionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FieldDimensionsRequest proto.InternalMessageInfo

func (m *FieldDimensionsRequest) GetShardIDs() []uint64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

func (m *FieldDimensionsRequest) GetMetric() []byte {
	if m != nil {
		return m.Metric
	}
	return nil
}

type FieldDimensionsResponse struct {
	Fields               []byte   `protobuf:"bytes,1,req,name=Fields" json:"Fields,omitempty"`
	Dimensions           []string `protobuf:"bytes,2,rep,name=Dimensions" json:"Dimensions,omitempty"`
	Err                  *string  `protobuf:"bytes,3,opt,name=Err" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldDimensionsResponse) Reset()         { *m = FieldDimensionsResponse{} }
func (m *FieldDimensionsResponse) String() string { return proto.CompactTextString(m) }
func (*FieldDimensionsResponse) ProtoMessage()    {}
func (*FieldDimensionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7438786364df21e1, []int{7}
}
func (m *FieldDimensionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldDimensionsResponse.Unmarshal(m, b)
}
func (m *FieldDimensionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldDimensionsResponse.Marshal(b, m, deterministic)
}
func (m *FieldDimensionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldDimensionsResponse.Merge(m, src)
}
func (m *FieldDimensionsResponse) XXX_Size() int {
	return xxx_messageInfo_FieldDimensionsResponse.Size(m)
}
func (m *FieldDimensionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldDimensionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FieldDimensionsResponse proto.InternalMessageInfo

func (m *FieldDimensionsResponse) GetFields() []byte {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *FieldDimensionsResponse) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *FieldDimensionsResponse) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*WriteShardRequest)(nil), "internal.WriteShardRequest")
	proto.RegisterType((*WriteShardResponse)(nil), "internal.WriteShardResponse")
	proto.RegisterType((*ExecuteStatementRequest)(nil), "internal.ExecuteStatementRequest")
	proto.RegisterType((*ExecuteStatementResponse)(nil), "internal.ExecuteStatementResponse")
	proto.RegisterType((*CreateIteratorRequest)(nil), "internal.CreateIteratorRequest")
	proto.RegisterType((*CreateIteratorResponse)(nil), "internal.CreateIteratorResponse")
	proto.RegisterType((*FieldDimensionsRequest)(nil), "internal.FieldDimensionsRequest")
	proto.RegisterType((*FieldDimensionsResponse)(nil), "internal.FieldDimensionsResponse")
}

func init() { proto.RegisterFile("internal/data.proto", fileDescriptor_7438786364df21e1) }

var fileDescriptor_7438786364df21e1 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0xc9, 0xbf, 0xda, 0x5c, 0xfa, 0xb0, 0x8e, 0x98, 0x1d, 0x16, 0x91, 0x90, 0xa7, 0x3c,
	0xe9, 0x77, 0x70, 0xbb, 0xe2, 0xc2, 0x6e, 0x94, 0x69, 0xc1, 0xe7, 0xb1, 0xb9, 0xe8, 0xc0, 0x26,
	0x13, 0x67, 0xee, 0x4a, 0x7d, 0xf4, 0xa3, 0xf8, 0x4d, 0x25, 0xd3, 0x99, 0x34, 0xb6, 0x08, 0xb2,
	0x6f, 0xf3, 0x3b, 0x37, 0x5c, 0xce, 0x3d, 0x39, 0xf0, 0x42, 0xf5, 0x84, 0xa6, 0x97, 0x0f, 0x6f,
	0x5b, 0x49, 0xf2, 0xcd, 0x60, 0x34, 0x69, 0xb6, 0x0c, 0x62, 0xf5, 0x2b, 0x82, 0xe7, 0x9f, 0x8d,
	0x22, 0xdc, 0x7c, 0x93, 0xa6, 0x15, 0xf8, 0xfd, 0x11, 0x2d, 0x31, 0x0e, 0xcf, 0x1c, 0xdf, 0xae,
	0x79, 0x54, 0xc6, 0x75, 0x2a, 0x02, 0xb2, 0x02, 0x16, 0x9f, 0xb4, 0xea, 0xc9, 0xf2, 0xb8, 0x4c,
	0xea, 0x95, 0xf0, 0xc4, 0xae, 0x60, 0xb9, 0x96, 0x24, 0xbf, 0x48, 0x8b, 0x3c, 0x29, 0xa3, 0x3a,
	0x17, 0x13, 0xb3, 0xd7, 0x00, 0x5b, 0xd5, 0xe1, 0x56, 0xdf, 0xa9, 0x1f, 0xc8, 0x53, 0x37, 0x9d,
	0x29, 0xd5, 0x3b, 0x60, 0x73, 0x0b, 0x76, 0xd0, 0xbd, 0x45, 0xc6, 0x20, 0xbd, 0xd6, 0x2d, 0x3a,
	0x03, 0x99, 0x70, 0xef, 0xd1, 0xd7, 0x3d, 0x5a, 0x2b, 0xbf, 0x22, 0x8f, 0xdd, 0x9a, 0x80, 0xd5,
	0x06, 0x2e, 0x6f, 0xf6, 0xb8, 0x7b, 0x24, 0xdc, 0x90, 0x24, 0xec, 0xb0, 0xa7, 0x70, 0xcc, 0x2b,
	0xc8, 0x27, 0xcd, 0x6d, 0xcb, 0xc5, 0x51, 0xf8, 0xcb, 0x78, 0xec, 0x86, 0x13, 0x57, 0x1f, 0x80,
	0x9f, 0x2f, 0x7d, 0x92, 0xbd, 0xdf, 0x11, 0xbc, 0xbc, 0x36, 0x28, 0x09, 0x6f, 0x09, 0x8d, 0x24,
	0x6d, 0x82, 0xbb, 0x2b, 0x58, 0xfa, 0x6c, 0x2d, 0x8f, 0xca, 0xa4, 0x4e, 0xc5, 0xc4, 0xec, 0x02,
	0x92, 0x8f, 0x03, 0x39, 0x5b, 0x2b, 0x31, 0x3e, 0x4f, 0x62, 0x1e, 0xe5, 0x7f, 0xc7, 0x3c, 0x4e,
	0x67, 0xca, 0x38, 0xbf, 0x47, 0x32, 0x6a, 0xd7, 0xc8, 0x0e, 0x79, 0x76, 0x98, 0x1f, 0x95, 0x6a,
	0x0f, 0xc5, 0xa9, 0x45, 0x7f, 0xeb, 0x05, 0x24, 0x37, 0xc6, 0xf0, 0xc8, 0xdd, 0x34, 0x3e, 0x83,
	0x8f, 0xed, 0xcf, 0xe1, 0x70, 0x6a, 0x26, 0x26, 0x76, 0xe5, 0x41, 0xa3, 0xd0, 0x36, 0xae, 0x09,
	0x99, 0x08, 0x38, 0x95, 0xa7, 0x71, 0x25, 0xc8, 0x7c, 0x79, 0x9a, 0xea, 0x0e, 0x8a, 0xf7, 0x0a,
	0x1f, 0xda, 0xb5, 0xea, 0xb0, 0xb7, 0x4a, 0xf7, 0xf6, 0x7f, 0xd2, 0x29, 0x60, 0x71, 0x70, 0xef,
	0x03, 0xf2, 0x54, 0xed, 0xe0, 0xf2, 0x6c, 0x9b, 0x3f, 0xa4, 0x80, 0x85, 0x1b, 0x59, 0xf7, 0xdb,
	0x56, 0xc2, 0xd3, 0x18, 0xcd, 0xf1, 0x6b, 0xd7, 0xec, 0x5c, 0xcc, 0x94, 0x10, 0x40, 0x32, 0x05,
	0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x72, 0x8e, 0x71, 0x57, 0x03, 0x00, 0x00,
}