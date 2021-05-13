// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: treasury/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryCurrenciesRequest is request type for the Query/Currencies RPC method
type QueryCurrenciesRequest struct {
	// pagination defines an optional pagination for the request
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCurrenciesRequest) Reset()         { *m = QueryCurrenciesRequest{} }
func (m *QueryCurrenciesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrenciesRequest) ProtoMessage()    {}
func (*QueryCurrenciesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8b5f9157b6b762, []int{0}
}
func (m *QueryCurrenciesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrenciesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrenciesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrenciesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrenciesRequest.Merge(m, src)
}
func (m *QueryCurrenciesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrenciesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrenciesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrenciesRequest proto.InternalMessageInfo

func (m *QueryCurrenciesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCurrenciesResponse is response type for the Query/Currencies RPC method
type QueryCurrenciesResponse struct {
	// Currencies defines all supported currency denom list
	Currencies *Currencies `protobuf:"bytes,1,opt,name=currencies,proto3" json:"currencies,omitempty"`
	// pagination defines the pagination in the response
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCurrenciesResponse) Reset()         { *m = QueryCurrenciesResponse{} }
func (m *QueryCurrenciesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrenciesResponse) ProtoMessage()    {}
func (*QueryCurrenciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8b5f9157b6b762, []int{1}
}
func (m *QueryCurrenciesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrenciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrenciesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrenciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrenciesResponse.Merge(m, src)
}
func (m *QueryCurrenciesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrenciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrenciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrenciesResponse proto.InternalMessageInfo

func (m *QueryCurrenciesResponse) GetCurrencies() *Currencies {
	if m != nil {
		return m.Currencies
	}
	return nil
}

func (m *QueryCurrenciesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCurrencyRequest is request type for the Query/Currency RPC method
type QueryCurrencyRequest struct {
	// denom defines the denom to query for
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryCurrencyRequest) Reset()         { *m = QueryCurrencyRequest{} }
func (m *QueryCurrencyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrencyRequest) ProtoMessage()    {}
func (*QueryCurrencyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8b5f9157b6b762, []int{2}
}
func (m *QueryCurrencyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrencyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrencyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrencyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrencyRequest.Merge(m, src)
}
func (m *QueryCurrencyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrencyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrencyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrencyRequest proto.InternalMessageInfo

func (m *QueryCurrencyRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// QueryCurrencyResponse is response type for the Query/Currency RPC method
type QueryCurrencyResponse struct {
	// Currency defines a currency info
	Currency *Currency `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (m *QueryCurrencyResponse) Reset()         { *m = QueryCurrencyResponse{} }
func (m *QueryCurrencyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrencyResponse) ProtoMessage()    {}
func (*QueryCurrencyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8b5f9157b6b762, []int{3}
}
func (m *QueryCurrencyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrencyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrencyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrencyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrencyResponse.Merge(m, src)
}
func (m *QueryCurrencyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrencyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrencyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrencyResponse proto.InternalMessageInfo

func (m *QueryCurrencyResponse) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

// QueryParamsRequest is request type for the Query/Params RPC method
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8b5f9157b6b762, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method
type QueryParamsResponse struct {
	// params defines the parameters of treasury module
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8b5f9157b6b762, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryCurrenciesRequest)(nil), "rizonworld.rizon.treasury.QueryCurrenciesRequest")
	proto.RegisterType((*QueryCurrenciesResponse)(nil), "rizonworld.rizon.treasury.QueryCurrenciesResponse")
	proto.RegisterType((*QueryCurrencyRequest)(nil), "rizonworld.rizon.treasury.QueryCurrencyRequest")
	proto.RegisterType((*QueryCurrencyResponse)(nil), "rizonworld.rizon.treasury.QueryCurrencyResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "rizonworld.rizon.treasury.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "rizonworld.rizon.treasury.QueryParamsResponse")
}

func init() { proto.RegisterFile("treasury/query.proto", fileDescriptor_7a8b5f9157b6b762) }

var fileDescriptor_7a8b5f9157b6b762 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xeb, 0xc1, 0xaa, 0xf1, 0x72, 0x33, 0xd9, 0x1f, 0xa2, 0x29, 0x40, 0xf8, 0xab, 0x89,
	0xda, 0xb4, 0x7c, 0x00, 0xa4, 0x21, 0xd8, 0x75, 0xe4, 0x80, 0x10, 0x27, 0xdc, 0xce, 0x0a, 0x91,
	0xd6, 0x38, 0xb3, 0x13, 0x20, 0x20, 0x2e, 0x5c, 0xb8, 0x22, 0x71, 0x40, 0x7c, 0x04, 0x3e, 0x09,
	0x3b, 0x4e, 0xe2, 0xc2, 0x09, 0xa1, 0x96, 0x0f, 0x82, 0x6a, 0xbf, 0x49, 0xd3, 0xc1, 0xd6, 0xee,
	0xe6, 0xd8, 0x7e, 0x9e, 0xe7, 0xf7, 0xbe, 0xaf, 0x15, 0xf0, 0x72, 0x2d, 0x85, 0x29, 0x74, 0xc9,
	0x0f, 0x0a, 0xa9, 0x4b, 0x96, 0x69, 0x95, 0x2b, 0x7a, 0x59, 0x27, 0x6f, 0x55, 0xfa, 0x5a, 0xe9,
	0xfd, 0x3d, 0x66, 0x97, 0xac, 0xba, 0xe6, 0x6f, 0x0d, 0x94, 0x19, 0x2a, 0xc3, 0xfb, 0xc2, 0x48,
	0xa7, 0xe1, 0xaf, 0xba, 0x7d, 0x99, 0x8b, 0x2e, 0xcf, 0x44, 0x9c, 0xa4, 0x22, 0x4f, 0x54, 0xea,
	0x6c, 0x7c, 0x2f, 0x56, 0xb1, 0xb2, 0x4b, 0x3e, 0x59, 0xe1, 0xee, 0x66, 0xac, 0x54, 0xbc, 0x2f,
	0xb9, 0xc8, 0x12, 0x2e, 0xd2, 0x54, 0xe5, 0x56, 0x62, 0xf0, 0x74, 0xbd, 0x06, 0xaa, 0x16, 0x78,
	0xb0, 0x5a, 0x1f, 0x64, 0x42, 0x8b, 0x21, 0xde, 0x0f, 0x5f, 0xc0, 0xda, 0x93, 0x09, 0xc5, 0xc3,
	0x42, 0x6b, 0x99, 0x0e, 0x12, 0x69, 0x22, 0x79, 0x50, 0x48, 0x93, 0xd3, 0xc7, 0x00, 0x53, 0xa2,
	0x0d, 0x72, 0x95, 0xdc, 0xb9, 0xd8, 0xbb, 0xc5, 0x1c, 0x3e, 0x9b, 0xe0, 0x33, 0x57, 0x32, 0xe2,
	0xb3, 0x5d, 0x11, 0x4b, 0xd4, 0x46, 0x0d, 0x65, 0xf8, 0x8d, 0xc0, 0xfa, 0x3f, 0x11, 0x26, 0x53,
	0xa9, 0x91, 0xf4, 0x11, 0xc0, 0xa0, 0xde, 0xc5, 0x8c, 0x9b, 0xec, 0xc4, 0xee, 0xb1, 0x86, 0x45,
	0x43, 0x48, 0x77, 0x66, 0x50, 0x97, 0xac, 0xcd, 0xed, 0xb9, 0xa8, 0x8e, 0x61, 0x86, 0xf5, 0x2e,
	0x78, 0x4d, 0xd4, 0xb2, 0xea, 0x85, 0x07, 0xcb, 0x7b, 0x32, 0x55, 0x43, 0x8b, 0x78, 0x21, 0x72,
	0x1f, 0xe1, 0x33, 0x58, 0x3d, 0x76, 0x1b, 0xcb, 0x7a, 0x00, 0x2b, 0x48, 0x57, 0x62, 0x51, 0xd7,
	0xe7, 0x17, 0x55, 0x46, 0xb5, 0x28, 0xf4, 0x80, 0x5a, 0xe7, 0x5d, 0x3b, 0x2a, 0xa4, 0x08, 0x9f,
	0xc2, 0xa5, 0x99, 0xdd, 0x3a, 0xad, 0xed, 0x46, 0x8a, 0x59, 0xd7, 0x4e, 0xc9, 0x72, 0xd2, 0xed,
	0xf3, 0x87, 0xbf, 0xae, 0xb4, 0x22, 0x94, 0xf5, 0xbe, 0x9f, 0x83, 0x65, 0x6b, 0x4c, 0xbf, 0x10,
	0x80, 0x69, 0x8f, 0x69, 0xf7, 0x14, 0xa7, 0xff, 0xbf, 0x1a, 0xbf, 0x77, 0x16, 0x89, 0x2b, 0x20,
	0x0c, 0x3f, 0xfc, 0xf8, 0xf3, 0x79, 0x69, 0x93, 0xfa, 0xdc, 0x0a, 0xea, 0x97, 0xcb, 0x1b, 0x23,
	0xfe, 0x4a, 0x60, 0xa5, 0x6a, 0x14, 0xe5, 0x0b, 0x86, 0x54, 0xf3, 0xf3, 0xef, 0x2d, 0x2e, 0x40,
	0xa6, 0x2d, 0xcb, 0x74, 0x83, 0x86, 0x27, 0x33, 0xf1, 0x77, 0xf6, 0x19, 0xbc, 0xa7, 0x1f, 0x09,
	0xb4, 0x5d, 0x63, 0x69, 0x67, 0x5e, 0xd0, 0xcc, 0x44, 0x7d, 0xb6, 0xe8, 0x75, 0xa4, 0x0a, 0x2c,
	0xd5, 0x06, 0x5d, 0x3b, 0x4e, 0xe5, 0x26, 0xb9, 0xbd, 0x73, 0x38, 0x0a, 0xc8, 0xd1, 0x28, 0x20,
	0xbf, 0x47, 0x01, 0xf9, 0x34, 0x0e, 0x5a, 0x47, 0xe3, 0xa0, 0xf5, 0x73, 0x1c, 0xb4, 0x9e, 0x77,
	0xe2, 0x24, 0x7f, 0x59, 0xf4, 0xd9, 0x40, 0x0d, 0x9d, 0xb6, 0x63, 0x43, 0xd1, 0xe7, 0xcd, 0xd4,
	0x29, 0x2f, 0x33, 0x69, 0xfa, 0x6d, 0xfb, 0x77, 0xb8, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x10,
	0xae, 0x76, 0xd3, 0xe0, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Currencies queries all supported currency denom list
	Currencies(ctx context.Context, in *QueryCurrenciesRequest, opts ...grpc.CallOption) (*QueryCurrenciesResponse, error)
	// Currency queries a currency info
	Currency(ctx context.Context, in *QueryCurrencyRequest, opts ...grpc.CallOption) (*QueryCurrencyResponse, error)
	// Params queries parameters of treasury
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Currencies(ctx context.Context, in *QueryCurrenciesRequest, opts ...grpc.CallOption) (*QueryCurrenciesResponse, error) {
	out := new(QueryCurrenciesResponse)
	err := c.cc.Invoke(ctx, "/rizonworld.rizon.treasury.Query/Currencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Currency(ctx context.Context, in *QueryCurrencyRequest, opts ...grpc.CallOption) (*QueryCurrencyResponse, error) {
	out := new(QueryCurrencyResponse)
	err := c.cc.Invoke(ctx, "/rizonworld.rizon.treasury.Query/Currency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/rizonworld.rizon.treasury.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Currencies queries all supported currency denom list
	Currencies(context.Context, *QueryCurrenciesRequest) (*QueryCurrenciesResponse, error)
	// Currency queries a currency info
	Currency(context.Context, *QueryCurrencyRequest) (*QueryCurrencyResponse, error)
	// Params queries parameters of treasury
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Currencies(ctx context.Context, req *QueryCurrenciesRequest) (*QueryCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Currencies not implemented")
}
func (*UnimplementedQueryServer) Currency(ctx context.Context, req *QueryCurrencyRequest) (*QueryCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Currency not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Currencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Currencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rizonworld.rizon.treasury.Query/Currencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Currencies(ctx, req.(*QueryCurrenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Currency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Currency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rizonworld.rizon.treasury.Query/Currency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Currency(ctx, req.(*QueryCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rizonworld.rizon.treasury.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rizonworld.rizon.treasury.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Currencies",
			Handler:    _Query_Currencies_Handler,
		},
		{
			MethodName: "Currency",
			Handler:    _Query_Currency_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "treasury/query.proto",
}

func (m *QueryCurrenciesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrenciesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrenciesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCurrenciesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrenciesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrenciesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Currencies != nil {
		{
			size, err := m.Currencies.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCurrencyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrencyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrencyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCurrencyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrencyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrencyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Currency != nil {
		{
			size, err := m.Currency.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCurrenciesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCurrenciesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Currencies != nil {
		l = m.Currencies.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCurrencyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCurrencyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Currency != nil {
		l = m.Currency.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCurrenciesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCurrenciesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrenciesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCurrenciesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCurrenciesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrenciesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Currencies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Currencies == nil {
				m.Currencies = &Currencies{}
			}
			if err := m.Currencies.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCurrencyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCurrencyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrencyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCurrencyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCurrencyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrencyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Currency", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Currency == nil {
				m.Currency = &Currency{}
			}
			if err := m.Currency.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
