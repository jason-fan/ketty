// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exthttp.proto

/*
Package ext_pb is a generated protocol buffer package.

It is generated from these files:
	exthttp.proto

It has these top-level messages:
	QueryReq
	JsonReq
	Req
	Rsp
*/
package ext_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import (
	ketty "github.com/yyzybb537/ketty"
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

type QueryReq struct {
	QVal int64 `protobuf:"varint,1,opt,name=qVal" json:"qVal,omitempty"`
}

func (m *QueryReq) Reset()                    { *m = QueryReq{} }
func (m *QueryReq) String() string            { return proto.CompactTextString(m) }
func (*QueryReq) ProtoMessage()               {}
func (*QueryReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryReq) GetQVal() int64 {
	if m != nil {
		return m.QVal
	}
	return 0
}

type JsonReq struct {
	JVal int64 `protobuf:"varint,1,opt,name=JVal" json:"JVal,omitempty"`
}

func (m *JsonReq) Reset()                    { *m = JsonReq{} }
func (m *JsonReq) String() string            { return proto.CompactTextString(m) }
func (*JsonReq) ProtoMessage()               {}
func (*JsonReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JsonReq) GetJVal() int64 {
	if m != nil {
		return m.JVal
	}
	return 0
}

type Req struct {
	Qr *QueryReq `protobuf:"bytes,1,opt,name=qr" json:"qr,omitempty"`
	Jr *JsonReq  `protobuf:"bytes,2,opt,name=jr" json:"jr,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Req) GetQr() *QueryReq {
	if m != nil {
		return m.Qr
	}
	return nil
}

func (m *Req) GetJr() *JsonReq {
	if m != nil {
		return m.Jr
	}
	return nil
}

type Rsp struct {
	Val int64 `protobuf:"varint,1,opt,name=val" json:"val,omitempty"`
}

func (m *Rsp) Reset()                    { *m = Rsp{} }
func (m *Rsp) String() string            { return proto.CompactTextString(m) }
func (*Rsp) ProtoMessage()               {}
func (*Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Rsp) GetVal() int64 {
	if m != nil {
		return m.Val
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryReq)(nil), "ext_pb.QueryReq")
	proto.RegisterType((*JsonReq)(nil), "ext_pb.JsonReq")
	proto.RegisterType((*Req)(nil), "ext_pb.Req")
	proto.RegisterType((*Rsp)(nil), "ext_pb.Rsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	Echo(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := grpc.Invoke(ctx, "/ext_pb.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Echo(context.Context, *Req) (*Rsp, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ext_pb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ext_pb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exthttp.proto",
}

// Reference imports to suppress errors if they are not otherwise used.
var _ ketty.Dummy

// This is a compile-time assertion to ensure that this generated file
// is compatible with the ketty package it is being compiled against.

type EchoServiceHandleT struct {
	desc *grpc.ServiceDesc
}

func (h *EchoServiceHandleT) Implement() interface{} {
	return h.desc
}

func (h *EchoServiceHandleT) ServiceName() string {
	return h.desc.ServiceName
}

var EchoServiceHandle = &EchoServiceHandleT{desc: &_EchoService_serviceDesc}

type KettyEchoServiceClient struct {
	client ketty.Client
}

func NewKettyEchoServiceClient(client ketty.Client) *KettyEchoServiceClient {
	return &KettyEchoServiceClient{client}
}

func (this *KettyEchoServiceClient) Echo(ctx context.Context, in *Req) (*Rsp, error) {
	out := new(Rsp)
	err := this.client.Invoke(ctx, EchoServiceHandle, "Echo", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (*QueryReq) KettyTransport() string {
	return "query"
}

func (*JsonReq) KettyMarshal() string {
	return "json"
}

func (*JsonReq) KettyTransport() string {
	return "body"
}

func (*Req) KettyHttpExtendMessage() {}

func init() { proto.RegisterFile("exthttp.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xcf, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x05, 0x60, 0xe2, 0x5c, 0x0a, 0x5c, 0x0b, 0x11, 0x79, 0xa1, 0xea, 0x42, 0xe5, 0x09, 0x96,
	0x20, 0xca, 0x96, 0x9d, 0xa5, 0x03, 0x12, 0x46, 0x62, 0x45, 0xa4, 0x5c, 0xa9, 0x04, 0x54, 0xff,
	0x52, 0x35, 0x6f, 0xc0, 0xcc, 0xd8, 0xb7, 0xf4, 0x1b, 0x20, 0xbb, 0x0d, 0xb0, 0x9d, 0x63, 0x7d,
	0xf2, 0xb1, 0xf1, 0x94, 0x36, 0x61, 0x19, 0x82, 0xa9, 0x8d, 0xd3, 0x41, 0x8b, 0x11, 0x6d, 0xc2,
	0xb3, 0x69, 0x27, 0xfc, 0x9d, 0x42, 0xe8, 0x77, 0x87, 0xf2, 0x0a, 0x8f, 0x1f, 0x3e, 0xc9, 0xf5,
	0x8a, 0xac, 0x10, 0x08, 0xf6, 0xe9, 0xe5, 0x63, 0x5c, 0x4c, 0x8b, 0xcb, 0x52, 0xe5, 0xdc, 0x9c,
	0x6c, 0x23, 0x1c, 0xda, 0x24, 0xe4, 0x35, 0x1e, 0xcd, 0xbd, 0x5e, 0xed, 0xe5, 0xfc, 0x9f, 0x4c,
	0xb9, 0xa9, 0xbe, 0x23, 0x40, 0xe7, 0xf5, 0x6a, 0x1b, 0x01, 0x5a, 0xfd, 0xda, 0xcb, 0x7b, 0x2c,
	0x13, 0x9e, 0x22, 0xb3, 0x2e, 0x53, 0x3e, 0xab, 0xea, 0xdd, 0x23, 0xea, 0x61, 0x54, 0x31, 0xeb,
	0xc4, 0x05, 0xb2, 0xce, 0x8d, 0x59, 0x16, 0x67, 0x83, 0xd8, 0x6f, 0x29, 0xd6, 0xb9, 0x06, 0xbe,
	0x22, 0x14, 0xf2, 0x1c, 0x4b, 0xe5, 0x8d, 0xa8, 0xb0, 0x5c, 0xff, 0x6e, 0xa7, 0x38, 0xbb, 0x41,
	0x7e, 0xb7, 0x58, 0xea, 0x47, 0x72, 0xeb, 0xb7, 0x05, 0x09, 0x89, 0x90, 0xaa, 0xe0, 0xc3, 0x55,
	0x8a, 0xec, 0xe4, 0xaf, 0x78, 0x23, 0x0f, 0xda, 0x51, 0xfe, 0xfe, 0xed, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf1, 0xd8, 0xf3, 0xeb, 0x24, 0x01, 0x00, 0x00,
}
