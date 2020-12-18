// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package first_website_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FirstWebsiteServiceClient is the client API for FirstWebsiteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FirstWebsiteServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type firstWebsiteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFirstWebsiteServiceClient(cc grpc.ClientConnInterface) FirstWebsiteServiceClient {
	return &firstWebsiteServiceClient{cc}
}

func (c *firstWebsiteServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/first_website_proto.FirstWebsiteService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirstWebsiteServiceServer is the server API for FirstWebsiteService service.
// All implementations must embed UnimplementedFirstWebsiteServiceServer
// for forward compatibility
type FirstWebsiteServiceServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	mustEmbedUnimplementedFirstWebsiteServiceServer()
}

// UnimplementedFirstWebsiteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFirstWebsiteServiceServer struct {
}

func (UnimplementedFirstWebsiteServiceServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedFirstWebsiteServiceServer) mustEmbedUnimplementedFirstWebsiteServiceServer() {}

// UnsafeFirstWebsiteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FirstWebsiteServiceServer will
// result in compilation errors.
type UnsafeFirstWebsiteServiceServer interface {
	mustEmbedUnimplementedFirstWebsiteServiceServer()
}

func RegisterFirstWebsiteServiceServer(s grpc.ServiceRegistrar, srv FirstWebsiteServiceServer) {
	s.RegisterService(&_FirstWebsiteService_serviceDesc, srv)
}

func _FirstWebsiteService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirstWebsiteServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/first_website_proto.FirstWebsiteService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirstWebsiteServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FirstWebsiteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "first_website_proto.FirstWebsiteService",
	HandlerType: (*FirstWebsiteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _FirstWebsiteService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "first_website.proto",
}
