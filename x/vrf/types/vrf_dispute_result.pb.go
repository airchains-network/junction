// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/vrf/vrf_dispute_result.proto

package types

import (
	encoding_binary "encoding/binary"
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

type VrfDisputeResult struct {
	Votes               []bool   `protobuf:"varint,1,rep,packed,name=votes,proto3" json:"votes,omitempty"`
	Members             []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	ConsentVote         uint32   `protobuf:"varint,3,opt,name=consentVote,proto3" json:"consentVote,omitempty"`
	DissentVote         uint32   `protobuf:"varint,4,opt,name=dissentVote,proto3" json:"dissentVote,omitempty"`
	AgreementPercentage float32  `protobuf:"fixed32,5,opt,name=agreementPercentage,proto3" json:"agreementPercentage,omitempty"`
	Result              bool     `protobuf:"varint,6,opt,name=result,proto3" json:"result,omitempty"`
	Message             string   `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *VrfDisputeResult) Reset()         { *m = VrfDisputeResult{} }
func (m *VrfDisputeResult) String() string { return proto.CompactTextString(m) }
func (*VrfDisputeResult) ProtoMessage()    {}
func (*VrfDisputeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac3ebb018761df3a, []int{0}
}
func (m *VrfDisputeResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VrfDisputeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VrfDisputeResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VrfDisputeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrfDisputeResult.Merge(m, src)
}
func (m *VrfDisputeResult) XXX_Size() int {
	return m.Size()
}
func (m *VrfDisputeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_VrfDisputeResult.DiscardUnknown(m)
}

var xxx_messageInfo_VrfDisputeResult proto.InternalMessageInfo

func (m *VrfDisputeResult) GetVotes() []bool {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *VrfDisputeResult) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *VrfDisputeResult) GetConsentVote() uint32 {
	if m != nil {
		return m.ConsentVote
	}
	return 0
}

func (m *VrfDisputeResult) GetDissentVote() uint32 {
	if m != nil {
		return m.DissentVote
	}
	return 0
}

func (m *VrfDisputeResult) GetAgreementPercentage() float32 {
	if m != nil {
		return m.AgreementPercentage
	}
	return 0
}

func (m *VrfDisputeResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *VrfDisputeResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*VrfDisputeResult)(nil), "junction.vrf.VrfDisputeResult")
}

func init() {
	proto.RegisterFile("junction/vrf/vrf_dispute_result.proto", fileDescriptor_ac3ebb018761df3a)
}

var fileDescriptor_ac3ebb018761df3a = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xeb, 0x96, 0xfe, 0x60, 0x40, 0x42, 0x06, 0x21, 0x4f, 0x96, 0x85, 0x84, 0x94, 0x85,
	0x16, 0xc4, 0x1b, 0x20, 0x36, 0x16, 0xe4, 0xa1, 0x03, 0x4b, 0x95, 0xa6, 0xb7, 0xa9, 0x81, 0xd8,
	0x91, 0x7d, 0x13, 0xe0, 0x2d, 0x78, 0x2c, 0xc6, 0x8e, 0x8c, 0x28, 0xd9, 0x78, 0x0a, 0x94, 0x98,
	0x56, 0x1d, 0x18, 0x3c, 0x7c, 0xc7, 0x9f, 0xa5, 0xe3, 0x43, 0x2f, 0x9e, 0x0a, 0x93, 0xa0, 0xb6,
	0x66, 0x52, 0xba, 0x65, 0x73, 0x66, 0x0b, 0xed, 0xf3, 0x02, 0x61, 0xe6, 0xc0, 0x17, 0x2f, 0x38,
	0xce, 0x9d, 0x45, 0xcb, 0x0e, 0x37, 0xda, 0xb8, 0x74, 0xcb, 0xf3, 0x1f, 0x42, 0x8f, 0xa7, 0x6e,
	0x79, 0x17, 0x4c, 0xd5, 0x8a, 0xec, 0x94, 0xf6, 0x4b, 0x8b, 0xe0, 0x39, 0x91, 0xbd, 0x68, 0xa4,
	0x02, 0x30, 0x4e, 0x87, 0x19, 0x64, 0x73, 0x70, 0x9e, 0x77, 0x65, 0x2f, 0xda, 0x57, 0x1b, 0x64,
	0x92, 0x1e, 0x24, 0xd6, 0x78, 0x30, 0x38, 0xb5, 0x08, 0xbc, 0x27, 0x49, 0x74, 0xa4, 0x76, 0xa3,
	0xc6, 0x58, 0x68, 0xbf, 0x35, 0xf6, 0x82, 0xb1, 0x13, 0xb1, 0x2b, 0x7a, 0x12, 0xa7, 0x0e, 0x20,
	0x03, 0x83, 0x0f, 0xe0, 0x12, 0x30, 0x18, 0xa7, 0xc0, 0xfb, 0x92, 0x44, 0x5d, 0xf5, 0xdf, 0x15,
	0x3b, 0xa3, 0x83, 0xf0, 0x31, 0x3e, 0x90, 0x24, 0x1a, 0xa9, 0x3f, 0x0a, 0x3d, 0xbd, 0x6f, 0x5e,
	0x0f, 0x25, 0x09, 0x3d, 0x5b, 0xbc, 0xbd, 0xff, 0xac, 0x04, 0x59, 0x57, 0x82, 0x7c, 0x57, 0x82,
	0x7c, 0xd4, 0xa2, 0xb3, 0xae, 0x45, 0xe7, 0xab, 0x16, 0x9d, 0xc7, 0xeb, 0x54, 0xe3, 0xaa, 0x98,
	0x8f, 0x13, 0x9b, 0x4d, 0x62, 0xed, 0x92, 0x55, 0xac, 0x8d, 0xbf, 0x34, 0x80, 0xaf, 0xd6, 0x3d,
	0x4f, 0xb6, 0xc3, 0xbe, 0xb5, 0xd3, 0xe2, 0x7b, 0x0e, 0x7e, 0x3e, 0x68, 0xe7, 0xbc, 0xf9, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0x56, 0xba, 0x84, 0x77, 0x01, 0x00, 0x00,
}

func (m *VrfDisputeResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VrfDisputeResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VrfDisputeResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintVrfDisputeResult(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Result {
		i--
		if m.Result {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.AgreementPercentage != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.AgreementPercentage))))
		i--
		dAtA[i] = 0x2d
	}
	if m.DissentVote != 0 {
		i = encodeVarintVrfDisputeResult(dAtA, i, uint64(m.DissentVote))
		i--
		dAtA[i] = 0x20
	}
	if m.ConsentVote != 0 {
		i = encodeVarintVrfDisputeResult(dAtA, i, uint64(m.ConsentVote))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintVrfDisputeResult(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.Votes[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintVrfDisputeResult(dAtA, i, uint64(len(m.Votes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVrfDisputeResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovVrfDisputeResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VrfDisputeResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Votes) > 0 {
		n += 1 + sovVrfDisputeResult(uint64(len(m.Votes))) + len(m.Votes)*1
	}
	if len(m.Members) > 0 {
		for _, s := range m.Members {
			l = len(s)
			n += 1 + l + sovVrfDisputeResult(uint64(l))
		}
	}
	if m.ConsentVote != 0 {
		n += 1 + sovVrfDisputeResult(uint64(m.ConsentVote))
	}
	if m.DissentVote != 0 {
		n += 1 + sovVrfDisputeResult(uint64(m.DissentVote))
	}
	if m.AgreementPercentage != 0 {
		n += 5
	}
	if m.Result {
		n += 2
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovVrfDisputeResult(uint64(l))
	}
	return n
}

func sovVrfDisputeResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVrfDisputeResult(x uint64) (n int) {
	return sovVrfDisputeResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VrfDisputeResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVrfDisputeResult
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
			return fmt.Errorf("proto: VrfDisputeResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VrfDisputeResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVrfDisputeResult
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
				m.Votes = append(m.Votes, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVrfDisputeResult
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthVrfDisputeResult
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthVrfDisputeResult
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.Votes) == 0 {
					m.Votes = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVrfDisputeResult
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
					m.Votes = append(m.Votes, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfDisputeResult
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
				return ErrInvalidLengthVrfDisputeResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfDisputeResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsentVote", wireType)
			}
			m.ConsentVote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfDisputeResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsentVote |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DissentVote", wireType)
			}
			m.DissentVote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfDisputeResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DissentVote |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgreementPercentage", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.AgreementPercentage = float32(math.Float32frombits(v))
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfDisputeResult
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
			m.Result = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVrfDisputeResult
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
				return ErrInvalidLengthVrfDisputeResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVrfDisputeResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVrfDisputeResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVrfDisputeResult
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
func skipVrfDisputeResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVrfDisputeResult
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
					return 0, ErrIntOverflowVrfDisputeResult
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
					return 0, ErrIntOverflowVrfDisputeResult
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
				return 0, ErrInvalidLengthVrfDisputeResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVrfDisputeResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVrfDisputeResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVrfDisputeResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVrfDisputeResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVrfDisputeResult = fmt.Errorf("proto: unexpected end of group")
)
