// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package grpc

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TimestampRange struct {
	Start                *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimestampRange) Reset()         { *m = TimestampRange{} }
func (m *TimestampRange) String() string { return proto.CompactTextString(m) }
func (*TimestampRange) ProtoMessage()    {}
func (*TimestampRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}
func (m *TimestampRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimestampRange.Unmarshal(m, b)
}
func (m *TimestampRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimestampRange.Marshal(b, m, deterministic)
}
func (m *TimestampRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimestampRange.Merge(m, src)
}
func (m *TimestampRange) XXX_Size() int {
	return xxx_messageInfo_TimestampRange.Size(m)
}
func (m *TimestampRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimestampRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimestampRange proto.InternalMessageInfo

func (m *TimestampRange) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TimestampRange) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*TimestampRange)(nil), "grpc.TimestampRange")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2f, 0x2a, 0x48, 0x96, 0x92,
	0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6,
	0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x94, 0x29, 0x15, 0x70, 0xf1, 0x85, 0xc0, 0x84, 0x82,
	0x12, 0xf3, 0xd2, 0x53, 0x85, 0x0c, 0xb8, 0x58, 0x8b, 0x4b, 0x12, 0x8b, 0x4a, 0x24, 0x18, 0x15,
	0x18, 0x35, 0xb8, 0x8d, 0xa4, 0xf4, 0x20, 0x46, 0xe8, 0xc1, 0x8c, 0xd0, 0x43, 0xa8, 0x87, 0x28,
	0x14, 0xd2, 0xe1, 0x62, 0x4e, 0xcd, 0x4b, 0x91, 0x60, 0x22, 0xa8, 0x1e, 0xa4, 0x2c, 0x89, 0x0d,
	0x2c, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x41, 0xeb, 0x49, 0xc8, 0xaf, 0x00, 0x00, 0x00,
}