// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: cronjob/v1/cronjob.proto

package cronjobv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CronJobService_Preempt_FullMethodName       = "/cronjob.v1.CronJobService/Preempt"
	CronJobService_ResetNextTime_FullMethodName = "/cronjob.v1.CronJobService/ResetNextTime"
	CronJobService_AddJob_FullMethodName        = "/cronjob.v1.CronJobService/AddJob"
	CronJobService_StopJob_FullMethodName       = "/cronjob.v1.CronJobService/StopJob"
	CronJobService_Release_FullMethodName       = "/cronjob.v1.CronJobService/Release"
)

// CronJobServiceClient is the client API for CronJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CronJobServiceClient interface {
	Preempt(ctx context.Context, in *PreemptRequest, opts ...grpc.CallOption) (*PreemptResponse, error)
	ResetNextTime(ctx context.Context, in *ResetNextTimeRequest, opts ...grpc.CallOption) (*ResetNextTimeResponse, error)
	AddJob(ctx context.Context, in *AddJobRequest, opts ...grpc.CallOption) (*AddJobResponse, error)
	StopJob(ctx context.Context, in *StopJobRequest, opts ...grpc.CallOption) (*StopJobRequest, error)
	Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
}

type cronJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCronJobServiceClient(cc grpc.ClientConnInterface) CronJobServiceClient {
	return &cronJobServiceClient{cc}
}

func (c *cronJobServiceClient) Preempt(ctx context.Context, in *PreemptRequest, opts ...grpc.CallOption) (*PreemptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PreemptResponse)
	err := c.cc.Invoke(ctx, CronJobService_Preempt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronJobServiceClient) ResetNextTime(ctx context.Context, in *ResetNextTimeRequest, opts ...grpc.CallOption) (*ResetNextTimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetNextTimeResponse)
	err := c.cc.Invoke(ctx, CronJobService_ResetNextTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronJobServiceClient) AddJob(ctx context.Context, in *AddJobRequest, opts ...grpc.CallOption) (*AddJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddJobResponse)
	err := c.cc.Invoke(ctx, CronJobService_AddJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronJobServiceClient) StopJob(ctx context.Context, in *StopJobRequest, opts ...grpc.CallOption) (*StopJobRequest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopJobRequest)
	err := c.cc.Invoke(ctx, CronJobService_StopJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronJobServiceClient) Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, CronJobService_Release_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CronJobServiceServer is the server API for CronJobService service.
// All implementations must embed UnimplementedCronJobServiceServer
// for forward compatibility
type CronJobServiceServer interface {
	Preempt(context.Context, *PreemptRequest) (*PreemptResponse, error)
	ResetNextTime(context.Context, *ResetNextTimeRequest) (*ResetNextTimeResponse, error)
	AddJob(context.Context, *AddJobRequest) (*AddJobResponse, error)
	StopJob(context.Context, *StopJobRequest) (*StopJobRequest, error)
	Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
	mustEmbedUnimplementedCronJobServiceServer()
}

// UnimplementedCronJobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCronJobServiceServer struct {
}

func (UnimplementedCronJobServiceServer) Preempt(context.Context, *PreemptRequest) (*PreemptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Preempt not implemented")
}
func (UnimplementedCronJobServiceServer) ResetNextTime(context.Context, *ResetNextTimeRequest) (*ResetNextTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetNextTime not implemented")
}
func (UnimplementedCronJobServiceServer) AddJob(context.Context, *AddJobRequest) (*AddJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddJob not implemented")
}
func (UnimplementedCronJobServiceServer) StopJob(context.Context, *StopJobRequest) (*StopJobRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopJob not implemented")
}
func (UnimplementedCronJobServiceServer) Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (UnimplementedCronJobServiceServer) mustEmbedUnimplementedCronJobServiceServer() {}

// UnsafeCronJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CronJobServiceServer will
// result in compilation errors.
type UnsafeCronJobServiceServer interface {
	mustEmbedUnimplementedCronJobServiceServer()
}

func RegisterCronJobServiceServer(s grpc.ServiceRegistrar, srv CronJobServiceServer) {
	s.RegisterService(&CronJobService_ServiceDesc, srv)
}

func _CronJobService_Preempt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreemptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServiceServer).Preempt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CronJobService_Preempt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServiceServer).Preempt(ctx, req.(*PreemptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronJobService_ResetNextTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetNextTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServiceServer).ResetNextTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CronJobService_ResetNextTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServiceServer).ResetNextTime(ctx, req.(*ResetNextTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronJobService_AddJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServiceServer).AddJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CronJobService_AddJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServiceServer).AddJob(ctx, req.(*AddJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronJobService_StopJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServiceServer).StopJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CronJobService_StopJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServiceServer).StopJob(ctx, req.(*StopJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronJobService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CronJobService_Release_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServiceServer).Release(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CronJobService_ServiceDesc is the grpc.ServiceDesc for CronJobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CronJobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cronjob.v1.CronJobService",
	HandlerType: (*CronJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Preempt",
			Handler:    _CronJobService_Preempt_Handler,
		},
		{
			MethodName: "ResetNextTime",
			Handler:    _CronJobService_ResetNextTime_Handler,
		},
		{
			MethodName: "AddJob",
			Handler:    _CronJobService_AddJob_Handler,
		},
		{
			MethodName: "StopJob",
			Handler:    _CronJobService_StopJob_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _CronJobService_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cronjob/v1/cronjob.proto",
}
