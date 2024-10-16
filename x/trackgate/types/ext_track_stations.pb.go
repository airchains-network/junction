// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/trackgate/ext_track_stations.proto

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

type ExtTrackStations struct {
	Operators                 []string `protobuf:"bytes,1,rep,name=operators,proto3" json:"operators,omitempty"`
	LatestPod                 uint64   `protobuf:"varint,2,opt,name=latestPod,proto3" json:"latestPod,omitempty"`
	LatestAcknowledgementHash string   `protobuf:"bytes,3,opt,name=latestAcknowledgementHash,proto3" json:"latestAcknowledgementHash,omitempty"`
	Name                      string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Id                        string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	StationType               string   `protobuf:"bytes,6,opt,name=stationType,proto3" json:"stationType,omitempty"`
	FheEnabled                bool     `protobuf:"varint,7,opt,name=fheEnabled,proto3" json:"fheEnabled,omitempty"`
	SequencerDetails          []byte   `protobuf:"bytes,8,opt,name=sequencerDetails,proto3" json:"sequencerDetails,omitempty"`
	DaDetails                 []byte   `protobuf:"bytes,9,opt,name=daDetails,proto3" json:"daDetails,omitempty"`
	ProverDetails             []byte   `protobuf:"bytes,10,opt,name=proverDetails,proto3" json:"proverDetails,omitempty"`
	StationSchemaKey          string   `protobuf:"bytes,11,opt,name=stationSchemaKey,proto3" json:"stationSchemaKey,omitempty"`
	Creator                   string   `protobuf:"bytes,12,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *ExtTrackStations) Reset()         { *m = ExtTrackStations{} }
func (m *ExtTrackStations) String() string { return proto.CompactTextString(m) }
func (*ExtTrackStations) ProtoMessage()    {}
func (*ExtTrackStations) Descriptor() ([]byte, []int) {
	return fileDescriptor_33d9c69d09280ebc, []int{0}
}
func (m *ExtTrackStations) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtTrackStations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtTrackStations.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtTrackStations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtTrackStations.Merge(m, src)
}
func (m *ExtTrackStations) XXX_Size() int {
	return m.Size()
}
func (m *ExtTrackStations) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtTrackStations.DiscardUnknown(m)
}

var xxx_messageInfo_ExtTrackStations proto.InternalMessageInfo

func (m *ExtTrackStations) GetOperators() []string {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (m *ExtTrackStations) GetLatestPod() uint64 {
	if m != nil {
		return m.LatestPod
	}
	return 0
}

func (m *ExtTrackStations) GetLatestAcknowledgementHash() string {
	if m != nil {
		return m.LatestAcknowledgementHash
	}
	return ""
}

func (m *ExtTrackStations) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExtTrackStations) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExtTrackStations) GetStationType() string {
	if m != nil {
		return m.StationType
	}
	return ""
}

func (m *ExtTrackStations) GetFheEnabled() bool {
	if m != nil {
		return m.FheEnabled
	}
	return false
}

func (m *ExtTrackStations) GetSequencerDetails() []byte {
	if m != nil {
		return m.SequencerDetails
	}
	return nil
}

func (m *ExtTrackStations) GetDaDetails() []byte {
	if m != nil {
		return m.DaDetails
	}
	return nil
}

func (m *ExtTrackStations) GetProverDetails() []byte {
	if m != nil {
		return m.ProverDetails
	}
	return nil
}

func (m *ExtTrackStations) GetStationSchemaKey() string {
	if m != nil {
		return m.StationSchemaKey
	}
	return ""
}

func (m *ExtTrackStations) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtTrackStations)(nil), "junction.trackgate.ExtTrackStations")
}

func init() {
	proto.RegisterFile("junction/trackgate/ext_track_stations.proto", fileDescriptor_33d9c69d09280ebc)
}

var fileDescriptor_33d9c69d09280ebc = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6f, 0xda, 0x30,
	0x18, 0x87, 0x31, 0x30, 0x20, 0x86, 0x4d, 0xc8, 0x27, 0x4f, 0x9a, 0xa2, 0x68, 0xda, 0x21, 0xda,
	0x34, 0x72, 0xd8, 0x61, 0x97, 0x5d, 0x56, 0x15, 0xa9, 0x52, 0x2f, 0x6d, 0xe0, 0xd4, 0x0b, 0x32,
	0xce, 0x5b, 0x92, 0x92, 0xd8, 0xa9, 0x6d, 0x0a, 0x7c, 0x8b, 0x7e, 0x8c, 0x7e, 0x94, 0x1e, 0x39,
	0xf6, 0x58, 0xc1, 0x17, 0xa9, 0x62, 0xfe, 0x84, 0x0a, 0xf5, 0xe6, 0xf7, 0x79, 0x1f, 0xd9, 0xef,
	0x6b, 0xfd, 0xf0, 0xaf, 0xbb, 0x99, 0xe0, 0x26, 0x91, 0x22, 0x30, 0x8a, 0xf1, 0xe9, 0x84, 0x19,
	0x08, 0x60, 0x61, 0x46, 0xb6, 0x1a, 0x69, 0xc3, 0x8a, 0x9e, 0xee, 0xe5, 0x4a, 0x1a, 0x49, 0xc8,
	0x5e, 0xee, 0x1d, 0xe4, 0xef, 0x4f, 0x35, 0xdc, 0xed, 0x2f, 0xcc, 0xb0, 0x00, 0x83, 0x9d, 0x4e,
	0xbe, 0x61, 0x47, 0xe6, 0xa0, 0x98, 0x91, 0x4a, 0x53, 0xe4, 0xd5, 0x7c, 0x27, 0x2c, 0x41, 0xd1,
	0x4d, 0x99, 0x01, 0x6d, 0xae, 0x64, 0x44, 0xab, 0x1e, 0xf2, 0xeb, 0x61, 0x09, 0xc8, 0x3f, 0xfc,
	0x75, 0x5b, 0xfc, 0xe7, 0x53, 0x21, 0xe7, 0x29, 0x44, 0x13, 0xc8, 0x40, 0x98, 0x0b, 0xa6, 0x63,
	0x5a, 0xf3, 0x90, 0xef, 0x84, 0x1f, 0x0b, 0x84, 0xe0, 0xba, 0x60, 0x19, 0xd0, 0xba, 0x15, 0xed,
	0x99, 0x7c, 0xc1, 0xd5, 0x24, 0xa2, 0x9f, 0x2c, 0xa9, 0x26, 0x11, 0xf1, 0x70, 0x7b, 0xb7, 0xd8,
	0x70, 0x99, 0x03, 0x6d, 0xd8, 0xc6, 0x31, 0x22, 0x2e, 0xc6, 0xb7, 0x31, 0xf4, 0x05, 0x1b, 0xa7,
	0x10, 0xd1, 0xa6, 0x87, 0xfc, 0x56, 0x78, 0x44, 0xc8, 0x4f, 0xdc, 0xd5, 0x70, 0x3f, 0x03, 0xc1,
	0x41, 0x9d, 0x83, 0x61, 0x49, 0xaa, 0x69, 0xcb, 0x43, 0x7e, 0x27, 0x3c, 0xe1, 0xc5, 0xb6, 0x11,
	0xdb, 0x4b, 0x8e, 0x95, 0x4a, 0x40, 0x7e, 0xe0, 0xcf, 0xb9, 0x92, 0x0f, 0xe5, 0x35, 0xd8, 0x1a,
	0xef, 0xa1, 0x7d, 0x6f, 0x3b, 0xde, 0x80, 0xc7, 0x90, 0xb1, 0x4b, 0x58, 0xd2, 0xb6, 0x1d, 0xfb,
	0x84, 0x13, 0x8a, 0x9b, 0x5c, 0x41, 0xf1, 0xd3, 0xb4, 0x63, 0x95, 0x7d, 0x79, 0x76, 0xfd, 0xbc,
	0x76, 0xd1, 0x6a, 0xed, 0xa2, 0xd7, 0xb5, 0x8b, 0x1e, 0x37, 0x6e, 0x65, 0xb5, 0x71, 0x2b, 0x2f,
	0x1b, 0xb7, 0x72, 0xf3, 0x77, 0x92, 0x98, 0x78, 0x36, 0xee, 0x71, 0x99, 0x05, 0x2c, 0x51, 0x3c,
	0x66, 0x89, 0xd0, 0xbf, 0x05, 0x98, 0xb9, 0x54, 0xd3, 0xe0, 0x10, 0x91, 0xc5, 0x51, 0x48, 0xcc,
	0x32, 0x07, 0x3d, 0x6e, 0xd8, 0x60, 0xfc, 0x79, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x00, 0x86,
	0x74, 0x47, 0x02, 0x00, 0x00,
}

func (m *ExtTrackStations) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtTrackStations) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtTrackStations) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.StationSchemaKey) > 0 {
		i -= len(m.StationSchemaKey)
		copy(dAtA[i:], m.StationSchemaKey)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.StationSchemaKey)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ProverDetails) > 0 {
		i -= len(m.ProverDetails)
		copy(dAtA[i:], m.ProverDetails)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.ProverDetails)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.DaDetails) > 0 {
		i -= len(m.DaDetails)
		copy(dAtA[i:], m.DaDetails)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.DaDetails)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.SequencerDetails) > 0 {
		i -= len(m.SequencerDetails)
		copy(dAtA[i:], m.SequencerDetails)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.SequencerDetails)))
		i--
		dAtA[i] = 0x42
	}
	if m.FheEnabled {
		i--
		if m.FheEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.StationType) > 0 {
		i -= len(m.StationType)
		copy(dAtA[i:], m.StationType)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.StationType)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LatestAcknowledgementHash) > 0 {
		i -= len(m.LatestAcknowledgementHash)
		copy(dAtA[i:], m.LatestAcknowledgementHash)
		i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.LatestAcknowledgementHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.LatestPod != 0 {
		i = encodeVarintExtTrackStations(dAtA, i, uint64(m.LatestPod))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Operators) > 0 {
		for iNdEx := len(m.Operators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Operators[iNdEx])
			copy(dAtA[i:], m.Operators[iNdEx])
			i = encodeVarintExtTrackStations(dAtA, i, uint64(len(m.Operators[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtTrackStations(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtTrackStations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtTrackStations) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Operators) > 0 {
		for _, s := range m.Operators {
			l = len(s)
			n += 1 + l + sovExtTrackStations(uint64(l))
		}
	}
	if m.LatestPod != 0 {
		n += 1 + sovExtTrackStations(uint64(m.LatestPod))
	}
	l = len(m.LatestAcknowledgementHash)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	l = len(m.StationType)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	if m.FheEnabled {
		n += 2
	}
	l = len(m.SequencerDetails)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	l = len(m.DaDetails)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	l = len(m.ProverDetails)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	l = len(m.StationSchemaKey)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovExtTrackStations(uint64(l))
	}
	return n
}

func sovExtTrackStations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtTrackStations(x uint64) (n int) {
	return sovExtTrackStations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtTrackStations) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtTrackStations
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
			return fmt.Errorf("proto: ExtTrackStations: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtTrackStations: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operators", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operators = append(m.Operators, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestPod", wireType)
			}
			m.LatestPod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestPod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestAcknowledgementHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestAcknowledgementHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FheEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
			m.FheEnabled = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencerDetails", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequencerDetails = append(m.SequencerDetails[:0], dAtA[iNdEx:postIndex]...)
			if m.SequencerDetails == nil {
				m.SequencerDetails = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaDetails", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaDetails = append(m.DaDetails[:0], dAtA[iNdEx:postIndex]...)
			if m.DaDetails == nil {
				m.DaDetails = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProverDetails", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProverDetails = append(m.ProverDetails[:0], dAtA[iNdEx:postIndex]...)
			if m.ProverDetails == nil {
				m.ProverDetails = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationSchemaKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationSchemaKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackStations
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
				return ErrInvalidLengthExtTrackStations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackStations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExtTrackStations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtTrackStations
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
func skipExtTrackStations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtTrackStations
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
					return 0, ErrIntOverflowExtTrackStations
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
					return 0, ErrIntOverflowExtTrackStations
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
				return 0, ErrInvalidLengthExtTrackStations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtTrackStations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtTrackStations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtTrackStations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtTrackStations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtTrackStations = fmt.Errorf("proto: unexpected end of group")
)
