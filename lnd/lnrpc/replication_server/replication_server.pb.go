// Code generated by protoc-gen-go. DO NOT EDIT.
// source: replication_server/replication_server.proto

package replication_server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetTokenOffersRequest struct {
	IssuerId             string   `protobuf:"bytes,1,opt,name=issuer_id,json=issuerId,proto3" json:"issuer_id,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint64   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenOffersRequest) Reset()         { *m = GetTokenOffersRequest{} }
func (m *GetTokenOffersRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenOffersRequest) ProtoMessage()    {}
func (*GetTokenOffersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_052aa1ffa0d701d2, []int{0}
}

func (m *GetTokenOffersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenOffersRequest.Unmarshal(m, b)
}
func (m *GetTokenOffersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenOffersRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenOffersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenOffersRequest.Merge(m, src)
}
func (m *GetTokenOffersRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenOffersRequest.Size(m)
}
func (m *GetTokenOffersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenOffersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenOffersRequest proto.InternalMessageInfo

func (m *GetTokenOffersRequest) GetIssuerId() string {
	if m != nil {
		return m.IssuerId
	}
	return ""
}

func (m *GetTokenOffersRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetTokenOffersRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type GetTokenOffersResponse struct {
	Offers               []*TokenOffer `protobuf:"bytes,1,rep,name=offers,proto3" json:"offers,omitempty"`
	Total                uint64        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetTokenOffersResponse) Reset()         { *m = GetTokenOffersResponse{} }
func (m *GetTokenOffersResponse) String() string { return proto.CompactTextString(m) }
func (*GetTokenOffersResponse) ProtoMessage()    {}
func (*GetTokenOffersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_052aa1ffa0d701d2, []int{1}
}

func (m *GetTokenOffersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenOffersResponse.Unmarshal(m, b)
}
func (m *GetTokenOffersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenOffersResponse.Marshal(b, m, deterministic)
}
func (m *GetTokenOffersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenOffersResponse.Merge(m, src)
}
func (m *GetTokenOffersResponse) XXX_Size() int {
	return xxx_messageInfo_GetTokenOffersResponse.Size(m)
}
func (m *GetTokenOffersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenOffersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenOffersResponse proto.InternalMessageInfo

func (m *GetTokenOffersResponse) GetOffers() []*TokenOffer {
	if m != nil {
		return m.Offers
	}
	return nil
}

func (m *GetTokenOffersResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type TokenOffer struct {
	IssuerId string `protobuf:"bytes,1,opt,name=issuer_id,json=issuerId,proto3" json:"issuer_id,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// TODO: discuss. Should we assume decimal prices?
	Price                uint64   `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenOffer) Reset()         { *m = TokenOffer{} }
func (m *TokenOffer) String() string { return proto.CompactTextString(m) }
func (*TokenOffer) ProtoMessage()    {}
func (*TokenOffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_052aa1ffa0d701d2, []int{2}
}

func (m *TokenOffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenOffer.Unmarshal(m, b)
}
func (m *TokenOffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenOffer.Marshal(b, m, deterministic)
}
func (m *TokenOffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenOffer.Merge(m, src)
}
func (m *TokenOffer) XXX_Size() int {
	return xxx_messageInfo_TokenOffer.Size(m)
}
func (m *TokenOffer) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenOffer.DiscardUnknown(m)
}

var xxx_messageInfo_TokenOffer proto.InternalMessageInfo

func (m *TokenOffer) GetIssuerId() string {
	if m != nil {
		return m.IssuerId
	}
	return ""
}

func (m *TokenOffer) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TokenOffer) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*GetTokenOffersRequest)(nil), "lnrpc.GetTokenOffersRequest")
	proto.RegisterType((*GetTokenOffersResponse)(nil), "lnrpc.GetTokenOffersResponse")
	proto.RegisterType((*TokenOffer)(nil), "lnrpc.TokenOffer")
}

func init() {
	proto.RegisterFile("replication_server/replication_server.proto", fileDescriptor_052aa1ffa0d701d2)
}

var fileDescriptor_052aa1ffa0d701d2 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x24, 0xd6, 0x16, 0xb3, 0x82, 0xd0, 0x45, 0x4b, 0xf0, 0x03, 0x42, 0x4e, 0x11, 0x31, 0x81,
	0x8a, 0x78, 0xf7, 0x22, 0x1e, 0x44, 0x58, 0xf5, 0xa0, 0x97, 0x92, 0x6c, 0x5e, 0xec, 0xd2, 0x34,
	0xbb, 0xee, 0x7b, 0xf1, 0xf7, 0x4b, 0x36, 0x91, 0xa0, 0x56, 0x7a, 0x08, 0x64, 0x66, 0x67, 0x77,
	0xe6, 0xcd, 0x63, 0x17, 0x16, 0x4c, 0xa5, 0x64, 0x46, 0x4a, 0xd7, 0x0b, 0x04, 0xfb, 0x09, 0x36,
	0xfd, 0x4b, 0x25, 0xc6, 0x6a, 0xd2, 0x7c, 0x5c, 0xd5, 0xd6, 0xc8, 0x28, 0x67, 0x47, 0x77, 0x40,
	0xcf, 0x7a, 0x05, 0xf5, 0x63, 0x59, 0x82, 0x45, 0x01, 0x1f, 0x0d, 0x20, 0xf1, 0x13, 0xe6, 0x2b,
	0xc4, 0x06, 0xec, 0x42, 0x15, 0x81, 0x17, 0x7a, 0xb1, 0x2f, 0xf6, 0x3a, 0xe2, 0xbe, 0xe0, 0x87,
	0x6c, 0x5c, 0xa9, 0xb5, 0xa2, 0x60, 0x27, 0xf4, 0xe2, 0x5d, 0xd1, 0x01, 0x3e, 0x63, 0x13, 0x5d,
	0x96, 0x08, 0x14, 0x8c, 0x1c, 0xdd, 0xa3, 0xe8, 0x95, 0xcd, 0x7e, 0x7b, 0xa0, 0xd1, 0x35, 0x02,
	0x3f, 0x77, 0x37, 0xc0, 0x62, 0xe0, 0x85, 0xa3, 0x78, 0x7f, 0x3e, 0x4d, 0x5c, 0xaa, 0x64, 0xd0,
	0x8a, 0x5e, 0xd0, 0x5a, 0x92, 0xa6, 0xac, 0xfa, 0xb6, 0x74, 0x20, 0x7a, 0x61, 0x6c, 0xd0, 0x6e,
	0xcd, 0x4c, 0xad, 0xd4, 0x3d, 0xe0, 0x8b, 0x0e, 0xb4, 0xac, 0xb1, 0x4a, 0x42, 0x1f, 0xb9, 0x03,
	0xf3, 0x9c, 0x4d, 0xc5, 0x50, 0xdc, 0x93, 0xeb, 0x8d, 0x3f, 0xb0, 0x83, 0x9f, 0x63, 0xf0, 0xd3,
	0x3e, 0xee, 0xc6, 0x06, 0x8f, 0xcf, 0xfe, 0x39, 0xed, 0x66, 0xbf, 0xbd, 0x79, 0xbb, 0x7e, 0x57,
	0xb4, 0x6c, 0xf2, 0x44, 0xea, 0x75, 0x6a, 0x56, 0x74, 0x29, 0x33, 0x5c, 0xb6, 0x3f, 0x45, 0x5a,
	0xd5, 0xed, 0x67, 0x8d, 0xdc, 0xb0, 0xbf, 0x7c, 0xe2, 0x16, 0x78, 0xf5, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x5d, 0x07, 0x8e, 0x5b, 0xef, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReplicationServerClient is the client API for ReplicationServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReplicationServerClient interface {
	// Returns available token offers
	GetTokenOffers(ctx context.Context, in *GetTokenOffersRequest, opts ...grpc.CallOption) (*GetTokenOffersResponse, error)
}

type replicationServerClient struct {
	cc *grpc.ClientConn
}

func NewReplicationServerClient(cc *grpc.ClientConn) ReplicationServerClient {
	return &replicationServerClient{cc}
}

func (c *replicationServerClient) GetTokenOffers(ctx context.Context, in *GetTokenOffersRequest, opts ...grpc.CallOption) (*GetTokenOffersResponse, error) {
	out := new(GetTokenOffersResponse)
	err := c.cc.Invoke(ctx, "/lnrpc.ReplicationServer/GetTokenOffers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServerServer is the server API for ReplicationServer service.
type ReplicationServerServer interface {
	// Returns available token offers
	GetTokenOffers(context.Context, *GetTokenOffersRequest) (*GetTokenOffersResponse, error)
}

// UnimplementedReplicationServerServer can be embedded to have forward compatible implementations.
type UnimplementedReplicationServerServer struct {
}

func (*UnimplementedReplicationServerServer) GetTokenOffers(ctx context.Context, req *GetTokenOffersRequest) (*GetTokenOffersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenOffers not implemented")
}

func RegisterReplicationServerServer(s *grpc.Server, srv ReplicationServerServer) {
	s.RegisterService(&_ReplicationServer_serviceDesc, srv)
}

func _ReplicationServer_GetTokenOffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenOffersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServerServer).GetTokenOffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lnrpc.ReplicationServer/GetTokenOffers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServerServer).GetTokenOffers(ctx, req.(*GetTokenOffersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReplicationServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lnrpc.ReplicationServer",
	HandlerType: (*ReplicationServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokenOffers",
			Handler:    _ReplicationServer_GetTokenOffers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replication_server/replication_server.proto",
}
