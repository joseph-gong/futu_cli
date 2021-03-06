// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Qot_UpdateOrderDetail.proto

package qotupdateorderdetail

import (
	fmt "fmt"
	_ "github.com/futuopen/ftapi4go/pb/common"
	qotcommon "github.com/futuopen/ftapi4go/pb/qotcommon"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
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

type S2C struct {
	Security       *qotcommon.Security    `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	OrderDetailAsk *qotcommon.OrderDetail `protobuf:"bytes,2,req,name=orderDetailAsk" json:"orderDetailAsk,omitempty"`
	OrderDetailBid *qotcommon.OrderDetail `protobuf:"bytes,3,req,name=orderDetailBid" json:"orderDetailBid,omitempty"`
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

func (m *S2C) GetSecurity() *qotcommon.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetOrderDetailAsk() *qotcommon.OrderDetail {
	if m != nil {
		return m.OrderDetailAsk
	}
	return nil
}

func (m *S2C) GetOrderDetailBid() *qotcommon.OrderDetail {
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
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe3, 0xa6, 0xbd, 0xb7, 0xb8, 0x08, 0xa1, 0x00, 0x22, 0x2a, 0xc8, 0x44, 0x9d, 0x32,
	0x80, 0x53, 0x45, 0x9d, 0x58, 0x10, 0x09, 0x2b, 0x42, 0xa4, 0xb0, 0xb0, 0xa0, 0x34, 0x71, 0x8b,
	0x05, 0xe9, 0x31, 0x8e, 0x33, 0xf4, 0x2d, 0x98, 0x78, 0x17, 0xde, 0xa0, 0x63, 0x47, 0x26, 0x84,
	0xda, 0x17, 0x41, 0x09, 0x29, 0x44, 0x55, 0x07, 0xd6, 0xff, 0xfb, 0xce, 0xf9, 0x8f, 0x6c, 0x7c,
	0x70, 0x0d, 0xea, 0xfe, 0x56, 0xc4, 0xa1, 0x62, 0x57, 0x32, 0x66, 0xf2, 0x82, 0xa9, 0x90, 0x3f,
	0x51, 0x21, 0x41, 0x81, 0xb1, 0xb7, 0x16, 0xb6, 0x37, 0x7d, 0x48, 0x12, 0x18, 0x7f, 0x4b, 0xed,
	0xed, 0x5c, 0xaa, 0x26, 0x9d, 0x37, 0x84, 0xf5, 0xbe, 0xeb, 0x1b, 0x5d, 0xdc, 0x4c, 0x59, 0x94,
	0x49, 0xae, 0x26, 0x26, 0xb2, 0x6a, 0x76, 0xcb, 0xdd, 0xa5, 0x15, 0xb9, 0x5f, 0xb2, 0xe0, 0xc7,
	0x32, 0xce, 0xf0, 0x16, 0xfc, 0x16, 0x9d, 0xa7, 0x8f, 0x66, 0xad, 0x98, 0xdb, 0xaf, 0xce, 0x55,
	0x4e, 0x09, 0x56, 0xf4, 0x95, 0x05, 0x1e, 0x8f, 0x4d, 0xfd, 0xef, 0x0b, 0x3c, 0x1e, 0x77, 0x5e,
	0x11, 0x6e, 0x06, 0x2c, 0x15, 0x30, 0x4e, 0x99, 0x41, 0xf0, 0x7f, 0xc9, 0xd4, 0xcd, 0x44, 0xb0,
	0xe2, 0xfe, 0xc6, 0x69, 0xfd, 0xa4, 0xd7, 0xed, 0x06, 0xcb, 0xd0, 0x38, 0xc4, 0xff, 0x24, 0x53,
	0x97, 0xe9, 0xc8, 0xac, 0x59, 0xc8, 0xde, 0xf0, 0xea, 0xd3, 0x8f, 0x23, 0x2d, 0x28, 0xb3, 0x7c,
	0x9a, 0x49, 0xe9, 0x43, 0xcc, 0x4c, 0xdd, 0x42, 0x76, 0xa3, 0xc4, 0xcb, 0xd0, 0x38, 0xc6, 0x7a,
	0xea, 0x46, 0x66, 0xdd, 0x42, 0x76, 0xcb, 0x6d, 0xd3, 0xf5, 0x1f, 0xd1, 0x77, 0xfd, 0x20, 0xd7,
	0xbc, 0x68, 0x3a, 0x27, 0x68, 0x36, 0x27, 0xe8, 0x73, 0x4e, 0xd0, 0xcb, 0x82, 0x68, 0xb3, 0x05,
	0xd1, 0xde, 0x17, 0x44, 0xc3, 0x3b, 0x11, 0x24, 0x74, 0x98, 0xa9, 0x8c, 0x82, 0x60, 0xe3, 0x50,
	0x70, 0x2a, 0x06, 0x77, 0xbd, 0x11, 0x57, 0x0f, 0xd9, 0x80, 0x46, 0x90, 0x38, 0x39, 0xcb, 0x91,
	0x33, 0x54, 0xa1, 0xe0, 0xbd, 0x11, 0x38, 0x62, 0xe0, 0x3c, 0x83, 0xca, 0x8a, 0xb6, 0xe2, 0x01,
	0xe2, 0xa2, 0xed, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x24, 0x82, 0x12, 0x84, 0x0e, 0x02, 0x00, 0x00,
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
				m.Security = &qotcommon.Security{}
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
				m.OrderDetailAsk = &qotcommon.OrderDetail{}
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
				m.OrderDetailBid = &qotcommon.OrderDetail{}
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
