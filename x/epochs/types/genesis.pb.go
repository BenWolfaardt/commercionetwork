// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commercionetwork/epochs/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type EpochInfo struct {
	Identifier            string        `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	StartTime             time.Time     `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	Duration              time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	CurrentEpoch          int64         `protobuf:"varint,4,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	CurrentEpochStartTime time.Time     `protobuf:"bytes,5,opt,name=current_epoch_start_time,json=currentEpochStartTime,proto3,stdtime" json:"current_epoch_start_time" yaml:"current_epoch_start_time"`
	EpochCountingStarted  bool          `protobuf:"varint,6,opt,name=epoch_counting_started,json=epochCountingStarted,proto3" json:"epoch_counting_started,omitempty"`
	CurrentEpochEnded     bool          `protobuf:"varint,7,opt,name=current_epoch_ended,json=currentEpochEnded,proto3" json:"current_epoch_ended,omitempty"`
}

func (m *EpochInfo) Reset()         { *m = EpochInfo{} }
func (m *EpochInfo) String() string { return proto.CompactTextString(m) }
func (*EpochInfo) ProtoMessage()    {}
func (*EpochInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f03e895ddd7de90, []int{0}
}
func (m *EpochInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochInfo.Merge(m, src)
}
func (m *EpochInfo) XXX_Size() int {
	return m.Size()
}
func (m *EpochInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EpochInfo proto.InternalMessageInfo

func (m *EpochInfo) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EpochInfo) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpochStartTime() time.Time {
	if m != nil {
		return m.CurrentEpochStartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetEpochCountingStarted() bool {
	if m != nil {
		return m.EpochCountingStarted
	}
	return false
}

func (m *EpochInfo) GetCurrentEpochEnded() bool {
	if m != nil {
		return m.CurrentEpochEnded
	}
	return false
}

// GenesisState defines the epochs module's genesis state.
type GenesisState struct {
	Epochs []EpochInfo `protobuf:"bytes,1,rep,name=epochs,proto3" json:"epochs"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f03e895ddd7de90, []int{1}
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

func (m *GenesisState) GetEpochs() []EpochInfo {
	if m != nil {
		return m.Epochs
	}
	return nil
}

func init() {
	proto.RegisterType((*EpochInfo)(nil), "commercionetwork.commercionetwork.epochs.EpochInfo")
	proto.RegisterType((*GenesisState)(nil), "commercionetwork.commercionetwork.epochs.GenesisState")
}

func init() {
	proto.RegisterFile("commercionetwork/epochs/genesis.proto", fileDescriptor_2f03e895ddd7de90)
}

var fileDescriptor_2f03e895ddd7de90 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0xef, 0xd8, 0xb5, 0x6e, 0x67, 0x57, 0x64, 0xc7, 0x55, 0x62, 0xc1, 0x24, 0x44, 0x84, 0x80,
	0x92, 0xe0, 0xae, 0x27, 0xbd, 0x55, 0x17, 0xf1, 0x68, 0xaa, 0x20, 0x5e, 0x4a, 0x9a, 0x4c, 0xd3,
	0xc1, 0xcd, 0x4c, 0x98, 0xbc, 0xa0, 0xbd, 0xf9, 0x11, 0xf6, 0xe8, 0x47, 0xda, 0xe3, 0x1e, 0x3d,
	0x55, 0x69, 0x6f, 0x1e, 0xf7, 0x2e, 0xc8, 0xcc, 0x24, 0xb5, 0xdb, 0x20, 0x7a, 0x4b, 0xde, 0xef,
	0xdf, 0xbc, 0xc7, 0x7b, 0xf8, 0x61, 0x22, 0xf2, 0x9c, 0xca, 0x84, 0x09, 0x4e, 0xe1, 0x93, 0x90,
	0x1f, 0x43, 0x5a, 0x88, 0x64, 0x56, 0x86, 0x19, 0xe5, 0xb4, 0x64, 0x65, 0x50, 0x48, 0x01, 0x82,
	0xf8, 0xdb, 0xb4, 0xa0, 0x55, 0x30, 0xba, 0xc1, 0x61, 0x26, 0x32, 0xa1, 0x45, 0xa1, 0xfa, 0x32,
	0xfa, 0x81, 0x9d, 0x09, 0x91, 0x9d, 0xd2, 0x50, 0xff, 0x4d, 0xaa, 0x69, 0x98, 0x56, 0x32, 0x06,
	0x26, 0x78, 0x8d, 0x3b, 0xdb, 0x38, 0xb0, 0x9c, 0x96, 0x10, 0xe7, 0x85, 0x21, 0x78, 0xbf, 0xba,
	0xb8, 0x7f, 0xa2, 0x12, 0x5e, 0xf3, 0xa9, 0x20, 0x36, 0xc6, 0x2c, 0xa5, 0x1c, 0xd8, 0x94, 0x51,
	0x69, 0x21, 0x17, 0xf9, 0xfd, 0x68, 0xa3, 0x42, 0xde, 0x63, 0x5c, 0x42, 0x2c, 0x61, 0xac, 0x6c,
	0xac, 0x6b, 0x2e, 0xf2, 0xf7, 0x8e, 0x06, 0x81, 0xc9, 0x08, 0x9a, 0x8c, 0xe0, 0x6d, 0x93, 0x31,
	0xbc, 0x7f, 0xbe, 0x70, 0x3a, 0x97, 0x0b, 0xe7, 0x60, 0x1e, 0xe7, 0xa7, 0xcf, 0xbc, 0x3f, 0x5a,
	0xef, 0xec, 0xbb, 0x83, 0xa2, 0xbe, 0x2e, 0x28, 0x3a, 0x99, 0xe1, 0xdd, 0xe6, 0xe9, 0x56, 0x57,
	0xfb, 0xde, 0x6b, 0xf9, 0xbe, 0xac, 0x09, 0xc3, 0x27, 0xca, 0xf6, 0xe7, 0xc2, 0x21, 0x8d, 0xe4,
	0xb1, 0xc8, 0x19, 0xd0, 0xbc, 0x80, 0xf9, 0xe5, 0xc2, 0xb9, 0x65, 0xc2, 0x1a, 0xcc, 0xfb, 0xaa,
	0xa2, 0xd6, 0xee, 0xe4, 0x01, 0xbe, 0x99, 0x54, 0x52, 0x52, 0x0e, 0x63, 0x3d, 0x5a, 0x6b, 0xc7,
	0x45, 0x7e, 0x37, 0xda, 0xaf, 0x8b, 0x7a, 0x18, 0xe4, 0x0b, 0xc2, 0xd6, 0x15, 0xd6, 0x78, 0xa3,
	0xef, 0xeb, 0xff, 0xec, 0xfb, 0x51, 0xdd, 0xb7, 0x63, 0x9e, 0xf2, 0x37, 0x27, 0x33, 0x85, 0x3b,
	0x9b, 0xc9, 0xa3, 0xf5, 0x44, 0x9e, 0xe2, 0xbb, 0x86, 0x9f, 0x88, 0x8a, 0x03, 0xe3, 0x99, 0x11,
	0xd2, 0xd4, 0xea, 0xb9, 0xc8, 0xdf, 0x8d, 0x0e, 0x35, 0xfa, 0xa2, 0x06, 0x47, 0x06, 0x23, 0x01,
	0xbe, 0x7d, 0x35, 0x8d, 0xf2, 0x94, 0xa6, 0xd6, 0x0d, 0x2d, 0x39, 0xd8, 0x4c, 0x3a, 0x51, 0x80,
	0x17, 0xe3, 0xfd, 0x57, 0x66, 0x23, 0x47, 0x10, 0x03, 0x25, 0x6f, 0x70, 0xcf, 0x2c, 0x9c, 0x85,
	0xdc, 0xae, 0xbf, 0x77, 0x74, 0x1c, 0xfc, 0xef, 0x86, 0x06, 0xeb, 0x35, 0x1a, 0xee, 0xa8, 0xf6,
	0xa3, 0xda, 0x68, 0xf8, 0xee, 0x7c, 0x69, 0xa3, 0x8b, 0xa5, 0x8d, 0x7e, 0x2c, 0x6d, 0x74, 0xb6,
	0xb2, 0x3b, 0x17, 0x2b, 0xbb, 0xf3, 0x6d, 0x65, 0x77, 0x3e, 0x3c, 0xcf, 0x18, 0xcc, 0xaa, 0x89,
	0x72, 0x0c, 0x5b, 0xf7, 0xd2, 0x2a, 0x7c, 0x6e, 0x4e, 0x08, 0xe6, 0x05, 0x2d, 0x27, 0x3d, 0x3d,
	0xf7, 0xe3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x4a, 0x93, 0x61, 0x6a, 0x03, 0x00, 0x00,
}

func (m *EpochInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpochEnded {
		i--
		if m.CurrentEpochEnded {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.EpochCountingStarted {
		i--
		if m.EpochCountingStarted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CurrentEpochStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.CurrentEpoch != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.Epochs) > 0 {
		for iNdEx := len(m.Epochs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Epochs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
func (m *EpochInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovGenesis(uint64(l))
	if m.CurrentEpoch != 0 {
		n += 1 + sovGenesis(uint64(m.CurrentEpoch))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime)
	n += 1 + l + sovGenesis(uint64(l))
	if m.EpochCountingStarted {
		n += 2
	}
	if m.CurrentEpochEnded {
		n += 2
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Epochs) > 0 {
		for _, e := range m.Epochs {
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
func (m *EpochInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EpochInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartTime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CurrentEpochStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochCountingStarted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EpochCountingStarted = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochEnded", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CurrentEpochEnded = bool(v != 0)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
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
			m.Epochs = append(m.Epochs, EpochInfo{})
			if err := m.Epochs[len(m.Epochs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
