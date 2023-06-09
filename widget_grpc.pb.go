// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: widget.proto

package puzzlewidgetservice

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

// WidgetClient is the client API for Widget service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WidgetClient interface {
	GetWidget(ctx context.Context, in *WidgetRequest, opts ...grpc.CallOption) (*WidgetResponse, error)
	Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
}

type widgetClient struct {
	cc grpc.ClientConnInterface
}

func NewWidgetClient(cc grpc.ClientConnInterface) WidgetClient {
	return &widgetClient{cc}
}

func (c *widgetClient) GetWidget(ctx context.Context, in *WidgetRequest, opts ...grpc.CallOption) (*WidgetResponse, error) {
	out := new(WidgetResponse)
	err := c.cc.Invoke(ctx, "/puzzlewidgetservice.Widget/GetWidget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *widgetClient) Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/puzzlewidgetservice.Widget/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WidgetServer is the server API for Widget service.
// All implementations must embed UnimplementedWidgetServer
// for forward compatibility
type WidgetServer interface {
	GetWidget(context.Context, *WidgetRequest) (*WidgetResponse, error)
	Process(context.Context, *ProcessRequest) (*ProcessResponse, error)
	mustEmbedUnimplementedWidgetServer()
}

// UnimplementedWidgetServer must be embedded to have forward compatible implementations.
type UnimplementedWidgetServer struct {
}

func (UnimplementedWidgetServer) GetWidget(context.Context, *WidgetRequest) (*WidgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWidget not implemented")
}
func (UnimplementedWidgetServer) Process(context.Context, *ProcessRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (UnimplementedWidgetServer) mustEmbedUnimplementedWidgetServer() {}

// UnsafeWidgetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WidgetServer will
// result in compilation errors.
type UnsafeWidgetServer interface {
	mustEmbedUnimplementedWidgetServer()
}

func RegisterWidgetServer(s grpc.ServiceRegistrar, srv WidgetServer) {
	s.RegisterService(&Widget_ServiceDesc, srv)
}

func _Widget_GetWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WidgetServer).GetWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzlewidgetservice.Widget/GetWidget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WidgetServer).GetWidget(ctx, req.(*WidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Widget_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WidgetServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/puzzlewidgetservice.Widget/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WidgetServer).Process(ctx, req.(*ProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Widget_ServiceDesc is the grpc.ServiceDesc for Widget service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Widget_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "puzzlewidgetservice.Widget",
	HandlerType: (*WidgetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWidget",
			Handler:    _Widget_GetWidget_Handler,
		},
		{
			MethodName: "Process",
			Handler:    _Widget_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "widget.proto",
}
