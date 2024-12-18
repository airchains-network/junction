// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/cipherpodledger/fhvms_meta.proto

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

type FhvmsMeta struct {
	ChainId                     string `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	ChainName                   string `protobuf:"bytes,2,opt,name=chainName,proto3" json:"chainName,omitempty"`
	ProofType                   string `protobuf:"bytes,3,opt,name=proofType,proto3" json:"proofType,omitempty"`
	DaProvider                  string `protobuf:"bytes,4,opt,name=daProvider,proto3" json:"daProvider,omitempty"`
	DaBlobId                    string `protobuf:"bytes,5,opt,name=daBlobId,proto3" json:"daBlobId,omitempty"`
	RelayerGAddress             string `protobuf:"bytes,6,opt,name=relayerGAddress,proto3" json:"relayerGAddress,omitempty"`
	RelayerAscAddress           string `protobuf:"bytes,7,opt,name=relayerAscAddress,proto3" json:"relayerAscAddress,omitempty"`
	PicContractAddress          string `protobuf:"bytes,8,opt,name=picContractAddress,proto3" json:"picContractAddress,omitempty"`
	AclContractAddress          string `protobuf:"bytes,9,opt,name=aclContractAddress,proto3" json:"aclContractAddress,omitempty"`
	TfheExecutorContractAddress string `protobuf:"bytes,10,opt,name=tfheExecutorContractAddress,proto3" json:"tfheExecutorContractAddress,omitempty"`
	KmsVerifierContractAddress  string `protobuf:"bytes,11,opt,name=kmsVerifierContractAddress,proto3" json:"kmsVerifierContractAddress,omitempty"`
	GatewayContractAddress      string `protobuf:"bytes,12,opt,name=gatewayContractAddress,proto3" json:"gatewayContractAddress,omitempty"`
	AscChildContractAddress     string `protobuf:"bytes,13,opt,name=ascChildContractAddress,proto3" json:"ascChildContractAddress,omitempty"`
	LatestVerifiedPodNumber     uint64 `protobuf:"varint,14,opt,name=latestVerifiedPodNumber,proto3" json:"latestVerifiedPodNumber,omitempty"`
	FinalityPodNumber           uint64 `protobuf:"varint,15,opt,name=finalityPodNumber,proto3" json:"finalityPodNumber,omitempty"`
	Status                      bool   `protobuf:"varint,16,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *FhvmsMeta) Reset()         { *m = FhvmsMeta{} }
func (m *FhvmsMeta) String() string { return proto.CompactTextString(m) }
func (*FhvmsMeta) ProtoMessage()    {}
func (*FhvmsMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_2997be451ff4723d, []int{0}
}
func (m *FhvmsMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FhvmsMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FhvmsMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FhvmsMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FhvmsMeta.Merge(m, src)
}
func (m *FhvmsMeta) XXX_Size() int {
	return m.Size()
}
func (m *FhvmsMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_FhvmsMeta.DiscardUnknown(m)
}

var xxx_messageInfo_FhvmsMeta proto.InternalMessageInfo

func (m *FhvmsMeta) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *FhvmsMeta) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *FhvmsMeta) GetProofType() string {
	if m != nil {
		return m.ProofType
	}
	return ""
}

func (m *FhvmsMeta) GetDaProvider() string {
	if m != nil {
		return m.DaProvider
	}
	return ""
}

func (m *FhvmsMeta) GetDaBlobId() string {
	if m != nil {
		return m.DaBlobId
	}
	return ""
}

func (m *FhvmsMeta) GetRelayerGAddress() string {
	if m != nil {
		return m.RelayerGAddress
	}
	return ""
}

func (m *FhvmsMeta) GetRelayerAscAddress() string {
	if m != nil {
		return m.RelayerAscAddress
	}
	return ""
}

func (m *FhvmsMeta) GetPicContractAddress() string {
	if m != nil {
		return m.PicContractAddress
	}
	return ""
}

func (m *FhvmsMeta) GetAclContractAddress() string {
	if m != nil {
		return m.AclContractAddress
	}
	return ""
}

func (m *FhvmsMeta) GetTfheExecutorContractAddress() string {
	if m != nil {
		return m.TfheExecutorContractAddress
	}
	return ""
}

func (m *FhvmsMeta) GetKmsVerifierContractAddress() string {
	if m != nil {
		return m.KmsVerifierContractAddress
	}
	return ""
}

func (m *FhvmsMeta) GetGatewayContractAddress() string {
	if m != nil {
		return m.GatewayContractAddress
	}
	return ""
}

func (m *FhvmsMeta) GetAscChildContractAddress() string {
	if m != nil {
		return m.AscChildContractAddress
	}
	return ""
}

func (m *FhvmsMeta) GetLatestVerifiedPodNumber() uint64 {
	if m != nil {
		return m.LatestVerifiedPodNumber
	}
	return 0
}

func (m *FhvmsMeta) GetFinalityPodNumber() uint64 {
	if m != nil {
		return m.FinalityPodNumber
	}
	return 0
}

func (m *FhvmsMeta) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*FhvmsMeta)(nil), "junction.cipherpodledger.FhvmsMeta")
}

func init() {
	proto.RegisterFile("junction/cipherpodledger/fhvms_meta.proto", fileDescriptor_2997be451ff4723d)
}

var fileDescriptor_2997be451ff4723d = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x63, 0x68, 0xd3, 0x64, 0xf9, 0x53, 0xd8, 0x43, 0x59, 0x15, 0x64, 0x45, 0x9c, 0x82,
	0x04, 0xc9, 0x01, 0x09, 0x71, 0x01, 0xd1, 0x56, 0x80, 0x7a, 0xa0, 0xaa, 0x22, 0x04, 0x12, 0x17,
	0x34, 0xd9, 0x1d, 0xc7, 0x4b, 0x6d, 0xaf, 0xb5, 0x3b, 0x6e, 0xeb, 0xb7, 0xe0, 0xb1, 0xe0, 0xd6,
	0x23, 0x47, 0x94, 0xbc, 0x08, 0xf2, 0x36, 0x4e, 0x91, 0xdd, 0xf4, 0xe6, 0xf9, 0x7e, 0xbf, 0x6f,
	0x64, 0xc9, 0x1e, 0xf6, 0xec, 0x47, 0x91, 0x49, 0xd2, 0x26, 0x1b, 0x4b, 0x9d, 0xc7, 0x68, 0x73,
	0xa3, 0x12, 0x54, 0x33, 0xb4, 0xe3, 0x28, 0x3e, 0x4d, 0xdd, 0xf7, 0x14, 0x09, 0x46, 0xb9, 0x35,
	0x64, 0xb8, 0xa8, 0xd5, 0x51, 0x43, 0x7d, 0xfa, 0x7b, 0x93, 0xf5, 0x3f, 0x54, 0xfa, 0x27, 0x24,
	0xe0, 0x82, 0x6d, 0xc9, 0x18, 0x74, 0x76, 0xa8, 0x44, 0x30, 0x08, 0x86, 0xfd, 0x49, 0x3d, 0xf2,
	0x27, 0xac, 0xef, 0x1f, 0x8f, 0x20, 0x45, 0x71, 0xcb, 0xb3, 0xab, 0xa0, 0xa2, 0xb9, 0x35, 0x26,
	0xfa, 0x5c, 0xe6, 0x28, 0x6e, 0x5f, 0xd2, 0x55, 0xc0, 0x43, 0xc6, 0x14, 0x1c, 0x5b, 0x73, 0xaa,
	0x15, 0x5a, 0xb1, 0xe1, 0xf1, 0x7f, 0x09, 0xdf, 0x65, 0x3d, 0x05, 0xfb, 0x89, 0x99, 0x1e, 0x2a,
	0xb1, 0xe9, 0xe9, 0x6a, 0xe6, 0x43, 0xb6, 0x6d, 0x31, 0x81, 0x12, 0xed, 0xc7, 0x3d, 0xa5, 0x2c,
	0x3a, 0x27, 0xba, 0x5e, 0x69, 0xc6, 0xfc, 0x39, 0x7b, 0xb8, 0x8c, 0xf6, 0x9c, 0xac, 0xdd, 0x2d,
	0xef, 0xb6, 0x01, 0x1f, 0x31, 0x9e, 0x6b, 0x79, 0x60, 0x32, 0xb2, 0x20, 0xa9, 0xd6, 0x7b, 0x5e,
	0xbf, 0x86, 0x54, 0x3e, 0xc8, 0xa4, 0xe9, 0xf7, 0x2f, 0xfd, 0x36, 0xe1, 0xef, 0xd8, 0x63, 0x8a,
	0x62, 0x7c, 0x7f, 0x8e, 0xb2, 0x20, 0x63, 0x9b, 0x45, 0xe6, 0x8b, 0x37, 0x29, 0xfc, 0x2d, 0xdb,
	0x3d, 0x49, 0xdd, 0x17, 0xb4, 0x3a, 0xd2, 0xd8, 0x5a, 0x70, 0xc7, 0x2f, 0xb8, 0xc1, 0xe0, 0xaf,
	0xd8, 0xce, 0x0c, 0x08, 0xcf, 0xa0, 0x6c, 0x76, 0xef, 0xfa, 0xee, 0x1a, 0xca, 0x5f, 0xb3, 0x47,
	0xe0, 0xe4, 0x41, 0xac, 0x13, 0xd5, 0x2c, 0xde, 0xf3, 0xc5, 0x75, 0xb8, 0x6a, 0x26, 0x40, 0xe8,
	0x68, 0xf9, 0x4a, 0xea, 0xd8, 0xa8, 0xa3, 0x22, 0x9d, 0xa2, 0x15, 0xf7, 0x07, 0xc1, 0x70, 0x63,
	0xb2, 0x0e, 0x57, 0xdf, 0x2e, 0xd2, 0x19, 0x24, 0x9a, 0xca, 0xab, 0xce, 0xb6, 0xef, 0xb4, 0x01,
	0xdf, 0x61, 0x5d, 0x47, 0x40, 0x85, 0x13, 0x0f, 0x06, 0xc1, 0xb0, 0x37, 0x59, 0x4e, 0xfb, 0x5f,
	0x7f, 0xcd, 0xc3, 0xe0, 0x62, 0x1e, 0x06, 0x7f, 0xe7, 0x61, 0xf0, 0x73, 0x11, 0x76, 0x2e, 0x16,
	0x61, 0xe7, 0xcf, 0x22, 0xec, 0x7c, 0x7b, 0x33, 0xd3, 0x14, 0x17, 0xd3, 0x91, 0x34, 0xe9, 0x18,
	0xb4, 0xf5, 0x3f, 0xae, 0x7b, 0x91, 0x21, 0x9d, 0x19, 0x7b, 0x32, 0x5e, 0xdd, 0xd1, 0x79, 0xeb,
	0x92, 0xa8, 0xcc, 0xd1, 0x4d, 0xbb, 0xfe, 0x8a, 0x5e, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xcf,
	0xe5, 0x68, 0xfe, 0x72, 0x03, 0x00, 0x00,
}

func (m *FhvmsMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FhvmsMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FhvmsMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status {
		i--
		if m.Status {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.FinalityPodNumber != 0 {
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(m.FinalityPodNumber))
		i--
		dAtA[i] = 0x78
	}
	if m.LatestVerifiedPodNumber != 0 {
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(m.LatestVerifiedPodNumber))
		i--
		dAtA[i] = 0x70
	}
	if len(m.AscChildContractAddress) > 0 {
		i -= len(m.AscChildContractAddress)
		copy(dAtA[i:], m.AscChildContractAddress)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.AscChildContractAddress)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.GatewayContractAddress) > 0 {
		i -= len(m.GatewayContractAddress)
		copy(dAtA[i:], m.GatewayContractAddress)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.GatewayContractAddress)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.KmsVerifierContractAddress) > 0 {
		i -= len(m.KmsVerifierContractAddress)
		copy(dAtA[i:], m.KmsVerifierContractAddress)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.KmsVerifierContractAddress)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.TfheExecutorContractAddress) > 0 {
		i -= len(m.TfheExecutorContractAddress)
		copy(dAtA[i:], m.TfheExecutorContractAddress)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.TfheExecutorContractAddress)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.AclContractAddress) > 0 {
		i -= len(m.AclContractAddress)
		copy(dAtA[i:], m.AclContractAddress)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.AclContractAddress)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.PicContractAddress) > 0 {
		i -= len(m.PicContractAddress)
		copy(dAtA[i:], m.PicContractAddress)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.PicContractAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.RelayerAscAddress) > 0 {
		i -= len(m.RelayerAscAddress)
		copy(dAtA[i:], m.RelayerAscAddress)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.RelayerAscAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.RelayerGAddress) > 0 {
		i -= len(m.RelayerGAddress)
		copy(dAtA[i:], m.RelayerGAddress)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.RelayerGAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DaBlobId) > 0 {
		i -= len(m.DaBlobId)
		copy(dAtA[i:], m.DaBlobId)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.DaBlobId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DaProvider) > 0 {
		i -= len(m.DaProvider)
		copy(dAtA[i:], m.DaProvider)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.DaProvider)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ProofType) > 0 {
		i -= len(m.ProofType)
		copy(dAtA[i:], m.ProofType)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.ProofType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintFhvmsMeta(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFhvmsMeta(dAtA []byte, offset int, v uint64) int {
	offset -= sovFhvmsMeta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FhvmsMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.ProofType)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.DaProvider)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.DaBlobId)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.RelayerGAddress)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.RelayerAscAddress)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.PicContractAddress)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.AclContractAddress)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.TfheExecutorContractAddress)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.KmsVerifierContractAddress)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.GatewayContractAddress)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	l = len(m.AscChildContractAddress)
	if l > 0 {
		n += 1 + l + sovFhvmsMeta(uint64(l))
	}
	if m.LatestVerifiedPodNumber != 0 {
		n += 1 + sovFhvmsMeta(uint64(m.LatestVerifiedPodNumber))
	}
	if m.FinalityPodNumber != 0 {
		n += 1 + sovFhvmsMeta(uint64(m.FinalityPodNumber))
	}
	if m.Status {
		n += 3
	}
	return n
}

func sovFhvmsMeta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFhvmsMeta(x uint64) (n int) {
	return sovFhvmsMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FhvmsMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFhvmsMeta
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
			return fmt.Errorf("proto: FhvmsMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FhvmsMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaProvider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaProvider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaBlobId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaBlobId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayerGAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayerGAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayerAscAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayerAscAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PicContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PicContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AclContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AclContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TfheExecutorContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TfheExecutorContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KmsVerifierContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KmsVerifierContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AscChildContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
				return ErrInvalidLengthFhvmsMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFhvmsMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AscChildContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestVerifiedPodNumber", wireType)
			}
			m.LatestVerifiedPodNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestVerifiedPodNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalityPodNumber", wireType)
			}
			m.FinalityPodNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalityPodNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFhvmsMeta
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
			m.Status = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFhvmsMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFhvmsMeta
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
func skipFhvmsMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFhvmsMeta
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
					return 0, ErrIntOverflowFhvmsMeta
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
					return 0, ErrIntOverflowFhvmsMeta
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
				return 0, ErrInvalidLengthFhvmsMeta
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFhvmsMeta
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFhvmsMeta
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFhvmsMeta        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFhvmsMeta          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFhvmsMeta = fmt.Errorf("proto: unexpected end of group")
)