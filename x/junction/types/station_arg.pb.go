// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/junction/station_arg.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type StationArg struct {
	TrackType string `protobuf:"bytes,1,opt,name=trackType,proto3" json:"trackType,omitempty"`
	DaType    string `protobuf:"bytes,2,opt,name=daType,proto3" json:"daType,omitempty"`
	Prover    string `protobuf:"bytes,3,opt,name=prover,proto3" json:"prover,omitempty"`
}

func (m *StationArg) Reset()         { *m = StationArg{} }
func (m *StationArg) String() string { return proto.CompactTextString(m) }
func (*StationArg) ProtoMessage()    {}
func (*StationArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_470a88e81ffac399, []int{0}
}
func (m *StationArg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StationArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StationArg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StationArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StationArg.Merge(m, src)
}
func (m *StationArg) XXX_Size() int {
	return m.Size()
}
func (m *StationArg) XXX_DiscardUnknown() {
	xxx_messageInfo_StationArg.DiscardUnknown(m)
}

var xxx_messageInfo_StationArg proto.InternalMessageInfo

func (m *StationArg) GetTrackType() string {
	if m != nil {
		return m.TrackType
	}
	return ""
}

func (m *StationArg) GetDaType() string {
	if m != nil {
		return m.DaType
	}
	return ""
}

func (m *StationArg) GetProver() string {
	if m != nil {
		return m.Prover
	}
	return ""
}

func init() {
	proto.RegisterType((*StationArg)(nil), "junction.junction.StationArg")
}

func init() {
	proto.RegisterFile("junction/junction/station_arg.proto", fileDescriptor_470a88e81ffac399)
}

var fileDescriptor_470a88e81ffac399 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x2a, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x87, 0x33, 0x8a, 0x4b, 0x12, 0x41, 0x74, 0x7c, 0x62, 0x51, 0xba,
	0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x20, 0x4c, 0x4e, 0x0f, 0xc6, 0x50, 0x8a, 0xe2, 0xe2,
	0x0a, 0x86, 0xa8, 0x73, 0x2c, 0x4a, 0x17, 0x92, 0xe1, 0xe2, 0x2c, 0x29, 0x4a, 0x4c, 0xce, 0x0e,
	0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x89, 0x71, 0xb1,
	0xa5, 0x24, 0x82, 0xa5, 0x98, 0xc0, 0x52, 0x50, 0x1e, 0x48, 0xbc, 0xa0, 0x28, 0xbf, 0x2c, 0xb5,
	0x48, 0x82, 0x19, 0x22, 0x0e, 0xe1, 0x39, 0x05, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x94, 0x59, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x62,
	0x66, 0x51, 0x72, 0x46, 0x62, 0x66, 0x5e, 0xb1, 0x6e, 0x5e, 0x6a, 0x49, 0x79, 0x7e, 0x51, 0x36,
	0xc2, 0x07, 0x15, 0x08, 0x66, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x1f, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0xd2, 0x75, 0x27, 0xee, 0x00, 0x00, 0x00,
}

func (m *StationArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StationArg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StationArg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Prover) > 0 {
		i -= len(m.Prover)
		copy(dAtA[i:], m.Prover)
		i = encodeVarintStationArg(dAtA, i, uint64(len(m.Prover)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DaType) > 0 {
		i -= len(m.DaType)
		copy(dAtA[i:], m.DaType)
		i = encodeVarintStationArg(dAtA, i, uint64(len(m.DaType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TrackType) > 0 {
		i -= len(m.TrackType)
		copy(dAtA[i:], m.TrackType)
		i = encodeVarintStationArg(dAtA, i, uint64(len(m.TrackType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStationArg(dAtA []byte, offset int, v uint64) int {
	offset -= sovStationArg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StationArg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TrackType)
	if l > 0 {
		n += 1 + l + sovStationArg(uint64(l))
	}
	l = len(m.DaType)
	if l > 0 {
		n += 1 + l + sovStationArg(uint64(l))
	}
	l = len(m.Prover)
	if l > 0 {
		n += 1 + l + sovStationArg(uint64(l))
	}
	return n
}

func sovStationArg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStationArg(x uint64) (n int) {
	return sovStationArg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StationArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStationArg
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
			return fmt.Errorf("proto: StationArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StationArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStationArg
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
				return ErrInvalidLengthStationArg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStationArg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrackType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStationArg
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
				return ErrInvalidLengthStationArg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStationArg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prover", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStationArg
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
				return ErrInvalidLengthStationArg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStationArg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prover = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStationArg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStationArg
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
func skipStationArg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStationArg
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
					return 0, ErrIntOverflowStationArg
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
					return 0, ErrIntOverflowStationArg
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
				return 0, ErrInvalidLengthStationArg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStationArg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStationArg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStationArg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStationArg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStationArg = fmt.Errorf("proto: unexpected end of group")
)
