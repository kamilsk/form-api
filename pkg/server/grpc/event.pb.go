// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Ignoring public import of Error from common.proto

// Ignoring public import of TimestampRange from common.proto

type LogEntry struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SchemaId             string               `protobuf:"bytes,3,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	InputId              string               `protobuf:"bytes,4,opt,name=input_id,json=inputId,proto3" json:"input_id,omitempty"`
	TemplateId           string               `protobuf:"bytes,5,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Identifier           string               `protobuf:"bytes,6,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Code                 int32                `protobuf:"varint,7,opt,name=code,proto3" json:"code,omitempty"`
	Context              []byte               `protobuf:"bytes,8,opt,name=context,proto3" json:"context,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_eaef3bcf5e6f8e8c, []int{0}
}
func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (dst *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(dst, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LogEntry) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *LogEntry) GetInputId() string {
	if m != nil {
		return m.InputId
	}
	return ""
}

func (m *LogEntry) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *LogEntry) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *LogEntry) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LogEntry) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *LogEntry) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type LogFilter struct {
	SchemaId             string          `protobuf:"bytes,3,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	InputId              string          `protobuf:"bytes,4,opt,name=input_id,json=inputId,proto3" json:"input_id,omitempty"`
	TemplateId           string          `protobuf:"bytes,5,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Identifier           string          `protobuf:"bytes,6,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Code                 int32           `protobuf:"varint,7,opt,name=code,proto3" json:"code,omitempty"`
	CreatedAt            *TimestampRange `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Limit                uint32          `protobuf:"varint,15,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LogFilter) Reset()         { *m = LogFilter{} }
func (m *LogFilter) String() string { return proto.CompactTextString(m) }
func (*LogFilter) ProtoMessage()    {}
func (*LogFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_eaef3bcf5e6f8e8c, []int{1}
}
func (m *LogFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogFilter.Unmarshal(m, b)
}
func (m *LogFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogFilter.Marshal(b, m, deterministic)
}
func (dst *LogFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogFilter.Merge(dst, src)
}
func (m *LogFilter) XXX_Size() int {
	return xxx_messageInfo_LogFilter.Size(m)
}
func (m *LogFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_LogFilter.DiscardUnknown(m)
}

var xxx_messageInfo_LogFilter proto.InternalMessageInfo

func (m *LogFilter) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *LogFilter) GetInputId() string {
	if m != nil {
		return m.InputId
	}
	return ""
}

func (m *LogFilter) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *LogFilter) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *LogFilter) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LogFilter) GetCreatedAt() *TimestampRange {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *LogFilter) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadLogsRequest struct {
	// Types that are valid to be assigned to Filter:
	//	*ReadLogsRequest_Id
	//	*ReadLogsRequest_Condition
	Filter               isReadLogsRequest_Filter `protobuf_oneof:"filter"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ReadLogsRequest) Reset()         { *m = ReadLogsRequest{} }
func (m *ReadLogsRequest) String() string { return proto.CompactTextString(m) }
func (*ReadLogsRequest) ProtoMessage()    {}
func (*ReadLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_eaef3bcf5e6f8e8c, []int{2}
}
func (m *ReadLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadLogsRequest.Unmarshal(m, b)
}
func (m *ReadLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadLogsRequest.Marshal(b, m, deterministic)
}
func (dst *ReadLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadLogsRequest.Merge(dst, src)
}
func (m *ReadLogsRequest) XXX_Size() int {
	return xxx_messageInfo_ReadLogsRequest.Size(m)
}
func (m *ReadLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadLogsRequest proto.InternalMessageInfo

type isReadLogsRequest_Filter interface {
	isReadLogsRequest_Filter()
}

type ReadLogsRequest_Id struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}
type ReadLogsRequest_Condition struct {
	Condition *LogFilter `protobuf:"bytes,2,opt,name=condition,proto3,oneof"`
}

func (*ReadLogsRequest_Id) isReadLogsRequest_Filter()        {}
func (*ReadLogsRequest_Condition) isReadLogsRequest_Filter() {}

func (m *ReadLogsRequest) GetFilter() isReadLogsRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ReadLogsRequest) GetId() uint32 {
	if x, ok := m.GetFilter().(*ReadLogsRequest_Id); ok {
		return x.Id
	}
	return 0
}

func (m *ReadLogsRequest) GetCondition() *LogFilter {
	if x, ok := m.GetFilter().(*ReadLogsRequest_Condition); ok {
		return x.Condition
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadLogsRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReadLogsRequest_OneofMarshaler, _ReadLogsRequest_OneofUnmarshaler, _ReadLogsRequest_OneofSizer, []interface{}{
		(*ReadLogsRequest_Id)(nil),
		(*ReadLogsRequest_Condition)(nil),
	}
}

func _ReadLogsRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadLogsRequest)
	// filter
	switch x := m.Filter.(type) {
	case *ReadLogsRequest_Id:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Id))
	case *ReadLogsRequest_Condition:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Condition); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReadLogsRequest.Filter has unexpected type %T", x)
	}
	return nil
}

func _ReadLogsRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadLogsRequest)
	switch tag {
	case 1: // filter.id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Filter = &ReadLogsRequest_Id{uint32(x)}
		return true, err
	case 2: // filter.condition
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogFilter)
		err := b.DecodeMessage(msg)
		m.Filter = &ReadLogsRequest_Condition{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReadLogsRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReadLogsRequest)
	// filter
	switch x := m.Filter.(type) {
	case *ReadLogsRequest_Id:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Id))
	case *ReadLogsRequest_Condition:
		s := proto.Size(x.Condition)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ReadLogsResponse struct {
	Logs                 []*LogEntry `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	Err                  *Error      `protobuf:"bytes,15,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReadLogsResponse) Reset()         { *m = ReadLogsResponse{} }
func (m *ReadLogsResponse) String() string { return proto.CompactTextString(m) }
func (*ReadLogsResponse) ProtoMessage()    {}
func (*ReadLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_eaef3bcf5e6f8e8c, []int{3}
}
func (m *ReadLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadLogsResponse.Unmarshal(m, b)
}
func (m *ReadLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadLogsResponse.Marshal(b, m, deterministic)
}
func (dst *ReadLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadLogsResponse.Merge(dst, src)
}
func (m *ReadLogsResponse) XXX_Size() int {
	return xxx_messageInfo_ReadLogsResponse.Size(m)
}
func (m *ReadLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadLogsResponse proto.InternalMessageInfo

func (m *ReadLogsResponse) GetLogs() []*LogEntry {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *ReadLogsResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type ListenLogsRequest struct {
	Filter               *LogFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListenLogsRequest) Reset()         { *m = ListenLogsRequest{} }
func (m *ListenLogsRequest) String() string { return proto.CompactTextString(m) }
func (*ListenLogsRequest) ProtoMessage()    {}
func (*ListenLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_eaef3bcf5e6f8e8c, []int{4}
}
func (m *ListenLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenLogsRequest.Unmarshal(m, b)
}
func (m *ListenLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenLogsRequest.Marshal(b, m, deterministic)
}
func (dst *ListenLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenLogsRequest.Merge(dst, src)
}
func (m *ListenLogsRequest) XXX_Size() int {
	return xxx_messageInfo_ListenLogsRequest.Size(m)
}
func (m *ListenLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenLogsRequest proto.InternalMessageInfo

func (m *ListenLogsRequest) GetFilter() *LogFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "grpc.LogEntry")
	proto.RegisterType((*LogFilter)(nil), "grpc.LogFilter")
	proto.RegisterType((*ReadLogsRequest)(nil), "grpc.ReadLogsRequest")
	proto.RegisterType((*ReadLogsResponse)(nil), "grpc.ReadLogsResponse")
	proto.RegisterType((*ListenLogsRequest)(nil), "grpc.ListenLogsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogClient interface {
	Read(ctx context.Context, in *ReadLogsRequest, opts ...grpc.CallOption) (*ReadLogsResponse, error)
	Listen(ctx context.Context, in *ListenLogsRequest, opts ...grpc.CallOption) (Log_ListenClient, error)
}

type logClient struct {
	cc *grpc.ClientConn
}

func NewLogClient(cc *grpc.ClientConn) LogClient {
	return &logClient{cc}
}

func (c *logClient) Read(ctx context.Context, in *ReadLogsRequest, opts ...grpc.CallOption) (*ReadLogsResponse, error) {
	out := new(ReadLogsResponse)
	err := c.cc.Invoke(ctx, "/grpc.Log/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Listen(ctx context.Context, in *ListenLogsRequest, opts ...grpc.CallOption) (Log_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Log_serviceDesc.Streams[0], "/grpc.Log/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &logListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Log_ListenClient interface {
	Recv() (*LogEntry, error)
	grpc.ClientStream
}

type logListenClient struct {
	grpc.ClientStream
}

func (x *logListenClient) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServer is the server API for Log service.
type LogServer interface {
	Read(context.Context, *ReadLogsRequest) (*ReadLogsResponse, error)
	Listen(*ListenLogsRequest, Log_ListenServer) error
}

func RegisterLogServer(s *grpc.Server, srv LogServer) {
	s.RegisterService(&_Log_serviceDesc, srv)
}

func _Log_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Log/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Read(ctx, req.(*ReadLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServer).Listen(m, &logListenServer{stream})
}

type Log_ListenServer interface {
	Send(*LogEntry) error
	grpc.ServerStream
}

type logListenServer struct {
	grpc.ServerStream
}

func (x *logListenServer) Send(m *LogEntry) error {
	return x.ServerStream.SendMsg(m)
}

var _Log_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Log_Read_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _Log_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_event_eaef3bcf5e6f8e8c) }

var fileDescriptor_event_eaef3bcf5e6f8e8c = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x14, 0xac, 0xfb, 0x9d, 0x97, 0x65, 0xbb, 0x58, 0x0b, 0x98, 0x22, 0xd8, 0x28, 0x17, 0x72, 0x6a,
	0x51, 0x57, 0x08, 0x21, 0x71, 0x01, 0x69, 0xd1, 0x56, 0xca, 0x01, 0x59, 0x70, 0x65, 0x95, 0x8d,
	0x5f, 0x83, 0xa5, 0xc4, 0x0e, 0x8e, 0x8b, 0xe0, 0xf7, 0xf2, 0x37, 0x38, 0xa0, 0xd8, 0xcd, 0xb6,
	0x2a, 0xfc, 0x80, 0xbd, 0xf9, 0xcd, 0x8c, 0x3c, 0x33, 0x7e, 0x86, 0x10, 0x7f, 0xa0, 0xb2, 0x8b,
	0xda, 0x68, 0xab, 0xe9, 0xb0, 0x30, 0x75, 0x3e, 0x3f, 0xc9, 0x75, 0x55, 0x69, 0xe5, 0xb1, 0xf9,
	0x45, 0xa1, 0x75, 0x51, 0xe2, 0xd2, 0x4d, 0xb7, 0xdb, 0xcd, 0xd2, 0xca, 0x0a, 0x1b, 0x9b, 0x55,
	0xb5, 0x17, 0xc4, 0x7f, 0x08, 0x4c, 0x53, 0x5d, 0x5c, 0x29, 0x6b, 0x7e, 0xd1, 0x53, 0xe8, 0x4b,
	0xc1, 0x48, 0x44, 0x92, 0x07, 0xbc, 0x2f, 0x05, 0x7d, 0x06, 0x41, 0x93, 0x7f, 0xc3, 0x2a, 0xbb,
	0x91, 0x82, 0x0d, 0x22, 0x92, 0x04, 0x7c, 0xea, 0x81, 0xb5, 0xa0, 0x4f, 0x61, 0x2a, 0x55, 0xbd,
	0xb5, 0x2d, 0x37, 0x74, 0xdc, 0xc4, 0xcd, 0x6b, 0x41, 0x2f, 0x20, 0xb4, 0x58, 0xd5, 0x65, 0x66,
	0xb1, 0x65, 0x47, 0x8e, 0x85, 0x0e, 0x5a, 0x0b, 0xfa, 0x02, 0x40, 0x0a, 0x54, 0x56, 0x6e, 0x24,
	0x1a, 0x36, 0xf6, 0xfc, 0x1e, 0xa1, 0x14, 0x86, 0xb9, 0x16, 0xc8, 0x26, 0x11, 0x49, 0x46, 0xdc,
	0x9d, 0x29, 0x83, 0x49, 0xae, 0x95, 0xc5, 0x9f, 0x96, 0x4d, 0x23, 0x92, 0x9c, 0xf0, 0x6e, 0xa4,
	0x6f, 0x01, 0x72, 0x83, 0x99, 0x45, 0x71, 0x93, 0x59, 0x16, 0x44, 0x24, 0x09, 0x57, 0xf3, 0x85,
	0x6f, 0xbe, 0xe8, 0x9a, 0x2f, 0x3e, 0x77, 0xcd, 0x79, 0xb0, 0x53, 0xbf, 0xb7, 0xf1, 0x6f, 0x02,
	0x41, 0xaa, 0x8b, 0x8f, 0xb2, 0xb4, 0x68, 0xee, 0x57, 0xdf, 0xcb, 0xff, 0xb4, 0x3a, 0x5f, 0xb4,
	0x3b, 0x3e, 0xa8, 0x92, 0xa9, 0x02, 0x0f, 0xfa, 0xd0, 0x73, 0x18, 0x95, 0xb2, 0x92, 0x96, 0xcd,
	0xdc, 0x12, 0xfd, 0x10, 0x7f, 0x85, 0x19, 0xc7, 0x4c, 0xa4, 0xba, 0x68, 0x38, 0x7e, 0xdf, 0x62,
	0x63, 0xe9, 0xd9, 0x7e, 0xd5, 0xd7, 0x3d, 0xb7, 0xec, 0x25, 0x04, 0xb9, 0x56, 0x42, 0x5a, 0xa9,
	0x15, 0xeb, 0x3b, 0xbb, 0x99, 0xb7, 0xbb, 0x7b, 0xa0, 0xeb, 0x1e, 0xdf, 0x6b, 0x3e, 0x4c, 0x61,
	0xbc, 0x71, 0x70, 0xfc, 0x05, 0xce, 0xf6, 0xf7, 0x37, 0xb5, 0x56, 0x0d, 0xd2, 0x18, 0x86, 0xa5,
	0x2e, 0x1a, 0x46, 0xa2, 0x41, 0x12, 0xae, 0x4e, 0xef, 0x6e, 0x72, 0x3f, 0x8d, 0x3b, 0x8e, 0x3e,
	0x87, 0x01, 0x1a, 0xe3, 0xb2, 0x86, 0xab, 0xd0, 0x4b, 0xae, 0x8c, 0xd1, 0x86, 0xb7, 0x78, 0xfc,
	0x0e, 0x1e, 0xa6, 0xb2, 0xb1, 0xa8, 0x0e, 0x83, 0xbf, 0xec, 0x5c, 0x5d, 0xf8, 0x7f, 0x33, 0xf2,
	0x1d, 0xbd, 0xda, 0xc2, 0x20, 0xd5, 0x05, 0x7d, 0x03, 0xc3, 0x36, 0x1b, 0x7d, 0xe4, 0x75, 0x47,
	0xef, 0x30, 0x7f, 0x7c, 0x0c, 0xfb, 0xf8, 0x71, 0x8f, 0xbe, 0x86, 0xb1, 0x77, 0xa7, 0x4f, 0x76,
	0x16, 0xc7, 0x59, 0xe6, 0x47, 0xad, 0xe2, 0xde, 0x2b, 0xf2, 0xa9, 0x77, 0x3b, 0x76, 0x5f, 0xee,
	0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xf6, 0x96, 0x07, 0x9d, 0x03, 0x00, 0x00,
}
