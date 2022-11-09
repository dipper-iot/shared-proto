// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: services/device/property_device.proto

package device

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PropertyDeviceServiceClient is the client API for PropertyDeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PropertyDeviceServiceClient interface {
	CountByDeviceId(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*IntResponse, error)
	GetByDeviceId(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*ListPropertyDeviceResponse, error)
	Create(ctx context.Context, in *PropertyDevice, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *PropertyDevice, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type propertyDeviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPropertyDeviceServiceClient(cc grpc.ClientConnInterface) PropertyDeviceServiceClient {
	return &propertyDeviceServiceClient{cc}
}

func (c *propertyDeviceServiceClient) CountByDeviceId(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*IntResponse, error) {
	out := new(IntResponse)
	err := c.cc.Invoke(ctx, "/device.PropertyDeviceService/CountByDeviceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyDeviceServiceClient) GetByDeviceId(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*ListPropertyDeviceResponse, error) {
	out := new(ListPropertyDeviceResponse)
	err := c.cc.Invoke(ctx, "/device.PropertyDeviceService/GetByDeviceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyDeviceServiceClient) Create(ctx context.Context, in *PropertyDevice, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/device.PropertyDeviceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyDeviceServiceClient) Update(ctx context.Context, in *PropertyDevice, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/device.PropertyDeviceService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyDeviceServiceClient) Delete(ctx context.Context, in *DeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/device.PropertyDeviceService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PropertyDeviceServiceServer is the server API for PropertyDeviceService service.
// All implementations must embed UnimplementedPropertyDeviceServiceServer
// for forward compatibility
type PropertyDeviceServiceServer interface {
	CountByDeviceId(context.Context, *DeviceRequest) (*IntResponse, error)
	GetByDeviceId(context.Context, *DeviceRequest) (*ListPropertyDeviceResponse, error)
	Create(context.Context, *PropertyDevice) (*emptypb.Empty, error)
	Update(context.Context, *PropertyDevice) (*emptypb.Empty, error)
	Delete(context.Context, *DeviceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPropertyDeviceServiceServer()
}

// UnimplementedPropertyDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPropertyDeviceServiceServer struct {
}

func (UnimplementedPropertyDeviceServiceServer) CountByDeviceId(context.Context, *DeviceRequest) (*IntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountByDeviceId not implemented")
}
func (UnimplementedPropertyDeviceServiceServer) GetByDeviceId(context.Context, *DeviceRequest) (*ListPropertyDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByDeviceId not implemented")
}
func (UnimplementedPropertyDeviceServiceServer) Create(context.Context, *PropertyDevice) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPropertyDeviceServiceServer) Update(context.Context, *PropertyDevice) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPropertyDeviceServiceServer) Delete(context.Context, *DeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPropertyDeviceServiceServer) mustEmbedUnimplementedPropertyDeviceServiceServer() {}

// UnsafePropertyDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PropertyDeviceServiceServer will
// result in compilation errors.
type UnsafePropertyDeviceServiceServer interface {
	mustEmbedUnimplementedPropertyDeviceServiceServer()
}

func RegisterPropertyDeviceServiceServer(s grpc.ServiceRegistrar, srv PropertyDeviceServiceServer) {
	s.RegisterService(&PropertyDeviceService_ServiceDesc, srv)
}

func _PropertyDeviceService_CountByDeviceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyDeviceServiceServer).CountByDeviceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.PropertyDeviceService/CountByDeviceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyDeviceServiceServer).CountByDeviceId(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyDeviceService_GetByDeviceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyDeviceServiceServer).GetByDeviceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.PropertyDeviceService/GetByDeviceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyDeviceServiceServer).GetByDeviceId(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyDeviceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertyDevice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyDeviceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.PropertyDeviceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyDeviceServiceServer).Create(ctx, req.(*PropertyDevice))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyDeviceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertyDevice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyDeviceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.PropertyDeviceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyDeviceServiceServer).Update(ctx, req.(*PropertyDevice))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyDeviceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyDeviceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.PropertyDeviceService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyDeviceServiceServer).Delete(ctx, req.(*DeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PropertyDeviceService_ServiceDesc is the grpc.ServiceDesc for PropertyDeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PropertyDeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "device.PropertyDeviceService",
	HandlerType: (*PropertyDeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountByDeviceId",
			Handler:    _PropertyDeviceService_CountByDeviceId_Handler,
		},
		{
			MethodName: "GetByDeviceId",
			Handler:    _PropertyDeviceService_GetByDeviceId_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PropertyDeviceService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PropertyDeviceService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PropertyDeviceService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/device/property_device.proto",
}