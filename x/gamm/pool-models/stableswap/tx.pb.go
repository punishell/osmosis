// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/pool-models/stableswap/tx.proto

package stableswap

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type MsgCreateStableswapPool struct {
	Sender               string                                   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	PoolParams           *PoolParams                              `protobuf:"bytes,2,opt,name=poolParams,proto3" json:"poolParams,omitempty" yaml:"pool_params"`
	InitialPoolLiquidity github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=initial_pool_liquidity,json=initialPoolLiquidity,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_pool_liquidity"`
	FuturePoolGovernor   string                                   `protobuf:"bytes,4,opt,name=future_pool_governor,json=futurePoolGovernor,proto3" json:"future_pool_governor,omitempty" yaml:"future_pool_governor"`
}

func (m *MsgCreateStableswapPool) Reset()         { *m = MsgCreateStableswapPool{} }
func (m *MsgCreateStableswapPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStableswapPool) ProtoMessage()    {}
func (*MsgCreateStableswapPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{0}
}
func (m *MsgCreateStableswapPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStableswapPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStableswapPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStableswapPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStableswapPool.Merge(m, src)
}
func (m *MsgCreateStableswapPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStableswapPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStableswapPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStableswapPool proto.InternalMessageInfo

func (m *MsgCreateStableswapPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgCreateStableswapPool) GetPoolParams() *PoolParams {
	if m != nil {
		return m.PoolParams
	}
	return nil
}

func (m *MsgCreateStableswapPool) GetInitialPoolLiquidity() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialPoolLiquidity
	}
	return nil
}

func (m *MsgCreateStableswapPool) GetFuturePoolGovernor() string {
	if m != nil {
		return m.FuturePoolGovernor
	}
	return ""
}

type MsgCreateStableswapPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateStableswapPoolResponse) Reset()         { *m = MsgCreateStableswapPoolResponse{} }
func (m *MsgCreateStableswapPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStableswapPoolResponse) ProtoMessage()    {}
func (*MsgCreateStableswapPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{1}
}
func (m *MsgCreateStableswapPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStableswapPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStableswapPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStableswapPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStableswapPoolResponse.Merge(m, src)
}
func (m *MsgCreateStableswapPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStableswapPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStableswapPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStableswapPoolResponse proto.InternalMessageInfo

func (m *MsgCreateStableswapPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateStableswapPool)(nil), "osmosis.gamm.poolmodels.stableswap.v1beta1.MsgCreateStableswapPool")
	proto.RegisterType((*MsgCreateStableswapPoolResponse)(nil), "osmosis.gamm.poolmodels.stableswap.v1beta1.MsgCreateStableswapPoolResponse")
}

func init() {
	proto.RegisterFile("osmosis/gamm/pool-models/stableswap/tx.proto", fileDescriptor_46b7c8a0f24de97c)
}

var fileDescriptor_46b7c8a0f24de97c = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x6b, 0x7a, 0x2a, 0xc2, 0x27, 0x06, 0xac, 0xea, 0x28, 0x45, 0x4a, 0xaa, 0xb0, 0x14,
	0x44, 0x6d, 0xae, 0x48, 0x20, 0x98, 0x50, 0x8a, 0x40, 0x27, 0x38, 0xa9, 0x84, 0xed, 0x96, 0xca,
	0x69, 0x4c, 0xb0, 0x48, 0xe2, 0x10, 0xbb, 0xe5, 0x3a, 0x32, 0xb3, 0xf0, 0x31, 0x10, 0x13, 0x1f,
	0xe3, 0xc6, 0x1b, 0x99, 0x02, 0x6a, 0xbf, 0x41, 0x3f, 0x01, 0xb2, 0x9d, 0xf4, 0x6e, 0xb8, 0x43,
	0x45, 0x62, 0xea, 0xdb, 0x27, 0xbf, 0xf7, 0x7d, 0xde, 0x3f, 0x86, 0xf7, 0x85, 0x4c, 0x85, 0xe4,
	0x92, 0xc4, 0x34, 0x4d, 0x49, 0x2e, 0x44, 0x32, 0x48, 0x45, 0xc4, 0x12, 0x49, 0xa4, 0xa2, 0x61,
	0xc2, 0xe4, 0x27, 0x9a, 0x13, 0x75, 0x8c, 0xf3, 0x42, 0x28, 0x81, 0xee, 0x55, 0x34, 0xd6, 0x34,
	0xd6, 0xb4, 0x85, 0xf1, 0x19, 0x8c, 0xe7, 0xfb, 0x21, 0x53, 0x74, 0xbf, 0xeb, 0x4c, 0x0d, 0x4c,
	0x42, 0x2a, 0x19, 0xa9, 0x44, 0x32, 0x15, 0x3c, 0xb3, 0xb5, 0xba, 0xed, 0x58, 0xc4, 0xc2, 0x84,
	0x44, 0x47, 0x95, 0xfa, 0x64, 0x9b, 0x7e, 0xce, 0xc2, 0x89, 0x26, 0x6c, 0xaa, 0xf7, 0xa5, 0x09,
	0x6f, 0x1e, 0xca, 0x78, 0x54, 0x30, 0xaa, 0xd8, 0xdb, 0x0d, 0x32, 0x16, 0x22, 0x41, 0x77, 0x61,
	0x4b, 0xb2, 0x2c, 0x62, 0x45, 0x07, 0xf4, 0x40, 0xff, 0x9a, 0x7f, 0x63, 0x5d, 0xba, 0xd7, 0x17,
	0x34, 0x4d, 0x9e, 0x7a, 0x56, 0xf7, 0x82, 0x0a, 0x40, 0x19, 0x84, 0xba, 0xe8, 0x98, 0x16, 0x34,
	0x95, 0x9d, 0x2b, 0x3d, 0xd0, 0xdf, 0x1d, 0x3e, 0xc2, 0xdb, 0x0f, 0x8e, 0xc7, 0x9b, 0x6c, 0x7f,
	0x6f, 0x5d, 0xba, 0xc8, 0xda, 0xe8, 0x9c, 0x49, 0x6e, 0x64, 0x2f, 0x38, 0xe7, 0x80, 0x3e, 0x03,
	0xb8, 0xc7, 0x33, 0xae, 0x38, 0x4d, 0xcc, 0x34, 0x93, 0x84, 0x7f, 0x9c, 0xf1, 0x88, 0xab, 0x45,
	0xa7, 0xd9, 0x6b, 0xf6, 0x77, 0x87, 0xb7, 0xb0, 0xdd, 0x24, 0xd6, 0x9b, 0xdc, 0xb8, 0x8c, 0x04,
	0xcf, 0xfc, 0x07, 0x27, 0xa5, 0xdb, 0xf8, 0xfe, 0xcb, 0xed, 0xc7, 0x5c, 0xbd, 0x9f, 0x85, 0x78,
	0x2a, 0x52, 0x52, 0xad, 0xdd, 0xfe, 0x0c, 0x64, 0xf4, 0x81, 0xa8, 0x45, 0xce, 0xa4, 0x49, 0x90,
	0x41, 0xbb, 0xb2, 0xd2, 0x4d, 0xbe, 0xae, 0x8d, 0xd0, 0x1b, 0xd8, 0x7e, 0x37, 0x53, 0xb3, 0x82,
	0xd9, 0x0e, 0x62, 0x31, 0x67, 0x45, 0x26, 0x8a, 0xce, 0x8e, 0x59, 0x96, 0xbb, 0x2e, 0xdd, 0xdb,
	0x76, 0x8a, 0x8b, 0x28, 0x2f, 0x40, 0x56, 0xd6, 0x35, 0x5f, 0xd6, 0xe2, 0x0b, 0xe8, 0x5e, 0x72,
	0x8c, 0x80, 0xc9, 0x5c, 0x64, 0x92, 0xa1, 0x3b, 0xf0, 0xaa, 0x29, 0xc4, 0x23, 0x73, 0x95, 0x1d,
	0x1f, 0x2e, 0x4b, 0xb7, 0xa5, 0x91, 0x83, 0xe7, 0x41, 0x4b, 0x7f, 0x3a, 0x88, 0x86, 0x3f, 0x00,
	0x6c, 0x1e, 0xca, 0x18, 0x7d, 0x03, 0xb0, 0x7d, 0xe1, 0x69, 0x47, 0xff, 0x72, 0x9b, 0x4b, 0x5a,
	0xea, 0xbe, 0xfa, 0x0f, 0x45, 0xea, 0xb9, 0xfc, 0xa3, 0x93, 0xa5, 0x03, 0x4e, 0x97, 0x0e, 0xf8,
	0xbd, 0x74, 0xc0, 0xd7, 0x95, 0xd3, 0x38, 0x5d, 0x39, 0x8d, 0x9f, 0x2b, 0xa7, 0x71, 0xf4, 0xec,
	0xdc, 0x9d, 0x2a, 0xc3, 0x41, 0x42, 0x43, 0x59, 0xff, 0x21, 0xf3, 0xc7, 0xe4, 0xf8, 0x6f, 0x4f,
	0x3f, 0x6c, 0x99, 0xb7, 0xfe, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x53, 0xea, 0x93,
	0xb8, 0x03, 0x00, 0x00,
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
	CreateStableswapPool(ctx context.Context, in *MsgCreateStableswapPool, opts ...grpc.CallOption) (*MsgCreateStableswapPoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateStableswapPool(ctx context.Context, in *MsgCreateStableswapPool, opts ...grpc.CallOption) (*MsgCreateStableswapPoolResponse, error) {
	out := new(MsgCreateStableswapPoolResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.poolmodels.stableswap.v1beta1.Msg/CreateStableswapPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateStableswapPool(context.Context, *MsgCreateStableswapPool) (*MsgCreateStableswapPoolResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateStableswapPool(ctx context.Context, req *MsgCreateStableswapPool) (*MsgCreateStableswapPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStableswapPool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateStableswapPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateStableswapPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateStableswapPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.poolmodels.stableswap.v1beta1.Msg/CreateStableswapPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateStableswapPool(ctx, req.(*MsgCreateStableswapPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.gamm.poolmodels.stableswap.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStableswapPool",
			Handler:    _Msg_CreateStableswapPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/gamm/pool-models/stableswap/tx.proto",
}

func (m *MsgCreateStableswapPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStableswapPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStableswapPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FuturePoolGovernor) > 0 {
		i -= len(m.FuturePoolGovernor)
		copy(dAtA[i:], m.FuturePoolGovernor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FuturePoolGovernor)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.InitialPoolLiquidity) > 0 {
		for iNdEx := len(m.InitialPoolLiquidity) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialPoolLiquidity[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolParams != nil {
		{
			size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateStableswapPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStableswapPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStableswapPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
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
func (m *MsgCreateStableswapPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolParams != nil {
		l = m.PoolParams.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.InitialPoolLiquidity) > 0 {
		for _, e := range m.InitialPoolLiquidity {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.FuturePoolGovernor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateStableswapPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateStableswapPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateStableswapPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStableswapPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
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
			if m.PoolParams == nil {
				m.PoolParams = &PoolParams{}
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPoolLiquidity", wireType)
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
			m.InitialPoolLiquidity = append(m.InitialPoolLiquidity, types.Coin{})
			if err := m.InitialPoolLiquidity[len(m.InitialPoolLiquidity)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuturePoolGovernor", wireType)
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
			m.FuturePoolGovernor = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateStableswapPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateStableswapPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStableswapPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
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
