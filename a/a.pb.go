// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: a/a.proto

package a

import (
	fmt "fmt"
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

type MsgPousljlwlybJloivfeuc struct {
	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B []byte `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C uint64 `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D string `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	E int64  `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
}

func (m *MsgPousljlwlybJloivfeuc) Reset()         { *m = MsgPousljlwlybJloivfeuc{} }
func (m *MsgPousljlwlybJloivfeuc) String() string { return proto.CompactTextString(m) }
func (*MsgPousljlwlybJloivfeuc) ProtoMessage()    {}
func (*MsgPousljlwlybJloivfeuc) Descriptor() ([]byte, []int) {
	return fileDescriptor_b544b42e4fa5387d, []int{0}
}
func (m *MsgPousljlwlybJloivfeuc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPousljlwlybJloivfeuc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPousljlwlybJloivfeuc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPousljlwlybJloivfeuc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPousljlwlybJloivfeuc.Merge(m, src)
}
func (m *MsgPousljlwlybJloivfeuc) XXX_Size() int {
	return m.Size()
}
func (m *MsgPousljlwlybJloivfeuc) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPousljlwlybJloivfeuc.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPousljlwlybJloivfeuc proto.InternalMessageInfo

func (m *MsgPousljlwlybJloivfeuc) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *MsgPousljlwlybJloivfeuc) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *MsgPousljlwlybJloivfeuc) GetC() uint64 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *MsgPousljlwlybJloivfeuc) GetD() string {
	if m != nil {
		return m.D
	}
	return ""
}

func (m *MsgPousljlwlybJloivfeuc) GetE() int64 {
	if m != nil {
		return m.E
	}
	return 0
}

type MsgMrtsdxfVexxblyqrzg struct {
	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B []byte `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C uint64 `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D string `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	E int64  `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
}

func (m *MsgMrtsdxfVexxblyqrzg) Reset()         { *m = MsgMrtsdxfVexxblyqrzg{} }
func (m *MsgMrtsdxfVexxblyqrzg) String() string { return proto.CompactTextString(m) }
func (*MsgMrtsdxfVexxblyqrzg) ProtoMessage()    {}
func (*MsgMrtsdxfVexxblyqrzg) Descriptor() ([]byte, []int) {
	return fileDescriptor_b544b42e4fa5387d, []int{1}
}
func (m *MsgMrtsdxfVexxblyqrzg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMrtsdxfVexxblyqrzg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMrtsdxfVexxblyqrzg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMrtsdxfVexxblyqrzg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMrtsdxfVexxblyqrzg.Merge(m, src)
}
func (m *MsgMrtsdxfVexxblyqrzg) XXX_Size() int {
	return m.Size()
}
func (m *MsgMrtsdxfVexxblyqrzg) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMrtsdxfVexxblyqrzg.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMrtsdxfVexxblyqrzg proto.InternalMessageInfo

func (m *MsgMrtsdxfVexxblyqrzg) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *MsgMrtsdxfVexxblyqrzg) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *MsgMrtsdxfVexxblyqrzg) GetC() uint64 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *MsgMrtsdxfVexxblyqrzg) GetD() string {
	if m != nil {
		return m.D
	}
	return ""
}

func (m *MsgMrtsdxfVexxblyqrzg) GetE() int64 {
	if m != nil {
		return m.E
	}
	return 0
}

type MsgLpctqrWokmmwtlm struct {
	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B []byte `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C uint64 `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D string `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	E int64  `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
}

func (m *MsgLpctqrWokmmwtlm) Reset()         { *m = MsgLpctqrWokmmwtlm{} }
func (m *MsgLpctqrWokmmwtlm) String() string { return proto.CompactTextString(m) }
func (*MsgLpctqrWokmmwtlm) ProtoMessage()    {}
func (*MsgLpctqrWokmmwtlm) Descriptor() ([]byte, []int) {
	return fileDescriptor_b544b42e4fa5387d, []int{2}
}
func (m *MsgLpctqrWokmmwtlm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLpctqrWokmmwtlm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLpctqrWokmmwtlm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLpctqrWokmmwtlm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLpctqrWokmmwtlm.Merge(m, src)
}
func (m *MsgLpctqrWokmmwtlm) XXX_Size() int {
	return m.Size()
}
func (m *MsgLpctqrWokmmwtlm) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLpctqrWokmmwtlm.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLpctqrWokmmwtlm proto.InternalMessageInfo

func (m *MsgLpctqrWokmmwtlm) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *MsgLpctqrWokmmwtlm) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *MsgLpctqrWokmmwtlm) GetC() uint64 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *MsgLpctqrWokmmwtlm) GetD() string {
	if m != nil {
		return m.D
	}
	return ""
}

func (m *MsgLpctqrWokmmwtlm) GetE() int64 {
	if m != nil {
		return m.E
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgPousljlwlybJloivfeuc)(nil), "tobrqchnip.pqvmqjx.hmeluptxh.MsgPousljlwlybJloivfeuc")
	proto.RegisterType((*MsgMrtsdxfVexxblyqrzg)(nil), "tobrqchnip.pqvmqjx.hmeluptxh.MsgMrtsdxfVexxblyqrzg")
	proto.RegisterType((*MsgLpctqrWokmmwtlm)(nil), "tobrqchnip.pqvmqjx.hmeluptxh.MsgLpctqrWokmmwtlm")
}

func init() { proto.RegisterFile("a/a.proto", fileDescriptor_b544b42e4fa5387d) }

var fileDescriptor_b544b42e4fa5387d = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0x31, 0x6a, 0xb4, 0x40,
	0x1c, 0x47, 0xfd, 0x7f, 0xbb, 0x5f, 0x60, 0x65, 0x2b, 0x21, 0xc4, 0x22, 0x0c, 0xb2, 0x95, 0x8d,
	0x5a, 0xe4, 0x06, 0xa9, 0x42, 0x88, 0x10, 0x2c, 0x12, 0x58, 0x08, 0x61, 0x66, 0x9c, 0x1d, 0x5d,
	0x67, 0x1c, 0x1d, 0xc7, 0x55, 0x73, 0x8a, 0x1c, 0x2b, 0xe5, 0x96, 0x29, 0x83, 0x5e, 0x24, 0x68,
	0xae, 0xb0, 0xe5, 0xfb, 0xf1, 0x78, 0xc5, 0xcf, 0xde, 0xe0, 0x08, 0x87, 0x95, 0x56, 0x46, 0x39,
	0xb7, 0x46, 0x11, 0x5d, 0xd3, 0xac, 0xcc, 0xab, 0xb0, 0xaa, 0x4f, 0xb2, 0x3e, 0xf6, 0x61, 0x26,
	0x99, 0x68, 0x2b, 0xd3, 0x67, 0xbb, 0x77, 0xfb, 0x26, 0x6e, 0xf8, 0xb3, 0x6a, 0x1b, 0x71, 0x14,
	0x9d, 0x18, 0xc8, 0xa3, 0x50, 0xf9, 0xe9, 0xc0, 0x5a, 0xea, 0x6c, 0x6d, 0xc0, 0x2e, 0x78, 0xe0,
	0x6f, 0x12, 0xc0, 0x33, 0x11, 0xf7, 0x9f, 0x07, 0xfe, 0x36, 0x01, 0x32, 0x13, 0x75, 0x57, 0x1e,
	0xf8, 0xeb, 0x04, 0x16, 0x33, 0x75, 0xd7, 0x7f, 0x66, 0x3a, 0x13, 0x73, 0xff, 0x7b, 0xe0, 0xaf,
	0x12, 0x60, 0xbb, 0x37, 0xfb, 0x3a, 0x6e, 0x78, 0xac, 0x4d, 0x93, 0xf6, 0x87, 0x17, 0xd6, 0xf7,
	0x44, 0x0c, 0xb5, 0xfe, 0xe0, 0x17, 0xca, 0xef, 0x6d, 0x27, 0x6e, 0xf8, 0x53, 0x45, 0x4d, 0xad,
	0x5f, 0x55, 0x21, 0x65, 0x67, 0x84, 0xbc, 0x4c, 0xfb, 0xfe, 0xe1, 0x6b, 0x44, 0x70, 0x1e, 0x11,
	0xfc, 0x8c, 0x08, 0x3e, 0x27, 0x64, 0x9d, 0x27, 0x64, 0x7d, 0x4f, 0xc8, 0xda, 0x87, 0x3c, 0x37,
	0x59, 0x4b, 0x42, 0xaa, 0x64, 0xa4, 0x19, 0x67, 0x65, 0x50, 0x32, 0xd3, 0x29, 0x5d, 0x44, 0xcb,
	0xe7, 0x01, 0x2e, 0x87, 0x80, 0xb0, 0x92, 0x66, 0x12, 0xeb, 0x22, 0xc2, 0xe4, 0x6a, 0x99, 0xef,
	0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x85, 0x02, 0x04, 0x97, 0x01, 0x00, 0x00,
}

func (m *MsgPousljlwlybJloivfeuc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPousljlwlybJloivfeuc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPousljlwlybJloivfeuc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.E != 0 {
		i = encodeVarintA(dAtA, i, uint64(m.E))
		i--
		dAtA[i] = 0x28
	}
	if len(m.D) > 0 {
		i -= len(m.D)
		copy(dAtA[i:], m.D)
		i = encodeVarintA(dAtA, i, uint64(len(m.D)))
		i--
		dAtA[i] = 0x22
	}
	if m.C != 0 {
		i = encodeVarintA(dAtA, i, uint64(m.C))
		i--
		dAtA[i] = 0x18
	}
	if len(m.B) > 0 {
		i -= len(m.B)
		copy(dAtA[i:], m.B)
		i = encodeVarintA(dAtA, i, uint64(len(m.B)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.A) > 0 {
		i -= len(m.A)
		copy(dAtA[i:], m.A)
		i = encodeVarintA(dAtA, i, uint64(len(m.A)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMrtsdxfVexxblyqrzg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMrtsdxfVexxblyqrzg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMrtsdxfVexxblyqrzg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.E != 0 {
		i = encodeVarintA(dAtA, i, uint64(m.E))
		i--
		dAtA[i] = 0x28
	}
	if len(m.D) > 0 {
		i -= len(m.D)
		copy(dAtA[i:], m.D)
		i = encodeVarintA(dAtA, i, uint64(len(m.D)))
		i--
		dAtA[i] = 0x22
	}
	if m.C != 0 {
		i = encodeVarintA(dAtA, i, uint64(m.C))
		i--
		dAtA[i] = 0x18
	}
	if len(m.B) > 0 {
		i -= len(m.B)
		copy(dAtA[i:], m.B)
		i = encodeVarintA(dAtA, i, uint64(len(m.B)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.A) > 0 {
		i -= len(m.A)
		copy(dAtA[i:], m.A)
		i = encodeVarintA(dAtA, i, uint64(len(m.A)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLpctqrWokmmwtlm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLpctqrWokmmwtlm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLpctqrWokmmwtlm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.E != 0 {
		i = encodeVarintA(dAtA, i, uint64(m.E))
		i--
		dAtA[i] = 0x28
	}
	if len(m.D) > 0 {
		i -= len(m.D)
		copy(dAtA[i:], m.D)
		i = encodeVarintA(dAtA, i, uint64(len(m.D)))
		i--
		dAtA[i] = 0x22
	}
	if m.C != 0 {
		i = encodeVarintA(dAtA, i, uint64(m.C))
		i--
		dAtA[i] = 0x18
	}
	if len(m.B) > 0 {
		i -= len(m.B)
		copy(dAtA[i:], m.B)
		i = encodeVarintA(dAtA, i, uint64(len(m.B)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.A) > 0 {
		i -= len(m.A)
		copy(dAtA[i:], m.A)
		i = encodeVarintA(dAtA, i, uint64(len(m.A)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintA(dAtA []byte, offset int, v uint64) int {
	offset -= sovA(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPousljlwlybJloivfeuc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.A)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.B)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.C != 0 {
		n += 1 + sovA(uint64(m.C))
	}
	l = len(m.D)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.E != 0 {
		n += 1 + sovA(uint64(m.E))
	}
	return n
}

func (m *MsgMrtsdxfVexxblyqrzg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.A)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.B)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.C != 0 {
		n += 1 + sovA(uint64(m.C))
	}
	l = len(m.D)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.E != 0 {
		n += 1 + sovA(uint64(m.E))
	}
	return n
}

func (m *MsgLpctqrWokmmwtlm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.A)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	l = len(m.B)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.C != 0 {
		n += 1 + sovA(uint64(m.C))
	}
	l = len(m.D)
	if l > 0 {
		n += 1 + l + sovA(uint64(l))
	}
	if m.E != 0 {
		n += 1 + sovA(uint64(m.E))
	}
	return n
}

func sovA(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozA(x uint64) (n int) {
	return sovA(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPousljlwlybJloivfeuc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
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
			return fmt.Errorf("proto: MsgPousljlwlybJloivfeuc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPousljlwlybJloivfeuc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B = append(m.B[:0], dAtA[iNdEx:postIndex]...)
			if m.B == nil {
				m.B = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			m.C = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.C |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			m.E = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.E |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthA
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthA
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
func (m *MsgMrtsdxfVexxblyqrzg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
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
			return fmt.Errorf("proto: MsgMrtsdxfVexxblyqrzg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMrtsdxfVexxblyqrzg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B = append(m.B[:0], dAtA[iNdEx:postIndex]...)
			if m.B == nil {
				m.B = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			m.C = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.C |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			m.E = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.E |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthA
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthA
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
func (m *MsgLpctqrWokmmwtlm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowA
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
			return fmt.Errorf("proto: MsgLpctqrWokmmwtlm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLpctqrWokmmwtlm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B = append(m.B[:0], dAtA[iNdEx:postIndex]...)
			if m.B == nil {
				m.B = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			m.C = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.C |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
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
				return ErrInvalidLengthA
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthA
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.D = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			m.E = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowA
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.E |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipA(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthA
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthA
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
func skipA(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowA
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
					return 0, ErrIntOverflowA
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
					return 0, ErrIntOverflowA
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
				return 0, ErrInvalidLengthA
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupA
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthA
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthA        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowA          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupA = fmt.Errorf("proto: unexpected end of group")
)
