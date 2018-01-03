// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	EchoRequest
	EchoReply
*/
package echo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// The request message containing the user's name.
type EchoRequest struct {
	Num int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

// The response message containing the greetings
type EchoReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *EchoReply) Reset()                    { *m = EchoReply{} }
func (m *EchoReply) String() string            { return proto.CompactTextString(m) }
func (*EchoReply) ProtoMessage()               {}
func (*EchoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "echo.EchoRequest")
	proto.RegisterType((*EchoReply)(nil), "echo.EchoReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	// Sends a greeting
	SayEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
	SayEchoS(ctx context.Context, opts ...grpc.CallOption) (Echo_SayEchoSClient, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) SayEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := grpc.Invoke(ctx, "/echo.Echo/SayEcho", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) SayEchoS(ctx context.Context, opts ...grpc.CallOption) (Echo_SayEchoSClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Echo_serviceDesc.Streams[0], c.cc, "/echo.Echo/SayEchoS", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoSayEchoSClient{stream}
	return x, nil
}

type Echo_SayEchoSClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoReply, error)
	grpc.ClientStream
}

type echoSayEchoSClient struct {
	grpc.ClientStream
}

func (x *echoSayEchoSClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoSayEchoSClient) Recv() (*EchoReply, error) {
	m := new(EchoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Echo service

type EchoServer interface {
	// Sends a greeting
	SayEcho(context.Context, *EchoRequest) (*EchoReply, error)
	SayEchoS(Echo_SayEchoSServer) error
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_SayEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).SayEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/SayEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).SayEcho(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_SayEchoS_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).SayEchoS(&echoSayEchoSServer{stream})
}

type Echo_SayEchoSServer interface {
	Send(*EchoReply) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoSayEchoSServer struct {
	grpc.ServerStream
}

func (x *echoSayEchoSServer) Send(m *EchoReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoSayEchoSServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayEcho",
			Handler:    _Echo_SayEcho_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayEchoS",
			Handler:       _Echo_SayEchoS_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xe4, 0xb9, 0xb8, 0x5d, 0x93,
	0x33, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xf3, 0x4a, 0x73,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x25, 0x55, 0x2e, 0x4e, 0x88, 0x82, 0x82,
	0x9c, 0x4a, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0xb0, 0x12, 0xce,
	0x20, 0x18, 0xd7, 0x28, 0x97, 0x8b, 0x05, 0xa4, 0x4c, 0x48, 0x9f, 0x8b, 0x3d, 0x38, 0xb1, 0x12,
	0xcc, 0x14, 0xd4, 0x03, 0xdb, 0x86, 0x64, 0xbc, 0x14, 0x3f, 0xb2, 0x50, 0x41, 0x4e, 0xa5, 0x12,
	0x83, 0x90, 0x09, 0x17, 0x07, 0x54, 0x43, 0x30, 0x71, 0x3a, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8,
	0xc0, 0x7e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x19, 0x56, 0x78, 0x38, 0xd1, 0x00, 0x00,
	0x00,
}
