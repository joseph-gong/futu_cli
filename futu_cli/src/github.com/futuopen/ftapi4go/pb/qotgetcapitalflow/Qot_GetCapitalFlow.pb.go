// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Qot_GetCapitalFlow.proto

package qotgetcapitalflow

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

type C2S struct {
	Security *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_b21fa699e235379e, []int{0}
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

func (m *C2S) GetSecurity() *qotcommon.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

type CapitalFlowItem struct {
	InFlow    float64 `protobuf:"fixed64,1,req,name=inFlow" json:"inFlow"`
	Time      string  `protobuf:"bytes,2,opt,name=time" json:"time"`
	Timestamp float64 `protobuf:"fixed64,3,opt,name=timestamp" json:"timestamp"`
}

func (m *CapitalFlowItem) Reset()         { *m = CapitalFlowItem{} }
func (m *CapitalFlowItem) String() string { return proto.CompactTextString(m) }
func (*CapitalFlowItem) ProtoMessage()    {}
func (*CapitalFlowItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b21fa699e235379e, []int{1}
}
func (m *CapitalFlowItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CapitalFlowItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CapitalFlowItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CapitalFlowItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CapitalFlowItem.Merge(m, src)
}
func (m *CapitalFlowItem) XXX_Size() int {
	return m.Size()
}
func (m *CapitalFlowItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CapitalFlowItem.DiscardUnknown(m)
}

var xxx_messageInfo_CapitalFlowItem proto.InternalMessageInfo

func (m *CapitalFlowItem) GetInFlow() float64 {
	if m != nil {
		return m.InFlow
	}
	return 0
}

func (m *CapitalFlowItem) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *CapitalFlowItem) GetTimestamp() float64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type S2C struct {
	FlowItemList       []*CapitalFlowItem `protobuf:"bytes,1,rep,name=flowItemList" json:"flowItemList,omitempty"`
	LastValidTime      string             `protobuf:"bytes,2,opt,name=lastValidTime" json:"lastValidTime"`
	LastValidTimestamp float64            `protobuf:"fixed64,3,opt,name=lastValidTimestamp" json:"lastValidTimestamp"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_b21fa699e235379e, []int{2}
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

func (m *S2C) GetFlowItemList() []*CapitalFlowItem {
	if m != nil {
		return m.FlowItemList
	}
	return nil
}

func (m *S2C) GetLastValidTime() string {
	if m != nil {
		return m.LastValidTime
	}
	return ""
}

func (m *S2C) GetLastValidTimestamp() float64 {
	if m != nil {
		return m.LastValidTimestamp
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
	return fileDescriptor_b21fa699e235379e, []int{3}
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
	return fileDescriptor_b21fa699e235379e, []int{4}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetCapitalFlow.C2S")
	proto.RegisterType((*CapitalFlowItem)(nil), "Qot_GetCapitalFlow.CapitalFlowItem")
	proto.RegisterType((*S2C)(nil), "Qot_GetCapitalFlow.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetCapitalFlow.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetCapitalFlow.Response")
}

func init() { proto.RegisterFile("Qot_GetCapitalFlow.proto", fileDescriptor_b21fa699e235379e) }

var fileDescriptor_b21fa699e235379e = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x77, 0x9a, 0xad, 0x6d, 0x5f, 0x2b, 0xca, 0x28, 0x18, 0x8a, 0xc4, 0x10, 0x2f, 0xab,
	0x60, 0x76, 0x1d, 0x16, 0x04, 0x8f, 0x0d, 0x58, 0x04, 0x3d, 0x38, 0x29, 0x1e, 0xbc, 0x68, 0x36,
	0x9d, 0x8d, 0x03, 0x99, 0xcc, 0x34, 0xf3, 0x82, 0xf4, 0x5b, 0x08, 0x7e, 0x11, 0x3f, 0x46, 0x8f,
	0x3d, 0x7a, 0x12, 0xd9, 0xfd, 0x22, 0x32, 0x69, 0x5a, 0x37, 0xdd, 0xf5, 0x14, 0xf2, 0xff, 0xfd,
	0x5f, 0xde, 0x8f, 0x17, 0xf0, 0x3f, 0x68, 0xfc, 0x7c, 0x2c, 0x30, 0xc9, 0x8c, 0xc4, 0xac, 0x7c,
	0x53, 0xea, 0x6f, 0xb1, 0xa9, 0x35, 0x6a, 0x4a, 0xd7, 0xc9, 0xe1, 0x41, 0xa2, 0x95, 0xd2, 0xd5,
	0x55, 0xe3, 0xf0, 0xbe, 0x6b, 0xac, 0x26, 0xd1, 0x2b, 0xf0, 0x12, 0x96, 0xd2, 0x09, 0xec, 0x5a,
	0x91, 0x37, 0xb5, 0xc4, 0x73, 0x9f, 0x84, 0x5b, 0xa3, 0x7d, 0xf6, 0x30, 0x5e, 0xe9, 0xa6, 0x1d,
	0xe3, 0x37, 0xad, 0x48, 0xc1, 0xbd, 0x95, 0x3d, 0x6f, 0x51, 0x28, 0xfa, 0x18, 0xee, 0xc8, 0xca,
	0xbd, 0xb5, 0x9f, 0x20, 0x47, 0xc3, 0x8b, 0xdf, 0x4f, 0x06, 0xbc, 0xcb, 0xa8, 0x0f, 0x43, 0x94,
	0x4a, 0xf8, 0x5b, 0x21, 0x19, 0xed, 0x75, 0xac, 0x4d, 0x68, 0x04, 0x7b, 0xee, 0x69, 0x31, 0x53,
	0xc6, 0xf7, 0x42, 0x72, 0x33, 0xfa, 0x2f, 0x8e, 0x7e, 0x12, 0xf0, 0x52, 0x96, 0xd0, 0x63, 0x38,
	0x98, 0x77, 0xfb, 0xde, 0x49, 0x8b, 0x3e, 0x09, 0xbd, 0xd1, 0x3e, 0x7b, 0x1a, 0x6f, 0x38, 0xca,
	0x2d, 0x3d, 0xde, 0x1b, 0xa4, 0xcf, 0xe1, 0x6e, 0x99, 0x59, 0xfc, 0x98, 0x95, 0xf2, 0xf4, 0xe4,
	0xb6, 0x57, 0x1f, 0xd1, 0x29, 0xd0, 0x5e, 0xb0, 0x6e, 0xba, 0x81, 0x47, 0x53, 0xd8, 0xe1, 0xe2,
	0xac, 0x11, 0x16, 0xe9, 0x33, 0xf0, 0x72, 0x66, 0xbb, 0xcb, 0x3e, 0xda, 0x28, 0xcb, 0x52, 0xee,
	0x3a, 0xd1, 0x0f, 0x02, 0xbb, 0x5c, 0x58, 0xa3, 0x2b, 0x2b, 0x68, 0x00, 0x3b, 0xb5, 0xc0, 0x93,
	0x73, 0x23, 0xda, 0xd9, 0xed, 0xd7, 0xc3, 0x17, 0xd3, 0xc9, 0x84, 0x5f, 0x87, 0xee, 0xe2, 0xb5,
	0xc0, 0xf7, 0xb6, 0xe8, 0xd9, 0x77, 0x99, 0x9b, 0x16, 0x75, 0x9d, 0xe8, 0x53, 0xd1, 0xba, 0x6e,
	0x77, 0xf8, 0x3a, 0x74, 0x56, 0x96, 0xe5, 0xfe, 0x30, 0x24, 0xff, 0xb3, 0x4a, 0x59, 0xc2, 0x5d,
	0xe7, 0xe8, 0xcb, 0xc5, 0x22, 0x20, 0x97, 0x8b, 0x80, 0xfc, 0x59, 0x04, 0xe4, 0xfb, 0x32, 0x18,
	0x5c, 0x2e, 0x83, 0xc1, 0xaf, 0x65, 0x30, 0x80, 0x07, 0xb9, 0x56, 0xf1, 0xbc, 0xc1, 0x26, 0xd6,
	0x46, 0x54, 0x99, 0x91, 0xb1, 0x99, 0x7d, 0x7a, 0x59, 0x48, 0xfc, 0xda, 0xcc, 0xe2, 0x5c, 0xab,
	0xb1, 0x63, 0x0e, 0x8d, 0xe7, 0x98, 0x19, 0x39, 0x2d, 0xf4, 0xd8, 0xcc, 0xc6, 0x67, 0x1a, 0x0b,
	0x81, 0xf9, 0xd5, 0x2a, 0xf7, 0x5f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x38, 0x1a, 0xc8, 0xc1,
	0xd7, 0x02, 0x00, 0x00,
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
	if m.Security == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("security")
	} else {
		{
			size, err := m.Security.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CapitalFlowItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CapitalFlowItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CapitalFlowItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Timestamp))))
	i--
	dAtA[i] = 0x19
	i -= len(m.Time)
	copy(dAtA[i:], m.Time)
	i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(len(m.Time)))
	i--
	dAtA[i] = 0x12
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.InFlow))))
	i--
	dAtA[i] = 0x9
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
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.LastValidTimestamp))))
	i--
	dAtA[i] = 0x19
	i -= len(m.LastValidTime)
	copy(dAtA[i:], m.LastValidTime)
	i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(len(m.LastValidTime)))
	i--
	dAtA[i] = 0x12
	if len(m.FlowItemList) > 0 {
		for iNdEx := len(m.FlowItemList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FlowItemList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(size))
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
			i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(size))
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
			i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(m.ErrCode))
	i--
	dAtA[i] = 0x18
	i -= len(m.RetMsg)
	copy(dAtA[i:], m.RetMsg)
	i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(len(m.RetMsg)))
	i--
	dAtA[i] = 0x12
	if m.RetType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	} else {
		i = encodeVarintQot_GetCapitalFlow(dAtA, i, uint64(*m.RetType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQot_GetCapitalFlow(dAtA []byte, offset int, v uint64) int {
	offset -= sovQot_GetCapitalFlow(v)
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
	if m.Security != nil {
		l = m.Security.Size()
		n += 1 + l + sovQot_GetCapitalFlow(uint64(l))
	}
	return n
}

func (m *CapitalFlowItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 9
	l = len(m.Time)
	n += 1 + l + sovQot_GetCapitalFlow(uint64(l))
	n += 9
	return n
}

func (m *S2C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FlowItemList) > 0 {
		for _, e := range m.FlowItemList {
			l = e.Size()
			n += 1 + l + sovQot_GetCapitalFlow(uint64(l))
		}
	}
	l = len(m.LastValidTime)
	n += 1 + l + sovQot_GetCapitalFlow(uint64(l))
	n += 9
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
		n += 1 + l + sovQot_GetCapitalFlow(uint64(l))
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
		n += 1 + sovQot_GetCapitalFlow(uint64(*m.RetType))
	}
	l = len(m.RetMsg)
	n += 1 + l + sovQot_GetCapitalFlow(uint64(l))
	n += 1 + sovQot_GetCapitalFlow(uint64(m.ErrCode))
	if m.S2C != nil {
		l = m.S2C.Size()
		n += 1 + l + sovQot_GetCapitalFlow(uint64(l))
	}
	return n
}

func sovQot_GetCapitalFlow(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQot_GetCapitalFlow(x uint64) (n int) {
	return sovQot_GetCapitalFlow(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return ErrIntOverflowQot_GetCapitalFlow
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
				return fmt.Errorf("proto: wrong wireType = %d for field Security", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_GetCapitalFlow
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
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
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
		default:
			iNdEx = preIndex
			skippy, err := skipQot_GetCapitalFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
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
func (m *CapitalFlowItem) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQot_GetCapitalFlow
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
			return fmt.Errorf("proto: CapitalFlowItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CapitalFlowItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field InFlow", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.InFlow = float64(math.Float64frombits(v))
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_GetCapitalFlow
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
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Timestamp = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipQot_GetCapitalFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("inFlow")
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
				return ErrIntOverflowQot_GetCapitalFlow
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
				return fmt.Errorf("proto: wrong wireType = %d for field FlowItemList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_GetCapitalFlow
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
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FlowItemList = append(m.FlowItemList, &CapitalFlowItem{})
			if err := m.FlowItemList[len(m.FlowItemList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastValidTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_GetCapitalFlow
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
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastValidTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastValidTimestamp", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.LastValidTimestamp = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipQot_GetCapitalFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
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
				return ErrIntOverflowQot_GetCapitalFlow
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
					return ErrIntOverflowQot_GetCapitalFlow
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
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
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
			skippy, err := skipQot_GetCapitalFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
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
				return ErrIntOverflowQot_GetCapitalFlow
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
					return ErrIntOverflowQot_GetCapitalFlow
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
					return ErrIntOverflowQot_GetCapitalFlow
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
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
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
					return ErrIntOverflowQot_GetCapitalFlow
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
					return ErrIntOverflowQot_GetCapitalFlow
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
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
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
			skippy, err := skipQot_GetCapitalFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_GetCapitalFlow
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
func skipQot_GetCapitalFlow(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQot_GetCapitalFlow
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
					return 0, ErrIntOverflowQot_GetCapitalFlow
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
					return 0, ErrIntOverflowQot_GetCapitalFlow
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
				return 0, ErrInvalidLengthQot_GetCapitalFlow
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQot_GetCapitalFlow
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQot_GetCapitalFlow
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
				next, err := skipQot_GetCapitalFlow(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQot_GetCapitalFlow
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
	ErrInvalidLengthQot_GetCapitalFlow = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQot_GetCapitalFlow   = fmt.Errorf("proto: integer overflow")
)
