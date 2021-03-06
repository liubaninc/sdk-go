// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: utxo/input.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Output_Purpose int32

const (
	Output_TRANSER   Output_Purpose = 0
	Output_CHANGE    Output_Purpose = 1
	Output_FEE       Output_Purpose = 2
	Output_ISSUE     Output_Purpose = 3
	Output_BURN      Output_Purpose = 4
	Output_Transient Output_Purpose = 5
)

var Output_Purpose_name = map[int32]string{
	0: "TRANSER",
	1: "CHANGE",
	2: "FEE",
	3: "ISSUE",
	4: "BURN",
	5: "Transient",
}

var Output_Purpose_value = map[string]int32{
	"TRANSER":   0,
	"CHANGE":    1,
	"FEE":       2,
	"ISSUE":     3,
	"BURN":      4,
	"Transient": 5,
}

func (x Output_Purpose) String() string {
	return proto.EnumName(Output_Purpose_name, int32(x))
}

func (Output_Purpose) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f017d14d9e0952f, []int{1, 0}
}

type Input struct {
	RefTx     string     `protobuf:"bytes,3,opt,name=refTx,proto3" json:"refTx,omitempty"`
	RefMsg    int32      `protobuf:"varint,4,opt,name=refMsg,proto3" json:"refMsg,omitempty"`
	RefOffset int32      `protobuf:"varint,5,opt,name=refOffset,proto3" json:"refOffset,omitempty"`
	Addr      string     `protobuf:"bytes,6,opt,name=addr,proto3" json:"addr,omitempty"`
	Amount    types.Coin `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount"`
	Frozen    int64      `protobuf:"varint,8,opt,name=frozen,proto3" json:"frozen,omitempty"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f017d14d9e0952f, []int{0}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Input.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return m.Size()
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetRefTx() string {
	if m != nil {
		return m.RefTx
	}
	return ""
}

func (m *Input) GetRefMsg() int32 {
	if m != nil {
		return m.RefMsg
	}
	return 0
}

func (m *Input) GetRefOffset() int32 {
	if m != nil {
		return m.RefOffset
	}
	return 0
}

func (m *Input) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Input) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *Input) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

type Output struct {
	Addr    string         `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Amount  types.Coin     `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount"`
	Frozen  int64          `protobuf:"varint,5,opt,name=frozen,proto3" json:"frozen,omitempty"`
	Purpose Output_Purpose `protobuf:"varint,6,opt,name=purpose,proto3,enum=liubaninc.m1.utxo.Output_Purpose" json:"purpose,omitempty"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f017d14d9e0952f, []int{1}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Output.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return m.Size()
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Output) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *Output) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *Output) GetPurpose() Output_Purpose {
	if m != nil {
		return m.Purpose
	}
	return Output_TRANSER
}

func init() {
	proto.RegisterEnum("liubaninc.m1.utxo.Output_Purpose", Output_Purpose_name, Output_Purpose_value)
	proto.RegisterType((*Input)(nil), "liubaninc.m1.utxo.Input")
	proto.RegisterType((*Output)(nil), "liubaninc.m1.utxo.Output")
}

func init() { proto.RegisterFile("utxo/input.proto", fileDescriptor_2f017d14d9e0952f) }

var fileDescriptor_2f017d14d9e0952f = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbb, 0x6e, 0xd4, 0x40,
	0x14, 0x86, 0x3d, 0xf8, 0x96, 0x3d, 0x11, 0xc8, 0x8c, 0x22, 0x64, 0x22, 0x64, 0x96, 0xad, 0xb6,
	0x61, 0x46, 0x1b, 0x0a, 0x0a, 0xaa, 0x6c, 0x64, 0x48, 0x0a, 0x76, 0x91, 0x77, 0xd3, 0xd0, 0xd9,
	0xeb, 0xb1, 0xb1, 0x88, 0x67, 0xac, 0xb9, 0xa0, 0xc0, 0x53, 0xf0, 0x38, 0x3c, 0x42, 0xca, 0x94,
	0x54, 0x08, 0xed, 0x3e, 0x06, 0x0d, 0xf2, 0x65, 0x01, 0x09, 0x2a, 0xba, 0x73, 0xce, 0xff, 0xfb,
	0xfc, 0x9f, 0x35, 0x07, 0x02, 0xa3, 0xaf, 0x05, 0xad, 0x78, 0x63, 0x34, 0x69, 0xa4, 0xd0, 0x02,
	0xdf, 0xbf, 0xaa, 0x4c, 0x96, 0xf2, 0x8a, 0x6f, 0x48, 0x3d, 0x23, 0xad, 0x7c, 0x7c, 0x54, 0x8a,
	0x52, 0x74, 0x2a, 0x6d, 0xab, 0xde, 0x78, 0x1c, 0x6d, 0x84, 0xaa, 0x85, 0xa2, 0x59, 0xaa, 0x18,
	0xfd, 0x30, 0xcb, 0x98, 0x4e, 0x67, 0x74, 0x23, 0x2a, 0xde, 0xeb, 0x93, 0x2f, 0x08, 0xdc, 0x8b,
	0x76, 0x31, 0x3e, 0x02, 0x57, 0xb2, 0x62, 0x7d, 0x1d, 0xda, 0x63, 0x34, 0x1d, 0x25, 0x7d, 0x83,
	0x1f, 0x80, 0x27, 0x59, 0xf1, 0x5a, 0x95, 0xa1, 0x33, 0x46, 0x53, 0x37, 0x19, 0x3a, 0xfc, 0x08,
	0x46, 0x92, 0x15, 0xcb, 0xa2, 0x50, 0x4c, 0x87, 0x6e, 0x27, 0xfd, 0x1e, 0x60, 0x0c, 0x4e, 0x9a,
	0xe7, 0x32, 0xf4, 0xba, 0x55, 0x5d, 0x8d, 0x9f, 0x83, 0x97, 0xd6, 0xc2, 0x70, 0x1d, 0xfa, 0x63,
	0x34, 0x3d, 0x3c, 0x79, 0x48, 0x7a, 0x34, 0xd2, 0xa2, 0x91, 0x01, 0x8d, 0x9c, 0x89, 0x8a, 0xcf,
	0x9d, 0x9b, 0x6f, 0x8f, 0xad, 0x64, 0xb0, 0xb7, 0x08, 0x85, 0x14, 0x9f, 0x18, 0x0f, 0x0f, 0xc6,
	0x68, 0x6a, 0x27, 0x43, 0x37, 0xf9, 0x81, 0xc0, 0x5b, 0x1a, 0xdd, 0xb2, 0xef, 0xf3, 0xec, 0x7f,
	0xe6, 0x39, 0xff, 0x9b, 0xe7, 0xfe, 0x99, 0x87, 0x5f, 0x80, 0xdf, 0x18, 0xd9, 0x08, 0xc5, 0xba,
	0xff, 0xba, 0x77, 0xf2, 0x84, 0xfc, 0xf5, 0x0a, 0xa4, 0x07, 0x22, 0x6f, 0x7a, 0x63, 0xb2, 0xff,
	0x62, 0xb2, 0x04, 0x7f, 0x98, 0xe1, 0x43, 0xf0, 0xd7, 0xc9, 0xe9, 0x62, 0x15, 0x27, 0x81, 0x85,
	0x01, 0xbc, 0xb3, 0xf3, 0xd3, 0xc5, 0xab, 0x38, 0x40, 0xd8, 0x07, 0xfb, 0x65, 0x1c, 0x07, 0x77,
	0xf0, 0x08, 0xdc, 0x8b, 0xd5, 0xea, 0x32, 0x0e, 0x6c, 0x7c, 0x00, 0xce, 0xfc, 0x32, 0x59, 0x04,
	0x0e, 0xbe, 0x0b, 0xa3, 0xb5, 0x4c, 0xb9, 0xaa, 0x18, 0xd7, 0x81, 0x3b, 0x3f, 0xbf, 0xd9, 0x46,
	0xe8, 0x76, 0x1b, 0xa1, 0xef, 0xdb, 0x08, 0x7d, 0xde, 0x45, 0xd6, 0xed, 0x2e, 0xb2, 0xbe, 0xee,
	0x22, 0xeb, 0x2d, 0x29, 0x2b, 0xfd, 0xce, 0x64, 0x64, 0x23, 0x6a, 0xfa, 0x0b, 0x90, 0xaa, 0xfc,
	0xfd, 0xd3, 0x52, 0xd0, 0x5a, 0xe4, 0xe6, 0x8a, 0x29, 0xda, 0x5d, 0x94, 0xfe, 0xd8, 0x30, 0x95,
	0x79, 0xdd, 0x25, 0x3c, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0x30, 0xde, 0xcf, 0x78, 0x66, 0x02,
	0x00, 0x00,
}

func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Input) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Frozen != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.Frozen))
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintInput(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintInput(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0x32
	}
	if m.RefOffset != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.RefOffset))
		i--
		dAtA[i] = 0x28
	}
	if m.RefMsg != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.RefMsg))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RefTx) > 0 {
		i -= len(m.RefTx)
		copy(dAtA[i:], m.RefTx)
		i = encodeVarintInput(dAtA, i, uint64(len(m.RefTx)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func (m *Output) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Output) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Output) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Purpose != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.Purpose))
		i--
		dAtA[i] = 0x30
	}
	if m.Frozen != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.Frozen))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintInput(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintInput(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func encodeVarintInput(dAtA []byte, offset int, v uint64) int {
	offset -= sovInput(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Input) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RefTx)
	if l > 0 {
		n += 1 + l + sovInput(uint64(l))
	}
	if m.RefMsg != 0 {
		n += 1 + sovInput(uint64(m.RefMsg))
	}
	if m.RefOffset != 0 {
		n += 1 + sovInput(uint64(m.RefOffset))
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovInput(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovInput(uint64(l))
	if m.Frozen != 0 {
		n += 1 + sovInput(uint64(m.Frozen))
	}
	return n
}

func (m *Output) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovInput(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovInput(uint64(l))
	if m.Frozen != 0 {
		n += 1 + sovInput(uint64(m.Frozen))
	}
	if m.Purpose != 0 {
		n += 1 + sovInput(uint64(m.Purpose))
	}
	return n
}

func sovInput(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInput(x uint64) (n int) {
	return sovInput(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInput
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
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefTx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
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
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefTx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefMsg", wireType)
			}
			m.RefMsg = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RefMsg |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefOffset", wireType)
			}
			m.RefOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RefOffset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
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
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			m.Frozen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frozen |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInput(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInput
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
func (m *Output) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInput
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
			return fmt.Errorf("proto: Output: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Output: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
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
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			m.Frozen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frozen |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Purpose", wireType)
			}
			m.Purpose = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Purpose |= Output_Purpose(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInput(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInput
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
func skipInput(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInput
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
					return 0, ErrIntOverflowInput
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
					return 0, ErrIntOverflowInput
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
				return 0, ErrInvalidLengthInput
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInput
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInput
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInput        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInput          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInput = fmt.Errorf("proto: unexpected end of group")
)
