// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wasm/contract_code.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	contract "github.com/liubaninc/sdk-go/modules/wasm/xmodel/contract"
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

type ContractCode_Status int32

const (
	ContractCode_SUBMIT  ContractCode_Status = 0
	ContractCode_APPROVE ContractCode_Status = 1
)

var ContractCode_Status_name = map[int32]string{
	0: "SUBMIT",
	1: "APPROVE",
}

var ContractCode_Status_value = map[string]int32{
	"SUBMIT":  0,
	"APPROVE": 1,
}

func (x ContractCode_Status) String() string {
	return proto.EnumName(ContractCode_Status_name, int32(x))
}

func (ContractCode_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f16d30825fee4598, []int{0, 0}
}

type ContractCode struct {
	Creator   string                 `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Hash      string                 `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Code      []byte                 `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Src       string                 `protobuf:"bytes,5,opt,name=src,proto3" json:"src,omitempty"`
	Abi       string                 `protobuf:"bytes,6,opt,name=abi,proto3" json:"abi,omitempty"`
	Desc      *contract.WasmCodeDesc `protobuf:"bytes,7,opt,name=desc,proto3" json:"desc,omitempty"`
	Status    ContractCode_Status    `protobuf:"varint,14,opt,name=status,proto3,enum=liubaninc.m1.wasm.ContractCode_Status" json:"status,omitempty"`
	Approvers []string               `protobuf:"bytes,15,rep,name=approvers,proto3" json:"approvers,omitempty"`
	Approved  []string               `protobuf:"bytes,16,rep,name=approved,proto3" json:"approved,omitempty"`
}

func (m *ContractCode) Reset()         { *m = ContractCode{} }
func (m *ContractCode) String() string { return proto.CompactTextString(m) }
func (*ContractCode) ProtoMessage()    {}
func (*ContractCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_f16d30825fee4598, []int{0}
}
func (m *ContractCode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractCode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCode.Merge(m, src)
}
func (m *ContractCode) XXX_Size() int {
	return m.Size()
}
func (m *ContractCode) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCode.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCode proto.InternalMessageInfo

func (m *ContractCode) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ContractCode) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ContractCode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContractCode) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ContractCode) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *ContractCode) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *ContractCode) GetDesc() *contract.WasmCodeDesc {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *ContractCode) GetStatus() ContractCode_Status {
	if m != nil {
		return m.Status
	}
	return ContractCode_SUBMIT
}

func (m *ContractCode) GetApprovers() []string {
	if m != nil {
		return m.Approvers
	}
	return nil
}

func (m *ContractCode) GetApproved() []string {
	if m != nil {
		return m.Approved
	}
	return nil
}

func init() {
	proto.RegisterEnum("liubaninc.m1.wasm.ContractCode_Status", ContractCode_Status_name, ContractCode_Status_value)
	proto.RegisterType((*ContractCode)(nil), "liubaninc.m1.wasm.ContractCode")
}

func init() { proto.RegisterFile("wasm/contract_code.proto", fileDescriptor_f16d30825fee4598) }

var fileDescriptor_f16d30825fee4598 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbd, 0x6e, 0xdb, 0x30,
	0x14, 0x85, 0x45, 0xdb, 0x95, 0x6b, 0xda, 0x70, 0x55, 0xb6, 0x03, 0x61, 0x14, 0x82, 0xec, 0xa1,
	0xd0, 0x52, 0x0a, 0x75, 0xd1, 0xb5, 0x40, 0xed, 0x16, 0x68, 0x87, 0xa2, 0x86, 0x9c, 0x1f, 0x20,
	0x4b, 0x40, 0x51, 0x84, 0x2c, 0xc4, 0x12, 0x05, 0x91, 0xca, 0xcf, 0x5b, 0xe4, 0xb1, 0x32, 0x7a,
	0xcc, 0x90, 0x21, 0xb0, 0x5f, 0x24, 0x20, 0xa5, 0x38, 0x09, 0x32, 0xe9, 0xdc, 0xef, 0x1e, 0x5d,
	0x9c, 0xcb, 0x0b, 0xf1, 0x05, 0x95, 0x59, 0xc0, 0x44, 0xae, 0x4a, 0xca, 0xd4, 0x29, 0x13, 0x31,
	0x27, 0x45, 0x29, 0x94, 0x40, 0xef, 0xd7, 0x69, 0x15, 0xd1, 0x3c, 0xcd, 0x19, 0xc9, 0xbe, 0x12,
	0x6d, 0x1b, 0x7d, 0x4c, 0x44, 0x22, 0x4c, 0x37, 0xd0, 0xaa, 0x36, 0x8e, 0x3e, 0x5c, 0x66, 0x22,
	0xe6, 0xeb, 0xa0, 0xfe, 0xd4, 0x70, 0x72, 0xd7, 0x82, 0x83, 0x79, 0x33, 0x75, 0x2e, 0x62, 0x8e,
	0x30, 0xec, 0xb2, 0x92, 0x53, 0x25, 0x4a, 0x0c, 0x3c, 0xe0, 0xf7, 0xc2, 0xc7, 0x12, 0x21, 0xd8,
	0x59, 0x51, 0xb9, 0xc2, 0x2d, 0x83, 0x8d, 0xd6, 0x2c, 0xa7, 0x19, 0xc7, 0xed, 0x9a, 0x69, 0xad,
	0x99, 0x8e, 0x87, 0x3b, 0x1e, 0xf0, 0x07, 0xa1, 0xd1, 0xc8, 0x81, 0x6d, 0x59, 0x32, 0xfc, 0xc6,
	0xd8, 0xb4, 0xd4, 0x84, 0x46, 0x29, 0xb6, 0x6b, 0x42, 0xa3, 0x14, 0x7d, 0x87, 0x9d, 0x98, 0x4b,
	0x86, 0xbb, 0x1e, 0xf0, 0xfb, 0xd3, 0x31, 0x79, 0xb1, 0x57, 0x13, 0xfa, 0x98, 0xca, 0x4c, 0xc7,
	0xfc, 0xc5, 0x25, 0x0b, 0x8d, 0x1d, 0xfd, 0x80, 0xb6, 0x54, 0x54, 0x55, 0x12, 0x0f, 0x3d, 0xe0,
	0x0f, 0xa7, 0x9f, 0xc9, 0xab, 0x07, 0x21, 0xcf, 0x37, 0x24, 0x4b, 0xe3, 0x0e, 0x9b, 0xbf, 0xd0,
	0x27, 0xd8, 0xa3, 0x45, 0x51, 0x8a, 0x73, 0x5e, 0x4a, 0xfc, 0xce, 0x6b, 0xfb, 0xbd, 0xf0, 0x09,
	0xa0, 0x11, 0x7c, 0xdb, 0x14, 0x31, 0x76, 0x4c, 0x73, 0x5f, 0x4f, 0xc6, 0xd0, 0xae, 0x67, 0x21,
	0x08, 0xed, 0xe5, 0xe1, 0xec, 0xdf, 0xdf, 0x03, 0xc7, 0x42, 0x7d, 0xd8, 0xfd, 0xb9, 0x58, 0x84,
	0xff, 0x8f, 0x7e, 0x3b, 0x60, 0xf6, 0xe7, 0x66, 0xeb, 0x82, 0xcd, 0xd6, 0x05, 0xf7, 0x5b, 0x17,
	0x5c, 0xef, 0x5c, 0x6b, 0xb3, 0x73, 0xad, 0xdb, 0x9d, 0x6b, 0x9d, 0x90, 0x24, 0x55, 0xab, 0x2a,
	0x22, 0x4c, 0x64, 0xc1, 0x3e, 0x70, 0x20, 0xe3, 0xb3, 0x2f, 0x89, 0x08, 0x32, 0x11, 0x57, 0x6b,
	0x2e, 0x03, 0x73, 0x74, 0x75, 0x55, 0x70, 0x19, 0xd9, 0xe6, 0x5e, 0xdf, 0x1e, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x3d, 0x81, 0xac, 0xc1, 0x09, 0x02, 0x00, 0x00,
}

func (m *ContractCode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractCode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractCode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Approved) > 0 {
		for iNdEx := len(m.Approved) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Approved[iNdEx])
			copy(dAtA[i:], m.Approved[iNdEx])
			i = encodeVarintContractCode(dAtA, i, uint64(len(m.Approved[iNdEx])))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if len(m.Approvers) > 0 {
		for iNdEx := len(m.Approvers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Approvers[iNdEx])
			copy(dAtA[i:], m.Approvers[iNdEx])
			i = encodeVarintContractCode(dAtA, i, uint64(len(m.Approvers[iNdEx])))
			i--
			dAtA[i] = 0x7a
		}
	}
	if m.Status != 0 {
		i = encodeVarintContractCode(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x70
	}
	if m.Desc != nil {
		{
			size, err := m.Desc.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintContractCode(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Abi) > 0 {
		i -= len(m.Abi)
		copy(dAtA[i:], m.Abi)
		i = encodeVarintContractCode(dAtA, i, uint64(len(m.Abi)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Src) > 0 {
		i -= len(m.Src)
		copy(dAtA[i:], m.Src)
		i = encodeVarintContractCode(dAtA, i, uint64(len(m.Src)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintContractCode(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintContractCode(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintContractCode(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintContractCode(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContractCode(dAtA []byte, offset int, v uint64) int {
	offset -= sovContractCode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractCode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovContractCode(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovContractCode(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovContractCode(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovContractCode(uint64(l))
	}
	l = len(m.Src)
	if l > 0 {
		n += 1 + l + sovContractCode(uint64(l))
	}
	l = len(m.Abi)
	if l > 0 {
		n += 1 + l + sovContractCode(uint64(l))
	}
	if m.Desc != nil {
		l = m.Desc.Size()
		n += 1 + l + sovContractCode(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovContractCode(uint64(m.Status))
	}
	if len(m.Approvers) > 0 {
		for _, s := range m.Approvers {
			l = len(s)
			n += 1 + l + sovContractCode(uint64(l))
		}
	}
	if len(m.Approved) > 0 {
		for _, s := range m.Approved {
			l = len(s)
			n += 2 + l + sovContractCode(uint64(l))
		}
	}
	return n
}

func sovContractCode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContractCode(x uint64) (n int) {
	return sovContractCode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractCode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContractCode
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
			return fmt.Errorf("proto: ContractCode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractCode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = append(m.Code[:0], dAtA[iNdEx:postIndex]...)
			if m.Code == nil {
				m.Code = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Src", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Src = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abi", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Abi = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Desc == nil {
				m.Desc = &contract.WasmCodeDesc{}
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ContractCode_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvers = append(m.Approvers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCode
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
				return ErrInvalidLengthContractCode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approved = append(m.Approved, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContractCode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContractCode
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
func skipContractCode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContractCode
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
					return 0, ErrIntOverflowContractCode
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
					return 0, ErrIntOverflowContractCode
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
				return 0, ErrInvalidLengthContractCode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContractCode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContractCode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContractCode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContractCode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContractCode = fmt.Errorf("proto: unexpected end of group")
)
