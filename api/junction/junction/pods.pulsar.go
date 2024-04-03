// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package junction

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Pods                        protoreflect.MessageDescriptor
	fd_Pods_podNumber              protoreflect.FieldDescriptor
	fd_Pods_merkleRootHash         protoreflect.FieldDescriptor
	fd_Pods_previousMerkleRootHash protoreflect.FieldDescriptor
	fd_Pods_zkProof                protoreflect.FieldDescriptor
	fd_Pods_witness                protoreflect.FieldDescriptor
	fd_Pods_timestamp              protoreflect.FieldDescriptor
	fd_Pods_isVerified             protoreflect.FieldDescriptor
)

func init() {
	file_junction_junction_pods_proto_init()
	md_Pods = File_junction_junction_pods_proto.Messages().ByName("Pods")
	fd_Pods_podNumber = md_Pods.Fields().ByName("podNumber")
	fd_Pods_merkleRootHash = md_Pods.Fields().ByName("merkleRootHash")
	fd_Pods_previousMerkleRootHash = md_Pods.Fields().ByName("previousMerkleRootHash")
	fd_Pods_zkProof = md_Pods.Fields().ByName("zkProof")
	fd_Pods_witness = md_Pods.Fields().ByName("witness")
	fd_Pods_timestamp = md_Pods.Fields().ByName("timestamp")
	fd_Pods_isVerified = md_Pods.Fields().ByName("isVerified")
}

var _ protoreflect.Message = (*fastReflection_Pods)(nil)

type fastReflection_Pods Pods

func (x *Pods) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Pods)(x)
}

func (x *Pods) slowProtoReflect() protoreflect.Message {
	mi := &file_junction_junction_pods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Pods_messageType fastReflection_Pods_messageType
var _ protoreflect.MessageType = fastReflection_Pods_messageType{}

type fastReflection_Pods_messageType struct{}

func (x fastReflection_Pods_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Pods)(nil)
}
func (x fastReflection_Pods_messageType) New() protoreflect.Message {
	return new(fastReflection_Pods)
}
func (x fastReflection_Pods_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Pods
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Pods) Descriptor() protoreflect.MessageDescriptor {
	return md_Pods
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Pods) Type() protoreflect.MessageType {
	return _fastReflection_Pods_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Pods) New() protoreflect.Message {
	return new(fastReflection_Pods)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Pods) Interface() protoreflect.ProtoMessage {
	return (*Pods)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Pods) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.PodNumber != uint64(0) {
		value := protoreflect.ValueOfUint64(x.PodNumber)
		if !f(fd_Pods_podNumber, value) {
			return
		}
	}
	if x.MerkleRootHash != "" {
		value := protoreflect.ValueOfString(x.MerkleRootHash)
		if !f(fd_Pods_merkleRootHash, value) {
			return
		}
	}
	if x.PreviousMerkleRootHash != "" {
		value := protoreflect.ValueOfString(x.PreviousMerkleRootHash)
		if !f(fd_Pods_previousMerkleRootHash, value) {
			return
		}
	}
	if len(x.ZkProof) != 0 {
		value := protoreflect.ValueOfBytes(x.ZkProof)
		if !f(fd_Pods_zkProof, value) {
			return
		}
	}
	if len(x.Witness) != 0 {
		value := protoreflect.ValueOfBytes(x.Witness)
		if !f(fd_Pods_witness, value) {
			return
		}
	}
	if x.Timestamp != "" {
		value := protoreflect.ValueOfString(x.Timestamp)
		if !f(fd_Pods_timestamp, value) {
			return
		}
	}
	if x.IsVerified != false {
		value := protoreflect.ValueOfBool(x.IsVerified)
		if !f(fd_Pods_isVerified, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Pods) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "junction.junction.Pods.podNumber":
		return x.PodNumber != uint64(0)
	case "junction.junction.Pods.merkleRootHash":
		return x.MerkleRootHash != ""
	case "junction.junction.Pods.previousMerkleRootHash":
		return x.PreviousMerkleRootHash != ""
	case "junction.junction.Pods.zkProof":
		return len(x.ZkProof) != 0
	case "junction.junction.Pods.witness":
		return len(x.Witness) != 0
	case "junction.junction.Pods.timestamp":
		return x.Timestamp != ""
	case "junction.junction.Pods.isVerified":
		return x.IsVerified != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.junction.Pods"))
		}
		panic(fmt.Errorf("message junction.junction.Pods does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pods) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "junction.junction.Pods.podNumber":
		x.PodNumber = uint64(0)
	case "junction.junction.Pods.merkleRootHash":
		x.MerkleRootHash = ""
	case "junction.junction.Pods.previousMerkleRootHash":
		x.PreviousMerkleRootHash = ""
	case "junction.junction.Pods.zkProof":
		x.ZkProof = nil
	case "junction.junction.Pods.witness":
		x.Witness = nil
	case "junction.junction.Pods.timestamp":
		x.Timestamp = ""
	case "junction.junction.Pods.isVerified":
		x.IsVerified = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.junction.Pods"))
		}
		panic(fmt.Errorf("message junction.junction.Pods does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Pods) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "junction.junction.Pods.podNumber":
		value := x.PodNumber
		return protoreflect.ValueOfUint64(value)
	case "junction.junction.Pods.merkleRootHash":
		value := x.MerkleRootHash
		return protoreflect.ValueOfString(value)
	case "junction.junction.Pods.previousMerkleRootHash":
		value := x.PreviousMerkleRootHash
		return protoreflect.ValueOfString(value)
	case "junction.junction.Pods.zkProof":
		value := x.ZkProof
		return protoreflect.ValueOfBytes(value)
	case "junction.junction.Pods.witness":
		value := x.Witness
		return protoreflect.ValueOfBytes(value)
	case "junction.junction.Pods.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfString(value)
	case "junction.junction.Pods.isVerified":
		value := x.IsVerified
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.junction.Pods"))
		}
		panic(fmt.Errorf("message junction.junction.Pods does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pods) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "junction.junction.Pods.podNumber":
		x.PodNumber = value.Uint()
	case "junction.junction.Pods.merkleRootHash":
		x.MerkleRootHash = value.Interface().(string)
	case "junction.junction.Pods.previousMerkleRootHash":
		x.PreviousMerkleRootHash = value.Interface().(string)
	case "junction.junction.Pods.zkProof":
		x.ZkProof = value.Bytes()
	case "junction.junction.Pods.witness":
		x.Witness = value.Bytes()
	case "junction.junction.Pods.timestamp":
		x.Timestamp = value.Interface().(string)
	case "junction.junction.Pods.isVerified":
		x.IsVerified = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.junction.Pods"))
		}
		panic(fmt.Errorf("message junction.junction.Pods does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pods) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "junction.junction.Pods.podNumber":
		panic(fmt.Errorf("field podNumber of message junction.junction.Pods is not mutable"))
	case "junction.junction.Pods.merkleRootHash":
		panic(fmt.Errorf("field merkleRootHash of message junction.junction.Pods is not mutable"))
	case "junction.junction.Pods.previousMerkleRootHash":
		panic(fmt.Errorf("field previousMerkleRootHash of message junction.junction.Pods is not mutable"))
	case "junction.junction.Pods.zkProof":
		panic(fmt.Errorf("field zkProof of message junction.junction.Pods is not mutable"))
	case "junction.junction.Pods.witness":
		panic(fmt.Errorf("field witness of message junction.junction.Pods is not mutable"))
	case "junction.junction.Pods.timestamp":
		panic(fmt.Errorf("field timestamp of message junction.junction.Pods is not mutable"))
	case "junction.junction.Pods.isVerified":
		panic(fmt.Errorf("field isVerified of message junction.junction.Pods is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.junction.Pods"))
		}
		panic(fmt.Errorf("message junction.junction.Pods does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Pods) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "junction.junction.Pods.podNumber":
		return protoreflect.ValueOfUint64(uint64(0))
	case "junction.junction.Pods.merkleRootHash":
		return protoreflect.ValueOfString("")
	case "junction.junction.Pods.previousMerkleRootHash":
		return protoreflect.ValueOfString("")
	case "junction.junction.Pods.zkProof":
		return protoreflect.ValueOfBytes(nil)
	case "junction.junction.Pods.witness":
		return protoreflect.ValueOfBytes(nil)
	case "junction.junction.Pods.timestamp":
		return protoreflect.ValueOfString("")
	case "junction.junction.Pods.isVerified":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.junction.Pods"))
		}
		panic(fmt.Errorf("message junction.junction.Pods does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Pods) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in junction.junction.Pods", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Pods) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pods) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Pods) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Pods) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Pods)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.PodNumber != 0 {
			n += 1 + runtime.Sov(uint64(x.PodNumber))
		}
		l = len(x.MerkleRootHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.PreviousMerkleRootHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ZkProof)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Witness)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Timestamp)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.IsVerified {
			n += 2
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Pods)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.IsVerified {
			i--
			if x.IsVerified {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x38
		}
		if len(x.Timestamp) > 0 {
			i -= len(x.Timestamp)
			copy(dAtA[i:], x.Timestamp)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Timestamp)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.Witness) > 0 {
			i -= len(x.Witness)
			copy(dAtA[i:], x.Witness)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Witness)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.ZkProof) > 0 {
			i -= len(x.ZkProof)
			copy(dAtA[i:], x.ZkProof)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ZkProof)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.PreviousMerkleRootHash) > 0 {
			i -= len(x.PreviousMerkleRootHash)
			copy(dAtA[i:], x.PreviousMerkleRootHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PreviousMerkleRootHash)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.MerkleRootHash) > 0 {
			i -= len(x.MerkleRootHash)
			copy(dAtA[i:], x.MerkleRootHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MerkleRootHash)))
			i--
			dAtA[i] = 0x12
		}
		if x.PodNumber != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.PodNumber))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Pods)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Pods: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Pods: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PodNumber", wireType)
				}
				x.PodNumber = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.PodNumber |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MerkleRootHash", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MerkleRootHash = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PreviousMerkleRootHash", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.PreviousMerkleRootHash = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ZkProof", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ZkProof = append(x.ZkProof[:0], dAtA[iNdEx:postIndex]...)
				if x.ZkProof == nil {
					x.ZkProof = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Witness", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Witness = append(x.Witness[:0], dAtA[iNdEx:postIndex]...)
				if x.Witness == nil {
					x.Witness = []byte{}
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Timestamp = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field IsVerified", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.IsVerified = bool(v != 0)
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: junction/junction/pods.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodNumber              uint64 `protobuf:"varint,1,opt,name=podNumber,proto3" json:"podNumber,omitempty"`
	MerkleRootHash         string `protobuf:"bytes,2,opt,name=merkleRootHash,proto3" json:"merkleRootHash,omitempty"`
	PreviousMerkleRootHash string `protobuf:"bytes,3,opt,name=previousMerkleRootHash,proto3" json:"previousMerkleRootHash,omitempty"`
	ZkProof                []byte `protobuf:"bytes,4,opt,name=zkProof,proto3" json:"zkProof,omitempty"`
	Witness                []byte `protobuf:"bytes,5,opt,name=witness,proto3" json:"witness,omitempty"`
	Timestamp              string `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsVerified             bool   `protobuf:"varint,7,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
}

func (x *Pods) Reset() {
	*x = Pods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_junction_junction_pods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pods) ProtoMessage() {}

// Deprecated: Use Pods.ProtoReflect.Descriptor instead.
func (*Pods) Descriptor() ([]byte, []int) {
	return file_junction_junction_pods_proto_rawDescGZIP(), []int{0}
}

func (x *Pods) GetPodNumber() uint64 {
	if x != nil {
		return x.PodNumber
	}
	return 0
}

func (x *Pods) GetMerkleRootHash() string {
	if x != nil {
		return x.MerkleRootHash
	}
	return ""
}

func (x *Pods) GetPreviousMerkleRootHash() string {
	if x != nil {
		return x.PreviousMerkleRootHash
	}
	return ""
}

func (x *Pods) GetZkProof() []byte {
	if x != nil {
		return x.ZkProof
	}
	return nil
}

func (x *Pods) GetWitness() []byte {
	if x != nil {
		return x.Witness
	}
	return nil
}

func (x *Pods) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Pods) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

var File_junction_junction_pods_proto protoreflect.FileDescriptor

var file_junction_junction_pods_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xf6, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70,
	0x6f, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x36, 0x0a, 0x16, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65,
	0x52, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x6b, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x7a, 0x6b, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0xab, 0x01, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x2e, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x50, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x22, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x4a, 0x4a, 0x58, 0xaa, 0x02, 0x11, 0x4a, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xca,
	0x02, 0x11, 0x4a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x4a, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x1d, 0x4a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x4a,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x4a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a,
	0x4a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_junction_junction_pods_proto_rawDescOnce sync.Once
	file_junction_junction_pods_proto_rawDescData = file_junction_junction_pods_proto_rawDesc
)

func file_junction_junction_pods_proto_rawDescGZIP() []byte {
	file_junction_junction_pods_proto_rawDescOnce.Do(func() {
		file_junction_junction_pods_proto_rawDescData = protoimpl.X.CompressGZIP(file_junction_junction_pods_proto_rawDescData)
	})
	return file_junction_junction_pods_proto_rawDescData
}

var file_junction_junction_pods_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_junction_junction_pods_proto_goTypes = []interface{}{
	(*Pods)(nil), // 0: junction.junction.Pods
}
var file_junction_junction_pods_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_junction_junction_pods_proto_init() }
func file_junction_junction_pods_proto_init() {
	if File_junction_junction_pods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_junction_junction_pods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pods); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_junction_junction_pods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_junction_junction_pods_proto_goTypes,
		DependencyIndexes: file_junction_junction_pods_proto_depIdxs,
		MessageInfos:      file_junction_junction_pods_proto_msgTypes,
	}.Build()
	File_junction_junction_pods_proto = out.File
	file_junction_junction_pods_proto_rawDesc = nil
	file_junction_junction_pods_proto_goTypes = nil
	file_junction_junction_pods_proto_depIdxs = nil
}
