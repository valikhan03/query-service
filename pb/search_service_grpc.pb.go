// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: search_service.proto

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	GetAuction(ctx context.Context, in *GetAuctionInfoRequest, opts ...grpc.CallOption) (*GetAuctionInfoResponse, error)
	SearchAuctions(ctx context.Context, in *SearchAuctionsRequest, opts ...grpc.CallOption) (*SearchAuctionsResponse, error)
	GetProduct(ctx context.Context, in *GetProductInfoRequest, opts ...grpc.CallOption) (*GetProductInfoResponse, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) GetAuction(ctx context.Context, in *GetAuctionInfoRequest, opts ...grpc.CallOption) (*GetAuctionInfoResponse, error) {
	out := new(GetAuctionInfoResponse)
	err := c.cc.Invoke(ctx, "/protobuf.SearchService/GetAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchAuctions(ctx context.Context, in *SearchAuctionsRequest, opts ...grpc.CallOption) (*SearchAuctionsResponse, error) {
	out := new(SearchAuctionsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.SearchService/SearchAuctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetProduct(ctx context.Context, in *GetProductInfoRequest, opts ...grpc.CallOption) (*GetProductInfoResponse, error) {
	out := new(GetProductInfoResponse)
	err := c.cc.Invoke(ctx, "/protobuf.SearchService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.SearchService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	GetAuction(context.Context, *GetAuctionInfoRequest) (*GetAuctionInfoResponse, error)
	SearchAuctions(context.Context, *SearchAuctionsRequest) (*SearchAuctionsResponse, error)
	GetProduct(context.Context, *GetProductInfoRequest) (*GetProductInfoResponse, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
	//mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) GetAuction(context.Context, *GetAuctionInfoRequest) (*GetAuctionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuction not implemented")
}
func (UnimplementedSearchServiceServer) SearchAuctions(context.Context, *SearchAuctionsRequest) (*SearchAuctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAuctions not implemented")
}
func (UnimplementedSearchServiceServer) GetProduct(context.Context, *GetProductInfoRequest) (*GetProductInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedSearchServiceServer) SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_GetAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuctionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SearchService/GetAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetAuction(ctx, req.(*GetAuctionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchAuctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAuctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchAuctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SearchService/SearchAuctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchAuctions(ctx, req.(*SearchAuctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SearchService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetProduct(ctx, req.(*GetProductInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SearchService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchProducts(ctx, req.(*SearchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuction",
			Handler:    _SearchService_GetAuction_Handler,
		},
		{
			MethodName: "SearchAuctions",
			Handler:    _SearchService_SearchAuctions_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _SearchService_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _SearchService_SearchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search_service.proto",
}
