// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Qot_GetStaticInfo.proto

package qotgetstaticinfo

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

type C2S struct {
	Market       int32                 `protobuf:"varint,1,opt,name=market" json:"market"`
	SecType      int32                 `protobuf:"varint,2,opt,name=secType" json:"secType"`
	SecurityList []*qotcommon.Security `protobuf:"bytes,3,rep,name=securityList" json:"securityList,omitempty"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d1a334bdcd68a6, []int{0}
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

func (m *C2S) GetMarket() int32 {
	if m != nil {
		return m.Market
	}
	return 0
}

func (m *C2S) GetSecType() int32 {
	if m != nil {
		return m.SecType
	}
	return 0
}

func (m *C2S) GetSecurityList() []*qotcommon.Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

type S2C struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo `protobuf:"bytes,1,rep,name=staticInfoList" json:"staticInfoList,omitempty"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d1a334bdcd68a6, []int{1}
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

func (m *S2C) GetStaticInfoList() []*qotcommon.SecurityStaticInfo {
	if m != nil {
		return m.StaticInfoList
	}
	return nil
}

type Request struct {
	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d1a334bdcd68a6, []int{2}
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
	RetType *int32 `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg  string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg"`
	ErrCode int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode"`
	S2C     *S2C   `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_52d1a334bdcd68a6, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetStaticInfo.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetStaticInfo.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetStaticInfo.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetStaticInfo.Response")
}

func init() { proto.RegisterFile("Qot_GetStaticInfo.proto", fileDescriptor_52d1a334bdcd68a6) }

var fileDescriptor_52d1a334bdcd68a6 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x71, 0x42, 0xe0, 0x12, 0x21, 0x30, 0x08, 0xac, 0x08, 0x1d, 0x96, 0x27, 0x2f,
	0xd8, 0xd1, 0x91, 0x01, 0x31, 0xc6, 0x12, 0x08, 0x89, 0x0c, 0x9c, 0x99, 0x58, 0x22, 0xc7, 0x9c,
	0x8d, 0x85, 0xec, 0xbb, 0xf8, 0x5e, 0x0f, 0x19, 0xfa, 0x1d, 0xaa, 0x7e, 0xaa, 0x8c, 0x19, 0x3b,
	0x55, 0x55, 0xf2, 0x45, 0xaa, 0x73, 0x92, 0xd6, 0x6d, 0xda, 0xf5, 0x79, 0x9e, 0xf7, 0xcf, 0xf3,
	0xc3, 0xef, 0x7f, 0x09, 0x98, 0x7f, 0xe7, 0x10, 0x41, 0x0c, 0x79, 0xf2, 0xa3, 0x4c, 0x85, 0x2f,
	0x2b, 0x01, 0xc2, 0x7a, 0x7d, 0x62, 0x8c, 0x86, 0xa1, 0x28, 0x0a, 0x51, 0xee, 0x03, 0xa3, 0x57,
	0x3a, 0xd0, 0x56, 0xdc, 0x33, 0x6c, 0x86, 0x34, 0xb2, 0x3e, 0xe0, 0x67, 0x45, 0x5c, 0xfd, 0xe7,
	0x60, 0x23, 0x07, 0x79, 0xbd, 0x69, 0x77, 0x7d, 0xf5, 0xd1, 0x60, 0x07, 0xcd, 0x22, 0xb8, 0xaf,
	0x78, 0xf2, 0x7b, 0x25, 0xb9, 0xdd, 0x69, 0xd9, 0x47, 0xd1, 0xfa, 0x82, 0x87, 0x8a, 0x27, 0x75,
	0x95, 0xc3, 0xea, 0x67, 0xae, 0xc0, 0x36, 0x1d, 0xd3, 0x1b, 0xd0, 0xb7, 0x7e, 0xeb, 0x5a, 0x74,
	0xf0, 0xd9, 0xbd, 0xa4, 0x3b, 0xc3, 0x66, 0x44, 0x43, 0xeb, 0x1b, 0x7e, 0xa9, 0x6e, 0x7f, 0x6e,
	0x56, 0xa0, 0x66, 0x05, 0x79, 0x6c, 0xc5, 0x5d, 0x3b, 0xf6, 0x60, 0xca, 0xfd, 0x8c, 0xfb, 0x8c,
	0x2f, 0x6b, 0xae, 0xc0, 0xf2, 0xb0, 0x99, 0x50, 0x65, 0x23, 0xa7, 0xe3, 0x0d, 0xe8, 0x3b, 0xff,
	0x14, 0x59, 0x48, 0x23, 0xa6, 0x23, 0xee, 0x05, 0xc2, 0xcf, 0x19, 0x57, 0x52, 0x94, 0x8a, 0xeb,
	0xaa, 0x15, 0x87, 0xa6, 0xaa, 0x1e, 0xed, 0x7d, 0xed, 0x7e, 0x9a, 0x8c, 0xc7, 0xec, 0x28, 0x6a,
	0x50, 0x15, 0x87, 0x99, 0xca, 0x1a, 0x12, 0x2f, 0x8e, 0xa0, 0xf6, 0x9a, 0x9e, 0xe6, 0x55, 0x15,
	0x8a, 0xbf, 0xdc, 0x36, 0xdb, 0xa0, 0x0e, 0xa2, 0x7e, 0x4a, 0xd1, 0xc4, 0xee, 0x3a, 0xe8, 0x89,
	0xa7, 0x22, 0x1a, 0x32, 0x1d, 0x99, 0xce, 0xd7, 0x5b, 0x82, 0x36, 0x5b, 0x82, 0xae, 0xb7, 0x04,
	0x9d, 0xef, 0x88, 0xb1, 0xd9, 0x11, 0xe3, 0x72, 0x47, 0x0c, 0xfc, 0x26, 0x11, 0x85, 0x9f, 0xd6,
	0x50, 0xfb, 0x42, 0xf2, 0x32, 0x96, 0xb9, 0x2f, 0x17, 0x7f, 0xc6, 0x59, 0x0e, 0xff, 0xea, 0x85,
	0x9f, 0x88, 0x22, 0xd0, 0x9e, 0xb6, 0x82, 0x14, 0x62, 0x99, 0x4f, 0x32, 0x11, 0xc8, 0x45, 0xb0,
	0x14, 0x90, 0x71, 0xd8, 0xe3, 0xca, 0xcb, 0x54, 0xdc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x09,
	0x10, 0x52, 0x45, 0x02, 0x00, 0x00,
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
	if len(m.SecurityList) > 0 {
		for iNdEx := len(m.SecurityList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SecurityList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(m.SecType))
	i--
	dAtA[i] = 0x10
	i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(m.Market))
	i--
	dAtA[i] = 0x8
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
	if len(m.StaticInfoList) > 0 {
		for iNdEx := len(m.StaticInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StaticInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
			i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(size))
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
			i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(m.ErrCode))
	i--
	dAtA[i] = 0x18
	i -= len(m.RetMsg)
	copy(dAtA[i:], m.RetMsg)
	i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(len(m.RetMsg)))
	i--
	dAtA[i] = 0x12
	if m.RetType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	} else {
		i = encodeVarintQot_GetStaticInfo(dAtA, i, uint64(*m.RetType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQot_GetStaticInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovQot_GetStaticInfo(v)
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
	n += 1 + sovQot_GetStaticInfo(uint64(m.Market))
	n += 1 + sovQot_GetStaticInfo(uint64(m.SecType))
	if len(m.SecurityList) > 0 {
		for _, e := range m.SecurityList {
			l = e.Size()
			n += 1 + l + sovQot_GetStaticInfo(uint64(l))
		}
	}
	return n
}

func (m *S2C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.StaticInfoList) > 0 {
		for _, e := range m.StaticInfoList {
			l = e.Size()
			n += 1 + l + sovQot_GetStaticInfo(uint64(l))
		}
	}
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
		n += 1 + l + sovQot_GetStaticInfo(uint64(l))
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
		n += 1 + sovQot_GetStaticInfo(uint64(*m.RetType))
	}
	l = len(m.RetMsg)
	n += 1 + l + sovQot_GetStaticInfo(uint64(l))
	n += 1 + sovQot_GetStaticInfo(uint64(m.ErrCode))
	if m.S2C != nil {
		l = m.S2C.Size()
		n += 1 + l + sovQot_GetStaticInfo(uint64(l))
	}
	return n
}

func sovQot_GetStaticInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQot_GetStaticInfo(x uint64) (n int) {
	return sovQot_GetStaticInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *C2S) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQot_GetStaticInfo
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
			}
			m.Market = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_GetStaticInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Market |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecType", wireType)
			}
			m.SecType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_GetStaticInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SecType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_GetStaticInfo
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
				return ErrInvalidLengthQot_GetStaticInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityList = append(m.SecurityList, &qotcommon.Security{})
			if err := m.SecurityList[len(m.SecurityList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQot_GetStaticInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
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
func (m *S2C) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQot_GetStaticInfo
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
				return fmt.Errorf("proto: wrong wireType = %d for field StaticInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_GetStaticInfo
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
				return ErrInvalidLengthQot_GetStaticInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StaticInfoList = append(m.StaticInfoList, &qotcommon.SecurityStaticInfo{})
			if err := m.StaticInfoList[len(m.StaticInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQot_GetStaticInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
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
func (m *Request) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQot_GetStaticInfo
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
					return ErrIntOverflowQot_GetStaticInfo
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
				return ErrInvalidLengthQot_GetStaticInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
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
			skippy, err := skipQot_GetStaticInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
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
				return ErrIntOverflowQot_GetStaticInfo
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
					return ErrIntOverflowQot_GetStaticInfo
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
					return ErrIntOverflowQot_GetStaticInfo
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
				return ErrInvalidLengthQot_GetStaticInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
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
					return ErrIntOverflowQot_GetStaticInfo
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
					return ErrIntOverflowQot_GetStaticInfo
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
				return ErrInvalidLengthQot_GetStaticInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
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
			skippy, err := skipQot_GetStaticInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetStaticInfo
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
func skipQot_GetStaticInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQot_GetStaticInfo
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
					return 0, ErrIntOverflowQot_GetStaticInfo
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
					return 0, ErrIntOverflowQot_GetStaticInfo
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
				return 0, ErrInvalidLengthQot_GetStaticInfo
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQot_GetStaticInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQot_GetStaticInfo
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
				next, err := skipQot_GetStaticInfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQot_GetStaticInfo
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
	ErrInvalidLengthQot_GetStaticInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQot_GetStaticInfo   = fmt.Errorf("proto: integer overflow")
)
