// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/helloworld.proto

package hello_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SubmitMessage struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitMessage) Reset()         { *m = SubmitMessage{} }
func (m *SubmitMessage) String() string { return proto.CompactTextString(m) }
func (*SubmitMessage) ProtoMessage()    {}
func (*SubmitMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d53fe9c48eadaad, []int{0}
}

func (m *SubmitMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitMessage.Unmarshal(m, b)
}
func (m *SubmitMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitMessage.Marshal(b, m, deterministic)
}
func (m *SubmitMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitMessage.Merge(m, src)
}
func (m *SubmitMessage) XXX_Size() int {
	return xxx_messageInfo_SubmitMessage.Size(m)
}
func (m *SubmitMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitMessage proto.InternalMessageInfo

func (m *SubmitMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubmitMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReceiveReply struct {
	Time                 string   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveReply) Reset()         { *m = ReceiveReply{} }
func (m *ReceiveReply) String() string { return proto.CompactTextString(m) }
func (*ReceiveReply) ProtoMessage()    {}
func (*ReceiveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d53fe9c48eadaad, []int{1}
}

func (m *ReceiveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveReply.Unmarshal(m, b)
}
func (m *ReceiveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveReply.Marshal(b, m, deterministic)
}
func (m *ReceiveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveReply.Merge(m, src)
}
func (m *ReceiveReply) XXX_Size() int {
	return xxx_messageInfo_ReceiveReply.Size(m)
}
func (m *ReceiveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveReply proto.InternalMessageInfo

func (m *ReceiveReply) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *ReceiveReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReceiveReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SubmitMessage)(nil), "hello_proto.SubmitMessage")
	proto.RegisterType((*ReceiveReply)(nil), "hello_proto.ReceiveReply")
}

func init() { proto.RegisterFile("proto/helloworld.proto", fileDescriptor_4d53fe9c48eadaad) }

var fileDescriptor_4d53fe9c48eadaad = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x03, 0x0b, 0x08,
	0x71, 0x83, 0x45, 0xe2, 0xc1, 0x1c, 0x25, 0x5b, 0x2e, 0xde, 0xe0, 0xd2, 0xa4, 0xdc, 0xcc, 0x12,
	0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x17, 0x22, 0x2d, 0xc1,
	0x04, 0x16, 0x86, 0x71, 0x95, 0x02, 0xb8, 0x78, 0x82, 0x52, 0x93, 0x53, 0x33, 0xcb, 0x52, 0x83,
	0x52, 0x0b, 0x72, 0x2a, 0x41, 0xba, 0x4b, 0x32, 0x11, 0xba, 0x41, 0x6c, 0xb8, 0x89, 0x4c, 0xd8,
	0x4d, 0x64, 0x46, 0x31, 0xd1, 0x28, 0x80, 0x8b, 0xcb, 0x39, 0x3f, 0x2f, 0x2d, 0xb5, 0xb8, 0x38,
	0x33, 0x3f, 0x4f, 0xc8, 0x89, 0x8b, 0x1d, 0xca, 0x13, 0x92, 0xd2, 0x43, 0x72, 0xb7, 0x1e, 0x8a,
	0xa3, 0xa5, 0x24, 0x51, 0xe4, 0x90, 0x5d, 0xa4, 0xc4, 0x90, 0xc4, 0x06, 0x16, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x9a, 0x74, 0x85, 0x87, 0x10, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfessionClient is the client API for Confession service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfessionClient interface {
	Confess(ctx context.Context, in *SubmitMessage, opts ...grpc.CallOption) (*ReceiveReply, error)
}

type confessionClient struct {
	cc *grpc.ClientConn
}

func NewConfessionClient(cc *grpc.ClientConn) ConfessionClient {
	return &confessionClient{cc}
}

func (c *confessionClient) Confess(ctx context.Context, in *SubmitMessage, opts ...grpc.CallOption) (*ReceiveReply, error) {
	out := new(ReceiveReply)
	err := c.cc.Invoke(ctx, "/hello_proto.Confession/Confess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfessionServer is the server API for Confession service.
type ConfessionServer interface {
	Confess(context.Context, *SubmitMessage) (*ReceiveReply, error)
}

// UnimplementedConfessionServer can be embedded to have forward compatible implementations.
type UnimplementedConfessionServer struct {
}

func (*UnimplementedConfessionServer) Confess(ctx context.Context, req *SubmitMessage) (*ReceiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confess not implemented")
}

func RegisterConfessionServer(s *grpc.Server, srv ConfessionServer) {
	s.RegisterService(&_Confession_serviceDesc, srv)
}

func _Confession_Confess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfessionServer).Confess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello_proto.Confession/Confess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfessionServer).Confess(ctx, req.(*SubmitMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Confession_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello_proto.Confession",
	HandlerType: (*ConfessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Confess",
			Handler:    _Confession_Confess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/helloworld.proto",
}
