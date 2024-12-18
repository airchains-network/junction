// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/trackgate/da_blob_data.proto

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

type DaBlobData struct {
	DaType       string `protobuf:"bytes,1,opt,name=daType,proto3" json:"daType,omitempty"`
	Height       uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Blob         []byte `protobuf:"bytes,3,opt,name=blob,proto3" json:"blob,omitempty"`
	StationRefId string `protobuf:"bytes,4,opt,name=stationRefId,proto3" json:"stationRefId,omitempty"`
}

func (m *DaBlobData) Reset()         { *m = DaBlobData{} }
func (m *DaBlobData) String() string { return proto.CompactTextString(m) }
func (*DaBlobData) ProtoMessage()    {}
func (*DaBlobData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e70f608921a7a48f, []int{0}
}
func (m *DaBlobData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DaBlobData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DaBlobData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DaBlobData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DaBlobData.Merge(m, src)
}
func (m *DaBlobData) XXX_Size() int {
	return m.Size()
}
func (m *DaBlobData) XXX_DiscardUnknown() {
	xxx_messageInfo_DaBlobData.DiscardUnknown(m)
}

var xxx_messageInfo_DaBlobData proto.InternalMessageInfo

func (m *DaBlobData) GetDaType() string {
	if m != nil {
		return m.DaType
	}
	return ""
}

func (m *DaBlobData) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *DaBlobData) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

func (m *DaBlobData) GetStationRefId() string {
	if m != nil {
		return m.StationRefId
	}
	return ""
}

func init() {
	proto.RegisterType((*DaBlobData)(nil), "junction.trackgate.DaBlobData")
}

func init() {
	proto.RegisterFile("junction/trackgate/da_blob_data.proto", fileDescriptor_e70f608921a7a48f)
}

var fileDescriptor_e70f608921a7a48f = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0x2a, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x29, 0x4a, 0x4c, 0xce, 0x4e, 0x4f, 0x2c, 0x49, 0xd5, 0x4f,
	0x49, 0x8c, 0x4f, 0xca, 0xc9, 0x4f, 0x8a, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x82, 0x29, 0xd3, 0x83, 0x2b, 0x53, 0x2a, 0xe1, 0xe2, 0x72, 0x49, 0x74, 0xca,
	0xc9, 0x4f, 0x72, 0x49, 0x2c, 0x49, 0x14, 0x12, 0xe3, 0x62, 0x4b, 0x49, 0x0c, 0xa9, 0x2c, 0x48,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x40, 0xe2, 0x19, 0xa9, 0x99, 0xe9, 0x19,
	0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x90, 0x10, 0x17, 0x0b, 0xc8, 0x12,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0x48, 0x89, 0x8b, 0xa7, 0xb8, 0x24, 0x11,
	0x64, 0x4d, 0x50, 0x6a, 0x9a, 0x67, 0x8a, 0x04, 0x0b, 0xd8, 0x24, 0x14, 0x31, 0xa7, 0xc0, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcc, 0x2c, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0x2b, 0xd6, 0xcd,
	0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x87, 0xfb, 0xb3, 0x02, 0xc9, 0xa7, 0x25, 0x95, 0x05,
	0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x3f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x06, 0xb0, 0xec,
	0x8f, 0x0c, 0x01, 0x00, 0x00,
}

func (m *DaBlobData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DaBlobData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DaBlobData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StationRefId) > 0 {
		i -= len(m.StationRefId)
		copy(dAtA[i:], m.StationRefId)
		i = encodeVarintDaBlobData(dAtA, i, uint64(len(m.StationRefId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Blob) > 0 {
		i -= len(m.Blob)
		copy(dAtA[i:], m.Blob)
		i = encodeVarintDaBlobData(dAtA, i, uint64(len(m.Blob)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Height != 0 {
		i = encodeVarintDaBlobData(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DaType) > 0 {
		i -= len(m.DaType)
		copy(dAtA[i:], m.DaType)
		i = encodeVarintDaBlobData(dAtA, i, uint64(len(m.DaType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDaBlobData(dAtA []byte, offset int, v uint64) int {
	offset -= sovDaBlobData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DaBlobData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DaType)
	if l > 0 {
		n += 1 + l + sovDaBlobData(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovDaBlobData(uint64(m.Height))
	}
	l = len(m.Blob)
	if l > 0 {
		n += 1 + l + sovDaBlobData(uint64(l))
	}
	l = len(m.StationRefId)
	if l > 0 {
		n += 1 + l + sovDaBlobData(uint64(l))
	}
	return n
}

func sovDaBlobData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDaBlobData(x uint64) (n int) {
	return sovDaBlobData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DaBlobData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDaBlobData
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
			return fmt.Errorf("proto: DaBlobData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DaBlobData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDaBlobData
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
				return ErrInvalidLengthDaBlobData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDaBlobData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDaBlobData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDaBlobData
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
				return ErrInvalidLengthDaBlobData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDaBlobData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blob = append(m.Blob[:0], dAtA[iNdEx:postIndex]...)
			if m.Blob == nil {
				m.Blob = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationRefId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDaBlobData
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
				return ErrInvalidLengthDaBlobData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDaBlobData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationRefId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDaBlobData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDaBlobData
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
func skipDaBlobData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDaBlobData
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
					return 0, ErrIntOverflowDaBlobData
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
					return 0, ErrIntOverflowDaBlobData
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
				return 0, ErrInvalidLengthDaBlobData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDaBlobData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDaBlobData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDaBlobData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDaBlobData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDaBlobData = fmt.Errorf("proto: unexpected end of group")
)
