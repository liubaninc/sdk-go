// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wasm/contract.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Contract_Status int32

const (
	Contract_NORMAL Contract_Status = 0
	Contract_FROZE  Contract_Status = 1
)

var Contract_Status_name = map[int32]string{
	0: "NORMAL",
	1: "FROZE",
}

var Contract_Status_value = map[string]int32{
	"NORMAL": 0,
	"FROZE":  1,
}

func (x Contract_Status) String() string {
	return proto.EnumName(Contract_Status_name, int32(x))
}

func (Contract_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_783ee2fe2c5dfbee, []int{0, 0}
}

type Contract struct {
	Creator  string          `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name     string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Hash     string          `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	InitArgs string          `protobuf:"bytes,5,opt,name=initArgs,proto3" json:"initArgs,omitempty"`
	Status   Contract_Status `protobuf:"varint,15,opt,name=status,proto3,enum=liubaninc.m1.wasm.Contract_Status" json:"status,omitempty"`
	Version  uint32          `protobuf:"varint,16,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_783ee2fe2c5dfbee, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Contract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contract) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Contract) GetInitArgs() string {
	if m != nil {
		return m.InitArgs
	}
	return ""
}

func (m *Contract) GetStatus() Contract_Status {
	if m != nil {
		return m.Status
	}
	return Contract_NORMAL
}

func (m *Contract) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterEnum("liubaninc.m1.wasm.Contract_Status", Contract_Status_name, Contract_Status_value)
	proto.RegisterType((*Contract)(nil), "liubaninc.m1.wasm.Contract")
}

func init() { proto.RegisterFile("wasm/contract.proto", fileDescriptor_783ee2fe2c5dfbee) }

var fileDescriptor_783ee2fe2c5dfbee = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x63, 0xd1, 0x86, 0xd6, 0x12, 0x50, 0x0c, 0x83, 0xd5, 0xc1, 0x44, 0x99, 0xb2, 0x60,
	0x0b, 0xd8, 0xd8, 0x0a, 0x02, 0x31, 0x00, 0x95, 0xc2, 0xd6, 0xcd, 0x49, 0xa3, 0xc4, 0xa2, 0xb1,
	0x2b, 0xdb, 0x01, 0xf1, 0x16, 0x3c, 0x16, 0x63, 0x47, 0x36, 0x50, 0xf2, 0x22, 0x28, 0x4e, 0xd3,
	0x85, 0xed, 0xee, 0xfe, 0xfb, 0xa5, 0x4f, 0x07, 0x4f, 0xde, 0xb9, 0x29, 0x59, 0xaa, 0xa4, 0xd5,
	0x3c, 0xb5, 0x74, 0xad, 0x95, 0x55, 0xe8, 0x78, 0x25, 0xaa, 0x84, 0x4b, 0x21, 0x53, 0x5a, 0x5e,
	0xd0, 0xb6, 0x31, 0x3d, 0xcd, 0x55, 0xae, 0xdc, 0x95, 0xb5, 0xaa, 0x2b, 0x86, 0x3f, 0x00, 0x8e,
	0x6e, 0xb7, 0xbf, 0x08, 0xc3, 0xfd, 0x54, 0x67, 0xdc, 0x2a, 0x8d, 0x41, 0x00, 0xa2, 0x71, 0xdc,
	0x5b, 0x84, 0xe0, 0x40, 0xf2, 0x32, 0xc3, 0x7b, 0x2e, 0x76, 0xba, 0xcd, 0x0a, 0x6e, 0x0a, 0x3c,
	0xe8, 0xb2, 0x56, 0xa3, 0x29, 0x1c, 0x09, 0x29, 0xec, 0x4c, 0xe7, 0x06, 0x0f, 0x5d, 0xbe, 0xf3,
	0xe8, 0x1a, 0xfa, 0xc6, 0x72, 0x5b, 0x19, 0x7c, 0x14, 0x80, 0xe8, 0xf0, 0x32, 0xa4, 0xff, 0x20,
	0x69, 0x8f, 0x42, 0x5f, 0x5c, 0x33, 0xde, 0x7e, 0xb4, 0x64, 0x6f, 0x99, 0x36, 0x42, 0x49, 0x3c,
	0x09, 0x40, 0x74, 0x10, 0xf7, 0x36, 0x3c, 0x83, 0x7e, 0xd7, 0x45, 0x10, 0xfa, 0xcf, 0xf3, 0xf8,
	0x69, 0xf6, 0x38, 0xf1, 0xd0, 0x18, 0x0e, 0xef, 0xe3, 0xf9, 0xe2, 0x6e, 0x02, 0x6e, 0x1e, 0xbe,
	0x6a, 0x02, 0x36, 0x35, 0x01, 0xbf, 0x35, 0x01, 0x9f, 0x0d, 0xf1, 0x36, 0x0d, 0xf1, 0xbe, 0x1b,
	0xe2, 0x2d, 0x68, 0x2e, 0x6c, 0x51, 0x25, 0x34, 0x55, 0x25, 0xdb, 0xa1, 0x30, 0xb3, 0x7c, 0x3d,
	0xcf, 0x15, 0x2b, 0xd5, 0xb2, 0x5a, 0x65, 0x86, 0xb9, 0x75, 0xed, 0xc7, 0x3a, 0x33, 0x89, 0xef,
	0x26, 0xbb, 0xfa, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xd3, 0x99, 0x82, 0x72, 0x01, 0x00, 0x00,
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.Status != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x78
	}
	if len(m.InitArgs) > 0 {
		i -= len(m.InitArgs)
		copy(dAtA[i:], m.InitArgs)
		i = encodeVarintContract(dAtA, i, uint64(len(m.InitArgs)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.InitArgs)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovContract(uint64(m.Status))
	}
	if m.Version != 0 {
		n += 2 + sovContract(uint64(m.Version))
	}
	return n
}

func sovContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContract(x uint64) (n int) {
	return sovContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitArgs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitArgs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Contract_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContract
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
func skipContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
				return 0, ErrInvalidLengthContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContract = fmt.Errorf("proto: unexpected end of group")
)