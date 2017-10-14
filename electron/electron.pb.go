// Code generated by protoc-gen-go. DO NOT EDIT.
// source: electron.proto

/*
Package electron is a generated protocol buffer package.

It is generated from these files:
	electron.proto

It has these top-level messages:
	Request
	Response
*/
package electron

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

type Request struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "electron.Request")
	proto.RegisterType((*Response)(nil), "electron.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Electron service

type ElectronClient interface {
	Read(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type electronClient struct {
	cc *grpc.ClientConn
}

func NewElectronClient(cc *grpc.ClientConn) ElectronClient {
	return &electronClient{cc}
}

func (c *electronClient) Read(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/electron.Electron/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Electron service

type ElectronServer interface {
	Read(context.Context, *Request) (*Response, error)
}

func RegisterElectronServer(s *grpc.Server, srv ElectronServer) {
	s.RegisterService(&_Electron_serviceDesc, srv)
}

func _Electron_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectronServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/electron.Electron/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectronServer).Read(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Electron_serviceDesc = grpc.ServiceDesc{
	ServiceName: "electron.Electron",
	HandlerType: (*ElectronServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Electron_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "electron.proto",
}

func init() { proto.RegisterFile("electron.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x49, 0x4d,
	0x2e, 0x29, 0xca, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x24,
	0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x94, 0xa4, 0xb8, 0x38, 0x82, 0x52, 0x8b,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xd1, 0xe5, 0x8c, 0xac, 0xb9, 0x38, 0x5c, 0xa1, 0x46, 0x08, 0xe9,
	0x73, 0xb1, 0x04, 0xa5, 0x26, 0xa6, 0x08, 0x09, 0xea, 0xc1, 0x6d, 0x81, 0x1a, 0x29, 0x25, 0x84,
	0x2c, 0x04, 0x31, 0x4a, 0x89, 0xc1, 0x49, 0x2f, 0x4a, 0x27, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0xbf, 0xa0, 0x28, 0xb1, 0x24, 0x35, 0x35, 0xbb, 0x20, 0x31, 0x2f, 0x25,
	0xb5, 0xd2, 0xd0, 0x44, 0x1f, 0xa6, 0x03, 0xce, 0x48, 0x62, 0x03, 0x3b, 0xda, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xab, 0x5c, 0xff, 0xf6, 0xc6, 0x00, 0x00, 0x00,
}