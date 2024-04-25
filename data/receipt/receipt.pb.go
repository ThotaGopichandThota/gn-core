// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: receipt.proto

package receipt

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_multiversx_mx_chain_core_go_data "github.com/ThotaGopichandThota/gn-core/data"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Receipt holds all the data needed for a transaction receipt
type Receipt struct {
	Value   *math_big.Int `protobuf:"bytes,1,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/ThotaGopichandThota/gn-core/data.BigIntCaster" json:"value"`
	SndAddr []byte        `protobuf:"bytes,2,opt,name=SndAddr,proto3" json:"sender"`
	Data    []byte        `protobuf:"bytes,3,opt,name=Data,proto3" json:"data,omitempty"`
	TxHash  []byte        `protobuf:"bytes,4,opt,name=TxHash,proto3" json:"txHash"`
}

func (m *Receipt) Reset()      { *m = Receipt{} }
func (*Receipt) ProtoMessage() {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace1d6eb38fad2c8, []int{0}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return m.Size()
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Receipt) GetSndAddr() []byte {
	if m != nil {
		return m.SndAddr
	}
	return nil
}

func (m *Receipt) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Receipt) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func init() {
	proto.RegisterType((*Receipt)(nil), "proto.Receipt")
}

func init() { proto.RegisterFile("receipt.proto", fileDescriptor_ace1d6eb38fad2c8) }

var fileDescriptor_ace1d6eb38fad2c8 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x8d, 0xbf, 0xaf, 0x3f, 0x92, 0x05, 0x0c, 0x99, 0x22, 0x86, 0x1b, 0x54, 0x21, 0xc4, 0x40,
	0x92, 0x81, 0xb1, 0x13, 0x05, 0x24, 0xba, 0x06, 0xc4, 0xc0, 0xe6, 0x24, 0x26, 0xb1, 0x54, 0xc7,
	0x95, 0x73, 0x53, 0x95, 0x8d, 0x47, 0xe0, 0x31, 0x10, 0x4f, 0xc2, 0xd8, 0xb1, 0x53, 0xa1, 0xae,
	0x90, 0x50, 0xa7, 0x3e, 0x02, 0xaa, 0x5b, 0x04, 0x2b, 0xd3, 0xf1, 0x39, 0xc7, 0x3e, 0x47, 0xbe,
	0x97, 0xee, 0x6a, 0x9e, 0x72, 0x31, 0xc4, 0x70, 0xa8, 0x15, 0x2a, 0xb7, 0x69, 0x61, 0x3f, 0xc8,
	0x05, 0x16, 0x75, 0x12, 0xa6, 0x4a, 0x46, 0xb9, 0xca, 0x55, 0x64, 0xe5, 0xa4, 0xbe, 0xb7, 0xcc,
	0x12, 0x7b, 0xda, 0xbc, 0xea, 0x7c, 0x10, 0xda, 0x8e, 0x37, 0x39, 0x6e, 0x41, 0x9b, 0xb7, 0x6c,
	0x50, 0x73, 0x8f, 0x1c, 0x90, 0xe3, 0x9d, 0x5e, 0xbc, 0x9c, 0xf9, 0xcd, 0xd1, 0x5a, 0x78, 0x79,
	0xf3, 0x2f, 0x25, 0xc3, 0x22, 0x4a, 0x44, 0x1e, 0xf6, 0x4b, 0xec, 0xfe, 0xea, 0x90, 0xf5, 0x00,
	0xc5, 0x88, 0xeb, 0x6a, 0x1c, 0xc9, 0x71, 0x90, 0x16, 0x4c, 0x94, 0x41, 0xaa, 0x34, 0x0f, 0x72,
	0x15, 0x65, 0x0c, 0x59, 0xd8, 0x13, 0x79, 0xbf, 0xc4, 0x73, 0x56, 0x21, 0xd7, 0xf1, 0xa6, 0xc0,
	0x3d, 0xa4, 0xed, 0xeb, 0x32, 0x3b, 0xcb, 0x32, 0xed, 0xfd, 0xb3, 0x5d, 0x74, 0x39, 0xf3, 0x5b,
	0x15, 0x2f, 0x33, 0xae, 0xe3, 0x6f, 0xcb, 0x3d, 0xa2, 0x8d, 0x0b, 0x86, 0xcc, 0xfb, 0x6f, 0xaf,
	0xb8, 0xcb, 0x99, 0xbf, 0xb7, 0x4e, 0x3c, 0x51, 0x52, 0x20, 0x97, 0x43, 0x7c, 0x88, 0xad, 0xef,
	0x76, 0x68, 0xeb, 0x66, 0x7c, 0xc5, 0xaa, 0xc2, 0x6b, 0xfc, 0x84, 0xa1, 0x55, 0xe2, 0xad, 0xd3,
	0xab, 0x27, 0x73, 0x70, 0xa6, 0x73, 0x70, 0x56, 0x73, 0x20, 0x8f, 0x06, 0xc8, 0xb3, 0x01, 0xf2,
	0x6a, 0x80, 0x4c, 0x0c, 0x90, 0xa9, 0x01, 0xf2, 0x6e, 0x80, 0x7c, 0x1a, 0x70, 0x56, 0x06, 0xc8,
	0xd3, 0x02, 0x9c, 0xc9, 0x02, 0x9c, 0xe9, 0x02, 0x9c, 0xbb, 0xbf, 0x7c, 0x36, 0xda, 0xae, 0xa4,
	0xbb, 0xc5, 0xa4, 0x65, 0xa7, 0x7c, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x24, 0x64, 0xda,
	0xac, 0x01, 0x00, 0x00,
}

func (this *Receipt) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Receipt)
	if !ok {
		that2, ok := that.(Receipt)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.SndAddr, that1.SndAddr) {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.TxHash, that1.TxHash) {
		return false
	}
	return true
}
func (this *Receipt) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&receipt.Receipt{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "SndAddr: "+fmt.Sprintf("%#v", this.SndAddr)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "TxHash: "+fmt.Sprintf("%#v", this.TxHash)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringReceipt(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Receipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Receipt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SndAddr) > 0 {
		i -= len(m.SndAddr)
		copy(dAtA[i:], m.SndAddr)
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.SndAddr)))
		i--
		dAtA[i] = 0x12
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintReceipt(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintReceipt(dAtA []byte, offset int, v uint64) int {
	offset -= sovReceipt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Receipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.SndAddr)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	return n
}

func sovReceipt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReceipt(x uint64) (n int) {
	return sovReceipt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Receipt) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Receipt{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`SndAddr:` + fmt.Sprintf("%v", this.SndAddr) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`TxHash:` + fmt.Sprintf("%v", this.TxHash) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringReceipt(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Receipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReceipt
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
			return fmt.Errorf("proto: Receipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SndAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SndAddr = append(m.SndAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.SndAddr == nil {
				m.SndAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReceipt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReceipt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReceipt
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
func skipReceipt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReceipt
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
					return 0, ErrIntOverflowReceipt
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
					return 0, ErrIntOverflowReceipt
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
				return 0, ErrInvalidLengthReceipt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReceipt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReceipt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReceipt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReceipt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReceipt = fmt.Errorf("proto: unexpected end of group")
)
