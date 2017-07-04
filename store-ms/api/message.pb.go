// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	SaveRequest
	SaveResult
*/
package api

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

type SaveRequest struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SaveRequest) Reset()                    { *m = SaveRequest{} }
func (m *SaveRequest) String() string            { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()               {}
func (*SaveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SaveRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *SaveRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SaveResult struct {
	IsOk  bool   `protobuf:"varint,1,opt,name=isOk" json:"isOk,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SaveResult) Reset()                    { *m = SaveResult{} }
func (m *SaveResult) String() string            { return proto.CompactTextString(m) }
func (*SaveResult) ProtoMessage()               {}
func (*SaveResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SaveResult) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func (m *SaveResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*SaveRequest)(nil), "api.SaveRequest")
	proto.RegisterType((*SaveResult)(nil), "api.SaveResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StoreService service

type StoreServiceClient interface {
	SaveFile(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResult, error)
}

type storeServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoreServiceClient(cc *grpc.ClientConn) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) SaveFile(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResult, error) {
	out := new(SaveResult)
	err := grpc.Invoke(ctx, "/api.StoreService/SaveFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoreService service

type StoreServiceServer interface {
	SaveFile(context.Context, *SaveRequest) (*SaveResult, error)
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_SaveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).SaveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StoreService/SaveFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).SaveFile(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveFile",
			Handler:    _StoreService_SaveFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x4d, 0xcb, 0x82, 0x40,
	0x10, 0xc7, 0x1f, 0x9f, 0x5e, 0xd0, 0xc9, 0x28, 0x86, 0x0e, 0xe2, 0x49, 0x3c, 0x79, 0x32, 0x28,
	0xe8, 0x16, 0xdd, 0xba, 0x06, 0xeb, 0x27, 0x98, 0x6a, 0x8a, 0x25, 0x6d, 0x6d, 0x77, 0xf5, 0xf3,
	0xc7, 0xae, 0x10, 0xde, 0xfe, 0xbf, 0x61, 0x7e, 0xf3, 0x02, 0xcb, 0x86, 0x8d, 0xa1, 0x27, 0x97,
	0xad, 0x56, 0x56, 0xe1, 0x84, 0x5a, 0x99, 0x1f, 0x61, 0x51, 0x51, 0xcf, 0x82, 0x3f, 0x1d, 0x1b,
	0x8b, 0x29, 0x84, 0x0f, 0x59, 0xf3, 0x9b, 0x1a, 0x4e, 0x82, 0x2c, 0x28, 0x22, 0xf1, 0x63, 0x44,
	0x98, 0xde, 0xc9, 0x52, 0xf2, 0x9f, 0x05, 0x45, 0x2c, 0x7c, 0xce, 0x0f, 0x00, 0x83, 0x6e, 0xba,
	0xda, 0xba, 0x0e, 0x69, 0x2e, 0x2f, 0x6f, 0x86, 0xc2, 0x67, 0xdc, 0xc0, 0x8c, 0xb5, 0x56, 0xda,
	0x6b, 0x91, 0x18, 0x60, 0x77, 0x82, 0xb8, 0xb2, 0x4a, 0x73, 0xc5, 0xba, 0x97, 0x37, 0xc6, 0x2d,
	0x84, 0x6e, 0xce, 0x59, 0xd6, 0x8c, 0xeb, 0x92, 0x5a, 0x59, 0x8e, 0xae, 0x4a, 0x57, 0xa3, 0x8a,
	0x5b, 0x94, 0xff, 0x5d, 0xe7, 0xfe, 0x87, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xed, 0x91,
	0xf1, 0xd4, 0x00, 0x00, 0x00,
}
