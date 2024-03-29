// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package generated

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

// CommentProtobufServiceClient is the client API for CommentProtobufService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentProtobufServiceClient interface {
	// /Attachments
	ListComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*CommentResponse, error)
	GetComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	UpdateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*CommentResponse, error)
	DeleteComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
}

type commentProtobufServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentProtobufServiceClient(cc grpc.ClientConnInterface) CommentProtobufServiceClient {
	return &commentProtobufServiceClient{cc}
}

func (c *commentProtobufServiceClient) ListComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/worklen.proto.CommentProtobufService/ListComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentProtobufServiceClient) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/worklen.proto.CommentProtobufService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentProtobufServiceClient) GetComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/worklen.proto.CommentProtobufService/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentProtobufServiceClient) UpdateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/worklen.proto.CommentProtobufService/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentProtobufServiceClient) DeleteComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/worklen.proto.CommentProtobufService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentProtobufServiceServer is the server API for CommentProtobufService service.
// All implementations should embed UnimplementedCommentProtobufServiceServer
// for forward compatibility
type CommentProtobufServiceServer interface {
	// /Attachments
	ListComment(context.Context, *CommentRequest) (*CommentResponse, error)
	CreateComment(context.Context, *Comment) (*CommentResponse, error)
	GetComment(context.Context, *CommentRequest) (*CommentResponse, error)
	UpdateComment(context.Context, *Comment) (*CommentResponse, error)
	DeleteComment(context.Context, *CommentRequest) (*CommentResponse, error)
}

// UnimplementedCommentProtobufServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCommentProtobufServiceServer struct {
}

func (UnimplementedCommentProtobufServiceServer) ListComment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComment not implemented")
}
func (UnimplementedCommentProtobufServiceServer) CreateComment(context.Context, *Comment) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentProtobufServiceServer) GetComment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCommentProtobufServiceServer) UpdateComment(context.Context, *Comment) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentProtobufServiceServer) DeleteComment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}

// UnsafeCommentProtobufServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentProtobufServiceServer will
// result in compilation errors.
type UnsafeCommentProtobufServiceServer interface {
	mustEmbedUnimplementedCommentProtobufServiceServer()
}

func RegisterCommentProtobufServiceServer(s grpc.ServiceRegistrar, srv CommentProtobufServiceServer) {
	s.RegisterService(&CommentProtobufService_ServiceDesc, srv)
}

func _CommentProtobufService_ListComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentProtobufServiceServer).ListComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worklen.proto.CommentProtobufService/ListComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentProtobufServiceServer).ListComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentProtobufService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentProtobufServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worklen.proto.CommentProtobufService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentProtobufServiceServer).CreateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentProtobufService_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentProtobufServiceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worklen.proto.CommentProtobufService/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentProtobufServiceServer).GetComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentProtobufService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentProtobufServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worklen.proto.CommentProtobufService/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentProtobufServiceServer).UpdateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentProtobufService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentProtobufServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worklen.proto.CommentProtobufService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentProtobufServiceServer).DeleteComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentProtobufService_ServiceDesc is the grpc.ServiceDesc for CommentProtobufService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentProtobufService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "worklen.proto.CommentProtobufService",
	HandlerType: (*CommentProtobufServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListComment",
			Handler:    _CommentProtobufService_ListComment_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _CommentProtobufService_CreateComment_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _CommentProtobufService_GetComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _CommentProtobufService_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentProtobufService_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
