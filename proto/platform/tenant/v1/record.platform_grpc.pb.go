// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: record.platform.proto

package v1

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
	Records_List_FullMethodName = "/chainguard.platform.tenant.Records/List"
)

// RecordsClient is the client API for Records service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordsClient interface {
	List(ctx context.Context, in *RecordFilter, opts ...grpc.CallOption) (*RecordList, error)
}

type recordsClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordsClient(cc grpc.ClientConnInterface) RecordsClient {
	return &recordsClient{cc}
}

func (c *recordsClient) List(ctx context.Context, in *RecordFilter, opts ...grpc.CallOption) (*RecordList, error) {
	out := new(RecordList)
	err := c.cc.Invoke(ctx, Records_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordsServer is the server API for Records service.
// All implementations must embed UnimplementedRecordsServer
// for forward compatibility
type RecordsServer interface {
	List(context.Context, *RecordFilter) (*RecordList, error)
	mustEmbedUnimplementedRecordsServer()
}

// UnimplementedRecordsServer must be embedded to have forward compatible implementations.
type UnimplementedRecordsServer struct {
}

func (UnimplementedRecordsServer) List(context.Context, *RecordFilter) (*RecordList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRecordsServer) mustEmbedUnimplementedRecordsServer() {}

// UnsafeRecordsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordsServer will
// result in compilation errors.
type UnsafeRecordsServer interface {
	mustEmbedUnimplementedRecordsServer()
}

func RegisterRecordsServer(s grpc.ServiceRegistrar, srv RecordsServer) {
	s.RegisterService(&Records_ServiceDesc, srv)
}

func _Records_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Records_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordsServer).List(ctx, req.(*RecordFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// Records_ServiceDesc is the grpc.ServiceDesc for Records service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Records_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.tenant.Records",
	HandlerType: (*RecordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Records_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "record.platform.proto",
}