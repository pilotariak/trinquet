// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Details
	Discipline
	Level
	League
	GetLeaguesRequest
	GetLeaguesResponse
	CreateLeagueResponse
	CreateLeagueRequest
	GetLeagueRequest
	GetLeagueResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

type Details struct {
	Website     string `protobuf:"bytes,1,opt,name=website" json:"website,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phonenumber string `protobuf:"bytes,4,opt,name=phonenumber" json:"phonenumber,omitempty"`
	Fax         string `protobuf:"bytes,5,opt,name=fax" json:"fax,omitempty"`
}

func (m *Details) Reset()                    { *m = Details{} }
func (m *Details) String() string            { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()               {}
func (*Details) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Details) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Details) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Details) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Details) GetPhonenumber() string {
	if m != nil {
		return m.Phonenumber
	}
	return ""
}

func (m *Details) GetFax() string {
	if m != nil {
		return m.Fax
	}
	return ""
}

type Discipline struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
}

func (m *Discipline) Reset()                    { *m = Discipline{} }
func (m *Discipline) String() string            { return proto.CompactTextString(m) }
func (*Discipline) ProtoMessage()               {}
func (*Discipline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Discipline) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Discipline) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type Level struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
}

func (m *Level) Reset()                    { *m = Level{} }
func (m *Level) String() string            { return proto.CompactTextString(m) }
func (*Level) ProtoMessage()               {}
func (*Level) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Level) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Level) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

// League define a pelota league
type League struct {
	Name        string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Details     *Details      `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
	Levels      []*Level      `protobuf:"bytes,3,rep,name=levels" json:"levels,omitempty"`
	Disciplines []*Discipline `protobuf:"bytes,4,rep,name=disciplines" json:"disciplines,omitempty"`
}

func (m *League) Reset()                    { *m = League{} }
func (m *League) String() string            { return proto.CompactTextString(m) }
func (*League) ProtoMessage()               {}
func (*League) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *League) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *League) GetDetails() *Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *League) GetLevels() []*Level {
	if m != nil {
		return m.Levels
	}
	return nil
}

func (m *League) GetDisciplines() []*Discipline {
	if m != nil {
		return m.Disciplines
	}
	return nil
}

type GetLeaguesRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetLeaguesRequest) Reset()                    { *m = GetLeaguesRequest{} }
func (m *GetLeaguesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLeaguesRequest) ProtoMessage()               {}
func (*GetLeaguesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetLeaguesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetLeaguesResponse struct {
	Leagues []*League `protobuf:"bytes,1,rep,name=leagues" json:"leagues,omitempty"`
}

func (m *GetLeaguesResponse) Reset()                    { *m = GetLeaguesResponse{} }
func (m *GetLeaguesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLeaguesResponse) ProtoMessage()               {}
func (*GetLeaguesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetLeaguesResponse) GetLeagues() []*League {
	if m != nil {
		return m.Leagues
	}
	return nil
}

type CreateLeagueResponse struct {
	Code   int32   `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	League *League `protobuf:"bytes,2,opt,name=league" json:"league,omitempty"`
}

func (m *CreateLeagueResponse) Reset()                    { *m = CreateLeagueResponse{} }
func (m *CreateLeagueResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateLeagueResponse) ProtoMessage()               {}
func (*CreateLeagueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateLeagueResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateLeagueResponse) GetLeague() *League {
	if m != nil {
		return m.League
	}
	return nil
}

type CreateLeagueRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Website string `protobuf:"bytes,2,opt,name=website" json:"website,omitempty"`
}

func (m *CreateLeagueRequest) Reset()                    { *m = CreateLeagueRequest{} }
func (m *CreateLeagueRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateLeagueRequest) ProtoMessage()               {}
func (*CreateLeagueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateLeagueRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateLeagueRequest) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

type GetLeagueRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetLeagueRequest) Reset()                    { *m = GetLeagueRequest{} }
func (m *GetLeagueRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLeagueRequest) ProtoMessage()               {}
func (*GetLeagueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetLeagueRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetLeagueResponse struct {
	League *League `protobuf:"bytes,1,opt,name=league" json:"league,omitempty"`
}

func (m *GetLeagueResponse) Reset()                    { *m = GetLeagueResponse{} }
func (m *GetLeagueResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLeagueResponse) ProtoMessage()               {}
func (*GetLeagueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetLeagueResponse) GetLeague() *League {
	if m != nil {
		return m.League
	}
	return nil
}

func init() {
	proto.RegisterType((*Details)(nil), "pb.Details")
	proto.RegisterType((*Discipline)(nil), "pb.Discipline")
	proto.RegisterType((*Level)(nil), "pb.Level")
	proto.RegisterType((*League)(nil), "pb.League")
	proto.RegisterType((*GetLeaguesRequest)(nil), "pb.GetLeaguesRequest")
	proto.RegisterType((*GetLeaguesResponse)(nil), "pb.GetLeaguesResponse")
	proto.RegisterType((*CreateLeagueResponse)(nil), "pb.CreateLeagueResponse")
	proto.RegisterType((*CreateLeagueRequest)(nil), "pb.CreateLeagueRequest")
	proto.RegisterType((*GetLeagueRequest)(nil), "pb.GetLeagueRequest")
	proto.RegisterType((*GetLeagueResponse)(nil), "pb.GetLeagueResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LeagueService service

type LeagueServiceClient interface {
	// List returns all available League
	List(ctx context.Context, in *GetLeaguesRequest, opts ...grpc.CallOption) (*GetLeaguesResponse, error)
	// Create creates a new league
	Create(ctx context.Context, in *CreateLeagueRequest, opts ...grpc.CallOption) (*CreateLeagueResponse, error)
	// Get return a league
	Get(ctx context.Context, in *GetLeagueRequest, opts ...grpc.CallOption) (*GetLeagueResponse, error)
}

type leagueServiceClient struct {
	cc *grpc.ClientConn
}

func NewLeagueServiceClient(cc *grpc.ClientConn) LeagueServiceClient {
	return &leagueServiceClient{cc}
}

func (c *leagueServiceClient) List(ctx context.Context, in *GetLeaguesRequest, opts ...grpc.CallOption) (*GetLeaguesResponse, error) {
	out := new(GetLeaguesResponse)
	err := grpc.Invoke(ctx, "/pb.LeagueService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leagueServiceClient) Create(ctx context.Context, in *CreateLeagueRequest, opts ...grpc.CallOption) (*CreateLeagueResponse, error) {
	out := new(CreateLeagueResponse)
	err := grpc.Invoke(ctx, "/pb.LeagueService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leagueServiceClient) Get(ctx context.Context, in *GetLeagueRequest, opts ...grpc.CallOption) (*GetLeagueResponse, error) {
	out := new(GetLeagueResponse)
	err := grpc.Invoke(ctx, "/pb.LeagueService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LeagueService service

type LeagueServiceServer interface {
	// List returns all available League
	List(context.Context, *GetLeaguesRequest) (*GetLeaguesResponse, error)
	// Create creates a new league
	Create(context.Context, *CreateLeagueRequest) (*CreateLeagueResponse, error)
	// Get return a league
	Get(context.Context, *GetLeagueRequest) (*GetLeagueResponse, error)
}

func RegisterLeagueServiceServer(s *grpc.Server, srv LeagueServiceServer) {
	s.RegisterService(&_LeagueService_serviceDesc, srv)
}

func _LeagueService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaguesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeagueServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LeagueService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeagueServiceServer).List(ctx, req.(*GetLeaguesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeagueService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLeagueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeagueServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LeagueService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeagueServiceServer).Create(ctx, req.(*CreateLeagueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeagueService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeagueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeagueServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LeagueService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeagueServiceServer).Get(ctx, req.(*GetLeagueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LeagueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LeagueService",
	HandlerType: (*LeagueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LeagueService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LeagueService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LeagueService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x95, 0x9d, 0x97, 0x7a, 0xad, 0x54, 0xe9, 0x4d, 0x5a, 0x4c, 0xc4, 0xa2, 0x8c, 0x78, 0x09,
	0x89, 0x18, 0xc2, 0x02, 0xa9, 0xdb, 0x56, 0x42, 0x42, 0x81, 0x85, 0xfb, 0x05, 0x76, 0x7c, 0x09,
	0x23, 0x39, 0xb6, 0xf1, 0x4c, 0x03, 0x12, 0x62, 0xc3, 0x86, 0x0f, 0x60, 0xc3, 0x7f, 0xf1, 0x0b,
	0xfc, 0x07, 0xcc, 0xcb, 0xc1, 0xe9, 0x03, 0x75, 0x13, 0xcd, 0xdc, 0x73, 0xee, 0xb9, 0xe7, 0x9e,
	0x78, 0x60, 0x28, 0xa8, 0xde, 0xf0, 0x25, 0xcd, 0xaa, 0xba, 0x94, 0x25, 0xfa, 0x55, 0x3a, 0xbd,
	0xb7, 0x2a, 0xcb, 0x55, 0x4e, 0x51, 0x52, 0xf1, 0x28, 0x29, 0x8a, 0x52, 0x26, 0x92, 0x97, 0x85,
	0xb0, 0x0c, 0xf6, 0xdd, 0x83, 0xc1, 0x19, 0xc9, 0x84, 0xe7, 0x02, 0x43, 0x18, 0x7c, 0xa2, 0x54,
	0x70, 0x49, 0xa1, 0x77, 0xec, 0x3d, 0xd9, 0x8b, 0x9b, 0xab, 0x46, 0x92, 0x2c, 0xab, 0x49, 0x88,
	0xd0, 0xb7, 0x88, 0xbb, 0xe2, 0x04, 0x7a, 0xb4, 0x56, 0xdd, 0x61, 0xc7, 0xd4, 0xed, 0x05, 0x8f,
	0x21, 0xa8, 0x3e, 0x94, 0x05, 0x15, 0x17, 0xeb, 0x94, 0xea, 0xb0, 0x6b, 0xb0, 0x76, 0x09, 0x47,
	0xd0, 0x79, 0x9f, 0x7c, 0x0e, 0x7b, 0x06, 0xd1, 0x47, 0x36, 0x07, 0x38, 0xe3, 0x62, 0xc9, 0xab,
	0x9c, 0x17, 0x84, 0xfb, 0xe0, 0xf3, 0xcc, 0xd9, 0x50, 0x27, 0x3d, 0x47, 0x72, 0x99, 0x93, 0x9b,
	0x6f, 0x2f, 0xec, 0x19, 0xf4, 0x16, 0xb4, 0xa1, 0xfc, 0x96, 0xf4, 0x9f, 0x1e, 0xf4, 0x17, 0x94,
	0xac, 0x2e, 0x08, 0x11, 0xba, 0x45, 0xb2, 0x6e, 0x16, 0x35, 0x67, 0x7c, 0x08, 0x83, 0xcc, 0x46,
	0x61, 0xda, 0x82, 0x79, 0x30, 0xab, 0xd2, 0x99, 0x4b, 0x27, 0x6e, 0x30, 0xbc, 0x0f, 0xfd, 0x5c,
	0x0f, 0x15, 0x6a, 0xe7, 0x8e, 0x62, 0xed, 0x69, 0x96, 0xb1, 0x11, 0x3b, 0x00, 0x9f, 0x43, 0x90,
	0x6d, 0x77, 0x11, 0x6a, 0x7f, 0xcd, 0xdb, 0x37, 0x6a, 0xdb, 0x72, 0xdc, 0xa6, 0xb0, 0xc7, 0x70,
	0xf0, 0x9a, 0xa4, 0x35, 0x27, 0x62, 0xfa, 0xa8, 0x7e, 0xe5, 0x75, 0x26, 0xd9, 0x09, 0x60, 0x9b,
	0x28, 0x2a, 0xf5, 0x5f, 0x12, 0x3e, 0x80, 0x41, 0x6e, 0x4b, 0x8a, 0xac, 0x87, 0x81, 0x35, 0xa5,
	0x4b, 0x71, 0x03, 0xb1, 0x77, 0x30, 0x39, 0xad, 0x29, 0x91, 0xe4, 0x80, 0xa6, 0x5b, 0xcd, 0x59,
	0x96, 0x99, 0x9d, 0xd3, 0x8b, 0xcd, 0x19, 0x99, 0xde, 0x52, 0xb3, 0x5c, 0x16, 0x6d, 0x41, 0x87,
	0xb0, 0x53, 0x18, 0xef, 0xea, 0xdd, 0x68, 0xbb, 0xfd, 0x6d, 0xf9, 0x3b, 0xdf, 0x16, 0x7b, 0x04,
	0xa3, 0xed, 0x42, 0xff, 0x5b, 0xfc, 0x55, 0x2b, 0xa1, 0xad, 0xf3, 0x7f, 0x2e, 0xbd, 0x9b, 0x5c,
	0xce, 0xff, 0x78, 0x30, 0xb4, 0xa5, 0x73, 0xfb, 0x38, 0xf0, 0x0d, 0x74, 0x17, 0x5c, 0x8d, 0x39,
	0xd4, 0xec, 0x2b, 0xb1, 0x4f, 0x8f, 0x2e, 0x97, 0xed, 0x30, 0x36, 0xfe, 0xf6, 0xeb, 0xf7, 0x0f,
	0x7f, 0x88, 0x41, 0xb4, 0x79, 0x11, 0xb9, 0x4c, 0xf1, 0x1c, 0xfa, 0x36, 0x03, 0xbc, 0xa3, 0xdb,
	0xae, 0xc9, 0x63, 0x1a, 0x5e, 0x05, 0x9c, 0xe2, 0x91, 0x51, 0x1c, 0xb1, 0xb6, 0xe2, 0x89, 0xf7,
	0x14, 0xdf, 0x42, 0x47, 0xcd, 0xc7, 0xc9, 0x8e, 0x91, 0x46, 0xee, 0xf0, 0x52, 0xd5, 0x69, 0xdd,
	0x35, 0x5a, 0x63, 0x3c, 0x68, 0xb4, 0x48, 0x44, 0x5f, 0x74, 0x72, 0x5f, 0xd3, 0xbe, 0x79, 0xeb,
	0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x74, 0x11, 0x2b, 0xa5, 0x1e, 0x04, 0x00, 0x00,
}
