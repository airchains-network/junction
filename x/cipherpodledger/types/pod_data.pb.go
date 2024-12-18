// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/cipherpodledger/pod_data.proto

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

type PodData struct {
	AscContractAddress string `protobuf:"bytes,1,opt,name=ascContractAddress,proto3" json:"ascContractAddress,omitempty"`
	PodNumber          uint64 `protobuf:"varint,2,opt,name=podNumber,proto3" json:"podNumber,omitempty"`
	DaBlobId           string `protobuf:"bytes,3,opt,name=daBlobId,proto3" json:"daBlobId,omitempty"`
	SubmittedBy        string `protobuf:"bytes,4,opt,name=submittedBy,proto3" json:"submittedBy,omitempty"`
	Status             string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp          int32  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// proof state update related fields
	ProvingNetwork  string `protobuf:"bytes,7,opt,name=provingNetwork,proto3" json:"provingNetwork,omitempty"`
	ZkFHEProof      []byte `protobuf:"bytes,8,opt,name=zkFHEProof,proto3" json:"zkFHEProof,omitempty"`
	ZkFHEWitness    []byte `protobuf:"bytes,9,opt,name=zkFHEWitness,proto3" json:"zkFHEWitness,omitempty"`
	IsProofVerified bool   `protobuf:"varint,10,opt,name=isProofVerified,proto3" json:"isProofVerified,omitempty"`
}

func (m *PodData) Reset()         { *m = PodData{} }
func (m *PodData) String() string { return proto.CompactTextString(m) }
func (*PodData) ProtoMessage()    {}
func (*PodData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b6b43847b732185, []int{0}
}
func (m *PodData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PodData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PodData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PodData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodData.Merge(m, src)
}
func (m *PodData) XXX_Size() int {
	return m.Size()
}
func (m *PodData) XXX_DiscardUnknown() {
	xxx_messageInfo_PodData.DiscardUnknown(m)
}

var xxx_messageInfo_PodData proto.InternalMessageInfo

func (m *PodData) GetAscContractAddress() string {
	if m != nil {
		return m.AscContractAddress
	}
	return ""
}

func (m *PodData) GetPodNumber() uint64 {
	if m != nil {
		return m.PodNumber
	}
	return 0
}

func (m *PodData) GetDaBlobId() string {
	if m != nil {
		return m.DaBlobId
	}
	return ""
}

func (m *PodData) GetSubmittedBy() string {
	if m != nil {
		return m.SubmittedBy
	}
	return ""
}

func (m *PodData) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PodData) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PodData) GetProvingNetwork() string {
	if m != nil {
		return m.ProvingNetwork
	}
	return ""
}

func (m *PodData) GetZkFHEProof() []byte {
	if m != nil {
		return m.ZkFHEProof
	}
	return nil
}

func (m *PodData) GetZkFHEWitness() []byte {
	if m != nil {
		return m.ZkFHEWitness
	}
	return nil
}

func (m *PodData) GetIsProofVerified() bool {
	if m != nil {
		return m.IsProofVerified
	}
	return false
}

func init() {
	proto.RegisterType((*PodData)(nil), "junction.cipherpodledger.PodData")
}

func init() {
	proto.RegisterFile("junction/cipherpodledger/pod_data.proto", fileDescriptor_4b6b43847b732185)
}

var fileDescriptor_4b6b43847b732185 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4e, 0x23, 0x31,
	0x14, 0x85, 0xe3, 0x6c, 0x7e, 0xbd, 0xd1, 0xae, 0xe4, 0x62, 0x65, 0xad, 0xd0, 0x68, 0x94, 0x02,
	0xa6, 0x61, 0xa6, 0xa0, 0xa6, 0x20, 0xfc, 0x08, 0x9a, 0x28, 0x9a, 0x82, 0x48, 0x34, 0xc8, 0x33,
	0x76, 0x12, 0x93, 0xcc, 0xd8, 0xb2, 0xef, 0x00, 0xe1, 0x29, 0x78, 0x2c, 0xca, 0x94, 0x94, 0x90,
	0xbc, 0x08, 0x8a, 0x81, 0x24, 0x04, 0xca, 0xf3, 0x9d, 0x63, 0xdf, 0xab, 0x7b, 0xf0, 0xde, 0x4d,
	0x91, 0xa7, 0x20, 0x55, 0x1e, 0xa5, 0x52, 0x8f, 0x84, 0xd1, 0x8a, 0x4f, 0x04, 0x1f, 0x0a, 0x13,
	0x69, 0xc5, 0xaf, 0x39, 0x03, 0x16, 0x6a, 0xa3, 0x40, 0x11, 0xfa, 0x19, 0x0c, 0xb7, 0x82, 0xed,
	0xd7, 0x32, 0xae, 0xf7, 0x14, 0x3f, 0x61, 0xc0, 0x48, 0x88, 0x09, 0xb3, 0xe9, 0xb1, 0xca, 0xc1,
	0xb0, 0x14, 0x8e, 0x38, 0x37, 0xc2, 0x5a, 0x8a, 0x7c, 0x14, 0x34, 0xe3, 0x1f, 0x1c, 0xb2, 0x83,
	0x9b, 0x5a, 0xf1, 0x6e, 0x91, 0x25, 0xc2, 0xd0, 0xb2, 0x8f, 0x82, 0x4a, 0xbc, 0x06, 0xe4, 0x3f,
	0x6e, 0x70, 0xd6, 0x99, 0xa8, 0xe4, 0x82, 0xd3, 0x5f, 0xee, 0x8f, 0x95, 0x26, 0x3e, 0xfe, 0x6d,
	0x8b, 0x24, 0x93, 0x00, 0x82, 0x77, 0xa6, 0xb4, 0xe2, 0xec, 0x4d, 0x44, 0xfe, 0xe1, 0x9a, 0x05,
	0x06, 0x85, 0xa5, 0x55, 0x67, 0x7e, 0xa8, 0xe5, 0x4c, 0x90, 0x99, 0xb0, 0xc0, 0x32, 0x4d, 0x6b,
	0x3e, 0x0a, 0xaa, 0xf1, 0x1a, 0x90, 0x5d, 0xfc, 0x47, 0x1b, 0x75, 0x2b, 0xf3, 0x61, 0x57, 0xc0,
	0x9d, 0x32, 0x63, 0x5a, 0x77, 0xaf, 0xb7, 0x28, 0xf1, 0x30, 0x7e, 0x18, 0x9f, 0x9d, 0x9f, 0xf6,
	0x8c, 0x52, 0x03, 0xda, 0xf0, 0x51, 0xd0, 0x8a, 0x37, 0x08, 0x69, 0xe3, 0x96, 0x53, 0x7d, 0x09,
	0xf9, 0xf2, 0x06, 0x4d, 0x97, 0xf8, 0xc2, 0x48, 0x80, 0xff, 0x4a, 0xeb, 0xe2, 0x97, 0xc2, 0xc8,
	0x81, 0x14, 0x9c, 0x62, 0x1f, 0x05, 0x8d, 0x78, 0x1b, 0x77, 0xfa, 0x4f, 0x73, 0x0f, 0xcd, 0xe6,
	0x1e, 0x7a, 0x99, 0x7b, 0xe8, 0x71, 0xe1, 0x95, 0x66, 0x0b, 0xaf, 0xf4, 0xbc, 0xf0, 0x4a, 0x57,
	0x87, 0x43, 0x09, 0xa3, 0x22, 0x09, 0x53, 0x95, 0x45, 0x4c, 0x9a, 0x74, 0xc4, 0x64, 0x6e, 0xf7,
	0xf3, 0xf7, 0x2d, 0xa3, 0x55, 0xbb, 0xf7, 0xdf, 0xfa, 0x85, 0xa9, 0x16, 0x36, 0xa9, 0xb9, 0x76,
	0x0f, 0xde, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0xa7, 0x80, 0xee, 0x08, 0x02, 0x00, 0x00,
}

func (m *PodData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PodData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsProofVerified {
		i--
		if m.IsProofVerified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if len(m.ZkFHEWitness) > 0 {
		i -= len(m.ZkFHEWitness)
		copy(dAtA[i:], m.ZkFHEWitness)
		i = encodeVarintPodData(dAtA, i, uint64(len(m.ZkFHEWitness)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ZkFHEProof) > 0 {
		i -= len(m.ZkFHEProof)
		copy(dAtA[i:], m.ZkFHEProof)
		i = encodeVarintPodData(dAtA, i, uint64(len(m.ZkFHEProof)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ProvingNetwork) > 0 {
		i -= len(m.ProvingNetwork)
		copy(dAtA[i:], m.ProvingNetwork)
		i = encodeVarintPodData(dAtA, i, uint64(len(m.ProvingNetwork)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Timestamp != 0 {
		i = encodeVarintPodData(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintPodData(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SubmittedBy) > 0 {
		i -= len(m.SubmittedBy)
		copy(dAtA[i:], m.SubmittedBy)
		i = encodeVarintPodData(dAtA, i, uint64(len(m.SubmittedBy)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DaBlobId) > 0 {
		i -= len(m.DaBlobId)
		copy(dAtA[i:], m.DaBlobId)
		i = encodeVarintPodData(dAtA, i, uint64(len(m.DaBlobId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PodNumber != 0 {
		i = encodeVarintPodData(dAtA, i, uint64(m.PodNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AscContractAddress) > 0 {
		i -= len(m.AscContractAddress)
		copy(dAtA[i:], m.AscContractAddress)
		i = encodeVarintPodData(dAtA, i, uint64(len(m.AscContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPodData(dAtA []byte, offset int, v uint64) int {
	offset -= sovPodData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PodData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AscContractAddress)
	if l > 0 {
		n += 1 + l + sovPodData(uint64(l))
	}
	if m.PodNumber != 0 {
		n += 1 + sovPodData(uint64(m.PodNumber))
	}
	l = len(m.DaBlobId)
	if l > 0 {
		n += 1 + l + sovPodData(uint64(l))
	}
	l = len(m.SubmittedBy)
	if l > 0 {
		n += 1 + l + sovPodData(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovPodData(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovPodData(uint64(m.Timestamp))
	}
	l = len(m.ProvingNetwork)
	if l > 0 {
		n += 1 + l + sovPodData(uint64(l))
	}
	l = len(m.ZkFHEProof)
	if l > 0 {
		n += 1 + l + sovPodData(uint64(l))
	}
	l = len(m.ZkFHEWitness)
	if l > 0 {
		n += 1 + l + sovPodData(uint64(l))
	}
	if m.IsProofVerified {
		n += 2
	}
	return n
}

func sovPodData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPodData(x uint64) (n int) {
	return sovPodData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PodData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPodData
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
			return fmt.Errorf("proto: PodData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AscContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
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
				return ErrInvalidLengthPodData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPodData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AscContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodNumber", wireType)
			}
			m.PodNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PodNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaBlobId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
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
				return ErrInvalidLengthPodData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPodData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaBlobId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmittedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
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
				return ErrInvalidLengthPodData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPodData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmittedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
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
				return ErrInvalidLengthPodData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPodData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvingNetwork", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
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
				return ErrInvalidLengthPodData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPodData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvingNetwork = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZkFHEProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPodData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPodData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZkFHEProof = append(m.ZkFHEProof[:0], dAtA[iNdEx:postIndex]...)
			if m.ZkFHEProof == nil {
				m.ZkFHEProof = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZkFHEWitness", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPodData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPodData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZkFHEWitness = append(m.ZkFHEWitness[:0], dAtA[iNdEx:postIndex]...)
			if m.ZkFHEWitness == nil {
				m.ZkFHEWitness = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsProofVerified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPodData
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
			m.IsProofVerified = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPodData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPodData
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
func skipPodData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPodData
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
					return 0, ErrIntOverflowPodData
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
					return 0, ErrIntOverflowPodData
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
				return 0, ErrInvalidLengthPodData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPodData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPodData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPodData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPodData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPodData = fmt.Errorf("proto: unexpected end of group")
)