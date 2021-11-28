// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package proto

import (
	fmt "fmt"
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

type Query struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Word                 string   `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Query) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type None struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *None) Reset()         { *m = None{} }
func (m *None) String() string { return proto.CompactTextString(m) }
func (*None) ProtoMessage()    {}
func (*None) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *None) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_None.Unmarshal(m, b)
}
func (m *None) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_None.Marshal(b, m, deterministic)
}
func (m *None) XXX_Merge(src proto.Message) {
	xxx_messageInfo_None.Merge(m, src)
}
func (m *None) XXX_Size() int {
	return xxx_messageInfo_None.Size(m)
}
func (m *None) XXX_DiscardUnknown() {
	xxx_messageInfo_None.DiscardUnknown(m)
}

var xxx_messageInfo_None proto.InternalMessageInfo

type Result struct {
	Mean                 string   `protobuf:"bytes,1,opt,name=mean,proto3" json:"mean,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetMean() string {
	if m != nil {
		return m.Mean
	}
	return ""
}

type Metric struct {
	Mem                  int32    `protobuf:"varint,1,opt,name=Mem,json=mem,proto3" json:"Mem,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=Ts,json=ts,proto3" json:"Ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
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

func (m *Metric) GetMem() int32 {
	if m != nil {
		return m.Mem
	}
	return 0
}

func (m *Metric) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type Msg struct {
	From                 int32    `protobuf:"varint,1,opt,name=From,json=from,proto3" json:"From,omitempty"`
	To                   int32    `protobuf:"varint,2,opt,name=To,json=to,proto3" json:"To,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=Content,json=content,proto3" json:"Content,omitempty"`
	Ts                   int64    `protobuf:"varint,4,opt,name=Ts,json=ts,proto3" json:"Ts,omitempty"`
	Onlines              []int32  `protobuf:"varint,5,rep,packed,name=Onlines,json=onlines,proto3" json:"Onlines,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Msg) GetTo() int32 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Msg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Msg) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *Msg) GetOnlines() []int32 {
	if m != nil {
		return m.Onlines
	}
	return nil
}

func init() {
	proto.RegisterType((*Query)(nil), "proto.Query")
	proto.RegisterType((*None)(nil), "proto.None")
	proto.RegisterType((*Result)(nil), "proto.Result")
	proto.RegisterType((*Metric)(nil), "proto.Metric")
	proto.RegisterType((*Msg)(nil), "proto.Msg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xbd, 0x4a, 0x04, 0x31,
	0x14, 0x85, 0x99, 0xfc, 0x0d, 0x73, 0x0b, 0x91, 0x54, 0x29, 0x2c, 0x86, 0x54, 0x83, 0x82, 0x8d,
	0x8f, 0x20, 0xd8, 0x8d, 0x62, 0xf0, 0x05, 0x74, 0xe7, 0xba, 0x04, 0x36, 0xb9, 0x92, 0x64, 0x11,
	0xdf, 0x5e, 0x72, 0x67, 0xb7, 0xca, 0x39, 0xf0, 0xe5, 0x3b, 0x5c, 0x98, 0x52, 0x3d, 0x3e, 0xfe,
	0x14, 0x6a, 0x64, 0x35, 0x3f, 0xfe, 0x01, 0xf4, 0xfb, 0x19, 0xcb, 0x9f, 0xbd, 0x01, 0x11, 0x37,
	0x37, 0xcc, 0xc3, 0xa2, 0x83, 0x88, 0x9b, 0xb5, 0xa0, 0x7e, 0xa9, 0x6c, 0x4e, 0xcc, 0xc3, 0x32,
	0x05, 0xce, 0xde, 0x80, 0x7a, 0xa5, 0x8c, 0xfe, 0x0e, 0x4c, 0xc0, 0x7a, 0x3e, 0xb5, 0x4e, 0x25,
	0xfc, 0xcc, 0xfc, 0x6f, 0x0a, 0x9c, 0xfd, 0x3d, 0x98, 0x15, 0x5b, 0x89, 0x07, 0x7b, 0x0b, 0x72,
	0xc5, 0x74, 0x91, 0xca, 0x84, 0xa9, 0xaf, 0x7c, 0x54, 0x76, 0xca, 0x20, 0x5a, 0xf5, 0x09, 0xe4,
	0x5a, 0x8f, 0x5d, 0xf3, 0x52, 0xe8, 0x4a, 0xaa, 0xef, 0x42, 0x3b, 0x4a, 0x8c, 0xea, 0x20, 0x1a,
	0x59, 0x07, 0xe3, 0x33, 0xe5, 0x86, 0xb9, 0x39, 0xc9, 0x6b, 0xe3, 0x61, 0xaf, 0x17, 0xa9, 0xba,
	0x4a, 0x3b, 0xf9, 0x96, 0x4f, 0x31, 0x63, 0x75, 0x7a, 0x96, 0x8b, 0x0e, 0x23, 0xed, 0xf5, 0xcb,
	0xf0, 0xd1, 0x4f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xcb, 0xb0, 0xc9, 0x08, 0x01, 0x00,
	0x00,
}