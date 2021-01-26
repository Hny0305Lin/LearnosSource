// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/status.proto

package server_status

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

type GetStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatus) Reset()         { *m = GetStatus{} }
func (m *GetStatus) String() string { return proto.CompactTextString(m) }
func (*GetStatus) ProtoMessage()    {}
func (*GetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_17c96f8603cffea9, []int{0}
}

func (m *GetStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatus.Unmarshal(m, b)
}
func (m *GetStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatus.Marshal(b, m, deterministic)
}
func (m *GetStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatus.Merge(m, src)
}
func (m *GetStatus) XXX_Size() int {
	return xxx_messageInfo_GetStatus.Size(m)
}
func (m *GetStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatus.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatus proto.InternalMessageInfo

type RetStatus struct {
	Msg                  uint64   `protobuf:"varint,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetStatus) Reset()         { *m = RetStatus{} }
func (m *RetStatus) String() string { return proto.CompactTextString(m) }
func (*RetStatus) ProtoMessage()    {}
func (*RetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_17c96f8603cffea9, []int{1}
}

func (m *RetStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetStatus.Unmarshal(m, b)
}
func (m *RetStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetStatus.Marshal(b, m, deterministic)
}
func (m *RetStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetStatus.Merge(m, src)
}
func (m *RetStatus) XXX_Size() int {
	return xxx_messageInfo_RetStatus.Size(m)
}
func (m *RetStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RetStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RetStatus proto.InternalMessageInfo

func (m *RetStatus) GetMsg() uint64 {
	if m != nil {
		return m.Msg
	}
	return 0
}

func init() {
	proto.RegisterType((*GetStatus)(nil), "server.status.get_status")
	proto.RegisterType((*RetStatus)(nil), "server.status.ret_status")
}

func init() { proto.RegisterFile("proto/status.proto", fileDescriptor_17c96f8603cffea9) }

var fileDescriptor_17c96f8603cffea9 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x03, 0x73, 0x84, 0x78, 0x8b, 0x53, 0x8b,
	0xca, 0x52, 0x8b, 0xf4, 0x20, 0x82, 0x4a, 0x3c, 0x5c, 0x5c, 0xe9, 0xa9, 0x25, 0xf1, 0x50, 0x9e,
	0x1c, 0x17, 0x57, 0x11, 0x9c, 0x27, 0x24, 0xc0, 0xc5, 0x9c, 0x5b, 0x9c, 0x2e, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x12, 0x04, 0x62, 0x1a, 0x79, 0x73, 0x71, 0xba, 0xa7, 0x96, 0x04, 0x43, 0xa4, 0xed,
	0xb8, 0xd8, 0xa0, 0x2c, 0x49, 0x3d, 0x14, 0x43, 0xf5, 0x10, 0x26, 0x4a, 0xa1, 0x4b, 0x21, 0x8c,
	0x4f, 0x62, 0x03, 0x3b, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x41, 0x3a, 0x29, 0xdb, 0xa6,
	0x00, 0x00, 0x00,
}
