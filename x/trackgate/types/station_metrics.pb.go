// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/trackgate/station_metrics.proto

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

type StationMetrics struct {
	TotalPodCount         uint64 `protobuf:"varint,1,opt,name=totalPodCount,proto3" json:"totalPodCount,omitempty"`
	TotalSchemaCount      uint64 `protobuf:"varint,2,opt,name=totalSchemaCount,proto3" json:"totalSchemaCount,omitempty"`
	TotalMigrationCount   uint64 `protobuf:"varint,3,opt,name=totalMigrationCount,proto3" json:"totalMigrationCount,omitempty"`
	TotalVerifiedPodCount uint64 `protobuf:"varint,4,opt,name=totalVerifiedPodCount,proto3" json:"totalVerifiedPodCount,omitempty"`
}

func (m *StationMetrics) Reset()         { *m = StationMetrics{} }
func (m *StationMetrics) String() string { return proto.CompactTextString(m) }
func (*StationMetrics) ProtoMessage()    {}
func (*StationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_eac1dbaf9d85c719, []int{0}
}
func (m *StationMetrics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StationMetrics.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StationMetrics.Merge(m, src)
}
func (m *StationMetrics) XXX_Size() int {
	return m.Size()
}
func (m *StationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_StationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_StationMetrics proto.InternalMessageInfo

func (m *StationMetrics) GetTotalPodCount() uint64 {
	if m != nil {
		return m.TotalPodCount
	}
	return 0
}

func (m *StationMetrics) GetTotalSchemaCount() uint64 {
	if m != nil {
		return m.TotalSchemaCount
	}
	return 0
}

func (m *StationMetrics) GetTotalMigrationCount() uint64 {
	if m != nil {
		return m.TotalMigrationCount
	}
	return 0
}

func (m *StationMetrics) GetTotalVerifiedPodCount() uint64 {
	if m != nil {
		return m.TotalVerifiedPodCount
	}
	return 0
}

func init() {
	proto.RegisterType((*StationMetrics)(nil), "junction.trackgate.StationMetrics")
}

func init() {
	proto.RegisterFile("junction/trackgate/station_metrics.proto", fileDescriptor_eac1dbaf9d85c719)
}

var fileDescriptor_eac1dbaf9d85c719 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0x2a, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x29, 0x4a, 0x4c, 0xce, 0x4e, 0x4f, 0x2c, 0x49, 0xd5, 0x2f,
	0x2e, 0x49, 0x04, 0x89, 0xc4, 0xe7, 0xa6, 0x96, 0x14, 0x65, 0x26, 0x17, 0xeb, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0x09, 0xc1, 0x54, 0xea, 0xc1, 0x55, 0x2a, 0x9d, 0x62, 0xe4, 0xe2, 0x0b, 0x86,
	0xa8, 0xf6, 0x85, 0x28, 0x16, 0x52, 0xe1, 0xe2, 0x2d, 0xc9, 0x2f, 0x49, 0xcc, 0x09, 0xc8, 0x4f,
	0x71, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x42, 0x15, 0x14, 0xd2,
	0xe2, 0x12, 0x00, 0x0b, 0x04, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0x42, 0x14, 0x32, 0x81, 0x15, 0x62,
	0x88, 0x0b, 0x19, 0x70, 0x09, 0x83, 0xc5, 0x7c, 0x33, 0xd3, 0x8b, 0xc0, 0x56, 0x41, 0x94, 0x33,
	0x83, 0x95, 0x63, 0x93, 0x12, 0x32, 0xe1, 0x12, 0x05, 0x0b, 0x87, 0xa5, 0x16, 0x65, 0xa6, 0x65,
	0xa6, 0xa6, 0xc0, 0xdd, 0xc2, 0x02, 0xd6, 0x83, 0x5d, 0xd2, 0x29, 0xf0, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xcc, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x13, 0x33, 0x8b, 0x92, 0x33, 0x12, 0x33, 0xf3, 0x8a, 0x75, 0xf3, 0x52, 0x4b, 0xca,
	0xf3, 0x8b, 0xb2, 0xf5, 0xe1, 0x21, 0x58, 0x81, 0x14, 0x86, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49,
	0x6c, 0xe0, 0xa0, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x15, 0x76, 0xc6, 0x56, 0x66, 0x01,
	0x00, 0x00,
}

func (m *StationMetrics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StationMetrics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StationMetrics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalVerifiedPodCount != 0 {
		i = encodeVarintStationMetrics(dAtA, i, uint64(m.TotalVerifiedPodCount))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalMigrationCount != 0 {
		i = encodeVarintStationMetrics(dAtA, i, uint64(m.TotalMigrationCount))
		i--
		dAtA[i] = 0x18
	}
	if m.TotalSchemaCount != 0 {
		i = encodeVarintStationMetrics(dAtA, i, uint64(m.TotalSchemaCount))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalPodCount != 0 {
		i = encodeVarintStationMetrics(dAtA, i, uint64(m.TotalPodCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStationMetrics(dAtA []byte, offset int, v uint64) int {
	offset -= sovStationMetrics(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StationMetrics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalPodCount != 0 {
		n += 1 + sovStationMetrics(uint64(m.TotalPodCount))
	}
	if m.TotalSchemaCount != 0 {
		n += 1 + sovStationMetrics(uint64(m.TotalSchemaCount))
	}
	if m.TotalMigrationCount != 0 {
		n += 1 + sovStationMetrics(uint64(m.TotalMigrationCount))
	}
	if m.TotalVerifiedPodCount != 0 {
		n += 1 + sovStationMetrics(uint64(m.TotalVerifiedPodCount))
	}
	return n
}

func sovStationMetrics(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStationMetrics(x uint64) (n int) {
	return sovStationMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StationMetrics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStationMetrics
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
			return fmt.Errorf("proto: StationMetrics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StationMetrics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalPodCount", wireType)
			}
			m.TotalPodCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStationMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalPodCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSchemaCount", wireType)
			}
			m.TotalSchemaCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStationMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalSchemaCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalMigrationCount", wireType)
			}
			m.TotalMigrationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStationMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalMigrationCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVerifiedPodCount", wireType)
			}
			m.TotalVerifiedPodCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStationMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVerifiedPodCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStationMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStationMetrics
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
func skipStationMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStationMetrics
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
					return 0, ErrIntOverflowStationMetrics
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
					return 0, ErrIntOverflowStationMetrics
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
				return 0, ErrInvalidLengthStationMetrics
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStationMetrics
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStationMetrics
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStationMetrics        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStationMetrics          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStationMetrics = fmt.Errorf("proto: unexpected end of group")
)
