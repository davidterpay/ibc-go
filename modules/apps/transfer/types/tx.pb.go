// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/transfer/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/davidterpay/ibc-go/modules/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// MsgTransfer defines a msg to transfer fungible tokens (i.e Coins) between
// ICS20 enabled chains. See ICS Spec here:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer#data-structures
type MsgTransfer struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// the tokens to be transferred
	Token types.Coin `protobuf:"bytes,3,opt,name=token,proto3" json:"token"`
	// the sender address
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types1.Height `protobuf:"bytes,6,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp in absolute nanoseconds since unix epoch.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,7,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	// optional memo
	Memo string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *MsgTransfer) Reset()         { *m = MsgTransfer{} }
func (m *MsgTransfer) String() string { return proto.CompactTextString(m) }
func (*MsgTransfer) ProtoMessage()    {}
func (*MsgTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_7401ed9bed2f8e09, []int{0}
}
func (m *MsgTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransfer.Merge(m, src)
}
func (m *MsgTransfer) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransfer proto.InternalMessageInfo

// MsgTransferResponse defines the Msg/Transfer response type.
type MsgTransferResponse struct {
	// sequence number of the transfer packet sent
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgTransferResponse) Reset()         { *m = MsgTransferResponse{} }
func (m *MsgTransferResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransferResponse) ProtoMessage()    {}
func (*MsgTransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7401ed9bed2f8e09, []int{1}
}
func (m *MsgTransferResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferResponse.Merge(m, src)
}
func (m *MsgTransferResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferResponse proto.InternalMessageInfo

func (m *MsgTransferResponse) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// MsgRegisterChain defines a msg to register a chain with the transfer module.
type MsgRegisterChain struct {
	// the chain id of the chain to be registered
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// the port of the chain to be registered
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// the channel of the chain to be registered
	Channel string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	// the signer of tx
	SenderAddress string `protobuf:"bytes,4,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
}

func (m *MsgRegisterChain) Reset()         { *m = MsgRegisterChain{} }
func (m *MsgRegisterChain) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterChain) ProtoMessage()    {}
func (*MsgRegisterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_7401ed9bed2f8e09, []int{2}
}
func (m *MsgRegisterChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterChain.Merge(m, src)
}
func (m *MsgRegisterChain) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterChain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterChain proto.InternalMessageInfo

func (m *MsgRegisterChain) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MsgRegisterChain) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgRegisterChain) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *MsgRegisterChain) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

// MsgRegisterChainResponse defines the Msg/RegisterChain response type.
type MsgRegisterChainResponse struct {
}

func (m *MsgRegisterChainResponse) Reset()         { *m = MsgRegisterChainResponse{} }
func (m *MsgRegisterChainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterChainResponse) ProtoMessage()    {}
func (*MsgRegisterChainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7401ed9bed2f8e09, []int{3}
}
func (m *MsgRegisterChainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterChainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterChainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterChainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterChainResponse.Merge(m, src)
}
func (m *MsgRegisterChainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterChainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterChainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterChainResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTransfer)(nil), "ibc.applications.transfer.v1.MsgTransfer")
	proto.RegisterType((*MsgTransferResponse)(nil), "ibc.applications.transfer.v1.MsgTransferResponse")
	proto.RegisterType((*MsgRegisterChain)(nil), "ibc.applications.transfer.v1.MsgRegisterChain")
	proto.RegisterType((*MsgRegisterChainResponse)(nil), "ibc.applications.transfer.v1.MsgRegisterChainResponse")
}

func init() {
	proto.RegisterFile("ibc/applications/transfer/v1/tx.proto", fileDescriptor_7401ed9bed2f8e09)
}

var fileDescriptor_7401ed9bed2f8e09 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x3d, 0x6f, 0xdb, 0x3c,
	0x10, 0xb6, 0x62, 0x27, 0xf1, 0xcb, 0xc0, 0x41, 0x5e, 0xb6, 0x0d, 0x14, 0x21, 0x95, 0x02, 0x01,
	0x01, 0xd2, 0xa1, 0x14, 0x94, 0xa2, 0x09, 0x90, 0xa9, 0x75, 0x96, 0x66, 0x08, 0xd0, 0x0a, 0x99,
	0xba, 0xb8, 0x12, 0x7d, 0x95, 0x89, 0x5a, 0xa2, 0x4a, 0xd2, 0x6e, 0xb3, 0x76, 0xea, 0xd8, 0x9f,
	0x90, 0x5f, 0xd2, 0x39, 0x63, 0xc6, 0x4e, 0x46, 0x91, 0x2c, 0x99, 0xf3, 0x0b, 0x0a, 0x92, 0xb2,
	0x6b, 0x77, 0xe8, 0xc7, 0xa4, 0xfb, 0x78, 0x4e, 0xcf, 0xdd, 0x3d, 0x24, 0xd1, 0x2e, 0xcb, 0x68,
	0x94, 0x56, 0xd5, 0x90, 0xd1, 0x54, 0x31, 0x5e, 0xca, 0x48, 0x89, 0xb4, 0x94, 0x6f, 0x41, 0x44,
	0xe3, 0x38, 0x52, 0x1f, 0x49, 0x25, 0xb8, 0xe2, 0x78, 0x9b, 0x65, 0x94, 0xcc, 0xc3, 0xc8, 0x14,
	0x46, 0xc6, 0xb1, 0x77, 0x3f, 0xe7, 0x39, 0x37, 0xc0, 0x48, 0x5b, 0xb6, 0xc6, 0xf3, 0x29, 0x97,
	0x05, 0x97, 0x51, 0x96, 0x4a, 0x88, 0xc6, 0x71, 0x06, 0x2a, 0x8d, 0x23, 0xca, 0x59, 0x59, 0xe7,
	0x03, 0x4d, 0x4d, 0xb9, 0x80, 0x88, 0x0e, 0x19, 0x94, 0x4a, 0x13, 0x5a, 0xcb, 0x02, 0xc2, 0xaf,
	0x4d, 0xb4, 0x76, 0x2a, 0xf3, 0xb3, 0x9a, 0x09, 0x1f, 0xa2, 0x35, 0xc9, 0x47, 0x82, 0x42, 0xaf,
	0xe2, 0x42, 0xb9, 0xce, 0x8e, 0xb3, 0xf7, 0x5f, 0x77, 0xf3, 0x6e, 0x12, 0xe0, 0xf3, 0xb4, 0x18,
	0x1e, 0x85, 0x73, 0xc9, 0x30, 0x41, 0xd6, 0x7b, 0xc9, 0x85, 0xc2, 0xcf, 0xd0, 0x7a, 0x9d, 0xa3,
	0x83, 0xb4, 0x2c, 0x61, 0xe8, 0x2e, 0x99, 0xda, 0xad, 0xbb, 0x49, 0xf0, 0x60, 0xa1, 0xb6, 0xce,
	0x87, 0x49, 0xc7, 0x06, 0x8e, 0xad, 0x8f, 0x9f, 0xa2, 0x65, 0xc5, 0xdf, 0x41, 0xe9, 0x36, 0x77,
	0x9c, 0xbd, 0xb5, 0xfd, 0x2d, 0x62, 0x67, 0x23, 0x7a, 0x36, 0x52, 0xcf, 0x46, 0x8e, 0x39, 0x2b,
	0xbb, 0xad, 0xcb, 0x49, 0xd0, 0x48, 0x2c, 0x1a, 0x6f, 0xa2, 0x15, 0x09, 0x65, 0x1f, 0x84, 0xdb,
	0xd2, 0x84, 0x49, 0xed, 0x61, 0x0f, 0xb5, 0x05, 0x50, 0x60, 0x63, 0x10, 0xee, 0xb2, 0xc9, 0xcc,
	0x7c, 0xfc, 0x06, 0xad, 0x2b, 0x56, 0x00, 0x1f, 0xa9, 0xde, 0x00, 0x58, 0x3e, 0x50, 0xee, 0x8a,
	0xe1, 0xf4, 0x88, 0xd6, 0x40, 0xef, 0x8b, 0xd4, 0x5b, 0x1a, 0xc7, 0xe4, 0x85, 0x41, 0x74, 0x1f,
	0x6a, 0xd2, 0x9f, 0xc3, 0x2c, 0xd6, 0x87, 0x49, 0xa7, 0x0e, 0x58, 0x34, 0x3e, 0x41, 0xff, 0x4f,
	0x11, 0xfa, 0x2b, 0x55, 0x5a, 0x54, 0xee, 0xea, 0x8e, 0xb3, 0xd7, 0xea, 0x6e, 0xdf, 0x4d, 0x02,
	0x77, 0xf1, 0x27, 0x33, 0x48, 0x98, 0x6c, 0xd4, 0xb1, 0xb3, 0x69, 0x08, 0x63, 0xd4, 0x2a, 0xa0,
	0xe0, 0x6e, 0xdb, 0x0c, 0x61, 0xec, 0xa3, 0xf6, 0xe7, 0x8b, 0xa0, 0x71, 0x7b, 0x11, 0x34, 0xc2,
	0x18, 0xdd, 0x9b, 0xd3, 0x2f, 0x01, 0x59, 0xf1, 0x52, 0x82, 0x9e, 0x5e, 0xc2, 0xfb, 0x11, 0x94,
	0x14, 0x8c, 0x88, 0xad, 0x64, 0xe6, 0x87, 0x9f, 0x1c, 0xb4, 0x71, 0x2a, 0xf3, 0x04, 0x72, 0x26,
	0x15, 0x88, 0xe3, 0x41, 0xca, 0x4a, 0xbc, 0x85, 0xda, 0x54, 0x1b, 0x3d, 0xd6, 0xb7, 0xaa, 0x27,
	0xab, 0xc6, 0x3f, 0xe9, 0xeb, 0x06, 0xcc, 0x61, 0x58, 0xb2, 0x0d, 0x68, 0x1b, 0xbb, 0x68, 0x75,
	0xaa, 0x73, 0x73, 0x86, 0x36, 0x32, 0xee, 0xa2, 0x75, 0xab, 0x40, 0x2f, 0xed, 0xf7, 0x05, 0x48,
	0x59, 0xeb, 0xd2, 0xb1, 0xd1, 0xe7, 0x36, 0x18, 0x7a, 0xc8, 0xfd, 0xb5, 0x87, 0x69, 0xf3, 0xfb,
	0xb7, 0x0e, 0x6a, 0x9e, 0xca, 0x1c, 0x0f, 0x50, 0x7b, 0x76, 0x30, 0x1f, 0x91, 0xdf, 0x5d, 0x0f,
	0x32, 0xb7, 0x03, 0x2f, 0xfe, 0x6b, 0xe8, 0x6c, 0x5d, 0x1f, 0x50, 0x67, 0x71, 0x1d, 0xe4, 0x8f,
	0xff, 0x58, 0xc0, 0x7b, 0x07, 0xff, 0x86, 0x9f, 0x12, 0x77, 0x5f, 0x5d, 0x5e, 0xfb, 0xce, 0xd5,
	0xb5, 0xef, 0x7c, 0xbf, 0xf6, 0x9d, 0x2f, 0x37, 0x7e, 0xe3, 0xea, 0xc6, 0x6f, 0x7c, 0xbb, 0xf1,
	0x1b, 0xaf, 0x0f, 0x73, 0xa6, 0x06, 0xa3, 0x8c, 0x50, 0x5e, 0x44, 0xf5, 0x2d, 0x67, 0x19, 0x7d,
	0x9c, 0xf3, 0x68, 0x7c, 0x10, 0x15, 0xbc, 0x3f, 0x1a, 0x82, 0xd4, 0xaf, 0xca, 0xdc, 0x6b, 0xa2,
	0xce, 0x2b, 0x90, 0xd9, 0x8a, 0xb9, 0xd9, 0x4f, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x08,
	0x1e, 0x72, 0x77, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Transfer defines a rpc handler method for MsgTransfer.
	Transfer(ctx context.Context, in *MsgTransfer, opts ...grpc.CallOption) (*MsgTransferResponse, error)
	// RegisterChain defines a rpc handler method for MsgRegisterChain.
	RegisterChain(ctx context.Context, in *MsgRegisterChain, opts ...grpc.CallOption) (*MsgRegisterChainResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Transfer(ctx context.Context, in *MsgTransfer, opts ...grpc.CallOption) (*MsgTransferResponse, error) {
	out := new(MsgTransferResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.transfer.v1.Msg/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterChain(ctx context.Context, in *MsgRegisterChain, opts ...grpc.CallOption) (*MsgRegisterChainResponse, error) {
	out := new(MsgRegisterChainResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.transfer.v1.Msg/RegisterChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Transfer defines a rpc handler method for MsgTransfer.
	Transfer(context.Context, *MsgTransfer) (*MsgTransferResponse, error)
	// RegisterChain defines a rpc handler method for MsgRegisterChain.
	RegisterChain(context.Context, *MsgRegisterChain) (*MsgRegisterChainResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Transfer(ctx context.Context, req *MsgTransfer) (*MsgTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (*UnimplementedMsgServer) RegisterChain(ctx context.Context, req *MsgRegisterChain) (*MsgRegisterChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterChain not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.transfer.v1.Msg/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Transfer(ctx, req.(*MsgTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.transfer.v1.Msg/RegisterChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterChain(ctx, req.(*MsgRegisterChain))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.transfer.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Msg_Transfer_Handler,
		},
		{
			MethodName: "RegisterChain",
			Handler:    _Msg_RegisterChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/transfer/v1/tx.proto",
}

func (m *MsgTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x42
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SenderAddress) > 0 {
		i -= len(m.SenderAddress)
		copy(dAtA[i:], m.SenderAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SenderAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterChainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterChainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterChainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Token.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgTransferResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovTx(uint64(m.Sequence))
	}
	return n
}

func (m *MsgRegisterChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SenderAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterChainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransferResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransferResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterChainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterChainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterChainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
