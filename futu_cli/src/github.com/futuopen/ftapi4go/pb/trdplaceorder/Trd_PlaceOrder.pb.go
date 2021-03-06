// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Trd_PlaceOrder.proto

package trdplaceorder

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	common "github.com/futuopen/ftapi4go/pb/common"
	trdcommon "github.com/futuopen/ftapi4go/pb/trdcommon"
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

type C2S struct {
	PacketID  *common.PacketID     `protobuf:"bytes,1,req,name=packetID" json:"packetID,omitempty"`
	Header    *trdcommon.TrdHeader `protobuf:"bytes,2,req,name=header" json:"header,omitempty"`
	TrdSide   int32                `protobuf:"varint,3,req,name=trdSide" json:"trdSide"`
	OrderType int32                `protobuf:"varint,4,req,name=orderType" json:"orderType"`
	Code      string               `protobuf:"bytes,5,req,name=code" json:"code"`
	Qty       float64              `protobuf:"fixed64,6,req,name=qty" json:"qty"`
	Price     float64              `protobuf:"fixed64,7,opt,name=price" json:"price"`
	//以下2个为调整价格使用，都传才有效，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
	AdjustPrice        bool    `protobuf:"varint,8,opt,name=adjustPrice" json:"adjustPrice"`
	AdjustSideAndLimit float64 `protobuf:"fixed64,9,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit"`
	SecMarket          int32   `protobuf:"varint,10,opt,name=secMarket" json:"secMarket"`
	Remark             string  `protobuf:"bytes,11,opt,name=remark" json:"remark"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1f985437abff19, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(m, src)
}
func (m *C2S) XXX_Size() int {
	return m.Size()
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetPacketID() *common.PacketID {
	if m != nil {
		return m.PacketID
	}
	return nil
}

func (m *C2S) GetHeader() *trdcommon.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *C2S) GetTrdSide() int32 {
	if m != nil {
		return m.TrdSide
	}
	return 0
}

func (m *C2S) GetOrderType() int32 {
	if m != nil {
		return m.OrderType
	}
	return 0
}

func (m *C2S) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *C2S) GetQty() float64 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *C2S) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *C2S) GetAdjustPrice() bool {
	if m != nil {
		return m.AdjustPrice
	}
	return false
}

func (m *C2S) GetAdjustSideAndLimit() float64 {
	if m != nil {
		return m.AdjustSideAndLimit
	}
	return 0
}

func (m *C2S) GetSecMarket() int32 {
	if m != nil {
		return m.SecMarket
	}
	return 0
}

func (m *C2S) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type S2C struct {
	Header  *trdcommon.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	OrderID uint64               `protobuf:"varint,2,opt,name=orderID" json:"orderID"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1f985437abff19, []int{1}
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

func (m *S2C) GetHeader() *trdcommon.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetOrderID() uint64 {
	if m != nil {
		return m.OrderID
	}
	return 0
}

type Request struct {
	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1f985437abff19, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	//以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType *int32 `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg  string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg"`
	ErrCode int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode"`
	S2C     *S2C   `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1f985437abff19, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Trd_PlaceOrder.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_PlaceOrder.S2C")
	proto.RegisterType((*Request)(nil), "Trd_PlaceOrder.Request")
	proto.RegisterType((*Response)(nil), "Trd_PlaceOrder.Response")
}

func init() { proto.RegisterFile("Trd_PlaceOrder.proto", fileDescriptor_dc1f985437abff19) }

var fileDescriptor_dc1f985437abff19 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x24, 0xfd, 0xe5, 0x22, 0x34, 0x79, 0x80, 0xac, 0x0a, 0x05, 0xab, 0x12, 0x28,
	0x07, 0x9a, 0x56, 0x51, 0x4f, 0xdc, 0x58, 0x76, 0x60, 0x12, 0x13, 0x55, 0xda, 0x13, 0x42, 0x42,
	0x69, 0xfc, 0xd6, 0x85, 0x92, 0xda, 0x73, 0x9c, 0xc3, 0xfe, 0x0a, 0x38, 0xf0, 0x47, 0xed, 0xb8,
	0x23, 0x27, 0x84, 0xda, 0x7f, 0x04, 0x39, 0x4d, 0xdb, 0x80, 0x76, 0xe0, 0x98, 0xcf, 0xe7, 0xeb,
	0xe7, 0xf7, 0x5e, 0x8c, 0x9f, 0xcc, 0x15, 0xff, 0x3c, 0xfd, 0x1a, 0x27, 0xf0, 0x41, 0x71, 0x50,
	0xbe, 0x54, 0x42, 0x0b, 0xf2, 0xf8, 0x6f, 0xda, 0x7f, 0x14, 0x8a, 0x2c, 0x13, 0xeb, 0x9d, 0xed,
	0x9f, 0x18, 0x5b, 0x27, 0x83, 0x1f, 0x36, 0xb6, 0xc3, 0x60, 0x46, 0x5e, 0xe3, 0x8e, 0x8c, 0x93,
	0x15, 0xe8, 0x8b, 0x73, 0x8a, 0x98, 0xe5, 0xf5, 0x82, 0x13, 0xbf, 0x0a, 0x4e, 0x2b, 0x1e, 0x1d,
	0x12, 0x64, 0x88, 0x5b, 0xd7, 0x10, 0x73, 0x50, 0xd4, 0x2a, 0xb3, 0x4f, 0xfd, 0x5a, 0xe1, 0xb9,
	0xe2, 0xef, 0x4a, 0x19, 0x55, 0x21, 0xe2, 0xe2, 0xb6, 0x56, 0x7c, 0x96, 0x72, 0xa0, 0x36, 0xb3,
	0xbc, 0xe6, 0x99, 0x73, 0xf7, 0xeb, 0x45, 0x23, 0xda, 0x43, 0x32, 0xc0, 0x5d, 0x61, 0xba, 0x9d,
	0xdf, 0x4a, 0xa0, 0x4e, 0x2d, 0x71, 0xc4, 0x84, 0x62, 0x27, 0x11, 0x1c, 0x68, 0x93, 0x59, 0x5e,
	0xb7, 0xd2, 0x25, 0x21, 0xcf, 0xb0, 0x7d, 0xa3, 0x6f, 0x69, 0x8b, 0x59, 0x1e, 0xaa, 0x84, 0x01,
	0xa4, 0x8f, 0x9b, 0x52, 0xa5, 0x09, 0xd0, 0x36, 0x43, 0x07, 0xb3, 0x43, 0xe4, 0x15, 0xee, 0xc5,
	0xfc, 0x4b, 0x91, 0xeb, 0x69, 0x99, 0xe8, 0x30, 0xe4, 0x75, 0xaa, 0x44, 0x5d, 0x90, 0x09, 0x26,
	0xbb, 0x4f, 0xd3, 0xe7, 0xdb, 0x35, 0x7f, 0x9f, 0x66, 0xa9, 0xa6, 0xdd, 0x5a, 0xc1, 0x07, 0xbc,
	0x99, 0x27, 0x87, 0xe4, 0x32, 0x56, 0x2b, 0xd0, 0x14, 0x33, 0x74, 0x9c, 0xe7, 0x80, 0xc9, 0x73,
	0xdc, 0x52, 0x90, 0xc5, 0x6a, 0x45, 0x7b, 0x0c, 0x1d, 0x26, 0xaa, 0xd8, 0x60, 0x8e, 0xed, 0x59,
	0x10, 0xd6, 0xf6, 0x8c, 0xfe, 0x73, 0xcf, 0xe5, 0xc2, 0x2e, 0xce, 0xa9, 0xc5, 0x90, 0xe7, 0xec,
	0xf7, 0x5c, 0xc1, 0xc1, 0x18, 0xb7, 0x23, 0xb8, 0x29, 0x20, 0xd7, 0xe4, 0x25, 0xb6, 0x93, 0x20,
	0xaf, 0xca, 0x9e, 0xfa, 0xff, 0xbc, 0xa5, 0x30, 0x98, 0x45, 0xc6, 0x0f, 0xbe, 0x21, 0xdc, 0x89,
	0x20, 0x97, 0x62, 0x9d, 0x83, 0x29, 0xaf, 0x40, 0x97, 0x3f, 0xc9, 0x9c, 0x6b, 0xbe, 0x71, 0x86,
	0x93, 0xf1, 0x38, 0xda, 0xc3, 0xdd, 0x48, 0xfa, 0x32, 0x5f, 0x96, 0xb7, 0xd7, 0x46, 0x32, 0xcc,
	0x9c, 0x06, 0xa5, 0x42, 0x51, 0x3e, 0x82, 0xe3, 0x4a, 0xf6, 0xd0, 0x74, 0x94, 0x07, 0x09, 0x75,
	0x18, 0x7a, 0xa8, 0xa3, 0x59, 0x10, 0x46, 0xc6, 0x9f, 0x7d, 0xba, 0xdb, 0xb8, 0xe8, 0x7e, 0xe3,
	0xa2, 0xdf, 0x1b, 0x17, 0x7d, 0xdf, 0xba, 0x8d, 0xfb, 0xad, 0xdb, 0xf8, 0xb9, 0x75, 0x1b, 0xf8,
	0x34, 0x11, 0x99, 0x7f, 0x55, 0xe8, 0xc2, 0x17, 0x12, 0xd6, 0xb1, 0x4c, 0x7d, 0xb9, 0xf8, 0x38,
	0x5c, 0xa6, 0xfa, 0xba, 0x58, 0xf8, 0x89, 0xc8, 0x46, 0xc6, 0x19, 0x35, 0xba, 0xd2, 0xb1, 0x4c,
	0x27, 0x4b, 0x31, 0x92, 0x8b, 0x91, 0x56, 0x5c, 0x9a, 0x6b, 0xca, 0x1d, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x27, 0x99, 0xf6, 0xb1, 0x55, 0x03, 0x00, 0x00,
}

func (m *C2S) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C2S) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *C2S) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Remark)
	copy(dAtA[i:], m.Remark)
	i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(len(m.Remark)))
	i--
	dAtA[i] = 0x5a
	i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(m.SecMarket))
	i--
	dAtA[i] = 0x50
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.AdjustSideAndLimit))))
	i--
	dAtA[i] = 0x49
	i--
	if m.AdjustPrice {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x40
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Price))))
	i--
	dAtA[i] = 0x39
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Qty))))
	i--
	dAtA[i] = 0x31
	i -= len(m.Code)
	copy(dAtA[i:], m.Code)
	i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(len(m.Code)))
	i--
	dAtA[i] = 0x2a
	i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(m.OrderType))
	i--
	dAtA[i] = 0x20
	i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(m.TrdSide))
	i--
	dAtA[i] = 0x18
	if m.Header == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("header")
	} else {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.PacketID == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("packetID")
	} else {
		{
			size, err := m.PacketID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(m.OrderID))
	i--
	dAtA[i] = 0x10
	if m.Header == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("header")
	} else {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.C2S == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("c2s")
	} else {
		{
			size, err := m.C2S.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(size))
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
			i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(m.ErrCode))
	i--
	dAtA[i] = 0x18
	i -= len(m.RetMsg)
	copy(dAtA[i:], m.RetMsg)
	i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(len(m.RetMsg)))
	i--
	dAtA[i] = 0x12
	if m.RetType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	} else {
		i = encodeVarintTrd_PlaceOrder(dAtA, i, uint64(*m.RetType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTrd_PlaceOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovTrd_PlaceOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *C2S) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PacketID != nil {
		l = m.PacketID.Size()
		n += 1 + l + sovTrd_PlaceOrder(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovTrd_PlaceOrder(uint64(l))
	}
	n += 1 + sovTrd_PlaceOrder(uint64(m.TrdSide))
	n += 1 + sovTrd_PlaceOrder(uint64(m.OrderType))
	l = len(m.Code)
	n += 1 + l + sovTrd_PlaceOrder(uint64(l))
	n += 9
	n += 9
	n += 2
	n += 9
	n += 1 + sovTrd_PlaceOrder(uint64(m.SecMarket))
	l = len(m.Remark)
	n += 1 + l + sovTrd_PlaceOrder(uint64(l))
	return n
}

func (m *S2C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovTrd_PlaceOrder(uint64(l))
	}
	n += 1 + sovTrd_PlaceOrder(uint64(m.OrderID))
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.C2S != nil {
		l = m.C2S.Size()
		n += 1 + l + sovTrd_PlaceOrder(uint64(l))
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
		n += 1 + sovTrd_PlaceOrder(uint64(*m.RetType))
	}
	l = len(m.RetMsg)
	n += 1 + l + sovTrd_PlaceOrder(uint64(l))
	n += 1 + sovTrd_PlaceOrder(uint64(m.ErrCode))
	if m.S2C != nil {
		l = m.S2C.Size()
		n += 1 + l + sovTrd_PlaceOrder(uint64(l))
	}
	return n
}

func sovTrd_PlaceOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTrd_PlaceOrder(x uint64) (n int) {
	return sovTrd_PlaceOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *C2S) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrd_PlaceOrder
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
			return fmt.Errorf("proto: C2S: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C2S: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
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
				return ErrInvalidLengthTrd_PlaceOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PacketID == nil {
				m.PacketID = &common.PacketID{}
			}
			if err := m.PacketID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
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
				return ErrInvalidLengthTrd_PlaceOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &trdcommon.TrdHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrdSide", wireType)
			}
			m.TrdSide = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrdSide |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			m.OrderType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
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
				return ErrInvalidLengthTrd_PlaceOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qty", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Qty = float64(math.Float64frombits(v))
			hasFields[0] |= uint64(0x00000020)
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Price = float64(math.Float64frombits(v))
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjustPrice", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AdjustPrice = bool(v != 0)
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjustSideAndLimit", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.AdjustSideAndLimit = float64(math.Float64frombits(v))
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecMarket", wireType)
			}
			m.SecMarket = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SecMarket |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remark", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
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
				return ErrInvalidLengthTrd_PlaceOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Remark = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrd_PlaceOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("packetID")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("header")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("trdSide")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("orderType")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("code")
	}
	if hasFields[0]&uint64(0x00000020) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("qty")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
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
				return ErrIntOverflowTrd_PlaceOrder
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
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
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
				return ErrInvalidLengthTrd_PlaceOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &trdcommon.TrdHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			m.OrderID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTrd_PlaceOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("header")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrd_PlaceOrder
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C2S", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrd_PlaceOrder
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
				return ErrInvalidLengthTrd_PlaceOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.C2S == nil {
				m.C2S = &C2S{}
			}
			if err := m.C2S.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipTrd_PlaceOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("c2s")
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
				return ErrIntOverflowTrd_PlaceOrder
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
					return ErrIntOverflowTrd_PlaceOrder
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
					return ErrIntOverflowTrd_PlaceOrder
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
				return ErrInvalidLengthTrd_PlaceOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
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
					return ErrIntOverflowTrd_PlaceOrder
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
					return ErrIntOverflowTrd_PlaceOrder
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
				return ErrInvalidLengthTrd_PlaceOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
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
			skippy, err := skipTrd_PlaceOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTrd_PlaceOrder
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
func skipTrd_PlaceOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrd_PlaceOrder
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
					return 0, ErrIntOverflowTrd_PlaceOrder
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
					return 0, ErrIntOverflowTrd_PlaceOrder
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
				return 0, ErrInvalidLengthTrd_PlaceOrder
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTrd_PlaceOrder
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrd_PlaceOrder
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
				next, err := skipTrd_PlaceOrder(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTrd_PlaceOrder
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
	ErrInvalidLengthTrd_PlaceOrder = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrd_PlaceOrder   = fmt.Errorf("proto: integer overflow")
)
