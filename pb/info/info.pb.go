// Code generated by protoc-gen-go.
// source: info.proto
// DO NOT EDIT!

/*
Package info is a generated protocol buffer package.

It is generated from these files:
	info.proto

It has these top-level messages:
	GetInfoRequest
	GetInfoResponse
	WebServiceStatus
*/
package info

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type GetInfoRequest struct {
}

func (m *GetInfoRequest) Reset()                    { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()               {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetInfoResponse struct {
	Version     string              `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Webservices []*WebServiceStatus `protobuf:"bytes,2,rep,name=webservices" json:"webservices,omitempty"`
}

func (m *GetInfoResponse) Reset()                    { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()               {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetInfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetInfoResponse) GetWebservices() []*WebServiceStatus {
	if m != nil {
		return m.Webservices
	}
	return nil
}

type WebServiceStatus struct {
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Status   string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Text     string `protobuf:"bytes,4,opt,name=text" json:"text,omitempty"`
}

func (m *WebServiceStatus) Reset()                    { *m = WebServiceStatus{} }
func (m *WebServiceStatus) String() string            { return proto.CompactTextString(m) }
func (*WebServiceStatus) ProtoMessage()               {}
func (*WebServiceStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WebServiceStatus) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *WebServiceStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WebServiceStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *WebServiceStatus) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInfoRequest)(nil), "info.GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "info.GetInfoResponse")
	proto.RegisterType((*WebServiceStatus)(nil), "info.WebServiceStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for InfoService service

type InfoServiceClient interface {
	Get(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type infoServiceClient struct {
	cc *grpc.ClientConn
}

func NewInfoServiceClient(cc *grpc.ClientConn) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) Get(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := grpc.Invoke(ctx, "/info.InfoService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InfoService service

type InfoServiceServer interface {
	Get(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

func RegisterInfoServiceServer(s *grpc.Server, srv InfoServiceServer) {
	s.RegisterService(&_InfoService_serviceDesc, srv)
}

func _InfoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.InfoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).Get(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "info.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _InfoService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}

func init() { proto.RegisterFile("info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4e, 0xc3, 0x40,
	0x0c, 0xc6, 0xd5, 0x26, 0xb4, 0xe0, 0x08, 0xa8, 0x2c, 0xa8, 0x4e, 0x11, 0x43, 0x95, 0xa9, 0x53,
	0x23, 0x95, 0x85, 0x95, 0x09, 0xb1, 0xa1, 0x74, 0x60, 0x4e, 0xc0, 0x54, 0x27, 0xc1, 0x5d, 0x88,
	0xdd, 0xc2, 0xcc, 0x2b, 0xf0, 0x68, 0xbc, 0x02, 0x0f, 0xc2, 0xfd, 0x09, 0x15, 0x74, 0xb3, 0x3f,
	0xff, 0xac, 0xcf, 0xfe, 0x00, 0xb4, 0x79, 0xb2, 0x8b, 0xb6, 0xb3, 0x62, 0x31, 0xf5, 0x75, 0x7e,
	0xb1, 0xb6, 0x76, 0xfd, 0x4c, 0x65, 0xdd, 0xea, 0xb2, 0x36, 0xc6, 0x4a, 0x2d, 0xda, 0x1a, 0x8e,
	0x4c, 0x31, 0x81, 0x93, 0x1b, 0x92, 0x5b, 0x07, 0x56, 0xf4, 0xba, 0x21, 0x96, 0x82, 0xe0, 0x74,
	0xa7, 0x70, 0xeb, 0x48, 0x42, 0x05, 0xe3, 0x2d, 0x75, 0xec, 0xd6, 0xd4, 0x60, 0x36, 0x98, 0x1f,
	0x55, 0xbf, 0x2d, 0x5e, 0x41, 0xf6, 0x46, 0x0d, 0x53, 0xb7, 0xd5, 0x0f, 0xc4, 0x6a, 0x38, 0x4b,
	0xe6, 0xd9, 0x72, 0xba, 0x08, 0x47, 0xdc, 0x53, 0xb3, 0x8a, 0x83, 0x95, 0xf3, 0xdc, 0x70, 0xf5,
	0x17, 0x2d, 0x0c, 0x4c, 0xf6, 0x01, 0xcc, 0xe1, 0x90, 0xcc, 0x63, 0x6b, 0xb5, 0x91, 0xde, 0x68,
	0xd7, 0x23, 0x42, 0x6a, 0xea, 0x17, 0x72, 0x16, 0x5e, 0x0f, 0x35, 0x4e, 0x61, 0xc4, 0x61, 0x53,
	0x25, 0x41, 0xed, 0x3b, 0xcf, 0x0a, 0xbd, 0x8b, 0x4a, 0x23, 0xeb, 0xeb, 0xe5, 0x1d, 0x64, 0xfe,
	0xa7, 0xde, 0x10, 0xaf, 0x21, 0x71, 0x5f, 0xe2, 0x59, 0x3c, 0xf5, 0x7f, 0x04, 0xf9, 0xf9, 0x9e,
	0x1a, 0x63, 0x28, 0x8e, 0x3f, 0xbe, 0xbe, 0x3f, 0x87, 0x63, 0x3c, 0x28, 0xfd, 0xb8, 0x19, 0x85,
	0x04, 0x2f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xec, 0xa0, 0x97, 0x73, 0x01, 0x00, 0x00,
}