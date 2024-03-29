// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: routeCo/routeCo.proto

package routeco

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

// RouteCoClient is the client API for RouteCo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteCoClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteCo_ListFeaturesClient, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteCo_RecordRouteClient, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteCo_RouteChatClient, error)
}

type routeCoClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteCoClient(cc grpc.ClientConnInterface) RouteCoClient {
	return &routeCoClient{cc}
}

func (c *routeCoClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/routeco.RouteCo/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeCoClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteCo_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteCo_ServiceDesc.Streams[0], "/routeco.RouteCo/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeCoListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteCo_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeCoListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeCoListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeCoClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteCo_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteCo_ServiceDesc.Streams[1], "/routeco.RouteCo/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeCoRecordRouteClient{stream}
	return x, nil
}

type RouteCo_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeCoRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeCoRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeCoRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeCoClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteCo_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteCo_ServiceDesc.Streams[2], "/routeco.RouteCo/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeCoRouteChatClient{stream}
	return x, nil
}

type RouteCo_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeCoRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeCoRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeCoRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteCoServer is the server API for RouteCo service.
// All implementations must embed UnimplementedRouteCoServer
// for forward compatibility
type RouteCoServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *Point) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(*Rectangle, RouteCo_ListFeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(RouteCo_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(RouteCo_RouteChatServer) error
	mustEmbedUnimplementedRouteCoServer()
}

// UnimplementedRouteCoServer must be embedded to have forward compatible implementations.
type UnimplementedRouteCoServer struct {
}

func (UnimplementedRouteCoServer) GetFeature(context.Context, *Point) (*Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedRouteCoServer) ListFeatures(*Rectangle, RouteCo_ListFeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (UnimplementedRouteCoServer) RecordRoute(RouteCo_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (UnimplementedRouteCoServer) RouteChat(RouteCo_RouteChatServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
}
func (UnimplementedRouteCoServer) mustEmbedUnimplementedRouteCoServer() {}

// UnsafeRouteCoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteCoServer will
// result in compilation errors.
type UnsafeRouteCoServer interface {
	mustEmbedUnimplementedRouteCoServer()
}

func RegisterRouteCoServer(s grpc.ServiceRegistrar, srv RouteCoServer) {
	s.RegisterService(&RouteCo_ServiceDesc, srv)
}

func _RouteCo_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteCoServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeco.RouteCo/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteCoServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteCo_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteCoServer).ListFeatures(m, &routeCoListFeaturesServer{stream})
}

type RouteCo_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeCoListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeCoListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteCo_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteCoServer).RecordRoute(&routeCoRecordRouteServer{stream})
}

type RouteCo_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeCoRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeCoRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeCoRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteCo_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteCoServer).RouteChat(&routeCoRouteChatServer{stream})
}

type RouteCo_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeCoRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeCoRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeCoRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteCo_ServiceDesc is the grpc.ServiceDesc for RouteCo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteCo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeco.RouteCo",
	HandlerType: (*RouteCoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteCo_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteCo_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteCo_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteCo_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "routeCo/routeCo.proto",
}
