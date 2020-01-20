// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Qot_UpdateOrderDetail.proto

package Qot_UpdateOrderDetail

import (
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	_ "pbgo/Common"
	Qot_Common "pbgo/Qot_Common"
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

type S2C struct {
	Security       *Qot_Common.Security    `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	OrderDetailAsk *Qot_Common.OrderDetail `protobuf:"bytes,2,req,name=orderDetailAsk" json:"orderDetailAsk,omitempty"`
	OrderDetailBid *Qot_Common.OrderDetail `protobuf:"bytes,3,req,name=orderDetailBid" json:"orderDetailBid,omitempty"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_754cbeee7853dcd5, []int{0}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(m, src)
}
func (m *S2C) XXX_Size() int {
	return m.Size()
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetOrderDetailAsk() *Qot_Common.OrderDetail {
	if m != nil {
		return m.OrderDetailAsk
	}
	return nil
}

func (m *S2C) GetOrderDetailBid() *Qot_Common.OrderDetail {
	if m != nil {
		return m.OrderDetailBid
	}
	return nil
}

type Response struct {
	RetType *int32 `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg  string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg"`
	ErrCode int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode"`
	S2C     *S2C   `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_754cbeee7853dcd5, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*S2C)(nil), "Qot_UpdateOrderDetail.S2C")
	proto.RegisterType((*Response)(nil), "Qot_UpdateOrderDetail.Response")
}

func init() { proto.RegisterFile("Qot_UpdateOrderDetail.proto", fileDescriptor_754cbeee7853dcd5) }

var fileDescriptor_754cbeee7853dcd5 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x0e, 0xcc, 0x2f, 0x89,
	0x0f, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0xf5, 0x2f, 0x4a, 0x49, 0x2d, 0x72, 0x49, 0x2d, 0x49, 0xcc,
	0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc5, 0x2a, 0x29, 0xc5, 0xe3, 0x9c, 0x9f,
	0x9b, 0x9b, 0x9f, 0x07, 0x51, 0x24, 0x25, 0x00, 0x52, 0x84, 0x2c, 0xa2, 0xb4, 0x93, 0x91, 0x8b,
	0x39, 0xd8, 0xc8, 0x59, 0xc8, 0x80, 0x8b, 0xa3, 0x38, 0x35, 0xb9, 0xb4, 0x28, 0xb3, 0xa4, 0x52,
	0x82, 0x51, 0x81, 0x49, 0x83, 0xdb, 0x48, 0x44, 0x0f, 0x49, 0x71, 0x30, 0x54, 0x2e, 0x08, 0xae,
	0x4a, 0xc8, 0x9e, 0x8b, 0x2f, 0x1f, 0x61, 0x91, 0x63, 0x71, 0xb6, 0x04, 0x13, 0x58, 0x9f, 0x38,
	0xb2, 0x3e, 0x24, 0xa7, 0x04, 0xa1, 0x29, 0x47, 0x33, 0xc0, 0x29, 0x33, 0x45, 0x82, 0x99, 0x78,
	0x03, 0x9c, 0x32, 0x53, 0x94, 0xa6, 0x31, 0x72, 0x71, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15,
	0xa7, 0x0a, 0xc9, 0x71, 0xb1, 0x17, 0xa5, 0x96, 0x84, 0x54, 0x16, 0xa4, 0x82, 0xdd, 0xcf, 0x6a,
	0xc5, 0xa2, 0x6b, 0x62, 0x60, 0x10, 0x04, 0x13, 0x14, 0x92, 0xe1, 0x62, 0x2b, 0x4a, 0x2d, 0xf1,
	0x2d, 0x4e, 0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x74, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08,
	0x2a, 0x06, 0xd2, 0x9d, 0x5a, 0x54, 0xe4, 0x9c, 0x9f, 0x92, 0x2a, 0xc1, 0xac, 0xc0, 0xa8, 0xc1,
	0x0a, 0x95, 0x86, 0x09, 0x0a, 0xe9, 0x70, 0x31, 0x17, 0x1b, 0x25, 0x4b, 0xb0, 0x28, 0x30, 0x6a,
	0x70, 0x1b, 0x49, 0xe9, 0x61, 0x8f, 0x88, 0x60, 0x23, 0xe7, 0x20, 0x90, 0x32, 0x27, 0xcf, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0x12, 0x4e, 0xce, 0xcf, 0xd5, 0x4b, 0x2b,
	0x2d, 0x29, 0xd5, 0xcb, 0x2f, 0x48, 0xcd, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x48, 0x8a, 0x92, 0x2a,
	0x48, 0x4a, 0xcf, 0xd7, 0xc7, 0x6a, 0x26, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x43, 0x02, 0xef,
	0xf4, 0x01, 0x00, 0x00,
}

func (m *S2C) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *S2C) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OrderDetailBid == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("orderDetailBid")
	} else {
		{
			size, err := m.OrderDetailBid.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQot_UpdateOrderDetail(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.OrderDetailAsk == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("orderDetailAsk")
	} else {
		{
			size, err := m.OrderDetailAsk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQot_UpdateOrderDetail(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Security == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("security")
	} else {
		{
			size, err := m.Security.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQot_UpdateOrderDetail(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.S2C != nil {
		{
			size, err := m.S2C.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQot_UpdateOrderDetail(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintQot_UpdateOrderDetail(dAtA, i, uint64(m.ErrCode))
	i--
	dAtA[i] = 0x18
	i -= len(m.RetMsg)
	copy(dAtA[i:], m.RetMsg)
	i = encodeVarintQot_UpdateOrderDetail(dAtA, i, uint64(len(m.RetMsg)))
	i--
	dAtA[i] = 0x12
	if m.RetType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	} else {
		i = encodeVarintQot_UpdateOrderDetail(dAtA, i, uint64(*m.RetType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQot_UpdateOrderDetail(dAtA []byte, offset int, v uint64) int {
	offset -= sovQot_UpdateOrderDetail(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *S2C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Security != nil {
		l = m.Security.Size()
		n += 1 + l + sovQot_UpdateOrderDetail(uint64(l))
	}
	if m.OrderDetailAsk != nil {
		l = m.OrderDetailAsk.Size()
		n += 1 + l + sovQot_UpdateOrderDetail(uint64(l))
	}
	if m.OrderDetailBid != nil {
		l = m.OrderDetailBid.Size()
		n += 1 + l + sovQot_UpdateOrderDetail(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RetType != nil {
		n += 1 + sovQot_UpdateOrderDetail(uint64(*m.RetType))
	}
	l = len(m.RetMsg)
	n += 1 + l + sovQot_UpdateOrderDetail(uint64(l))
	n += 1 + sovQot_UpdateOrderDetail(uint64(m.ErrCode))
	if m.S2C != nil {
		l = m.S2C.Size()
		n += 1 + l + sovQot_UpdateOrderDetail(uint64(l))
	}
	return n
}

func sovQot_UpdateOrderDetail(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQot_UpdateOrderDetail(x uint64) (n int) {
	return sovQot_UpdateOrderDetail(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *S2C) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQot_UpdateOrderDetail
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
			return fmt.Errorf("proto: S2C: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Security", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderDetail
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
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Security == nil {
				m.Security = &Qot_Common.Security{}
			}
			if err := m.Security.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderDetailAsk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderDetail
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
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OrderDetailAsk == nil {
				m.OrderDetailAsk = &Qot_Common.OrderDetail{}
			}
			if err := m.OrderDetailAsk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderDetailBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderDetail
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
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OrderDetailBid == nil {
				m.OrderDetailBid = &Qot_Common.OrderDetail{}
			}
			if err := m.OrderDetailBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipQot_UpdateOrderDetail(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("security")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("orderDetailAsk")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("orderDetailBid")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQot_UpdateOrderDetail
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetType", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderDetail
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RetType = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderDetail
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
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RetMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderDetail
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S2C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderDetail
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
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.S2C == nil {
				m.S2C = &S2C{}
			}
			if err := m.S2C.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQot_UpdateOrderDetail(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_UpdateOrderDetail
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQot_UpdateOrderDetail(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQot_UpdateOrderDetail
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
					return 0, ErrIntOverflowQot_UpdateOrderDetail
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQot_UpdateOrderDetail
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
				return 0, ErrInvalidLengthQot_UpdateOrderDetail
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQot_UpdateOrderDetail
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQot_UpdateOrderDetail
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipQot_UpdateOrderDetail(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQot_UpdateOrderDetail
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthQot_UpdateOrderDetail = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQot_UpdateOrderDetail   = fmt.Errorf("proto: integer overflow")
)