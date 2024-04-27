// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: grpc/post_service.proto

package post_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostService_CreatePost_FullMethodName        = "/PostService/CreatePost"
	PostService_GetPost_FullMethodName           = "/PostService/GetPost"
	PostService_GetPaginatedPosts_FullMethodName = "/PostService/GetPaginatedPosts"
	PostService_UpdatePost_FullMethodName        = "/PostService/UpdatePost"
	PostService_DeletePost_FullMethodName        = "/PostService/DeletePost"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*PostId, error)
	GetPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	GetPaginatedPosts(ctx context.Context, in *PaginatedPostRequest, opts ...grpc.CallOption) (PostService_GetPaginatedPostsClient, error)
	UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Error, error)
	DeletePost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Error, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*PostId, error) {
	out := new(PostId)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, PostService_GetPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPaginatedPosts(ctx context.Context, in *PaginatedPostRequest, opts ...grpc.CallOption) (PostService_GetPaginatedPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PostService_ServiceDesc.Streams[0], PostService_GetPaginatedPosts_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &postServiceGetPaginatedPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PostService_GetPaginatedPostsClient interface {
	Recv() (*Post, error)
	grpc.ClientStream
}

type postServiceGetPaginatedPostsClient struct {
	grpc.ClientStream
}

func (x *postServiceGetPaginatedPostsClient) Recv() (*Post, error) {
	m := new(Post)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, PostService_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, PostService_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	CreatePost(context.Context, *Post) (*PostId, error)
	GetPost(context.Context, *PostRequest) (*PostResponse, error)
	GetPaginatedPosts(*PaginatedPostRequest, PostService_GetPaginatedPostsServer) error
	UpdatePost(context.Context, *Post) (*Error, error)
	DeletePost(context.Context, *PostRequest) (*Error, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) CreatePost(context.Context, *Post) (*PostId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) GetPost(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServiceServer) GetPaginatedPosts(*PaginatedPostRequest, PostService_GetPaginatedPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPaginatedPosts not implemented")
}
func (UnimplementedPostServiceServer) UpdatePost(context.Context, *Post) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *PostRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPaginatedPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PaginatedPostRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostServiceServer).GetPaginatedPosts(m, &postServiceGetPaginatedPostsServer{stream})
}

type PostService_GetPaginatedPostsServer interface {
	Send(*Post) error
	grpc.ServerStream
}

type postServiceGetPaginatedPostsServer struct {
	grpc.ServerStream
}

func (x *postServiceGetPaginatedPostsServer) Send(m *Post) error {
	return x.ServerStream.SendMsg(m)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPaginatedPosts",
			Handler:       _PostService_GetPaginatedPosts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/post_service.proto",
}
