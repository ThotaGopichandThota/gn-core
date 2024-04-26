// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metrics.proto

package metrics

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
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

type Metric struct {
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Metric_ValUint64
	//	*Metric_ValString
	Value isMetric_Value `protobuf_oneof:"Value"`
}

func (m *Metric) Reset()      { *m = Metric{} }
func (*Metric) ProtoMessage() {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}
func (m *Metric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return m.Size()
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

type isMetric_Value interface {
	isMetric_Value()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type Metric_ValUint64 struct {
	ValUint64 uint64 `protobuf:"varint,2,opt,name=ValUint64,proto3,oneof" json:"ValUint64,omitempty"`
}
type Metric_ValString struct {
	ValString string `protobuf:"bytes,3,opt,name=ValString,proto3,oneof" json:"ValString,omitempty"`
}

func (*Metric_ValUint64) isMetric_Value() {}
func (*Metric_ValString) isMetric_Value() {}

func (m *Metric) GetValue() isMetric_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Metric) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Metric) GetValUint64() uint64 {
	if x, ok := m.GetValue().(*Metric_ValUint64); ok {
		return x.ValUint64
	}
	return 0
}

func (m *Metric) GetValString() string {
	if x, ok := m.GetValue().(*Metric_ValString); ok {
		return x.ValString
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Metric) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Metric_ValUint64)(nil),
		(*Metric_ValString)(nil),
	}
}

type MetricsList struct {
	Metrics []Metric `protobuf:"bytes,1,rep,name=Metrics,proto3" json:"Metrics"`
}

func (m *MetricsList) Reset()      { *m = MetricsList{} }
func (*MetricsList) ProtoMessage() {}
func (*MetricsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{1}
}
func (m *MetricsList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetricsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MetricsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsList.Merge(m, src)
}
func (m *MetricsList) XXX_Size() int {
	return m.Size()
}
func (m *MetricsList) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsList.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsList proto.InternalMessageInfo

func (m *MetricsList) GetMetrics() []Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Metric)(nil), "metrics.Metric")
	proto.RegisterType((*MetricsList)(nil), "metrics.MetricsList")
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72) }

var fileDescriptor_6039342a2ba47b72 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x29,
	0xca, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x74, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1,
	0xf2, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xf4, 0x29, 0x25, 0x73, 0xb1, 0xf9,
	0x82, 0x75, 0x0a, 0x09, 0x70, 0x31, 0x7b, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x81, 0x98, 0x42, 0x72, 0x5c, 0x9c, 0x61, 0x89, 0x39, 0xa1, 0x99, 0x79, 0x25, 0x66, 0x26, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x2c, 0x1e, 0x0c, 0x41, 0x08, 0x21, 0xa8, 0x7c, 0x70, 0x49, 0x51, 0x66,
	0x5e, 0xba, 0x04, 0x33, 0x48, 0x1f, 0x54, 0x1e, 0x22, 0xe4, 0xc4, 0xce, 0xc5, 0x1a, 0x96, 0x98,
	0x53, 0x9a, 0xaa, 0x64, 0xc7, 0xc5, 0x0d, 0xb1, 0xa4, 0xd8, 0x27, 0xb3, 0xb8, 0x44, 0x48, 0x9f,
	0x8b, 0x1d, 0xca, 0x95, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd7, 0x83, 0x79, 0x06, 0x22,
	0xee, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x4c, 0x95, 0x93, 0xe3, 0x85, 0x87, 0x72, 0x0c,
	0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39,
	0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0xbc, 0xf1, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x17, 0x8f, 0xe4, 0x18, 0x3e, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x60, 0xc1, 0x92, 0xc4, 0x06, 0xf6, 0xae, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x7f, 0x2e, 0x70, 0x37, 0x01, 0x00, 0x00,
}

func (this *Metric) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metric)
	if !ok {
		that2, ok := that.(Metric)
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
	if this.Key != that1.Key {
		return false
	}
	if that1.Value == nil {
		if this.Value != nil {
			return false
		}
	} else if this.Value == nil {
		return false
	} else if !this.Value.Equal(that1.Value) {
		return false
	}
	return true
}
func (this *Metric_ValUint64) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metric_ValUint64)
	if !ok {
		that2, ok := that.(Metric_ValUint64)
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
	if this.ValUint64 != that1.ValUint64 {
		return false
	}
	return true
}
func (this *Metric_ValString) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metric_ValString)
	if !ok {
		that2, ok := that.(Metric_ValString)
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
	if this.ValString != that1.ValString {
		return false
	}
	return true
}
func (this *MetricsList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetricsList)
	if !ok {
		that2, ok := that.(MetricsList)
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
	if len(this.Metrics) != len(that1.Metrics) {
		return false
	}
	for i := range this.Metrics {
		if !this.Metrics[i].Equal(&that1.Metrics[i]) {
			return false
		}
	}
	return true
}
func (this *Metric) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&metrics.Metric{")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	if this.Value != nil {
		s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Metric_ValUint64) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&metrics.Metric_ValUint64{` +
		`ValUint64:` + fmt.Sprintf("%#v", this.ValUint64) + `}`}, ", ")
	return s
}
func (this *Metric_ValString) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&metrics.Metric_ValString{` +
		`ValString:` + fmt.Sprintf("%#v", this.ValString) + `}`}, ", ")
	return s
}
func (this *MetricsList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&metrics.MetricsList{")
	if this.Metrics != nil {
		vs := make([]Metric, len(this.Metrics))
		for i := range vs {
			vs[i] = this.Metrics[i]
		}
		s = append(s, "Metrics: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMetrics(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Metric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metric) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metric) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metric_ValUint64) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metric_ValUint64) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintMetrics(dAtA, i, uint64(m.ValUint64))
	i--
	dAtA[i] = 0x10
	return len(dAtA) - i, nil
}
func (m *Metric_ValString) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metric_ValString) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ValString)
	copy(dAtA[i:], m.ValString)
	i = encodeVarintMetrics(dAtA, i, uint64(len(m.ValString)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *MetricsList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricsList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetricsList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for iNdEx := len(m.Metrics) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Metrics[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetrics(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetrics(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetrics(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Metric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *Metric_ValUint64) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMetrics(uint64(m.ValUint64))
	return n
}
func (m *Metric_ValString) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValString)
	n += 1 + l + sovMetrics(uint64(l))
	return n
}
func (m *MetricsList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Metrics) > 0 {
		for _, e := range m.Metrics {
			l = e.Size()
			n += 1 + l + sovMetrics(uint64(l))
		}
	}
	return n
}

func sovMetrics(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetrics(x uint64) (n int) {
	return sovMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Metric) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Metric{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Metric_ValUint64) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Metric_ValUint64{`,
		`ValUint64:` + fmt.Sprintf("%v", this.ValUint64) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Metric_ValString) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Metric_ValString{`,
		`ValString:` + fmt.Sprintf("%v", this.ValString) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MetricsList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForMetrics := "[]Metric{"
	for _, f := range this.Metrics {
		repeatedStringForMetrics += strings.Replace(strings.Replace(f.String(), "Metric", "Metric", 1), `&`, ``, 1) + ","
	}
	repeatedStringForMetrics += "}"
	s := strings.Join([]string{`&MetricsList{`,
		`Metrics:` + repeatedStringForMetrics + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMetrics(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Metric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
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
			return fmt.Errorf("proto: Metric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValUint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Metric_ValUint64{v}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Metric_ValString{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
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
func (m *MetricsList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
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
			return fmt.Errorf("proto: MetricsList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricsList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = append(m.Metrics, Metric{})
			if err := m.Metrics[len(m.Metrics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetrics
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
func skipMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetrics
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
					return 0, ErrIntOverflowMetrics
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
					return 0, ErrIntOverflowMetrics
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
				return 0, ErrInvalidLengthMetrics
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetrics
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetrics
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetrics        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetrics          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetrics = fmt.Errorf("proto: unexpected end of group")
)
