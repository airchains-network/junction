// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/trackgate/ext_sequencer_check.proto

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

type ExtSequencerCheck struct {
	PodNumber          string `protobuf:"bytes,1,opt,name=podNumber,proto3" json:"podNumber,omitempty"`
	ExtTrackStationId  string `protobuf:"bytes,2,opt,name=extTrackStationId,proto3" json:"extTrackStationId,omitempty"`
	VerificationStatus bool   `protobuf:"varint,3,opt,name=verificationStatus,proto3" json:"verificationStatus,omitempty"`
	Namespace          string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *ExtSequencerCheck) Reset()         { *m = ExtSequencerCheck{} }
func (m *ExtSequencerCheck) String() string { return proto.CompactTextString(m) }
func (*ExtSequencerCheck) ProtoMessage()    {}
func (*ExtSequencerCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aff7cb8c2d558ff, []int{0}
}
func (m *ExtSequencerCheck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtSequencerCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtSequencerCheck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtSequencerCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtSequencerCheck.Merge(m, src)
}
func (m *ExtSequencerCheck) XXX_Size() int {
	return m.Size()
}
func (m *ExtSequencerCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtSequencerCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ExtSequencerCheck proto.InternalMessageInfo

func (m *ExtSequencerCheck) GetPodNumber() string {
	if m != nil {
		return m.PodNumber
	}
	return ""
}

func (m *ExtSequencerCheck) GetExtTrackStationId() string {
	if m != nil {
		return m.ExtTrackStationId
	}
	return ""
}

func (m *ExtSequencerCheck) GetVerificationStatus() bool {
	if m != nil {
		return m.VerificationStatus
	}
	return false
}

func (m *ExtSequencerCheck) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtSequencerCheck)(nil), "junction.trackgate.ExtSequencerCheck")
}

func init() {
	proto.RegisterFile("junction/trackgate/ext_sequencer_check.proto", fileDescriptor_5aff7cb8c2d558ff)
}

var fileDescriptor_5aff7cb8c2d558ff = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x40, 0x88, 0x66, 0xab, 0xa7, 0x0c, 0xc8, 0xaa, 0x98, 0x3a, 0x94, 0x64, 0x60,
	0x60, 0x07, 0x31, 0xb0, 0x20, 0xd1, 0x32, 0xb1, 0x54, 0xce, 0xf5, 0x68, 0x4c, 0x14, 0x3b, 0xd8,
	0x67, 0x08, 0x6f, 0xc1, 0x8b, 0xf0, 0x1e, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x22, 0xc8, 0x41, 0x0d,
	0x48, 0x74, 0xfd, 0xff, 0xef, 0xee, 0x97, 0xbe, 0x78, 0xf6, 0xe8, 0x35, 0x90, 0x32, 0x3a, 0x23,
	0x2b, 0xa1, 0x5c, 0x4b, 0xc2, 0x0c, 0x1b, 0x5a, 0x3a, 0x7c, 0xf2, 0xa8, 0x01, 0xed, 0x12, 0x0a,
	0x84, 0x32, 0xad, 0xad, 0x21, 0xc3, 0xf9, 0x96, 0x4e, 0x07, 0xfa, 0xe4, 0x9d, 0xc5, 0xe3, 0xab,
	0x86, 0x16, 0xdb, 0x83, 0xcb, 0xc0, 0xf3, 0xe3, 0x78, 0x54, 0x9b, 0xd5, 0x8d, 0xaf, 0x72, 0xb4,
	0x09, 0x9b, 0xb0, 0xe9, 0x68, 0xfe, 0x1b, 0xf0, 0x59, 0x3c, 0xc6, 0x86, 0xee, 0xc2, 0x8f, 0x05,
	0xc9, 0xf0, 0xf0, 0x7a, 0x95, 0xec, 0xf5, 0xd4, 0xff, 0x82, 0xa7, 0x31, 0x7f, 0x46, 0xab, 0x1e,
	0x14, 0xf4, 0x49, 0x28, 0xbc, 0x4b, 0xf6, 0x27, 0x6c, 0x7a, 0x34, 0xdf, 0xd1, 0x84, 0x6d, 0x2d,
	0x2b, 0x74, 0xb5, 0x04, 0x4c, 0x0e, 0x7e, 0xb6, 0x87, 0xe0, 0xe2, 0xf6, 0xa3, 0x15, 0x6c, 0xd3,
	0x0a, 0xf6, 0xd5, 0x0a, 0xf6, 0xd6, 0x89, 0x68, 0xd3, 0x89, 0xe8, 0xb3, 0x13, 0xd1, 0xfd, 0xf9,
	0x5a, 0x51, 0xe1, 0xf3, 0x14, 0x4c, 0x95, 0x49, 0x65, 0xa1, 0x90, 0x4a, 0xbb, 0x53, 0x8d, 0xf4,
	0x62, 0x6c, 0x99, 0x0d, 0xa2, 0x9a, 0x3f, 0xaa, 0xe8, 0xb5, 0x46, 0x97, 0x1f, 0xf6, 0x76, 0xce,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x53, 0xb0, 0x93, 0xe3, 0x4d, 0x01, 0x00, 0x00,
}

func (m *ExtSequencerCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtSequencerCheck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtSequencerCheck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintExtSequencerCheck(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x22
	}
	if m.VerificationStatus {
		i--
		if m.VerificationStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ExtTrackStationId) > 0 {
		i -= len(m.ExtTrackStationId)
		copy(dAtA[i:], m.ExtTrackStationId)
		i = encodeVarintExtSequencerCheck(dAtA, i, uint64(len(m.ExtTrackStationId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PodNumber) > 0 {
		i -= len(m.PodNumber)
		copy(dAtA[i:], m.PodNumber)
		i = encodeVarintExtSequencerCheck(dAtA, i, uint64(len(m.PodNumber)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtSequencerCheck(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtSequencerCheck(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtSequencerCheck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PodNumber)
	if l > 0 {
		n += 1 + l + sovExtSequencerCheck(uint64(l))
	}
	l = len(m.ExtTrackStationId)
	if l > 0 {
		n += 1 + l + sovExtSequencerCheck(uint64(l))
	}
	if m.VerificationStatus {
		n += 2
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovExtSequencerCheck(uint64(l))
	}
	return n
}

func sovExtSequencerCheck(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtSequencerCheck(x uint64) (n int) {
	return sovExtSequencerCheck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtSequencerCheck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtSequencerCheck
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
			return fmt.Errorf("proto: ExtSequencerCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtSequencerCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtSequencerCheck
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
				return ErrInvalidLengthExtSequencerCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtSequencerCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtTrackStationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtSequencerCheck
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
				return ErrInvalidLengthExtSequencerCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtSequencerCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtTrackStationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtSequencerCheck
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
			m.VerificationStatus = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtSequencerCheck
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
				return ErrInvalidLengthExtSequencerCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtSequencerCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExtSequencerCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtSequencerCheck
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
func skipExtSequencerCheck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtSequencerCheck
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
					return 0, ErrIntOverflowExtSequencerCheck
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
					return 0, ErrIntOverflowExtSequencerCheck
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
				return 0, ErrInvalidLengthExtSequencerCheck
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtSequencerCheck
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtSequencerCheck
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtSequencerCheck        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtSequencerCheck          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtSequencerCheck = fmt.Errorf("proto: unexpected end of group")
)
