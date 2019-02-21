// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post.proto

package post

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

// The request message containing the id
type GetPostRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostRequest) Reset()         { *m = GetPostRequest{} }
func (m *GetPostRequest) String() string { return proto.CompactTextString(m) }
func (*GetPostRequest) ProtoMessage()    {}
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{0}
}

func (m *GetPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPostRequest.Unmarshal(m, b)
}
func (m *GetPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPostRequest.Marshal(b, m, deterministic)
}
func (m *GetPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostRequest.Merge(m, src)
}
func (m *GetPostRequest) XXX_Size() int {
	return xxx_messageInfo_GetPostRequest.Size(m)
}
func (m *GetPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostRequest proto.InternalMessageInfo

func (m *GetPostRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The response message containing the post
type GetPostResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostResponse) Reset()         { *m = GetPostResponse{} }
func (m *GetPostResponse) String() string { return proto.CompactTextString(m) }
func (*GetPostResponse) ProtoMessage()    {}
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{1}
}

func (m *GetPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPostResponse.Unmarshal(m, b)
}
func (m *GetPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPostResponse.Marshal(b, m, deterministic)
}
func (m *GetPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostResponse.Merge(m, src)
}
func (m *GetPostResponse) XXX_Size() int {
	return xxx_messageInfo_GetPostResponse.Size(m)
}
func (m *GetPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostResponse proto.InternalMessageInfo

func (m *GetPostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

// The request message containing the id
type CreatePostRequest struct {
	AuthorID             int64    `protobuf:"varint,1,opt,name=authorID,proto3" json:"authorID,omitempty"`
	Headline             string   `protobuf:"bytes,2,opt,name=headline,proto3" json:"headline,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{2}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetAuthorID() int64 {
	if m != nil {
		return m.AuthorID
	}
	return 0
}

func (m *CreatePostRequest) GetHeadline() string {
	if m != nil {
		return m.Headline
	}
	return ""
}

func (m *CreatePostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// The response message containing the post
type CreatePostResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostResponse) Reset()         { *m = CreatePostResponse{} }
func (m *CreatePostResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePostResponse) ProtoMessage()    {}
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{3}
}

func (m *CreatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostResponse.Unmarshal(m, b)
}
func (m *CreatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostResponse.Marshal(b, m, deterministic)
}
func (m *CreatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostResponse.Merge(m, src)
}
func (m *CreatePostResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePostResponse.Size(m)
}
func (m *CreatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostResponse proto.InternalMessageInfo

func (m *CreatePostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

// The request message containing the id
type DeletePostRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRequest) Reset()         { *m = DeletePostRequest{} }
func (m *DeletePostRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePostRequest) ProtoMessage()    {}
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{4}
}

func (m *DeletePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRequest.Unmarshal(m, b)
}
func (m *DeletePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRequest.Marshal(b, m, deterministic)
}
func (m *DeletePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRequest.Merge(m, src)
}
func (m *DeletePostRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePostRequest.Size(m)
}
func (m *DeletePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRequest proto.InternalMessageInfo

func (m *DeletePostRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// The response message containing the post
type DeletePostResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostResponse) Reset()         { *m = DeletePostResponse{} }
func (m *DeletePostResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePostResponse) ProtoMessage()    {}
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{5}
}

func (m *DeletePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostResponse.Unmarshal(m, b)
}
func (m *DeletePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostResponse.Marshal(b, m, deterministic)
}
func (m *DeletePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostResponse.Merge(m, src)
}
func (m *DeletePostResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePostResponse.Size(m)
}
func (m *DeletePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostResponse proto.InternalMessageInfo

type Post struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	AuthorID             int64    `protobuf:"varint,2,opt,name=authorID,proto3" json:"authorID,omitempty"`
	Headline             string   `protobuf:"bytes,3,opt,name=headline,proto3" json:"headline,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{6}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Post) GetAuthorID() int64 {
	if m != nil {
		return m.AuthorID
	}
	return 0
}

func (m *Post) GetHeadline() string {
	if m != nil {
		return m.Headline
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPostRequest)(nil), "post.GetPostRequest")
	proto.RegisterType((*GetPostResponse)(nil), "post.GetPostResponse")
	proto.RegisterType((*CreatePostRequest)(nil), "post.CreatePostRequest")
	proto.RegisterType((*CreatePostResponse)(nil), "post.CreatePostResponse")
	proto.RegisterType((*DeletePostRequest)(nil), "post.DeletePostRequest")
	proto.RegisterType((*DeletePostResponse)(nil), "post.DeletePostResponse")
	proto.RegisterType((*Post)(nil), "post.Post")
}

func init() { proto.RegisterFile("post.proto", fileDescriptor_e114ad14deab1dd1) }

var fileDescriptor_e114ad14deab1dd1 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x0b, 0x25, 0x54, 0xa7, 0x49, 0x4d, 0x37, 0x8d, 0x12, 0x4c, 0x0c, 0x59, 0x2f, 0xbd,
	0x08, 0xb1, 0x7a, 0xf3, 0xa6, 0x35, 0x4d, 0x6f, 0x06, 0x9f, 0x80, 0x96, 0xb1, 0x90, 0x20, 0x8b,
	0x30, 0x98, 0xf4, 0x95, 0x7c, 0x4a, 0xb3, 0xbb, 0x42, 0xa0, 0x8d, 0xc6, 0xdb, 0xce, 0xec, 0xbf,
	0xff, 0xcc, 0xb7, 0x33, 0x00, 0x85, 0xa8, 0xc8, 0x2f, 0x4a, 0x41, 0x82, 0x59, 0xf2, 0xec, 0x5e,
	0xee, 0x84, 0xd8, 0x65, 0x18, 0xa8, 0xdc, 0xa6, 0x7e, 0x0b, 0xf0, 0xbd, 0xa0, 0xbd, 0x96, 0x70,
	0x0f, 0x26, 0x2b, 0xa4, 0x17, 0x51, 0x51, 0x88, 0x1f, 0x35, 0x56, 0xc4, 0x26, 0x60, 0xa6, 0xb1,
	0x63, 0x78, 0xc6, 0x7c, 0x18, 0x9a, 0x69, 0xcc, 0x6f, 0xe1, 0xac, 0x55, 0x54, 0x85, 0xc8, 0x2b,
	0x64, 0x57, 0xa0, 0x9c, 0x95, 0x68, 0xbc, 0x00, 0x5f, 0x95, 0x54, 0x0a, 0x95, 0xe7, 0x08, 0xd3,
	0xa7, 0x12, 0x23, 0xc2, 0xae, 0xaf, 0x0b, 0x27, 0x51, 0x4d, 0x89, 0x28, 0xd7, 0xcb, 0x1f, 0xf7,
	0x36, 0x96, 0x77, 0x09, 0x46, 0x71, 0x96, 0xe6, 0xe8, 0x98, 0x9e, 0x31, 0x3f, 0x0d, 0xdb, 0x98,
	0x39, 0x30, 0xda, 0x8a, 0x9c, 0x30, 0x27, 0x67, 0xa8, 0xae, 0x9a, 0x90, 0xdf, 0x03, 0xeb, 0x96,
	0xf9, 0x67, 0x73, 0xd7, 0x30, 0x5d, 0x62, 0x86, 0xfd, 0xe6, 0x0e, 0xa1, 0x67, 0xc0, 0xba, 0x22,
	0x6d, 0xcd, 0x13, 0xb0, 0x64, 0x2c, 0xd5, 0x2d, 0x84, 0xa9, 0xdb, 0x6f, 0xd1, 0xcc, 0x3f, 0xd0,
	0x86, 0xbf, 0xa3, 0x59, 0x3d, 0xb4, 0xc5, 0x97, 0x01, 0x63, 0x59, 0xea, 0x15, 0xcb, 0xcf, 0x74,
	0x8b, 0xec, 0x06, 0x46, 0x2b, 0xa4, 0xc7, 0xfd, 0x3a, 0x66, 0x33, 0x4d, 0xd4, 0x9f, 0x9a, 0xdb,
	0xe1, 0xe4, 0x03, 0x16, 0x80, 0xad, 0x7f, 0x86, 0x5d, 0xe8, 0xfc, 0xd1, 0x38, 0x0e, 0x1e, 0x3c,
	0x80, 0xad, 0x79, 0x9b, 0x07, 0x47, 0x5f, 0xe4, 0x9e, 0xfb, 0x7a, 0x8f, 0xfc, 0x66, 0x8f, 0xfc,
	0x67, 0xb9, 0x47, 0x7c, 0xb0, 0xb1, 0x55, 0xe6, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x06,
	0xe9, 0xde, 0x7b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	// Get a single Post by id
	GetById(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error)
	// Create a single Post validates if the passed AuthorId is valid
	Create(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error)
	// Delete a single Post by id
	Delete(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type postServiceClient struct {
	cc *grpc.ClientConn
}

func NewPostServiceClient(cc *grpc.ClientConn) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) GetById(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/post.PostService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Create(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/post.PostService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Delete(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/post.PostService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	// Get a single Post by id
	GetById(context.Context, *GetPostRequest) (*Post, error)
	// Create a single Post validates if the passed AuthorId is valid
	Create(context.Context, *CreatePostRequest) (*Post, error)
	// Delete a single Post by id
	Delete(context.Context, *DeletePostRequest) (*empty.Empty, error)
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetById(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Create(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Delete(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _PostService_GetById_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PostService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PostService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}