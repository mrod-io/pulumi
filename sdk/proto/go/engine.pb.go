// Code generated by protoc-gen-go.
// source: engine.proto
// DO NOT EDIT!

package lumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// LogSeverity is the severity level of a log message.  Errors are fatal; all others are informational.
type LogSeverity int32

const (
	LogSeverity_DEBUG   LogSeverity = 0
	LogSeverity_INFO    LogSeverity = 1
	LogSeverity_WARNING LogSeverity = 2
	LogSeverity_ERROR   LogSeverity = 3
)

var LogSeverity_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
}
var LogSeverity_value = map[string]int32{
	"DEBUG":   0,
	"INFO":    1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x LogSeverity) String() string {
	return proto.EnumName(LogSeverity_name, int32(x))
}
func (LogSeverity) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type QueryResourcesResponse struct {
	Obj *google_protobuf.Struct `protobuf:"bytes,1,opt,name=obj" json:"obj,omitempty"`
}

func (m *QueryResourcesResponse) Reset()                    { *m = QueryResourcesResponse{} }
func (m *QueryResourcesResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResourcesResponse) ProtoMessage()               {}
func (*QueryResourcesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *QueryResourcesResponse) GetObj() *google_protobuf.Struct {
	if m != nil {
		return m.Obj
	}
	return nil
}

type LogRequest struct {
	Severity LogSeverity `protobuf:"varint,1,opt,name=severity,enum=lumirpc.LogSeverity" json:"severity,omitempty"`
	Message  string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *LogRequest) Reset()                    { *m = LogRequest{} }
func (m *LogRequest) String() string            { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()               {}
func (*LogRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *LogRequest) GetSeverity() LogSeverity {
	if m != nil {
		return m.Severity
	}
	return LogSeverity_DEBUG
}

func (m *LogRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReadLocationRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *ReadLocationRequest) Reset()                    { *m = ReadLocationRequest{} }
func (m *ReadLocationRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadLocationRequest) ProtoMessage()               {}
func (*ReadLocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ReadLocationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ReadLocationsRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *ReadLocationsRequest) Reset()                    { *m = ReadLocationsRequest{} }
func (m *ReadLocationsRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadLocationsRequest) ProtoMessage()               {}
func (*ReadLocationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ReadLocationsRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ReadLocationsResponse struct {
	Properties *google_protobuf.Struct `protobuf:"bytes,1,opt,name=properties" json:"properties,omitempty"`
}

func (m *ReadLocationsResponse) Reset()                    { *m = ReadLocationsResponse{} }
func (m *ReadLocationsResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadLocationsResponse) ProtoMessage()               {}
func (*ReadLocationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ReadLocationsResponse) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryResourcesResponse)(nil), "lumirpc.QueryResourcesResponse")
	proto.RegisterType((*LogRequest)(nil), "lumirpc.LogRequest")
	proto.RegisterType((*ReadLocationRequest)(nil), "lumirpc.ReadLocationRequest")
	proto.RegisterType((*ReadLocationsRequest)(nil), "lumirpc.ReadLocationsRequest")
	proto.RegisterType((*ReadLocationsResponse)(nil), "lumirpc.ReadLocationsResponse")
	proto.RegisterEnum("lumirpc.LogSeverity", LogSeverity_name, LogSeverity_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Engine service

type EngineClient interface {
	// Log logs a global message in the engine, including errors and warnings.
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// ReadLocation reads the value from a location identified by a token in the current program.
	ReadLocation(ctx context.Context, in *ReadLocationRequest, opts ...grpc.CallOption) (*google_protobuf.Value, error)
	// ReadLocations reads a list of static or module variables from a single parent token.
	ReadLocations(ctx context.Context, in *ReadLocationsRequest, opts ...grpc.CallOption) (*ReadLocationsResponse, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/lumirpc.Engine/Log", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) ReadLocation(ctx context.Context, in *ReadLocationRequest, opts ...grpc.CallOption) (*google_protobuf.Value, error) {
	out := new(google_protobuf.Value)
	err := grpc.Invoke(ctx, "/lumirpc.Engine/ReadLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) ReadLocations(ctx context.Context, in *ReadLocationsRequest, opts ...grpc.CallOption) (*ReadLocationsResponse, error) {
	out := new(ReadLocationsResponse)
	err := grpc.Invoke(ctx, "/lumirpc.Engine/ReadLocations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Engine service

type EngineServer interface {
	// Log logs a global message in the engine, including errors and warnings.
	Log(context.Context, *LogRequest) (*google_protobuf1.Empty, error)
	// ReadLocation reads the value from a location identified by a token in the current program.
	ReadLocation(context.Context, *ReadLocationRequest) (*google_protobuf.Value, error)
	// ReadLocations reads a list of static or module variables from a single parent token.
	ReadLocations(context.Context, *ReadLocationsRequest) (*ReadLocationsResponse, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.Engine/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_ReadLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).ReadLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.Engine/ReadLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).ReadLocation(ctx, req.(*ReadLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_ReadLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).ReadLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lumirpc.Engine/ReadLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).ReadLocations(ctx, req.(*ReadLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lumirpc.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Engine_Log_Handler,
		},
		{
			MethodName: "ReadLocation",
			Handler:    _Engine_ReadLocation_Handler,
		},
		{
			MethodName: "ReadLocations",
			Handler:    _Engine_ReadLocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "engine.proto",
}

func init() { proto.RegisterFile("engine.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x4b, 0x6f, 0xda, 0x40,
	0x14, 0x85, 0x6d, 0x28, 0xaf, 0x0b, 0xad, 0xac, 0x81, 0x52, 0xe4, 0xd2, 0x0a, 0x79, 0x45, 0x1f,
	0x32, 0x15, 0xad, 0x54, 0xa9, 0xbb, 0x26, 0x31, 0x08, 0xc9, 0x02, 0x32, 0x28, 0x8f, 0xad, 0x71,
	0x6e, 0x2c, 0x27, 0xe0, 0x71, 0x66, 0xc6, 0x91, 0xf8, 0xa7, 0xf9, 0x39, 0x11, 0x7e, 0x20, 0x87,
	0x90, 0x64, 0x39, 0x73, 0xbe, 0x39, 0x73, 0xcf, 0xb9, 0xd0, 0xc0, 0xc0, 0xf3, 0x03, 0x34, 0x43,
	0xce, 0x24, 0x23, 0x95, 0x55, 0xb4, 0xf6, 0x79, 0xe8, 0xea, 0x9f, 0x3d, 0xc6, 0xbc, 0x15, 0x0e,
	0xe2, 0xeb, 0x65, 0x74, 0x3d, 0xc0, 0x75, 0x28, 0x37, 0x09, 0xa5, 0x77, 0xf7, 0x45, 0x21, 0x79,
	0xe4, 0xca, 0x44, 0x35, 0x8e, 0xa1, 0x7d, 0x1a, 0x21, 0xdf, 0x50, 0x14, 0x2c, 0xe2, 0x2e, 0x0a,
	0x8a, 0x22, 0x64, 0x81, 0x40, 0xf2, 0x0d, 0x8a, 0x6c, 0x79, 0xd3, 0x51, 0x7b, 0x6a, 0xbf, 0x3e,
	0xfc, 0x64, 0x26, 0x2e, 0x66, 0xe6, 0x62, 0x2e, 0x62, 0x17, 0xba, 0x65, 0x8c, 0x4b, 0x00, 0x9b,
	0x79, 0x14, 0xef, 0x22, 0x14, 0x92, 0xfc, 0x82, 0xaa, 0xc0, 0x7b, 0xe4, 0xbe, 0xdc, 0xc4, 0xaf,
	0x3f, 0x0c, 0x5b, 0x66, 0x3a, 0xa9, 0x69, 0x33, 0x6f, 0x91, 0x6a, 0x74, 0x47, 0x91, 0x0e, 0x54,
	0xd6, 0x28, 0x84, 0xe3, 0x61, 0xa7, 0xd0, 0x53, 0xfb, 0x35, 0x9a, 0x1d, 0x8d, 0x1f, 0xd0, 0xa4,
	0xe8, 0x5c, 0xd9, 0xcc, 0x75, 0xa4, 0xcf, 0x82, 0xec, 0x8b, 0x16, 0x94, 0x24, 0xbb, 0xc5, 0x20,
	0xf6, 0xaf, 0xd1, 0xe4, 0x60, 0xfc, 0x84, 0x56, 0x1e, 0x16, 0xaf, 0xd3, 0x73, 0xf8, 0xb8, 0x47,
	0xa7, 0xc1, 0xff, 0x02, 0x84, 0x9c, 0x85, 0xc8, 0xa5, 0x8f, 0xe2, 0xad, 0xfc, 0x39, 0xf4, 0xfb,
	0x3f, 0xa8, 0xe7, 0xf2, 0x91, 0x1a, 0x94, 0x4e, 0xac, 0xa3, 0xb3, 0xb1, 0xa6, 0x90, 0x2a, 0xbc,
	0x9b, 0x4c, 0x47, 0x33, 0x4d, 0x25, 0x75, 0xa8, 0x5c, 0xfc, 0xa7, 0xd3, 0xc9, 0x74, 0xac, 0x15,
	0xb6, 0x84, 0x45, 0xe9, 0x8c, 0x6a, 0xc5, 0xe1, 0x83, 0x0a, 0x65, 0x2b, 0x5e, 0x2e, 0xf9, 0x03,
	0x45, 0x9b, 0x79, 0xa4, 0x99, 0x2f, 0x2d, 0x8d, 0xa2, 0xb7, 0x9f, 0xcd, 0x61, 0x6d, 0x57, 0x6d,
	0x28, 0x64, 0x04, 0x8d, 0x7c, 0x1c, 0xd2, 0xdd, 0x3d, 0x3f, 0x50, 0xe0, 0x01, 0x9f, 0x73, 0x67,
	0x15, 0xa1, 0xa1, 0x90, 0x39, 0xbc, 0x7f, 0x52, 0x0b, 0xf9, 0x72, 0xd0, 0x28, 0x2b, 0x57, 0xff,
	0xfa, 0x92, 0x9c, 0xb4, 0x69, 0x28, 0xcb, 0x72, 0xfc, 0xc7, 0xef, 0xc7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x36, 0xfb, 0x2b, 0xc6, 0xbd, 0x02, 0x00, 0x00,
}
