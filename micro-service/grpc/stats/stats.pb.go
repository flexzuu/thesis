// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stats/stats.proto

package stats

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2e, 0x49, 0x2c,
	0x29, 0xd6, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x92, 0x26,
	0x97, 0x60, 0x50, 0x7e, 0x69, 0x5e, 0x4a, 0x48, 0x51, 0x66, 0x41, 0x50, 0x6a, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x72, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x6b, 0x10, 0x84, 0xe3, 0x64, 0x1e, 0x65, 0x9a, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x9f, 0x96, 0x93, 0x5a, 0x51, 0x55, 0x5a, 0xaa, 0x9f, 0x94, 0x9a, 0x97, 0x9c,
	0x91, 0x9b, 0x58, 0x94, 0xad, 0x9f, 0x9b, 0x99, 0x5c, 0x94, 0xaf, 0x5b, 0x9c, 0x5a, 0x54, 0x96,
	0x99, 0x9c, 0xaa, 0x9f, 0x5e, 0x54, 0x90, 0x0c, 0xb1, 0x30, 0x89, 0x0d, 0x6c, 0xa3, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x60, 0x22, 0x46, 0xb2, 0x86, 0x00, 0x00, 0x00,
}
