// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: internal/pb/analytics.proto

package pb

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

// AnalyticsClient is the client API for Analytics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsClient interface {
	SendEvent(ctx context.Context, in *Event_Request, opts ...grpc.CallOption) (*Event_Response, error)
}

type analyticsClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsClient(cc grpc.ClientConnInterface) AnalyticsClient {
	return &analyticsClient{cc}
}

func (c *analyticsClient) SendEvent(ctx context.Context, in *Event_Request, opts ...grpc.CallOption) (*Event_Response, error) {
	out := new(Event_Response)
	err := c.cc.Invoke(ctx, "/cloudquery.backend.analytics.Analytics/SendEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServer is the server API for Analytics service.
// All implementations must embed UnimplementedAnalyticsServer
// for forward compatibility
type AnalyticsServer interface {
	SendEvent(context.Context, *Event_Request) (*Event_Response, error)
	mustEmbedUnimplementedAnalyticsServer()
}

// UnimplementedAnalyticsServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServer struct {
}

func (UnimplementedAnalyticsServer) SendEvent(context.Context, *Event_Request) (*Event_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEvent not implemented")
}
func (UnimplementedAnalyticsServer) mustEmbedUnimplementedAnalyticsServer() {}

// UnsafeAnalyticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsServer will
// result in compilation errors.
type UnsafeAnalyticsServer interface {
	mustEmbedUnimplementedAnalyticsServer()
}

func RegisterAnalyticsServer(s grpc.ServiceRegistrar, srv AnalyticsServer) {
	s.RegisterService(&Analytics_ServiceDesc, srv)
}

func _Analytics_SendEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).SendEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudquery.backend.analytics.Analytics/SendEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).SendEvent(ctx, req.(*Event_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Analytics_ServiceDesc is the grpc.ServiceDesc for Analytics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analytics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudquery.backend.analytics.Analytics",
	HandlerType: (*AnalyticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEvent",
			Handler:    _Analytics_SendEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pb/analytics.proto",
}
