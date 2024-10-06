// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/ticket.proto

package proto

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
	TicketService_SubmitPurchase_FullMethodName     = "/ticket.TicketService/SubmitPurchase"
	TicketService_ShowReceipt_FullMethodName        = "/ticket.TicketService/ShowReceipt"
	TicketService_ModifySeat_FullMethodName         = "/ticket.TicketService/ModifySeat"
	TicketService_RemoveUser_FullMethodName         = "/ticket.TicketService/RemoveUser"
	TicketService_ViewUsersBySection_FullMethodName = "/ticket.TicketService/ViewUsersBySection"
)

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// TicketService service definition
type TicketServiceClient interface {
	SubmitPurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Receipt, error)
	ShowReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserList, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*Response, error)
	RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error)
	ViewUsersBySection(ctx context.Context, in *ViewUsersBySectionRequest, opts ...grpc.CallOption) (*ViewUsersBySectionResponse, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) SubmitPurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Receipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Receipt)
	err := c.cc.Invoke(ctx, TicketService_SubmitPurchase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ShowReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserList)
	err := c.cc.Invoke(ctx, TicketService_ShowReceipt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, TicketService_ModifySeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, TicketService_RemoveUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ViewUsersBySection(ctx context.Context, in *ViewUsersBySectionRequest, opts ...grpc.CallOption) (*ViewUsersBySectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ViewUsersBySectionResponse)
	err := c.cc.Invoke(ctx, TicketService_ViewUsersBySection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility.
//
// TicketService service definition
type TicketServiceServer interface {
	SubmitPurchase(context.Context, *PurchaseRequest) (*Receipt, error)
	ShowReceipt(context.Context, *UserRequest) (*UserList, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*Response, error)
	RemoveUser(context.Context, *UserRequest) (*Response, error)
	ViewUsersBySection(context.Context, *ViewUsersBySectionRequest) (*ViewUsersBySectionResponse, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTicketServiceServer struct{}

func (UnimplementedTicketServiceServer) SubmitPurchase(context.Context, *PurchaseRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPurchase not implemented")
}
func (UnimplementedTicketServiceServer) ShowReceipt(context.Context, *UserRequest) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowReceipt not implemented")
}
func (UnimplementedTicketServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTicketServiceServer) RemoveUser(context.Context, *UserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketServiceServer) ViewUsersBySection(context.Context, *ViewUsersBySectionRequest) (*ViewUsersBySectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewUsersBySection not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}
func (UnimplementedTicketServiceServer) testEmbeddedByValue()                       {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	// If the following call pancis, it indicates UnimplementedTicketServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_SubmitPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).SubmitPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_SubmitPurchase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).SubmitPurchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ShowReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ShowReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ShowReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ShowReceipt(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ModifySeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).RemoveUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ViewUsersBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewUsersBySectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ViewUsersBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ViewUsersBySection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ViewUsersBySection(ctx, req.(*ViewUsersBySectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticket.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitPurchase",
			Handler:    _TicketService_SubmitPurchase_Handler,
		},
		{
			MethodName: "ShowReceipt",
			Handler:    _TicketService_ShowReceipt_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TicketService_ModifySeat_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketService_RemoveUser_Handler,
		},
		{
			MethodName: "ViewUsersBySection",
			Handler:    _TicketService_ViewUsersBySection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ticket.proto",
}