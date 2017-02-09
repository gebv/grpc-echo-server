// Code generated by protoc-gen-go.
// source: src/apps/echo/grpc.proto
// DO NOT EDIT!

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	src/apps/echo/grpc.proto

It has these top-level messages:
	EchoDTO
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

type EchoDTO struct {
	Str string `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
	Raw []byte `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (m *EchoDTO) Reset()                    { *m = EchoDTO{} }
func (m *EchoDTO) String() string            { return proto.CompactTextString(m) }
func (*EchoDTO) ProtoMessage()               {}
func (*EchoDTO) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoDTO) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *EchoDTO) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func init() {
	proto.RegisterType((*EchoDTO)(nil), "echo.EchoDTO")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Server service

type ServerClient interface {
	Echo(ctx context.Context, in *EchoDTO, opts ...grpc.CallOption) (*EchoDTO, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Echo(ctx context.Context, in *EchoDTO, opts ...grpc.CallOption) (*EchoDTO, error) {
	out := new(EchoDTO)
	err := grpc.Invoke(ctx, "/echo.Server/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerServer interface {
	Echo(context.Context, *EchoDTO) (*EchoDTO, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Server/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Echo(ctx, req.(*EchoDTO))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Server_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/apps/echo/grpc.proto",
}

func init() { proto.RegisterFile("src/apps/echo/grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2e, 0x4a, 0xd6,
	0x4f, 0x2c, 0x28, 0x28, 0xd6, 0x4f, 0x4d, 0xce, 0xc8, 0xd7, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x09, 0x28, 0xe9, 0x72, 0xb1, 0xbb, 0x26, 0x67, 0xe4,
	0xbb, 0x84, 0xf8, 0x0b, 0x09, 0x70, 0x31, 0x17, 0x97, 0x14, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0x98, 0x20, 0x91, 0xa2, 0xc4, 0x72, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x10,
	0xd3, 0xc8, 0x80, 0x8b, 0x2d, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x48, 0x48, 0x8d, 0x8b, 0x05, 0xa4,
	0x51, 0x88, 0x57, 0x0f, 0x64, 0x8e, 0x1e, 0xd4, 0x10, 0x29, 0x54, 0xae, 0x12, 0x43, 0x12, 0x1b,
	0xd8, 0x36, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x4b, 0xef, 0x9e, 0x89, 0x00, 0x00,
	0x00,
}
