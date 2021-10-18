// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wasm/base_ext.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/liubaninc/sdk-go/modules/utxo/types"
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

type BaseExt struct {
	InputExts  []*contract.InputExt  `protobuf:"bytes,10,rep,name=inputExts,proto3" json:"inputExts,omitempty"`
	OutputExts []*contract.OutputExt `protobuf:"bytes,11,rep,name=outputExts,proto3" json:"outputExts,omitempty"`
}

func (m *BaseExt) Reset()         { *m = BaseExt{} }
func (m *BaseExt) String() string { return proto.CompactTextString(m) }
func (*BaseExt) ProtoMessage()    {}
func (*BaseExt) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e51919f157afee, []int{0}
}
func (m *BaseExt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseExt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseExt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseExt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseExt.Merge(m, src)
}
func (m *BaseExt) XXX_Size() int {
	return m.Size()
}
func (m *BaseExt) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseExt.DiscardUnknown(m)
}

var xxx_messageInfo_BaseExt proto.InternalMessageInfo

func (m *BaseExt) GetInputExts() []*contract.InputExt {
	if m != nil {
		return m.InputExts
	}
	return nil
}

func (m *BaseExt) GetOutputExts() []*contract.OutputExt {
	if m != nil {
		return m.OutputExts
	}
	return nil
}

// 预执行的请求结构
type InvokeRequest struct {
	ModuleName     string                    `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	ContractName   string                    `protobuf:"bytes,2,opt,name=contract_name,json=contractName,proto3" json:"contract_name,omitempty"`
	MethodName     string                    `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	Args           string                    `protobuf:"bytes,4,opt,name=args,proto3" json:"args,omitempty"`
	ResourceLimits []*contract.ResourceLimit `protobuf:"bytes,5,rep,name=resource_limits,json=resourceLimits,proto3" json:"resource_limits,omitempty"`
	// amount is the amount transfer to the contract
	// attention: In one transaction, transfer to only one contract is allowed
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *InvokeRequest) Reset()         { *m = InvokeRequest{} }
func (m *InvokeRequest) String() string { return proto.CompactTextString(m) }
func (*InvokeRequest) ProtoMessage()    {}
func (*InvokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e51919f157afee, []int{1}
}
func (m *InvokeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InvokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InvokeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InvokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeRequest.Merge(m, src)
}
func (m *InvokeRequest) XXX_Size() int {
	return m.Size()
}
func (m *InvokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeRequest proto.InternalMessageInfo

func (m *InvokeRequest) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *InvokeRequest) GetContractName() string {
	if m != nil {
		return m.ContractName
	}
	return ""
}

func (m *InvokeRequest) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *InvokeRequest) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *InvokeRequest) GetResourceLimits() []*contract.ResourceLimit {
	if m != nil {
		return m.ResourceLimits
	}
	return nil
}

func (m *InvokeRequest) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

// 预执行的返回结构
type InvokeResponse struct {
	Base      *types1.Base                 `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	BaseExt   *BaseExt                     `protobuf:"bytes,2,opt,name=baseExt,proto3" json:"baseExt,omitempty"`
	Response  [][]byte                     `protobuf:"bytes,3,rep,name=response,proto3" json:"response,omitempty"`
	GasUsed   int64                        `protobuf:"varint,4,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Requests  []*InvokeRequest             `protobuf:"bytes,5,rep,name=requests,proto3" json:"requests,omitempty"`
	Responses []*contract.ContractResponse `protobuf:"bytes,6,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (m *InvokeResponse) Reset()         { *m = InvokeResponse{} }
func (m *InvokeResponse) String() string { return proto.CompactTextString(m) }
func (*InvokeResponse) ProtoMessage()    {}
func (*InvokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38e51919f157afee, []int{2}
}
func (m *InvokeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InvokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InvokeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InvokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeResponse.Merge(m, src)
}
func (m *InvokeResponse) XXX_Size() int {
	return m.Size()
}
func (m *InvokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeResponse proto.InternalMessageInfo

func (m *InvokeResponse) GetBase() *types1.Base {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InvokeResponse) GetBaseExt() *BaseExt {
	if m != nil {
		return m.BaseExt
	}
	return nil
}

func (m *InvokeResponse) GetResponse() [][]byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *InvokeResponse) GetGasUsed() int64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *InvokeResponse) GetRequests() []*InvokeRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *InvokeResponse) GetResponses() []*contract.ContractResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseExt)(nil), "liubaninc.m1.wasm.BaseExt")
	proto.RegisterType((*InvokeRequest)(nil), "liubaninc.m1.wasm.InvokeRequest")
	proto.RegisterType((*InvokeResponse)(nil), "liubaninc.m1.wasm.InvokeResponse")
}

func init() { proto.RegisterFile("wasm/base_ext.proto", fileDescriptor_38e51919f157afee) }

var fileDescriptor_38e51919f157afee = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x8e, 0xd3, 0x3c,
	0x14, 0x6d, 0xda, 0x7e, 0xed, 0xd4, 0x9d, 0x1f, 0x7d, 0x1e, 0x24, 0x32, 0x95, 0x48, 0xab, 0x22,
	0xa4, 0x4a, 0x68, 0x1c, 0x5a, 0xd8, 0x81, 0x58, 0xb4, 0x1a, 0x89, 0x11, 0x08, 0x24, 0x4b, 0x6c,
	0xd8, 0x54, 0x4e, 0x62, 0x65, 0xa2, 0x36, 0x71, 0xc9, 0x75, 0x86, 0xf2, 0x02, 0xac, 0x79, 0x0e,
	0x24, 0xde, 0x80, 0x07, 0x98, 0xe5, 0x2c, 0x59, 0x01, 0x6a, 0x5f, 0x04, 0xe5, 0xc6, 0xe9, 0xb4,
	0xa2, 0xab, 0xc4, 0xd7, 0xe7, 0xdc, 0x6b, 0x9f, 0x73, 0x4c, 0x4e, 0x3f, 0x09, 0x88, 0x5d, 0x4f,
	0x80, 0x9c, 0xca, 0xa5, 0x66, 0x8b, 0x54, 0x69, 0x45, 0xff, 0x9f, 0x47, 0x99, 0x27, 0x92, 0x28,
	0xf1, 0x59, 0x3c, 0x64, 0x39, 0xa2, 0x73, 0x2f, 0x54, 0xa1, 0xc2, 0x5d, 0x37, 0xff, 0x2b, 0x80,
	0x9d, 0xd3, 0x65, 0xac, 0x02, 0x39, 0x77, 0x8b, 0x8f, 0x29, 0x9e, 0x64, 0x7a, 0xa9, 0xb0, 0xa5,
	0x29, 0x38, 0xbe, 0x82, 0x58, 0x01, 0x96, 0xdc, 0xeb, 0xa1, 0x27, 0xb5, 0x18, 0xba, 0xbe, 0x8a,
	0x92, 0x62, 0xbf, 0xff, 0xc5, 0x22, 0xcd, 0xb1, 0x00, 0x79, 0xb1, 0xd4, 0xf4, 0x39, 0x69, 0x45,
	0xc9, 0x22, 0xd3, 0x17, 0x4b, 0x0d, 0x36, 0xe9, 0xd5, 0x06, 0xed, 0xd1, 0x03, 0xb6, 0x73, 0x1c,
	0x33, 0xeb, 0xd2, 0xa0, 0xf8, 0x1d, 0x9e, 0xbe, 0x24, 0x44, 0x65, 0xba, 0x64, 0xb7, 0x91, 0xed,
	0xec, 0x65, 0xbf, 0x2b, 0x61, 0x7c, 0x8b, 0xd1, 0xff, 0x51, 0x25, 0x47, 0x97, 0xc9, 0xb5, 0x9a,
	0x49, 0x2e, 0x3f, 0x66, 0x12, 0x34, 0xed, 0x92, 0x76, 0xac, 0x82, 0x6c, 0x2e, 0xa7, 0x89, 0x88,
	0xa5, 0x6d, 0xf5, 0xac, 0x41, 0x8b, 0x93, 0xa2, 0xf4, 0x56, 0xc4, 0x92, 0x3e, 0x24, 0x47, 0xbe,
	0x4a, 0x74, 0x2a, 0x7c, 0x5d, 0x40, 0xaa, 0x08, 0x39, 0x2c, 0x8b, 0x08, 0xca, 0xbb, 0x48, 0x7d,
	0xa5, 0x82, 0x02, 0x52, 0x33, 0x5d, 0xb0, 0x84, 0x00, 0x4a, 0xea, 0x22, 0x0d, 0xc1, 0xae, 0xe3,
	0x0e, 0xfe, 0xd3, 0xd7, 0xe4, 0x24, 0x95, 0xa0, 0xb2, 0xd4, 0x97, 0xd3, 0x79, 0x14, 0x47, 0x1a,
	0xec, 0xff, 0xf0, 0x46, 0xfd, 0xbd, 0x37, 0xe2, 0x06, 0xfb, 0x26, 0x87, 0xf2, 0xe3, 0x74, 0x7b,
	0x09, 0xd4, 0x27, 0x0d, 0x11, 0xab, 0x2c, 0xd1, 0x76, 0x03, 0x7b, 0x9c, 0xb1, 0xc2, 0x13, 0x86,
	0x36, 0x19, 0x4f, 0xd8, 0x44, 0x45, 0xc9, 0xf8, 0xc9, 0xcd, 0xaf, 0x6e, 0xe5, 0xdb, 0xef, 0xee,
	0x20, 0x8c, 0xf4, 0x55, 0xe6, 0x31, 0x5f, 0xc5, 0xae, 0x31, 0xb0, 0xf8, 0x9c, 0x43, 0x30, 0x73,
	0xf5, 0xe7, 0x85, 0x04, 0x24, 0x00, 0x37, 0xad, 0xfb, 0xdf, 0xab, 0xe4, 0xb8, 0x94, 0x0f, 0x16,
	0x2a, 0x01, 0x49, 0x1f, 0x93, 0x7a, 0x3e, 0x01, 0x85, 0x6b, 0x8f, 0xee, 0xef, 0x9e, 0x3c, 0xcf,
	0x09, 0xcb, 0x8d, 0xe7, 0x08, 0xa2, 0xcf, 0x48, 0xd3, 0x2b, 0x62, 0x80, 0x2a, 0xb6, 0x47, 0x1d,
	0xf6, 0x4f, 0x10, 0x99, 0x09, 0x0a, 0x2f, 0xa1, 0xb4, 0x43, 0x0e, 0x52, 0x33, 0xce, 0xae, 0xf5,
	0x6a, 0x83, 0x43, 0xbe, 0x59, 0xd3, 0x33, 0x72, 0x10, 0x0a, 0x98, 0x66, 0x20, 0x03, 0xd4, 0xb6,
	0xc6, 0x9b, 0xa1, 0x80, 0xf7, 0x20, 0x03, 0xfa, 0x22, 0xa7, 0xa1, 0xc9, 0xa5, 0xae, 0xbd, 0x3d,
	0xd3, 0x76, 0xd2, 0xc0, 0x37, 0x0c, 0x3a, 0x21, 0xad, 0x72, 0x08, 0x18, 0x49, 0x1f, 0xed, 0xb5,
	0x65, 0x62, 0x72, 0x50, 0x2a, 0xc2, 0xef, 0x78, 0xe3, 0x57, 0x37, 0x2b, 0xc7, 0xba, 0x5d, 0x39,
	0xd6, 0x9f, 0x95, 0x63, 0x7d, 0x5d, 0x3b, 0x95, 0xdb, 0xb5, 0x53, 0xf9, 0xb9, 0x76, 0x2a, 0x1f,
	0xd8, 0x96, 0xf6, 0x9b, 0xae, 0x2e, 0x04, 0xb3, 0xf3, 0x50, 0xb9, 0x45, 0xfa, 0xc0, 0xc5, 0x97,
	0x8b, 0x3e, 0x78, 0x0d, 0x7c, 0x48, 0x4f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x67, 0xd2, 0xf7,
	0x27, 0xce, 0x03, 0x00, 0x00,
}

func (m *BaseExt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseExt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseExt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OutputExts) > 0 {
		for iNdEx := len(m.OutputExts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutputExts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBaseExt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.InputExts) > 0 {
		for iNdEx := len(m.InputExts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InputExts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBaseExt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	return len(dAtA) - i, nil
}

func (m *InvokeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InvokeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InvokeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBaseExt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ResourceLimits) > 0 {
		for iNdEx := len(m.ResourceLimits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourceLimits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBaseExt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Args) > 0 {
		i -= len(m.Args)
		copy(dAtA[i:], m.Args)
		i = encodeVarintBaseExt(dAtA, i, uint64(len(m.Args)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MethodName) > 0 {
		i -= len(m.MethodName)
		copy(dAtA[i:], m.MethodName)
		i = encodeVarintBaseExt(dAtA, i, uint64(len(m.MethodName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractName) > 0 {
		i -= len(m.ContractName)
		copy(dAtA[i:], m.ContractName)
		i = encodeVarintBaseExt(dAtA, i, uint64(len(m.ContractName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintBaseExt(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InvokeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InvokeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InvokeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Responses) > 0 {
		for iNdEx := len(m.Responses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Responses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBaseExt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBaseExt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.GasUsed != 0 {
		i = encodeVarintBaseExt(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Response) > 0 {
		for iNdEx := len(m.Response) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Response[iNdEx])
			copy(dAtA[i:], m.Response[iNdEx])
			i = encodeVarintBaseExt(dAtA, i, uint64(len(m.Response[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.BaseExt != nil {
		{
			size, err := m.BaseExt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBaseExt(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Base != nil {
		{
			size, err := m.Base.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBaseExt(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBaseExt(dAtA []byte, offset int, v uint64) int {
	offset -= sovBaseExt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseExt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.InputExts) > 0 {
		for _, e := range m.InputExts {
			l = e.Size()
			n += 1 + l + sovBaseExt(uint64(l))
		}
	}
	if len(m.OutputExts) > 0 {
		for _, e := range m.OutputExts {
			l = e.Size()
			n += 1 + l + sovBaseExt(uint64(l))
		}
	}
	return n
}

func (m *InvokeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovBaseExt(uint64(l))
	}
	l = len(m.ContractName)
	if l > 0 {
		n += 1 + l + sovBaseExt(uint64(l))
	}
	l = len(m.MethodName)
	if l > 0 {
		n += 1 + l + sovBaseExt(uint64(l))
	}
	l = len(m.Args)
	if l > 0 {
		n += 1 + l + sovBaseExt(uint64(l))
	}
	if len(m.ResourceLimits) > 0 {
		for _, e := range m.ResourceLimits {
			l = e.Size()
			n += 1 + l + sovBaseExt(uint64(l))
		}
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovBaseExt(uint64(l))
		}
	}
	return n
}

func (m *InvokeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Base != nil {
		l = m.Base.Size()
		n += 1 + l + sovBaseExt(uint64(l))
	}
	if m.BaseExt != nil {
		l = m.BaseExt.Size()
		n += 1 + l + sovBaseExt(uint64(l))
	}
	if len(m.Response) > 0 {
		for _, b := range m.Response {
			l = len(b)
			n += 1 + l + sovBaseExt(uint64(l))
		}
	}
	if m.GasUsed != 0 {
		n += 1 + sovBaseExt(uint64(m.GasUsed))
	}
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovBaseExt(uint64(l))
		}
	}
	if len(m.Responses) > 0 {
		for _, e := range m.Responses {
			l = e.Size()
			n += 1 + l + sovBaseExt(uint64(l))
		}
	}
	return n
}

func sovBaseExt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBaseExt(x uint64) (n int) {
	return sovBaseExt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseExt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBaseExt
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
			return fmt.Errorf("proto: BaseExt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseExt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputExts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InputExts = append(m.InputExts, &contract.InputExt{})
			if err := m.InputExts[len(m.InputExts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputExts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutputExts = append(m.OutputExts, &contract.OutputExt{})
			if err := m.OutputExts[len(m.OutputExts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBaseExt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBaseExt
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
func (m *InvokeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBaseExt
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
			return fmt.Errorf("proto: InvokeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InvokeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceLimits = append(m.ResourceLimits, &contract.ResourceLimit{})
			if err := m.ResourceLimits[len(m.ResourceLimits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBaseExt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBaseExt
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
func (m *InvokeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBaseExt
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
			return fmt.Errorf("proto: InvokeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InvokeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Base == nil {
				m.Base = &types1.Base{}
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseExt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseExt == nil {
				m.BaseExt = &BaseExt{}
			}
			if err := m.BaseExt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = append(m.Response, make([]byte, postIndex-iNdEx))
			copy(m.Response[len(m.Response)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, &InvokeRequest{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Responses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseExt
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
				return ErrInvalidLengthBaseExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Responses = append(m.Responses, &contract.ContractResponse{})
			if err := m.Responses[len(m.Responses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBaseExt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBaseExt
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
func skipBaseExt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBaseExt
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
					return 0, ErrIntOverflowBaseExt
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
					return 0, ErrIntOverflowBaseExt
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
				return 0, ErrInvalidLengthBaseExt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBaseExt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBaseExt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBaseExt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBaseExt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBaseExt = fmt.Errorf("proto: unexpected end of group")
)