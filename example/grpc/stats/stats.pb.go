// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stats/stats.proto

package stats

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type RoundTripResponse struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoundTripResponse) Reset()         { *m = RoundTripResponse{} }
func (m *RoundTripResponse) String() string { return proto.CompactTextString(m) }
func (*RoundTripResponse) ProtoMessage()    {}
func (*RoundTripResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d537a116482db188, []int{0}
}

func (m *RoundTripResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundTripResponse.Unmarshal(m, b)
}
func (m *RoundTripResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundTripResponse.Marshal(b, m, deterministic)
}
func (m *RoundTripResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundTripResponse.Merge(m, src)
}
func (m *RoundTripResponse) XXX_Size() int {
	return xxx_messageInfo_RoundTripResponse.Size(m)
}
func (m *RoundTripResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundTripResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoundTripResponse proto.InternalMessageInfo

func (m *RoundTripResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*RoundTripResponse)(nil), "stats.RoundTripResponse")
}

func init() { proto.RegisterFile("stats/stats.proto", fileDescriptor_d537a116482db188) }

var fileDescriptor_d537a116482db188 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4b, 0x05, 0x31,
	0x10, 0xc4, 0xdf, 0x15, 0xb1, 0x48, 0xf7, 0x82, 0xc8, 0xe3, 0x6c, 0xe4, 0x2a, 0x2d, 0x4c, 0x40,
	0x11, 0xb1, 0x13, 0xc1, 0x2f, 0x10, 0xad, 0xec, 0x2e, 0x71, 0x2f, 0x17, 0xbc, 0x64, 0x43, 0xfe,
	0x88, 0x5a, 0xfb, 0xc1, 0xe5, 0x12, 0xb4, 0x91, 0x6b, 0x16, 0x66, 0xd8, 0x99, 0xdf, 0xb2, 0x74,
	0x9f, 0xf2, 0x98, 0x93, 0xa8, 0x93, 0x87, 0x88, 0x19, 0x19, 0xa9, 0xa2, 0x3f, 0x35, 0x88, 0x66,
	0x01, 0x51, 0x4d, 0x55, 0x26, 0x01, 0x2e, 0xe4, 0xcf, 0xb6, 0x33, 0x5c, 0xd0, 0xbd, 0xc4, 0xe2,
	0x5f, 0x9f, 0xa3, 0x0d, 0x12, 0x52, 0x40, 0x9f, 0x80, 0x1d, 0x53, 0xa2, 0xb1, 0xf8, 0x7c, 0xe8,
	0xce, 0xba, 0x73, 0x22, 0x9b, 0xb8, 0xfa, 0xee, 0x28, 0x79, 0x5a, 0x1b, 0xd9, 0x3d, 0xa5, 0x7f,
	0xa1, 0xc4, 0x4e, 0x78, 0x03, 0xf0, 0x5f, 0x00, 0x7f, 0x5c, 0x01, 0xfd, 0x81, 0xb7, 0x63, 0xfe,
	0xf5, 0x0f, 0x3b, 0x76, 0x47, 0x89, 0x84, 0x04, 0x79, 0x33, 0xbc, 0xe1, 0x0f, 0xbb, 0x87, 0xdb,
	0x97, 0x1b, 0x63, 0xf3, 0x5c, 0x14, 0xd7, 0xe8, 0xc4, 0xb4, 0xc0, 0xc7, 0x57, 0x29, 0x42, 0x81,
	0xd7, 0xb3, 0x1b, 0xe3, 0x9b, 0x70, 0x56, 0x47, 0xbc, 0x4c, 0x10, 0xdf, 0xad, 0x06, 0x61, 0x62,
	0xd0, 0xed, 0x29, 0xea, 0xa8, 0x56, 0x5d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x7c, 0x45,
	0xcc, 0x2a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsClient interface {
	RoundTrips(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RoundTripResponse, error)
	Reset(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type statsClient struct {
	cc *grpc.ClientConn
}

func NewStatsClient(cc *grpc.ClientConn) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) RoundTrips(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RoundTripResponse, error) {
	out := new(RoundTripResponse)
	err := c.cc.Invoke(ctx, "/stats.Stats/RoundTrips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Reset(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/stats.Stats/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServer is the server API for Stats service.
type StatsServer interface {
	RoundTrips(context.Context, *empty.Empty) (*RoundTripResponse, error)
	Reset(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterStatsServer(s *grpc.Server, srv StatsServer) {
	s.RegisterService(&_Stats_serviceDesc, srv)
}

func _Stats_RoundTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).RoundTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.Stats/RoundTrips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).RoundTrips(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.Stats/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Reset(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stats.Stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoundTrips",
			Handler:    _Stats_RoundTrips_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Stats_Reset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stats/stats.proto",
}
