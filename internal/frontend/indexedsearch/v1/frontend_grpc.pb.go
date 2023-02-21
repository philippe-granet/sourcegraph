// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: frontend.proto

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

// IndexedSearchConfigurationServiceClient is the client API for IndexedSearchConfigurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexedSearchConfigurationServiceClient interface {
	// SearchConfiguration returns the current indexing configuration for the specified repositories.
	SearchConfiguration(ctx context.Context, in *SearchConfigurationRequest, opts ...grpc.CallOption) (*SearchConfigurationResponse, error)
	// List returns the list of repositories that the client should index.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// RepositoryRank returns the rank vector for the specified repository.
	RepositoryRank(ctx context.Context, in *RepositoryRankRequest, opts ...grpc.CallOption) (*RepositoryRankResponse, error)
	// DocumentRanks returns the rank vectors for all documents in the specified repository.
	DocumentRanks(ctx context.Context, in *DocumentRanksRequest, opts ...grpc.CallOption) (*DocumentRanksResponse, error)
	// UpdateIndexStatus informs the server that the caller has indexed the specified repositories
	// at the specified commits.
	UpdateIndexStatus(ctx context.Context, in *UpdateIndexStatusRequest, opts ...grpc.CallOption) (*UpdateIndexStatusResponse, error)
}

type indexedSearchConfigurationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexedSearchConfigurationServiceClient(cc grpc.ClientConnInterface) IndexedSearchConfigurationServiceClient {
	return &indexedSearchConfigurationServiceClient{cc}
}

func (c *indexedSearchConfigurationServiceClient) SearchConfiguration(ctx context.Context, in *SearchConfigurationRequest, opts ...grpc.CallOption) (*SearchConfigurationResponse, error) {
	out := new(SearchConfigurationResponse)
	err := c.cc.Invoke(ctx, "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/SearchConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexedSearchConfigurationServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexedSearchConfigurationServiceClient) RepositoryRank(ctx context.Context, in *RepositoryRankRequest, opts ...grpc.CallOption) (*RepositoryRankResponse, error) {
	out := new(RepositoryRankResponse)
	err := c.cc.Invoke(ctx, "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/RepositoryRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexedSearchConfigurationServiceClient) DocumentRanks(ctx context.Context, in *DocumentRanksRequest, opts ...grpc.CallOption) (*DocumentRanksResponse, error) {
	out := new(DocumentRanksResponse)
	err := c.cc.Invoke(ctx, "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/DocumentRanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexedSearchConfigurationServiceClient) UpdateIndexStatus(ctx context.Context, in *UpdateIndexStatusRequest, opts ...grpc.CallOption) (*UpdateIndexStatusResponse, error) {
	out := new(UpdateIndexStatusResponse)
	err := c.cc.Invoke(ctx, "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/UpdateIndexStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexedSearchConfigurationServiceServer is the server API for IndexedSearchConfigurationService service.
// All implementations must embed UnimplementedIndexedSearchConfigurationServiceServer
// for forward compatibility
type IndexedSearchConfigurationServiceServer interface {
	// SearchConfiguration returns the current indexing configuration for the specified repositories.
	SearchConfiguration(context.Context, *SearchConfigurationRequest) (*SearchConfigurationResponse, error)
	// List returns the list of repositories that the client should index.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// RepositoryRank returns the rank vector for the specified repository.
	RepositoryRank(context.Context, *RepositoryRankRequest) (*RepositoryRankResponse, error)
	// DocumentRanks returns the rank vectors for all documents in the specified repository.
	DocumentRanks(context.Context, *DocumentRanksRequest) (*DocumentRanksResponse, error)
	// UpdateIndexStatus informs the server that the caller has indexed the specified repositories
	// at the specified commits.
	UpdateIndexStatus(context.Context, *UpdateIndexStatusRequest) (*UpdateIndexStatusResponse, error)
	mustEmbedUnimplementedIndexedSearchConfigurationServiceServer()
}

// UnimplementedIndexedSearchConfigurationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIndexedSearchConfigurationServiceServer struct {
}

func (UnimplementedIndexedSearchConfigurationServiceServer) SearchConfiguration(context.Context, *SearchConfigurationRequest) (*SearchConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchConfiguration not implemented")
}
func (UnimplementedIndexedSearchConfigurationServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedIndexedSearchConfigurationServiceServer) RepositoryRank(context.Context, *RepositoryRankRequest) (*RepositoryRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryRank not implemented")
}
func (UnimplementedIndexedSearchConfigurationServiceServer) DocumentRanks(context.Context, *DocumentRanksRequest) (*DocumentRanksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocumentRanks not implemented")
}
func (UnimplementedIndexedSearchConfigurationServiceServer) UpdateIndexStatus(context.Context, *UpdateIndexStatusRequest) (*UpdateIndexStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIndexStatus not implemented")
}
func (UnimplementedIndexedSearchConfigurationServiceServer) mustEmbedUnimplementedIndexedSearchConfigurationServiceServer() {
}

// UnsafeIndexedSearchConfigurationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexedSearchConfigurationServiceServer will
// result in compilation errors.
type UnsafeIndexedSearchConfigurationServiceServer interface {
	mustEmbedUnimplementedIndexedSearchConfigurationServiceServer()
}

func RegisterIndexedSearchConfigurationServiceServer(s grpc.ServiceRegistrar, srv IndexedSearchConfigurationServiceServer) {
	s.RegisterService(&IndexedSearchConfigurationService_ServiceDesc, srv)
}

func _IndexedSearchConfigurationService_SearchConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexedSearchConfigurationServiceServer).SearchConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/SearchConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexedSearchConfigurationServiceServer).SearchConfiguration(ctx, req.(*SearchConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexedSearchConfigurationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexedSearchConfigurationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexedSearchConfigurationServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexedSearchConfigurationService_RepositoryRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexedSearchConfigurationServiceServer).RepositoryRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/RepositoryRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexedSearchConfigurationServiceServer).RepositoryRank(ctx, req.(*RepositoryRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexedSearchConfigurationService_DocumentRanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentRanksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexedSearchConfigurationServiceServer).DocumentRanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/DocumentRanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexedSearchConfigurationServiceServer).DocumentRanks(ctx, req.(*DocumentRanksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexedSearchConfigurationService_UpdateIndexStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIndexStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexedSearchConfigurationServiceServer).UpdateIndexStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.indexedsearch.v1.IndexedSearchConfigurationService/UpdateIndexStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexedSearchConfigurationServiceServer).UpdateIndexStatus(ctx, req.(*UpdateIndexStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IndexedSearchConfigurationService_ServiceDesc is the grpc.ServiceDesc for IndexedSearchConfigurationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndexedSearchConfigurationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "frontend.indexedsearch.v1.IndexedSearchConfigurationService",
	HandlerType: (*IndexedSearchConfigurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchConfiguration",
			Handler:    _IndexedSearchConfigurationService_SearchConfiguration_Handler,
		},
		{
			MethodName: "List",
			Handler:    _IndexedSearchConfigurationService_List_Handler,
		},
		{
			MethodName: "RepositoryRank",
			Handler:    _IndexedSearchConfigurationService_RepositoryRank_Handler,
		},
		{
			MethodName: "DocumentRanks",
			Handler:    _IndexedSearchConfigurationService_DocumentRanks_Handler,
		},
		{
			MethodName: "UpdateIndexStatus",
			Handler:    _IndexedSearchConfigurationService_UpdateIndexStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "frontend.proto",
}
