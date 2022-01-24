// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the incentives module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module
	Params            Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Gauges            []Gauge            `protobuf:"bytes,2,rep,name=gauges,proto3" json:"gauges"`
	LockableDurations []time.Duration    `protobuf:"bytes,3,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
	GenesisReward     []GenesisReward    `protobuf:"bytes,4,rep,name=genesis_reward,json=genesisReward,proto3" json:"genesis_reward"`
	PeriodLockReward  []PeriodLockReward `protobuf:"bytes,5,rep,name=period_lock_reward,json=periodLockReward,proto3" json:"period_lock_reward"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a288ccc95d977d2d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGauges() []Gauge {
	if m != nil {
		return m.Gauges
	}
	return nil
}

func (m *GenesisState) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

func (m *GenesisState) GetGenesisReward() []GenesisReward {
	if m != nil {
		return m.GenesisReward
	}
	return nil
}

func (m *GenesisState) GetPeriodLockReward() []PeriodLockReward {
	if m != nil {
		return m.PeriodLockReward
	}
	return nil
}

type GenesisReward struct {
	CurrentReward    CurrentReward      `protobuf:"bytes,1,opt,name=current_reward,json=currentReward,proto3" json:"current_reward"`
	HistoricalReward []HistoricalReward `protobuf:"bytes,2,rep,name=historical_reward,json=historicalReward,proto3" json:"historical_reward"`
}

func (m *GenesisReward) Reset()         { *m = GenesisReward{} }
func (m *GenesisReward) String() string { return proto.CompactTextString(m) }
func (*GenesisReward) ProtoMessage()    {}
func (*GenesisReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_a288ccc95d977d2d, []int{1}
}
func (m *GenesisReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisReward.Merge(m, src)
}
func (m *GenesisReward) XXX_Size() int {
	return m.Size()
}
func (m *GenesisReward) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisReward.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisReward proto.InternalMessageInfo

func (m *GenesisReward) GetCurrentReward() CurrentReward {
	if m != nil {
		return m.CurrentReward
	}
	return CurrentReward{}
}

func (m *GenesisReward) GetHistoricalReward() []HistoricalReward {
	if m != nil {
		return m.HistoricalReward
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.incentives.GenesisState")
	proto.RegisterType((*GenesisReward)(nil), "osmosis.incentives.GenesisReward")
}

func init() { proto.RegisterFile("osmosis/incentives/genesis.proto", fileDescriptor_a288ccc95d977d2d) }

var fileDescriptor_a288ccc95d977d2d = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0x87, 0x93, 0xdb, 0x4b, 0x07, 0x5f, 0x2e, 0xe2, 0x5a, 0x0c, 0x69, 0x87, 0xa4, 0x44, 0x20,
	0x75, 0x21, 0x91, 0xca, 0x00, 0x62, 0x2c, 0x48, 0x65, 0xa8, 0x10, 0x2a, 0x03, 0x88, 0xa5, 0x72,
	0x52, 0xe3, 0x5a, 0x4d, 0xe2, 0xc8, 0x76, 0x80, 0xbe, 0x05, 0x23, 0x03, 0x0f, 0xc2, 0x23, 0x74,
	0xec, 0xc8, 0x54, 0x50, 0xfb, 0x06, 0x3c, 0x01, 0x8a, 0xff, 0x40, 0xd3, 0x86, 0xad, 0x3e, 0xe7,
	0xf3, 0xef, 0x7c, 0xf5, 0x09, 0x18, 0x30, 0x91, 0x33, 0x41, 0x45, 0x4c, 0x8b, 0x14, 0x17, 0x92,
	0x7e, 0xc4, 0x22, 0x26, 0xb8, 0xc0, 0x82, 0x8a, 0xa8, 0xe4, 0x4c, 0x32, 0x08, 0x0d, 0x11, 0xfd,
	0x23, 0xfa, 0xf7, 0x08, 0x23, 0x4c, 0xb5, 0xe3, 0xfa, 0x97, 0x26, 0xfb, 0x3e, 0x61, 0x8c, 0x64,
	0x38, 0x56, 0xa7, 0xa4, 0xfa, 0x10, 0x2f, 0x2a, 0x8e, 0x24, 0x65, 0x85, 0xe9, 0x07, 0x2d, 0xb3,
	0x4a, 0xc4, 0x51, 0x2e, 0x6c, 0x40, 0x9b, 0x0c, 0xaa, 0x08, 0xd6, 0xfd, 0xf0, 0x5b, 0x07, 0xdc,
	0x9e, 0x68, 0xb9, 0x37, 0x12, 0x49, 0x0c, 0x9f, 0x82, 0xae, 0x0e, 0xf0, 0xdc, 0x81, 0x3b, 0xbc,
	0x1a, 0xf5, 0xa3, 0x73, 0xd9, 0xe8, 0xb5, 0x22, 0xc6, 0x97, 0x9b, 0x5d, 0xe0, 0xcc, 0x0c, 0x0f,
	0x9f, 0x80, 0xae, 0x4a, 0x16, 0xde, 0xc5, 0xa0, 0x33, 0xbc, 0x1a, 0xf5, 0xda, 0x6e, 0x4e, 0x6a,
	0xc2, 0x5e, 0xd4, 0x38, 0x64, 0x00, 0x66, 0x2c, 0x5d, 0xa1, 0x24, 0xc3, 0x73, 0xfb, 0xff, 0x84,
	0xd7, 0x31, 0x21, 0xfa, 0x05, 0x22, 0xfb, 0x02, 0xd1, 0x0b, 0x43, 0x8c, 0x1f, 0xd6, 0x21, 0xbf,
	0x77, 0x41, 0x6f, 0x8d, 0xf2, 0xec, 0x59, 0x78, 0x1e, 0x11, 0x7e, 0xfd, 0x19, 0xb8, 0xb3, 0x1b,
	0xdb, 0xb0, 0x17, 0x05, 0x7c, 0x05, 0xee, 0x98, 0x85, 0xcc, 0x39, 0xfe, 0x84, 0xf8, 0xc2, 0xbb,
	0x54, 0xc3, 0xee, 0xb7, 0x1a, 0x6b, 0x72, 0xa6, 0x40, 0x63, 0x7e, 0x4d, 0x8e, 0x8b, 0xf0, 0x1d,
	0x80, 0x25, 0xe6, 0x94, 0x2d, 0xe6, 0xf5, 0x2c, 0x9b, 0x79, 0x4b, 0x65, 0x3e, 0x68, 0x7d, 0x3f,
	0x45, 0x4f, 0x59, 0xba, 0x6a, 0xc4, 0xde, 0x2d, 0x4f, 0xea, 0xe1, 0x77, 0x17, 0x5c, 0x37, 0x04,
	0x6a, 0xf7, 0xb4, 0xe2, 0x1c, 0x17, 0xd2, 0xce, 0xd1, 0x7b, 0x6a, 0x75, 0x7f, 0xae, 0xc9, 0xa6,
	0x7b, 0x7a, 0x5c, 0x84, 0x6f, 0xc1, 0xcd, 0x92, 0x0a, 0xc9, 0x38, 0x4d, 0x51, 0x66, 0x23, 0x2f,
	0xfe, 0xaf, 0xfe, 0xf2, 0x2f, 0xdc, 0x54, 0x5f, 0x9e, 0xd6, 0xa7, 0x9b, 0xbd, 0xef, 0x6e, 0xf7,
	0xbe, 0xfb, 0x6b, 0xef, 0xbb, 0x5f, 0x0e, 0xbe, 0xb3, 0x3d, 0xf8, 0xce, 0x8f, 0x83, 0xef, 0xbc,
	0x1f, 0x11, 0x2a, 0x97, 0x55, 0x12, 0xa5, 0x2c, 0x8f, 0xcd, 0x84, 0x47, 0x19, 0x4a, 0x84, 0x3d,
	0xc4, 0x9f, 0x8f, 0xbf, 0x56, 0xb9, 0x2e, 0xb1, 0x48, 0xba, 0x6a, 0xff, 0x8f, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x05, 0x80, 0x91, 0x59, 0x5d, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PeriodLockReward) > 0 {
		for iNdEx := len(m.PeriodLockReward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodLockReward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.GenesisReward) > 0 {
		for iNdEx := len(m.GenesisReward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisReward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.LockableDurations) > 0 {
		for iNdEx := len(m.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockableDurations[iNdEx], dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockableDurations[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintGenesis(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Gauges) > 0 {
		for iNdEx := len(m.Gauges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Gauges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HistoricalReward) > 0 {
		for iNdEx := len(m.HistoricalReward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HistoricalReward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.CurrentReward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Gauges) > 0 {
		for _, e := range m.Gauges {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GenesisReward) > 0 {
		for _, e := range m.GenesisReward {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PeriodLockReward) > 0 {
		for _, e := range m.PeriodLockReward {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CurrentReward.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.HistoricalReward) > 0 {
		for _, e := range m.HistoricalReward {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gauges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gauges = append(m.Gauges, Gauge{})
			if err := m.Gauges[len(m.Gauges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockableDurations = append(m.LockableDurations, time.Duration(0))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&(m.LockableDurations[len(m.LockableDurations)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisReward = append(m.GenesisReward, GenesisReward{})
			if err := m.GenesisReward[len(m.GenesisReward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodLockReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodLockReward = append(m.PeriodLockReward, PeriodLockReward{})
			if err := m.PeriodLockReward[len(m.PeriodLockReward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HistoricalReward = append(m.HistoricalReward, HistoricalReward{})
			if err := m.HistoricalReward[len(m.HistoricalReward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)