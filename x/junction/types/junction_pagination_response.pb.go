// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/junction/junction_pagination_response.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type JunctionPaginationResponse struct {
	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Order  string `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	Total  uint64 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *JunctionPaginationResponse) Reset()         { *m = JunctionPaginationResponse{} }
func (m *JunctionPaginationResponse) String() string { return proto.CompactTextString(m) }
func (*JunctionPaginationResponse) ProtoMessage()    {}
func (*JunctionPaginationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d44250ae4a5de4, []int{0}
}
func (m *JunctionPaginationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JunctionPaginationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JunctionPaginationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JunctionPaginationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunctionPaginationResponse.Merge(m, src)
}
func (m *JunctionPaginationResponse) XXX_Size() int {
	return m.Size()
}
func (m *JunctionPaginationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JunctionPaginationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JunctionPaginationResponse proto.InternalMessageInfo

func (m *JunctionPaginationResponse) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *JunctionPaginationResponse) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *JunctionPaginationResponse) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *JunctionPaginationResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*JunctionPaginationResponse)(nil), "junction.junction.JunctionPaginationResponse")
}

func init() {
	proto.RegisterFile("junction/junction/junction_pagination_response.proto", fileDescriptor_a5d44250ae4a5de4)
}

var fileDescriptor_a5d44250ae4a5de4 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc9, 0x2a, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0xc7, 0x60, 0xc4, 0x17, 0x24, 0xa6, 0x67, 0xe6, 0x25, 0x82, 0x99,
	0x45, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x82,
	0x30, 0x35, 0x7a, 0x30, 0x86, 0x52, 0x19, 0x97, 0x94, 0x17, 0x94, 0x1d, 0x00, 0xd7, 0x17, 0x04,
	0xd5, 0x26, 0x24, 0xc6, 0xc5, 0x96, 0x9f, 0x96, 0x56, 0x9c, 0x5a, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x12, 0x04, 0xe5, 0x09, 0x89, 0x70, 0xb1, 0xe6, 0x64, 0xe6, 0x66, 0x96, 0x48, 0x30, 0x81,
	0x85, 0x21, 0x1c, 0x90, 0x68, 0x7e, 0x51, 0x4a, 0x6a, 0x91, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x84, 0x03, 0x12, 0x2d, 0xc9, 0x2f, 0x49, 0xcc, 0x91, 0x60, 0x81, 0xa8, 0x05, 0x73, 0x9c,
	0x02, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x2c, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x31, 0xb3, 0x28, 0x39, 0x23, 0x31, 0x33, 0xaf,
	0x58, 0x37, 0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x1b, 0xe1, 0xdd, 0x0a, 0x04, 0xb3, 0xa4, 0xb2,
	0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x47, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xef,
	0x18, 0x48, 0x1b, 0x01, 0x00, 0x00,
}

func (m *JunctionPaginationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JunctionPaginationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JunctionPaginationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Total != 0 {
		i = encodeVarintJunctionPaginationResponse(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Order) > 0 {
		i -= len(m.Order)
		copy(dAtA[i:], m.Order)
		i = encodeVarintJunctionPaginationResponse(dAtA, i, uint64(len(m.Order)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Limit != 0 {
		i = encodeVarintJunctionPaginationResponse(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x10
	}
	if m.Offset != 0 {
		i = encodeVarintJunctionPaginationResponse(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintJunctionPaginationResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovJunctionPaginationResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *JunctionPaginationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovJunctionPaginationResponse(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovJunctionPaginationResponse(uint64(m.Limit))
	}
	l = len(m.Order)
	if l > 0 {
		n += 1 + l + sovJunctionPaginationResponse(uint64(l))
	}
	if m.Total != 0 {
		n += 1 + sovJunctionPaginationResponse(uint64(m.Total))
	}
	return n
}

func sovJunctionPaginationResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJunctionPaginationResponse(x uint64) (n int) {
	return sovJunctionPaginationResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *JunctionPaginationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJunctionPaginationResponse
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
			return fmt.Errorf("proto: JunctionPaginationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JunctionPaginationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJunctionPaginationResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJunctionPaginationResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJunctionPaginationResponse
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
				return ErrInvalidLengthJunctionPaginationResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJunctionPaginationResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Order = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJunctionPaginationResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipJunctionPaginationResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJunctionPaginationResponse
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
func skipJunctionPaginationResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJunctionPaginationResponse
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
					return 0, ErrIntOverflowJunctionPaginationResponse
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
					return 0, ErrIntOverflowJunctionPaginationResponse
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
				return 0, ErrInvalidLengthJunctionPaginationResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJunctionPaginationResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJunctionPaginationResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJunctionPaginationResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJunctionPaginationResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJunctionPaginationResponse = fmt.Errorf("proto: unexpected end of group")
)
