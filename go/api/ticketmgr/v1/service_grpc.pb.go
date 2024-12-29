// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: ticketmgr/v1/service.proto

package ticketmgrv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TicketMgrService_CreateTicket_FullMethodName = "/ticketmgr.v1.TicketMgrService/CreateTicket"
	TicketMgrService_UpdateTicket_FullMethodName = "/ticketmgr.v1.TicketMgrService/UpdateTicket"
	TicketMgrService_DeleteTicket_FullMethodName = "/ticketmgr.v1.TicketMgrService/DeleteTicket"
)

// TicketMgrServiceClient is the client API for TicketMgrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketMgrServiceClient interface {
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	UpdateTicket(ctx context.Context, in *UpdateTicketRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error)
	DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error)
}

type ticketMgrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketMgrServiceClient(cc grpc.ClientConnInterface) TicketMgrServiceClient {
	return &ticketMgrServiceClient{cc}
}

func (c *ticketMgrServiceClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, TicketMgrService_CreateTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketMgrServiceClient) UpdateTicket(ctx context.Context, in *UpdateTicketRequest, opts ...grpc.CallOption) (*UpdateTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTicketResponse)
	err := c.cc.Invoke(ctx, TicketMgrService_UpdateTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketMgrServiceClient) DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTicketResponse)
	err := c.cc.Invoke(ctx, TicketMgrService_DeleteTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketMgrServiceServer is the server API for TicketMgrService service.
// All implementations must embed UnimplementedTicketMgrServiceServer
// for forward compatibility.
type TicketMgrServiceServer interface {
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	UpdateTicket(context.Context, *UpdateTicketRequest) (*UpdateTicketResponse, error)
	DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error)
	mustEmbedUnimplementedTicketMgrServiceServer()
}

// UnimplementedTicketMgrServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTicketMgrServiceServer struct{}

func (UnimplementedTicketMgrServiceServer) CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedTicketMgrServiceServer) UpdateTicket(context.Context, *UpdateTicketRequest) (*UpdateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicket not implemented")
}
func (UnimplementedTicketMgrServiceServer) DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (UnimplementedTicketMgrServiceServer) mustEmbedUnimplementedTicketMgrServiceServer() {}
func (UnimplementedTicketMgrServiceServer) testEmbeddedByValue()                          {}

// UnsafeTicketMgrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketMgrServiceServer will
// result in compilation errors.
type UnsafeTicketMgrServiceServer interface {
	mustEmbedUnimplementedTicketMgrServiceServer()
}

func RegisterTicketMgrServiceServer(s grpc.ServiceRegistrar, srv TicketMgrServiceServer) {
	// If the following call pancis, it indicates UnimplementedTicketMgrServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TicketMgrService_ServiceDesc, srv)
}

func _TicketMgrService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketMgrServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketMgrService_CreateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketMgrServiceServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketMgrService_UpdateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketMgrServiceServer).UpdateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketMgrService_UpdateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketMgrServiceServer).UpdateTicket(ctx, req.(*UpdateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketMgrService_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketMgrServiceServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketMgrService_DeleteTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketMgrServiceServer).DeleteTicket(ctx, req.(*DeleteTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketMgrService_ServiceDesc is the grpc.ServiceDesc for TicketMgrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketMgrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticketmgr.v1.TicketMgrService",
	HandlerType: (*TicketMgrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _TicketMgrService_CreateTicket_Handler,
		},
		{
			MethodName: "UpdateTicket",
			Handler:    _TicketMgrService_UpdateTicket_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _TicketMgrService_DeleteTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticketmgr/v1/service.proto",
}