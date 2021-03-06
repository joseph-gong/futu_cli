// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Qot_UpdateOrderBook.proto

package qotupdateorderbook

import (
	encoding_binary "encoding/binary"
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
	Security                *qotcommon.Security    `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	OrderBookAskList        []*qotcommon.OrderBook `protobuf:"bytes,2,rep,name=orderBookAskList" json:"orderBookAskList,omitempty"`
	OrderBookBidList        []*qotcommon.OrderBook `protobuf:"bytes,3,rep,name=orderBookBidList" json:"orderBookBidList,omitempty"`
	SvrRecvTimeBid          string                 `protobuf:"bytes,4,opt,name=svrRecvTimeBid" json:"svrRecvTimeBid"`
	SvrRecvTimeBidTimestamp float64                `protobuf:"fixed64,5,opt,name=svrRecvTimeBidTimestamp" json:"svrRecvTimeBidTimestamp"`
	SvrRecvTimeAsk          string                 `protobuf:"bytes,6,opt,name=svrRecvTimeAsk" json:"svrRecvTimeAsk"`
	SvrRecvTimeAskTimestamp float64                `protobuf:"fixed64,7,opt,name=svrRecvTimeAskTimestamp" json:"svrRecvTimeAskTimestamp"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb6eb95eb88e215, []int{0}
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

func (m *S2C) GetOrderBookAskList() []*qotcommon.OrderBook {
	if m != nil {
		return m.OrderBookAskList
	}
	return nil
}

func (m *S2C) GetOrderBookBidList() []*qotcommon.OrderBook {
	if m != nil {
		return m.OrderBookBidList
	}
	return nil
}

func (m *S2C) GetSvrRecvTimeBid() string {
	if m != nil {
		return m.SvrRecvTimeBid
	}
	return ""
}

func (m *S2C) GetSvrRecvTimeBidTimestamp() float64 {
	if m != nil {
		return m.SvrRecvTimeBidTimestamp
	}
	return 0
}

func (m *S2C) GetSvrRecvTimeAsk() string {
	if m != nil {
		return m.SvrRecvTimeAsk
	}
	return ""
}

func (m *S2C) GetSvrRecvTimeAskTimestamp() float64 {
	if m != nil {
		return m.SvrRecvTimeAskTimestamp
	}
	return 0
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
	return fileDescriptor_2eb6eb95eb88e215, []int{1}
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
	proto.RegisterType((*S2C)(nil), "Qot_UpdateOrderBook.S2C")
	proto.RegisterType((*Response)(nil), "Qot_UpdateOrderBook.Response")
}

func init() { proto.RegisterFile("Qot_UpdateOrderBook.proto", fileDescriptor_2eb6eb95eb88e215) }

var fileDescriptor_2eb6eb95eb88e215 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0x4d, 0xbb, 0xbb, 0x4e, 0x45, 0x96, 0x59, 0xc5, 0x71, 0x91, 0x18, 0xf6, 0x14,
	0x44, 0x93, 0x12, 0xf6, 0xe4, 0x41, 0x48, 0x72, 0x55, 0xc4, 0x74, 0xbd, 0x78, 0x91, 0xfc, 0x99,
	0xd6, 0x21, 0xa4, 0xef, 0x38, 0x33, 0x29, 0xf4, 0x5b, 0x78, 0xf0, 0xe2, 0x37, 0xea, 0xb1, 0x47,
	0x4f, 0x22, 0xed, 0x17, 0x91, 0xc4, 0xb4, 0xa6, 0xb5, 0x05, 0x4f, 0x81, 0xe7, 0xf9, 0xbd, 0xef,
	0x6f, 0xc8, 0x0c, 0x7e, 0xf2, 0x1e, 0xf4, 0xa7, 0x0f, 0x22, 0x4f, 0x34, 0x7b, 0x27, 0x73, 0x26,
	0x43, 0x80, 0xc2, 0x15, 0x12, 0x34, 0x90, 0xab, 0x23, 0xd5, 0xf5, 0xfd, 0x08, 0xca, 0x12, 0x66,
	0x7f, 0x90, 0xeb, 0xcb, 0x1a, 0xe9, 0x26, 0x37, 0xdf, 0x4d, 0x6c, 0x8e, 0xfd, 0x88, 0x8c, 0xf0,
	0x85, 0x62, 0x59, 0x25, 0xb9, 0x5e, 0x50, 0x64, 0xf7, 0x9c, 0xa1, 0xff, 0xd0, 0xed, 0xc0, 0xe3,
	0xb6, 0x8b, 0x77, 0x14, 0x09, 0xf0, 0x25, 0x6c, 0x35, 0x81, 0x2a, 0xde, 0x70, 0xa5, 0x69, 0xcf,
	0x36, 0x9d, 0xa1, 0xff, 0xa8, 0x3b, 0xb9, 0x3b, 0x4a, 0xfc, 0x0f, 0xbe, 0xb7, 0x22, 0xe4, 0x79,
	0xb3, 0xc2, 0xfc, 0xbf, 0x15, 0x2d, 0x4e, 0x5e, 0xe0, 0x07, 0x6a, 0x2e, 0x63, 0x96, 0xcd, 0xef,
	0x78, 0xc9, 0x42, 0x9e, 0xd3, 0xbe, 0x8d, 0x9c, 0x7b, 0x61, 0x7f, 0xf9, 0xf3, 0x99, 0x11, 0x1f,
	0x74, 0xe4, 0x35, 0x7e, 0xbc, 0x9f, 0xd4, 0x1f, 0xa5, 0x93, 0x52, 0xd0, 0x81, 0x8d, 0x1c, 0xd4,
	0x8e, 0x9d, 0x82, 0x0e, 0x6c, 0x81, 0x2a, 0xe8, 0xd9, 0x09, 0x5b, 0xa0, 0x8a, 0x03, 0x5b, 0xa0,
	0x8a, 0xbf, 0xb6, 0xf3, 0x13, 0xb6, 0x2e, 0x74, 0xf3, 0x0d, 0xe1, 0x8b, 0x98, 0x29, 0x01, 0x33,
	0xc5, 0x88, 0x85, 0xcf, 0x25, 0xd3, 0x77, 0x0b, 0xc1, 0x9a, 0xfb, 0x19, 0xbc, 0xea, 0xbf, 0xbc,
	0x1d, 0x8d, 0xe2, 0x6d, 0x48, 0x9e, 0xe2, 0x33, 0xc9, 0xf4, 0x5b, 0x35, 0xa5, 0xbd, 0xce, 0x91,
	0xda, 0xac, 0x9e, 0x66, 0x52, 0x46, 0x90, 0x33, 0x6a, 0xda, 0xc8, 0x19, 0xb4, 0xf5, 0x36, 0x24,
	0xcf, 0xb1, 0xa9, 0xfc, 0xac, 0xf9, 0x77, 0x43, 0x9f, 0xba, 0xc7, 0x1e, 0xd9, 0xd8, 0x8f, 0xe2,
	0x1a, 0x0a, 0x93, 0xe5, 0xda, 0x42, 0xab, 0xb5, 0x85, 0x7e, 0xad, 0x2d, 0xf4, 0x75, 0x63, 0x19,
	0xab, 0x8d, 0x65, 0xfc, 0xd8, 0x58, 0x06, 0xbe, 0xca, 0xa0, 0x74, 0x27, 0x95, 0xae, 0x5c, 0x10,
	0x6c, 0x96, 0x08, 0xee, 0x8a, 0xf4, 0xa3, 0x3f, 0xe5, 0xfa, 0x73, 0x95, 0xba, 0x19, 0x94, 0x5e,
	0xdd, 0xd5, 0x95, 0x37, 0xd1, 0x89, 0xe0, 0xb7, 0x53, 0xf0, 0x44, 0xea, 0x7d, 0x01, 0x5d, 0x35,
	0xae, 0xe6, 0x72, 0x53, 0x80, 0xe2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0xf1, 0x47, 0x66,
	0xe6, 0x02, 0x00, 0x00,
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
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SvrRecvTimeAskTimestamp))))
	i--
	dAtA[i] = 0x39
	i -= len(m.SvrRecvTimeAsk)
	copy(dAtA[i:], m.SvrRecvTimeAsk)
	i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(len(m.SvrRecvTimeAsk)))
	i--
	dAtA[i] = 0x32
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SvrRecvTimeBidTimestamp))))
	i--
	dAtA[i] = 0x29
	i -= len(m.SvrRecvTimeBid)
	copy(dAtA[i:], m.SvrRecvTimeBid)
	i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(len(m.SvrRecvTimeBid)))
	i--
	dAtA[i] = 0x22
	if len(m.OrderBookBidList) > 0 {
		for iNdEx := len(m.OrderBookBidList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderBookBidList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OrderBookAskList) > 0 {
		for iNdEx := len(m.OrderBookAskList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderBookAskList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
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
			i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(size))
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
			i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(m.ErrCode))
	i--
	dAtA[i] = 0x18
	i -= len(m.RetMsg)
	copy(dAtA[i:], m.RetMsg)
	i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(len(m.RetMsg)))
	i--
	dAtA[i] = 0x12
	if m.RetType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	} else {
		i = encodeVarintQot_UpdateOrderBook(dAtA, i, uint64(*m.RetType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQot_UpdateOrderBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovQot_UpdateOrderBook(v)
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
		n += 1 + l + sovQot_UpdateOrderBook(uint64(l))
	}
	if len(m.OrderBookAskList) > 0 {
		for _, e := range m.OrderBookAskList {
			l = e.Size()
			n += 1 + l + sovQot_UpdateOrderBook(uint64(l))
		}
	}
	if len(m.OrderBookBidList) > 0 {
		for _, e := range m.OrderBookBidList {
			l = e.Size()
			n += 1 + l + sovQot_UpdateOrderBook(uint64(l))
		}
	}
	l = len(m.SvrRecvTimeBid)
	n += 1 + l + sovQot_UpdateOrderBook(uint64(l))
	n += 9
	l = len(m.SvrRecvTimeAsk)
	n += 1 + l + sovQot_UpdateOrderBook(uint64(l))
	n += 9
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RetType != nil {
		n += 1 + sovQot_UpdateOrderBook(uint64(*m.RetType))
	}
	l = len(m.RetMsg)
	n += 1 + l + sovQot_UpdateOrderBook(uint64(l))
	n += 1 + sovQot_UpdateOrderBook(uint64(m.ErrCode))
	if m.S2C != nil {
		l = m.S2C.Size()
		n += 1 + l + sovQot_UpdateOrderBook(uint64(l))
	}
	return n
}

func sovQot_UpdateOrderBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQot_UpdateOrderBook(x uint64) (n int) {
	return sovQot_UpdateOrderBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return ErrIntOverflowQot_UpdateOrderBook
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
					return ErrIntOverflowQot_UpdateOrderBook
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
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
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
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookAskList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderBook
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
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookAskList = append(m.OrderBookAskList, &qotcommon.OrderBook{})
			if err := m.OrderBookAskList[len(m.OrderBookAskList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookBidList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderBook
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
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookBidList = append(m.OrderBookBidList, &qotcommon.OrderBook{})
			if err := m.OrderBookBidList[len(m.OrderBookBidList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SvrRecvTimeBid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderBook
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
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SvrRecvTimeBid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SvrRecvTimeBidTimestamp", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SvrRecvTimeBidTimestamp = float64(math.Float64frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SvrRecvTimeAsk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateOrderBook
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
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SvrRecvTimeAsk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SvrRecvTimeAskTimestamp", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SvrRecvTimeAskTimestamp = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipQot_UpdateOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
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
				return ErrIntOverflowQot_UpdateOrderBook
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
					return ErrIntOverflowQot_UpdateOrderBook
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
					return ErrIntOverflowQot_UpdateOrderBook
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
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
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
					return ErrIntOverflowQot_UpdateOrderBook
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
					return ErrIntOverflowQot_UpdateOrderBook
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
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
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
			skippy, err := skipQot_UpdateOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_UpdateOrderBook
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
func skipQot_UpdateOrderBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQot_UpdateOrderBook
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
					return 0, ErrIntOverflowQot_UpdateOrderBook
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
					return 0, ErrIntOverflowQot_UpdateOrderBook
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
				return 0, ErrInvalidLengthQot_UpdateOrderBook
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQot_UpdateOrderBook
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQot_UpdateOrderBook
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
				next, err := skipQot_UpdateOrderBook(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQot_UpdateOrderBook
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
	ErrInvalidLengthQot_UpdateOrderBook = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQot_UpdateOrderBook   = fmt.Errorf("proto: integer overflow")
)
