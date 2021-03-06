// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: utxo/token.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/bank/types"
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

type Token struct {
	Creator string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Base    string                                   `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Issued  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=issued,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"issued"`
	Burned  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=burned,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"burned"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9bea73c90a1b05, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Token) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Token) GetIssued() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Issued
	}
	return nil
}

func (m *Token) GetBurned() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Burned
	}
	return nil
}

type TokenDetail struct {
	Creator  string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Metadata *types1.Metadata                         `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Issued   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=issued,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"issued"`
	Burned   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=burned,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"burned"`
}

func (m *TokenDetail) Reset()         { *m = TokenDetail{} }
func (m *TokenDetail) String() string { return proto.CompactTextString(m) }
func (*TokenDetail) ProtoMessage()    {}
func (*TokenDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9bea73c90a1b05, []int{1}
}
func (m *TokenDetail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenDetail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenDetail.Merge(m, src)
}
func (m *TokenDetail) XXX_Size() int {
	return m.Size()
}
func (m *TokenDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TokenDetail proto.InternalMessageInfo

func (m *TokenDetail) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TokenDetail) GetMetadata() *types1.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TokenDetail) GetIssued() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Issued
	}
	return nil
}

func (m *TokenDetail) GetBurned() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Burned
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "liubaninc.m1.utxo.Token")
	proto.RegisterType((*TokenDetail)(nil), "liubaninc.m1.utxo.TokenDetail")
}

func init() { proto.RegisterFile("utxo/token.proto", fileDescriptor_fd9bea73c90a1b05) }

var fileDescriptor_fd9bea73c90a1b05 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x52, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x4d, 0xda, 0x52, 0xc0, 0x5d, 0x20, 0x62, 0x08, 0x95, 0x70, 0xab, 0x4e, 0x5d, 0x6a, 0xd3,
	0x32, 0xb1, 0x16, 0x06, 0x16, 0x96, 0x8a, 0x89, 0xcd, 0x49, 0xac, 0x60, 0xa5, 0xf1, 0x55, 0xb1,
	0x8d, 0xe0, 0x2f, 0xf8, 0x0e, 0xc4, 0x87, 0x74, 0xec, 0xc8, 0x04, 0xa8, 0xfd, 0x10, 0x90, 0x9d,
	0x10, 0x90, 0x90, 0xd8, 0x58, 0x98, 0x7c, 0xe7, 0x77, 0xf7, 0xde, 0xe9, 0xe9, 0xa1, 0x3d, 0xa3,
	0xef, 0x80, 0x6a, 0xc8, 0xb8, 0x24, 0x8b, 0x02, 0x34, 0x04, 0xfb, 0x73, 0x61, 0x22, 0x26, 0x85,
	0x8c, 0x49, 0x3e, 0x26, 0x16, 0xee, 0x1e, 0xa4, 0x90, 0x82, 0x43, 0xa9, 0xad, 0xca, 0xc1, 0x2e,
	0x8e, 0x41, 0xe5, 0xa0, 0x68, 0xc4, 0x64, 0x46, 0x6f, 0xc7, 0x11, 0xd7, 0x6c, 0xec, 0x9a, 0x1f,
	0xb8, 0xe2, 0x35, 0x1e, 0x83, 0xa8, 0x84, 0x06, 0xef, 0x3e, 0xda, 0xba, 0xb2, 0xc2, 0x41, 0x88,
	0xb6, 0xe3, 0x82, 0x33, 0x0d, 0x45, 0xe8, 0xf7, 0xfd, 0xe1, 0xee, 0xec, 0xb3, 0x0d, 0x02, 0xd4,
	0xb2, 0xeb, 0x61, 0xc3, 0x7d, 0xbb, 0x3a, 0x88, 0x51, 0x5b, 0x28, 0x65, 0x78, 0x12, 0x36, 0xfb,
	0xcd, 0x61, 0x67, 0x72, 0x48, 0x4a, 0x21, 0x62, 0x51, 0x52, 0x09, 0x91, 0x33, 0x10, 0x72, 0x7a,
	0xbc, 0x7c, 0xe9, 0x79, 0x8f, 0xaf, 0xbd, 0x61, 0x2a, 0xf4, 0x8d, 0x89, 0x48, 0x0c, 0x39, 0xad,
	0xae, 0x2a, 0x9f, 0x91, 0x4a, 0x32, 0xaa, 0xef, 0x17, 0x5c, 0xb9, 0x05, 0x35, 0xab, 0xa8, 0xad,
	0x48, 0x64, 0x0a, 0xc9, 0x93, 0xb0, 0xf5, 0x07, 0x22, 0x25, 0xf5, 0xe0, 0xa9, 0x81, 0x3a, 0xce,
	0x81, 0x73, 0xae, 0x99, 0x98, 0xff, 0xe2, 0xc3, 0x29, 0xda, 0xc9, 0xb9, 0x66, 0x09, 0xd3, 0xcc,
	0x79, 0xd1, 0x99, 0x1c, 0x7d, 0x1d, 0x24, 0xb3, 0xfa, 0xa0, 0xcb, 0x6a, 0x68, 0x56, 0x8f, 0xff,
	0x1f, 0xbb, 0xa6, 0x17, 0xcb, 0x35, 0xf6, 0x57, 0x6b, 0xec, 0xbf, 0xad, 0xb1, 0xff, 0xb0, 0xc1,
	0xde, 0x6a, 0x83, 0xbd, 0xe7, 0x0d, 0xf6, 0xae, 0xc9, 0x37, 0xae, 0x3a, 0xbe, 0x54, 0x25, 0xd9,
	0x28, 0x05, 0x9a, 0x43, 0x62, 0xe6, 0x5c, 0xd1, 0x32, 0xe9, 0x96, 0x37, 0x6a, 0xbb, 0x04, 0x9e,
	0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xd7, 0x01, 0xe8, 0xfe, 0x02, 0x00, 0x00,
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Burned) > 0 {
		for iNdEx := len(m.Burned) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Burned[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintToken(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Issued) > 0 {
		for iNdEx := len(m.Issued) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issued[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintToken(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Base) > 0 {
		i -= len(m.Base)
		copy(dAtA[i:], m.Base)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Base)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenDetail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenDetail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenDetail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Burned) > 0 {
		for iNdEx := len(m.Burned) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Burned[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintToken(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Issued) > 0 {
		for iNdEx := len(m.Issued) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issued[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintToken(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintToken(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Base)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if len(m.Issued) > 0 {
		for _, e := range m.Issued {
			l = e.Size()
			n += 1 + l + sovToken(uint64(l))
		}
	}
	if len(m.Burned) > 0 {
		for _, e := range m.Burned {
			l = e.Size()
			n += 1 + l + sovToken(uint64(l))
		}
	}
	return n
}

func (m *TokenDetail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovToken(uint64(l))
	}
	if len(m.Issued) > 0 {
		for _, e := range m.Issued {
			l = e.Size()
			n += 1 + l + sovToken(uint64(l))
		}
	}
	if len(m.Burned) > 0 {
		for _, e := range m.Burned {
			l = e.Size()
			n += 1 + l + sovToken(uint64(l))
		}
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Base = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issued = append(m.Issued, types.Coin{})
			if err := m.Issued[len(m.Issued)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burned", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Burned = append(m.Burned, types.Coin{})
			if err := m.Burned[len(m.Burned)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func (m *TokenDetail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: TokenDetail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenDetail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &types1.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issued = append(m.Issued, types.Coin{})
			if err := m.Issued[len(m.Issued)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burned", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Burned = append(m.Burned, types.Coin{})
			if err := m.Burned[len(m.Burned)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
